package main

import "fmt"

// --- 1. Define the Interface ---
// An interface describes *behavior* (methods).
type Speaker interface {
	Speak() string
}

// --- 2. Implement the Interface (Implicitly!) ---

type Dog struct {
	Name string
}

// Dog implements Speaker because it has a Speak method
func (d Dog) Speak() string {
	return "Woof! I am " + d.Name
}

type Cat struct {
	Name string
}

// Cat implements Speaker too
func (c Cat) Speak() string {
	return "Meow. " + c.Name + " here."
}

// --- 3. Use the Interface ---
// This function accepts ANYTHING that knows how to Speak.
func MakeItTalk(s Speaker) {
	fmt.Println("Speaking:", s.Speak())
}

func main() {
	d := Dog{Name: "Buddy"}
	c := Cat{Name: "Whiskers"}

	// Polymorphism in action
	MakeItTalk(d)
	MakeItTalk(c)

	// Interface Slice
	animals := []Speaker{d, c}
	fmt.Printf("We have %d animals.\n", len(animals))
}
