package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Simple CLI Calculator
	// Usage: go run main.go 5 + 3

	// os.Args holds command line arguments
	// os.Args[0] is the program name
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <num1> <operator> <num2>")
		return
	}

	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	// ParseFloat string -> float64
	n1, err1 := strconv.ParseFloat(num1Str, 64)
	n2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers.")
		return
	}

	var result float64

	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		if n2 == 0 {
			fmt.Println("Error: Cannot divide by zero.")
			return
		}
		result = n1 / n2
	default:
		fmt.Println("Error: Invalid operator. Use +, -, *, /")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", n1, operator, n2, result)
}
