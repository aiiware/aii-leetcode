package linkedlists

import "leetcode/utils"

// RemoveElements solves LeetCode problem 0203: Remove Linked List Elements
// Difficulty: Easy
// Tags: Linked List, Recursion
//
// Given the head of a linked list and an integer val, remove all the nodes
// of the linked list that have Node.val == val, and return the new head.
//
// Example 1:
// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]
//
// Example 2:
// Input: head = [], val = 1
// Output: []
//
// Example 3:
// Input: head = [7,7,7,7], val = 7
// Output: []
//
// Constraints:
// - The number of nodes in the list is in the range [0, 10^4].
// - 1 <= Node.val <= 50
// - 0 <= val <= 50
//
// Time complexity: O(n), Space complexity: O(1) for iterative, O(n) for recursive
func RemoveElements(head *utils.ListNode, val int) *utils.ListNode {
	// Handle edge case: empty list
	if head == nil {
		return nil
	}

	// Create a dummy node to simplify edge cases (when head needs to be removed)
	dummy := &utils.ListNode{Next: head}
	prev := dummy
	current := head

	for current != nil {
		if current.Val == val {
			// Skip the current node
			prev.Next = current.Next
		} else {
			// Move prev pointer forward
			prev = current
		}
		// Move current pointer forward
		current = current.Next
	}

	return dummy.Next
}

// RemoveElementsRecursive solves the same problem using recursion
func RemoveElementsRecursive(head *utils.ListNode, val int) *utils.ListNode {
	// Base case: empty list
	if head == nil {
		return nil
	}

	// Recursively process the rest of the list
	head.Next = RemoveElementsRecursive(head.Next, val)

	// If current node has the target value, skip it
	if head.Val == val {
		return head.Next
	}

	return head
}

// RemoveElementsTwoPointers is an alternative iterative approach without dummy node
func RemoveElementsTwoPointers(head *utils.ListNode, val int) *utils.ListNode {
	// First, handle the case where head needs to be removed
	for head != nil && head.Val == val {
		head = head.Next
	}

	// If list is empty after removing head nodes
	if head == nil {
		return nil
	}

	// Now head is guaranteed not to have the target value
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			// Skip the next node
			current.Next = current.Next.Next
		} else {
			// Move to next node
			current = current.Next
		}
	}

	return head
}