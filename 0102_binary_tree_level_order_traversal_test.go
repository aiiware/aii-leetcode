package leetcode

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected [][]int
	}{
		{
			name:     "Example 1: Standard tree",
			root:     []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Example 2: Single node",
			root:     []*int{IntPtr(1)},
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: Empty tree",
			root:     []*int{},
			expected: [][]int{},
		},
		{
			name:     "Complete binary tree",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Tree with negative values",
			root:     []*int{IntPtr(-10), IntPtr(9), IntPtr(20), IntPtr(-5), IntPtr(-3), IntPtr(15), IntPtr(7)},
			expected: [][]int{{-10}, {9, 20}, {-5, -3, 15, 7}},
		},
		{
			name:     "Tree with mixed null values",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4), IntPtr(5), nil, nil, nil, IntPtr(6)},
			expected: [][]int{{1}, {2, 3}, {4, 5}, {6}},
		},
		{
			name:     "Large tree (10 nodes)",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7), IntPtr(8), IntPtr(9), IntPtr(10)},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10}},
		},
		{
			name:     "Tree with zero values",
			root:     []*int{IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0)},
			expected: [][]int{{0}, {0, 0}, {0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeFromSlice(tt.root)
			result := levelOrder(root)

			if !MatrixEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAllLevelOrderImplementations(t *testing.T) {
	testCases := []struct {
		name string
		root []*int
	}{
		{"Empty tree", []*int{}},
		{"Single node", []*int{IntPtr(1)}},
		{"Standard tree", []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)}},
		{"Complete tree", []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := NewTreeFromSlice(tc.root)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*TreeNode) [][]int
			}{
				{"BFS", levelOrderBFS},
				{"DFS", levelOrderDFS},
				{"TwoQueues", levelOrderTwoQueues},
				{"RecursiveWithQueue", levelOrderRecursiveWithQueue},
				{"Optimized", levelOrderOptimized},
				{"Main", levelOrder},
			}

			// Get reference result from BFS (our main implementation)
			reference := levelOrderBFS(root)

			for _, impl := range implementations {
				result := impl.fn(root)
				if !MatrixEqual(result, reference) {
					t.Errorf("%s implementation failed for %s: got %v, expected %v",
						impl.name, tc.name, result, reference)
				}
			}
		})
	}
}

func TestLevelOrderWithHelperFunctions(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Left-skewed tree with 3 nodes",
			root:     createLeftSkewedTreeTest(3),
			expected: [][]int{{3}, {2}, {1}},
		},
		{
			name:     "Right-skewed tree with 3 nodes",
			root:     createRightSkewedTreeTest(3),
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "Left-skewed tree with 4 nodes",
			root:     createLeftSkewedTreeTest(4),
			expected: [][]int{{4}, {3}, {2}, {1}},
		},
		{
			name:     "Right-skewed tree with 4 nodes",
			root:     createRightSkewedTreeTest(4),
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Complete tree with 7 nodes",
			root:     createCompleteTree(7),
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrder(tt.root)

			if !MatrixEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrderEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected [][]int
	}{
		{
			name:     "Nil root passed directly",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Unbalanced tree",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), nil, nil, nil, IntPtr(5)},
			expected: [][]int{{1}, {2, 3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			if tt.root != nil {
				root = NewTreeFromSlice(tt.root)
			}
			result := levelOrder(root)

			if !MatrixEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestLevelOrderPerformance(t *testing.T) {
	// Create a large complete binary tree (1023 nodes, 10 levels)
	n := 1023
	vals := make([]*int, n)
	for i := 0; i < n; i++ {
		val := i
		vals[i] = &val
	}

	root := NewTreeFromSlice(vals)

	// Quick sanity check
	result := levelOrder(root)
	if len(result) != 10 { // 2^10 - 1 = 1023 nodes
		t.Errorf("Expected 10 levels for 1023-node complete tree, got %d", len(result))
	}

	// Verify total nodes
	totalNodes := 0
	for _, level := range result {
		totalNodes += len(level)
	}
	if totalNodes != n {
		t.Errorf("Expected %d total nodes, got %d", n, totalNodes)
	}
}

// Benchmark tests for different implementations
func BenchmarkLevelOrderBFS(b *testing.B) {
	// Create a medium-sized tree for benchmarking
	vals := make([]*int, 511) // 511 nodes, 9 levels
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderBFS(root)
	}
}

func BenchmarkLevelOrderDFS(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderDFS(root)
	}
}

func BenchmarkLevelOrderTwoQueues(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderTwoQueues(root)
	}
}

func BenchmarkLevelOrderOptimized(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderOptimized(root)
	}
}

func BenchmarkLevelOrderRecursiveWithQueue(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderRecursiveWithQueue(root)
	}
}