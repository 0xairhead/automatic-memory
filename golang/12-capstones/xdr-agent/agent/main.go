package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Config
const (
	ServerURL  = "http://localhost:9090/audit"
	AgentID    = "agent-macbook-01"
	NumWorkers = 3
)

type Alert struct {
	AgentID   string `json:"agent_id"`
	EventType string `json:"event_type"`
	Details   string `json:"details"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	fmt.Println("🛡️  XDR Agent Starting...")

	// 1. Setup Channels
	alertQueue := make(chan Alert, 100)
	var wg sync.WaitGroup

	// 2. Start Worker Pool (Network Senders)
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			senderWorker(id, alertQueue)
		}(i)
	}

	// 3. Start Monitors
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go fileMonitor(ctx, &wg, alertQueue) // Monitor 1
	wg.Add(1)
	go processMonitor(ctx, &wg, alertQueue) // Monitor 2

	// 4. Wait for Shutdown Signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n🛑 Shutdown signal received. Stopping monitors...")
	cancel() // Stop monitors

	// Close queue after monitors are done (we need to wait for monitors aka "producers" to finish first ideally)
	// But mostly monitors respect context.
	// For simplicity, we'll give them a moment to cleanup or just close queue if we know they are done.
	// Correct pattern: Wait for monitors to return, then close(alertQueue), then wait for workers.
	// Implementing simple timeout for demo:
	time.Sleep(1 * time.Second)
	close(alertQueue)

	wg.Wait()
	fmt.Println("Agent exited gracefully.")
}

// --- Monitors (Producers) ---

func fileMonitor(ctx context.Context, wg *sync.WaitGroup, alerts chan<- Alert) {
	defer wg.Done()
	fmt.Println("Checking File Integrity...")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Simulate finding a file change
			if rand.Float32() < 0.3 {
				alerts <- Alert{
					AgentID:   AgentID,
					EventType: "FILE_MODIFIED",
					Details:   "/etc/passwd accessed by unknown user",
					Timestamp: time.Now().Unix(),
				}
			}
		}
	}
}

func processMonitor(ctx context.Context, wg *sync.WaitGroup, alerts chan<- Alert) {
	defer wg.Done()
	fmt.Println("Monitoring Processes...")
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Simulate a suspicious process
			if rand.Float32() < 0.2 {
				alerts <- Alert{
					AgentID:   AgentID,
					EventType: "UNAUTHORIZED_ACCESS",
					Details:   "Process 'miner_x' started (PID: 9999)",
					Timestamp: time.Now().Unix(),
				}
			}
		}
	}
}

// --- Workers (Consumers) ---

func senderWorker(id int, alerts <-chan Alert) {
	for alert := range alerts {
		sendAlert(alert)
	}
	fmt.Printf("Worker %d stopped.\n", id)
}

func sendAlert(alert Alert) {
	data, _ := json.Marshal(alert)
	resp, err := http.Post(ServerURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("⚠️  Failed to send alert: %v\n", err)
		return
	}
	defer resp.Body.Close()
	// fmt.Printf("Sent alert: %s\n", alert.EventType)
}
