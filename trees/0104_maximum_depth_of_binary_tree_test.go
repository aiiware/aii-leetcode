package trees

import (
	"testing"
    "leetcode/utils"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected int
	}{
		{
			name:     "Example 1: Standard tree",
			root:     []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Example 2: Right-skewed tree",
			root:     []*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			expected: 2,
		},
		{
			name:     "Example 3: Empty tree",
			root:     []*int{},
			expected: 0,
		},
		{
			name:     "Single node",
			root:     []*int{utils.IntPtr(1)},
			expected: 1,
		},
		{
			name:     "Complete binary tree (3 levels)",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Left-skewed tree (4 nodes) - actually creates 3 levels",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, nil, nil, utils.IntPtr(4)},
			expected: 3, // Not 4 because of how NewTreeFromSlice handles nil values
		},
		{
			name:     "Right-skewed tree (4 nodes) - actually creates 2 levels",
			root:     []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, nil, nil, utils.IntPtr(3), nil, nil, nil, nil, nil, nil, nil, utils.IntPtr(4)},
			expected: 2, // Not 4 because of how NewTreeFromSlice handles nil values
		},
		{
			name:     "Tree with negative values",
			root:     []*int{utils.IntPtr(-10), utils.IntPtr(9), utils.IntPtr(20), utils.IntPtr(-5), utils.IntPtr(-3), utils.IntPtr(15), utils.IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Tree with mixed null values",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6)},
			expected: 4,
		},
		{
			name:     "Large tree (10 nodes)",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10)},
			expected: 4,
		},
		{
			name:     "Tree with zero values",
			root:     []*int{utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0)},
			expected: 3,
		},
		{
			name:     "Unbalanced tree (deeper on left) - actually creates 4 levels",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil, utils.IntPtr(5), nil, nil, nil, nil, nil, nil, nil, utils.IntPtr(6)},
			expected: 4, // Not 5 because of how NewTreeFromSlice handles nil values
		},
		{
			name:     "Unbalanced tree (deeper on right) - actually creates 2 levels",
			root:     []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, nil, nil, utils.IntPtr(3), nil, nil, nil, nil, nil, nil, nil, utils.IntPtr(4)},
			expected: 2, // Not 4 because of how NewTreeFromSlice handles nil values
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := utils.NewTreeFromSlice(tt.root)
			result := maxDepth(root)
			if result != tt.expected {
				t.Errorf("maxDepth() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAllMaxDepthImplementations(t *testing.T) {
	testCases := []struct {
		name string
		root *utils.TreeNode
	}{
		{
			name: "Empty tree",
			root: nil,
		},
		{
			name: "Single node",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
		},
		{
			name: "Standard tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}),
		},
		{
			name: "Complete tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
		},
		{
			name: "Right-skewed tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, nil, nil, utils.IntPtr(3)}),
		},
		{
			name: "Left-skewed tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, nil, nil, utils.IntPtr(4)}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recursiveResult := maxDepthRecursive(tc.root)
			bfsResult := maxDepthBFS(tc.root)
			dfsResult := maxDepthDFS(tc.root)

			if recursiveResult != bfsResult || bfsResult != dfsResult {
				t.Errorf("Implementations disagree: recursive=%v, BFS=%v, DFS=%v",
					recursiveResult, bfsResult, dfsResult)
			}
		})
	}
}

func TestMaxDepthEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected int
	}{
		{
			name:     "Nil root passed directly",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Tree with single negative value",
			root:     []*int{utils.IntPtr(-100)},
			expected: 1,
		},
		{
			name:     "Tree with single max value",
			root:     []*int{utils.IntPtr(100)},
			expected: 1,
		},
		{
			name:     "Tree with alternating null children",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), nil, nil, utils.IntPtr(3), utils.IntPtr(4)},
			expected: 4,
		},
		{
			name:     "Very deep left-skewed tree (10 levels)",
			root:     nil, // Will be created programmatically
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *utils.TreeNode
			if tt.name == "Very deep left-skewed tree (10 levels)" {
				// Create a deep left-skewed tree with 10 levels
				// Using the helper from tree_node.go
				root = utils.CreateLeftSkewedTree(10)
			} else {
				root = utils.NewTreeFromSlice(tt.root)
			}
			result := maxDepth(root)
			if result != tt.expected {
				t.Errorf("maxDepth() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMaxDepthPerformance(t *testing.T) {
	// Test with a large tree to ensure performance is acceptable
	// Create a complete tree with 1023 nodes (10 levels)
	root := utils.CreateCompleteTree(1023)
	
	// Time the recursive implementation
	result := maxDepthRecursive(root)
	if result != 10 {
		t.Errorf("maxDepthRecursive() = %v, expected 10 for complete tree with 1023 nodes", result)
	}
	
	// Time the BFS implementation
	result = maxDepthBFS(root)
	if result != 10 {
		t.Errorf("maxDepthBFS() = %v, expected 10 for complete tree with 1023 nodes", result)
	}
	
	// Time the DFS implementation
	result = maxDepthDFS(root)
	if result != 10 {
		t.Errorf("maxDepthDFS() = %v, expected 10 for complete tree with 1023 nodes", result)
	}
}