package leetcode

import (
	"testing"
)

func TestFindMinBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Basic rotated cases
		{
			name:     "Example 1: [3,4,5,1,2]",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Example 2: [4,5,6,7,0,1,2]",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Example 3: [11,13,15,17] (not rotated)",
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		},
		// Edge cases
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Two elements rotated",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "Two elements not rotated",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "Three elements rotated at end",
			nums:     []int{2, 3, 1},
			expected: 1,
		},
		{
			name:     "Three elements rotated at middle",
			nums:     []int{3, 1, 2},
			expected: 1,
		},
		// Larger arrays
		{
			name:     "Large rotated array",
			nums:     []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "Rotated at beginning",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 1,
		},
		{
			name:     "Rotated at end",
			nums:     []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 1,
		},
		// Negative numbers
		{
			name:     "With negative numbers",
			nums:     []int{4, 5, 6, 7, -2, -1, 0, 1, 2},
			expected: -2,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-5, -4, -3, -2, -1},
			expected: -5,
		},
		{
			name:     "Negative numbers rotated",
			nums:     []int{-1, -5, -4, -3, -2},
			expected: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinBinarySearch(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinBinarySearch(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinLinear(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Simple case",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Already sorted",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinLinear(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinLinear(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Main function test 1",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Main function test 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.nums)
			if result != tt.expected {
				t.Errorf("findMin(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinAlt(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Alternative implementation test 1",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Alternative implementation test 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Already sorted",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinAlt(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinAlt(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinRecursive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Recursive test 1",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "Recursive test 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Recursive test 3 - single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Recursive test 4 - not rotated",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Recursive test 5 - negative numbers",
			nums:     []int{4, 5, 6, 7, -2, -1, 0, 1, 2},
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinRecursive(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinRecursive(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindMinOnePass(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "One-pass test 1",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "One-pass test 2",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "One-pass test 3 - not rotated",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "One-pass test 4 - rotated at end",
			nums:     []int{2, 3, 4, 5, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinOnePass(tt.nums)
			if result != tt.expected {
				t.Errorf("findMinOnePass(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkFindMinBinarySearch(b *testing.B) {
	nums := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		findMinBinarySearch(nums)
	}
}

func BenchmarkFindMinLinear(b *testing.B) {
	nums := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		findMinLinear(nums)
	}
}

func BenchmarkFindMinAlt(b *testing.B) {
	nums := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		findMinAlt(nums)
	}
}

func BenchmarkFindMinRecursive(b *testing.B) {
	nums := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		findMinRecursive(nums)
	}
}

func BenchmarkFindMinOnePass(b *testing.B) {
	nums := []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		findMinOnePass(nums)
	}
}