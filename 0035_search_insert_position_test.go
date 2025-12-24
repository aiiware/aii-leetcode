package leetcode

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "Example 4",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element less than target",
			nums:     []int{1},
			target:   2,
			expected: 1,
		},
		{
			name:     "Single element greater than target",
			nums:     []int{3},
			target:   2,
			expected: 0,
		},
		{
			name:     "Single element equal to target",
			nums:     []int{2},
			target:   2,
			expected: 0,
		},
		{
			name:     "Insert at beginning",
			nums:     []int{2, 4, 6, 8},
			target:   1,
			expected: 0,
		},
		{
			name:     "Insert at end",
			nums:     []int{2, 4, 6, 8},
			target:   10,
			expected: 4,
		},
		{
			name:     "Insert in middle",
			nums:     []int{1, 3, 5, 7, 9},
			target:   6,
			expected: 3,
		},
		{
			name:     "Duplicate values",
			nums:     []int{1, 3, 3, 3, 5},
			target:   3,
			expected: 1, // First occurrence
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test binary search approach
			result := SearchInsert(test.nums, test.target)
			if result != test.expected {
				t.Errorf("SearchInsert(%v, %d) = %d, expected %d", 
					test.nums, test.target, result, test.expected)
			}

			// Test linear search approach
			result = SearchInsertLinear(test.nums, test.target)
			if result != test.expected {
				t.Errorf("SearchInsertLinear(%v, %d) = %d, expected %d", 
					test.nums, test.target, result, test.expected)
			}

			// Test recursive approach
			result = SearchInsertRecursive(test.nums, test.target)
			if result != test.expected {
				t.Errorf("SearchInsertRecursive(%v, %d) = %d, expected %d", 
					test.nums, test.target, result, test.expected)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	target := 16
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInsert(nums, target)
	}
}

func BenchmarkSearchInsertLinear(b *testing.B) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	target := 16
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInsertLinear(nums, target)
	}
}

func BenchmarkSearchInsertRecursive(b *testing.B) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	target := 16
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInsertRecursive(nums, target)
	}
}