package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "With zeros",
			nums:     []int{0, 2, 3, 4},
			expected: []int{24, 0, 0, 0},
		},
		{
			name:     "Multiple zeros",
			nums:     []int{0, 0, 2, 3},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-2, -3, -4},
			expected: []int{12, 8, 6},
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-2, 3, -4, 5},
			expected: []int{-60, 40, -30, 24},
		},
		{
			name:     "Large numbers",
			nums:     []int{10, 20, 30, 40},
			expected: []int{24000, 12000, 8000, 6000},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProductExceptSelf(tt.nums)
			assert.Equal(t, tt.expected, result,
				"ProductExceptSelf(%v) = %v, expected %v",
				tt.nums, result, tt.expected)
		})
	}
}

func TestProductExceptSelf_EdgeCases(t *testing.T) {
	t.Run("All zeros except one", func(t *testing.T) {
		result := ProductExceptSelf([]int{0, 0, 5, 0})
		assert.Equal(t, []int{0, 0, 0, 0}, result)
	})

	t.Run("Single zero", func(t *testing.T) {
		result := ProductExceptSelf([]int{0})
		assert.Equal(t, []int{1}, result)
	})

	t.Run("Large array with pattern", func(t *testing.T) {
		n := 100
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = i + 1
		}

		// Calculate expected using brute force for verification
		expected := make([]int, n)
		for i := 0; i < n; i++ {
			product := 1
			for j := 0; j < n; j++ {
				if j != i {
					product *= nums[j]
				}
			}
			expected[i] = product
		}

		result := ProductExceptSelf(nums)
		assert.Equal(t, expected, result)
	})

	t.Run("Array with 1 and -1 only", func(t *testing.T) {
		result := ProductExceptSelf([]int{1, -1, 1, -1, 1})
		// Correct calculation: [1, -1, 1, -1, 1]
		// For index 0: (-1) * 1 * (-1) * 1 = 1
		// For index 1: 1 * 1 * (-1) * 1 = -1
		// For index 2: 1 * (-1) * (-1) * 1 = 1
		// For index 3: 1 * (-1) * 1 * 1 = -1
		// For index 4: 1 * (-1) * 1 * (-1) = 1
		assert.Equal(t, []int{1, -1, 1, -1, 1}, result)
	})
}

func BenchmarkProductExceptSelf(b *testing.B) {
	// Create a large array for benchmarking
	n := 10000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = (i % 10) + 1 // Values 1-10
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProductExceptSelf(nums)
	}
}