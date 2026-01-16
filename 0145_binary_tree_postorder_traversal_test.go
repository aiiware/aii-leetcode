package leetcode

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
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
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "Right skewed tree",
			tree:     []*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3), nil, nil, nil, nil, nil, nil, nil, IntPtr(4)},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "Complete binary tree",
			tree:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name:     "Example 1 from LeetCode",
			tree:     []*int{IntPtr(1), nil, IntPtr(2), nil, nil, IntPtr(3)},
			expected: []int{3, 2, 1},
		},
		{
			name:     "Complex tree",
			tree:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4), IntPtr(5), IntPtr(6), nil, nil, IntPtr(7), IntPtr(8)},
			expected: []int{7, 8, 4, 2, 5, 6, 3, 1},
		},
		{
			name:     "Tree with negative values",
			tree:     []*int{IntPtr(-1), IntPtr(2), IntPtr(-3), IntPtr(4), IntPtr(-5)},
			expected: []int{4, -5, 2, -3, -1},
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
				{"PostorderTraversal", PostorderTraversal},
				{"PostorderTraversalIterative", PostorderTraversalIterative},
				{"PostorderTraversalTwoStacks", PostorderTraversalTwoStacks},
				{"PostorderTraversalReversePreorder", PostorderTraversalReversePreorder},
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

func BenchmarkPostorderTraversal(b *testing.B) {
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
			PostorderTraversal(root)
		}
	})

	b.Run("Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostorderTraversalIterative(root)
		}
	})

	b.Run("TwoStacks", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostorderTraversalTwoStacks(root)
		}
	})

	b.Run("ReversePreorder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PostorderTraversalReversePreorder(root)
		}
	})
}

func TestPostorderTraversalEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := PostorderTraversal(nil)
		if len(result) != 0 {
			t.Errorf("PostorderTraversal(nil) = %v, expected []", result)
		}

		result = PostorderTraversalIterative(nil)
		if len(result) != 0 {
			t.Errorf("PostorderTraversalIterative(nil) = %v, expected []", result)
		}

		result = PostorderTraversalTwoStacks(nil)
		if len(result) != 0 {
			t.Errorf("PostorderTraversalTwoStacks(nil) = %v, expected []", result)
		}

		result = PostorderTraversalReversePreorder(nil)
		if len(result) != 0 {
			t.Errorf("PostorderTraversalReversePreorder(nil) = %v, expected []", result)
		}
	})

	t.Run("Tree with only left children", func(t *testing.T) {
		// Build tree: 1 -> 2 -> 3 -> 4
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		root.Left.Left.Left = &TreeNode{Val: 4}

		expected := []int{4, 3, 2, 1}

		result := PostorderTraversal(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree: got %v, expected %v", result, expected)
		}

		result = PostorderTraversalIterative(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree (iterative): got %v, expected %v", result, expected)
		}

		result = PostorderTraversalTwoStacks(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree (two stacks): got %v, expected %v", result, expected)
		}

		result = PostorderTraversalReversePreorder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Left-skewed tree (reverse preorder): got %v, expected %v", result, expected)
		}
	})

	t.Run("Tree with only right children", func(t *testing.T) {
		// Build tree: 1 -> 2 -> 3 -> 4
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 3}
		root.Right.Right.Right = &TreeNode{Val: 4}

		expected := []int{4, 3, 2, 1}

		result := PostorderTraversal(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree: got %v, expected %v", result, expected)
		}

		result = PostorderTraversalIterative(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree (iterative): got %v, expected %v", result, expected)
		}

		result = PostorderTraversalTwoStacks(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree (two stacks): got %v, expected %v", result, expected)
		}

		result = PostorderTraversalReversePreorder(root)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Right-skewed tree (reverse preorder): got %v, expected %v", result, expected)
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
		PostorderTraversal(root)
		PostorderTraversalIterative(root)
		PostorderTraversalTwoStacks(root)
		PostorderTraversalReversePreorder(root)

		// Verify tree structure is unchanged
		if !root.Equal(original) {
			t.Error("Tree structure was modified by traversal functions")
		}
	})
}

func TestPostorderTraversalComparison(t *testing.T) {
	// Test that all implementations produce the same result
	testCases := []struct {
		name string
		tree []*int
	}{
		{"Empty", []*int{}},
		{"Single", []*int{IntPtr(1)}},
		{"Balanced", []*int{IntPtr(1), IntPtr(2), IntPtr(3)}},
		{"Complex", []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := NewTreeFromSlice(tc.tree)

			results := []struct {
				name string
				fn   func(*TreeNode) []int
			}{
				{"Recursive", PostorderTraversal},
				{"Iterative", PostorderTraversalIterative},
				{"TwoStacks", PostorderTraversalTwoStacks},
				{"ReversePreorder", PostorderTraversalReversePreorder},
			}

			// Get first result
			firstResult := results[0].fn(root)

			// Compare all others with first
			for i := 1; i < len(results); i++ {
				result := results[i].fn(root)
				if !reflect.DeepEqual(result, firstResult) {
					t.Errorf("%s() = %v, %s() = %v, expected same results",
						results[0].name, firstResult, results[i].name, result)
				}
			}
		})
	}
}