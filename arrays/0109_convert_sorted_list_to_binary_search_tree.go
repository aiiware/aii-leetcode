package arrays

// SortedListToBST solves LeetCode problem 0109: Convert Sorted List to Binary Search Tree
// Difficulty: Medium
// Tags: Linked List, Depth-First Search, Binary Search Tree, Divide and Conquer
//
// Given the head of a singly linked list where elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
//
// Time complexity: O(n log n), Space complexity: O(log n) for recursion stack
func SortedListToBST(head *ListNode) *TreeNode {
	// Handle empty list case
	if head == nil {
		return nil
	}

	// Handle single node case
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// Use the fast and slow pointer technique to find the middle node
	slow := head
	fast := head
	var prev *ListNode

	// Find the middle node
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split the list into two halves
	prev.Next = nil

	// Create root node from the middle element
	root := &TreeNode{Val: slow.Val}

	// Recursively build left and right subtrees
	root.Left = SortedListToBST(head)
	root.Right = SortedListToBST(slow.Next)

	return root
}
