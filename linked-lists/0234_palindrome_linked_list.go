package linkedlists

import "leetcode/utils"

// IsPalindrome solves LeetCode problem 0234: Palindrome Linked List
// Difficulty: Easy
// Tags: Linked List, Two Pointers, Stack, Recursion
//
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
//
// Example 1:
// Input: head = [1,2,2,1]
// Output: true
//
// Example 2:
// Input: head = [1,2]
// Output: false
//
// Constraints:
// - The number of nodes in the list is in the range [1, 10^5].
// - 0 <= Node.val <= 9
//
// Follow up: Could you do it in O(n) time and O(1) space?
//
// Time complexity: O(n), Space complexity: O(1) for optimal solution
func IsPalindrome(head *utils.ListNode) bool {
	// Edge cases: empty list or single node
	if head == nil || head.Next == nil {
		return true
	}

	// Step 1: Find the middle of the linked list using slow/fast pointers
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half of the list
	secondHalf := reverseHalf(slow)
	firstHalf := head

	// Step 3: Compare the first half with the reversed second half
	temp := secondHalf
	result := true
	for temp != nil {
		if firstHalf.Val != temp.Val {
			result = false
			break
		}
		firstHalf = firstHalf.Next
		temp = temp.Next
	}

	// Step 4: Restore the list (optional but good practice)
	reverseHalf(secondHalf)

	return result
}

// Helper function to reverse a linked list (renamed to avoid conflict)
func reverseHalf(head *utils.ListNode) *utils.ListNode {
	var prev *utils.ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// IsPalindromeStack solves using stack approach (O(n) space)
func IsPalindromeStack(head *utils.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Push all values to a stack
	stack := []int{}
	current := head
	for current != nil {
		stack = append(stack, current.Val)
		current = current.Next
	}

	// Compare with original list
	current = head
	for i := len(stack) - 1; i >= 0; i-- {
		if current.Val != stack[i] {
			return false
		}
		current = current.Next
	}

	return true
}

// IsPalindromeRecursive solves using recursion (O(n) space for recursion stack)
func IsPalindromeRecursive(head *utils.ListNode) bool {
	frontPointer := head

	var recursivelyCheck func(*utils.ListNode) bool
	recursivelyCheck = func(currentNode *utils.ListNode) bool {
		if currentNode != nil {
			if !recursivelyCheck(currentNode.Next) {
				return false
			}
			if currentNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}

	return recursivelyCheck(head)
}