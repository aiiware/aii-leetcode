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

	t.Run("Large symmetric valley (mountain shape)", func(t *testing.T) {
		// This creates a mountain shape: 0, 1, 2, ..., 500, 499, ..., 1, 0
		// Since it's a mountain (increasing then decreasing), no water is trapped
		height := make([]int, 1001)
		for i := range height {
			if i < 500 {
				height[i] = i
			} else {
				height[i] = 1000 - i
			}
		}
		result := Trap(height)
		// No water trapped because it's a mountain shape, not a valley
		assert.Equal(t, 0, result)
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
		// For alternating [10, 1, 10, 1, ...], the DP approach calculates:
		// At each 1 (odd indices), water = min(leftMax, rightMax) - 1
		// At indices 1, 3, 5, ..., 97: leftMax=10, rightMax=10, water = 10-1 = 9 (49 positions)
		// At index 99 (last): leftMax=10, rightMax=1, water = min(10,1)-1 = 0
		// Total = 49 * 9 = 441
		assert.Equal(t, 441, result)
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
