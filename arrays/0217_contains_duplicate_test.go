package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1: Contains duplicate",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Example 2: No duplicate",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Example 3: Multiple duplicates",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "Two identical elements",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "Negative numbers with duplicate",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "Large numbers",
			nums:     []int{1000000, 2000000, 3000000, 1000000},
			expected: true,
		},
		{
			name:     "Zero values",
			nums:     []int{0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-5, 10, -5, 20},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicate(tt.nums)
			assert.Equal(t, tt.expected, result,
				"ContainsDuplicate(%v) = %v, expected %v",
				tt.nums, result, tt.expected)
		})
	}
}

func TestContainsDuplicateSorting(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Contains duplicate",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "No duplicate",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Multiple duplicates",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicateSorting(tt.nums)
			assert.Equal(t, tt.expected, result,
				"ContainsDuplicateSorting(%v) = %v, expected %v",
				tt.nums, result, tt.expected)
		})
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	// Test with large array containing duplicates
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	// Add a duplicate at the end
	nums[9999] = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicate(nums)
	}
}

func BenchmarkContainsDuplicateSorting(b *testing.B) {
	// Test with large array
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicateSorting(nums)
	}
}
