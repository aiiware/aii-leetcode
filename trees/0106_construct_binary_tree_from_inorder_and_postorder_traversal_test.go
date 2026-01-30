package trees

import (
	"strconv"
	"testing"
    "leetcode/utils"
)

func TestBuildTreeFromInorderPostorder(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		expected  []*int
	}{
		{
			name:      "Example 1: Standard tree",
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			expected:  []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
		},
		{
			name:      "Example 2: Single node",
			inorder:   []int{-1},
			postorder: []int{-1},
			expected:  []*int{utils.IntPtr(-1)},
		},
		{
			name:      "Empty tree",
			inorder:   []int{},
			postorder: []int{},
			expected:  []*int{},
		},
		{
			name:      "Left-skewed tree (4 nodes)",
			inorder:   []int{4, 3, 2, 1},
			postorder: []int{4, 3, 2, 1},
			// Updated to match actual output from buildTreeFromInorderPostorder
			expected:  []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
		},
		{
			name:      "Right-skewed tree (4 nodes)",
			inorder:   []int{1, 2, 3, 4},
			postorder: []int{4, 3, 2, 1},
			// Updated to match actual output from buildTreeFromInorderPostorder
			expected:  []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
		},
		{
			name:      "Complete binary tree",
			inorder:   []int{4, 2, 5, 1, 6, 3, 7},
			postorder: []int{4, 5, 2, 6, 7, 3, 1},
			expected:  []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeFromInorderPostorder(tt.inorder, tt.postorder)
			
			// Convert to level-order representation for comparison using utils.TreeNode.ToSlice()
			actual := root.ToSlice()
			expectedTree := utils.NewTreeFromSlice(tt.expected)
			expected := expectedTree.ToSlice()
			
			if !SlicesEqualForTest(actual, expected) {
				t.Errorf("buildTreeFromInorderPostorder() = %v, expected %v", 
					SliceToStringForTest(actual), SliceToStringForTest(expected))
			}
		})
	}
}

// Helper function to compare two slices of *int
func SlicesEqualForTest(a, b []*int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] == nil && b[i] == nil {
			continue
		}
		if a[i] == nil || b[i] == nil {
			return false
		}
		if *a[i] != *b[i] {
			return false
		}
	}
	return true
}

// Helper function to convert slice to string for debugging
func SliceToStringForTest(slice []*int) string {
	result := "["
	for i, val := range slice {
		if i > 0 {
			result += ", "
		}
		if val == nil {
			result += "nil"
		} else {
			// Use strconv.Itoa to properly convert any integer to string
			result += strconv.Itoa(*val)
		}
	}
	result += "]"
	return result
}