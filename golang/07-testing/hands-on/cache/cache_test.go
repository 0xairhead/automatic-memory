package cache

import (
	"sync"
	"testing"
)

func TestCacheOperations(t *testing.T) {
	c := New()

	// 1. Test Set & Get
	c.Set("user", "phoenix")
	val, ok := c.Get("user")
	if !ok || val != "phoenix" {
		t.Errorf("Get(user) = %s, %v; want phoenix, true", val, ok)
	}

	// 2. Test Delete
	c.Delete("user")
	_, ok = c.Get("user")
	if ok {
		t.Error("Get(user) should return false after delete")
	}
}

// TestConcurrency runs many goroutines to check for race conditions
// Use: go test -race
func TestConcurrency(t *testing.T) {
	c := New()
	var wg sync.WaitGroup

	// 100 writers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Set("key", "value")
		}()
	}

	// 100 readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("key")
		}()
	}

	wg.Wait()
}

func BenchmarkCacheSet(b *testing.B) {
	c := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set("key", "value")
	}
}

func BenchmarkCacheGet(b *testing.B) {
	c := New()
	c.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get("key")
	}
}
