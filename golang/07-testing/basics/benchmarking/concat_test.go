package benchmarking

import "testing"

// generateSlice creates a slice of "n" strings
func generateSlice(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = "a"
	}
	return res
}

// Benchmark naming: BenchmarkXxx(b *testing.B)
// Run with: go test -bench=.

func BenchmarkCatPlus(b *testing.B) {
	parts := generateSlice(1000)
	// Reset timer to ignore setup time
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CatPlus(parts)
	}
}

func BenchmarkCatBuilder(b *testing.B) {
	parts := generateSlice(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CatBuilder(parts)
	}
}

func BenchmarkCatBuffer(b *testing.B) {
	parts := generateSlice(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CatBuffer(parts)
	}
}
