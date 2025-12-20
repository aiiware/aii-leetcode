package leetcode

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"Example 1", []int{-1, 2, 1, -4}, 1, 2},
		{"Example 2", []int{0, 0, 0}, 1, 0},
		{"Exact match", []int{1, 1, 1, 0}, -100, 2},
		{"Negative target", []int{-1, 2, 1, -4}, -100, -4},
		{"Simple negative", []int{-1, 0, 1}, 0, 0},
		{"Small array", []int{1, 1, 1}, 3, 3},
		{"Exact zero", []int{-1, 0, 1, 2}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ThreeSumClosest(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("ThreeSumClosest(%v, %d) = %d, expected %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func BenchmarkThreeSumClosest(b *testing.B) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	for i := 0; i < b.N; i++ {
		ThreeSumClosest(nums, target)
	}
}

func BenchmarkThreeSumClosestLarge(b *testing.B) {
	nums := make([]int, 100)
	for i := 0; i < 100; i++ {
		nums[i] = (i % 50) - 25
	}
	target := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ThreeSumClosest(nums, target)
	}
}
