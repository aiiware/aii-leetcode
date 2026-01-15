package leetcode

// SortList solves LeetCode problem 0148: Sort List
// Difficulty: Medium
// Tags: Linked List, Two Pointers, Divide and Conquer, Sorting, Merge Sort
//
// Given the head of a linked list, return the list after sorting it in ascending order.
//
// Follow up: Can you sort the linked list in O(n log n) time and O(1) memory (i.e., constant space)?
//
// Example 1:
// Input: head = [4,2,1,3]
// Output: [1,2,3,4]
//
// Example 2:
// Input: head = [-1,5,3,4,0]
// Output: [-1,0,3,4,5]
//
// Example 3:
// Input: head = []
// Output: []
//
// Constraints:
// - The number of nodes in the list is in the range [0, 5 * 10^4].
// - -10^5 <= Node.val <= 10^5
//
// Time complexity: O(n log n), Space complexity: O(1) for merge sort approach
func SortList(head *ListNode) *ListNode {
	// Base case: empty list or single node
	if head == nil || head.Next == nil {
		return head
	}

	// Find the middle of the list
	mid := findMiddle(head)
	
	// Split the list into two halves
	right := mid.Next
	mid.Next = nil
	
	// Recursively sort both halves
	leftSorted := SortList(head)
	rightSorted := SortList(right)
	
	// Merge the sorted halves
	return merge(leftSorted, rightSorted)
}

// findMiddle finds the middle node of a linked list using slow/fast pointer technique
func findMiddle(head *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := head, head
	
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	// For even number of nodes, return the first middle node
	if prev != nil {
		return prev
	}
	return slow
}

// merge merges two sorted linked lists
func merge(l1, l2 *ListNode) *ListNode {
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
	
	// Append remaining nodes
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	
	return dummy.Next
}

// SortListBottomUp is an alternative bottom-up merge sort implementation
// This approach uses O(1) space and O(n log n) time
func SortListBottomUp(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	// Get the length of the list
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	
	dummy := &ListNode{Next: head}
	
	// Bottom-up merge sort
	for step := 1; step < length; step <<= 1 {
		prev := dummy
		current = dummy.Next
		
		for current != nil {
			// Get the first sublist
			left := current
			right := split(left, step)
			current = split(right, step)
			
			// Merge the two sublists
			merged := merge(left, right)
			
			// Connect the merged list
			prev.Next = merged
			
			// Move prev to the end of merged list
			for prev.Next != nil {
				prev = prev.Next
			}
		}
	}
	
	return dummy.Next
}

// split splits the list after n nodes and returns the head of the second part
func split(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	
	for i := 1; i < n && head.Next != nil; i++ {
		head = head.Next
	}
	
	second := head.Next
	head.Next = nil
	return second
}