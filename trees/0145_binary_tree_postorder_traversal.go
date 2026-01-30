package trees

import "leetcode/utils"

// PostorderTraversal solves LeetCode problem 0145: Binary Tree Postorder Traversal
// Difficulty: Easy
// Tags: Tree, Depth-First Search, Binary Tree, Stack
//
// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
// Postorder traversal: Left -> Right -> Root
//
// Example 1:
// Input: root = [1,null,2,3]
// Output: [3,2,1]
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [1]
// Output: [1]
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 100].
// - -100 <= Node.val <= 100
//
// Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Time complexity: O(n), Space complexity: O(n) for recursion stack or explicit stack
func PostorderTraversal(root *utils.TreeNode) []int {
	// Recursive solution
	result := []int{}
	postorderRecursive(root, &result)
	return result
}

// PostorderTraversalIterative solves the problem using iterative approach with stack
func PostorderTraversalIterative(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*utils.TreeNode{root}
	var lastVisited *utils.TreeNode

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		// If we're at a leaf node or we've already visited the children
		if (node.Left == nil && node.Right == nil) ||
			(lastVisited != nil && (lastVisited == node.Right || lastVisited == node.Left)) {
			// Visit the node
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			lastVisited = node
		} else {
			// Push right child first (so left is processed first)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return result
}

// PostorderTraversalTwoStacks solves the problem using two stacks approach
func PostorderTraversalTwoStacks(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack1 := []*utils.TreeNode{root}
	stack2 := []*utils.TreeNode{}

	// Process using first stack, push to second stack
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)

		// Push left then right to stack1
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}

	// Process second stack (which gives postorder)
	for len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		result = append(result, node.Val)
	}

	return result
}

// PostorderTraversalReversePreorder solves by doing reverse preorder
func PostorderTraversalReversePreorder(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*utils.TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Append to front (or we can append and reverse later)
		result = append([]int{node.Val}, result...)
		
		// Push left then right (opposite of preorder)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return result
}

// Helper function for recursive postorder traversal
func postorderRecursive(node *utils.TreeNode, result *[]int) {
	if node == nil {
		return
	}
	postorderRecursive(node.Left, result)
	postorderRecursive(node.Right, result)
	*result = append(*result, node.Val)
}