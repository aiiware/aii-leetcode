package leetcode

// Problem 0102: Binary Tree Level Order Traversal
//
// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
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

// levelOrderBFS is the standard BFS solution using a queue.
// This is the most intuitive solution for level order traversal.
// Time complexity: O(n), Space complexity: O(n) worst case (complete tree)
func levelOrderBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// levelOrderDFS uses DFS with level tracking.
// This approach uses recursion to traverse the tree while tracking the current level.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func levelOrderDFS(root *TreeNode) [][]int {
	result := [][]int{}
	dfs(root, 0, &result)
	return result
}

// dfs helper function for DFS solution
func dfs(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// If we need a new level, add it
	if level >= len(*result) {
		*result = append(*result, []int{})
	}

	// Add current node value to its level
	(*result)[level] = append((*result)[level], node.Val)

	// Recursively process children at next level
	dfs(node.Left, level+1, result)
	dfs(node.Right, level+1, result)
}

// levelOrderTwoQueues uses two queues to separate levels.
// This approach uses two queues to clearly separate current and next levels.
// Time complexity: O(n), Space complexity: O(n)
func levelOrderTwoQueues(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		level := make([]int, 0, len(currentLevel))
		nextLevel := []*TreeNode{}

		// Process all nodes in current level
		for _, node := range currentLevel {
			level = append(level, node.Val)

			// Add children to next level
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		result = append(result, level)
		currentLevel = nextLevel
	}

	return result
}

// levelOrderRecursiveWithQueue uses recursion with queue state.
// This is a hybrid approach that uses recursion but maintains queue state.
// Time complexity: O(n), Space complexity: O(n)
func levelOrderRecursiveWithQueue(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	processLevel(queue, &result)
	return result
}

// processLevel helper for recursive queue solution
func processLevel(nodes []*TreeNode, result *[][]int) {
	if len(nodes) == 0 {
		return
	}

	level := make([]int, 0, len(nodes))
	nextLevel := []*TreeNode{}

	// Process current level
	for _, node := range nodes {
		level = append(level, node.Val)

		// Collect children for next level
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}

	// Add current level to result
	*result = append(*result, level)

	// Recursively process next level
	processLevel(nextLevel, result)
}

// levelOrderOptimized is an optimized BFS solution with pre-allocation.
// This version pre-allocates slices for better performance.
// Time complexity: O(n), Space complexity: O(n)
func levelOrderOptimized(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[i] // Direct access without shifting

			level[i] = node.Val

			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Remove processed nodes from queue
		queue = queue[levelSize:]
		result = append(result, level)
	}

	return result
}

// levelOrder is the main entry point that uses the BFS solution.
// This follows LeetCode's expected function signature.
func levelOrder(root *TreeNode) [][]int {
	return levelOrderBFS(root)
}