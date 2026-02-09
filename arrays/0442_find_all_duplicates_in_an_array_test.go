package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{2, 3},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 1, 2},
			expected: []int{1},
		},
		{
			name:     "Example 3",
			nums:     []int{1},
			expected: []int{},
		},
		{
			name:     "No duplicates",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "All duplicates",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Single duplicate at end",
			nums:     []int{1, 2, 3, 4, 5, 5},
			expected: []int{5},
		},
		{
			name:     "Single duplicate at beginning",
			nums:     []int{1, 1, 2, 3, 4, 5},
			expected: []int{1},
		},
		{
			name:     "Multiple duplicates mixed",
			nums:     []int{3, 2, 1, 3, 2, 4, 5, 5},
			expected: []int{3, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums since the function modifies the input
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			result := FindDuplicates(numsCopy)

			// Sort both slices for comparison (order doesn't matter in problem)
			sortSlice(result)
			sortSlice(tt.expected)

			assert.Equal(t, tt.expected, result,
				"FindDuplicates(%v) = %v, expected %v", tt.nums, result, tt.expected)
		})
	}
}

func BenchmarkFindDuplicates(b *testing.B) {
	// Create a large test case for benchmarking
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = (i % 50000) + 1 // Creates duplicates
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		FindDuplicates(numsCopy)
	}
}

// Helper function to sort int slices
func sortSlice(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
