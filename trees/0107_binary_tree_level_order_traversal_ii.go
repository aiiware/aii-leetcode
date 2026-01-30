package trees

import "leetcode/utils"

// Problem 0107: Binary Tree Level Order Traversal II
//
// Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values.
// (i.e., from left to right, level by level from leaf to root).
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: [[15,7],[9,20],[3]]
// Explanation:
// Level 0: [3] (root)
// Level 1: [9,20]
// Level 2: [15,7]
// Bottom-up: [[15,7],[9,20],[3]]
//
// Example 2:
// Input: root = [1]
// Output: [[1]]
//
// Example 3:
// Input: root = []
// Output: []
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 2000].
// - -1000 <= Node.val <= 1000

// levelOrderBottomBFS is the standard BFS solution.
// This approach performs regular level order traversal and then reverses the result.
// Time complexity: O(n), Space complexity: O(n) worst case (complete tree)
func levelOrderBottomBFS(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Add node value to level slice
			level = append(level, node.Val)

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Add level to result
		result = append(result, level)
	}

	// Reverse the result to get bottom-up order
	reverseSlice2D(result)
	return result
}

// levelOrderBottomDFS is the DFS solution.
// This approach uses depth-first search to traverse the tree and build levels.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func levelOrderBottomDFS(root *utils.TreeNode) [][]int {
	result := [][]int{}
	var dfs func(node *utils.TreeNode, depth int)
	
	dfs = func(node *utils.TreeNode, depth int) {
		if node == nil {
			return
		}
		
		// If we need a new level, add it
		if depth >= len(result) {
			result = append(result, []int{})
		}
		
		// Add node value to its level
		result[depth] = append(result[depth], node.Val)
		
		// Recursively traverse left and right subtrees
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	
	dfs(root, 0)
	
	// Reverse the result to get bottom-up order
	reverseSlice2D(result)
	return result
}

// levelOrderBottomOptimized is an optimized BFS solution that builds result in reverse order.
// This approach adds each new level to the beginning of the result instead of appending and reversing.
// Time complexity: O(n), Space complexity: O(n)
func levelOrderBottomOptimized(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Add node value to level slice
			level = append(level, node.Val)

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Add level to beginning of result (reverse order)
		result = append([][]int{level}, result...)
	}

	return result
}

// levelOrderBottom is the main function.
func levelOrderBottom(root *utils.TreeNode) [][]int {
	return levelOrderBottomOptimized(root)
}

// Helper function to reverse a 2D slice
func reverseSlice2D(slice [][]int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}