package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// Struct Tags (Metadata for JSON/YAML)
type Product struct {
	Name  string  `json:"product_name"`
	Price float64 `json:"price_usd"`
	// "omitempty": Don't include field in JSON if zero-value (0, "", false)
	Stock int `json:"stock,omitempty"`
	// "-": Ignore this field in JSON
	SecretCode string `json:"-"`
}

func main() {
	// --- 1. Creating Structs ---
	u1 := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		IsActive:  true,
	}
	// Missing fields get zero values
	u2 := User{FirstName: "Jane"}

	fmt.Printf("User 1: %+v\n", u1) // %+v prints field names
	fmt.Printf("User 2: %+v\n", u2) // ID=0, IsActive=false

	// --- 2. Anonymous / Embedded Structs (Composition) ---
	// Go doesn't have inheritance. We use embedding.
	type Admin struct {
		User  // Embeds all fields of User
		Level string
	}

	admin := Admin{
		User:  User{FirstName: "Root", IsActive: true},
		Level: "Super",
	}

	// Access embedded fields directly!
	fmt.Println("Admin Name:", admin.FirstName) // admin.User.FirstName also works

	// --- 3. Struct Tags & JSON ---
	p := Product{
		Name:       "Monitor",
		Price:      199.99,
		Stock:      0,       // Zero value, will be omitted
		SecretCode: "12345", // Will be ignored
	}

	// Marshaling (Struct -> JSON)
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Returns byte slice, convert to string
	fmt.Println("JSON:", string(jsonData))
}
