package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "Single element found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element not found",
			nums:     []int{5},
			target:   2,
			expected: -1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "First element",
			nums:     []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Last element",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Middle element",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Large array",
			nums:     makeRangeLocal(1, 1000),
			target:   750,
			expected: 749,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-10, -5, -3, 0, 2, 7},
			target:   -3,
			expected: 2,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-100, -50, -25, -10, -5},
			target:   -25,
			expected: 2,
		},
		{
			name:     "Duplicate numbers - first occurrence",
			nums:     []int{1, 2, 2, 2, 3, 4},
			target:   2,
			expected: 2, // Binary search finds one of the 2's (index 2)
		},
		{
			name:     "Even length array",
			nums:     []int{1, 3, 5, 7},
			target:   5,
			expected: 2,
		},
		{
			name:     "Odd length array",
			nums:     []int{1, 3, 5, 7, 9},
			target:   7,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result,
				"BinarySearch(%v, %d) = %d, expected %d",
				tt.nums, tt.target, result, tt.expected)
		})
	}
}

func TestBinarySearch_EdgeCases(t *testing.T) {
	t.Run("Target less than all elements", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50}
		result := BinarySearch(nums, 5)
		assert.Equal(t, -1, result)
	})

	t.Run("Target greater than all elements", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50}
		result := BinarySearch(nums, 55)
		assert.Equal(t, -1, result)
	})

	t.Run("Array with one element not found", func(t *testing.T) {
		result := BinarySearch([]int{42}, 100)
		assert.Equal(t, -1, result)
	})

	t.Run("Array with two elements found first", func(t *testing.T) {
		result := BinarySearch([]int{1, 2}, 1)
		assert.Equal(t, 0, result)
	})

	t.Run("Array with two elements found second", func(t *testing.T) {
		result := BinarySearch([]int{1, 2}, 2)
		assert.Equal(t, 1, result)
	})
}

func BenchmarkBinarySearch(b *testing.B) {
	// Create a large sorted array for benchmarking
	nums := makeRangeLocal(1, 1000000)
	target := 750000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(nums, target)
	}
}

// Helper function to create a range of integers
// Renamed to avoid conflict with makeRange in arrays/0004_median_sorted_arrays_test.go
func makeRangeLocal(start, end int) []int {
	if start > end {
		return []int{}
	}
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}