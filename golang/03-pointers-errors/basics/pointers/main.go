package main

import "fmt"

func main() {
	// --- 1. Basic Pointers ---
	a := 10
	var p *int = &a // & gets the address

	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", p)
	fmt.Printf("Value at address p (*p): %d\n", *p) // Dereferencing

	// Changing value via pointer
	*p = 20
	fmt.Println("New value of a:", a) // 20

	// --- 2. Pass by Value vs Pass by Pointer ---
	val := 100
	modifyValue(val)
	fmt.Println("After modifyValue:", val) // Still 100 (Copy)

	modifyPointer(&val)
	fmt.Println("After modifyPointer:", val) // 200 (Modified)

	// --- 3. Nil Pointers (Danger Zone) ---
	var nilPtr *int // Zero value is nil
	if nilPtr == nil {
		fmt.Println("nilPtr is nil. Safety check passed.")
	}
	// fmt.Println(*nilPtr) // PANIC! runtime error: invalid memory address or nil pointer dereference
}

func modifyValue(n int) {
	n = 0 // Only changes local copy
}

func modifyPointer(n *int) {
	*n = 200 // Changes the value at the address
}
