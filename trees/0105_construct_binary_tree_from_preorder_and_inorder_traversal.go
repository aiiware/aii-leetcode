package trees

import "leetcode/utils"

// Problem 0105: Construct Binary Tree from Preorder and Inorder Traversal
//
// Given two integer arrays preorder and inorder where preorder is the preorder traversal
// of a binary tree and inorder is the inorder traversal of the same tree,
// construct and return the binary tree.
//
// Example 1:
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
// Explanation:
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// Example 2:
// Input: preorder = [-1], inorder = [-1]
// Output: [-1]
//
// Constraints:
// - 1 <= preorder.length <= 3000
// - inorder.length == preorder.length
// - -3000 <= preorder[i], inorder[i] <= 3000
// - preorder and inorder consist of unique values.
// - Each value of inorder also appears in preorder.
// - preorder is guaranteed to be the preorder traversal of the tree.
// - inorder is guaranteed to be the inorder traversal of the tree.

// buildTreeRecursive is the standard recursive solution.
// This approach uses the property that in preorder traversal, the first element is always the root.
// In inorder traversal, all elements left of the root are in the left subtree,
// and all elements right of the root are in the right subtree.
// Time complexity: O(n^2) in worst case (unbalanced tree), O(n log n) average (balanced tree)
// Space complexity: O(n) for recursion stack
func buildTreeRecursive(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// The first element in preorder is always the root
	rootVal := preorder[0]
	root := &utils.TreeNode{Val: rootVal}

	// Find the root in inorder traversal
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// Recursively build left and right subtrees
	// Left subtree: inorder elements before rootIndex, preorder elements after root (same count)
	// Right subtree: inorder elements after rootIndex, preorder elements after left subtree
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]

	root.Left = buildTreeRecursive(leftPreorder, leftInorder)
	root.Right = buildTreeRecursive(rightPreorder, rightInorder)

	return root
}

// buildTreeOptimized is an optimized version using a hash map for O(1) lookups.
// This approach preprocesses the inorder array into a map for fast root index lookup.
// Time complexity: O(n), Space complexity: O(n) for the map
func buildTreeOptimized(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Create a map from value to index in inorder traversal
	inorderIndex := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderIndex[val] = i
	}

	// Use a helper function with indices to avoid slicing
	var build func(preStart, preEnd, inStart, inEnd int) *utils.TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *utils.TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		// The first element in current preorder range is the root
		rootVal := preorder[preStart]
		root := &utils.TreeNode{Val: rootVal}

		// Find root index in inorder using the map
		rootIdx := inorderIndex[rootVal]

		// Calculate sizes of left and right subtrees
		leftSize := rootIdx - inStart

		// Recursively build left and right subtrees
		root.Left = build(preStart+1, preStart+leftSize, inStart, rootIdx-1)
		root.Right = build(preStart+leftSize+1, preEnd, rootIdx+1, inEnd)

		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

// buildTreeIterative is an iterative solution using a stack.
// This approach simulates the recursive process using an explicit stack.
// Time complexity: O(n), Space complexity: O(n) for the stack
func buildTreeIterative(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &utils.TreeNode{Val: preorder[0]}
	stack := []*utils.TreeNode{root}

	inorderIndex := 0

	for i := 1; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		preorderVal := preorder[i]

		// If current node's value doesn't match inorder value,
		// we're still building the left subtree
		if node.Val != inorder[inorderIndex] {
			node.Left = &utils.TreeNode{Val: preorderVal}
			stack = append(stack, node.Left)
		} else {
			// We've reached a point where we need to build right subtree
			// Pop nodes from stack until we find a node whose value
			// doesn't match inorder[inorderIndex]
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}

			// The right child is the current preorder value
			node.Right = &utils.TreeNode{Val: preorderVal}
			stack = append(stack, node.Right)
		}
	}

	return root
}

// buildTree is the main function that delegates to the optimized solution.
func buildTree(preorder []int, inorder []int) *utils.TreeNode {
	return buildTreeOptimized(preorder, inorder)
}