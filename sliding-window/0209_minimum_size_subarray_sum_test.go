package sliding_window

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		// Test case 1: Basic example from LeetCode
		{
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2, // [4, 3]
		},
		// Test case 2: Single element equals target
		{
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1, // [4]
		},
		// Test case 3: Entire array needed
		{
			target:   11,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 3, // [3, 4, 5] or [2, 3, 4, 5] but minimum is 3
		},
		// Test case 4: No valid subarray
		{
			target:   100,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		// Test case 5: First element alone satisfies
		{
			target:   5,
			nums:     []int{5, 1, 1, 1, 1},
			expected: 1,
		},
		// Test case 6: Last element alone satisfies
		{
			target:   5,
			nums:     []int{1, 1, 1, 1, 5},
			expected: 1,
		},
		// Test case 7: Multiple possible subarrays
		{
			target:   6,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 2, // [2, 4] or [3, 4] or [1, 5] etc.
		},
		// Test case 8: Empty array
		{
			target:   5,
			nums:     []int{},
			expected: 0,
		},
		// Test case 9: All elements needed
		{
			target:   15,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		// Test case 10: Large numbers
		{
			target:   80,
			nums:     []int{10, 20, 30, 40, 50},
			expected: 2, // [30, 50] or [40, 50]
		},
		// Test case 11: Zero target (edge case)
		{
			target:   0,
			nums:     []int{1, 2, 3},
			expected: 1, // Any single element satisfies sum >= 0
		},
		// Test case 12: All zeros
		{
			target:   0,
			nums:     []int{0, 0, 0, 0},
			expected: 1,
		},
	}

	for i, test := range tests {
		result := minSubArrayLen(test.target, test.nums)
		if result != test.expected {
			t.Errorf("Test case %d failed: target=%d, nums=%v, expected=%d, got=%d",
				i+1, test.target, test.nums, test.expected, result)
		}

		// Also test brute force for comparison
		bruteResult := minSubArrayLenBruteForce(test.target, test.nums)
		if result != bruteResult {
			t.Errorf("Test case %d: Sliding window doesn't match brute force: sliding=%d, brute=%d",
				i+1, result, bruteResult)
		}
	}
}

func TestMinSubArrayLenEdgeCases(t *testing.T) {
	// Test with nil slice
	result := minSubArrayLen(5, nil)
	if result != 0 {
		t.Errorf("Expected 0 for nil slice, got %d", result)
	}

	// Test with single element that satisfies
	result = minSubArrayLen(3, []int{5})
	if result != 1 {
		t.Errorf("Expected 1 for single element that satisfies, got %d", result)
	}

	// Test with single element that doesn't satisfy
	result = minSubArrayLen(10, []int{5})
	if result != 0 {
		t.Errorf("Expected 0 for single element that doesn't satisfy, got %d", result)
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = (i % 10) + 1 // Values 1-10
	}
	target := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minSubArrayLen(target, nums)
	}
}

func BenchmarkMinSubArrayLenBruteForce(b *testing.B) {
	// Smaller test case for brute force (it's O(n^2))
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i % 10) + 1
	}
	target := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minSubArrayLenBruteForce(target, nums)
	}
}