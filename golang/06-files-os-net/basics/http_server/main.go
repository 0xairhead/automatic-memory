package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// --- Server Handler ---
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Server! Time: %s", time.Now().Format(time.RFC3339))
}

func main() {
	// 1. Defining Routes
	// http.HandleFunc registers a handler to the DefaultServeMux
	http.HandleFunc("/", helloHandler)

	// 2. Starting Server in a Goroutine (so we can act as client below)
	go func() {
		fmt.Println("Server starting on :8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// Give server a moment to start
	time.Sleep(1 * time.Second)

	// --- 3. HTTP Client ---
	fmt.Println("\nClient: Making request...")
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Client Error:", err)
		return
	}
	defer resp.Body.Close() // ALWAYS close body

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Client: Received Status: %s\n", resp.Status)
	fmt.Printf("Client: Received Body: %s\n", string(body))
}
