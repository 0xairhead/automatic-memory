package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Vulnerable function
func readFileUnsafe(baseDir, userInput string) {
	// DANGER: Simple concatenation allows "../"
	fullPath := baseDir + "/" + userInput
	fmt.Println("Unsafe Path:", fullPath)
}

// Secure function
func readFileSafe(baseDir, userInput string) error {
	// 1. Clean the path (removes ../)
	cleanInput := filepath.Clean(userInput)

	// 2. Join with base
	fullPath := filepath.Join(baseDir, cleanInput)

	// 3. Verify it's still inside baseDir!
	// Prevents: "base/" + "../etc/passwd" -> "/etc/passwd"
	// But Join handles some of this. The critical check is HasPrefix.

	absBase, _ := filepath.Abs(baseDir)
	absPath, _ := filepath.Abs(fullPath)

	if !strings.HasPrefix(absPath, absBase) {
		return fmt.Errorf("security alert: path traversal attempt detected! (%s)", userInput)
	}

	fmt.Println("Safe Path:", absPath)
	return nil
}

func main() {
	base := "/var/www/html"
	attackerInput := "../../etc/shadow"

	fmt.Println("--- Unsafe ---")
	readFileUnsafe(base, attackerInput)

	fmt.Println("\n--- Safe ---")
	err := readFileSafe(base, attackerInput)
	if err != nil {
		fmt.Println("BLOCKED:", err)
	}

	validInput := "images/logo.png"
	_ = readFileSafe(base, validInput)
}
