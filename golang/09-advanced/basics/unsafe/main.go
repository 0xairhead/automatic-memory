package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// --- 1. Size of types ---
	var x int64 = 10
	fmt.Println("Size of x (int64):", unsafe.Sizeof(x), "bytes")

	// --- 2. Reinterpreting Memory (Dangerous!) ---
	// Convert float64 bits directly to uint64 (like C++ reinterpret_cast)
	f := 3.14

	// Step 1: Get pointer to float
	ptrF := &f

	// Step 2: Convert to universally unsafe pointer
	unsafePtr := unsafe.Pointer(ptrF)

	// Step 3: Cast to pointer of target type
	ptrUint := (*uint64)(unsafePtr)

	// Step 4: Dereference
	bits := *ptrUint

	fmt.Printf("Float: %f\n", f)
	fmt.Printf("Bits (uint64): %b\n", bits)

	// Note: We modified nothing, just viewed the same bytes differently.
	// Modifying via unsafe pointer is possible but easy to segfault.
}
