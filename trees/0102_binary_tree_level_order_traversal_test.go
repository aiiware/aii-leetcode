package trees

import (
	"testing"
    "leetcode/utils"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected [][]int
	}{
		{
			name:     "Example 1: Standard tree",
			root:     []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Example 2: Single node",
			root:     []*int{utils.IntPtr(1)},
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3: Empty tree",
			root:     []*int{},
			expected: [][]int{},
		},
		{
			name:     "Complete binary tree",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Tree with negative values",
			root:     []*int{utils.IntPtr(-10), utils.IntPtr(9), utils.IntPtr(20), utils.IntPtr(-5), utils.IntPtr(-3), utils.IntPtr(15), utils.IntPtr(7)},
			expected: [][]int{{-10}, {9, 20}, {-5, -3, 15, 7}},
		},
		{
			name:     "Tree with mixed null values",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6)},
			expected: [][]int{{1}, {2, 3}, {4, 5}, {6}},
		},
		{
			name:     "Large tree (10 nodes)",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10)},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10}},
		},
		{
			name:     "Tree with zero values",
			root:     []*int{utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0), utils.IntPtr(0)},
			expected: [][]int{{0}, {0, 0}, {0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := utils.NewTreeFromSlice(tt.root)
			result := levelOrder(root)

			if !utils.MatrixEqual(result, tt.expected) {
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
		{"Single node", []*int{utils.IntPtr(1)}},
		{"Standard tree", []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}},
		{"Complete tree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := utils.NewTreeFromSlice(tc.root)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*utils.TreeNode) [][]int
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
				if !utils.MatrixEqual(result, reference) {
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
		root     *utils.TreeNode
		expected [][]int
	}{
		{
			name:     "Left-skewed tree with 3 nodes",
			root:     utils.CreateLeftSkewedTree(3),
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "Right-skewed tree with 3 nodes",
			root:     utils.CreateRightSkewedTree(3),
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "Left-skewed tree with 4 nodes",
			root:     utils.CreateLeftSkewedTree(4),
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Right-skewed tree with 4 nodes",
			root:     utils.CreateRightSkewedTree(4),
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Complete tree with 7 nodes",
			root:     utils.CreateCompleteTree(7),
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := levelOrder(tt.root)

			if !utils.MatrixEqual(result, tt.expected) {
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
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil, utils.IntPtr(5)},
			expected: [][]int{{1}, {2, 3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *utils.TreeNode
			if tt.root != nil {
				root = utils.NewTreeFromSlice(tt.root)
			}
			result := levelOrder(root)

			if !utils.MatrixEqual(result, tt.expected) {
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

	root := utils.NewTreeFromSlice(vals)

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
	root := utils.NewTreeFromSlice(vals)

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
	root := utils.NewTreeFromSlice(vals)

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
	root := utils.NewTreeFromSlice(vals)

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
	root := utils.NewTreeFromSlice(vals)

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
	root := utils.NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrderRecursiveWithQueue(root)
	}
}