package main

import (
	"fmt"
	"math"
)

// Global variable (must use var)
var Version = "1.0.0"

func main() {
	// --- 1. Variable Declaration Styles ---

	// Style A: Verbose (good when you don't have a value yet)
	var age int
	age = 30

	// Style B: Type inference (rarely used for simple types)
	var name = "Gopher"

	// Style C: Short declaration operator (Standard for local vars)
	// Only works inside functions!
	score := 95.5

	fmt.Printf("Name: %s, Age: %d, Score: %.1f\n", name, age, score)

	// --- 2. Constants ---
	const Pi = 3.14159
	// Pi = 3.14 // Error: cannot assign to Pi

	// --- 3. Basic Types ---
	var isCool bool = true            // boolean
	var initial rune = 'G'            // rune (alias for int32, represents a Unicode code point)
	var bit byte = 255                // byte (alias for uint8)
	var complexNum complex64 = 1 + 2i // complex number

	fmt.Printf("Types: %T %T %T %T\n", isCool, initial, bit, complexNum)
	fmt.Println("Rune 'G' value:", initial) // Prints 71 (ASCII/Unicode value)

	// --- 4. Type Conversion (Casting) ---
	// Go implies NOTHING. You must cast explicitly.
	var x int = 42
	var y float64 = float64(x) // Must convert int to float explicitly
	var z uint = uint(y)

	fmt.Println("Conversions:", x, y, z)
	fmt.Println("Max Int64:", math.MaxInt64)
}
