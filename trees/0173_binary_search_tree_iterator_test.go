package trees

import (
	"testing"
    "leetcode/utils"
)

// Helper function to create a tree from slice representation
func createTree(nums []interface{}) *utils.TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}
	
	root := &utils.TreeNode{Val: nums[0].(int)}
	queue := []*utils.TreeNode{root}
	i := 1
	
	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]
		
		// Left child
		if i < len(nums) && nums[i] != nil {
			node.Left = &utils.TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		
		// Right child
		if i < len(nums) && nums[i] != nil {
			node.Right = &utils.TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	
	return root
}

func TestBSTIterator(t *testing.T) {
	tests := []struct {
		name     string
		tree     []interface{}
		expected []int
	}{
		{
			name:     "Example 1 from LeetCode",
			tree:     []interface{}{7, 3, 15, nil, nil, 9, 20},
			expected: []int{3, 7, 9, 15, 20},
		},
		{
			name:     "Single node tree",
			tree:     []interface{}{5},
			expected: []int{5},
		},
		{
			name:     "Left-skewed tree",
			tree:     []interface{}{5, 3, nil, 1, nil},
			expected: []int{1, 3, 5},
		},
		{
			name:     "Right-skewed tree",
			tree:     []interface{}{1, nil, 2, nil, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complete binary tree",
			tree:     []interface{}{4, 2, 6, 1, 3, 5, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Tree with duplicate values",
			tree:     []interface{}{2, 1, 3, 1, 2, 2, 4},
			expected: []int{1, 1, 2, 2, 2, 3, 4},
		},
		{
			name:     "Large tree",
			tree:     []interface{}{10, 5, 15, 3, 7, 13, 18, 1, nil, 6, 8, 12, 14, 17, 20},
			expected: []int{1, 3, 5, 6, 7, 8, 10, 12, 13, 14, 15, 17, 18, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.tree)
			iterator := BSTIteratorConstructor(root)
			
			// Test iteration
			for i, expectedVal := range tt.expected {
				if !iterator.HasNext() {
					t.Errorf("HasNext() = false, expected true at position %d", i)
				}
				
				val := iterator.Next()
				if val != expectedVal {
					t.Errorf("Next() = %d, expected %d at position %d", val, expectedVal, i)
				}
			}
			
			// Test that HasNext returns false after all elements
			if iterator.HasNext() {
				t.Errorf("HasNext() = true after all elements, expected false")
			}
		})
	}
}

func TestBSTIterator_EmptyTree(t *testing.T) {
	// Test with nil root
	iterator := BSTIteratorConstructor(nil)
	if iterator.HasNext() {
		t.Error("HasNext() should return false for empty tree")
	}
}

func TestBSTIterator_NextWithoutHasNext(t *testing.T) {
	// According to problem statement, next() calls will always be valid
	// But we should test basic functionality
	root := createTree([]interface{}{2, 1, 3})
	iterator := BSTIteratorConstructor(root)
	
	// Should get 1, 2, 3 in order
	expected := []int{1, 2, 3}
	for _, exp := range expected {
		val := iterator.Next()
		if val != exp {
			t.Errorf("Next() = %d, expected %d", val, exp)
		}
	}
}

func TestBSTIterator_MixedCalls(t *testing.T) {
	root := createTree([]interface{}{7, 3, 15, nil, nil, 9, 20})
	iterator := BSTIteratorConstructor(root)
	
	// Test mixed calls like in the example
	if !iterator.HasNext() {
		t.Error("HasNext() should return true initially")
	}
	
	val := iterator.Next()
	if val != 3 {
		t.Errorf("First Next() = %d, expected 3", val)
	}
	
	val = iterator.Next()
	if val != 7 {
		t.Errorf("Second Next() = %d, expected 7", val)
	}
	
	if !iterator.HasNext() {
		t.Error("HasNext() should return true after two calls")
	}
	
	val = iterator.Next()
	if val != 9 {
		t.Errorf("Third Next() = %d, expected 9", val)
	}
	
	if !iterator.HasNext() {
		t.Error("HasNext() should return true after three calls")
	}
	
	val = iterator.Next()
	if val != 15 {
		t.Errorf("Fourth Next() = %d, expected 15", val)
	}
	
	if !iterator.HasNext() {
		t.Error("HasNext() should return true after four calls")
	}
	
	val = iterator.Next()
	if val != 20 {
		t.Errorf("Fifth Next() = %d, expected 20", val)
	}
	
	if iterator.HasNext() {
		t.Error("HasNext() should return false after all elements")
	}
}

// Benchmark tests
func BenchmarkBSTIterator(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	// Simple balanced tree (not perfectly balanced, but good for benchmarking)
	root := createBalancedTree(1, 1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := BSTIteratorConstructor(root)
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}

func BenchmarkBSTIterator_Next(b *testing.B) {
	root := createBalancedTree(1, 1000)
	iterator := BSTIteratorConstructor(root)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset iterator for each iteration
		iterator = BSTIteratorConstructor(root)
		for j := 0; j < 100; j++ {
			if iterator.HasNext() {
				iterator.Next()
			}
		}
	}
}

// Helper to create a balanced BST for benchmarking
func createBalancedTree(start, end int) *utils.TreeNode {
	if start > end {
		return nil
	}
	
	mid := (start + end) / 2
	root := &utils.TreeNode{Val: mid}
	root.Left = createBalancedTree(start, mid-1)
	root.Right = createBalancedTree(mid+1, end)
	
	return root
}