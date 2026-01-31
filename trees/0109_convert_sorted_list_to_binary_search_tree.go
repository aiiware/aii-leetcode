package trees


/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// Problem 0109: Convert Sorted List to Binary Search Tree
//
// Given the head of a singly linked list where elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
//
// Example 1:
// Input: head = [-10,-3,0,5,9]
// Output: [0,-3,9,-10,null,5]
// Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
//
// Example 2:
// Input: head = []
// Output: []
//
// Constraints:
// - The number of nodes in head is in the range [0, 2 * 10^4].
// - -10^5 <= Node.val <= 10^5

// sortedListToBSTSlowFast is the two-pointer (slow-fast) solution.
// This approach uses the slow-fast pointer technique to find the middle of the list,
// then recursively builds left and right subtrees.
// Time complexity: O(n log n), Space complexity: O(log n) for recursion stack
func sortedListToBSTSlowFast(head *utils.ListNode) *utils.TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &utils.TreeNode{Val: head.Val}
	}

	// Find the middle of the list using slow-fast pointers
	var prev *utils.ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Disconnect the left half from the middle node
	if prev != nil {
		prev.Next = nil
	}

	// Create root with middle node value
	root := &utils.TreeNode{Val: slow.Val}

	// Recursively build left and right subtrees
	root.Left = sortedListToBSTSlowFast(head)
	root.Right = sortedListToBSTSlowFast(slow.Next)

	return root
}

// sortedListToBSTArray is the array conversion solution.
// This approach converts the linked list to an array first,
// then uses the same approach as sortedArrayToBST.
// Time complexity: O(n), Space complexity: O(n) for the array
func sortedListToBSTArray(head *utils.ListNode) *utils.TreeNode {
	// Convert linked list to array
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	// Use local sortedArrayToBST function
	return sortedArrayToBSTLocal(nums)
}

// sortedListToBSTInorder is the inorder simulation solution.
// This approach simulates inorder traversal to build the BST.
// Time complexity: O(n), Space complexity: O(log n) for recursion stack
func sortedListToBSTInorder(head *utils.ListNode) *utils.TreeNode {
	// Count the number of nodes
	n := 0
	current := head
	for current != nil {
		n++
		current = current.Next
	}

	// Use inorder traversal to build the tree
	current = head
	var build func(start, end int) *utils.TreeNode
	build = func(start, end int) *utils.TreeNode {
		if start > end {
			return nil
		}

		mid := (start + end) / 2

		// Build left subtree
		left := build(start, mid-1)

		// Create root node
		root := &utils.TreeNode{Val: current.Val}
		current = current.Next

		// Build right subtree
		right := build(mid+1, end)

		root.Left = left
		root.Right = right
		return root
	}

	return build(0, n-1)
}

// sortedListToBST is the main function.
func sortedListToBST(head *utils.ListNode) *utils.TreeNode {
	return sortedListToBSTInorder(head)
}