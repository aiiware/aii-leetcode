package leetcode

import (
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "Not rotated",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			target:   4,
			expected: 3,
		},
		{
			name:     "Rotated at beginning",
			nums:     []int{3, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "Rotated at end",
			nums:     []int{1, 3},
			target:   3,
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Single element found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Complex rotation",
			nums:     []int{5, 6, 7, 8, 9, 1, 2, 3, 4},
			target:   6,
			expected: 1,
		},
		{
			name:     "Target in second half",
			nums:     []int{5, 6, 7, 8, 9, 1, 2, 3, 4},
			target:   3,
			expected: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test single pass approach
			result := SearchInRotatedSortedArray(test.nums, test.target)
			if result != test.expected {
				t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d, expected %d", 
					test.nums, test.target, result, test.expected)
			}

			// Test two-pass approach
			result = SearchInRotatedSortedArrayTwoPass(test.nums, test.target)
			if result != test.expected {
				t.Errorf("SearchInRotatedSortedArrayTwoPass(%v, %d) = %d, expected %d", 
					test.nums, test.target, result, test.expected)
			}
		})
	}
}

func BenchmarkSearchInRotatedSortedArray(b *testing.B) {
	nums := []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	target := 7
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInRotatedSortedArray(nums, target)
	}
}

func BenchmarkSearchInRotatedSortedArrayTwoPass(b *testing.B) {
	nums := []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	target := 7
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInRotatedSortedArrayTwoPass(nums, target)
	}
}