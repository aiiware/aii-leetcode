package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "All zeros except last",
			nums:     []int{0, 0, 0, 0, 1},
			expected: false,
		},
		{
			name:     "Can jump directly",
			nums:     []int{4, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Large jump at beginning",
			nums:     []int{5, 0, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Zero at second position",
			nums:     []int{1, 0, 2, 0, 0},
			expected: false,
		},
		{
			name:     "Progressive jumps",
			nums:     []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "Can barely reach",
			nums:     []int{2, 0, 1, 0, 1},
			expected: false,
		},
		{
			name:     "Can reach with exact jump",
			nums:     []int{2, 0, 1, 0, 2},
			expected: false,
		},
		{
			name:     "Large numbers",
			nums:     []int{10, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Complex case reachable",
			nums:     []int{2, 3, 1, 1, 2, 0, 1},
			expected: true,
		},
		{
			name:     "Complex case unreachable",
			nums:     []int{2, 3, 1, 0, 2, 0, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanJump(tt.nums)
			assert.Equal(t, tt.expected, result,
				"CanJump(%v) = %v, expected %v",
				tt.nums, result, tt.expected)
		})
	}
}

func TestCanJump_EdgeCases(t *testing.T) {
	t.Run("All zeros", func(t *testing.T) {
		result := CanJump([]int{0, 0, 0, 0, 0})
		assert.Equal(t, false, result)
	})

	t.Run("Single large jump", func(t *testing.T) {
		result := CanJump([]int{100})
		assert.Equal(t, true, result)
	})

	t.Run("Decreasing jumps", func(t *testing.T) {
		result := CanJump([]int{5, 4, 3, 2, 1, 0})
		assert.Equal(t, true, result)
	})

	t.Run("Increasing jumps", func(t *testing.T) {
		result := CanJump([]int{1, 2, 3, 4, 5})
		assert.Equal(t, true, result)
	})

	t.Run("Zero at last position", func(t *testing.T) {
		result := CanJump([]int{1, 2, 3, 0})
		assert.Equal(t, true, result)
	})
}

func BenchmarkCanJump(b *testing.B) {
	// Create a large array for benchmarking
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = 10000 - i // Decreasing jumps
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanJump(nums)
	}
}
