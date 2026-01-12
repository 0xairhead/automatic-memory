package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// --- 1. Environment Variables ---
	// Set one simply for demo
	os.Setenv("APP_PORT", "9090")

	// Get existing
	user := os.Getenv("USER")
	port := os.Getenv("APP_PORT")
	missing := os.Getenv("NOT_REAL")

	fmt.Printf("User: %s, Port: %s, Missing: '%s'\n", user, port, missing)

	// --- 2. Signal Handling (Graceful Shutdown) ---
	// Create a channel to listen for signals
	sigs := make(chan os.Signal, 1)

	// Notify this channel on SIGINT (Ctrl+C) or SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	// Background goroutine to wait for signal
	go func() {
		sig := <-sigs // Block until signal received
		fmt.Println("\nReceived signal:", sig)
		fmt.Println("Cleaning up resources...")
		time.Sleep(1 * time.Second) // Simulate cleanup
		fmt.Println("Cleanup done.")
		done <- true
	}()

	fmt.Println("App running... (Press Ctrl+C to stop)")
	<-done // Block main until cleanup is done
	fmt.Println("Exiting.")
}
