package leetcode

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
