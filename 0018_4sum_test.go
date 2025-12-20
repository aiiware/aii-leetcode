package leetcode

import (
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			"Example 1",
			[]int{1000000000, 1000000000, 1000000000, 1000000000},
			-294967296,
			[][]int{},
		},
		{
			"Simple case",
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			"All same",
			[]int{2, 2, 2, 2, 2},
			8,
			[][]int{{2, 2, 2, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FourSum(tt.nums, tt.target)
			if !quadrupletSlicesEqual(result, tt.expected) {
				t.Errorf("FourSum(%v, %d) = %v, expected %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

// Helper function to compare quadruplet slices
func quadrupletSlicesEqual(a, b [][]int) bool {
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

func BenchmarkFourSum(b *testing.B) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	for i := 0; i < b.N; i++ {
		FourSum(nums, target)
	}
}

func BenchmarkFourSumLarge(b *testing.B) {
	nums := make([]int, 50)
	for i := 0; i < 50; i++ {
		nums[i] = (i % 25) - 12
	}
	target := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FourSum(nums, target)
	}
}
