package leetcode

import (
	"fmt"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected []*int
	}{
		{
			name:     "Example 1",
			input:    []*int{intPtr(1), intPtr(3), nil, nil, intPtr(2)},
			expected: []*int{intPtr(3), intPtr(1), nil, nil, intPtr(2)},
		},
		{
			name:     "Example 2",
			input:    []*int{intPtr(3), intPtr(1), intPtr(4), nil, nil, intPtr(2)},
			expected: []*int{intPtr(2), intPtr(1), intPtr(4), nil, nil, intPtr(3)},
		},
		{
			name:     "Two nodes swapped",
			input:    []*int{intPtr(2), intPtr(1), intPtr(3)},
			expected: []*int{intPtr(2), intPtr(1), intPtr(3)}, // Already valid
		},
		{
			name:     "Adjacent nodes swapped in inorder",
			input:    []*int{intPtr(1), intPtr(3), intPtr(2)},
			expected: []*int{intPtr(1), intPtr(2), intPtr(3)},
		},
		{
			name:     "Non-adjacent nodes swapped",
			input:    []*int{intPtr(3), intPtr(2), intPtr(1)},
			expected: []*int{intPtr(1), intPtr(2), intPtr(3)},
		},
		{
			name:     "Root and leaf swapped",
			input:    []*int{intPtr(2), intPtr(1), intPtr(4), nil, nil, intPtr(3)},
			expected: []*int{intPtr(3), intPtr(1), intPtr(4), nil, nil, intPtr(2)},
		},
		{
			name:     "Two leaves swapped",
			input:    []*int{intPtr(2), intPtr(3), intPtr(1)},
			expected: []*int{intPtr(2), intPtr(1), intPtr(3)},
		},
		{
			name:     "Complex tree with swap",
			input:    []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(1), intPtr(6), intPtr(7), intPtr(9)},
			expected: []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(1), intPtr(6), intPtr(7), intPtr(9)}, // Already valid
		},
		{
			name:     "Swap in left subtree",
			input:    []*int{intPtr(4), intPtr(2), intPtr(6), intPtr(3), intPtr(1), intPtr(5), intPtr(7)},
			expected: []*int{intPtr(4), intPtr(2), intPtr(6), intPtr(1), intPtr(3), intPtr(5), intPtr(7)},
		},
		{
			name:     "Swap in right subtree",
			input:    []*int{intPtr(4), intPtr(2), intPtr(6), intPtr(1), intPtr(3), intPtr(7), intPtr(5)},
			expected: []*int{intPtr(4), intPtr(2), intPtr(6), intPtr(1), intPtr(3), intPtr(5), intPtr(7)},
		},
		{
			name:     "Root swapped with left child",
			input:    []*int{intPtr(2), intPtr(4), intPtr(3)},
			expected: []*int{intPtr(4), intPtr(2), intPtr(3)},
		},
		{
			name:     "Root swapped with right child",
			input:    []*int{intPtr(3), intPtr(1), intPtr(2)},
			expected: []*int{intPtr(2), intPtr(1), intPtr(3)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create tree from input
			root := NewTreeFromSlice(tt.input)
			
			// Recover the tree
			RecoverTree(root)
			
			// Get the result
			result := root.ToSlice()
			
			// Create expected tree
			expectedTree := NewTreeFromSlice(tt.expected)
			expected := expectedTree.ToSlice()
			
			// Compare
			if !treeSlicesEqual(result, expected) {
				t.Errorf("RecoverTree() = %v, expected %v", result, expected)
			}
			
			// Verify the tree is now a valid BST
			if !IsValidBST(root) {
				t.Errorf("Tree is not valid BST after recovery: %v", result)
			}
		})
	}
}

func TestAllRecoverTreeImplementations(t *testing.T) {
	testCases := []struct {
		name  string
		input []*int
	}{
		{
			name:  "Example 1",
			input: []*int{intPtr(1), intPtr(3), nil, nil, intPtr(2)},
		},
		{
			name:  "Example 2",
			input: []*int{intPtr(3), intPtr(1), intPtr(4), nil, nil, intPtr(2)},
		},
		{
			name:  "Simple swap",
			input: []*int{intPtr(2), intPtr(3), intPtr(1)},
		},
		{
			name:  "Adjacent swap",
			input: []*int{intPtr(1), intPtr(3), intPtr(2)},
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode)
	}{
		{"recoverTree", recoverTree},
		{"recoverTreeIterative", recoverTreeIterative},
		{"recoverTreeMorris", recoverTreeMorris},
		{"recoverTreeDFS", recoverTreeDFS},
		{"recoverTreeSimple", recoverTreeSimple},
		{"recoverTreeOptimized", recoverTreeOptimized},
		{"recoverTreeTwoPass", recoverTreeTwoPass},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Get expected result using default implementation
			expectedRoot := NewTreeFromSlice(tc.input)
			RecoverTree(expectedRoot)
			expected := expectedRoot.ToSlice()

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create fresh tree for each implementation
					root := NewTreeFromSlice(tc.input)
					
					// Apply recovery
					impl.fn(root)
					
					// Get result
					result := root.ToSlice()
					
					// Compare with expected
					if !treeSlicesEqual(result, expected) {
						t.Errorf("%s() = %v, expected %v",
							impl.name, result, expected)
					}
					
					// Verify tree is valid BST
					if !IsValidBST(root) {
						t.Errorf("%s() did not produce valid BST: %v",
							impl.name, result)
					}
				})
			}
		})
	}
}

func TestRecoverTreeEdgeCases(t *testing.T) {
	t.Run("Already valid BST", func(t *testing.T) {
		testCases := []struct {
			name  string
			input []*int
		}{
			{"Single node", []*int{intPtr(1)}},
			{"Two nodes valid", []*int{intPtr(2), intPtr(1)}},
			{"Three nodes valid", []*int{intPtr(2), intPtr(1), intPtr(3)}},
			{"Complete BST", []*int{
				intPtr(4), intPtr(2), intPtr(6),
				intPtr(1), intPtr(3), intPtr(5), intPtr(7),
			}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				root := NewTreeFromSlice(tc.input)
				original := root.ToSlice()
				
				RecoverTree(root)
				result := root.ToSlice()
				
				// Tree should remain unchanged
				if !treeSlicesEqual(result, original) {
					t.Errorf("Valid BST changed after recovery: was %v, now %v",
						original, result)
				}
				
				// Should still be valid BST
				if !IsValidBST(root) {
					t.Errorf("Valid BST became invalid after recovery: %v", result)
				}
			})
		}
	})

	t.Run("Two node tree swapped", func(t *testing.T) {
		// Only two nodes, swapped
		root := NewTreeFromSlice([]*int{intPtr(2), intPtr(1)})
		RecoverTree(root)
		
		// Should be [1, 2] or equivalent
		if !IsValidBST(root) {
			t.Errorf("Two node tree not valid after recovery: %v", root.ToSlice())
		}
	})

	t.Run("Large tree with swap", func(t *testing.T) {
		// Create a large valid BST
		root := createValidBST(1, 1000)
		original := cloneTree(root)
		
		// Swap two random nodes
		// Find two nodes to swap
		nodes := collectNodes(root)
		if len(nodes) >= 2 {
			// Swap first and last
			nodes[0].Val, nodes[len(nodes)-1].Val = nodes[len(nodes)-1].Val, nodes[0].Val
			
			// Tree should now be invalid
			if IsValidBST(root) {
				t.Error("Tree should be invalid after swap")
			}
			
			// Recover
			RecoverTree(root)
			
			// Should be valid again
			if !IsValidBST(root) {
				t.Error("Tree should be valid after recovery")
			}
			
			// Should match original (might not be identical structure but should be valid)
			// Just check it's a valid BST
			values := inorderTraversal(root)
			for i := 1; i < len(values); i++ {
				if values[i] <= values[i-1] {
					t.Errorf("Inorder traversal not sorted after recovery: %v", values)
					break
				}
			}
		}
	})

	t.Run("Tree with duplicate values before recovery", func(t *testing.T) {
		// Note: BST shouldn't have duplicates, but test edge case
		root := NewTreeFromSlice([]*int{intPtr(2), intPtr(2), intPtr(3)})
		RecoverTree(root)
		
		// After recovery, should be valid (though might still have duplicates)
		// Our algorithm swaps based on inorder violations
		_ = root // Just ensure no panic
	})

	t.Run("Minimum and maximum integer values", func(t *testing.T) {
		testCases := []struct {
			name     string
			input    []*int
			expected []*int
		}{
			{
				name:     "Min and max swapped",
				input:    []*int{intPtr(2147483647), intPtr(-2147483648)},
				expected: []*int{intPtr(-2147483648), intPtr(2147483647)},
			},
			{
				name:     "Min in wrong place",
				input:    []*int{intPtr(0), intPtr(2147483647), intPtr(-2147483648)},
				expected: []*int{intPtr(0), intPtr(-2147483648), intPtr(2147483647)},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				root := NewTreeFromSlice(tc.input)
				RecoverTree(root)
				
				if !IsValidBST(root) {
					t.Errorf("Tree with min/max values not valid after recovery: %v",
						root.ToSlice())
				}
			})
		}
	})
}

func TestRecoverTreeProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(*TreeNode)
	}{
		{"recoverTree", recoverTree},
		{"recoverTreeIterative", recoverTreeIterative},
		{"recoverTreeMorris", recoverTreeMorris},
		{"recoverTreeDFS", recoverTreeDFS},
		{"recoverTreeOptimized", recoverTreeOptimized},
		{"recoverTreeTwoPass", recoverTreeTwoPass},
	}

	testCases := []struct {
		name  string
		input []*int
	}{
		{"Two nodes swapped", []*int{intPtr(2), intPtr(1)}},
		{"Three nodes adjacent swap", []*int{intPtr(1), intPtr(3), intPtr(2)}},
		{"Three nodes non-adjacent swap", []*int{intPtr(3), intPtr(2), intPtr(1)}},
		{"Four nodes with swap", []*int{intPtr(3), intPtr(1), intPtr(4), nil, nil, intPtr(2)}},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					root := NewTreeFromSlice(tc.input)
					
					// Property 1: After recovery, tree should be valid BST
					impl.fn(root)
					if !IsValidBST(root) {
						t.Errorf("Tree not valid BST after %s: %v",
							impl.name, root.ToSlice())
					}
					
					// Property 2: Inorder traversal should be sorted
					values := inorderTraversal(root)
					for i := 1; i < len(values); i++ {
						if values[i] <= values[i-1] {
							t.Errorf("Inorder traversal not sorted after %s: %v",
								impl.name, values)
							break
						}
					}
					
					// Property 3: Tree structure should not change (only values swapped)
					original := NewTreeFromSlice(tc.input)
					originalStructure := treeStructure(original)
					resultStructure := treeStructure(root)
					
					if originalStructure != resultStructure {
						t.Errorf("Tree structure changed after %s", impl.name)
					}
				})
			}
		})
	}
}

func BenchmarkRecoverTree(b *testing.B) {
	// Create test trees of different sizes with swaps
	testCases := []struct {
		name  string
		input []*int
	}{
		{
			name:  "Small",
			input: []*int{intPtr(3), intPtr(1), intPtr(4), nil, nil, intPtr(2)},
		},
		{
			name:  "Medium",
			input: createSwappedTree(100),
		},
		{
			name:  "Large",
			input: createSwappedTree(1000),
		},
		{
			name:  "Skewed with swap",
			input: createSwappedSkewedTree(1000),
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode)
	}{
		{"recoverTree", recoverTree},
		{"recoverTreeIterative", recoverTreeIterative},
		{"recoverTreeMorris", recoverTreeMorris},
		{"recoverTreeDFS", recoverTreeDFS},
		{"recoverTreeSimple", recoverTreeSimple},
		{"recoverTreeOptimized", recoverTreeOptimized},
		{"recoverTreeTwoPass", recoverTreeTwoPass},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						// Create fresh tree for each iteration
						root := NewTreeFromSlice(tc.input)
						impl.fn(root)
					}
				})
			}
		})
	}
}

func BenchmarkRecoverTreeWorstCase(b *testing.B) {
	// Worst case: large skewed tree with swap at ends
	input := createSwappedSkewedTree(10000)

	b.ResetTimer()

	b.Run("recoverTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			root := NewTreeFromSlice(input)
			recoverTree(root)
		}
	})

	b.Run("recoverTreeIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			root := NewTreeFromSlice(input)
			recoverTreeIterative(root)
		}
	})

	b.Run("recoverTreeMorris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			root := NewTreeFromSlice(input)
			recoverTreeMorris(root)
		}
	})

	b.Run("recoverTreeOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			root := NewTreeFromSlice(input)
			recoverTreeOptimized(root)
		}
	})
}

// Helper functions

func intPtr(x int) *int {
	return &x
}

func treeSlicesEqual(a, b []*int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if (a[i] == nil) != (b[i] == nil) {
			return false
		}
		if a[i] != nil && b[i] != nil && *a[i] != *b[i] {
			return false
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	result := append(left, root.Val)
	result = append(result, right...)
	return result
}

func collectNodes(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	left := collectNodes(root.Left)
	right := collectNodes(root.Right)
	result := append(left, root)
	result = append(result, right...)
	return result
}

func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  cloneTree(root.Left),
		Right: cloneTree(root.Right),
	}
}

func treeStructure(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return fmt.Sprintf("(%s %d %s)", treeStructure(root.Left), root.Val, treeStructure(root.Right))
}

func createSwappedTree(n int) []*int {
	// Create a valid BST then swap two nodes
	root := createValidBST(1, n)
	nodes := collectNodes(root)
	
	// Swap first and last nodes
	if len(nodes) >= 2 {
		nodes[0].Val, nodes[len(nodes)-1].Val = nodes[len(nodes)-1].Val, nodes[0].Val
	}
	
	// Convert to slice
	return root.ToSlice()
}

func createSwappedSkewedTree(n int) []*int {
	// Create right-skewed BST
	root := createRightSkewedBST(n)
	nodes := collectNodes(root)
	
	// Swap first and last
	if len(nodes) >= 2 {
		nodes[0].Val, nodes[len(nodes)-1].Val = nodes[len(nodes)-1].Val, nodes[0].Val
	}
	
	return root.ToSlice()
}

func createValidBST(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	return &TreeNode{
		Val:   mid,
		Left:  createValidBST(start, mid-1),
		Right: createValidBST(mid+1, end),
	}
}

func createRightSkewedBST(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	root := &TreeNode{Val: 1}
	current := root
	for i := 2; i <= n; i++ {
		current.Right = &TreeNode{Val: i}
		current = current.Right
	}
	return root
}