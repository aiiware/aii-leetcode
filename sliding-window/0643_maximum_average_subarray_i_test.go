package sliding_window

import (
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{5},
			k:        1,
			expected: 5.0,
		},
		{
			name:     "All positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: 4.0, // (3+4+5)/3 = 12/3 = 4
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        2,
			expected: -1.5, // (-1 + -2)/2 = -3/2 = -1.5
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-1, 2, -3, 4, -5, 6},
			k:        3,
			expected: 1.6666666666666667, // (4 + -5 + 6)/3 = 5/3 ≈ 1.6667
		},
		{
			name:     "k equals array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: 3.0, // (1+2+3+4+5)/5 = 15/5 = 3
		},
		{
			name:     "k = 1",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 5.0, // Maximum single element is 5
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        0,
			expected: 0.0,
		},
		{
			name:     "Single element with k=1",
			nums:     []int{10},
			k:        1,
			expected: 10.0,
		},
		{
			name:     "Large numbers",
			nums:     []int{10000, 20000, 30000, 40000},
			k:        2,
			expected: 35000.0, // (30000+40000)/2 = 70000/2 = 35000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxAverage(tt.nums, tt.k)
			// Compare with tolerance for floating point errors
			if !floatEquals(result, tt.expected, 1e-5) {
				t.Errorf("findMaxAverage(%v, %d) = %f, expected %f", tt.nums, tt.k, result, tt.expected)
			}

			// Also test brute force for comparison (skip for large arrays)
			if len(tt.nums) <= 20 {
				bruteResult := findMaxAverageBruteForce(tt.nums, tt.k)
				if !floatEquals(result, bruteResult, 1e-5) {
					t.Errorf("Sliding window doesn't match brute force: sliding=%f, brute=%f", result, bruteResult)
				}
			}
		})
	}
}

func floatEquals(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}

func BenchmarkFindMaxAverage(b *testing.B) {
	// Create a large test case
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMaxAverage(nums, k)
	}
}

func BenchmarkFindMaxAverageBruteForce(b *testing.B) {
	// Smaller test case for brute force (it's O(n*k))
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i % 100
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMaxAverageBruteForce(nums, k)
	}
}