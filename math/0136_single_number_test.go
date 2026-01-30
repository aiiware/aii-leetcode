package math

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "Example 2",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -2},
			expected: -2,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{5, -3, 5, 2, -3},
			expected: 2,
		},
		{
			name:     "Large array",
			nums:     []int{1, 2, 3, 4, 5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Zero is the single number",
			nums:     []int{0, 1, 1},
			expected: 0,
		},
		{
			name:     "All zeros except one",
			nums:     []int{0, 0, 0, 0, 7},
			expected: 7,
		},
		{
			name:     "Maximum values",
			nums:     []int{30000, 30000, -30000},
			expected: -30000,
		},
		{
			name:     "Minimum values",
			nums:     []int{-30000, -30000, 30000},
			expected: 30000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SingleNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("SingleNumber(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small array",
			nums: []int{2, 2, 1},
		},
		{
			name: "Medium array",
			nums: []int{4, 1, 2, 1, 2, 3, 3, 5, 5, 6, 6, 7, 7},
		},
		{
			name: "Large array",
			nums: make([]int, 10001), // Odd number to have a single number
		},
		{
			name: "Very large array",
			nums: make([]int, 100001), // Odd number to have a single number
		},
	}

	// Initialize large arrays
	// For array of size 10001: pairs from 1 to 5000, and single number 5001
	for i := 0; i < 5000; i++ {
		testCases[2].nums[2*i] = i + 1
		testCases[2].nums[2*i+1] = i + 1
	}
	testCases[2].nums[10000] = 5001

	// For array of size 100001: pairs from 1 to 50000, and single number 50001
	for i := 0; i < 50000; i++ {
		testCases[3].nums[2*i] = i + 1
		testCases[3].nums[2*i+1] = i + 1
	}
	testCases[3].nums[100000] = 50001

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SingleNumber(tc.nums)
			}
		})
	}
}