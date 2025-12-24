package leetcode

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Already largest",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Complex case",
			nums:     []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
		{
			name:     "With duplicates",
			nums:     []int{2, 3, 1, 3, 3},
			expected: []int{2, 3, 3, 1, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test optimized version
			nums := make([]int, len(test.nums))
			copy(nums, test.nums)
			NextPermutation(nums)
			
			if !reflect.DeepEqual(nums, test.expected) {
				t.Errorf("NextPermutation(%v) = %v, expected %v", test.nums, nums, test.expected)
			}

			// Test brute force version
			nums2 := make([]int, len(test.nums))
			copy(nums2, test.nums)
			NextPermutationBruteForce(nums2)
			
			if !reflect.DeepEqual(nums2, test.expected) {
				t.Errorf("NextPermutationBruteForce(%v) = %v, expected %v", test.nums, nums2, test.expected)
			}
		})
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		NextPermutation(testNums)
	}
}

func BenchmarkNextPermutationBruteForce(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		NextPermutationBruteForce(testNums)
	}
}