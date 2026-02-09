package arrays

import (
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		// Test case 1: Basic example
		{
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		// Test case 2: Single element window
		{
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        1,
			expected: []int{1, 3, -1, -3, 5, 3, 6, 7},
		},
		// Test case 3: Window size equals array length
		{
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        8,
			expected: []int{7},
		},
		// Test case 4: Decreasing sequence
		{
			nums:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:        3,
			expected: []int{9, 8, 7, 6, 5, 4, 3},
		},
		// Test case 5: Increasing sequence
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:        3,
			expected: []int{3, 4, 5, 6, 7, 8, 9},
		},
		// Test case 6: Empty array
		{
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
		// Test case 7: Single element array
		{
			nums:     []int{5},
			k:        1,
			expected: []int{5},
		},
		// Test case 8: All same elements
		{
			nums:     []int{2, 2, 2, 2, 2, 2},
			k:        3,
			expected: []int{2, 2, 2, 2},
		},
		// Test case 9: Mixed positive and negative
		{
			nums:     []int{-7, -8, 7, 5, 7, 1, 6, 0},
			k:        4,
			expected: []int{7, 7, 7, 7, 7},
		},
	}

	for i, test := range tests {
		result := maxSlidingWindow(test.nums, test.k)
		if !equalSlices(result, test.expected) {
			t.Errorf("Test case %d failed: nums=%v, k=%d, expected=%v, got=%v",
				i+1, test.nums, test.k, test.expected, result)
		}

		// Also test brute force for comparison (on small inputs only)
		if len(test.nums) <= 20 {
			bruteResult := maxSlidingWindowBruteForce(test.nums, test.k)
			if !equalSlices(result, bruteResult) {
				t.Errorf("Test case %d: Deque solution doesn't match brute force: deque=%v, brute=%v",
					i+1, result, bruteResult)
			}
		}
	}
}

func TestMaxSlidingWindowEdgeCases(t *testing.T) {
	// Test with k = 0
	result := maxSlidingWindow([]int{1, 2, 3}, 0)
	if len(result) != 0 {
		t.Errorf("Expected empty result for k=0, got %v", result)
	}

	// Test with nil slice
	result = maxSlidingWindow(nil, 3)
	if len(result) != 0 {
		t.Errorf("Expected empty result for nil slice, got %v", result)
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindow(nums, k)
	}
}

func BenchmarkMaxSlidingWindowBruteForce(b *testing.B) {
	// Smaller test case for brute force (it's O(n*k))
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSlidingWindowBruteForce(nums, k)
	}
}
