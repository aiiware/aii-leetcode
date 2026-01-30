package trees

import (
	"reflect"
	"testing"
    "leetcode/utils"
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
			tree:     []*int{utils.IntPtr(1)},
			expected: []int{1},
		},
		{
			name:     "Left skewed tree",
			tree:     []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Right skewed tree",
			tree:     []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Complete binary tree",
			tree:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:     "Example 1 from LeetCode",
			tree:     []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complex tree",
			tree:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8)},
			expected: []int{1, 2, 4, 7, 8, 3, 5, 6},
		},
		{
			name:     "Tree with negative values",
			tree:     []*int{utils.IntPtr(-1), utils.IntPtr(2), utils.IntPtr(-3), utils.IntPtr(4), utils.IntPtr(-5)},
			expected: []int{-1, 2, 4, -5, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := utils.NewTreeFromSlice(tt.tree)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*utils.TreeNode) []int
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
	root := utils.NewTreeFromSlice(treeSlice)

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
		root := &utils.TreeNode{Val: 1}
		root.Left = &utils.TreeNode{Val: 2}
		root.Left.Left = &utils.TreeNode{Val: 3}
		root.Left.Left.Left = &utils.TreeNode{Val: 4}

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
		root := &utils.TreeNode{Val: 1}
		root.Right = &utils.TreeNode{Val: 2}
		root.Right.Right = &utils.TreeNode{Val: 3}
		root.Right.Right.Right = &utils.TreeNode{Val: 4}

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
		root := &utils.TreeNode{Val: 1}
		root.Left = &utils.TreeNode{Val: 2}
		root.Right = &utils.TreeNode{Val: 3}
		root.Left.Left = &utils.TreeNode{Val: 4}
		root.Left.Right = &utils.TreeNode{Val: 5}

		// Save original structure
		original := utils.CloneTree(root)

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