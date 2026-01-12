package main

import "fmt"

func main() {
	// --- Panic & Recover ---
	// DO NOT use panic for normal error handling.
	// Use it for unrecoverable state (e.g. init failure)
	// or in library code to catch internal bugs.

	fmt.Println("Start")

	// Call a function that crashes
	safeCall()

	fmt.Println("End (This runs because we recovered!)")
}

func safeCall() {
	// DEFER is LIFO.
	// RECOVER must be called inside a deferred function.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("Something terrible happened!")
	// fmt.Println("Unreachable")
}
