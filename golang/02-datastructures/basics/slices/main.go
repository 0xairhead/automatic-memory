package main

import "fmt"

func main() {
	// --- 1. Arrays (Fixed Size - rarely used directly) ---
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	// arr[3] = 40 // Compile error: out of bounds
	fmt.Println("Array:", arr)

	// --- 2. Slices (Dynamic Size - the standard) ---
	// A slice is a "view" into an underlying array.
	var s []int = []int{1, 2, 3}
	fmt.Printf("Slice: %v, Len: %d, Cap: %d\n", s, len(s), cap(s))

	// Appending
	s = append(s, 4, 5) // Might resize underlying array
	fmt.Printf("After append: %v, Len: %d, Cap: %d\n", s, len(s), cap(s))

	// --- 3. Creating with make() ---
	// make([]type, len, cap)
	// Useful to pre-allocate memory for performance
	users := make([]string, 2, 5) // Length 2, Capacity 5
	users[0] = "Alice"
	users[1] = "Bob"
	// users[2] = "Charlie" // Panic! Index out of range (len is 2)

	users = append(users, "Charlie") // Now len is 3
	fmt.Printf("Users: %v, Cap: %d\n", users, cap(users))

	// --- 4. Slicing Slices ---
	// slice[low:high] (inclusive low, exclusive high)
	numbers := []int{0, 1, 2, 3, 4, 5}
	sub := numbers[1:4] // Index 1, 2, 3
	fmt.Println("Sub-slice [1:4]:", sub)

	// Important: Sub-slices share the same memory!
	sub[0] = 999
	fmt.Println("Original numbers after modifying sub:", numbers)
	// numbers[1] becomes 999
}
