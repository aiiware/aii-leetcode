package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		target        int
		validResults  [][]int // Multiple valid results if algorithm can return different valid pairs
	}{
		{
			name:         "Example 1",
			nums:         []int{2, 7, 11, 15},
			target:       9,
			validResults: [][]int{{0, 1}},
		},
		{
			name:         "Example 2",
			nums:         []int{3, 2, 4},
			target:       6,
			validResults: [][]int{{1, 2}},
		},
		{
			name:         "Example 3",
			nums:         []int{3, 3},
			target:       6,
			validResults: [][]int{{0, 1}},
		},
		{
			name:         "Negative numbers",
			nums:         []int{-1, -2, -3, -4, -5},
			target:       -8,
			validResults: [][]int{{2, 4}},
		},
		{
			name:         "Mixed positive and negative",
			nums:         []int{-10, 7, 19, 15},
			target:       9,
			validResults: [][]int{{0, 2}},
		},
		{
			name:         "Zero target",
			nums:         []int{0, 4, 3, 0},
			target:       0,
			validResults: [][]int{{0, 3}},
		},
		{
			name:  "Large numbers",
			nums:  []int{1000000, 500000, 1500000},
			target: 2000000,
			// Both [0, 2] and [1, 2] are valid solutions
			validResults: [][]int{{0, 2}, {1, 2}},
		},
		{
			name:         "Duplicate numbers different indices",
			nums:         []int{1, 2, 3, 1, 5},
			target:       4,
			validResults: [][]int{{0, 2}, {2, 3}}, // Both [0, 2] and [2, 3] are valid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			
			// Check if result is one of the valid results
			found := false
			for _, validResult := range tt.validResults {
				if len(result) == 2 && len(validResult) == 2 &&
					result[0] == validResult[0] && result[1] == validResult[1] {
					found = true
					break
				}
			}
			
			assert.True(t, found, 
				"TwoSum(%v, %d) = %v, expected one of %v", 
				tt.nums, tt.target, result, tt.validResults)
			
			// Verify the result is mathematically correct
			if len(result) == 2 {
				assert.Equal(t, tt.target, tt.nums[result[0]] + tt.nums[result[1]],
					"Indices %v: %d + %d should equal %d",
					result, tt.nums[result[0]], tt.nums[result[1]], tt.target)
			}
		})
	}
}

func TestTwoSum_EdgeCases(t *testing.T) {
	t.Run("Empty slice returns empty", func(t *testing.T) {
		result := TwoSum([]int{}, 0)
		assert.Empty(t, result)
	})

	t.Run("Single element returns empty", func(t *testing.T) {
		result := TwoSum([]int{1}, 2)
		assert.Empty(t, result)
	})

	t.Run("No solution returns empty", func(t *testing.T) {
		// Note: According to problem, exactly one solution exists
		// This test is just to verify our implementation handles it gracefully
		result := TwoSum([]int{1, 2, 3}, 100)
		assert.Empty(t, result)
	})
}

func BenchmarkTwoSum(b *testing.B) {
	// Create a large slice for benchmarking
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	target := 19997 // Should find indices 9998 and 9999

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}