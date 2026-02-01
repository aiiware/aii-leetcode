package trees

import "leetcode/utils"

// DiameterOfBinaryTree solves LeetCode problem 0543: Diameter of Binary Tree
// Difficulty: Easy
// Tags: Tree, DFS, Binary Tree
//
// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two
// nodes in a tree. This path may or may not pass through the root.
// The length of a path between two nodes is represented by the number of edges
// between them.
//
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func DiameterOfBinaryTree(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter := 0

	// DFS helper function that returns the height of the tree
	var dfs func(node *utils.TreeNode) int
	dfs = func(node *utils.TreeNode) int {
		if node == nil {
			return 0
		}

		// Get heights of left and right subtrees
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)

		// Update max diameter (path through current node)
		// Diameter at this node = leftHeight + rightHeight
		currentDiameter := leftHeight + rightHeight
		if currentDiameter > maxDiameter {
			maxDiameter = currentDiameter
		}

		// Return height of current node
		// Height = 1 + max(leftHeight, rightHeight)
		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	// Start DFS from root
	dfs(root)

	return maxDiameter
}