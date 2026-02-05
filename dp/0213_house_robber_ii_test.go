package dp

import (
	"testing"
)

func TestRobII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1 - circular adjacency",
			nums:     []int{2, 3, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "Single house",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Two houses - rob first",
			nums:     []int{2, 1},
			expected: 2,
		},
		{
			name:     "Two houses - rob second",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "Four houses - circular case",
			nums:     []int{1, 3, 1, 3},
			expected: 6,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Large values with circular constraint",
			nums:     []int{100, 50, 400, 200, 100},
			expected: 500,
		},
		{
			name:     "Case where first and last are both optimal but adjacent",
			nums:     []int{10, 1, 1, 10},
			expected: 11,
		},
		{
			name:     "All houses same value",
			nums:     []int{5, 5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "Alternating high values",
			nums:     []int{100, 1, 100, 1, 100},
			expected: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test main implementation
			result := RobII(tt.nums)
			if result != tt.expected {
				t.Errorf("RobII(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}

			// Test DP array implementation
			resultDP := RobIIDPArray(tt.nums)
			if resultDP != tt.expected {
				t.Errorf("RobIIDPArray(%v) = %d, expected %d", tt.nums, resultDP, tt.expected)
			}

			// Test recursive implementation
			resultRecursive := RobIIRecursive(tt.nums)
			if resultRecursive != tt.expected {
				t.Errorf("RobIIRecursive(%v) = %d, expected %d", tt.nums, resultRecursive, tt.expected)
			}
		})
	}
}

func TestRobIIEdgeCases(t *testing.T) {
	// Test with maximum constraints (100 houses, max value 1000)
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = 1000
	}
	// For circular houses: can't rob first and last, so we rob 50 houses
	expected := 50 * 1000
	result := RobII(nums)
	if result != expected {
		t.Errorf("RobII(large array) = %d, expected %d", result, expected)
	}

	// Test with alternating pattern that would be optimal in non-circular case
	nums2 := make([]int, 99) // odd number to test edge case
	for i := range nums2 {
		if i%2 == 0 {
			nums2[i] = 1000
		} else {
			nums2[i] = 1
		}
	}
	// In circular case, we have to skip either first or last house
	// So we get (n-1)/2 houses * 1000
	expected2 := 49 * 1000
	result2 := RobII(nums2)
	if result2 != expected2 {
		t.Errorf("RobII(alternating pattern) = %d, expected %d", result2, expected2)
	}
}

func BenchmarkRobII(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2, 7, 9, 3, 1, 5, 8, 4, 6}
	for i := 0; i < b.N; i++ {
		RobII(nums)
	}
}

func BenchmarkRobIIDPArray(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2, 7, 9, 3, 1, 5, 8, 4, 6}
	for i := 0; i < b.N; i++ {
		RobIIDPArray(nums)
	}
}

func BenchmarkRobIIRecursive(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2, 7, 9, 3, 1, 5, 8, 4, 6}
	for i := 0; i < b.N; i++ {
		RobIIRecursive(nums)
	}
}