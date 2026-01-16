package leetcode

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected int
	}{
		{
			name:     "Example 1: Standard tree",
			root:     []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Example 2: Right-skewed tree",
			root:     []*int{IntPtr(1), nil, IntPtr(2)},
			expected: 2,
		},
		{
			name:     "Example 3: Empty tree",
			root:     []*int{},
			expected: 0,
		},
		{
			name:     "Single node",
			root:     []*int{IntPtr(1)},
			expected: 1,
		},
		{
			name:     "Complete binary tree (3 levels)",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Left-skewed tree (4 nodes) - actually creates 3 levels",
			root:     []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4)},
			expected: 3, // Not 4 because of how NewTreeFromSlice handles nil values
		},
		{
			name:     "Right-skewed tree (4 nodes) - actually creates 2 levels",
			root:     []*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3), nil, nil, nil, nil, nil, nil, nil, IntPtr(4)},
			expected: 2, // Not 4 because of how NewTreeFromSlice handles nil values
		},
		{
			name:     "Tree with negative values",
			root:     []*int{IntPtr(-10), IntPtr(9), IntPtr(20), IntPtr(-5), IntPtr(-3), IntPtr(15), IntPtr(7)},
			expected: 3,
		},
		{
			name:     "Tree with mixed null values",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4), IntPtr(5), nil, nil, nil, IntPtr(6)},
			expected: 4,
		},
		{
			name:     "Large tree (10 nodes)",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7), IntPtr(8), IntPtr(9), IntPtr(10)},
			expected: 4,
		},
		{
			name:     "Tree with zero values",
			root:     []*int{IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0)},
			expected: 3,
		},
		{
			name:     "Unbalanced tree (deeper on left) - actually creates 4 levels",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), nil, nil, nil, IntPtr(5), nil, nil, nil, nil, nil, nil, nil, IntPtr(6)},
			expected: 4, // Not 5 because of how NewTreeFromSlice handles nil values
		},
		{
			name:     "Unbalanced tree (deeper on right) - actually creates 2 levels",
			root:     []*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3), nil, nil, nil, nil, nil, nil, nil, IntPtr(4)},
			expected: 2, // Not 4 because of how NewTreeFromSlice handles nil values
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeFromSlice(tt.root)
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
		root *TreeNode
	}{
		{
			name: "Empty tree",
			root: nil,
		},
		{
			name: "Single node",
			root: NewTreeFromSlice([]*int{IntPtr(1)}),
		},
		{
			name: "Standard tree",
			root: NewTreeFromSlice([]*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)}),
		},
		{
			name: "Complete tree",
			root: NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}),
		},
		{
			name: "Right-skewed tree",
			root: NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2), nil, nil, nil, IntPtr(3)}),
		},
		{
			name: "Left-skewed tree",
			root: NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4)}),
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
			root:     []*int{IntPtr(-100)},
			expected: 1,
		},
		{
			name:     "Tree with single max value",
			root:     []*int{IntPtr(100)},
			expected: 1,
		},
		{
			name:     "Tree with alternating null children",
			root:     []*int{IntPtr(1), IntPtr(2), nil, nil, IntPtr(3), IntPtr(4)},
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
			var root *TreeNode
			if tt.name == "Very deep left-skewed tree (10 levels)" {
				// Create a deep left-skewed tree with 10 levels
				// Using the helper from tree_node.go
				root = CreateLeftSkewedTree(10)
			} else {
				root = NewTreeFromSlice(tt.root)
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
	root := CreateCompleteTree(1023)
	
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