package trees

import "leetcode/utils"

// SubtreeOfAnotherTree solves LeetCode problem 0572: Subtree of Another Tree
// Difficulty: Easy
// Tags: Tree, DFS, Binary Tree
//
// Given the roots of two binary trees root and subRoot, return true if there is a
// subtree of root with the same structure and node values of subRoot and false otherwise.
//
// A subtree of a binary tree tree is a tree that consists of a node in tree and all
// of this node's descendants. The tree tree could also be considered as a subtree of itself.
//
// Time complexity: O(m * n) where m is nodes in root, n is nodes in subRoot
// Space complexity: O(h) where h is the height of root
func SubtreeOfAnotherTree(root *utils.TreeNode, subRoot *utils.TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot == nil {
		return true // empty tree is subtree of any tree
	}

	// Check if current tree matches subRoot
	if isSameTreeLocal(root, subRoot) {
		return true
	}

	// Recursively check left and right subtrees
	return SubtreeOfAnotherTree(root.Left, subRoot) || SubtreeOfAnotherTree(root.Right, subRoot)
}

// isSameTreeLocal checks if two trees are identical
// Renamed to avoid conflict with isSameTree in trees/0100_same_tree.go
func isSameTreeLocal(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTreeLocal(p.Left, q.Left) && isSameTreeLocal(p.Right, q.Right)
}