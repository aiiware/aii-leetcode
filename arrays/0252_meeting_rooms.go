package arrays

import "sort"

// CanAttendMeetings solves LeetCode problem 0252: Meeting Rooms
// Difficulty: Easy
// Tags: Array, Sorting, Intervals
//
// Given an array of meeting time intervals where intervals[i] = [starti, endi],
// determine if a person could attend all meetings.
//
// Time complexity: O(n log n), Space complexity: O(1) or O(n) for sorting
func CanAttendMeetings(intervals [][]int) bool {
	// Sort intervals by start time
	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	// Check for overlaps
	for i := 1; i < len(sorted); i++ {
		if sorted[i][0] < sorted[i-1][1] {
			// Current meeting starts before previous meeting ends
			return false
		}
	}

	return true
}
