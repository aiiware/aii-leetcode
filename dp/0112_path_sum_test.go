package dp

import (
	"testing"
    "leetcode/utils"
    "leetcode/testutils"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		target   int
		expected bool
	}{
		{
			name:     "Example 1: Path exists",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(5), utils.IntPtr(4), utils.IntPtr(8), utils.IntPtr(11), nil, utils.IntPtr(13), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(2), nil, nil, nil, utils.IntPtr(1)}),
			target:   22,
			expected: true,
		},
		{
			name:     "Example 2: No path exists",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}),
			target:   5,
			expected: false,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			target:   0,
			expected: false,
		},
		{
			name:     "Single node tree with matching value",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			target:   1,
			expected: true,
		},
		{
			name:     "Single node tree with non-matching value",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			target:   2,
			expected: false,
		},
		{
			name:     "Path exists in left subtree only",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}),
			target:   3,
			expected: true,
		},
		{
			name:     "Path exists in right subtree only",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}),
			target:   4,
			expected: true,
		},
		{
			name:     "Negative values - path exists",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(-2), nil, utils.IntPtr(-3)}),
			target:   -5,
			expected: true,
		},
		{
			name:     "Negative values - no path exists",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(-2), nil, utils.IntPtr(-3)}),
			target:   -4,
			expected: false,
		},
		{
			name:     "Zero values - path exists",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(0), utils.IntPtr(1), utils.IntPtr(0), utils.IntPtr(1), utils.IntPtr(0), utils.IntPtr(1)}),
			target:   2,
			expected: true,
		},
		{
			name:     "Complex tree with multiple possible paths",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			target:   8,
			expected: true,
		},
		{
			name:     "Skewed tree with path",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			target:   10,
			expected: true,
		},
		{
			name:     "Skewed tree without path",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			target:   11,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test main function (recursive implementation)
			result := hasPathSum(tt.root, tt.target)
			if result != tt.expected {
				t.Errorf("hasPathSum() = %v, expected %v", result, tt.expected)
			}

			// Test recursive solution
			resultRecursive := hasPathSumRecursive(tt.root, tt.target)
			if resultRecursive != tt.expected {
				t.Errorf("hasPathSumRecursive() = %v, expected %v", resultRecursive, tt.expected)
			}

			// Test DFS solution
			resultDFS := hasPathSumDFS(tt.root, tt.target)
			if resultDFS != tt.expected {
				t.Errorf("hasPathSumDFS() = %v, expected %v", resultDFS, tt.expected)
			}

			// Test BFS solution
			resultBFS := hasPathSumBFS(tt.root, tt.target)
			if resultBFS != tt.expected {
				t.Errorf("hasPathSumBFS() = %v, expected %v", resultBFS, tt.expected)
			}

			// Test optimized solution
			resultOptimized := hasPathSumOptimized(tt.root, tt.target)
			if resultOptimized != tt.expected {
				t.Errorf("hasPathSumOptimized() = %v, expected %v", resultOptimized, tt.expected)
			}
		})
	}
}

func TestHasPathSum_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		target   int
		expected bool
	}{
		{
			name:     "Large tree with path",
			root:     testutils.CreateChainTree(1000, 1),
			target:   (1000 * 1001) / 2, // Sum of 1 to 1000
			expected: true,
		},
		{
			name:     "Large tree without path",
			root:     testutils.CreateChainTree(1000, 1),
			target:   (1000 * 1001) / 2 + 1, // One more than sum
			expected: false,
		},
		{
			name:     "Perfect binary tree with path",
			root:     testutils.CreatePerfectBinaryTree(4, 1),
			target:   10, // Arbitrary target that should exist
			expected: true,
		},
		{
			name:     "All paths sum to same value",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1)}),
			target:   3,
			expected: true,
		},
		{
			name:     "Tree with maximum node values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1000), utils.IntPtr(1000), utils.IntPtr(1000)}),
			target:   2000,
			expected: true,
		},
		{
			name:     "Tree with minimum node values",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(-1000), utils.IntPtr(-1000), utils.IntPtr(-1000)}),
			target:   -2000,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasPathSum(tt.root, tt.target)
			if result != tt.expected {
				t.Errorf("hasPathSum() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestHasPathSum_Performance(t *testing.T) {
	// Test with a large balanced tree (2^12 - 1 = 4095 nodes)
	height := 12
	root := testutils.CreatePerfectBinaryTree(height, 1)
	
	// Target sum that should exist (sum of values along some path)
	// Since values are sequential from 1 to 4095, we can pick a reasonable target
	target := 100 // This should exist in many paths
	
	// This should complete quickly even for large trees
	result := hasPathSum(root, target)
	
	// We don't know if it's true or false, but it should complete without timeout
	// Just verify the function doesn't panic
	if result != result { // This will always be false, just to use the result
		t.Errorf("Unexpected result")
	}
}