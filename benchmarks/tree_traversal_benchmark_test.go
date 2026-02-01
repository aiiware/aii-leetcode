package benchmarks

import (
	"leetcode/trees"
	"leetcode/utils"
	"testing"
)

// createTestTree creates a balanced binary tree of given depth
func createTestTree(depth int) *utils.TreeNode {
	if depth <= 0 {
		return nil
	}
	
	root := &utils.TreeNode{Val: 1}
	
	// Create left and right subtrees recursively
	if depth > 1 {
		root.Left = createTestTree(depth - 1)
		root.Right = createTestTree(depth - 1)
		
		// Adjust values to make them unique
		if root.Left != nil {
			root.Left.Val = root.Val * 2
		}
		if root.Right != nil {
			root.Right.Val = root.Val*2 + 1
		}
	}
	
	return root
}

// BenchmarkInorderTraversal benchmarks inorder traversal
func BenchmarkInorderTraversal(b *testing.B) {
	tree := createTestTree(10) // Tree with ~1023 nodes

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trees.InorderTraversal(tree)
	}
}

// BenchmarkPreorderTraversal benchmarks preorder traversal
func BenchmarkPreorderTraversal(b *testing.B) {
	tree := createTestTree(10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trees.PreorderTraversal(tree)
	}
}

// BenchmarkTreeTraversalSmall benchmarks with small tree
func BenchmarkTreeTraversalSmall(b *testing.B) {
	tree := createTestTree(5) // Small tree (~31 nodes)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trees.InorderTraversal(tree)
	}
}

// BenchmarkTreeTraversalLarge benchmarks with large tree
func BenchmarkTreeTraversalLarge(b *testing.B) {
	tree := createTestTree(12) // Large tree (~4095 nodes)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trees.InorderTraversal(tree)
	}
}

// BenchmarkTreeTraversalSkewed benchmarks with skewed tree
func BenchmarkTreeTraversalSkewed(b *testing.B) {
	// Create a right-skewed tree
	root := &utils.TreeNode{Val: 1}
	current := root
	for i := 2; i <= 1000; i++ {
		current.Right = &utils.TreeNode{Val: i}
		current = current.Right
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trees.InorderTraversal(root)
	}
}

// BenchmarkTreeTraversalComplete benchmarks with complete tree
func BenchmarkTreeTraversalComplete(b *testing.B) {
	// Use utils.CreateCompleteTree
	tree := utils.CreateCompleteTree(1023) // Complete tree with 1023 nodes

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trees.InorderTraversal(tree)
	}
}