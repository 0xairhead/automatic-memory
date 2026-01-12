package main

import "fmt"

// --- 1. Without Generics (Repeated Code) ---
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// --- 2. With Generics (DRY) ---

// Number constraint interface
type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m.
// K is comparable (required for map keys)
// V is Number (int64 or float64)
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// --- 3. Generic Types ---
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	lastIdx := len(s.items) - 1
	val := s.items[lastIdx]
	s.items = s.items[:lastIdx] // Remove
	return val
}

func main() {
	// Maps
	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	// Generic Stack
	s := Stack[string]{}
	s.Push("World")
	s.Push("Hello")
	fmt.Println("Stack Pop:", s.Pop())
	fmt.Println("Stack Pop:", s.Pop())
}
