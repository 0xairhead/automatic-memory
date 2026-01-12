package main

import "fmt"

func main() {
	// 1. Basic call
	res := add(5, 10)
	fmt.Println("Sum:", res)

	// 2. Multiple returns
	quotient, remainder := safeDiv(10, 3)
	fmt.Printf("10/3 = %d (rem %d)\n", quotient, remainder)

	// Blank identifier (_) to ignore values
	q, _ := safeDiv(10, 3)
	fmt.Println("Just quotient:", q)

	// 3. Named returns (use carefully)
	x, y := split(17)
	fmt.Println("Split 17:", x, y) // 7, 10

	// 4. Variadic functions
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Variadic sum:", total)
}

// Basic function
func add(x int, y int) int {
	return x + y
}

// Multiple return values (Power feature!)
func safeDiv(a, b int) (int, int) {
	if b == 0 {
		return 0, 0
	}
	return a / b, a % b
}

// Named return values
// "x" and "y" are initialized to 0.
// A "naked" return returns them automatically.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Returns x, y
}

// Variadic function (takes any number of ints)
func sumAll(nums ...int) int {
	total := 0
	// nums is a slice []int
	for _, num := range nums {
		total += num
	}
	return total
}
