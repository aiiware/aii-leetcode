package arrays

import "sort"

// Interval represents a closed interval [start, end]
type Interval struct {
	Start int
	End   int
}

// Merge solves LeetCode problem 0056: Merge Intervals
// Difficulty: Medium
// Tags: Array, Sorting, Interval
//
// Given an array of intervals where intervals[i] = [starti, endi],
// merge all overlapping intervals, and return an array of the
// non-overlapping intervals that cover all the intervals in the input.
//
// Example 1:
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//
// Example 2:
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
// Time complexity: O(n log n), Space complexity: O(n) for sorting
func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Convert to Interval struct for easier sorting
	intervalList := make([]Interval, len(intervals))
	for i, interval := range intervals {
		intervalList[i] = Interval{Start: interval[0], End: interval[1]}
	}

	// Sort intervals by start time
	sort.Slice(intervalList, func(i, j int) bool {
		return intervalList[i].Start < intervalList[j].Start
	})

	// Merge intervals
	result := [][]int{}
	current := intervalList[0]

	for i := 1; i < len(intervalList); i++ {
		next := intervalList[i]

		// If intervals overlap, merge them
		if current.End >= next.Start {
			// Update the end to the maximum of both ends
			if next.End > current.End {
				current.End = next.End
			}
		} else {
			// No overlap, add current to result and move to next
			result = append(result, []int{current.Start, current.End})
			current = next
		}
	}

	// Add the last interval
	result = append(result, []int{current.Start, current.End})

	return result
}
