package leetcode

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Example 1",
			root:     NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), IntPtr(3)}),
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Example 3",
			root:     NewTreeFromSlice([]*int{IntPtr(1)}),
			expected: []int{1},
		},
		{
			name:     "Complete binary tree",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}),
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:     "Left skewed tree",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)}),
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "Right skewed tree",
			root:     NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)}),
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Tree with negative values",
			root:     NewTreeFromSlice([]*int{IntPtr(-10), IntPtr(5), IntPtr(-20), IntPtr(-3), IntPtr(0)}),
			expected: []int{-3, 5, 0, -10, -20},
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 42},
			expected: []int{42},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Balanced tree",
			root:     NewTreeFromSlice([]*int{IntPtr(10), IntPtr(5), IntPtr(15), IntPtr(3), IntPtr(7), IntPtr(12), IntPtr(18)}),
			expected: []int{3, 5, 7, 10, 12, 15, 18},
		},
		{
			name:     "Tree with duplicates",
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(1), IntPtr(2), IntPtr(1), IntPtr(2)}),
			expected: []int{1, 1, 2, 2, 2},
		},
		{
			name:     "Complex tree 1",
			root:     NewTreeFromSlice([]*int{IntPtr(5), IntPtr(3), IntPtr(8), IntPtr(2), IntPtr(4), IntPtr(6), IntPtr(9), IntPtr(1), nil, nil, nil, nil, nil, IntPtr(7)}),
			expected: []int{1, 2, 3, 4, 5, 6, 8, 7, 9}, // Fixed: 7 is left child of 9, so order is 6, 8, 7, 9
		},
		{
			name:     "Tree with null values in middle",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4), nil, IntPtr(5)}),
			expected: []int{2, 4, 1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InorderTraversal(tt.root)
			if !SlicesEqual(result, tt.expected) {
				t.Errorf("InorderTraversal() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAllInorderTraversalImplementations(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Example 1",
			root: NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), IntPtr(3)}),
		},
		{
			name: "Complete tree",
			root: NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}),
		},
		{
			name: "Left skewed",
			root: NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), nil, IntPtr(3)}),
		},
		{
			name: "Right skewed",
			root: NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3)}),
		},
		{
			name: "Single node",
			root: &TreeNode{Val: 5},
		},
		{
			name: "Empty tree",
			root: nil,
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{"inorderTraversal", inorderTraversal},
		{"inorderTraversalIterative", inorderTraversalIterative},
		{"inorderTraversalMorris", inorderTraversalMorris},
		{"inorderTraversalDFS", inorderTraversalDFS},
		{"inorderTraversalSimple", inorderTraversalSimple},
		{"inorderTraversalOptimized", inorderTraversalOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := InorderTraversal(tc.root)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.root)
					if !SlicesEqual(result, expected) {
						t.Errorf("%s() = %v, expected %v",
							impl.name, result, expected)
					}
				})
			}
		})
	}
}

func TestInorderTraversalEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		result := InorderTraversal(nil)
		if len(result) != 0 {
			t.Errorf("InorderTraversal(nil) should return empty slice, got %v", result)
		}
	})

	t.Run("Single node tree", func(t *testing.T) {
		root := &TreeNode{Val: 100}
		result := InorderTraversal(root)
		expected := []int{100}
		if !SlicesEqual(result, expected) {
			t.Errorf("InorderTraversal(single node) = %v, expected %v", result, expected)
		}
	})

	t.Run("Large tree", func(t *testing.T) {
		// Create a tree with 100 nodes (complete binary tree of height 7)
		nodes := make([]*int, 127) // 2^7 - 1 = 127 nodes for complete tree
		for i := range nodes {
			val := i + 1
			nodes[i] = &val
		}
		root := NewTreeFromSlice(nodes)

		result := InorderTraversal(root)
		
		// Check length
		expectedLength := 127
		if len(result) != expectedLength {
			t.Errorf("Expected %d nodes, got %d", expectedLength, len(result))
		}

		// Note: A complete binary tree created from array is NOT a BST,
		// so inorder traversal won't be sorted. Removing the sorted check.
		// Just verify we got all values 1..127
		valueCount := make(map[int]bool)
		for _, val := range result {
			valueCount[val] = true
		}
		for i := 1; i <= 127; i++ {
			if !valueCount[i] {
				t.Errorf("Missing value %d in traversal", i)
			}
		}
	})

	t.Run("Tree with all same values", func(t *testing.T) {
		root := NewTreeFromSlice([]*int{
			IntPtr(5), IntPtr(5), IntPtr(5),
			IntPtr(5), IntPtr(5), IntPtr(5), IntPtr(5),
		})
		result := InorderTraversal(root)
		
		// Should have 7 nodes all with value 5
		if len(result) != 7 {
			t.Errorf("Expected 7 nodes, got %d", len(result))
		}
		for _, val := range result {
			if val != 5 {
				t.Errorf("Expected all 5s, got %d", val)
			}
		}
	})

	t.Run("Unbalanced tree", func(t *testing.T) {
		// Create a very unbalanced tree (linked list)
		root := &TreeNode{Val: 1}
		current := root
		for i := 2; i <= 100; i++ {
			current.Right = &TreeNode{Val: i}
			current = current.Right
		}

		result := InorderTraversal(root)
		
		// Should be 1, 2, 3, ..., 100
		if len(result) != 100 {
			t.Errorf("Expected 100 nodes, got %d", len(result))
		}
		for i, val := range result {
			if val != i+1 {
				t.Errorf("At index %d: expected %d, got %d", i, i+1, val)
			}
		}
	})
}

func TestInorderTraversalProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{"inorderTraversal", inorderTraversal},
		{"inorderTraversalIterative", inorderTraversalIterative},
		{"inorderTraversalMorris", inorderTraversalMorris},
		{"inorderTraversalDFS", inorderTraversalDFS},
		{"inorderTraversalOptimized", inorderTraversalOptimized},
	}

	testTrees := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Simple tree",
			root: NewTreeFromSlice([]*int{IntPtr(2), IntPtr(1), IntPtr(3)}),
		},
		{
			name: "Larger tree",
			root: NewTreeFromSlice([]*int{
				IntPtr(4), IntPtr(2), IntPtr(6),
				IntPtr(1), IntPtr(3), IntPtr(5), IntPtr(7),
			}),
		},
		{
			name: "Skewed left",
			root: NewTreeFromSlice([]*int{IntPtr(3), IntPtr(2), nil, IntPtr(1)}),
		},
		{
			name: "Skewed right",
			root: NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3)}),
		},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, tc := range testTrees {
				t.Run(tc.name, func(t *testing.T) {
					result := impl.fn(tc.root)

					// Property 1: Number of nodes in traversal should match tree size
					expectedSize := countNodes(tc.root)
					if len(result) != expectedSize {
						t.Errorf("Expected %d nodes, got %d", expectedSize, len(result))
					}

					// Property 2: For BST, inorder traversal should be sorted
					if isBST(tc.root) {
						for i := 1; i < len(result); i++ {
							if result[i] < result[i-1] {
								t.Errorf("Result not sorted at index %d: %d < %d",
									i, result[i], result[i-1])
							}
						}
					}

					// Property 3: All values should be in the tree
					valuesInTree := getAllValues(tc.root)
					for _, val := range result {
						if !contains(valuesInTree, val) {
							t.Errorf("Value %d in result not found in tree", val)
						}
					}

					// Property 4: No duplicates beyond what's in the tree
					resultCount := make(map[int]int)
					treeCount := make(map[int]int)
					
					for _, val := range result {
						resultCount[val]++
					}
					countValues(tc.root, treeCount)
					
					for val, count := range resultCount {
						if count != treeCount[val] {
							t.Errorf("Value %d appears %d times in result, %d times in tree",
								val, count, treeCount[val])
						}
					}
				})
			}
		})
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	// Create test trees of different sizes and shapes
	testCases := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Small balanced",
			root: NewTreeFromSlice([]*int{
				IntPtr(4), IntPtr(2), IntPtr(6),
				IntPtr(1), IntPtr(3), IntPtr(5), IntPtr(7),
			}),
		},
		{
			name: "Medium balanced",
			root: CreateCompleteTree(31), // 2^5 - 1 nodes
		},
		{
			name: "Large balanced",
			root: CreateCompleteTree(127), // 2^7 - 1 nodes
		},
		{
			name: "Skewed right",
			root: CreateRightSkewedTree(100),
		},
		{
			name: "Skewed left",
			root: CreateLeftSkewedTree(100),
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{"inorderTraversal", inorderTraversal},
		{"inorderTraversalIterative", inorderTraversalIterative},
		{"inorderTraversalMorris", inorderTraversalMorris},
		{"inorderTraversalDFS", inorderTraversalDFS},
		{"inorderTraversalOptimized", inorderTraversalOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.root)
					}
				})
			}
		})
	}
}

func BenchmarkInorderTraversalWorstCase(b *testing.B) {
	// Worst case: skewed tree (linked list)
	root := CreateRightSkewedTree(1000)

	b.ResetTimer()

	b.Run("inorderTraversal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversal(root)
		}
	})

	b.Run("inorderTraversalIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalIterative(root)
		}
	})

	b.Run("inorderTraversalMorris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalMorris(root)
		}
	})

	b.Run("inorderTraversalOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			inorderTraversalOptimized(root)
		}
	})
}

// Helper functions for property-based testing

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func isBST(root *TreeNode) bool {
	return isBSTHelper(root, nil, nil)
}

func isBSTHelper(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}
	
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	
	return isBSTHelper(node.Left, min, &node.Val) && 
	       isBSTHelper(node.Right, &node.Val, max)
}

func getAllValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	left := getAllValues(root.Left)
	right := getAllValues(root.Right)
	
	result := make([]int, 0, len(left)+1+len(right))
	result = append(result, left...)
	result = append(result, root.Val)
	result = append(result, right...)
	
	return result
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func countValues(root *TreeNode, counts map[int]int) {
	if root == nil {
		return
	}
	
	counts[root.Val]++
	countValues(root.Left, counts)
	countValues(root.Right, counts)
}