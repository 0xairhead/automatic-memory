package main

import (
	"fmt"
)

// Define a custom error struct
type DetailedError struct {
	Code    int
	Message string
	Path    string
}

// Implement the strict Error interface: Error() string
func (e *DetailedError) Error() string {
	return fmt.Sprintf("Error %d: %s (at %s)", e.Code, e.Message, e.Path)
}

func readFile(path string) error {
	// Simulate missing file
	return &DetailedError{
		Code:    404,
		Message: "File not found",
		Path:    path,
	}
}

func main() {
	err := readFile("/etc/shadow")
	if err != nil {
		fmt.Println("Standard print:", err)

		// Type Assertion: Check if err is *DetailedError
		if e, ok := err.(*DetailedError); ok {
			fmt.Println("Custom field access - Code:", e.Code)
			fmt.Println("Custom field access - Path:", e.Path)
		}
	}
}
