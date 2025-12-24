package leetcode

import "container/heap"

// MergeKLists solves LeetCode problem 0023: Merge k Sorted Lists
// Difficulty: Hard
// Tags: Linked List, Divide and Conquer, Heap (Priority Queue)
//
// You are given an array of k linked-lists lists, each linked-list is sorted
// in ascending order. Merge all the linked-lists into one sorted linked-list
// and return it.
//
// Time complexity: O(N log k) where N is total nodes, k is number of lists
// Space complexity: O(k) for the heap
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// Create a min-heap
	h := &minHeap{}
	heap.Init(h)

	// Push the first node of each list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	// Create dummy node for result
	dummy := &ListNode{}
	current := dummy

	// Process the heap
	for h.Len() > 0 {
		// Get the smallest node
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next

		// Push the next node from the same list
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// minHeap implements heap.Interface for ListNode pointers
type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeKListsDivideConquer uses divide and conquer approach
func MergeKListsDivideConquer(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	interval := 1
	for interval < len(lists) {
		for i := 0; i+interval < len(lists); i += interval * 2 {
			lists[i] = MergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}