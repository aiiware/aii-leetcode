package arrays

// RecoverBinarySearchTree solves LeetCode problem 0099: Recover Binary Search Tree
// Difficulty: Hard
// Tags: Tree, Depth-First Search, Binary Search Tree
//
// You are given the root of a binary search tree, where exactly two nodes of the tree
// were swapped by mistake. Recover the tree without changing its structure.
//
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func RecoverBinarySearchTree(root *TreeNode) {
	// Initialize variables to track the misplaced nodes
	var first, second, prev *TreeNode

	// In-order traversal to find the misplaced nodes
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Traverse left subtree
		inorder(node.Left)

		// Check if current node is misplaced
		if prev != nil && prev.Val > node.Val {
			if first == nil {
				// First misplaced node
				first = prev
			}
			// Second misplaced node
			second = node
		}

		// Update previous node
		prev = node

		// Traverse right subtree
		inorder(node.Right)
	}

	// Perform in-order traversal
	inorder(root)

	// Swap the values of the misplaced nodes
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
