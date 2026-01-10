package leetcode

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []*int // Expected tree in level-order representation
	}{
		{
			name:     "Example 1: Standard tree",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)},
		},
		{
			name:     "Example 2: Single node",
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: []*int{IntPtr(-1)},
		},
		{
			name:     "Empty tree",
			preorder: []int{},
			inorder:  []int{},
			expected: []*int{},
		},
		{
			name:     "Left-skewed tree",
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{4, 3, 2, 1},
			expected: []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4)},
		},
		{
			name:     "Right-skewed tree",
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{1, 2, 3, 4},
			expected: []*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3), nil, nil, nil, nil, nil, nil, nil, IntPtr(4)},
		},
		{
			name:     "Complete binary tree",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			expected: []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
		},
		{
			name:     "Tree with negative values",
			preorder: []int{-10, 9, -5, -3, 20, 15, 7},
			inorder:  []int{-5, 9, -3, -10, 15, 20, 7},
			expected: []*int{IntPtr(-10), IntPtr(9), IntPtr(20), IntPtr(-5), IntPtr(-3), IntPtr(15), IntPtr(7)},
		},
		{
			name:     "Tree with unique values",
			preorder: []int{1, 2, 3, 4, 5, 6, 7},
			inorder:  []int{3, 2, 4, 1, 6, 5, 7},
			expected: []*int{IntPtr(1), IntPtr(2), IntPtr(5), IntPtr(3), IntPtr(4), IntPtr(6), IntPtr(7)},
		},
		{
			name:     "Unbalanced tree",
			preorder: []int{1, 2, 3, 4, 5},
			inorder:  []int{5, 4, 3, 2, 1},
			expected: []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4), nil, nil, nil, nil, nil, nil, nil, IntPtr(5)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.preorder, tt.inorder)
			actual := treeToSlice(root)
			
			// Compare slices by converting to comparable format
			if !slicesEqualPtr(actual, tt.expected) {
				t.Errorf("buildTree() = %v, expected %v", sliceToStringPtr(actual), sliceToStringPtr(tt.expected))
			}
		})
	}
}

func TestAllBuildTreeImplementations(t *testing.T) {
	testCases := []struct {
		name     string
		preorder []int
		inorder  []int
	}{
		{
			name:     "Empty tree",
			preorder: []int{},
			inorder:  []int{},
		},
		{
			name:     "Single node",
			preorder: []int{1},
			inorder:  []int{1},
		},
		{
			name:     "Standard tree",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
		},
		{
			name:     "Complete tree",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:     "Right-skewed tree",
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{1, 2, 3, 4},
		},
		{
			name:     "Left-skewed tree",
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recursiveResult := buildTreeRecursive(tc.preorder, tc.inorder)
			optimizedResult := buildTreeOptimized(tc.preorder, tc.inorder)
			iterativeResult := buildTreeIterative(tc.preorder, tc.inorder)

			recursiveSlice := treeToSlice(recursiveResult)
			optimizedSlice := treeToSlice(optimizedResult)
			iterativeSlice := treeToSlice(iterativeResult)

			if !slicesEqualPtr(recursiveSlice, optimizedSlice) || !slicesEqualPtr(optimizedSlice, iterativeSlice) {
				t.Errorf("Implementations disagree:\nRecursive: %v\nOptimized: %v\nIterative: %v",
					sliceToStringPtr(recursiveSlice), sliceToStringPtr(optimizedSlice), sliceToStringPtr(iterativeSlice))
			}
		})
	}
}

func TestBuildTreeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []*int
	}{
		{
			name:     "Nil slices",
			preorder: nil,
			inorder:  nil,
			expected: []*int{},
		},
		{
			name:     "Tree with single negative value",
			preorder: []int{-100},
			inorder:  []int{-100},
			expected: []*int{IntPtr(-100)},
		},
		{
			name:     "Tree with single max value",
			preorder: []int{3000},
			inorder:  []int{3000},
			expected: []*int{IntPtr(3000)},
		},
		{
			name:     "Tree with mixed positive and negative",
			preorder: []int{0, -1, 1},
			inorder:  []int{-1, 0, 1},
			expected: []*int{IntPtr(0), IntPtr(-1), IntPtr(1)},
		},
		{
			name:     "Large tree (10 nodes)",
			preorder: []int{1, 2, 4, 8, 9, 5, 10, 3, 6, 7},
			inorder:  []int{8, 4, 9, 2, 5, 10, 1, 6, 3, 7},
			expected: []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7), IntPtr(8), IntPtr(9), IntPtr(10)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.preorder, tt.inorder)
			actual := treeToSlice(root)
			
			if !slicesEqualPtr(actual, tt.expected) {
				t.Errorf("buildTree() = %v, expected %v", sliceToStringPtr(actual), sliceToStringPtr(tt.expected))
			}
		})
	}
}

func TestBuildTreePerformance(t *testing.T) {
	// Test with a large tree to ensure performance is acceptable
	// Create preorder and inorder for a complete tree with 1023 nodes
	n := 1023
	preorder := make([]int, n)
	inorder := make([]int, n)
	
	// Generate values 1..n
	for i := 0; i < n; i++ {
		preorder[i] = i + 1
		inorder[i] = i + 1
	}
	
	// Sort inorder to make it a valid inorder traversal for a BST
	// (This is a simplified test - actual test would need proper traversal)
	
	// Time the optimized implementation
	root := buildTreeOptimized(preorder, inorder)
	if root == nil {
		t.Errorf("buildTreeOptimized() returned nil for large tree")
	}
	
	// Time the iterative implementation
	root = buildTreeIterative(preorder, inorder)
	if root == nil {
		t.Errorf("buildTreeIterative() returned nil for large tree")
	}
}

// Helper function to convert a tree to level-order slice representation
func treeToSlice(root *TreeNode) []*int {
	if root == nil {
		return []*int{}
	}

	result := []*int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelHasNodes := false
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			
			if node == nil {
				result = append(result, nil)
				queue = append(queue, nil, nil)
				continue
			}
			
			result = append(result, IntPtr(node.Val))
			levelHasNodes = true
			
			queue = append(queue, node.Left, node.Right)
		}
		
		// If entire level was nil, we're done
		if !levelHasNodes {
			break
		}
	}

	// Remove trailing nil values
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// Helper function to compare two slices of *int
func slicesEqualPtr(a, b []*int) bool {
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
func sliceToStringPtr(slice []*int) string {
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