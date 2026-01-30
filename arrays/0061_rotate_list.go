package arrays

import "leetcode/utils"

// RotateRight solves LeetCode problem 0061: Rotate List
// Difficulty: Medium
// Tags: Linked List, Two Pointers
//
// Given the head of a linked list, rotate the list to the right by k places.
//
// Example 1:
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
//
// Example 2:
// Input: head = [0,1,2], k = 4
// Output: [2,0,1]
//
// Time complexity: O(n), Space complexity: O(1)
func RotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Step 1: Calculate the length of the list
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Step 2: Normalize k (k might be larger than length)
	k %= length
	if k == 0 {
		return head // No rotation needed
	}

	// Step 3: Find the new tail (at position length - k - 1)
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}

	// Step 4: Perform rotation
	newHead := newTail.Next
	newTail.Next = nil
	tail.Next = head

	return newHead
}