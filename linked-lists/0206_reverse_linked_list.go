package linkedlists

import "leetcode/utils"

// ReverseList solves LeetCode problem 0206: Reverse Linked List
// Difficulty: Easy
// Tags: Linked List, Recursion
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
// Example 2:
// Input: head = [1,2]
// Output: [2,1]
//
// Example 3:
// Input: head = []
// Output: []
//
// Constraints:
// - The number of nodes in the list is the range [0, 5000].
// - -5000 <= Node.val <= 5000
//
// Time complexity: O(n), Space complexity: O(1) for iterative, O(n) for recursive
func ReverseList(head *utils.ListNode) *utils.ListNode {
	// Iterative approach (preferred for space efficiency)
	var prev *utils.ListNode
	current := head

	for current != nil {
		// Store next node before we change current.Next
		next := current.Next
		// Reverse the link
		current.Next = prev
		// Move pointers forward
		prev = current
		current = next
	}

	// prev is now the new head
	return prev
}

// ReverseListRecursive solves the same problem using recursion
func ReverseListRecursive(head *utils.ListNode) *utils.ListNode {
	// Base case: empty list or single node
	if head == nil || head.Next == nil {
		return head
	}

	// Recursively reverse the rest of the list
	reversed := ReverseListRecursive(head.Next)

	// Make the next node point to current node
	head.Next.Next = head
	// Break the original link
	head.Next = nil

	return reversed
}