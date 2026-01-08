package leetcode

// Problem 0092: Reverse Linked List II
//
// Given the head of a singly linked list and two integers left and right where left <= right, 
// reverse the nodes of the list from position left to position right, and return the reversed list.
//
// Example 1:
// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]
//
// Example 2:
// Input: head = [5], left = 1, right = 1
// Output: [5]
//
// Constraints:
// - The number of nodes in the list is n.
// - 1 <= n <= 500
// - -500 <= Node.val <= 500
// - 1 <= left <= right <= n

// reverseBetween is the main solution function.
// Time complexity: O(n), Space complexity: O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// Create a dummy node to handle edge cases
	dummy := &ListNode{Next: head}
	prev := dummy

	// Move prev to the node before the reversal section
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Start reversing from current
	current := prev.Next
	var next *ListNode

	// Reverse the section from left to right
	for i := 0; i < right-left; i++ {
		next = current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

// reverseBetweenTwoPass uses two-pass approach: find section, reverse, reconnect.
func reverseBetweenTwoPass(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// Find the node before left and the node at right
	dummy := &ListNode{Next: head}
	beforeLeft := dummy
	for i := 1; i < left; i++ {
		beforeLeft = beforeLeft.Next
	}

	// Find the node at right
	rightNode := beforeLeft
	for i := 0; i <= right-left; i++ {
		rightNode = rightNode.Next
	}

	// Extract the section to reverse
	sectionStart := beforeLeft.Next
	afterRight := rightNode.Next
	
	// Disconnect the section
	beforeLeft.Next = nil
	rightNode.Next = nil

	// Reverse the section
	reversedStart := reverseCompleteList(sectionStart)

	// Reconnect
	beforeLeft.Next = reversedStart
	sectionStart.Next = afterRight

	return dummy.Next
}

// reverseCompleteList reverses a complete linked list
func reverseCompleteList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// reverseBetweenRecursive uses recursion to reverse the section.
func reverseBetweenRecursive(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// Handle the case when left == 1 separately
	if left == 1 {
		return reverseN(head, right)
	}

	// Recursively handle the rest
	head.Next = reverseBetweenRecursive(head.Next, left-1, right-1)
	return head
}

// reverseN reverses the first n nodes of a linked list
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		return head
	}

	// Reverse the first n-1 nodes
	newHead := reverseN(head.Next, n-1)
	
	// Reconnect
	successor := head.Next.Next
	head.Next.Next = head
	head.Next = successor
	
	return newHead
}

// reverseBetweenStack uses stack to reverse the section.
func reverseBetweenStack(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// Move to node before left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Push nodes from left to right onto stack
	stack := []*ListNode{}
	current := prev.Next
	for i := 0; i <= right-left; i++ {
		stack = append(stack, current)
		current = current.Next
	}

	// Pop from stack and reconnect
	afterRight := current
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev.Next = node
		prev = prev.Next
	}

	// Connect to the rest
	prev.Next = afterRight

	return dummy.Next
}

// reverseBetweenInPlace uses in-place reversal with three pointers.
func reverseBetweenInPlace(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// Find the node before left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Reverse the section
	start := prev.Next
	then := start.Next

	for i := 0; i < right-left; i++ {
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = start.Next
	}

	return dummy.Next
}

// reverseBetweenOptimized is the optimized version.
func reverseBetweenOptimized(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// Move prev to position left-1
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Reverse from left to right
	curr := prev.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

// ReverseBetween is the public interface function.
// It uses the optimized solution by default.
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	return reverseBetweenOptimized(head, left, right)
}