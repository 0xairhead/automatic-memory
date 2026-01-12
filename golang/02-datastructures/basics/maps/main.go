package main

import "fmt"

func main() {
	// --- 1. Creating Maps ---
	// map[KeyType]ValueType

	// Init style A: Literal
	scores := map[string]int{
		"Alice": 95,
		"Bob":   82,
	}

	// Init style B: make()
	// Always use make if not initializing immediately!
	// otherwise it's nil and assigning panics.
	cache := make(map[string]string)
	cache["home"] = "index.html"
	cache["about"] = "about.html"

	fmt.Println("Scores:", scores)
	fmt.Println("Cache:", cache)

	// --- 2. Accessing & Checking Keys ---
	aliceScore := scores["Alice"]
	fmt.Println("Alice:", aliceScore)

	// What if key doesn't exist?
	// Returns "zero value" for the type (0 for int)
	charlieScore := scores["Charlie"]
	fmt.Println("Charlie (not in map):", charlieScore)

	// How to distinguish 0 vs Missing? uses "ok" idiom
	val, ok := scores["Charlie"]
	if !ok {
		fmt.Println("Charlie does not exist in map!")
	} else {
		fmt.Println("Charlie's score:", val)
	}

	// --- 3. Deleting ---
	delete(scores, "Bob")
	fmt.Println("After deleting Bob:", scores)

	// --- 4. Iterating (Random Order!) ---
	scores["Dave"] = 100
	scores["Eve"] = 42

	fmt.Println("Iterating:")
	for key, value := range scores {
		fmt.Printf("  %s -> %d\n", key, value)
	}
}
