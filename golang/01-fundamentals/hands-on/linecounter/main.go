package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Hands-on 2: File Line Counter
	// Usage: go run main.go <filename>

	if len(os.Args) < 2 {
		fmt.Println("Usage: linecounter <filename>")
		return
	}

	fileName := os.Args[1]

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	// defer ensures the file is closed when the function exits
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Loop through each line
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("File '%s' has %d lines.\n", fileName, lineCount)
}
