package dp


/*
Difficulty: Hard
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// Problem 0113: Path Sum II
//
// Given the root of a binary tree and an integer targetSum, return all root-to-leaf
// paths where the sum of the node values in the path equals targetSum.
//
// Each path should be returned as a list of the node values, not node references.
//
// Example 1:
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: [[5,4,11,2],[5,8,4,5]]
// Explanation:
//    5
//   / \
//  4   8
// /   / \
// 11  13  4
// /  \    / \
// 7    2  5   1
// There are two paths whose sum equals targetSum:
// 5 + 4 + 11 + 2 = 22
// 5 + 8 + 4 + 5 = 22
//
// Example 2:
// Input: root = [1,2,3], targetSum = 5
// Output: []
//
// Example 3:
// Input: root = [1,2], targetSum = 0
// Output: []
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 5000].
// - -1000 <= Node.val <= 1000
// - -1000 <= targetSum <= 1000

// pathSum returns all root-to-leaf paths where the sum equals targetSum.
// This is a DFS recursive solution that tracks the current path.
// Time complexity: O(n^2) in worst case (skewed tree with many valid paths)
// Space complexity: O(n) for recursion stack and path storage
func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	result := [][]int{}
	currentPath := []int{}
	
	// Helper function for DFS traversal
	var dfs func(node *utils.TreeNode, remainingSum int)
	dfs = func(node *utils.TreeNode, remainingSum int) {
		if node == nil {
			return
		}
		
		// Add current node to path
		currentPath = append(currentPath, node.Val)
		
		// Check if we're at a leaf node
		if node.Left == nil && node.Right == nil {
			// If remaining sum equals node value, we found a valid path
			if remainingSum == node.Val {
				// Make a copy of current path to avoid mutation issues
				pathCopy := make([]int, len(currentPath))
				copy(pathCopy, currentPath)
				result = append(result, pathCopy)
			}
		} else {
			// Continue DFS traversal with updated remaining sum
			newRemaining := remainingSum - node.Val
			dfs(node.Left, newRemaining)
			dfs(node.Right, newRemaining)
		}
		
		// Backtrack: remove current node from path
		currentPath = currentPath[:len(currentPath)-1]
	}
	
	dfs(root, targetSum)
	return result
}

// pathSumIterative is an alternative iterative solution using stack.
// This approach uses explicit stack to avoid recursion overhead.
// Time complexity: O(n^2) in worst case
// Space complexity: O(n) for stack and path storage
func pathSumIterative(root *utils.TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	result := [][]int{}
	
	// Stack stores tuple of (node, remainingSum, path)
	type stackItem struct {
		node         *utils.TreeNode
		remainingSum int
		path         []int
	}
	
	stack := []stackItem{
		{node: root, remainingSum: targetSum, path: []int{}},
	}
	
	for len(stack) > 0 {
		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		node := item.node
		remainingSum := item.remainingSum
		path := item.path
		
		if node == nil {
			continue
		}
		
		// Add current node to path
		newPath := make([]int, len(path)+1)
		copy(newPath, path)
		newPath[len(path)] = node.Val
		
		// Check if we're at a leaf node
		if node.Left == nil && node.Right == nil {
			// If remaining sum equals node value, we found a valid path
			if remainingSum == node.Val {
				result = append(result, newPath)
			}
		} else {
			// Push children to stack with updated remaining sum
			newRemaining := remainingSum - node.Val
			if node.Left != nil {
				stack = append(stack, stackItem{
					node:         node.Left,
					remainingSum: newRemaining,
					path:         newPath,
				})
			}
			if node.Right != nil {
				stack = append(stack, stackItem{
					node:         node.Right,
					remainingSum: newRemaining,
					path:         newPath,
				})
			}
		}
	}
	
	return result
}