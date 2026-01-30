package arrays

import (
	"testing"
)

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name:     "Example 1",
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected: 3,
		},
		{
			name:     "Example 2",
			points:   [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			expected: 4,
		},
		{
			name:     "Single point",
			points:   [][]int{{0, 0}},
			expected: 1,
		},
		{
			name:     "Two points",
			points:   [][]int{{0, 0}, {1, 1}},
			expected: 2,
		},
		{
			name:     "All points on same line",
			points:   [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
			expected: 5,
		},
		{
			name:     "No three points on same line",
			points:   [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			expected: 2,
		},
		{
			name:     "Vertical line",
			points:   [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			expected: 4,
		},
		{
			name:     "Horizontal line",
			points:   [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			expected: 4,
		},
		{
			name:     "Duplicate points",
			points:   [][]int{{0, 0}, {0, 0}, {1, 1}, {1, 1}},
			expected: 4,
		},
		{
			name:     "Complex case 1",
			points:   [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {0, 1}, {1, 0}},
			expected: 4,
		},
		{
			name:     "Complex case 2",
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}, {1, 2}, {2, 3}, {3, 4}},
			expected: 3,
		},
		{
			name:     "Negative coordinates",
			points:   [][]int{{-1, -1}, {-2, -2}, {-3, -3}, {0, 0}},
			expected: 4,
		},
		{
			name:     "Mixed coordinates",
			points:   [][]int{{-1, 1}, {0, 0}, {1, -1}, {2, -2}},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxPoints(tt.points)
			if result != tt.expected {
				t.Errorf("MaxPoints(%v) = %d, expected %d", tt.points, result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxPoints(b *testing.B) {
	// Create a set of points for benchmarking
	points := make([][]int, 100)
	for i := 0; i < 100; i++ {
		points[i] = []int{i, i * i} // Points on a parabola
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxPoints(points)
	}
}

func TestGcd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"gcd(12, 8)", 12, 8, 4},
		{"gcd(8, 12)", 8, 12, 4},
		{"gcd(17, 13)", 17, 13, 1},
		{"gcd(0, 5)", 0, 5, 5},
		{"gcd(5, 0)", 5, 0, 5},
		{"gcd(-12, 8)", -12, 8, 4},
		{"gcd(12, -8)", 12, -8, 4},
		{"gcd(-12, -8)", -12, -8, 4},
		{"gcd(1, 1)", 1, 1, 1},
		{"gcd(100, 75)", 100, 75, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcd(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("gcd(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}