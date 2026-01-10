package leetcode

import (
	"reflect"
	"testing"
)

func TestLevelOrderBottomBFS(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Example 1: [3,9,20,null,null,15,7]",
			root:     NewTreeFromSlice([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)}),
			expected: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:     "Example 2: [1]",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Single node tree",
			root:     NewTreeFromSlice([]*int{intPtr(5)}),
			expected: [][]int{{5}},
		},
		{
			name:     "Left-skewed tree",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4), nil}),
			expected: [][]int{{4}, {3}, {2}, {1}},
		},
		{
			name:     "Right-skewed tree",
			root:     NewTreeFromSlice([]*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)}),
			expected: [][]int{{4}, {3}, {2}, {1}},
		},
		{
			name:     "Complete binary tree",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)}),
			expected: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
		},
		{
			name:     "Tree with negative values",
			root:     NewTreeFromSlice([]*int{intPtr(-10), intPtr(-5), intPtr(-15), intPtr(-2), intPtr(-7), intPtr(-12), intPtr(-20)}),
			expected: [][]int{{-2, -7, -12, -20}, {-5, -15}, {-10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrderBottomBFS(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrderBottomBFS() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrderBottomDFS(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Example 1: [3,9,20,null,null,15,7]",
			root:     NewTreeFromSlice([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)}),
			expected: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:     "Example 2: [1]",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Complete binary tree",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)}),
			expected: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrderBottomDFS(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrderBottomDFS() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrderBottomOptimized(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Example 1: [3,9,20,null,null,15,7]",
			root:     NewTreeFromSlice([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)}),
			expected: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:     "Example 2: [1]",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Left-skewed tree",
			root:     NewTreeFromSlice([]*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4), nil}),
			expected: [][]int{{4}, {3}, {2}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrderBottomOptimized(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrderBottomOptimized() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrderBottom(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Example 1: [3,9,20,null,null,15,7]",
			root:     NewTreeFromSlice([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)}),
			expected: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:     "Example 2: [1]",
			root:     NewTreeFromSlice([]*int{intPtr(1)}),
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: empty tree",
			root:     nil,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrderBottom(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrderBottom() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestReverseSlice2D(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "Empty slice",
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			name:     "Single element",
			input:    [][]int{{1, 2, 3}},
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "Two elements",
			input:    [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{3, 4}, {1, 2}},
		},
		{
			name:     "Three elements",
			input:    [][]int{{1}, {2, 3}, {4, 5, 6}},
			expected: [][]int{{4, 5, 6}, {2, 3}, {1}},
		},
		{
			name:     "Four elements",
			input:    [][]int{{1}, {2}, {3}, {4}},
			expected: [][]int{{4}, {3}, {2}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the original test case
			inputCopy := make([][]int, len(tt.input))
			copy(inputCopy, tt.input)
			
			reverseSlice2D(inputCopy)
			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("reverseSlice2D() = %v, expected %v", inputCopy, tt.expected)
			}
		})
	}
}

func BenchmarkLevelOrderBottomBFS(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderBottomBFS(root)
	}
}

func BenchmarkLevelOrderBottomDFS(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderBottomDFS(root)
	}
}

func BenchmarkLevelOrderBottomOptimized(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderBottomOptimized(root)
	}
}