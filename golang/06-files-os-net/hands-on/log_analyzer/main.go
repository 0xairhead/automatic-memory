package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	logFile := "app.log"
	createDummyLog(logFile)
	defer os.Remove(logFile) // Cleanup

	fmt.Printf("Analyzing %s...\n", logFile)

	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorCount := 0
	lineNum := 0

	fmt.Println("--- Errors Found ---")
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Simple parsing logic: Look for "ERROR"
		if strings.Contains(line, "ERROR") {
			errorCount++
			fmt.Printf("Line %d: %s\n", lineNum, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}

	fmt.Printf("--------------------\nTotal Errors: %d\n", errorCount)
}

func createDummyLog(filename string) {
	content := `[INFO] Server started
[INFO] User logged in
[ERROR] Database timeout
[INFO] Request processed
[ERROR] Payment gateway failed
[WARN] High memory usage
[INFO] User logged out`

	os.WriteFile(filename, []byte(content), 0644)
}
