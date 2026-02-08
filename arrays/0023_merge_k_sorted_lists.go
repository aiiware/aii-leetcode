package arrays

// MergeKLists merges k sorted linked lists into one sorted list
// Difficulty: Hard
// Tags: Linked List, Divide and Conquer, Heap
//
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked lists into one sorted linked list and return it.
//
// Time complexity: O(N log k) where N is total number of nodes and k is number of lists
// Space complexity: O(1) if we don't count output list
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	// Use divide and conquer approach
	return mergeLists(lists, 0, len(lists)-1)
}

// mergeLists recursively merges lists using divide and conquer approach
func mergeLists(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	mid := start + (end-start)/2
	left := mergeLists(lists, start, mid)
	right := mergeLists(lists, mid+1, end)

	return mergeTwoLists(left, right)
}

// mergeTwoLists merges two sorted linked lists
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}