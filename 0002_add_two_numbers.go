package leetcode

// AddTwoNumbers solves LeetCode problem 0002: Add Two Numbers
// Difficulty: Medium
// Tags: Linked List, Math, Recursion
//
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Time complexity: O(max(m, n)), Space complexity: O(max(m, n))
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	// Traverse both lists
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		// Add value from l1 if exists
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// Add value from l2 if exists
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Calculate carry and digit
		carry = sum / 10
		digit := sum % 10

		// Create new node with digit
		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummy.Next
}