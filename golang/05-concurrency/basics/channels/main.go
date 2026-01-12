package main

import (
	"fmt"
	"time"
)

func main() {
	// --- 1. Unbuffered Channel (Synchronous) ---
	// Requires receiver to be ready when sending
	ch := make(chan string)

	go func() {
		fmt.Println("Sender: Sending Ping...")
		ch <- "Ping" // Blocks here until someone receives!
		fmt.Println("Sender: Sent Ping!")
	}()

	time.Sleep(1 * time.Second) // Simulate delay
	fmt.Println("Receiver: Waiting...")
	msg := <-ch // Blocks here until someone sends!
	fmt.Println("Receiver: Got", msg)

	// --- 2. Buffered Channel (Asynchronous capacity) ---
	// Doesn't block send until full.
	bufCh := make(chan int, 2) // Capacity 2

	bufCh <- 1
	bufCh <- 2
	// bufCh <- 3 // Would block (deadlock in single thread!) if uncommented

	fmt.Println("Buffered Read 1:", <-bufCh)
	fmt.Println("Buffered Read 2:", <-bufCh)

	// --- 3. Iterating over channels ---
	queue := make(chan string, 5)
	queue <- "Job 1"
	queue <- "Job 2"
	close(queue) // CLOSING is crucial for range loops!

	for item := range queue {
		fmt.Println("Processing:", item)
	}
}
