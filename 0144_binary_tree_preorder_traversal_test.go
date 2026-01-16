package leetcode

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		tree     []*int
		expected []int
	}{
		{
			name:     "Empty tree",
			tree:     []*int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			tree:     []*int{IntPtr(1)},
			expected: []int{1},
		},
		{
			name:     "Left skewed tree",
			tree:     []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Right skewed tree",
			tree:     []*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3), nil, nil, nil, nil, nil, nil, nil, IntPtr(4)},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Complete binary tree",
			tree:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:     "Example 1 from LeetCode",
			tree:     []*int{IntPtr(1), nil, IntPtr(2), nil, nil, IntPtr(3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complex tree",
			tree:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4), IntPtr(5), IntPtr(6), nil, nil, IntPtr(7), IntPtr(8)},
			expected: []int{1, 2, 4, 7, 8, 3, 5, 6},
		},
		{
			name:     "Tree with negative values",
			tree:     []*int{IntPtr(-1), IntPtr(2), IntPtr(-3), IntPtr(4), IntPtr(-5)},
			expected: []int{-1, 2, 4, -5, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeFromSlice(tt.tree)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*TreeNode) []int
			}{
				{"PreorderTraversal", PreorderTraversal},
				{"PreorderTraversalIterative", PreorderTraversalIterative},
				{"PreorderTraversalMorris", PreorderTraversalMorris},
			}

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(root)
					if !reflect.DeepEqual(result, tt.expected) {
						t.Errorf("%s() = %v, expected %v", impl.name, result, tt.expected)
					}
				})
			}
		})
	}
}

func BenchmarkPreorderTraversal(b *testing.B) {
	// Create a complete binary tree with 1023 nodes (10 levels)
	levels := 10
	totalNodes := 1<<levels - 1
	treeSlice := make([]*int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		val := i + 1
		treeSlice[i] = &val
	}
	root := NewTreeFromSlice(treeSlice)

	b.Run("Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PreorderTraversal(root)
		}
	})

	b.Run("Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PreorderTraversalIterative(root)
		}
	})

	b.Run("Morris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PreorderTraversalMorris(root)
		}
	})
}

func TestPreorderTraversalEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := PreorderTraversal(nil)
		if len(result) != 0 {
			t.Errorf("PreorderTraversal(nil) = %v, expected []", result)
		}

		result = PreorderTraversalIterative(nil)
		if len(result) != 0 {
			t.Errorf("PreorderTraversalIterative(nil) = %v, expected []", result)
		}

		result = PreorderTraversalMorris(nil)
		if len(result) != 0 {
			t.Errorf("PreorderTraversalMorris(nil) = %v, expected []", result)
		}
	})

	t.Run("Tree with only left children", func(t *testing.T) {
		// Build tree: 1 -> 2 -> 3 -> 4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}

		expected := []int{1, 2, 3, 4}

		result := PreorderTraversal(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree: got %v, expected %v", result, expected)
		}

		result = PreorderTraversalIterative(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree (iterative): got %v, expected %v", result, expected)
		}

		result = PreorderTraversalMorris(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree (Morris): got %v, expected %v", result, expected)
		}
	})

	t.Run("Tree with only right children", func(t *testing.T) {
		// Build tree: 1 -> 2 -> 3 -> 4
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 3}
		root.Right.Right.Right = &TreeNode{Val: 4}

		expected := []int{1, 2, 3, 4}

		result := PreorderTraversal(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree: got %v, expected %v", result, expected)
		}

		result = PreorderTraversalIterative(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree (iterative): got %v, expected %v", result, expected)
		}

		result = PreorderTraversalMorris(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree (Morris): got %v, expected %v", result, expected)
		}
	})

	t.Run("Verify tree structure is not modified", func(t *testing.T) {
		// Create a tree
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3}
		root.Left.Left = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 5}

		// Save original structure
		original := CloneTree(root)

		// Run all traversal methods
		PreorderTraversal(root)
		PreorderTraversalIterative(root)
		PreorderTraversalMorris(root)

		// Verify tree structure is unchanged
		if !root.Equal(original) {
			t.Error("Tree structure was modified by traversal functions")
		}
	})
}