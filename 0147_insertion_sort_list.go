package leetcode

// InsertionSortList solves LeetCode problem 0147: Insertion Sort List
// Difficulty: Medium
// Tags: Linked List, Sorting
//
// Given the head of a singly linked list, sort the list using insertion sort,
// and return the sorted list's head.
//
// The steps of the insertion sort algorithm:
// 1. Insertion sort iterates, consuming one input element each repetition and
//    growing a sorted output list.
// 2. At each iteration, insertion sort removes one element from the input data,
//    finds the location it belongs within the sorted list, and inserts it there.
// 3. It repeats until no input elements remain.
//
// Example 1:
// Input: head = [4,2,1,3]
// Output: [1,2,3,4]
//
// Example 2:
// Input: head = [-1,5,3,4,0]
// Output: [-1,0,3,4,5]
//
// Constraints:
// - The number of nodes in the list is in the range [1, 5000].
// - -5000 <= Node.val <= 5000
//
// Time complexity: O(n^2) worst case, O(n) best case (already sorted)
// Space complexity: O(1)
func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Create a dummy node to serve as the start of the sorted list
	dummy := &ListNode{Next: head}
	
	// The last node of the sorted part
	lastSorted := head
	current := head.Next

	for current != nil {
		if lastSorted.Val <= current.Val {
			// Current node is in correct position
			lastSorted = current
			current = current.Next
		} else {
			// Need to insert current node into sorted part
			prev := dummy
			
			// Find the insertion position
			for prev.Next.Val <= current.Val {
				prev = prev.Next
			}
			
			// Insert current between prev and prev.Next
			lastSorted.Next = current.Next
			current.Next = prev.Next
			prev.Next = current
			
			// Move to next node to process
			current = lastSorted.Next
		}
	}

	return dummy.Next
}

// InsertionSortListOptimized is an optimized version with cleaner code
func InsertionSortListOptimized(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Create a dummy node for the sorted list
	dummy := &ListNode{}
	current := head

	for current != nil {
		// Keep track of next node to process
		next := current.Next
		
		// Find insertion position in sorted list
		prev := dummy
		for prev.Next != nil && prev.Next.Val < current.Val {
			prev = prev.Next
		}
		
		// Insert current node
		current.Next = prev.Next
		prev.Next = current
		
		// Move to next node
		current = next
	}

	return dummy.Next
}

// InsertionSortListWithArray solves by converting to array, sorting, and rebuilding
// This is not the intended solution but shows an alternative approach
// Time complexity: O(n log n), Space complexity: O(n)
func InsertionSortListWithArray(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Convert linked list to array
	nodes := []*ListNode{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	// Sort array using insertion sort (to match problem requirement)
	for i := 1; i < len(nodes); i++ {
		key := nodes[i]
		j := i - 1
		
		// Move elements of nodes[0..i-1] that are greater than key.Val
		// to one position ahead of their current position
		for j >= 0 && nodes[j].Val > key.Val {
			nodes[j+1] = nodes[j]
			j--
		}
		nodes[j+1] = key
	}

	// Rebuild linked list from sorted array
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil

	return nodes[0]
}