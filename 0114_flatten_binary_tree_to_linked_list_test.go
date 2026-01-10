package leetcode

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1: Complex tree",
			root: NewTreeFromSlice([]*int{
				intPtr(1), intPtr(2), intPtr(5),
				intPtr(3), intPtr(4), nil, intPtr(6),
			}),
			expected: NewTreeFromSlice([]*int{
				intPtr(1), nil, intPtr(2), nil, intPtr(3),
				nil, intPtr(4), nil, intPtr(5), nil, intPtr(6),
			}),
		},
		{
			name:     "Example 2: Empty tree",
			root:     nil,
			expected: nil,
		},
		{
			name:     "Example 3: Single node",
			root:     NewTreeFromSlice([]*int{intPtr(0)}),
			expected: NewTreeFromSlice([]*int{intPtr(0)}),
		},
		{
			name: "Example 4: Left-skewed tree",
			root: NewTreeFromSlice([]*int{
				intPtr(1), intPtr(2), nil,
				intPtr(3), nil,
			}),
			expected: NewTreeFromSlice([]*int{
				intPtr(1), nil, intPtr(2), nil, intPtr(3),
			}),
		},
		{
			name: "Example 5: Right-skewed tree",
			// Tree: 1 -> 2 -> 3
			// Complete representation: [1, nil, 2, nil, nil, nil, 3]
			// But NewTreeFromSlice doesn't handle nil nodes in queue
			// So we need to create the tree manually
			root: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 2}
				root.Right.Right = &TreeNode{Val: 3}
				return root
			}(),
			expected: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 2}
				root.Right.Right = &TreeNode{Val: 3}
				return root
			}(),
		},
		{
			name: "Example 6: Complete binary tree",
			root: NewTreeFromSlice([]*int{
				intPtr(1), intPtr(2), intPtr(3),
				intPtr(4), intPtr(5), intPtr(6), intPtr(7),
			}),
			expected: NewTreeFromSlice([]*int{
				intPtr(1), nil, intPtr(2), nil, intPtr(4),
				nil, intPtr(5), nil, intPtr(3), nil, intPtr(6), nil, intPtr(7),
			}),
		},
		{
			name: "Example 7: Tree with negative values",
			root: NewTreeFromSlice([]*int{
				intPtr(-1), intPtr(-2), intPtr(-3),
				intPtr(-4), intPtr(-5), intPtr(-6), intPtr(-7),
			}),
			expected: NewTreeFromSlice([]*int{
				intPtr(-1), nil, intPtr(-2), nil, intPtr(-4),
				nil, intPtr(-5), nil, intPtr(-3), nil, intPtr(-6), nil, intPtr(-7),
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test recursive flatten
			root1 := copyTree(tt.root)
			flatten(root1)
			if !treesEqualFlattened(root1, tt.expected) {
				t.Errorf("flatten() failed\nGot values: %v\nExpected values: %v",
					getFlattenedValues(root1), getFlattenedValues(tt.expected))
			}

			// Test iterative flatten
			root2 := copyTree(tt.root)
			flattenIterative(root2)
			if !treesEqualFlattened(root2, tt.expected) {
				t.Errorf("flattenIterative() failed\nGot values: %v\nExpected values: %v",
					getFlattenedValues(root2), getFlattenedValues(tt.expected))
			}

			// Test Morris flatten
			root3 := copyTree(tt.root)
			flattenMorris(root3)
			if !treesEqualFlattened(root3, tt.expected) {
				t.Errorf("flattenMorris() failed\nGot values: %v\nExpected values: %v",
					getFlattenedValues(root3), getFlattenedValues(tt.expected))
			}

			// Test reverse post-order flatten
			root4 := copyTree(tt.root)
			flattenReversePostOrder(root4)
			if !treesEqualFlattened(root4, tt.expected) {
				t.Errorf("flattenReversePostOrder() failed\nGot values: %v\nExpected values: %v",
					getFlattenedValues(root4), getFlattenedValues(tt.expected))
			}
		})
	}
}

func TestFlatten_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected *TreeNode
	}{
		{
			name: "Already flattened tree",
			root: NewTreeFromSlice([]*int{
				intPtr(1), nil, intPtr(2), nil, intPtr(3),
			}),
			expected: NewTreeFromSlice([]*int{
				intPtr(1), nil, intPtr(2), nil, intPtr(3),
			}),
		},
		{
			name: "Tree with only left children",
			root: NewTreeFromSlice([]*int{
				intPtr(3), intPtr(2), nil,
				intPtr(1), nil,
			}),
			expected: NewTreeFromSlice([]*int{
				intPtr(3), nil, intPtr(2), nil, intPtr(1),
			}),
		},
		{
			name: "Tree with only right children",
			root: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 2}
				root.Right.Right = &TreeNode{Val: 3}
				return root
			}(),
			expected: func() *TreeNode {
				root := &TreeNode{Val: 1}
				root.Right = &TreeNode{Val: 2}
				root.Right.Right = &TreeNode{Val: 3}
				return root
			}(),
		},
		{
			name: "Large tree",
			root: func() *TreeNode {
				// Create a tree with 15 nodes
				vals := make([]*int, 15)
				for i := range vals {
					vals[i] = intPtr(i + 1)
				}
				return NewTreeFromSlice(vals)
			}(),
			expected: func() *TreeNode {
				// Expected flattened list: 1,2,4,8,9,5,10,11,3,6,12,13,7,14,15
				// This is the pre-order traversal of a complete binary tree
				vals := []*int{
					intPtr(1), nil, intPtr(2), nil, intPtr(4), nil, intPtr(8),
					nil, intPtr(9), nil, intPtr(5), nil, intPtr(10), nil, intPtr(11),
					nil, intPtr(3), nil, intPtr(6), nil, intPtr(12), nil, intPtr(13),
					nil, intPtr(7), nil, intPtr(14), nil, intPtr(15),
				}
				return NewTreeFromSlice(vals)
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*TreeNode)
			}{
				{"flatten", flatten},
				{"flattenIterative", flattenIterative},
				{"flattenMorris", flattenMorris},
				{"flattenReversePostOrder", flattenReversePostOrder},
			}

			for _, impl := range implementations {
				rootCopy := copyTree(tt.root)
				impl.fn(rootCopy)
				if !treesEqualFlattened(rootCopy, tt.expected) {
					t.Errorf("%s() failed for %s\nGot values: %v\nExpected values: %v",
						impl.name, tt.name, getFlattenedValues(rootCopy), getFlattenedValues(tt.expected))
				}
			}
		})
	}
}

func TestFlatten_PreservesValues(t *testing.T) {
	// Create a complex tree
	root := NewTreeFromSlice([]*int{
		intPtr(5), intPtr(4), intPtr(8),
		intPtr(11), nil, intPtr(13), intPtr(4),
		intPtr(7), intPtr(2), nil, nil, intPtr(5), intPtr(1),
	})

	// Get pre-order traversal before flattening
	preOrderBefore := preOrderTraversal(root)

	// Flatten the tree
	flatten(root)

	// Get values from flattened list
	flattenedValues := []int{}
	current := root
	for current != nil {
		flattenedValues = append(flattenedValues, current.Val)
		current = current.Right
	}

	// Check that flattened list matches pre-order traversal
	if !reflect.DeepEqual(flattenedValues, preOrderBefore) {
		t.Errorf("Flattened list doesn't match pre-order traversal\nFlattened: %v\nPre-order: %v",
			flattenedValues, preOrderBefore)
	}

	// Check that all left pointers are nil
	current = root
	for current != nil {
		if current.Left != nil {
			t.Errorf("Node with value %d has non-nil left pointer", current.Val)
		}
		current = current.Right
	}
}

// Helper functions

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

// treesEqualFlattened checks if two flattened trees (right-linked lists) are equal
func treesEqualFlattened(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return treesEqualFlattened(a.Right, b.Right) // Only check right since left should be nil
}

func getFlattenedValues(root *TreeNode) []int {
	result := []int{}
	current := root
	for current != nil {
		result = append(result, current.Val)
		current = current.Right
	}
	return result
}

func preOrderTraversal(root *TreeNode) []int {
	result := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}