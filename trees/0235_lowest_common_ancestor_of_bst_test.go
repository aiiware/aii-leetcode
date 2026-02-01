package trees

import (
	"testing"
	
	"leetcode/utils"
)

func TestLowestCommonAncestorBST(t *testing.T) {
	// Build a sample BST: [6,2,8,0,4,7,9,-1,-1,3,5]
	//        6
	//       / \
	//      2   8
	//     / \ / \
	//    0  4 7  9
	//      / \
	//     3   5
	root := buildBST([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5})

	// Get references to nodes
	node2 := findNode(root, 2)
	node8 := findNode(root, 8)
	node4 := findNode(root, 4)
	node3 := findNode(root, 3)
	node5 := findNode(root, 5)

	tests := []struct {
		name   string
		root   *utils.TreeNode
		p      *utils.TreeNode
		q      *utils.TreeNode
		expect int
	}{
		{
			name:   "Example 1: LCA of 2 and 8 is 6",
			root:   root,
			p:      node2,
			q:      node8,
			expect: 6,
		},
		{
			name:   "Example 2: LCA of 2 and 4 is 2",
			root:   root,
			p:      node2,
			q:      node4,
			expect: 2,
		},
		{
			name:   "LCA of 3 and 5 is 4",
			root:   root,
			p:      node3,
			q:      node5,
			expect: 4,
		},
		{
			name:   "LCA of 3 and 9 is 6",
			root:   root,
			p:      node3,
			q:      findNode(root, 9),
			expect: 6,
		},
		{
			name:   "Same node",
			root:   root,
			p:      node2,
			q:      node2,
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test iterative version
			result := LowestCommonAncestorBST(tt.root, tt.p, tt.q)
			if result == nil || result.Val != tt.expect {
				t.Errorf("LowestCommonAncestorBST() = %v, want %v", getVal(result), tt.expect)
			}

			// Test recursive version
			resultRec := LowestCommonAncestorBSTRecursive(tt.root, tt.p, tt.q)
			if resultRec == nil || resultRec.Val != tt.expect {
				t.Errorf("LowestCommonAncestorBSTRecursive() = %v, want %v", getVal(resultRec), tt.expect)
			}
		})
	}
}

// Helper function to build BST from slice representation
// -1 represents null, assumes valid BST level order
func buildBST(vals []int) *utils.TreeNode {
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

// Helper to find node with given value in BST
func findNode(root *utils.TreeNode, val int) *utils.TreeNode {
	current := root
	for current != nil {
		if val == current.Val {
			return current
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

// Helper to get value or -1 if nil
func getVal(node *utils.TreeNode) int {
	if node == nil {
		return -1
	}
	return node.Val
}