package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "All duplicates",
			nums:     []int{1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-5, -4, -3, -2, -1, 0},
			expected: 6,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-1, 0, 1, 2, -2, -3},
			expected: 6,
		},
		{
			name:     "Multiple sequences",
			nums:     []int{1, 2, 3, 10, 11, 12, 20, 21},
			expected: 3,
		},
		{
			name:     "Large gap between sequences",
			nums:     []int{1, 2, 3, 100, 101, 102, 103},
			expected: 4,
		},
		{
			name:     "Unsorted with duplicates",
			nums:     []int{5, 3, 3, 2, 4, 4, 1, 1},
			expected: 5,
		},
		{
			name:     "All consecutive",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 10,
		},
		{
			name:     "No consecutive",
			nums:     []int{1, 3, 5, 7, 9},
			expected: 1,
		},
		{
			name:     "With zero",
			nums:     []int{0, -1, 1, -2, 2},
			expected: 5,
		},
		{
			name:     "Large numbers",
			nums:     []int{1000000000, 999999999, 999999998},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestConsecutive(tt.nums)
			assert.Equal(t, tt.expected, result,
				"LongestConsecutive(%v) = %d, expected %d",
				tt.nums, result, tt.expected)
		})
	}
}

func TestLongestConsecutive_EdgeCases(t *testing.T) {
	t.Run("Very large array", func(t *testing.T) {
		// Create array with 100000 elements, all consecutive
		nums := make([]int, 100000)
		for i := range nums {
			nums[i] = i - 50000 // From -50000 to 49999
		}

		result := LongestConsecutive(nums)
		assert.Equal(t, 100000, result)
	})

	t.Run("Array with minimum values", func(t *testing.T) {
		nums := []int{-1000000000, -999999999, -999999998}
		result := LongestConsecutive(nums)
		assert.Equal(t, 3, result)
	})

	t.Run("Array with maximum values", func(t *testing.T) {
		nums := []int{1000000000, 999999999, 999999998}
		result := LongestConsecutive(nums)
		assert.Equal(t, 3, result)
	})

	t.Run("Random order with long sequence", func(t *testing.T) {
		nums := []int{50, 1, 49, 2, 48, 3, 47, 4, 46, 5, 45, 6, 44, 7, 43, 8, 42, 9, 41, 10}
		// The sequence 1-10 is length 10, but there might be longer
		result := LongestConsecutive(nums)
		assert.Equal(t, 10, result) // 1-10 or 41-50 both length 10
	})

	t.Run("All same number", func(t *testing.T) {
		nums := make([]int, 1000)
		for i := range nums {
			nums[i] = 42
		}
		result := LongestConsecutive(nums)
		assert.Equal(t, 1, result)
	})
}

func BenchmarkLongestConsecutive(b *testing.B) {
	// Create a large array for benchmarking
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i % 1000 // Creates many sequences of length up to 1000
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestConsecutive(nums)
	}
}