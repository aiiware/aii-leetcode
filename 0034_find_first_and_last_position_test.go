package leetcode

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "Example 3",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "Single occurrence",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: []int{2, 2},
		},
		{
			name:     "Multiple occurrences",
			nums:     []int{1, 2, 2, 2, 3, 4},
			target:   2,
			expected: []int{1, 3},
		},
		{
			name:     "Target at beginning",
			nums:     []int{1, 1, 2, 3, 4},
			target:   1,
			expected: []int{0, 1},
		},
		{
			name:     "Target at end",
			nums:     []int{1, 2, 3, 4, 4},
			target:   4,
			expected: []int{3, 4},
		},
		{
			name:     "All same elements",
			nums:     []int{2, 2, 2, 2, 2},
			target:   2,
			expected: []int{0, 4},
		},
		{
			name:     "Not found in middle",
			nums:     []int{1, 3, 5, 7, 9},
			target:   4,
			expected: []int{-1, -1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test two-pass approach
			result := SearchRange(test.nums, test.target)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("SearchRange(%v, %d) = %v, expected %v", 
					test.nums, test.target, result, test.expected)
			}

			// Test single-pass approach
			result = SearchRangeSinglePass(test.nums, test.target)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("SearchRangeSinglePass(%v, %d) = %v, expected %v", 
					test.nums, test.target, result, test.expected)
			}
		})
	}
}

func BenchmarkSearchRange(b *testing.B) {
	nums := []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 6, 7, 8, 9, 10, 10, 10}
	target := 3
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchRange(nums, target)
	}
}

func BenchmarkSearchRangeSinglePass(b *testing.B) {
	nums := []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 6, 7, 8, 9, 10, 10, 10}
	target := 3
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchRangeSinglePass(nums, target)
	}
}