package leetcode

import (
	"testing"
)

func TestFindMinBinarySearchII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Basic rotated cases (no duplicates)
		{
			name:     "Example 1: [1,3,5]",
			nums:     []int{1, 3, 5},
			expected: 1,
		},
		{
			name:     "Example 2: [2,2,2,0,1]",
			nums:     []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "Example 3: [3,3,1,3]",
			nums:     []int{3, 3, 1, 3},
			expected: 1,
		},
		// Edge cases with duplicates
		{
			name:     "All duplicates",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Two elements with duplicates",
			nums:     []int{2, 2},
			expected: 2,
		},
		{
			name:     "Two elements different",
			nums:     []int{2, 1},
			expected: 1,
		},
		// Cases from problem 153 (no duplicates)
		{
			name:     "No duplicates: [3,4,5,1,2]",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "No duplicates: [4,5,6,7,0,1,2]",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		// Complex duplicate cases
		{
			name:     "Many duplicates at beginning",
			nums:     []int{2, 2, 2, 2, 3, 4, 5, 0, 1},
			expected: 0,
		},
		{
			name:     "Many duplicates at end",
			nums:     []int{3, 4, 5, 0, 1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "Duplicates around pivot",
			nums:     []int{3, 3, 3, 1, 1, 1, 2, 2, 2},
			expected: 1,
		},
		{
			name:     "Alternating duplicates",
			nums:     []int{2, 3, 2, 2, 2, 2, 2},
			expected: 2,
		},
		// Negative numbers with duplicates
		{
			name:     "Negative numbers with duplicates",
			nums:     []int{-1, -1, -1, -2, -2, -2},
			expected: -2,
		},
		{
			name:     "Mixed positive and negative with duplicates",
			nums:     []int{1, 1, 1, -1, 0, 0, 0},
			expected: -1,
		},
		// Already sorted with duplicates
		{
			name:     "Already sorted ascending with duplicates",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: 1,
		},
		// Large array with many duplicates
		{
			name:     "Large array with pattern",
			nums:     []int{5, 5, 5, 5, 5, 1, 2, 3, 4, 5, 5, 5},
			expected: 1,
		},
		// Tricky case: duplicates make binary search ambiguous
		{
			name:     "Tricky: [10,1,10,10,10]",
			nums:     []int{10, 1, 10, 10, 10},
			expected: 1,
		},
		{
			name:     "Tricky: [3,3,3,3,3,3,3,3,1,3]",
			nums:     []int{3, 3, 3, 3, 3, 3, 3, 3, 1, 3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinBinarySearchII(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinBinarySearchII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinLinearII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Simple case with duplicates",
			nums:     []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "All duplicates",
			nums:     []int{3, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinLinearII(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinLinearII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Main function test 1",
			nums:     []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "Main function test 2",
			nums:     []int{3, 3, 1, 3},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinII(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinRecursiveII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Recursive test 1",
			nums:     []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "Recursive test 2",
			nums:     []int{3, 3, 1, 3},
			expected: 1,
		},
		{
			name:     "Recursive test 3 - all duplicates",
			nums:     []int{4, 4, 4, 4},
			expected: 4,
		},
		{
			name:     "Recursive test 4 - no duplicates",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinRecursiveII(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinRecursiveII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinEarlyExitII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Early exit test 1",
			nums:     []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "Early exit test 2 - many duplicates",
			nums:     []int{3, 3, 3, 3, 3, 1, 2, 3, 3, 3},
			expected: 1,
		},
		{
			name:     "Early exit test 3 - already sorted",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: 1,
		},
		{
			name:     "Early exit test 4 - tricky case",
			nums:     []int{10, 1, 10, 10, 10},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinEarlyExitII(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinEarlyExitII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinTwoPointersII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Two pointers test 1",
			nums:     []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "Two pointers test 2 - many duplicates at ends",
			nums:     []int{3, 3, 3, 3, 1, 2, 3, 3, 3, 3},
			expected: 1,
		},
		{
			name:     "Two pointers test 3 - alternating duplicates",
			nums:     []int{2, 3, 2, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "Two pointers test 4 - no duplicates",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinTwoPointersII(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinTwoPointersII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkFindMinBinarySearchII(b *testing.B) {
	// Test with array that has duplicates
	nums := []int{5, 5, 5, 5, 5, 1, 2, 3, 4, 5, 5, 5}
	for i := 0; i < b.N; i++ {
		findMinBinarySearchII(nums)
	}
}

func BenchmarkFindMinLinearII(b *testing.B) {
	nums := []int{5, 5, 5, 5, 5, 1, 2, 3, 4, 5, 5, 5}
	for i := 0; i < b.N; i++ {
		findMinLinearII(nums)
	}
}

func BenchmarkFindMinRecursiveII(b *testing.B) {
	nums := []int{5, 5, 5, 5, 5, 1, 2, 3, 4, 5, 5, 5}
	for i := 0; i < b.N; i++ {
		findMinRecursiveII(nums)
	}
}

func BenchmarkFindMinEarlyExitII(b *testing.B) {
	nums := []int{5, 5, 5, 5, 5, 1, 2, 3, 4, 5, 5, 5}
	for i := 0; i < b.N; i++ {
		findMinEarlyExitII(nums)
	}
}

func BenchmarkFindMinTwoPointersII(b *testing.B) {
	nums := []int{5, 5, 5, 5, 5, 1, 2, 3, 4, 5, 5, 5}
	for i := 0; i < b.N; i++ {
		findMinTwoPointersII(nums)
	}
}

// Comparative benchmark: array with many duplicates vs few duplicates
func BenchmarkFindMinManyDuplicatesII(b *testing.B) {
	numsManyDups := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2}
	numsFewDups := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}

	b.Run("ManyDuplicates", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			findMinBinarySearchII(numsManyDups)
		}
	})

	b.Run("FewDuplicates", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			findMinBinarySearchII(numsFewDups)
		}
	})
}