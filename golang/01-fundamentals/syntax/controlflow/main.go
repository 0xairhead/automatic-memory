package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// --- 1. If / Else ---
	x := 10
	if x > 5 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is small")
	}

	// Unique to Go: Statement initialization
	// "err" scope is limited to the if/else block
	if err := someCheck(); err != nil {
		fmt.Println("Error:", err)
	}

	// --- 2. For Loops (The ONLY loop in Go) ---
	// Basic loop
	fmt.Print("Counting: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// While-style loop
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("Sum is:", sum)

	// Infinite loop (commented out to prevent hanging)
	// for {
	//     break
	// }

	// --- 3. Switch ---
	// No "break" needed! It's automatic.
	// Use "fallthrough" if you strictly need it.
	fmt.Print("OS: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Tagless switch (cleaner if-else chain)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func someCheck() error {
	return nil // Simulate no error
}
