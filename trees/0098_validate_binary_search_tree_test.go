package trees

import (
	"testing"
    "leetcode/utils"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected bool
	}{
		{
			name:     "Example 1: Valid BST",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)}),
			expected: true,
		},
		{
			name:     "Example 2: Invalid BST",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(4), nil, nil, utils.IntPtr(3), utils.IntPtr(6)}),
			expected: false,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Left child equal to parent",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3)}),
			expected: false,
		},
		{
			name:     "Right child equal to parent",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(2)}),
			expected: false,
		},
		{
			name:     "Left child greater than parent",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(1)}),
			expected: false,
		},
		{
			name:     "Right child less than parent",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(1)}),
			expected: false,
		},
		{
			name:     "Valid BST with negative values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Invalid BST in right subtree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), nil, nil, utils.IntPtr(6), utils.IntPtr(20)}),
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
			root:     utils.NewTreeFromSlice([]*int{
				utils.IntPtr(8), 
				utils.IntPtr(3), utils.IntPtr(10), 
				utils.IntPtr(1), utils.IntPtr(6), nil, utils.IntPtr(14),
				nil, nil, utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(13), nil, nil, nil,
			}),
			expected: true,
		},
		{
			name:     "Invalid: right child of left subtree greater than root",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(1), utils.IntPtr(5), utils.IntPtr(0), utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(6), nil, nil, nil, utils.IntPtr(3)}),
			expected: false, // 3 in left subtree equals root
		},
		{
			name:     "Skewed right valid BST",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			expected: true,
		},
		{
			name:     "Skewed left valid BST",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(4), utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Complete valid BST",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Tree with duplicate values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(2)}),
			expected: false,
		},
		{
			name:     "Minimum integer values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(-2147483648), nil, utils.IntPtr(2147483647)}),
			expected: true,
		},
		{
			name:     "Invalid: node in left subtree greater than ancestor",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(1), utils.IntPtr(8), utils.IntPtr(12), utils.IntPtr(20), nil, nil, nil, utils.IntPtr(11)}),
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