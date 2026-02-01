package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "k = 1 (largest element)",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        1,
			expected: 6,
		},
		{
			name:     "k = len(nums) (smallest element)",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        6,
			expected: 1,
		},
		{
			name:     "All same elements",
			nums:     []int{5, 5, 5, 5, 5},
			k:        3,
			expected: 5,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        2,
			expected: -2,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-10, 5, 0, -3, 8, 2},
			k:        3,
			expected: 2,
		},
		{
			name:     "Single element",
			nums:     []int{42},
			k:        1,
			expected: 42,
		},
		{
			name:     "Sorted ascending",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        5,
			expected: 6,
		},
		{
			name:     "Sorted descending",
			nums:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			k:        5,
			expected: 6,
		},
		{
			name:     "With duplicates",
			nums:     []int{2, 2, 2, 1, 1, 1, 3, 3, 3},
			k:        5,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result,
				"KthLargest(%v, %d) = %d, expected %d",
				tt.nums, tt.k, result, tt.expected)
		})
	}
}

func TestKthLargest_EdgeCases(t *testing.T) {
	t.Run("Empty array returns -1", func(t *testing.T) {
		result := KthLargest([]int{}, 1)
		assert.Equal(t, -1, result)
	})

	t.Run("k = 0 returns -1", func(t *testing.T) {
		result := KthLargest([]int{1, 2, 3}, 0)
		assert.Equal(t, -1, result)
	})

	t.Run("k > len(nums) returns -1", func(t *testing.T) {
		result := KthLargest([]int{1, 2, 3}, 5)
		assert.Equal(t, -1, result)
	})

	t.Run("k negative returns -1", func(t *testing.T) {
		result := KthLargest([]int{1, 2, 3}, -1)
		assert.Equal(t, -1, result)
	})

	t.Run("Large array", func(t *testing.T) {
		n := 1000
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = i + 1
		}
		k := 250
		expected := n - k + 1 // Since array is 1..n, kth largest is n-k+1
		
		result := KthLargest(nums, k)
		assert.Equal(t, expected, result)
	})

	t.Run("All zeros", func(t *testing.T) {
		result := KthLargest([]int{0, 0, 0, 0, 0}, 3)
		assert.Equal(t, 0, result)
	})
}

func BenchmarkKthLargest(b *testing.B) {
	// Create a large array for benchmarking
	n := 10000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KthLargest(nums, k)
	}
}