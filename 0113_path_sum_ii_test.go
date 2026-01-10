package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		target   int
		expected [][]int
	}{
		{
			name:   "Example 1: Multiple valid paths",
			root:   NewTreeFromSlice([]*int{intPtr(5), intPtr(4), intPtr(8), intPtr(11), nil, intPtr(13), intPtr(4), intPtr(7), intPtr(2), nil, nil, intPtr(5), intPtr(1)}),
			target: 22,
			expected: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
		{
			name:     "Example 2: No valid paths",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(2), intPtr(3)}),
			target:   5,
			expected: [][]int{},
		},
		{
			name:     "Example 3: Single node tree with matching value",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			target:   1,
			expected: [][]int{{1}},
		},
		{
			name:     "Example 4: Single node tree with non-matching value",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			target:   2,
			expected: [][]int{},
		},
		{
			name:     "Example 5: Empty tree",
			root:     nil,
			target:   0,
			expected: [][]int{},
		},
		{
			name:   "Example 6: Complex tree with multiple paths",
			root:   NewTreeFromSlice([]*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)}),
			target: 8,
			expected: [][]int{
				{1, 2, 5},
			},
		},
		{
			name:   "Example 7: Negative values",
			root:   NewTreeFromSlice([]*int{intPtr(-2), nil, intPtr(-3)}),
			target: -5,
			expected: [][]int{
				{-2, -3},
			},
		},
		{
			name:   "Example 8: Zero values",
			root:   NewTreeFromSlice([]*int{intPtr(0), intPtr(1), intPtr(0), intPtr(1), intPtr(0), intPtr(1)}),
			target: 2,
			expected: [][]int{
				{0, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test recursive solution
			result := pathSum(tt.root, tt.target)
			if !equalPathSlices(result, tt.expected) {
				t.Errorf("pathSum() = %v, expected %v", result, tt.expected)
			}

			// Test iterative solution
			resultIterative := pathSumIterative(tt.root, tt.target)
			if !equalPathSlices(resultIterative, tt.expected) {
				t.Errorf("pathSumIterative() = %v, expected %v", resultIterative, tt.expected)
			}
		})
	}
}

func TestPathSum_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		target   int
		expected [][]int
	}{
		{
			name:     "Large target sum",
			root:     NewTreeFromSlice([]*int{intPtr(1000), intPtr(1000), intPtr(1000)}),
			target:   2000,
			expected: [][]int{{1000, 1000}, {1000, 1000}},
		},
		{
			name:     "All paths valid in complete binary tree",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(1), intPtr(1), intPtr(1), intPtr(1), intPtr(1), intPtr(1)}),
			target:   3,
			expected: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		},
		{
			name:     "Skewed tree with single path",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4)}),
			target:   10,
			expected: [][]int{{1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pathSum(tt.root, tt.target)
			if !equalPathSlices(result, tt.expected) {
				t.Errorf("pathSum() = %v, expected %v", result, tt.expected)
			}

			resultIterative := pathSumIterative(tt.root, tt.target)
			if !equalPathSlices(resultIterative, tt.expected) {
				t.Errorf("pathSumIterative() = %v, expected %v", resultIterative, tt.expected)
			}
		})
	}
}

func TestPathSum_Performance(t *testing.T) {
	// Create a balanced tree with 15 nodes (height 4)
	// All values are 1, target is 4 (paths of length 4)
	vals := make([]*int, 15)
	for i := range vals {
		vals[i] = intPtr(1)
	}
	root := NewTreeFromSlice(vals)
	target := 4

	// There should be 8 paths of length 4 in a complete binary tree with 15 nodes
	// (all leaf nodes at depth 4)
	expectedCount := 8

	result := pathSum(root, target)
	if len(result) != expectedCount {
		t.Errorf("pathSum() returned %d paths, expected %d", len(result), expectedCount)
	}

	// Verify all paths have sum 4
	for _, path := range result {
		sum := 0
		for _, val := range path {
			sum += val
		}
		if sum != target {
			t.Errorf("path %v has sum %d, expected %d", path, sum, target)
		}
		if len(path) != 4 {
			t.Errorf("path %v has length %d, expected 4", path, len(path))
		}
	}
}

// Helper function to compare path slices ignoring order
func equalPathSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	
	// Sort both slices
	sort.Slice(a, func(i, j int) bool {
		return comparePaths(a[i], a[j]) < 0
	})
	sort.Slice(b, func(i, j int) bool {
		return comparePaths(b[i], b[j]) < 0
	})
	
	return reflect.DeepEqual(a, b)
}

// Helper function to compare two paths lexicographically
func comparePaths(a, b []int) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return len(a) - len(b)
}