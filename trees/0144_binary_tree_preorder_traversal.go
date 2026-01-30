package trees

import "leetcode/utils"

// PreorderTraversal solves LeetCode problem 0144: Binary Tree Preorder Traversal
// Difficulty: Easy
// Tags: Tree, Depth-First Search, Binary Tree, Stack
//
// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
// Preorder traversal: Root -> Left -> Right
//
// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,2,3]
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [1]
// Output: [1]
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 100].
// - -100 <= Node.val <= 100
//
// Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Time complexity: O(n), Space complexity: O(n) for recursion stack or explicit stack
func PreorderTraversal(root *utils.TreeNode) []int {
	// Recursive solution
	result := []int{}
	preorderRecursive(root, &result)
	return result
}

// PreorderTraversalIterative solves the problem using iterative approach with stack
func PreorderTraversalIterative(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*utils.TreeNode{root}

	for len(stack) > 0 {
		// Pop from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Process current node
		result = append(result, node.Val)

		// Push right child first (so left is processed first)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// PreorderTraversalMorris solves the problem using Morris traversal (O(1) space)
func PreorderTraversalMorris(root *utils.TreeNode) []int {
	result := []int{}
	current := root

	for current != nil {
		if current.Left == nil {
			// If no left child, visit current node and go to right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find the inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// First time visiting this predecessor
				// Visit current node before traversing left subtree
				result = append(result, current.Val)
				// Create a temporary link to current
				predecessor.Right = current
				current = current.Left
			} else {
				// We've already visited this predecessor
				// Remove the temporary link
				predecessor.Right = nil
				current = current.Right
			}
		}
	}

	return result
}

// Helper function for recursive preorder traversal
func preorderRecursive(node *utils.TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preorderRecursive(node.Left, result)
	preorderRecursive(node.Right, result)
}