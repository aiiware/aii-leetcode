package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "Example 2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			name:     "Empty array",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			height:   []int{5},
			expected: 0,
		},
		{
			name:     "Two elements",
			height:   []int{5, 3},
			expected: 0,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Decreasing heights",
			height:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Flat terrain",
			height:   []int{3, 3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Valley in middle",
			height:   []int{5, 1, 1, 1, 5},
			expected: 12, // 4 units between each pair: (5-1)*3 = 12
		},
		{
			name:     "Multiple peaks",
			height:   []int{0, 2, 0, 1, 0, 3, 0, 2, 0},
			expected: 7, // 2 + 1 + 2 + 2 = 7
		},
		{
			name:     "All zeros",
			height:   []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "One tall building in middle",
			height:   []int{1, 0, 2, 0, 1},
			expected: 2, // 1 unit on left, 1 unit on right
		},
		{
			name:     "Complex terrain",
			height:   []int{2, 0, 3, 0, 4, 0, 1},
			expected: 6, // 2 + 3 + 1 = 6
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Trap(tt.height)
			assert.Equal(t, tt.expected, result,
				"Trap(%v) = %d, expected %d",
				tt.height, result, tt.expected)
		})
	}
}

func TestTrap_EdgeCases(t *testing.T) {
	t.Run("Nil slice returns 0", func(t *testing.T) {
		result := Trap(nil)
		assert.Equal(t, 0, result)
	})

	t.Run("Large symmetric valley", func(t *testing.T) {
		height := make([]int, 1001)
		for i := range height {
			if i < 500 {
				height[i] = i
			} else {
				height[i] = 1000 - i
			}
		}
		result := Trap(height)
		// The valley should trap water in the middle
		// Maximum height is 500 at index 500
		// Water trapped = sum of (500 - height[i]) for i where height[i] < 500
		assert.True(t, result > 0, "Should trap some water")
	})

	t.Run("Alternating high-low", func(t *testing.T) {
		height := make([]int, 100)
		for i := range height {
			if i%2 == 0 {
				height[i] = 10
			} else {
				height[i] = 1
			}
		}
		result := Trap(height)
		// Between each pair of high points (10), there's a low point (1)
		// Water trapped = (10-1) * number of low points between highs
		// There are 49 low points between highs (indices 1, 3, 5, ..., 97, 99)
		// But the first and last low points don't have walls on both sides
		// So actual trapped water is less
		assert.Equal(t, 9*48, result) // 48 interior low points * 9 units each
	})
}

func BenchmarkTrap(b *testing.B) {
	// Create a large elevation map for benchmarking
	height := make([]int, 10000)
	for i := range height {
		// Create a wave pattern
		height[i] = (i % 100) * 2
		if height[i] > 100 {
			height[i] = 200 - height[i]
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trap(height)
	}
}