package leetcode

// Problem 0106: Construct Binary Tree from Inorder and Postorder Traversal
//
// Given two integer arrays inorder and postorder where inorder is the inorder traversal
// of a binary tree and postorder is the postorder traversal of the same tree,
// construct and return the binary tree.
//
// Example 1:
// Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// Output: [3,9,20,null,null,15,7]
// Explanation:
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// Example 2:
// Input: inorder = [-1], postorder = [-1]
// Output: [-1]
//
// Constraints:
// - 1 <= inorder.length <= 3000
// - postorder.length == inorder.length
// - -3000 <= inorder[i], postorder[i] <= 3000
// - inorder and postorder consist of unique values.
// - Each value of postorder also appears in inorder.
// - inorder is guaranteed to be the inorder traversal of the tree.
// - postorder is guaranteed to be the postorder traversal of the tree.

// buildTreeFromInorderPostorderRecursive is the standard recursive solution.
// This approach uses the property that in postorder traversal, the last element is always the root.
// In inorder traversal, all elements left of the root are in the left subtree,
// and all elements right of the root are in the right subtree.
// Time complexity: O(n^2) in worst case (unbalanced tree), O(n log n) average (balanced tree)
// Space complexity: O(n) for recursion stack
func buildTreeFromInorderPostorderRecursive(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// The last element in postorder is always the root
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder traversal
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// Recursively build left and right subtrees
	// Left subtree: inorder elements before rootIndex, postorder elements with same count
	// Right subtree: inorder elements after rootIndex, postorder elements before last element
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder):len(postorder)-1]

	root.Left = buildTreeFromInorderPostorderRecursive(leftInorder, leftPostorder)
	root.Right = buildTreeFromInorderPostorderRecursive(rightInorder, rightPostorder)

	return root
}

// buildTreeFromInorderPostorderOptimized is an optimized version using a hash map for O(1) lookups.
// This approach preprocesses the inorder array into a map for fast root index lookup.
// Time complexity: O(n), Space complexity: O(n) for the map
func buildTreeFromInorderPostorderOptimized(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// Create a map from value to index in inorder traversal
	inorderIndex := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderIndex[val] = i
	}

	// Use a helper function with indices to avoid slicing
	var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
	build = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd || postStart > postEnd {
			return nil
		}

		// The last element in current postorder range is the root
		rootVal := postorder[postEnd]
		root := &TreeNode{Val: rootVal}

		// Find root index in inorder using the map
		rootIdx := inorderIndex[rootVal]

		// Calculate sizes of left and right subtrees
		leftSize := rootIdx - inStart

		// Recursively build left and right subtrees
		root.Left = build(inStart, rootIdx-1, postStart, postStart+leftSize-1)
		root.Right = build(rootIdx+1, inEnd, postStart+leftSize, postEnd-1)

		return root
	}

	return build(0, len(inorder)-1, 0, len(postorder)-1)
}

// buildTreeFromInorderPostorder is the main function.
func buildTreeFromInorderPostorder(inorder []int, postorder []int) *TreeNode {
	return buildTreeFromInorderPostorderOptimized(inorder, postorder)
}