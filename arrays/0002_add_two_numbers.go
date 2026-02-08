package arrays

// AddTwoNumbers solves LeetCode problem 0002: Add Two Numbers
// Difficulty: Medium
// Tags: Linked List, Math
//
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Time complexity: O(max(m,n)), Space complexity: O(max(m,n))
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy head for the result list
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	// Iterate through both lists
	for l1 != nil || l2 != nil || carry != 0 {
		// Get values from current nodes
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		// Calculate sum and new carry
		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10

		// Create new node with sum and advance
		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	return dummyHead.Next
}