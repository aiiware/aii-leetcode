package arrays

import (
	"container/heap"
	"sort"
)

// MinHeap implements heap.Interface for a min-heap of integers
type MinHeapII []int

func (h MinHeapII) Len() int           { return len(h) }
func (h MinHeapII) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapII) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapII) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapII) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinMeetingRooms solves LeetCode problem 0253: Meeting Rooms II
// Difficulty: Medium
// Tags: Array, Sorting, Heap, Greedy, Intervals
//
// Given an array of meeting time intervals where intervals[i] = [starti, endi],
// return the minimum number of conference rooms required.
//
// Time complexity: O(n log n), Space complexity: O(n)
func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals by start time
	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	// Min-heap to track end times of meetings in progress
	h := &MinHeapII{}
	heap.Init(h)

	// Add first meeting's end time
	heap.Push(h, sorted[0][1])

	// Process remaining meetings
	for i := 1; i < len(sorted); i++ {
		start, end := sorted[i][0], sorted[i][1]

		// If the earliest ending meeting ends before or when this meeting starts,
		// we can reuse that room
		if (*h)[0] <= start {
			heap.Pop(h)
		}

		// Add current meeting's end time to heap
		heap.Push(h, end)
	}

	return h.Len()
}
