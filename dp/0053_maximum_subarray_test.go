package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "All positive",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "All negative",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -1,
		},
		{
			name:     "Mixed with zero",
			nums:     []int{-2, 0, -1, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "Single negative",
			nums:     []int{-10},
			expected: -10,
		},
		{
			name:     "Large array",
			nums:     []int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10},
			expected: 9,
		},
		{
			name:     "Alternating positive negative",
			nums:     []int{1, -1, 1, -1, 1, -1, 1},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSubArray(tt.nums)
			assert.Equal(t, tt.expected, result,
				"MaxSubArray(%v) = %d, expected %d",
				tt.nums, result, tt.expected)
		})
	}
}

func TestMaxSubArray_EdgeCases(t *testing.T) {
	t.Run("All zeros", func(t *testing.T) {
		result := MaxSubArray([]int{0, 0, 0, 0, 0})
		assert.Equal(t, 0, result)
	})

	t.Run("One positive rest negative", func(t *testing.T) {
		result := MaxSubArray([]int{-10, -5, -3, 7, -2, -1})
		assert.Equal(t, 7, result)
	})

	t.Run("Max at beginning", func(t *testing.T) {
		result := MaxSubArray([]int{10, -2, -3, -4, -5})
		assert.Equal(t, 10, result)
	})

	t.Run("Max at end", func(t *testing.T) {
		result := MaxSubArray([]int{-5, -4, -3, -2, 10})
		assert.Equal(t, 10, result)
	})
}

func BenchmarkMaxSubArray(b *testing.B) {
	// Create a large array for benchmarking
	nums := make([]int, 10000)
	for i := range nums {
		if i%3 == 0 {
			nums[i] = -i
		} else {
			nums[i] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxSubArray(nums)
	}
}