package arrays

import (
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		// Test case 1: Basic example from LeetCode
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true, // nums[0] == nums[3], distance = 3
		},
		// Test case 2: Same example but k too small
		{
			nums:     []int{1, 2, 3, 1},
			k:        2,
			expected: false, // distance = 3 > 2
		},
		// Test case 3: Multiple duplicates within range
		{
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true, // nums[2] == nums[3], distance = 1
		},
		// Test case 4: No duplicates
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: false,
		},
		// Test case 5: Duplicate at exact distance k
		{
			nums:     []int{1, 2, 3, 4, 1},
			k:        4,
			expected: true, // distance = 4 == k
		},
		// Test case 6: Empty array
		{
			nums:     []int{},
			k:        3,
			expected: false,
		},
		// Test case 7: Single element
		{
			nums:     []int{1},
			k:        0,
			expected: false, // k=0, need distinct indices
		},
		// Test case 8: k=0 always false (need distinct indices)
		{
			nums:     []int{1, 1, 1, 1},
			k:        0,
			expected: false,
		},
		// Test case 9: Negative k
		{
			nums:     []int{1, 2, 3, 1},
			k:        -1,
			expected: false,
		},
		// Test case 10: Large array with duplicates
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
			k:        10,
			expected: true, // distance = 10 == k
		},
		// Test case 11: Multiple duplicates, first one within range
		{
			nums:     []int{1, 2, 1, 3, 1, 4, 1},
			k:        2,
			expected: true, // nums[0] == nums[2], distance = 2
		},
		// Test case 12: All same numbers
		{
			nums:     []int{7, 7, 7, 7, 7},
			k:        1,
			expected: true,
		},
	}

	for i, test := range tests {
		// Test main implementation
		result := containsNearbyDuplicate(test.nums, test.k)
		if result != test.expected {
			t.Errorf("Test case %d failed (hash map): nums=%v, k=%d, expected=%v, got=%v",
				i+1, test.nums, test.k, test.expected, result)
		}

		// Test sliding window implementation
		windowResult := containsNearbyDuplicateSlidingWindow(test.nums, test.k)
		if windowResult != test.expected {
			t.Errorf("Test case %d failed (sliding window): nums=%v, k=%d, expected=%v, got=%v",
				i+1, test.nums, test.k, test.expected, windowResult)
		}

		// Test brute force for comparison (skip for large arrays)
		if len(test.nums) <= 20 {
			bruteResult := containsNearbyDuplicateBruteForce(test.nums, test.k)
			if bruteResult != test.expected {
				t.Errorf("Test case %d: Brute force mismatch: expected=%v, brute=%v",
					i+1, test.expected, bruteResult)
			}
		}

		// Ensure both implementations agree
		if result != windowResult {
			t.Errorf("Test case %d: Implementations disagree: hashmap=%v, window=%v",
				i+1, result, windowResult)
		}
	}
}

func BenchmarkContainsNearbyDuplicate(b *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100 // Many duplicates
	}
	k := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsNearbyDuplicate(nums, k)
	}
}

func BenchmarkContainsNearbyDuplicateSlidingWindow(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsNearbyDuplicateSlidingWindow(nums, k)
	}
}

func BenchmarkContainsNearbyDuplicateBruteForce(b *testing.B) {
	// Smaller test case for brute force (it's O(n*k))
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsNearbyDuplicateBruteForce(nums, k)
	}
}
