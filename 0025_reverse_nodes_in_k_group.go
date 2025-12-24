package leetcode

// ReverseKGroup solves LeetCode problem 0025: Reverse Nodes in k-Group
// Difficulty: Hard
// Tags: Linked List, Recursion
//
// Given the head of a linked list, reverse the nodes of the list k at a time,
// and return the modified list. k is a positive integer and is less than or
// equal to the length of the linked list. If the number of nodes is not a
// multiple of k then left-out nodes, in the end, should remain as it is.
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
// Time complexity: O(n), Space complexity: O(n/k) for recursion stack
func ReverseKGroup(head *ListNode, k int) *ListNode {
	// Count the number of nodes
	count := 0
	current := head
	for current != nil && count < k {
		current = current.Next
		count++
	}

	// If we have k nodes, reverse them
	if count == k {
		// Reverse first k nodes
		reversedHead := reverseList(head, k)

		// Recursively reverse the remaining list
		head.Next = ReverseKGroup(current, k)

		return reversedHead
	}

	// Not enough nodes to reverse, return as is
	return head
}

// reverseList reverses the first k nodes of a linked list
func reverseList(head *ListNode, k int) *ListNode {
	var prev *ListNode
	current := head

	for i := 0; i < k; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// ReverseKGroupIterative is an iterative implementation
func ReverseKGroupIterative(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		// Check if there are k nodes left
		kth := getKthNode(prevGroupEnd, k)
		if kth == nil {
			break
		}

		nextGroupStart := kth.Next
		groupStart := prevGroupEnd.Next

		// Reverse the group
		reverseListIterative(groupStart, kth)

		// Connect the reversed group
		prevGroupEnd.Next = kth
		groupStart.Next = nextGroupStart

		// Move to next group
		prevGroupEnd = groupStart
	}

	return dummy.Next
}

// getKthNode returns the kth node from start (1-indexed)
func getKthNode(start *ListNode, k int) *ListNode {
	current := start
	for i := 0; i < k && current != nil; i++ {
		current = current.Next
	}
	return current
}

// reverseListIterative reverses nodes from start to end (inclusive)
func reverseListIterative(start, end *ListNode) {
	var prev *ListNode
	current := start
	stop := end.Next

	for current != stop {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
}