package leetcode

// SwapPairs solves LeetCode problem 0024: Swap Nodes in Pairs
// Difficulty: Medium
// Tags: Linked List, Recursion
//
// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)
//
// Time complexity: O(n), Space complexity: O(1) (iterative) or O(n) (recursive)
func SwapPairs(head *ListNode) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{Next: head}
	current := dummy

	// Swap pairs while there are at least two nodes left
	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		// Swap the pair
		first.Next = second.Next
		second.Next = first
		current.Next = second

		// Move to the next pair
		current = first
	}

	return dummy.Next
}

// SwapPairsRecursive is a recursive implementation
func SwapPairsRecursive(head *ListNode) *ListNode {
	// Base case: empty list or single node
	if head == nil || head.Next == nil {
		return head
	}

	// Nodes to be swapped
	first := head
	second := head.Next

	// Swap the pair recursively
	first.Next = SwapPairsRecursive(second.Next)
	second.Next = first

	return second
}