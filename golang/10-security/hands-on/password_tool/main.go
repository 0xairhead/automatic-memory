package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// GenerateFromPassword does both salting and hashing
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // Cost 14 is secure but slow
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go hash <password>")
		fmt.Println("       go run main.go check <hash> <password>")
		return
	}

	command := os.Args[1]

	switch command {
	case "hash":
		password := os.Args[2]
		hash, _ := HashPassword(password)
		fmt.Println(hash)

	case "check":
		hash := os.Args[2]
		password := os.Args[3]
		match := CheckPasswordHash(password, hash)
		if match {
			fmt.Println("✅ Password matches!")
		} else {
			fmt.Println("❌ Invalid password.")
		}

	default:
		fmt.Println("Unknown command")
	}
}
