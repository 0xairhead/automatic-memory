package main

import (
	"encoding/json"
	"fmt"
)

// Define the shape of our data
// Note: Capitalized fields = Exported (Visible to JSON package)
type Weather struct {
	City      string  `json:"city"`
	Temp      float64 `json:"temperature"`
	Condition string  `json:"condition"`
}

func main() {
	// Hands-on: Parse JSON into structs

	// 1. Simulate receiving JSON payload (e.g. from API)
	jsonBlob := `
	{
		"city": "London",
		"temperature": 15.5,
		"condition": "Cloudy",
		"extra_field": "ignored"
	}`

	// 2. Variable to hold the data
	var w Weather

	// 3. Unmarshal (Parse)
	// Requires byte slice
	err := json.Unmarshal([]byte(jsonBlob), &w)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Printf("Parsed Data:\nCity: %s\nTemp: %.1f\nDesc: %s\n",
		w.City, w.Temp, w.Condition)
}
