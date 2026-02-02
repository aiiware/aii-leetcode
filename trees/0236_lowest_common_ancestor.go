package trees

import "leetcode/utils"

// LowestCommonAncestor solves LeetCode problem 0236: Lowest Common Ancestor of a Binary Tree
// Difficulty: Medium
// Tags: Tree, Depth-First Search, Binary Tree
//
// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
//
// According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined
// between two nodes p and q as the lowest node in T that has both p and q as descendants
// (where we allow a node to be a descendant of itself)."
//
// Example 1:
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.
//
// Example 2:
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself.
//
// Example 3:
// Input: root = [1,2], p = 1, q = 2
// Output: 1
//
// Constraints:
// - The number of nodes in the tree is in the range [2, 10^5].
// - -10^9 <= Node.val <= 10^9
// - All Node.val are unique.
// - p != q
// - p and q will exist in the tree.
//
// Time complexity: O(n), Space complexity: O(n) for recursion stack
func LowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	// Base case: if root is nil or root is one of p or q, return root
	if root == nil || root == p || root == q {
		return root
	}

	// Recursively search in left and right subtrees
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// If both left and right are non-nil, current root is LCA
	if left != nil && right != nil {
		return root
	}

	// Otherwise, return non-nil child (if any)
	if left != nil {
		return left
	}
	return right
}

// LowestCommonAncestorIterative solves using iterative approach with parent pointers
func LowestCommonAncestorIterative(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// Stack for DFS
	stack := []*utils.TreeNode{root}
	
	// Map to store parent pointers
	parent := make(map[*utils.TreeNode]*utils.TreeNode)
	parent[root] = nil

	// Traverse until we find both p and q
	for (parent[p] == nil || parent[q] == nil) && len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
	}

	// Collect ancestors of p
	ancestors := make(map[*utils.TreeNode]bool)
	for node := p; node != nil; node = parent[node] {
		ancestors[node] = true
	}

	// Find first common ancestor with q
	for node := q; node != nil; node = parent[node] {
		if ancestors[node] {
			return node
		}
	}

	return nil
}

// LowestCommonAncestorBSTForComparison solves the simpler case for Binary Search Tree (problem 0235)
// Included here for comparison - renamed to avoid conflict with 0235
func LowestCommonAncestorBSTForComparison(root, p, q *utils.TreeNode) *utils.TreeNode {
	// For BST, we can use the property: left < root < right
	current := root

	for current != nil {
		// If both p and q are greater than current, LCA is in right subtree
		if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
		} else if p.Val < current.Val && q.Val < current.Val {
			// If both p and q are less than current, LCA is in left subtree
			current = current.Left
		} else {
			// We have found the split point, or one of p/q equals current
			return current
		}
	}

	return nil
}

// FindNode finds a node with given value in the tree (helper function)
func FindNode(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := FindNode(root.Left, val); left != nil {
		return left
	}
	return FindNode(root.Right, val)
}