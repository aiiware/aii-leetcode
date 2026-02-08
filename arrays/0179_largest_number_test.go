package arrays

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{10, 2},
			expected: "210",
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{3, 30, 34, 5, 9},
			expected: "9534330",
		},
		{
			name:     "All same digits",
			nums:     []int{1, 1, 1},
			expected: "111",
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: "1",
		},
		{
			name:     "Single element - zero",
			nums:     []int{0},
			expected: "0",
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: "0",
		},
		{
			name:     "All zeros with different counts",
			nums:     []int{0, 0, 0, 0, 0},
			expected: "0",
		},
		{
			name:     "Numbers with leading zeros in comparison",
			nums:     []int{10, 2, 20},
			expected: "22010",
		},
		{
			name:     "Numbers where order matters by first digit",
			nums:     []int{9, 90, 900, 9000},
			expected: "9909009000",
		},
		{
			name:     "Mixed single and multi-digit numbers",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: "98765432110",
		},
		{
			name:     "Large numbers",
			nums:     []int{123, 456, 789},
			expected: "789456123",
		},
		{
			name:     "Numbers with same first digit",
			nums:     []int{838, 83},
			expected: "83883",
		},
		{
			name:     "Numbers with same first digit - different lengths",
			nums:     []int{828, 82},
			expected: "82882",
		},
		{
			name:     "Edge case with 10^9",
			nums:     []int{1000000000, 999999999},
			expected: "9999999991000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("largestNumber() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestLargestNumberBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{10, 2},
			expected: "210",
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{3, 30, 34, 5, 9},
			expected: "9534330",
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: "0",
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: "1",
		},
		{
			name:     "Numbers with same first digit",
			nums:     []int{838, 83},
			expected: "83883",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestNumberBubbleSort(tt.nums)
			if result != tt.expected {
				t.Errorf("largestNumberBubbleSort() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestLargestNumberMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{10, 2},
			expected: "210",
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{3, 30, 34, 5, 9},
			expected: "9534330",
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: "0",
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestNumberMergeSort(tt.nums)
			if result != tt.expected {
				t.Errorf("largestNumberMergeSort() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestLargestNumberQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{10, 2},
			expected: "210",
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{3, 30, 34, 5, 9},
			expected: "9534330",
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: "0",
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestNumberQuickSort(tt.nums)
			if result != tt.expected {
				t.Errorf("largestNumberQuickSort() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func BenchmarkLargestNumber(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small array (2 elements)",
			nums: []int{10, 2},
		},
		{
			name: "Medium array (10 elements)",
			nums: []int{3, 30, 34, 5, 9, 1, 2, 4, 6, 8},
		},
		{
			name: "Large array (100 elements)",
			nums: func() []int {
				arr := make([]int, 100)
				for i := range arr {
					arr[i] = i + 1
				}
				return arr
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name+"-Standard", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)
				largestNumber(numsCopy)
			}
		})

		b.Run(tc.name+"-BubbleSort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)
				largestNumberBubbleSort(numsCopy)
			}
		})

		b.Run(tc.name+"-MergeSort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)
				largestNumberMergeSort(numsCopy)
			}
		})

		b.Run(tc.name+"-QuickSort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)
				largestNumberQuickSort(numsCopy)
			}
		})
	}
}
