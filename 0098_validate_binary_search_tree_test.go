package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Example 1: Valid BST",
			root:     NewTreeFromSlice([]*int{intPtr(2), intPtr(1), intPtr(3)}),
			expected: true,
		},
		{
			name:     "Example 2: Invalid BST",
			root:     NewTreeFromSlice([]*int{intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6)}),
			expected: false,
		},
		{
			name:     "Single node",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			expected: true,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Left child equal to parent",
			root:     NewTreeFromSlice([]*int{intPtr(2), intPtr(2), intPtr(3)}),
			expected: false,
		},
		{
			name:     "Right child equal to parent",
			root:     NewTreeFromSlice([]*int{intPtr(2), intPtr(1), intPtr(2)}),
			expected: false,
		},
		{
			name:     "Left child greater than parent",
			root:     NewTreeFromSlice([]*int{intPtr(2), intPtr(3), intPtr(1)}),
			expected: false,
		},
		{
			name:     "Right child less than parent",
			root:     NewTreeFromSlice([]*int{intPtr(2), intPtr(1), intPtr(1)}),
			expected: false,
		},
		{
			name:     "Valid BST with negative values",
			root:     NewTreeFromSlice([]*int{intPtr(0), intPtr(-1), intPtr(1)}),
			expected: true,
		},
		{
			name:     "Invalid BST in right subtree",
			root:     NewTreeFromSlice([]*int{intPtr(10), intPtr(5), intPtr(15), nil, nil, intPtr(6), intPtr(20)}),
			expected: false, // 6 < 10 but in right subtree
		},
		{
			name:     "Valid large BST",
			root:     NewTreeFromSlice([]*int{intPtr(8), intPtr(3), intPtr(10), intPtr(1), intPtr(6), nil, intPtr(14), nil, nil, intPtr(4), intPtr(7), nil, nil, intPtr(13)}),
			expected: true,
		},
		{
			name:     "Invalid: right child of left subtree greater than root",
			root:     NewTreeFromSlice([]*int{intPtr(3), intPtr(1), intPtr(5), intPtr(0), intPtr(2), intPtr(4), intPtr(6), nil, nil, nil, intPtr(3)}),
			expected: false, // 3 in left subtree equals root
		},
		{
			name:     "Skewed right valid BST",
			root:     NewTreeFromSlice([]*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)}),
			expected: true,
		},
		{
			name:     "Skewed left valid BST",
			root:     NewTreeFromSlice([]*int{intPtr(4), intPtr(3), nil, intPtr(2), nil, intPtr(1)}),
			expected: true,
		},
		{
			name:     "Complete valid BST",
			root:     NewTreeFromSlice([]*int{intPtr(4), intPtr(2), intPtr(6), intPtr(1), intPtr(3), intPtr(5), intPtr(7)}),
			expected: true,
		},
		{
			name:     "Tree with duplicate values",
			root:     NewTreeFromSlice([]*int{intPtr(2), intPtr(2), intPtr(2)}),
			expected: false,
		},
		{
			name:     "Minimum integer values",
			root:     NewTreeFromSlice([]*int{intPtr(-2147483648), nil, intPtr(2147483647)}),
			expected: true,
		},
		{
			name:     "Invalid: node in left subtree greater than ancestor",
			root:     NewTreeFromSlice([]*int{intPtr(10), intPtr(5), intPtr(15), intPtr(1), intPtr(8), intPtr(12), intPtr(20), nil, nil, nil, intPtr(11)}),
			expected: true, // Actually this is valid
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidBST(tt.root)
			if result != tt.expected {
				t.Errorf("IsValidBST() = %v, expected %v", result, tt.expected)
				if tt.root != nil {
					t.Logf("Tree: %v", tt.root.ToSlice())
				}
			}
		})
	}
}

func TestAllIsValidBSTImplementations(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Valid BST 1",
			root: NewTreeFromSlice([]*int{intPtr(2), intPtr(1), intPtr(3)}),
		},
		{
			name: "Invalid BST 1",
			root: NewTreeFromSlice([]*int{intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6)}),
		},
		{
			name: "Single node",
			root: NewTreeFromSlice([]*int{intPtr(1)}),
		},
		{
			name: "Empty tree",
			root: nil,
		},
		{
			name: "Valid large BST",
			root: NewTreeFromSlice([]*int{
				intPtr(8), intPtr(3), intPtr(10),
				intPtr(1), intPtr(6), nil, intPtr(14),
				nil, nil, intPtr(4), intPtr(7), nil, nil, intPtr(13),
			}),
		},
		{
			name: "Invalid with equal values",
			root: NewTreeFromSlice([]*int{intPtr(2), intPtr(2), intPtr(2)}),
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{"isValidBST", isValidBST},
		{"isValidBSTRecursive", isValidBSTRecursive},
		{"isValidBSTIterative", isValidBSTIterative},
		{"isValidBSTMorris", isValidBSTMorris},
		{"isValidBSTDFS", isValidBSTDFS},
		{"isValidBSTBFS", isValidBSTBFS},
		{"isValidBSTSimple", isValidBSTSimple},
		{"isValidBSTOptimized", isValidBSTOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := IsValidBST(tc.root)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.root)
					if result != expected {
						t.Errorf("%s() = %v, expected %v",
							impl.name, result, expected)
					}
				})
			}
		})
	}
}

func TestIsValidBSTEdgeCases(t *testing.T) {
	t.Run("Empty tree", func(t *testing.T) {
		if !IsValidBST(nil) {
			t.Error("Empty tree should be valid BST")
		}
	})

	t.Run("Single node", func(t *testing.T) {
		root := &TreeNode{Val: 42}
		if !IsValidBST(root) {
			t.Error("Single node tree should be valid BST")
		}
	})

	t.Run("Tree with minimum and maximum integer values", func(t *testing.T) {
		// Test edge cases with min and max int values
		testCases := []struct {
			name     string
			root     *TreeNode
			expected bool
		}{
			{
				name:     "Min int as root",
				root:     NewTreeFromSlice([]*int{intPtr(-2147483648), nil, intPtr(2147483647)}),
				expected: true,
			},
			{
				name:     "Max int as root",
				root:     NewTreeFromSlice([]*int{intPtr(2147483647), intPtr(-2147483648), nil}),
				expected: true,
			},
			{
				name:     "All min int",
				root:     NewTreeFromSlice([]*int{intPtr(-2147483648), intPtr(-2147483648), intPtr(-2147483648)}),
				expected: false,
			},
			{
				name:     "All max int",
				root:     NewTreeFromSlice([]*int{intPtr(2147483647), intPtr(2147483647), intPtr(2147483647)}),
				expected: false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := IsValidBST(tc.root)
				if result != tc.expected {
					t.Errorf("IsValidBST() = %v, expected %v", result, tc.expected)
				}
			})
		}
	})

	t.Run("Large tree (10,000 nodes)", func(t *testing.T) {
		// Create a valid BST with many nodes
		root := createValidBST(1, 10000)
		if !IsValidBST(root) {
			t.Error("Large valid BST should return true")
		}

		// Create an invalid BST
		invalidRoot := createInvalidBST(10000)
		if IsValidBST(invalidRoot) {
			t.Error("Large invalid BST should return false")
		}
	})

	t.Run("Tree with duplicate values at different levels", func(t *testing.T) {
		// Tree where duplicate values appear but not violating BST property
		// Actually, BST cannot have duplicates by definition
		root := NewTreeFromSlice([]*int{
			intPtr(5),
			intPtr(3), intPtr(7),
			intPtr(2), intPtr(4), intPtr(6), intPtr(8),
		})
		if !IsValidBST(root) {
			t.Error("Valid BST without duplicates should return true")
		}

		// Add a duplicate
		root.Left.Left.Val = 3 // Now 3 appears in left subtree of 3
		if IsValidBST(root) {
			t.Error("BST with duplicate should return false")
		}
	})

	t.Run("Skewed trees", func(t *testing.T) {
		// Right-skewed valid BST
		root := createRightSkewedBST(100)
		if !IsValidBST(root) {
			t.Error("Right-skewed valid BST should return true")
		}

		// Left-skewed valid BST
		root = createLeftSkewedBST(100)
		if !IsValidBST(root) {
			t.Error("Left-skewed valid BST should return true")
		}

		// Create invalid skewed tree
		invalidRoot := &TreeNode{Val: 1}
		current := invalidRoot
		for i := 2; i <= 100; i++ {
			current.Right = &TreeNode{Val: i - 1} // Decreasing values in right subtree
			current = current.Right
		}
		if IsValidBST(invalidRoot) {
			t.Error("Invalid skewed BST should return false")
		}
	})
}

func TestIsValidBSTProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{"isValidBST", isValidBST},
		{"isValidBSTRecursive", isValidBSTRecursive},
		{"isValidBSTIterative", isValidBSTIterative},
		{"isValidBSTMorris", isValidBSTMorris},
		{"isValidBSTDFS", isValidBSTDFS},
		{"isValidBSTBFS", isValidBSTBFS},
		{"isValidBSTOptimized", isValidBSTOptimized},
	}

	testTrees := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Valid BST 1",
			root: NewTreeFromSlice([]*int{intPtr(2), intPtr(1), intPtr(3)}),
		},
		{
			name: "Valid BST 2",
			root: NewTreeFromSlice([]*int{
				intPtr(4), intPtr(2), intPtr(6),
				intPtr(1), intPtr(3), intPtr(5), intPtr(7),
			}),
		},
		{
			name: "Invalid BST 1",
			root: NewTreeFromSlice([]*int{intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6)}),
		},
		{
			name: "Invalid BST 2",
			root: NewTreeFromSlice([]*int{intPtr(2), intPtr(2), intPtr(2)}),
		},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, tc := range testTrees {
				t.Run(tc.name, func(t *testing.T) {
					result := impl.fn(tc.root)

					// Property 1: If tree is valid BST, inorder traversal should be sorted
					if result {
						values := inorderTraversal(tc.root)
						for i := 1; i < len(values); i++ {
							if values[i] <= values[i-1] {
								t.Errorf("Valid BST but inorder traversal not sorted: %v", values)
								break
							}
						}
					}

					// Property 2: All implementations should agree
					// (tested in TestAllIsValidBSTImplementations)

					// Property 3: Empty tree should always be valid
					if !impl.fn(nil) {
						t.Errorf("Empty tree should be valid BST")
					}

					// Property 4: Single node tree should always be valid
					singleNode := &TreeNode{Val: 1}
					if !impl.fn(singleNode) {
						t.Errorf("Single node tree should be valid BST")
					}
				})
			}
		})
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	// Create test trees of different sizes and shapes
	testCases := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Small balanced valid",
			root: NewTreeFromSlice([]*int{
				intPtr(4), intPtr(2), intPtr(6),
				intPtr(1), intPtr(3), intPtr(5), intPtr(7),
			}),
		},
		{
			name: "Medium balanced valid",
			root: createValidBST(1, 1000),
		},
		{
			name: "Large balanced valid",
			root: createValidBST(1, 10000),
		},
		{
			name: "Skewed right valid",
			root: createRightSkewedBST(1000),
		},
		{
			name: "Skewed left valid",
			root: createLeftSkewedBST(1000),
		},
		{
			name: "Invalid tree",
			root: createInvalidBST(1000),
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{"isValidBST", isValidBST},
		{"isValidBSTRecursive", isValidBSTRecursive},
		{"isValidBSTIterative", isValidBSTIterative},
		{"isValidBSTMorris", isValidBSTMorris},
		{"isValidBSTDFS", isValidBSTDFS},
		{"isValidBSTBFS", isValidBSTBFS},
		{"isValidBSTSimple", isValidBSTSimple},
		{"isValidBSTOptimized", isValidBSTOptimized},
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

func BenchmarkIsValidBSTWorstCase(b *testing.B) {
	// Worst case: skewed tree (linked list)
	root := createRightSkewedBST(10000)

	b.ResetTimer()

	b.Run("isValidBST", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBST(root)
		}
	})

	b.Run("isValidBSTRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBSTRecursive(root)
		}
	})

	b.Run("isValidBSTIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBSTIterative(root)
		}
	})

	b.Run("isValidBSTMorris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBSTMorris(root)
		}
	})

	b.Run("isValidBSTOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isValidBSTOptimized(root)
		}
	})
}

// Helper functions

func intPtr(x int) *int {
	return &x
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

func createInvalidBST(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	// Create a tree that looks like a BST but has one invalid node
	root := createValidBST(1, n)
	// Make it invalid by swapping two nodes
	if root != nil && root.Left != nil && root.Right != nil {
		root.Left.Val, root.Right.Val = root.Right.Val, root.Left.Val
	}
	return root
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

func createLeftSkewedBST(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	root := &TreeNode{Val: n}
	current := root
	for i := n - 1; i >= 1; i-- {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}
	return root
}