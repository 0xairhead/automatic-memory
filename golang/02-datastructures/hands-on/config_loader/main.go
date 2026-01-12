package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port  int    `json:"server_port"`
	DBUrl string `json:"database_url"`
	Debug bool   `json:"debug_mode"`
}

func main() {
	// Hands-on: Config Loader
	// 1. Read file
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// 2. Parse JSON
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// 3. Use Config
	fmt.Println("--- App Configuration ---")
	fmt.Printf("Listening on port: %d\n", cfg.Port)
	fmt.Printf("Connecting to DB:  %s\n", cfg.DBUrl)
	if cfg.Debug {
		fmt.Println("Warning: Debug mode is ON")
	}
}
