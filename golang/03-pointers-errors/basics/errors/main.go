package main

import (
	"errors"
	"fmt"
	"os"
)

// --- 1. Basic Error ---
// Errors are values! Type is "error", an interface.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Standard checking pattern
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result2, err2 := divide(10, 2)
	if err2 != nil {
		// handle error
	}
	fmt.Println("Result2:", result2)

	// --- 2. Wrapping Errors (Go 1.13+) ---
	if err := openConfig("missing.json"); err != nil {
		// %w wraps the error
		fmt.Printf("Main failed: %v\n", err)

		// Unwrapping
		originalErr := errors.Unwrap(err)
		fmt.Println("Original cause:", originalErr)
	}
}

func openConfig(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// Annotate error with context
		return fmt.Errorf("loading config failed: %w", err)
	}
	return nil
}
