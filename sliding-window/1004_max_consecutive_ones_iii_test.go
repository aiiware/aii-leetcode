package sliding_window

import (
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			k:        0,
			expected: 5,
		},
		{
			name:     "All zeros with k=0",
			nums:     []int{0, 0, 0, 0, 0},
			k:        0,
			expected: 0,
		},
		{
			name:     "All zeros with k=5",
			nums:     []int{0, 0, 0, 0, 0},
			k:        5,
			expected: 5,
		},
		{
			name:     "Mixed with k=0",
			nums:     []int{1, 0, 1, 0, 1},
			k:        0,
			expected: 1,
		},
		{
			name:     "Mixed with k=1",
			nums:     []int{1, 0, 1, 0, 1},
			k:        1,
			expected: 3,
		},
		{
			name:     "Mixed with k=2",
			nums:     []int{1, 0, 1, 0, 1},
			k:        2,
			expected: 5,
		},
		{
			name:     "Single element 1 with k=0",
			nums:     []int{1},
			k:        0,
			expected: 1,
		},
		{
			name:     "Single element 0 with k=0",
			nums:     []int{0},
			k:        0,
			expected: 0,
		},
		{
			name:     "Single element 0 with k=1",
			nums:     []int{0},
			k:        1,
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        0,
			expected: 0,
		},
		{
			name:     "k larger than array length",
			nums:     []int{0, 0, 0},
			k:        5,
			expected: 3,
		},
		{
			name:     "Alternating pattern",
			nums:     []int{1, 0, 1, 0, 1, 0, 1},
			k:        2,
			expected: 5,
		},
		{
			name:     "Long sequence of ones with zeros in between",
			nums:     []int{1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1},
			k:        2,
			expected: 11, // Can flip both zeros to get all ones
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestOnes(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("longestOnes(%v, %d) = %d, expected %d", tt.nums, tt.k, result, tt.expected)
			}

			// Also test optimized version for comparison
			optimizedResult := longestOnesOptimized(tt.nums, tt.k)
			if result != optimizedResult {
				t.Errorf("Standard and optimized versions don't match: standard=%d, optimized=%d", result, optimizedResult)
			}

			// Also test brute force for comparison (skip for large arrays)
			if len(tt.nums) <= 20 {
				bruteResult := longestOnesBruteForce(tt.nums, tt.k)
				if result != bruteResult {
					t.Errorf("Sliding window doesn't match brute force: sliding=%d, brute=%d", result, bruteResult)
				}
			}
		})
	}
}

func TestLongestOnesOptimized(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			k:        0,
			expected: 5,
		},
		{
			name:     "All zeros with k=5",
			nums:     []int{0, 0, 0, 0, 0},
			k:        5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestOnesOptimized(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("longestOnesOptimized(%v, %d) = %d, expected %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestOnes(b *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 2 // Alternating 0 and 1
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestOnes(nums, k)
	}
}

func BenchmarkLongestOnesOptimized(b *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 2 // Alternating 0 and 1
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestOnesOptimized(nums, k)
	}
}

func BenchmarkLongestOnesBruteForce(b *testing.B) {
	// Smaller test case for brute force (it's O(n^2))
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i % 2
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestOnesBruteForce(nums, k)
	}
}