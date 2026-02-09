package arrays

/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// Problem 0108: Convert Sorted Array to Binary Search Tree
//
// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
//
// A height-balanced binary tree is a binary tree in which the depth of the two subtrees
// of every node never differs by more than one.
//
// Example 1:
// Input: nums = [-10,-3,0,5,9]
// Output: [0,-3,9,-10,null,5]
// Explanation: [0,-10,5,null,-3,null,9] is also accepted:
//     0
//    / \
//  -10  5
//    \   \
//    -3   9
//
// Example 2:
// Input: nums = [1,3]
// Output: [3,1]
// Explanation: [1,null,3] is also accepted:
//   1         3
//    \       /
//     3     1
//
// Constraints:
// - 1 <= nums.length <= 10^4
// - -10^4 <= nums[i] <= 10^4
// - nums is sorted in ascending order.

// sortedArrayToBSTRecursive is the standard recursive solution.
// This approach uses binary search to find the middle element as root,
// then recursively builds left and right subtrees from left and right halves.
// Time complexity: O(n), Space complexity: O(log n) for recursion stack
func sortedArrayToBSTRecursive(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Find middle element
	mid := len(nums) / 2

	// Create root with middle element
	root := &utils.TreeNode{Val: nums[mid]}

	// Recursively build left and right subtrees
	root.Left = sortedArrayToBSTRecursive(nums[:mid])
	root.Right = sortedArrayToBSTRecursive(nums[mid+1:])

	return root
}

// sortedArrayToBSTIterative is an iterative solution using a stack.
// This approach simulates the recursive process using an explicit stack.
// Time complexity: O(n), Space complexity: O(log n) for the stack
func sortedArrayToBSTIterative(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Stack stores tuples of (start, end, parent, isLeft)
	type stackItem struct {
		start  int
		end    int
		parent *utils.TreeNode
		isLeft bool
	}

	stack := []stackItem{{0, len(nums) - 1, nil, false}}
	var root *utils.TreeNode

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if item.start > item.end {
			continue
		}

		// Find middle element
		mid := (item.start + item.end) / 2
		node := &utils.TreeNode{Val: nums[mid]}

		// Set as root or attach to parent
		if item.parent == nil {
			root = node
		} else {
			if item.isLeft {
				item.parent.Left = node
			} else {
				item.parent.Right = node
			}
		}

		// Push right subtree first (so left gets processed first due to LIFO)
		stack = append(stack, stackItem{mid + 1, item.end, node, false})
		stack = append(stack, stackItem{item.start, mid - 1, node, true})
	}

	return root
}

// sortedArrayToBST is the main function.
func sortedArrayToBST(nums []int) *utils.TreeNode {
	return sortedArrayToBSTRecursive(nums)
}
