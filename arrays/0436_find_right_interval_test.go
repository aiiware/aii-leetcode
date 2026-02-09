package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRightInterval(t *testing.T) {
	tests := []struct {
		name     string
		intervals [][]int
		expected []int
	}{
		{
			name: "Example 1",
			intervals: [][]int{
				{1, 2},
			},
			expected: []int{-1},
		},
		{
			name: "Example 2",
			intervals: [][]int{
				{3, 4},
				{2, 3},
				{1, 2},
			},
			expected: []int{-1, 0, 1},
		},
		{
			name: "Example 3",
			intervals: [][]int{
				{1, 4},
				{2, 3},
				{3, 4},
			},
			expected: []int{-1, 2, -1},
		},
		{
			name: "Single interval",
			intervals: [][]int{
				{1, 3},
			},
			expected: []int{-1},
		},
		{
			name: "Two intervals, second is right interval",
			intervals: [][]int{
				{1, 3},
				{3, 5},
			},
			expected: []int{1, -1},
		},
		{
			name: "Two intervals, first is right interval",
			intervals: [][]int{
				{3, 5},
				{1, 3},
			},
			expected: []int{-1, 0},
		},
		{
			name: "Multiple intervals with exact matches",
			intervals: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			expected: []int{1, 2, 3, -1},
		},
		{
			name: "Intervals with gaps",
			intervals: [][]int{
				{1, 3},
				{5, 7},
				{9, 11},
			},
			expected: []int{1, 2, -1}, // [5,7] has start=5 ≥ 3, [9,11] has start=9 ≥ 7
		},
		{
			name: "Intervals with overlapping right intervals",
			intervals: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
				{4, 7},
			},
			expected: []int{3, -1, -1, -1}, // [4,7] has start=4 ≥ 4
		},
		{
			name: "Complex case 1",
			intervals: [][]int{
				{10, 20},
				{5, 15},
				{15, 25},
				{20, 30},
			},
			expected: []int{3, 2, -1, -1}, // Sorted: [5,15](1), [10,20](0), [15,25](2), [20,30](3)
		},
		{
			name: "Complex case 2",
			intervals: [][]int{
				{4, 5},
				{2, 3},
				{1, 2},
				{3, 4},
			},
			expected: []int{-1, 3, 1, 0},
		},
		{
			name: "Empty intervals",
			intervals: [][]int{},
			expected: []int{},
		},
		{
			name: "All intervals have right intervals",
			intervals: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
			},
			expected: []int{1, 2, -1},
		},
		{
			name: "Large start values",
			intervals: [][]int{
				{100, 200},
				{200, 300},
				{300, 400},
			},
			expected: []int{1, 2, -1},
		},
		{
			name: "Negative start values",
			intervals: [][]int{
				{-5, -3},
				{-3, -1},
				{-1, 1},
			},
			expected: []int{1, 2, -1},
		},
		{
			name: "Mixed positive and negative",
			intervals: [][]int{
				{-2, 0},
				{0, 2},
				{2, 4},
			},
			expected: []int{1, 2, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindRightInterval(tt.intervals)
			assert.Equal(t, tt.expected, result,
				"FindRightInterval(%v) = %v, expected %v",
				tt.intervals, result, tt.expected)
		})
	}
}

func TestFindRightIntervalWithMap(t *testing.T) {
	tests := []struct {
		name     string
		intervals [][]int
		expected []int
	}{
		{
			name: "Example 1",
			intervals: [][]int{
				{1, 2},
			},
			expected: []int{-1},
		},
		{
			name: "Example 2",
			intervals: [][]int{
				{3, 4},
				{2, 3},
				{1, 2},
			},
			expected: []int{-1, 0, 1},
		},
		{
			name: "Example 3",
			intervals: [][]int{
				{1, 4},
				{2, 3},
				{3, 4},
			},
			expected: []int{-1, 2, -1},
		},
		{
			name: "Complex case",
			intervals: [][]int{
				{4, 5},
				{2, 3},
				{1, 2},
				{3, 4},
			},
			expected: []int{-1, 3, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindRightIntervalWithMap(tt.intervals)
			assert.Equal(t, tt.expected, result,
				"FindRightIntervalWithMap(%v) = %v, expected %v",
				tt.intervals, result, tt.expected)
		})
	}
}

func TestFindRightInterval_Consistency(t *testing.T) {
	// Test that both implementations give the same result
	testCases := [][][]int{
		{{1, 2}},
		{{3, 4}, {2, 3}, {1, 2}},
		{{1, 4}, {2, 3}, {3, 4}},
		{{4, 5}, {2, 3}, {1, 2}, {3, 4}},
		{{10, 20}, {5, 15}, {15, 25}, {20, 30}},
		{{-5, -3}, {-3, -1}, {-1, 1}},
		{},
		{{1, 2}, {2, 3}, {3, 4}},
	}

	for _, intervals := range testCases {
		t.Run("", func(t *testing.T) {
			result1 := FindRightInterval(intervals)
			result2 := FindRightIntervalWithMap(intervals)
			assert.Equal(t, result1, result2,
				"Both implementations should give same result for %v: %v vs %v",
				intervals, result1, result2)
		})
	}
}

func TestFindRightInterval_EdgeCases(t *testing.T) {
	t.Run("Empty intervals", func(t *testing.T) {
		result := FindRightInterval([][]int{})
		assert.Empty(t, result)
	})

	t.Run("Single interval with no right interval", func(t *testing.T) {
		result := FindRightInterval([][]int{{1, 5}})
		assert.Equal(t, []int{-1}, result)
	})

	t.Run("Single interval with itself as right interval (start == end)", func(t *testing.T) {
		// Note: According to problem, right interval j must have start_j >= end_i
		// If start == end, it's a valid right interval
		result := FindRightInterval([][]int{{5, 5}})
		assert.Equal(t, []int{0}, result) // Interval itself is the right interval
	})

	t.Run("Multiple intervals with same start time (shouldn't happen per problem)", func(t *testing.T) {
		// Problem states each start_i is unique, but let's test robustness
		intervals := [][]int{
			{1, 3},
			{1, 4}, // Same start as first (invalid per problem)
			{2, 5},
		}
		// The map-based implementation will have issues with duplicate keys
		// The slice-based implementation should handle it
		result := FindRightInterval(intervals)
		// We can't assert exact values due to non-deterministic sorting of equal starts
		assert.Len(t, result, 3)
	})

	t.Run("Very large intervals", func(t *testing.T) {
		intervals := [][]int{
			{1, 1000000},
			{1000000, 2000000},
			{2000000, 3000000},
		}
		result := FindRightInterval(intervals)
		expected := []int{1, 2, -1}
		assert.Equal(t, expected, result)
	})
}

func BenchmarkFindRightInterval(b *testing.B) {
	// Create test data
	intervals := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		intervals[i] = []int{i * 2, i*2 + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindRightInterval(intervals)
	}
}

func BenchmarkFindRightIntervalWithMap(b *testing.B) {
	// Create test data
	intervals := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		intervals[i] = []int{i * 2, i*2 + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindRightIntervalWithMap(intervals)
	}
}

func BenchmarkFindRightInterval_WorstCase(b *testing.B) {
	// Worst case: all intervals need binary search
	intervals := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		// Each interval ends before the next starts
		intervals[i] = []int{i * 3, i*3 + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindRightInterval(intervals)
	}
}

func BenchmarkFindRightInterval_BestCase(b *testing.B) {
	// Best case: intervals are already sorted by start
	intervals := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		intervals[i] = []int{i, i + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindRightInterval(intervals)
	}
}