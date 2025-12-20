package leetcode

// RemoveNthFromEnd solves LeetCode problem 0019: Remove Nth Node From End of List
// Difficulty: Medium
// Tags: Linked List, Two Pointers
//
// Given the head of a linked list, remove the nth node from the end of the list
// and return its head.
//
// Example:
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
//
// Time complexity: O(n), Space complexity: O(1)
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	// Create a dummy node to handle edge cases (like removing the head)
	dummy := &ListNode{Val: 0, Next: head}

	// Use two pointers with a gap of n nodes between them
	first := dummy
	second := dummy

	// Move first pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		if first == nil {
			return head
		}
		first = first.Next
	}

	// Move both pointers until first reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node
	if second.Next != nil {
		second.Next = second.Next.Next
	}

	return dummy.Next
}

// Alternative approach using single pass with a counter
func RemoveNthFromEndAlt(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	// Count total nodes
	current := head
	count := 0
	for current != nil {
		count++
		current = current.Next
	}

	// If we need to remove the head
	if count == n {
		return head.Next
	}

	// Find the node before the one to remove
	current = head
	for i := 0; i < count-n-1; i++ {
		current = current.Next
	}

	// Remove the node
	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}
