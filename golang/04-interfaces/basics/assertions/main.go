package main

import "fmt"

func main() {
	// --- 1. The Empty Interface (any) ---
	// Can hold ANY value.
	var i interface{} = "hello"

	fmt.Printf("Underlying value: %v, Type: %T\n", i, i)

	// --- 2. Type Assertion ---
	// Extract the string out of the interface
	s, ok := i.(string)
	if ok {
		fmt.Println("It is a string:", s)
	}

	// Unsafe assertion (will panic if wrong)
	// hasNum := i.(int) // PANIC!

	// Safe assertion for wrong type
	num, ok := i.(int)
	if !ok {
		fmt.Println("It is NOT an int")
		fmt.Println("Zero value of int:", num) // 0
	}

	// --- 3. Type Switch ---
	describe(42)
	describe("Golang")
	describe(true)
}

func describe(i interface{}) {
	// Switch based on TYPE, not value
	switch v := i.(type) {
	case int:
		fmt.Printf("I got an Integer: %d\n", v)
	case string:
		fmt.Printf("I got a String of len %d\n", len(v))
	default:
		fmt.Printf("I don't know type %T\n", v)
	}
}
