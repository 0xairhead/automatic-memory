package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// --- 1. Context with Timeout ---
	// "Stop this operation if it takes longer than 2s"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Release resources

	// Simulate a call
	result, err := slowOperation(ctx)
	if err != nil {
		fmt.Println("Main: Operation failed:", err)
	} else {
		fmt.Println("Main: Success:", result)
	}
}

func slowOperation(ctx context.Context) (string, error) {
	// Simulate work using select
	// This pattern allows us to listen to the Done() channel
	select {
	case <-time.After(3 * time.Second): // Takes 3s
		return "Completed!", nil
	case <-ctx.Done(): // Context cancelled/timed out
		return "", ctx.Err() // returns "context deadline exceeded"
	}
}
