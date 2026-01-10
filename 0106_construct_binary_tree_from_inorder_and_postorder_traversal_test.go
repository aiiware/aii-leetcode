package leetcode

import (
	"testing"
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
			expected:  []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)},
		},
		{
			name:      "Example 2: Single node",
			inorder:   []int{-1},
			postorder: []int{-1},
			expected:  []*int{IntPtr(-1)},
		},
		{
			name:      "Empty tree",
			inorder:   []int{},
			postorder: []int{},
			expected:  []*int{},
		},
		{
			name:      "Left-skewed tree",
			inorder:   []int{4, 3, 2, 1},
			postorder: []int{4, 3, 2, 1},
			expected:  []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4)},
		},
		{
			name:      "Right-skewed tree",
			inorder:   []int{1, 2, 3, 4},
			postorder: []int{4, 3, 2, 1},
			expected:  []*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3), nil, nil, nil, nil, nil, nil, nil, IntPtr(4)},
		},
		{
			name:      "Complete binary tree",
			inorder:   []int{4, 2, 5, 1, 6, 3, 7},
			postorder: []int{4, 5, 2, 6, 7, 3, 1},
			expected:  []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeFromInorderPostorder(tt.inorder, tt.postorder)
			
			// Convert to level-order representation for comparison
			actual := TreeToLevelOrder(root)
			expectedTree := NewTreeFromSlice(tt.expected)
			expected := TreeToLevelOrder(expectedTree)
			
			if !SlicesEqualForTest(actual, expected) {
				t.Errorf("buildTreeFromInorderPostorder() = %v, expected %v", 
					SliceToStringForTest(actual), SliceToStringForTest(expected))
			}
		})
	}
}

// TreeToLevelOrder converts a tree to level-order slice representation
func TreeToLevelOrder(root *TreeNode) []*int {
	if root == nil {
		return []*int{}
	}

	result := []*int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, IntPtr(node.Val))
		queue = append(queue, node.Left, node.Right)
	}

	// Remove trailing nil values
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
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
			result += string(rune('0' + *val)) // Simple conversion for single digits
		}
	}
	result += "]"
	return result
}