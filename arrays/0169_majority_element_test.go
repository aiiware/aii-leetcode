package arrays

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Basic test cases
		{
			name:     "Example 1",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "All same elements",
			nums:     []int{7, 7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "Majority at beginning",
			nums:     []int{4, 4, 4, 4, 1, 2},
			expected: 4,
		},
		{
			name:     "Majority at end",
			nums:     []int{1, 2, 3, 5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "Even length array",
			nums:     []int{3, 3, 3, 3, 2, 2}, // 4 out of 6
			expected: 3,
		},
		{
			name:     "Large array",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2}, // 8 out of 12
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -1, 2, 3},
			expected: -1,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-5, -5, -5, -5, 1, 2, 3},
			expected: -5,
		},
		{
			name:     "Zero as majority",
			nums:     []int{0, 0, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "Alternating pattern",
			nums:     []int{1, 2, 1, 2, 1, 2, 1}, // 4 ones, 3 twos
			expected: 1,
		},
	}

	// Test all solution functions
	solutionFuncs := []struct {
		name string
		fn   func([]int) int
	}{
		{"BoyerMoore", majorityElementBoyerMoore},
		{"HashMap", majorityElementHashMap},
		{"Sorting", majorityElementSorting},
		{"DivideConquer", majorityElementDivideConquer},
		{"BitManipulation", majorityElementBitManipulation},
		{"WithVerification", majorityElementWithVerification},
		{"BuiltinSort", majorityElementBuiltinSort},
		{"Random", majorityElementRandom},
	}

	for _, solution := range solutionFuncs {
		t.Run(solution.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					result := solution.fn(tt.nums)
					if result != tt.expected {
						t.Errorf("%s(%v) = %d, expected %d", solution.name, tt.nums, result, tt.expected)
					}
				})
			}
		})
	}
}

func TestMajorityElementEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Minimum size (1 element)",
			nums:     []int{42},
			expected: 42,
		},
		{
			name:     "Two elements same",
			nums:     []int{7, 7},
			expected: 7,
		},
		{
			name:     "Three elements with clear majority",
			nums:     []int{2, 3, 2},
			expected: 2,
		},
		{
			name:     "Large numbers",
			nums:     []int{1000000, 1000000, 999999},
			expected: 1000000,
		},
		{
			name:     "Max int values",
			nums:     []int{2147483647, 2147483647, -2147483648},
			expected: 2147483647,
		},
	}

	for _, solution := range []struct {
		name string
		fn   func([]int) int
	}{
		{"BoyerMoore", majorityElementBoyerMoore},
		{"HashMap", majorityElementHashMap},
	} {
		t.Run(solution.name, func(t *testing.T) {
			for _, tc := range edgeCases {
				t.Run(tc.name, func(t *testing.T) {
					result := solution.fn(tc.nums)
					if result != tc.expected {
						t.Errorf("%s(%v) = %d, expected %d", solution.name, tc.nums, result, tc.expected)
					}
				})
			}
		})
	}
}

func BenchmarkMajorityElementBoyerMoore(b *testing.B) {
	nums := make([]int, 10000)
	// Create array with majority element 42
	for i := 0; i < 10000; i++ {
		if i < 5001 {
			nums[i] = 42 // Majority
		} else {
			nums[i] = i // Other elements
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElementBoyerMoore(nums)
	}
}

func BenchmarkMajorityElementHashMap(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		if i < 5001 {
			nums[i] = 42
		} else {
			nums[i] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElementHashMap(nums)
	}
}

func BenchmarkMajorityElementSorting(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		if i < 5001 {
			nums[i] = 42
		} else {
			nums[i] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElementSorting(nums)
	}
}

func BenchmarkMajorityElementDivideConquer(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		if i < 5001 {
			nums[i] = 42
		} else {
			nums[i] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElementDivideConquer(nums)
	}
}

func TestAllSolutionsConsistency(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 2, 2, 2, 5, 4, 2},
		{3, 3, 4, 2, 4, 4, 2, 4, 4},
		{6, 5, 5},
		{2, 2, 1, 1, 1, 2, 2},
		{1},
		{0, 0, 0},
		{-1, -1, 2},
	}

	solutions := []struct {
		name string
		fn   func([]int) int
	}{
		{"BoyerMoore", majorityElementBoyerMoore},
		{"HashMap", majorityElementHashMap},
		{"Sorting", majorityElementSorting},
		{"DivideConquer", majorityElementDivideConquer},
		{"BitManipulation", majorityElementBitManipulation},
		{"WithVerification", majorityElementWithVerification},
		{"BuiltinSort", majorityElementBuiltinSort},
	}

	for _, nums := range testCases {
		// Get result from first solution
		expected := majorityElementBoyerMoore(nums)

		for _, solution := range solutions {
			result := solution.fn(nums)
			if result != expected {
				t.Errorf("Inconsistency for nums=%v: %s=%d, expected=%d",
					nums, solution.name, result, expected)
			}
		}
	}
}

// Test that the main function works correctly
func TestMainFunction(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{5, 5, 5, 2, 3}, 5},
	}

	for _, tt := range tests {
		result := majorityElement(tt.nums)
		if result != tt.expected {
			t.Errorf("majorityElement(%v) = %d, expected %d", tt.nums, result, tt.expected)
		}
	}
}
