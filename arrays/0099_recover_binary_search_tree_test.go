package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverBinarySearchTree(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		// Create tree: [1,3,null,null,2]
		//       1
		//      / \
		//     3   2
		// Should become: [3,1,null,null,2]
		//       3
		//      / \
		//     1   2

		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 3}
		root.Right = &TreeNode{Val: 2}

		RecoverBinarySearchTree(root)

		// Check that tree is now valid BST
		assert.Equal(t, 3, root.Val)
		assert.Equal(t, 1, root.Left.Val)
		assert.Equal(t, 2, root.Right.Val)
	})

	t.Run("Example 2", func(t *testing.T) {
		// Create tree: [3,1,4,null,null,2]
		//       3
		//      / \
		//     1   4
		//        / \
		//       2   5
		// Should become: [2,1,4,null,null,3,5]
		//       2
		//      / \
		//     1   4
		//        / \
		//       3   5

		root := &TreeNode{Val: 3}
		root.Left = &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 5}

		RecoverBinarySearchTree(root)

		// Check that tree is now valid BST
		assert.Equal(t, 2, root.Val)
		assert.Equal(t, 1, root.Left.Val)
		assert.Equal(t, 4, root.Right.Val)
		assert.Equal(t, 3, root.Right.Left.Val)
		assert.Equal(t, 5, root.Right.Right.Val)
	})

	t.Run("Single node", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		RecoverBinarySearchTree(root)
		assert.Equal(t, 1, root.Val)
	})
}

// BenchmarkRecoverBinarySearchTree benchmarks the RecoverBinarySearchTree function
func BenchmarkRecoverBinarySearchTree(b *testing.B) {
	// Create a larger tree for benchmarking
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RecoverBinarySearchTree(root)
	}
}
