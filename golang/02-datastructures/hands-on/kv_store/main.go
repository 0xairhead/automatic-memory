package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Simple In-Memory DB
	store := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("SimpleKV v0.1")
	fmt.Println("Commands: SET <key> <value> | GET <key> | DEL <key> | EXIT")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])

		switch command {
		case "SET":
			if len(parts) < 3 {
				fmt.Println("Usage: SET <key> <value>")
				continue
			}
			key := parts[1]
			value := strings.Join(parts[2:], " ")
			store[key] = value
			fmt.Println("OK")

		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET <key>")
				continue
			}
			key := parts[1]
			val, ok := store[key]
			if !ok {
				fmt.Println("(nil)")
			} else {
				fmt.Println(val)
			}

		case "DEL":
			if len(parts) != 2 {
				fmt.Println("Usage: DEL <key>")
				continue
			}
			key := parts[1]
			delete(store, key)
			fmt.Println("OK")

		case "EXIT":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
