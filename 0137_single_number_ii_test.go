package leetcode

import (
	"testing"
)

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -1, -2},
			expected: -2,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{3, 3, 3, -4, -5, -5, -5},
			expected: -4,
		},
		{
			name:     "Large array",
			nums:     []int{1, 1, 1, 2, 2, 2, 3, 4, 4, 4},
			expected: 3,
		},
		{
			name:     "Zero is the single number",
			nums:     []int{0, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "All zeros except one",
			nums:     []int{0, 0, 0, 7},
			expected: 7,
		},
		{
			name:     "Maximum values",
			nums:     []int{2147483647, 2147483647, 2147483647, -2147483648},
			expected: -2147483648,
		},
		{
			name:     "Minimum values",
			nums:     []int{-2147483648, -2147483648, -2147483648, 2147483647},
			expected: 2147483647,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SingleNumberII(tt.nums)
			if result != tt.expected {
				t.Errorf("SingleNumberII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
			
			// Also test the optimized version
			resultOpt := SingleNumberIIOptimized(tt.nums)
			if resultOpt != tt.expected {
				t.Errorf("SingleNumberIIOptimized(%v) = %d, expected %d", tt.nums, resultOpt, tt.expected)
			}
		})
	}
}

func TestSingleNumberIIOptimized(t *testing.T) {
	// Additional tests specifically for the optimized version
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Complex bit patterns",
			nums:     []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4},
			expected: 4,
		},
		{
			name:     "Multiple of 3 plus one",
			nums:     []int{10, 10, 10, 20, 20, 20, 30, 30, 30, 40, 40, 40, 50},
			expected: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SingleNumberIIOptimized(tt.nums)
			if result != tt.expected {
				t.Errorf("SingleNumberIIOptimized(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkSingleNumberII(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small array",
			nums: []int{2, 2, 3, 2},
		},
		{
			name: "Medium array",
			nums: []int{0, 1, 0, 1, 0, 1, 99, 2, 2, 2, 3, 3, 3},
		},
		{
			name: "Large array",
			nums: make([]int, 10000), // 3333 triples + 1 single
		},
		{
			name: "Very large array",
			nums: make([]int, 100000), // 33333 triples + 1 single
		},
	}

	// Initialize large arrays
	// For array of size 10000: 3333 triples of numbers 1-3333, and single number 3334
	for i := 0; i < 3333; i++ {
		val := i + 1
		testCases[2].nums[3*i] = val
		testCases[2].nums[3*i+1] = val
		testCases[2].nums[3*i+2] = val
	}
	testCases[2].nums[9999] = 3334

	// For array of size 100000: 33333 triples of numbers 1-33333, and single number 33334
	for i := 0; i < 33333; i++ {
		val := i + 1
		testCases[3].nums[3*i] = val
		testCases[3].nums[3*i+1] = val
		testCases[3].nums[3*i+2] = val
	}
	testCases[3].nums[99999] = 33334

	for _, tc := range testCases {
		b.Run(tc.name+"_bitcount", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SingleNumberII(tc.nums)
			}
		})
		
		b.Run(tc.name+"_optimized", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SingleNumberIIOptimized(tc.nums)
			}
		})
	}
}