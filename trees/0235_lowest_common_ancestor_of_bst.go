package trees

import "leetcode/utils"

// LowestCommonAncestorBST solves LeetCode problem 0235: Lowest Common Ancestor of a Binary Search Tree
// Difficulty: Medium
// Tags: Tree, BST, DFS
//
// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
//
// According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes
// p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a
// descendant of itself)."
//
// Time complexity: O(h) where h is the height of the tree
// Space complexity: O(1) for iterative, O(h) for recursive call stack
func LowestCommonAncestorBST(root, p, q *utils.TreeNode) *utils.TreeNode {
	// Iterative approach using BST properties
	current := root

	for current != nil {
		// If both p and q are greater than current, LCA is in right subtree
		if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
		} else if p.Val < current.Val && q.Val < current.Val {
			// If both p and q are less than current, LCA is in left subtree
			current = current.Left
		} else {
			// We have found the split point, or one of the nodes equals current
			// This is the LCA
			return current
		}
	}

	return nil // Should never reach here for valid BST and nodes
}

// LowestCommonAncestorBSTRecursive is the recursive version
func LowestCommonAncestorBSTRecursive(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// If both nodes are in right subtree
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestorBSTRecursive(root.Right, p, q)
	}

	// If both nodes are in left subtree
	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestorBSTRecursive(root.Left, p, q)
	}

	// We have found the LCA (split point or one node equals root)
	return root
}