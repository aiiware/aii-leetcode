package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			"Example 1",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"Example 2",
			[]int{0},
			[][]int{},
		},
		{
			"All zeros",
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			"Negative numbers",
			[]int{-2, 0, 1, 1, 2},
			[][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			"With duplicates",
			[]int{-1, 0, 1, 2, -1, -4, -1},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"Single triplet",
			[]int{-1, 0, 1},
			[][]int{{-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ThreeSum(tt.nums)
			if !tripletSlicesEqual(result, tt.expected) {
				t.Errorf("ThreeSum(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// Helper function to compare triplet slices
func tripletSlicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func BenchmarkThreeSum(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	for i := 0; i < b.N; i++ {
		ThreeSum(nums)
	}
}

func BenchmarkThreeSumLarge(b *testing.B) {
	// Create a larger test case
	nums := make([]int, 100)
	for i := 0; i < 100; i++ {
		nums[i] = (i % 50) - 25
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSum(nums)
	}
}
