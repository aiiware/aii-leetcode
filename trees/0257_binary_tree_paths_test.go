package trees

import (
	"leetcode/utils"
	"reflect"
	"sort"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected []string
	}{
		{
			name: "Example 1",
			root: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
					Right: &utils.TreeNode{
						Val: 5,
					},
				},
				Right: &utils.TreeNode{
					Val: 3,
				},
			},
			expected: []string{"1->2->5", "1->3"},
		},
		{
			name: "Example 2",
			root: &utils.TreeNode{
				Val: 1,
			},
			expected: []string{"1"},
		},
		{
			name: "Complete binary tree",
			root: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
					Left: &utils.TreeNode{
						Val: 4,
					},
					Right: &utils.TreeNode{
						Val: 5,
					},
				},
				Right: &utils.TreeNode{
					Val: 3,
					Left: &utils.TreeNode{
						Val: 6,
					},
					Right: &utils.TreeNode{
						Val: 7,
					},
				},
			},
			expected: []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"},
		},
		{
			name: "Single path tree",
			root: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
					Left: &utils.TreeNode{
						Val: 3,
						Left: &utils.TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: []string{"1->2->3->4"},
		},
		{
			name: "Tree with negative values",
			root: &utils.TreeNode{
				Val: -1,
				Left: &utils.TreeNode{
					Val: 2,
				},
				Right: &utils.TreeNode{
					Val: -3,
				},
			},
			expected: []string{"-1->2", "-1->-3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinaryTreePaths(tt.root)
			
			// Sort both slices for comparison
			sort.Strings(result)
			sort.Strings(tt.expected)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BinaryTreePaths() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestBinaryTreePaths_NilRoot(t *testing.T) {
	result := BinaryTreePaths(nil)
	if len(result) != 0 {
		t.Errorf("BinaryTreePaths(nil) = %v, expected empty slice", result)
	}
}