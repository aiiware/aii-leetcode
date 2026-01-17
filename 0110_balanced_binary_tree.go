package leetcode

// Problem 0110: Balanced Binary Tree
//
// Given a binary tree, determine if it is height-balanced.
//
// A height-balanced binary tree is a binary tree in which the depth of the two subtrees
// of every node never differs by more than one.
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: true
// Explanation:
//    3
//   / \
//  9  20
//    /  \
//   15   7
// The depth of left subtree of node 3 is 1, right subtree is 2, difference is 1.
//
// Example 2:
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false
// Explanation:
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
// The depth of left subtree of node 2 is 3, right subtree is 1, difference is 2.
//
// Example 3:
// Input: root = []
// Output: true
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 5000].
// - -10^4 <= Node.val <= 10^4

// Helper function
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// isBalancedTopDown is the top-down recursive solution.
// This approach checks balance at each node by calculating heights of left and right subtrees,
// then recursively checks balance for left and right subtrees.
// Time complexity: O(n^2) worst case (skewed tree), O(n log n) average (balanced tree)
// Space complexity: O(n) for recursion stack
func isBalancedTopDown(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// Check if current node is balanced
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	// Recursively check left and right subtrees
	return isBalancedTopDown(root.Left) && isBalancedTopDown(root.Right)
}

// isBalancedBottomUp is the bottom-up recursive solution (optimized).
// This approach calculates height and checks balance in a single pass.
// Returns (height, balanced) for each subtree.
// Time complexity: O(n), Space complexity: O(n) for recursion stack
func isBalancedBottomUp(root *TreeNode) bool {
	_, balanced := checkBalance(root)
	return balanced
}

// checkBalance returns (height, balanced) for a subtree
func checkBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	// Check left subtree
	leftHeight, leftBalanced := checkBalance(node.Left)
	if !leftBalanced {
		return 0, false
	}

	// Check right subtree
	rightHeight, rightBalanced := checkBalance(node.Right)
	if !rightBalanced {
		return 0, false
	}

	// Check if current node is balanced
	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	// Return height of current node
	return max(leftHeight, rightHeight) + 1, true
}

// isBalancedIterative is an iterative solution using postorder traversal.
// This approach uses a stack to simulate postorder traversal.
// Time complexity: O(n), Space complexity: O(n) for the stack
func isBalancedIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*TreeNode{root}
	heights := make(map[*TreeNode]int)
	var lastVisited *TreeNode

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		// If leaf node or both children processed
		if (node.Left == nil && node.Right == nil) ||
			(lastVisited != nil && (lastVisited == node.Left || lastVisited == node.Right)) {
			stack = stack[:len(stack)-1]

			leftHeight := heights[node.Left]
			rightHeight := heights[node.Right]

			if abs(leftHeight-rightHeight) > 1 {
				return false
			}

			heights[node] = max(leftHeight, rightHeight) + 1
			lastVisited = node
		} else {
			// Push right child first (so left gets processed first due to LIFO)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return true
}

// isBalanced is the main function.
func isBalanced(root *TreeNode) bool {
	return isBalancedBottomUp(root)
}

// Helper function to calculate height of a tree
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}