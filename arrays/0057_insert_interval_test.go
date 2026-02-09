package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name: "Example 1",
			intervals: [][]int{
				{1, 3},
				{6, 9},
			},
			newInterval: []int{2, 5},
			expected: [][]int{
				{1, 5},
				{6, 9},
			},
		},
		{
			name: "Example 2",
			intervals: [][]int{
				{1, 2},
				{3, 5},
				{6, 7},
				{8, 10},
				{12, 16},
			},
			newInterval: []int{4, 8},
			expected: [][]int{
				{1, 2},
				{3, 10},
				{12, 16},
			},
		},
		{
			name:        "Empty intervals",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected: [][]int{
				{5, 7},
			},
		},
		{
			name: "Insert at beginning",
			intervals: [][]int{
				{5, 7},
				{8, 10},
			},
			newInterval: []int{1, 3},
			expected: [][]int{
				{1, 3},
				{5, 7},
				{8, 10},
			},
		},
		{
			name: "Insert at end",
			intervals: [][]int{
				{1, 3},
				{5, 7},
			},
			newInterval: []int{8, 10},
			expected: [][]int{
				{1, 3},
				{5, 7},
				{8, 10},
			},
		},
		{
			name: "Merge with single interval",
			intervals: [][]int{
				{1, 5},
			},
			newInterval: []int{2, 7},
			expected: [][]int{
				{1, 7},
			},
		},
		{
			name: "Completely inside existing interval",
			intervals: [][]int{
				{1, 10},
			},
			newInterval: []int{3, 7},
			expected: [][]int{
				{1, 10},
			},
		},
		{
			name: "Overlap with multiple intervals",
			intervals: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
			newInterval: []int{2, 7},
			expected: [][]int{
				{1, 8},
			},
		},
		{
			name: "No overlap in middle",
			intervals: [][]int{
				{1, 2},
				{5, 6},
				{9, 10},
			},
			newInterval: []int{3, 4},
			expected: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{9, 10},
			},
		},
		{
			name: "Touching intervals merge",
			intervals: [][]int{
				{1, 2},
				{3, 4},
			},
			newInterval: []int{2, 3},
			expected: [][]int{
				{1, 4},
			},
		},
		{
			name: "Extend existing interval",
			intervals: [][]int{
				{1, 3},
				{6, 9},
			},
			newInterval: []int{3, 6},
			expected: [][]int{
				{1, 9},
			},
		},
		{
			name: "New interval covers all",
			intervals: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			newInterval: []int{0, 10},
			expected: [][]int{
				{0, 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Insert(tt.intervals, tt.newInterval)
			assert.Equal(t, tt.expected, result,
				"Insert(%v, %v) = %v, expected %v",
				tt.intervals, tt.newInterval, result, tt.expected)
		})
	}
}

func TestInsert_EdgeCases(t *testing.T) {
	t.Run("Negative intervals", func(t *testing.T) {
		intervals := [][]int{
			{-10, -5},
			{-3, -1},
		}
		newInterval := []int{-7, -2}
		expected := [][]int{
			{-10, -1},
		}
		result := Insert(intervals, newInterval)
		assert.Equal(t, expected, result)
	})

	t.Run("Single point interval insertion", func(t *testing.T) {
		intervals := [][]int{
			{1, 1},
			{3, 3},
			{5, 5},
		}
		newInterval := []int{2, 2}
		expected := [][]int{
			{1, 1},
			{2, 2},
			{3, 3},
			{5, 5},
		}
		result := Insert(intervals, newInterval)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge single point with interval", func(t *testing.T) {
		intervals := [][]int{
			{1, 3},
			{5, 7},
		}
		newInterval := []int{3, 5}
		expected := [][]int{
			{1, 7},
		}
		result := Insert(intervals, newInterval)
		assert.Equal(t, expected, result)
	})

	t.Run("Large intervals", func(t *testing.T) {
		intervals := [][]int{
			{100, 200},
			{300, 400},
			{500, 600},
		}
		newInterval := []int{250, 550}
		expected := [][]int{
			{100, 200},
			{250, 600},
		}
		result := Insert(intervals, newInterval)
		assert.Equal(t, expected, result)
	})
}

func BenchmarkInsert(b *testing.B) {
	// Create a large set of intervals for benchmarking
	intervals := make([][]int, 10000)
	for i := 0; i < 10000; i++ {
		start := i * 10
		end := start + 5
		intervals[i] = []int{start, end}
	}
	newInterval := []int{25000, 35000}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Insert(intervals, newInterval)
	}
}
