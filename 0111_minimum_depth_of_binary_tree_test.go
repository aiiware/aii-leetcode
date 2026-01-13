package leetcode

import (
	"testing"
)

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Example 1: Balanced tree",
			root:     NewTreeFromSlice([]*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)}),
			expected: 2,
		},
		{
			name:     "Example 2: Right-skewed tree",
			root:     NewTreeFromSlice([]*int{IntPtr(2), nil, IntPtr(3), nil, IntPtr(4), nil, IntPtr(5), nil, IntPtr(6)}),
			expected: 5,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node tree",
			root:     NewTreeFromSlice([]*int{IntPtr(1)}),
			expected: 1,
		},
		{
			name:     "Left-skewed tree",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)}),
			expected: 4,
		},
		{
			name:     "Tree with only left child",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), nil}),
			expected: 2,
		},
		{
			name:     "Tree with only right child",
			root:     NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2)}),
			expected: 2,
		},
		{
			name:     "Complete binary tree",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}),
			expected: 3,
		},
		{
			name:     "Tree with negative values",
			root:     NewTreeFromSlice([]*int{IntPtr(-10), IntPtr(-20), IntPtr(-30)}),
			expected: 2,
		},
		{
			name:     "Tree with zero values",
			root:     NewTreeFromSlice([]*int{IntPtr(0), IntPtr(0), IntPtr(0)}),
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test main function (BFS implementation)
			result := minDepth(tt.root)
			if result != tt.expected {
				t.Errorf("minDepth() = %v, expected %v", result, tt.expected)
			}

			// Test recursive solution
			resultRecursive := minDepthRecursive(tt.root)
			if resultRecursive != tt.expected {
				t.Errorf("minDepthRecursive() = %v, expected %v", resultRecursive, tt.expected)
			}

			// Test BFS solution (same as main function)
			resultBFS := minDepthBFS(tt.root)
			if resultBFS != tt.expected {
				t.Errorf("minDepthBFS() = %v, expected %v", resultBFS, tt.expected)
			}

			// Test DFS solution
			resultDFS := minDepthDFS(tt.root)
			if resultDFS != tt.expected {
				t.Errorf("minDepthDFS() = %v, expected %v", resultDFS, tt.expected)
			}

			// Test optimized solution
			resultOptimized := minDepthOptimized(tt.root)
			if resultOptimized != tt.expected {
				t.Errorf("minDepthOptimized() = %v, expected %v", resultOptimized, tt.expected)
			}
		})
	}
}

func TestMinDepth_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Large tree (1000 nodes in a chain)",
			root:     createChainTree(1000, 1),
			expected: 1000,
		},
		{
			name:     "Perfect binary tree height 3 (7 nodes)",
			root:     createPerfectBinaryTree(3, 1),
			expected: 3,
		},
		{
			name:     "Perfect binary tree height 4 (15 nodes)",
			root:     createPerfectBinaryTree(4, 1),
			expected: 4,
		},
		{
			name:     "Tree where left subtree is deeper but right has leaf earlier",
			root:     NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), nil, nil, nil, IntPtr(5), nil}),
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDepth(tt.root)
			if result != tt.expected {
				t.Errorf("minDepth() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMinDepth_Performance(t *testing.T) {
	// Test with a large balanced tree (2^10 - 1 = 1023 nodes)
	height := 10
	root := createPerfectBinaryTree(height, 1)
	
	// The minimum depth should be the height of the tree
	expected := height
	
	result := minDepth(root)
	if result != expected {
		t.Errorf("minDepth() = %v, expected %v for perfect binary tree of height %d", result, expected, height)
	}
}

// Helper function to create a chain tree (completely skewed)
func createChainTree(n int, startVal int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	root := &TreeNode{Val: startVal}
	current := root
	for i := 1; i < n; i++ {
		current.Right = &TreeNode{Val: startVal + i}
		current = current.Right
	}
	return root
}

// Helper function to create a perfect binary tree
func createPerfectBinaryTree(height int, startVal int) *TreeNode {
	if height <= 0 {
		return nil
	}
	
	var build func(depth int, val *int) *TreeNode
	build = func(depth int, val *int) *TreeNode {
		if depth > height {
			return nil
		}
		
		node := &TreeNode{Val: *val}
		*val++
		
		node.Left = build(depth+1, val)
		node.Right = build(depth+1, val)
		
		return node
	}
	
	val := startVal
	return build(1, &val)
}