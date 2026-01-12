package main

import (
	"fmt"
	"sync"
	"time"
)

// Helper to simulate work
func printMessage(msg string, count int, wg *sync.WaitGroup) {
	// 3. Defer Done() to ensure it runs even if function panics
	defer wg.Done()

	for i := 0; i < count; i++ {
		fmt.Println(msg, i)
		time.Sleep(100 * time.Millisecond) // Simulate slow work
	}
}

func main() {
	// --- 1. Sequential Execution (Slow) ---
	// start := time.Now()
	// printMessage("Sequential", 3, &sync.WaitGroup{}) // Mocking wg for demo
	// fmt.Println("Seq took:", time.Since(start))

	// --- 2. Concurrent Execution via Goroutines ---
	var wg sync.WaitGroup

	fmt.Println("Starting Goroutines...")

	// Launching goroutine 1
	wg.Add(1) // Increment counter
	go printMessage("🚀 Fast", 5, &wg)

	// Launching goroutine 2
	wg.Add(1)
	go printMessage("🐢 Slow", 5, &wg)

	// Launching an anonymous function
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("👻 I am a ghost goroutine!")
	}()

	// --- 3. Wait for completion ---
	// Without this, main() would exit immediately, killing other goroutines!
	fmt.Println("Waiting for work to finish...")
	wg.Wait()
	fmt.Println("All done!")
}
