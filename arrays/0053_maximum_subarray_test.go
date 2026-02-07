package arrays

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
			expected: -1, // The maximum subarray is the single element -1
		},
		{
			name:     "Mixed with zero",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "Single negative",
			nums:     []int{-10},
			expected: -10,
		},
		{
			name:     "Large array",
			nums:     []int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10},
			expected: 9, // The subarray [9] or [7, -8, 9] = 8
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
	t.Run("Nil slice returns 0", func(t *testing.T) {
		result := MaxSubArray(nil)
		assert.Equal(t, 0, result)
	})

	t.Run("Large positive sum", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := range nums {
			nums[i] = 1
		}
		result := MaxSubArray(nums)
		assert.Equal(t, 1000, result)
	})

	t.Run("Large alternating", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := range nums {
			if i%2 == 0 {
				nums[i] = 1000
			} else {
				nums[i] = -999
			}
		}
		result := MaxSubArray(nums)
		// Best is to take all even indices (1000 each)
		assert.Equal(t, 500*1000, result) // 500 even indices
	})
}

func BenchmarkMaxSubArray(b *testing.B) {
	// Create a large array for benchmarking
	nums := make([]int, 10000)
	for i := range nums {
		// Mix of positive and negative numbers
		if i%3 == 0 {
			nums[i] = -i
		} else if i%3 == 1 {
			nums[i] = i
		} else {
			nums[i] = 0
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxSubArray(nums)
	}
}