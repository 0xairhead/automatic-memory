package mathutils

import "testing"

// Helper for simple test
func TestAddBasic(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

// --- Table-Driven Tests (The Go Way) ---
func TestAddTableDriven(t *testing.T) {
	// 1. Define the table
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive numbers", 2, 2, 4},
		{"negative numbers", -1, -5, -6},
		{"mixed numbers", -5, 5, 0},
		{"zeros", 0, 0, 0},
	}

	// 2. Iterate
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
	}

	for _, tt := range tests {
		t.Run("checking number", func(t *testing.T) {
			if got := IsEven(tt.n); got != tt.want {
				t.Errorf("IsEven(%d) = %v; want %v", tt.n, got, tt.want)
			}
		})
	}
}
