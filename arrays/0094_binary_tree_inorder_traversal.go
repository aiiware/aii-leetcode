package arrays

// InorderTraversal solves LeetCode problem 0094: Binary Tree Inorder Traversal
// Difficulty: Easy
// Tags: Tree, Depth-First Search, Binary Tree
//
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func InorderTraversal(root *TreeNode) []int {
	// Handle empty tree case
	if root == nil {
		return []int{}
	}

	// Use a stack for iterative traversal
	stack := []*TreeNode{}
	result := []int{}

	// Process nodes iteratively
	current := root
	for current != nil || len(stack) > 0 {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Process the top node from stack
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Move to right subtree
		current = current.Right
	}

	return result
}
