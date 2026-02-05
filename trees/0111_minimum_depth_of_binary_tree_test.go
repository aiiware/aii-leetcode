package trees

import (
	"testing"
    "leetcode/utils"
    "leetcode/testutils"
)

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected int
	}{
		{
			name:     "Example 1: Balanced tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}),
			expected: 2,
		},
		{
			name:     "Example 2: Right-skewed tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), nil, utils.IntPtr(6)}),
			expected: 5,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: 1,
		},
		{
			name:     "Left-skewed tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			expected: 4,
		},
		{
			name:     "Tree with only left child",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil}),
			expected: 2,
		},
		{
			name:     "Tree with only right child",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2)}),
			expected: 2,
		},
		{
			name:     "Complete binary tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			expected: 3,
		},
		{
			name:     "Tree with negative values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(-10), utils.IntPtr(-20), utils.IntPtr(-30)}),
			expected: 2,
		},
		{
			name:     "Tree with zero values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0)}),
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
		root     *utils.TreeNode
		expected int
	}{
		{
			name:     "Large tree (1000 nodes in a chain)",
			root:     testutils.CreateChainTree(1000, 1),
			expected: 1000,
		},
		{
			name:     "Perfect binary tree height 3 (7 nodes)",
			root:     testutils.CreatePerfectBinaryTree(3, 1),
			expected: 3,
		},
		{
			name:     "Perfect binary tree height 4 (15 nodes)",
			root:     testutils.CreatePerfectBinaryTree(4, 1),
			expected: 4,
		},
		{
			name:     "Tree where left subtree is deeper but right has leaf earlier",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil, utils.IntPtr(5), nil}),
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
	root := testutils.CreatePerfectBinaryTree(height, 1)
	
	// The minimum depth should be the height of the tree
	expected := height
	
	result := minDepth(root)
	if result != expected {
		t.Errorf("minDepth() = %v, expected %v for perfect binary tree of height %d", result, expected, height)
	}
}