package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		median float64
	}{
		{
			name:   "Example 1",
			nums1:  []int{1, 3},
			nums2:  []int{2},
			median: 2.0,
		},
		{
			name:   "Example 2",
			nums1:  []int{1, 2},
			nums2:  []int{3, 4},
			median: 2.5,
		},
		{
			name:   "Empty first array",
			nums1:  []int{},
			nums2:  []int{1, 2, 3, 4, 5},
			median: 3.0,
		},
		{
			name:   "Empty second array",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{},
			median: 3.0,
		},
		{
			name:   "Single element arrays",
			nums1:  []int{5},
			nums2:  []int{10},
			median: 7.5,
		},
		{
			name:   "Large numbers",
			nums1:  []int{100, 200, 300},
			nums2:  []int{150, 250, 350},
			median: 225.0,
		},
		{
			name:   "Negative numbers",
			nums1:  []int{-5, -3, -1},
			nums2:  []int{-4, -2, 0},
			median: -2.5,
		},
		{
			name:   "Mixed positive and negative",
			nums1:  []int{-10, -5, 0, 5, 10},
			nums2:  []int{-7, -3, 3, 7},
			median: 0.0,
		},
		{
			name:   "Different sizes",
			nums1:  []int{1, 2, 3, 4, 5},
			nums2:  []int{6, 7, 8, 9, 10, 11, 12},
			median: 6.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMedianSortedArrays(tt.nums1, tt.nums2)
			assert.InDelta(t, tt.median, result, 0.0001, "Median should match expected value")
		})
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	benchmarks := []struct {
		name  string
		nums1 []int
		nums2 []int
	}{
		{
			name:  "Small arrays",
			nums1: []int{1, 3, 5, 7, 9},
			nums2: []int{2, 4, 6, 8, 10},
		},
		{
			name:  "Medium arrays",
			nums1: makeRange(0, 100),
			nums2: makeRange(50, 150),
		},
		{
			name:  "Large arrays",
			nums1: makeRange(0, 1000),
			nums2: makeRange(500, 1500),
		},
		{
			name:  "Very different sizes",
			nums1: makeRange(0, 10),
			nums2: makeRange(0, 1000),
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMedianSortedArrays(bm.nums1, bm.nums2)
			}
		})
	}
}

// Helper function to create a range of integers
func makeRange(start, end int) []int {
	result := make([]int, end-start)
	for i := range result {
		result[i] = start + i
	}
	return result
}