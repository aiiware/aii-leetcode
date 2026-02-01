package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  bool
	}{
		{
			name:      "Example 1 - Can attend",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  false,
		},
		{
			name:      "Example 2 - Can attend",
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  true,
		},
		{
			name:      "Single meeting",
			intervals: [][]int{{1, 5}},
			expected:  true,
		},
		{
			name:      "Empty intervals",
			intervals: [][]int{},
			expected:  true,
		},
		{
			name:      "Back-to-back meetings",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			expected:  true,
		},
		{
			name:      "Overlapping at boundary",
			intervals: [][]int{{1, 3}, {3, 4}, {4, 5}},
			expected:  true,
		},
		{
			name:      "Multiple overlaps",
			intervals: [][]int{{1, 5}, {2, 3}, {4, 6}, {7, 8}},
			expected:  false,
		},
		{
			name:      "Nested intervals",
			intervals: [][]int{{1, 10}, {2, 5}, {6, 8}},
			expected:  false,
		},
		{
			name:      "Same start time",
			intervals: [][]int{{1, 3}, {1, 4}, {2, 5}},
			expected:  false,
		},
		{
			name:      "Large intervals",
			intervals: [][]int{{100, 200}, {150, 250}, {300, 400}},
			expected:  false,
		},
		{
			name:      "Negative times",
			intervals: [][]int{{-10, -5}, {-7, -3}, {-2, 0}},
			expected:  false,
		},
		{
			name:      "Zero duration meetings",
			intervals: [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanAttendMeetings(tt.intervals)
			assert.Equal(t, tt.expected, result,
				"CanAttendMeetings(%v) = %v, expected %v",
				tt.intervals, result, tt.expected)
		})
	}
}

func TestCanAttendMeetings_EdgeCases(t *testing.T) {
	t.Run("Very large number of meetings", func(t *testing.T) {
		intervals := make([][]int, 1000)
		for i := range intervals {
			// Create non-overlapping meetings
			intervals[i] = []int{i * 2, i*2 + 1}
		}
		result := CanAttendMeetings(intervals)
		assert.True(t, result)
	})

	t.Run("All meetings at same time", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {1, 3}, {1, 3}}
		result := CanAttendMeetings(intervals)
		assert.False(t, result)
	})

	t.Run("Meeting ends when next starts", func(t *testing.T) {
		intervals := [][]int{{1, 2}, {2, 3}, {3, 4}}
		result := CanAttendMeetings(intervals)
		assert.True(t, result, "Meetings that end exactly when next starts should be attendable")
	})

	t.Run("Single point in time meetings", func(t *testing.T) {
		intervals := [][]int{{1, 1}, {1, 1}, {2, 2}}
		result := CanAttendMeetings(intervals)
		assert.True(t, result, "Zero-duration meetings at same time don't overlap")
	})
}

func BenchmarkCanAttendMeetings(b *testing.B) {
	// Create a large set of intervals for benchmarking
	intervals := make([][]int, 10000)
	for i := range intervals {
		// Create some overlapping pattern
		if i%3 == 0 {
			intervals[i] = []int{i * 2, i*2 + 3}
		} else {
			intervals[i] = []int{i * 2, i*2 + 1}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanAttendMeetings(intervals)
	}
}

func BenchmarkCanAttendMeetings_Sorted(b *testing.B) {
	// Already sorted intervals (best case for algorithm after sorting)
	intervals := make([][]int, 10000)
	for i := range intervals {
		intervals[i] = []int{i, i + 1}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanAttendMeetings(intervals)
	}
}

func BenchmarkCanAttendMeetings_AllOverlapping(b *testing.B) {
	// All intervals overlap (worst case - needs to check all)
	intervals := make([][]int, 10000)
	for i := range intervals {
		intervals[i] = []int{0, 10}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanAttendMeetings(intervals)
	}
}