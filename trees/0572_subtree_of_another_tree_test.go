package trees

import (
	"testing"
	
	"leetcode/utils"
)

func TestSubtreeOfAnotherTree(t *testing.T) {
	tests := []struct {
		name   string
		root   *utils.TreeNode
		sub    *utils.TreeNode
		expect bool
	}{
		{
			name:   "Example 1",
			root:   buildTreeFromSlice([]int{3, 4, 5, 1, 2}),
			sub:    buildTreeFromSlice([]int{4, 1, 2}),
			expect: true,
		},
		{
			name:   "Example 2",
			root:   buildTreeFromSlice([]int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0}),
			sub:    buildTreeFromSlice([]int{4, 1, 2}),
			expect: false,
		},
		{
			name:   "Empty subRoot",
			root:   buildTreeFromSlice([]int{1, 2, 3}),
			sub:    nil,
			expect: true,
		},
		{
			name:   "Empty root with non-empty subRoot",
			root:   nil,
			sub:    buildTreeFromSlice([]int{1}),
			expect: false,
		},
		{
			name:   "Same tree",
			root:   buildTreeFromSlice([]int{1, 2, 3}),
			sub:    buildTreeFromSlice([]int{1, 2, 3}),
			expect: true,
		},
		{
			name:   "Subtree in left branch",
			root:   buildTreeFromSlice([]int{3, 4, 5, 1, 2, -1, -1, 0}),
			sub:    buildTreeFromSlice([]int{4, 1, 2}),
			expect: false, // Changed from true to false because node 1 has left child 0 in root
		},
		{
			name:   "Not a subtree",
			root:   buildTreeFromSlice([]int{1, 2, 3}),
			sub:    buildTreeFromSlice([]int{2, -1, 4}),
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubtreeOfAnotherTree(tt.root, tt.sub)
			if result != tt.expect {
				t.Errorf("SubtreeOfAnotherTree() = %v, want %v", result, tt.expect)
			}
		})
	}
}

// Helper function to build tree from slice representation
// -1 represents null
// Renamed to avoid conflict with buildTree in trees/0105_construct_binary_tree_from_preorder_and_inorder_traversal.go
func buildTreeFromSlice(vals []int) *utils.TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &utils.TreeNode{Val: vals[0]}
	queue := []*utils.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != -1 {
			node.Left = &utils.TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != -1 {
			node.Right = &utils.TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}