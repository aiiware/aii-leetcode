package arrays

// Insert solves LeetCode problem 0057: Insert Interval
// Difficulty: Medium
// Tags: Array, Interval
//
// You are given an array of non-overlapping intervals where intervals[i] = [starti, endi]
// represent the start and the end of the ith interval and intervals is sorted in
// ascending order by starti. You are also given an interval newInterval = [start, end]
// that represents the start and end of another interval.
//
// Insert newInterval into intervals such that intervals is still sorted in ascending
// order by starti and intervals still does not have any overlapping intervals
// (merge overlapping intervals if necessary).
//
// Return intervals after the insertion.
//
// Example 1:
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
//
// Example 2:
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
//
// Time complexity: O(n), Space complexity: O(n)
func Insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)

	// Add all intervals that come before the new interval
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Merge all overlapping intervals
	mergedStart := newInterval[0]
	mergedEnd := newInterval[1]

	for i < n && intervals[i][0] <= newInterval[1] {
		// Update the merged interval boundaries
		if intervals[i][0] < mergedStart {
			mergedStart = intervals[i][0]
		}
		if intervals[i][1] > mergedEnd {
			mergedEnd = intervals[i][1]
		}
		i++
	}

	// Add the merged interval
	result = append(result, []int{mergedStart, mergedEnd})

	// Add all remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
