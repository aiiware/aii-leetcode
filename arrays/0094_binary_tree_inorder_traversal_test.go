package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	// Test case 1: Empty tree
	t.Run("Empty tree", func(t *testing.T) {
		result := InorderTraversal(nil)
		assert.Equal(t, []int{}, result)
	})

	// Test case 2: Single node
	t.Run("Single node", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		result := InorderTraversal(root)
		assert.Equal(t, []int{1}, result)
	})

	// Test case 3: Complete binary tree
	t.Run("Complete binary tree", func(t *testing.T) {
		//       1
		//      / \
		//     2   3
		//    / \
		//   4   5
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 5}
		
		result := InorderTraversal(root)
		assert.Equal(t, []int{4, 2, 5, 1, 3}, result)
	})

	// Test case 4: Left skewed tree
	t.Run("Left skewed tree", func(t *testing.T) {
		//   1
		//  /
		// 2
		///
		//3
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		
		result := InorderTraversal(root)
		assert.Equal(t, []int{3, 2, 1}, result)
	})

	// Test case 5: Right skewed tree
	t.Run("Right skewed tree", func(t *testing.T) {
		// 1
		//  \
		//   2
		//    \
		//     3
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 3}
		
		result := InorderTraversal(root)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	// Test case 6: Complex tree
	t.Run("Complex tree", func(t *testing.T) {
		//       5
		//      / \
		//     3   8
		//    / \ / \
		//   2  4 7  9
		//  /     /
		// 1     6
		root := &TreeNode{Val: 5}
		root.Left = &TreeNode{Val: 3}
		root.Right = &TreeNode{Val: 8}
		root.Left.Left = &TreeNode{Val: 2}
		root.Left.Right = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 7}
		root.Right.Right = &TreeNode{Val: 9}
		root.Left.Left.Left = &TreeNode{Val: 1}
		root.Right.Left.Left = &TreeNode{Val: 6}
		
		result := InorderTraversal(root)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func BenchmarkInorderTraversal(b *testing.B) {
	// Create a moderately sized tree for benchmarking
	root := &TreeNode{Val: 1}
	current := root
	
	// Create a left-skewed tree for benchmarking
	for i := 2; i <= 100; i++ {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InorderTraversal(root)
	}
}