package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"
	content := []byte("Hello, File System!\nThis is line 2.")

	// --- 1. Writing to a file ---
	// Permission 0644: Owner Read/Write, Group Read, Others Read
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully.")

	// --- 2. Reading from a file ---
	readContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read content:\n%s\n", string(readContent))

	// --- 3. Checking if file exists / Stat ---
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error checking file:", err)
		}
	} else {
		fmt.Printf("File Info: Name=%s, Size=%d bytes, Mode=%s\n",
			info.Name(), info.Size(), info.Mode())
	}

	// Cleanup
	// os.Remove(filename)
}
