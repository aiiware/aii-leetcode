package arrays

import "testing"

func TestPreorderTraversal(t *testing.T) {
	// Test case 1: Empty tree
	root := (*TreeNode)(nil)
	result := PreorderTraversal(root)
	expected := []int{}
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 2: Single node
	root = &TreeNode{Val: 1}
	result = PreorderTraversal(root)
	expected = []int{1}
	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test case 3: Complete tree [1,null,2,3]
	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	result = PreorderTraversal(root)
	expected = []int{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}