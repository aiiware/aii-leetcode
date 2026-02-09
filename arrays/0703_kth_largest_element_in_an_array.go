package arrays

import (
	"container/heap"
)

// KthLargest solves LeetCode problem 0703: Kth Largest Element in an Array
// Difficulty: Medium
// Tags: Array, Divide and Conquer, Heap (Priority Queue), Quickselect
//
// Given an integer array nums and an integer k, return the kth largest element
// in the array.
//
// Note that it is the kth largest element in the sorted order, not the kth
// distinct element.
//
// Time complexity: O(n log k) for heap approach, Space complexity: O(k)
func KthLargest(nums []int, k int) int {
	if len(nums) == 0 || k < 1 || k > len(nums) {
		return -1 // Invalid input
	}

	// Use a min-heap to keep track of the k largest elements
	h := &minHeap{}
	heap.Init(h)

	// Add first k elements to heap
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	// For remaining elements, if larger than heap's minimum, replace it
	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	// The root of the min-heap is the kth largest element
	return (*h)[0]
}

// minHeap implements heap.Interface for a min-heap
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
