package arrays

import "sort"

// FindRightInterval solves LeetCode problem 436: Find Right Interval
// Difficulty: Medium
// Tags: Array, Binary Search, Sorting
//
// You are given an array of intervals where intervals[i] = [start_i, end_i]
// and each start_i is unique.
//
// The right interval for an interval i is an interval j such that
// start_j >= end_i and start_j is minimized.
//
// Return an array of indices representing the right interval for each interval i.
// If no right interval exists for interval i, then put -1 at index i.
//
// Time complexity: O(n log n), Space complexity: O(n)
func FindRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}

	// Create a slice of interval indices with their start times
	type intervalInfo struct {
		index int
		start int
	}

	intervalInfos := make([]intervalInfo, n)
	for i, interval := range intervals {
		intervalInfos[i] = intervalInfo{
			index: i,
			start: interval[0],
		}
	}

	// Sort by start time
	sort.Slice(intervalInfos, func(i, j int) bool {
		return intervalInfos[i].start < intervalInfos[j].start
	})

	// Extract sorted start times for binary search
	starts := make([]int, n)
	indexMap := make([]int, n) // Maps from sorted position to original index
	for i, info := range intervalInfos {
		starts[i] = info.start
		indexMap[i] = info.index
	}

	// Result array
	result := make([]int, n)

	// For each interval, find the right interval using binary search
	for i, interval := range intervals {
		end := interval[1]

		// Binary search for the smallest start >= end
		pos := sort.Search(n, func(j int) bool {
			return starts[j] >= end
		})

		if pos < n {
			// Found a right interval
			result[i] = indexMap[pos]
		} else {
			// No right interval found
			result[i] = -1
		}
	}

	return result
}

// FindRightIntervalWithMap solves the same problem using a map for O(1) lookup
// after sorting. This is more efficient when we need to map back to original indices.
func FindRightIntervalWithMap(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}

	// Map from start value to original index
	startToIndex := make(map[int]int, n)
	starts := make([]int, n)

	for i, interval := range intervals {
		start := interval[0]
		startToIndex[start] = i
		starts[i] = start
	}

	// Sort starts for binary search
	sort.Ints(starts)

	// Result array
	result := make([]int, n)

	// For each interval, find the right interval
	for i, interval := range intervals {
		end := interval[1]

		// Binary search for the smallest start >= end
		pos := sort.SearchInts(starts, end)

		if pos < n {
			// Found a right interval, get its original index
			rightStart := starts[pos]
			result[i] = startToIndex[rightStart]
		} else {
			// No right interval found
			result[i] = -1
		}
	}

	return result
}