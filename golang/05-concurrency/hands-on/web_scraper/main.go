package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FetchResult struct {
	URL  string
	Size int
	Err  error
}

// Simulates an HTTP request
func fetchURL(url string, wg *sync.WaitGroup, results chan<- FetchResult) {
	defer wg.Done()

	// Simulate network latency
	delay := rand.Intn(500)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// Simulate error
	if delay > 400 {
		results <- FetchResult{URL: url, Err: fmt.Errorf("timeout")}
		return
	}

	results <- FetchResult{URL: url, Size: len(url) * 10, Err: nil}
}

func main() {
	urls := []string{
		"http://google.com",
		"http://golang.org",
		"http://reddit.com",
		"http://github.com",
		"http://stackoverflow.com",
	}

	var wg sync.WaitGroup
	results := make(chan FetchResult, len(urls))

	fmt.Println("🚀 Starting concurrent scrape...")
	start := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg, results)
	}

	// Wait in a separate goroutine so we can close string channel without blocking main
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results as they come in
	for res := range results {
		if res.Err != nil {
			fmt.Printf("❌ Failed: %s (%v)\n", res.URL, res.Err)
		} else {
			fmt.Printf("✅ Fetched: %s (Size: %d)\n", res.URL, res.Size)
		}
	}

	fmt.Printf("Done in %v\n", time.Since(start))
}
