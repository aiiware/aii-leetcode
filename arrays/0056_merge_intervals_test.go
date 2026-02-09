package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name: "Example 1",
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			expected: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "Example 2",
			intervals: [][]int{
				{1, 4},
				{4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
		{
			name:      "Empty intervals",
			intervals: [][]int{},
			expected:  [][]int{},
		},
		{
			name: "Single interval",
			intervals: [][]int{
				{1, 3},
			},
			expected: [][]int{
				{1, 3},
			},
		},
		{
			name: "All overlapping",
			intervals: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
				{4, 7},
			},
			expected: [][]int{
				{1, 7},
			},
		},
		{
			name: "No overlapping",
			intervals: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
			expected: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
		},
		{
			name: "Multiple merges",
			intervals: [][]int{
				{1, 3},
				{2, 4},
				{5, 7},
				{6, 8},
				{10, 12},
			},
			expected: [][]int{
				{1, 4},
				{5, 8},
				{10, 12},
			},
		},
		{
			name: "Unsorted intervals",
			intervals: [][]int{
				{15, 18},
				{8, 10},
				{2, 6},
				{1, 3},
			},
			expected: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "Complete overlap",
			intervals: [][]int{
				{1, 10},
				{2, 3},
				{4, 5},
				{6, 7},
			},
			expected: [][]int{
				{1, 10},
			},
		},
		{
			name: "Edge case: touching intervals",
			intervals: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
		{
			name: "Large intervals",
			intervals: [][]int{
				{1, 100},
				{50, 150},
				{200, 300},
				{250, 350},
			},
			expected: [][]int{
				{1, 150},
				{200, 350},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.intervals)
			assert.Equal(t, tt.expected, result,
				"Merge(%v) = %v, expected %v",
				tt.intervals, result, tt.expected)
		})
	}
}

func TestMerge_EdgeCases(t *testing.T) {
	t.Run("Negative intervals", func(t *testing.T) {
		intervals := [][]int{
			{-10, -5},
			{-7, -3},
			{-2, 0},
		}
		expected := [][]int{
			{-10, -3},
			{-2, 0},
		}
		result := Merge(intervals)
		assert.Equal(t, expected, result)
	})

	t.Run("Mixed positive negative", func(t *testing.T) {
		intervals := [][]int{
			{-5, 0},
			{-1, 3},
			{2, 5},
			{4, 7},
		}
		expected := [][]int{
			{-5, 7},
		}
		result := Merge(intervals)
		assert.Equal(t, expected, result)
	})

	t.Run("Single point intervals", func(t *testing.T) {
		intervals := [][]int{
			{1, 1},
			{2, 2},
			{3, 3},
		}
		expected := [][]int{
			{1, 1},
			{2, 2},
			{3, 3},
		}
		result := Merge(intervals)
		assert.Equal(t, expected, result)
	})

	t.Run("Overlapping single points", func(t *testing.T) {
		intervals := [][]int{
			{1, 1},
			{1, 1},
			{2, 2},
		}
		expected := [][]int{
			{1, 1},
			{2, 2},
		}
		result := Merge(intervals)
		assert.Equal(t, expected, result)
	})
}

func BenchmarkMerge(b *testing.B) {
	// Create a large set of intervals for benchmarking
	intervals := make([][]int, 10000)
	for i := 0; i < 10000; i++ {
		start := i * 2
		end := start + (i % 5) + 1
		intervals[i] = []int{start, end}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Merge(intervals)
	}
}
