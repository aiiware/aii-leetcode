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
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(1), IntPtr(3)}),
			expected: true,
		},
		{
			name:     "Example 2: Invalid BST",
			root:     NewTreeFromSlice([]*int{IntPtr(5), IntPtr(1), IntPtr(4), nil, nil, IntPtr(3), IntPtr(6)}),
			expected: false,
		},
		{
			name:     "Single node",
			root:     NewTreeFromSlice([]*int{IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Left child equal to parent",
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(2), IntPtr(3)}),
			expected: false,
		},
		{
			name:     "Right child equal to parent",
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(1), IntPtr(2)}),
			expected: false,
		},
		{
			name:     "Left child greater than parent",
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(3), IntPtr(1)}),
			expected: false,
		},
		{
			name:     "Right child less than parent",
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(1), IntPtr(1)}),
			expected: false,
		},
		{
			name:     "Valid BST with negative values",
			root:     NewTreeFromSlice([]*int{IntPtr(0), IntPtr(-1), IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Invalid BST in right subtree",
			root:     NewTreeFromSlice([]*int{IntPtr(10), IntPtr(5), IntPtr(15), nil, nil, IntPtr(6), IntPtr(20)}),
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
				IntPtr(8), 
				IntPtr(3), IntPtr(10), 
				IntPtr(1), IntPtr(6), nil, IntPtr(14),
				nil, nil, IntPtr(4), IntPtr(7), IntPtr(13), nil, nil, nil,
			}),
			expected: true,
		},
		{
			name:     "Invalid: right child of left subtree greater than root",
			root:     NewTreeFromSlice([]*int{IntPtr(3), IntPtr(1), IntPtr(5), IntPtr(0), IntPtr(2), IntPtr(4), IntPtr(6), nil, nil, nil, IntPtr(3)}),
			expected: false, // 3 in left subtree equals root
		},
		{
			name:     "Skewed right valid BST",
			root:     NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)}),
			expected: true,
		},
		{
			name:     "Skewed left valid BST",
			root:     NewTreeFromSlice([]*int{IntPtr(4), IntPtr(3), nil, IntPtr(2), nil, IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Complete valid BST",
			root:     NewTreeFromSlice([]*int{IntPtr(4), IntPtr(2), IntPtr(6), IntPtr(1), IntPtr(3), IntPtr(5), IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Tree with duplicate values",
			root:     NewTreeFromSlice([]*int{IntPtr(2), IntPtr(2), IntPtr(2)}),
			expected: false,
		},
		{
			name:     "Minimum integer values",
			root:     NewTreeFromSlice([]*int{IntPtr(-2147483648), nil, IntPtr(2147483647)}),
			expected: true,
		},
		{
			name:     "Invalid: node in left subtree greater than ancestor",
			root:     NewTreeFromSlice([]*int{IntPtr(10), IntPtr(5), IntPtr(15), IntPtr(1), IntPtr(8), IntPtr(12), IntPtr(20), nil, nil, nil, IntPtr(11)}),
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