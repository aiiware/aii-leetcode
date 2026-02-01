package trees

import (
	"testing"
	
	"leetcode/utils"
)

func TestSerializeDeserializeBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		tree *utils.TreeNode
	}{
		{
			name: "Simple tree",
			tree: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
				},
				Right: &utils.TreeNode{
					Val: 3,
					Left: &utils.TreeNode{
						Val: 4,
					},
					Right: &utils.TreeNode{
						Val: 5,
					},
				},
			},
		},
		{
			name: "Single node",
			tree: &utils.TreeNode{Val: 1},
		},
		{
			name: "Empty tree",
			tree: nil,
		},
		{
			name: "Left skewed tree",
			tree: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
					Left: &utils.TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "Right skewed tree",
			tree: &utils.TreeNode{
				Val: 1,
				Right: &utils.TreeNode{
					Val: 2,
					Right: &utils.TreeNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "Complete tree",
			tree: &utils.TreeNode{
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
		},
		{
			name: "Tree with nulls in middle",
			tree: &utils.TreeNode{
				Val: 1,
				Left: &utils.TreeNode{
					Val: 2,
				},
				Right: &utils.TreeNode{
					Val: 3,
					Left: &utils.TreeNode{
						Val: 4,
					},
					// Right is nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := Constructor()

			// Test serialization
			serialized := codec.Serialize(tt.tree)

			// Test deserialization
			deserialized := codec.Deserialize(serialized)

			// Verify the trees are equal
			if !areTreesEqual(tt.tree, deserialized) {
				t.Errorf("Serialize/Deserialize failed\nOriginal: %v\nSerialized: %s\nDeserialized: %v",
					tt.tree, serialized, deserialized)
			}

			// Test convenience functions
			serialized2 := SerializeTree(tt.tree)
			deserialized2 := DeserializeTree(serialized2)

			if !areTreesEqual(tt.tree, deserialized2) {
				t.Errorf("Convenience functions failed\nOriginal: %v\nSerialized: %s\nDeserialized: %v",
					tt.tree, serialized2, deserialized2)
			}
		})
	}
}

func TestSerializeDeserializeEdgeCases(t *testing.T) {
	codec := Constructor()

	// Test empty string
	tree := codec.Deserialize("")
	if tree != nil {
		t.Errorf("Deserialize(\"\") = %v, want nil", tree)
	}

	// Test single null
	tree = codec.Deserialize("null")
	if tree != nil {
		t.Errorf("Deserialize(\"null\") = %v, want nil", tree)
	}

	// Test serialization of nil
	serialized := codec.Serialize(nil)
	if serialized != "" {
		t.Errorf("Serialize(nil) = %s, want \"\"", serialized)
	}
}

// Helper function to check if two trees are identical
// Renamed to avoid conflict with isSameTreeLocal in trees/0572_subtree_of_another_tree.go
func areTreesEqual(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return areTreesEqual(p.Left, q.Left) && areTreesEqual(p.Right, q.Right)
}