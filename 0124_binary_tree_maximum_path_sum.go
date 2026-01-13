package leetcode

// MaxPathSum solves LeetCode problem 0124: Binary Tree Maximum Path Sum
// Difficulty: Hard
// Tags: Tree, Depth-First Search, Dynamic Programming
//
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes
// in the sequence has an edge connecting them. A node can only appear in the sequence
// at most once. Note that the path does not need to pass through the root.
//
// The path sum of a path is the sum of the node's values in the path.
//
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
// Example 1:
// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
//
// Example 2:
// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
//
// Constraints:
// The number of nodes in the tree is in the range [1, 3 * 10^4].
// -1000 <= Node.val <= 1000
//
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Initialize with minimum value
	maxSum := -1 << 31

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Recursively get max path sum from left and right children
		leftSum := max(0, dfs(node.Left))  // Only take positive contributions
		rightSum := max(0, dfs(node.Right)) // Only take positive contributions

		// Current path sum if we use current node as the "root" of the path
		currentPathSum := node.Val + leftSum + rightSum

		// Update global maximum
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		// Return the maximum path sum that can be extended to parent
		// We can only extend one side (left or right) to parent
		return node.Val + max(leftSum, rightSum)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}