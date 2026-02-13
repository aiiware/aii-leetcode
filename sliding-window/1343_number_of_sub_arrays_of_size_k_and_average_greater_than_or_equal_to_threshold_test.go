package sliding_window

import (
	"testing"
)

func TestNumOfSubarrays(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		k         int
		threshold int
		expected  int
	}{
		{
			name:      "Example 1 from LeetCode",
			arr:       []int{2, 2, 2, 2, 5, 5, 5, 8},
			k:         3,
			threshold: 4,
			expected:  3,
		},
		{
			name:      "Example 2 from LeetCode",
			arr:       []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3},
			k:         3,
			threshold: 5,
			expected:  6,
		},
		{
			name:      "All elements below threshold",
			arr:       []int{1, 2, 3, 4, 5},
			k:         3,
			threshold: 10,
			expected:  0,
		},
		{
			name:      "All elements above threshold",
			arr:       []int{10, 20, 30, 40, 50},
			k:         2,
			threshold: 5,
			expected:  4, // All 4 windows have average >= 5
		},
		{
			name:      "k equals array length",
			arr:       []int{10, 20, 30},
			k:         3,
			threshold: 20,
			expected:  1, // Average is 20, which equals threshold
		},
		{
			name:      "k = 1",
			arr:       []int{1, 2, 3, 4, 5},
			k:         1,
			threshold: 3,
			expected:  3, // Elements 3, 4, 5 are >= 3
		},
		{
			name:      "Single element with k=1",
			arr:       []int{10},
			k:         1,
			threshold: 10,
			expected:  1,
		},
		{
			name:      "Threshold = 0",
			arr:       []int{1, 2, 3, 4, 5},
			k:         3,
			threshold: 0,
			expected:  3, // All windows have average >= 0
		},
		{
			name:      "Mixed values",
			arr:       []int{1, 3, 5, 7, 9},
			k:         2,
			threshold: 4,
			expected:  3, // Windows [3,5], [5,7], [7,9] have average >= 4
		},
		{
			name:      "Large threshold",
			arr:       []int{100, 200, 300, 400, 500},
			k:         2,
			threshold: 350,
			expected:  2, // Windows [300,400], [400,500] have average >= 350
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numOfSubarrays(tt.arr, tt.k, tt.threshold)
			if result != tt.expected {
				t.Errorf("numOfSubarrays(%v, %d, %d) = %d, expected %d", tt.arr, tt.k, tt.threshold, result, tt.expected)
			}

			// Also test brute force for comparison (skip for large arrays)
			if len(tt.arr) <= 20 {
				bruteResult := numOfSubarraysBruteForce(tt.arr, tt.k, tt.threshold)
				if result != bruteResult {
					t.Errorf("Sliding window doesn't match brute force: sliding=%d, brute=%d", result, bruteResult)
				}
			}
		})
	}
}

func BenchmarkNumOfSubarrays(b *testing.B) {
	// Create a large test case
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i % 100
	}
	k := 100
	threshold := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numOfSubarrays(arr, k, threshold)
	}
}

func BenchmarkNumOfSubarraysBruteForce(b *testing.B) {
	// Smaller test case for brute force (it's O(n*k))
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i % 100
	}
	k := 10
	threshold := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numOfSubarraysBruteForce(arr, k, threshold)
	}
}