package trees

import (
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
			input:    []*int{intPtr99(1), intPtr99(3), nil, nil, intPtr99(2)},
			expected: []*int{intPtr99(3), intPtr99(1), nil, nil, intPtr99(2)},
		},
		{
			name:     "Example 2",
			input:    []*int{intPtr99(3), intPtr99(1), intPtr99(4), nil, nil, intPtr99(2)},
			expected: []*int{intPtr99(2), intPtr99(1), intPtr99(4), nil, nil, intPtr99(3)},
		},
		{
			name:     "Two nodes swapped",
			input:    []*int{intPtr99(2), intPtr99(1), intPtr99(3)},
			expected: []*int{intPtr99(2), intPtr99(1), intPtr99(3)}, // Already valid
		},
		{
			name:     "Adjacent nodes swapped in inorder - FIXED",
			// Tree: 1 with right child 3, and 3 has left child 2
			// Inorder: 1, 2, 3 (but 2 and 3 are swapped to give 1, 3, 2)
			input:    []*int{intPtr99(1), nil, intPtr99(3), intPtr99(2)},
			expected: []*int{intPtr99(1), nil, intPtr99(2), intPtr99(3)},
		},
		{
			name:     "Non-adjacent nodes swapped - FIXED",
			// Tree: 3 with left child 2, and 2 has left child 1
			// Inorder: 1, 2, 3 (but 1 and 3 are swapped to give 3, 2, 1)
			input:    []*int{intPtr99(3), intPtr99(2), nil, intPtr99(1)},
			expected: []*int{intPtr99(1), intPtr99(2), nil, intPtr99(3)},
		},
		{
			name:     "Root and leaf swapped",
			input:    []*int{intPtr99(2), intPtr99(1), intPtr99(4), nil, nil, intPtr99(3)},
			expected: []*int{intPtr99(3), intPtr99(1), intPtr99(4), nil, nil, intPtr99(2)},
		},
		{
			name:     "Two leaves swapped - FIXED",
			// Tree: 2 with left child 3 and right child 1 (invalid BST)
			// Should become: 2 with left child 1 and right child 3
			input:    []*int{intPtr99(2), intPtr99(3), intPtr99(1)},
			expected: []*int{intPtr99(2), intPtr99(1), intPtr99(3)},
		},
		{
			name:     "Complex tree with swap",
			input:    []*int{intPtr99(5), intPtr99(3), intPtr99(8), intPtr99(1), intPtr99(6), intPtr99(7), intPtr99(9)},
			expected: []*int{intPtr99(5), intPtr99(3), intPtr99(8), intPtr99(1), intPtr99(6), intPtr99(7), intPtr99(9)}, // Already valid
		},
		{
			name:     "Swap in left subtree",
			input:    []*int{intPtr99(4), intPtr99(2), intPtr99(6), intPtr99(3), intPtr99(1), intPtr99(5), intPtr99(7)},
			expected: []*int{intPtr99(4), intPtr99(2), intPtr99(6), intPtr99(1), intPtr99(3), intPtr99(5), intPtr99(7)},
		},
		{
			name:     "Swap in right subtree",
			input:    []*int{intPtr99(4), intPtr99(2), intPtr99(6), intPtr99(1), intPtr99(3), intPtr99(7), intPtr99(5)},
			expected: []*int{intPtr99(4), intPtr99(2), intPtr99(6), intPtr99(1), intPtr99(3), intPtr99(5), intPtr99(7)},
		},
		{
			name:     "Root swapped with left child - FIXED",
			// Tree: 2 with left child 4 and right child 3 (invalid)
			// Should become: 4 with left child 2 and right child 3
			input:    []*int{intPtr99(2), intPtr99(4), intPtr99(3)},
			expected: []*int{intPtr99(4), intPtr99(2), intPtr99(3)},
		},
		{
			name:     "Root swapped with right child - FIXED",
			// Tree: 3 with left child 1 and right child 2 (invalid - right child < root)
			// Should become: 2 with left child 1 and right child 3
			input:    []*int{intPtr99(3), intPtr99(1), intPtr99(2)},
			expected: []*int{intPtr99(2), intPtr99(1), intPtr99(3)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create tree from input
			root := utils.NewTreeFromSlice(tt.input)
			
			// Recover the tree
			RecoverTree(root)
			
			// Get the result
			result := root.ToSlice()
			
			// Create expected tree
			expectedTree := utils.NewTreeFromSlice(tt.expected)
			expected := expectedTree.ToSlice()
			
			// Compare
			if !treeSlicesEqual99(result, expected) {
				t.Errorf("RecoverTree() = %v, expected %v", result, expected)
			}
			
			// Verify the tree is now a valid BST
			if !IsValidBST(root) {
				t.Errorf("Tree is not valid BST after recovery: %v", result)
			}
		})
	}
}

func intPtr99(x int) *int {
	return &x
}

func treeSlicesEqual99(a, b []*int) bool {
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