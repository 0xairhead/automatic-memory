package main

import (
	"fmt"
)

// SafeRunner executes a function and ensures it doesn't crash the app
func SafeRunner(task func(), name string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[%s] CRASHED! Reason: %v\n", name, r)
		}
	}()

	fmt.Printf("[%s] Starting...\n", name)
	task()
	fmt.Printf("[%s] Completed successfully.\n", name)
}

func riskyTask() {
	panic("Out of memory")
}

func safeTask() {
	fmt.Println("Doing some work...")
}

func main() {
	// 1. Run a safe task
	SafeRunner(safeTask, "Worker-1")

	fmt.Println("---")

	// 2. Run a risky task
	SafeRunner(riskyTask, "Worker-2")

	fmt.Println("---")
	fmt.Println("Main program continues running!")
}
