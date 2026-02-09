package arrays

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{"Example 1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"Example 2", []int{1, 1}, 1},
		{"Two equal heights", []int{2, 2, 2, 2}, 6},
		{"Single large height", []int{1, 100, 1}, 2},
		{"Increasing sequence", []int{1, 2, 3, 4, 5}, 6},
		{"Decreasing sequence", []int{5, 4, 3, 2, 1}, 6},
		{"Large height at edges", []int{100, 1, 1, 100}, 300},
		// New edge cases
		{"Empty array", []int{}, 0},
		{"Single element", []int{5}, 0},
		{"All zeros", []int{0, 0, 0, 0}, 0},
		{"Very large numbers", []int{1000000, 500000, 1000000}, 2000000}, // Fixed: min(1000000,1000000)*2 = 2000000
		{"Alternating heights", []int{1, 100, 1, 100, 1}, 200},           // Fixed: min(100,100)*2 = 200
		{"Peak in middle", []int{1, 2, 100, 2, 1}, 4},
		{"Valley in middle", []int{100, 1, 2, 1, 100}, 400},
		{"Large array with pattern", []int{1, 3, 5, 7, 9, 7, 5, 3, 1}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxArea(tt.height)
			if result != tt.expected {
				t.Errorf("MaxArea(%v) = %d, expected %d", tt.height, result, tt.expected)
			}
		})
	}
}

func TestMaxArea_EdgeCases(t *testing.T) {
	t.Run("Nil slice", func(t *testing.T) {
		result := MaxArea(nil)
		if result != 0 {
			t.Errorf("MaxArea(nil) = %d, expected 0", result)
		}
	})

	t.Run("Two elements with zero height", func(t *testing.T) {
		result := MaxArea([]int{0, 0})
		if result != 0 {
			t.Errorf("MaxArea([0, 0]) = %d, expected 0", result)
		}
	})

	t.Run("Extremely large array", func(t *testing.T) {
		// Create array with 10000 elements
		height := make([]int, 10000)
		for i := range height {
			height[i] = i % 100
		}
		result := MaxArea(height)
		// The maximum area should be between indices 99 and 9999:
		// height[99] = 99, height[9999] = 99, distance = 9900
		// area = min(99, 99) * 9900 = 99 * 9900 = 980100
		expected := 980100
		if result != expected {
			t.Errorf("MaxArea(large array) = %d, expected %d", result, expected)
		}
	})

	t.Run("All same height large array", func(t *testing.T) {
		height := make([]int, 1000)
		for i := range height {
			height[i] = 50
		}
		result := MaxArea(height)
		// Maximum area is between first and last: 50 * 999 = 49950
		expected := 50 * 999
		if result != expected {
			t.Errorf("MaxArea(all same) = %d, expected %d", result, expected)
		}
	})
}

func BenchmarkMaxArea(b *testing.B) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxAreaLarge(b *testing.B) {
	// Create a large test case
	height := make([]int, 1000)
	for i := 0; i < len(height); i++ {
		height[i] = (i % 100) + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxAreaVeryLarge(b *testing.B) {
	// Create a very large test case
	height := make([]int, 10000)
	for i := 0; i < len(height); i++ {
		height[i] = (i % 200) + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}
