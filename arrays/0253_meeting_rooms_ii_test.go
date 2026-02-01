package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  2,
		},
		{
			name:      "Example 2",
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  1,
		},
		{
			name:      "Single meeting",
			intervals: [][]int{{1, 5}},
			expected:  1,
		},
		{
			name:      "Empty intervals",
			intervals: [][]int{},
			expected:  0,
		},
		{
			name:      "Back-to-back meetings",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			expected:  1,
		},
		{
			name:      "All meetings overlap",
			intervals: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}},
			expected:  4,
		},
		{
			name:      "Nested intervals",
			intervals: [][]int{{1, 10}, {2, 5}, {6, 8}, {11, 12}},
			expected:  2,
		},
		{
			name:      "Same start time",
			intervals: [][]int{{1, 3}, {1, 4}, {1, 5}},
			expected:  3,
		},
		{
			name:      "Mixed pattern",
			intervals: [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}},
			expected:  2,
		},
		{
			name:      "Zero duration meetings",
			intervals: [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected:  1,
		},
		{
			name:      "Complex overlapping",
			intervals: [][]int{{1, 4}, {2, 5}, {3, 6}, {5, 7}, {6, 8}, {7, 9}},
			expected:  3,
		},
		{
			name:      "Large gaps between meetings",
			intervals: [][]int{{1, 2}, {100, 101}, {200, 201}},
			expected:  1,
		},
		{
			name:      "Negative times",
			intervals: [][]int{{-10, -5}, {-7, -3}, {-2, 0}},
			expected:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinMeetingRooms(tt.intervals)
			assert.Equal(t, tt.expected, result,
				"MinMeetingRooms(%v) = %d, expected %d",
				tt.intervals, result, tt.expected)
		})
	}
}

func TestMinMeetingRooms_EdgeCases(t *testing.T) {
	t.Run("Meeting ends when next starts", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {2, 3}, {3, 4}}
		result := MinMeetingRooms(intervals)
		assert.Equal(t, 1, result, "Should be able to reuse room when meeting ends exactly when next starts")
	})

	t.Run("All meetings at same time", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {1, 3}, {1, 3}}
		result := MinMeetingRooms(intervals)
		assert.Equal(t, 3, result)
	})

	t.Run("Single point in time meetings", func(t *testing.T) {
		intervals := [][]int{{1, 1}, {1, 1}, {2, 2}}
		result := MinMeetingRooms(intervals)
		assert.Equal(t, 1, result, "Zero-duration meetings at same time don't need separate rooms")
	})

	t.Run("Very large number of meetings", func(t *testing.T) {
		intervals := make([][]int, 1000)
		for i := range intervals {
			// Create pattern: every 3 meetings overlap
			if i%3 == 0 {
				intervals[i] = []int{i, i + 5}
			} else {
				intervals[i] = []int{i * 10, i*10 + 1}
			}
		}
		result := MinMeetingRooms(intervals)
		// Should need at most 3 rooms for the overlapping pattern
		assert.True(t, result >= 1 && result <= 3)
	})
}

func BenchmarkMinMeetingRooms(b *testing.B) {
	// Create a large set of intervals for benchmarking
	intervals := make([][]int, 10000)
	for i := range intervals {
		// Create overlapping pattern
		if i%4 == 0 {
			intervals[i] = []int{i, i + 10}
		} else {
			intervals[i] = []int{i * 2, i*2 + 1}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinMeetingRooms(intervals)
	}
}

func BenchmarkMinMeetingRooms_AllNonOverlapping(b *testing.B) {
	// All intervals are non-overlapping (best case)
	intervals := make([][]int, 10000)
	for i := range intervals {
		intervals[i] = []int{i * 2, i*2 + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinMeetingRooms(intervals)
	}
}

func BenchmarkMinMeetingRooms_AllOverlapping(b *testing.B) {
	// All intervals overlap (worst case)
	intervals := make([][]int, 10000)
	for i := range intervals {
		intervals[i] = []int{0, 10}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinMeetingRooms(intervals)
	}
}

func BenchmarkMinMeetingRooms_Small(b *testing.B) {
	// Small number of intervals
	intervals := [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinMeetingRooms(intervals)
	}
}