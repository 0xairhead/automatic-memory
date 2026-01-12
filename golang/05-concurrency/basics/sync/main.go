package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	// Lock ensuring only one goroutine enters at a time
	c.mu.Lock()
	defer c.mu.Unlock() // Always defer Unlock!

	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// simulate 1000 concurrent increments
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc() // Safe access
		}()
	}

	wg.Wait()
	fmt.Println("Final Count (Should be 1000):", counter.Value())
}
