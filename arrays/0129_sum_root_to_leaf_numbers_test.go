package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"leetcode/utils"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected int
	}{
		{
			name:     "Example 1",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}),
			expected: 25,
		},
		{
			name:     "Example 2",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(4), utils.IntPtr(9), utils.IntPtr(0), utils.IntPtr(5), utils.IntPtr(1)}),
			expected: 1026,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(5)}),
			expected: 5,
		},
		{
			name:     "Complete binary tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			expected: 522, // Paths: 124 + 125 + 136 + 137 = 522
		},
		{
			name:     "Right skewed tree",
			root:     utils.CreateRightSkewedTree(3),
			expected: 123, // Only path: 1->2->3 = 123
		},
		{
			name:     "Left skewed tree",
			root:     utils.CreateLeftSkewedTree(3),
			expected: 123, // Only path: 1->2->3 = 123
		},
		{
			name:     "Tree with zeros",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(0), utils.IntPtr(2)}),
			expected: 22, // Paths: 10 + 12 = 22
		},
		{
			name:     "Nil root",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Tree with all same digits",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(9), utils.IntPtr(9), utils.IntPtr(9)}),
			expected: 198, // Paths: 99 + 99 = 198
		},
		{
			name:     "Unbalanced tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4)}),
			expected: 247, // Paths: 123 + 124 = 247
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumNumbers(tt.root)
			assert.Equal(t, tt.expected, result,
				"SumNumbers() = %d, expected %d",
				result, tt.expected)
		})
	}
}

func TestSumNumbers_EdgeCases(t *testing.T) {
	t.Run("Maximum depth tree", func(t *testing.T) {
		// Create a tree with depth 10 (maximum per constraints)
		root := &utils.TreeNode{Val: 1}
		current := root
		for i := 2; i <= 10; i++ {
			current.Right = &utils.TreeNode{Val: i % 10} // Use modulo to keep digits 0-9
			current = current.Right
		}

		result := SumNumbers(root)
		// Should be a valid 32-bit integer
		assert.True(t, result >= 0 && result <= 1<<31-1)
	})

	t.Run("Tree with 1000 nodes", func(t *testing.T) {
		// Create a complete tree with 1000 nodes (maximum per constraints)
		// A complete binary tree with 10 levels has 1023 nodes
		// Let's create 10 levels but only 1000 nodes
		root := utils.CreateCompleteTree(1000)
		// All nodes have values 1-1000, but we need digits 0-9
		// Let's modify values to be 0-9
		var modifyValues func(node *utils.TreeNode)
		modifyValues = func(node *utils.TreeNode) {
			if node == nil {
				return
			}
			node.Val = node.Val % 10
			modifyValues(node.Left)
			modifyValues(node.Right)
		}
		modifyValues(root)

		result := SumNumbers(root)
		// Just verify it computes without error
		assert.True(t, result >= 0)
	})

	t.Run("All zeros tree", func(t *testing.T) {
		root := utils.CreateCompleteTree(7) // 7 nodes
		// Set all values to 0
		var setZero func(node *utils.TreeNode)
		setZero = func(node *utils.TreeNode) {
			if node == nil {
				return
			}
			node.Val = 0
			setZero(node.Left)
			setZero(node.Right)
		}
		setZero(root)

		result := SumNumbers(root)
		assert.Equal(t, 0, result)
	})

	t.Run("Tree with only leaf nodes having values", func(t *testing.T) {
		// Root and internal nodes are 0, leaves are non-zero
		root := &utils.TreeNode{Val: 0}
		root.Left = &utils.TreeNode{Val: 0}
		root.Right = &utils.TreeNode{Val: 0}
		root.Left.Left = &utils.TreeNode{Val: 1}
		root.Left.Right = &utils.TreeNode{Val: 2}
		root.Right.Left = &utils.TreeNode{Val: 3}
		root.Right.Right = &utils.TreeNode{Val: 4}

		result := SumNumbers(root)
		// Paths: 001 + 002 + 003 + 004 = 10
		assert.Equal(t, 10, result)
	})
}

func BenchmarkSumNumbers(b *testing.B) {
	// Create a tree for benchmarking
	root := utils.CreateCompleteTree(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumNumbers(root)
	}
}
