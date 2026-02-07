package trees

import (
	"leetcode/utils"
	"strconv"
)

// BinaryTreePaths solves LeetCode problem 0257: Binary Tree Paths
// Difficulty: Easy
// Tags: Tree, Depth-First Search, String
//
// Given the root of a binary tree, return all root-to-leaf paths in any order.
// A leaf is a node with no children.
//
// Example 1:
// Input: root = [1,2,3,null,5]
// Output: ["1->2->5","1->3"]
//
// Example 2:
// Input: root = [1]
// Output: ["1"]
//
// Constraints:
// - The number of nodes in the tree is in the range [1, 100].
// - -100 <= Node.val <= 100
//
// Time complexity: O(n), Space complexity: O(n) for recursion stack
func BinaryTreePaths(root *utils.TreeNode) []string {
	var result []string
	if root == nil {
		return result
	}
	
	dfsPaths(root, "", &result)
	return result
}

func dfsPaths(node *utils.TreeNode, currentPath string, result *[]string) {
	// Add current node value to path
	var newPath string
	if currentPath == "" {
		newPath = strconv.Itoa(node.Val)
	} else {
		newPath = currentPath + "->" + strconv.Itoa(node.Val)
	}
	
	// If leaf node, add path to result
	if node.Left == nil && node.Right == nil {
		*result = append(*result, newPath)
		return
	}
	
	// Recursively traverse left and right subtrees
	if node.Left != nil {
		dfsPaths(node.Left, newPath, result)
	}
	if node.Right != nil {
		dfsPaths(node.Right, newPath, result)
	}
}