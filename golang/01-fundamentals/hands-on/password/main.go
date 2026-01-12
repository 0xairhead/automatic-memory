package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Hands-on 3: Password Strength Checker
	if len(os.Args) != 2 {
		fmt.Println("Usage: password <password>")
		return
	}

	password := os.Args[1]
	length := len(password)
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	score := 0
	if length >= 8 {
		score++
	}
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSpecial {
		score++
	}

	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Score: %d/5\n", score)

	switch score {
	case 5:
		fmt.Println("Strength: Strong 💪")
	case 3, 4:
		fmt.Println("Strength: Moderate 😐")
	default:
		fmt.Println("Strength: Weak 🤢")
	}
}
