package leetcode

import (
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
			name:     "Valid large BST - FIXED",
			// Correct representation of:
			//         8
			//        / \
			//       3   10
			//      / \    \
			//     1   6    14
			//        / \   /
			//       4   7 13
			root:     NewTreeFromSlice([]*int{
				intPtr(8), 
				intPtr(3), intPtr(10), 
				intPtr(1), intPtr(6), nil, intPtr(14),
				nil, nil, intPtr(4), intPtr(7), intPtr(13), nil, nil, nil,
			}),
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
			expected: false,
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