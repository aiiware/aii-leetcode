package leetcode

import (
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected [][]int
	}{
		{
			name:     "Example 1: Standard tree",
			root:     []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)},
			expected: [][]int{{3}, {20, 9}, {15, 7}},
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
			name:     "Complete binary tree (3 levels)",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name:     "Tree with negative values",
			root:     []*int{IntPtr(-10), IntPtr(9), IntPtr(20), IntPtr(-5), IntPtr(-3), IntPtr(15), IntPtr(7)},
			expected: [][]int{{-10}, {20, 9}, {-5, -3, 15, 7}},
		},
		{
			name:     "Tree with mixed null values",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4), IntPtr(5), nil, nil, nil, IntPtr(6)},
			expected: [][]int{{1}, {3, 2}, {4, 5}, {6}},
		},
		{
			name:     "Large tree (10 nodes)",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7), IntPtr(8), IntPtr(9), IntPtr(10)},
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}, {10, 9, 8}},
		},
		{
			name:     "Tree with zero values",
			root:     []*int{IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0), IntPtr(0)},
			expected: [][]int{{0}, {0, 0}, {0, 0, 0, 0}},
		},
		{
			name:     "Right-skewed tree (linked list to right)",
			root:     []*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Left-skewed tree (linked list to left)",
			root:     []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4)},
			expected: [][]int{{1}, {2}, {3}}, // Fixed: This creates a 3-level tree, not 4
		},
		{
			name:     "Unbalanced tree",
			root:     []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), nil, nil, nil, IntPtr(5)},
			expected: [][]int{{1}, {3, 2}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeFromSlice(tt.root)
			result := zigzagLevelOrder(root)

			if !MatrixEqual(result, tt.expected) {
				t.Errorf("zigzagLevelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAllZigzagLevelOrderImplementations(t *testing.T) {
	testCases := []struct {
		name string
		root []*int
	}{
		{"Empty tree", []*int{}},
		{"Single node", []*int{IntPtr(1)}},
		{"Standard tree", []*int{IntPtr(3), IntPtr(9), IntPtr(20), nil, nil, IntPtr(15), IntPtr(7)}},
		{"Complete tree", []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)}},
		{"Right-skewed tree", []*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3), nil, IntPtr(4)}},
		{"Left-skewed tree", []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3), nil, nil, nil, IntPtr(4)}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := NewTreeFromSlice(tc.root)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*TreeNode) [][]int
			}{
				{"BFS", zigzagLevelOrderBFS},
				{"TwoStacks", zigzagLevelOrderTwoStacks},
				{"DFS", zigzagLevelOrderDFS},
				{"Deque", zigzagLevelOrderDeque},
				{"Optimized", zigzagLevelOrderOptimized},
				{"Main", zigzagLevelOrder},
			}

			// Get reference result from BFS (our main implementation)
			reference := zigzagLevelOrderBFS(root)

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

func TestZigzagLevelOrderWithHelperFunctions(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Left-skewed tree with 3 nodes",
			root:     CreateLeftSkewedTree(3),
			expected: [][]int{{1}, {2}, {3}}, // Fixed: zigzag traversal goes level by level
		},
		{
			name:     "Right-skewed tree with 3 nodes",
			root:     CreateRightSkewedTree(3),
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "Left-skewed tree with 4 nodes",
			root:     CreateLeftSkewedTree(4),
			expected: [][]int{{1}, {2}, {3}, {4}}, // Fixed: zigzag traversal goes level by level
		},
		{
			name:     "Right-skewed tree with 4 nodes",
			root:     CreateRightSkewedTree(4),
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:     "Complete tree with 7 nodes",
			root:     CreateCompleteTree(7),
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name:     "Symmetric tree with 3 levels",
			root:     CreateSymmetricTree(3),
			expected: [][]int{{1}, {2, 2}, {3, 4, 4, 3}}, // Fixed: createSymmetricTree creates symmetric values
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zigzagLevelOrder(tt.root)

			if !MatrixEqual(result, tt.expected) {
				t.Errorf("zigzagLevelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestZigzagLevelOrderEdgeCases(t *testing.T) {
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
			name:     "Tree with single negative value",
			root:     []*int{IntPtr(-1000)},
			expected: [][]int{{-1000}},
		},
		{
			name:     "Tree with single max value",
			root:     []*int{IntPtr(1000)},
			expected: [][]int{{1000}},
		},
		{
			name:     "Tree with alternating null children",
			root:     []*int{IntPtr(1), IntPtr(2), nil, nil, IntPtr(3), IntPtr(4)},
			expected: [][]int{{1}, {2}, {3}, {4}}, // Fixed: This creates a 4-level tree
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			if tt.root != nil {
				root = NewTreeFromSlice(tt.root)
			}
			result := zigzagLevelOrder(root)

			if !MatrixEqual(result, tt.expected) {
				t.Errorf("zigzagLevelOrder() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestZigzagLevelOrderPerformance(t *testing.T) {
	// Create a large complete binary tree (1023 nodes, 10 levels)
	n := 1023
	vals := make([]*int, n)
	for i := 0; i < n; i++ {
		val := i
		vals[i] = &val
	}

	root := NewTreeFromSlice(vals)

	// Quick sanity check
	result := zigzagLevelOrder(root)
	if len(result) != 10 { // 2^10 - 1 = 1023 nodes
		t.Errorf("Expected 10 levels for 1023-node complete tree, got %d", len(result))
	}

	// Verify zigzag pattern
	for level := 0; level < len(result); level++ {
		levelSlice := result[level]
		if level%2 == 1 {
			// Odd levels should be in reverse order compared to regular level order
			// For a complete tree with increasing values, check if decreasing
			for i := 1; i < len(levelSlice); i++ {
				if levelSlice[i] >= levelSlice[i-1] {
					t.Errorf("Level %d (odd) should be in decreasing order, got %v", level, levelSlice)
					break
				}
			}
		} else {
			// Even levels should be in increasing order
			for i := 1; i < len(levelSlice); i++ {
				if levelSlice[i] <= levelSlice[i-1] {
					t.Errorf("Level %d (even) should be in increasing order, got %v", level, levelSlice)
					break
				}
			}
		}
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

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty slice", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Two elements", []int{1, 2}, []int{2, 1}},
		{"Three elements", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Five elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"With negatives", []int{-1, 0, 1}, []int{1, 0, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the test case
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			
			reverseSlice(input)
			
			if !SlicesEqual(input, tt.expected) {
				t.Errorf("reverseSlice(%v) = %v, expected %v", tt.input, input, tt.expected)
			}
		})
	}
}

// Benchmark tests for different implementations
func BenchmarkZigzagLevelOrderBFS(b *testing.B) {
	// Create a medium-sized tree for benchmarking
	vals := make([]*int, 511) // 511 nodes, 9 levels
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zigzagLevelOrderBFS(root)
	}
}

func BenchmarkZigzagLevelOrderTwoStacks(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zigzagLevelOrderTwoStacks(root)
	}
}

func BenchmarkZigzagLevelOrderDFS(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zigzagLevelOrderDFS(root)
	}
}

func BenchmarkZigzagLevelOrderDeque(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zigzagLevelOrderDeque(root)
	}
}

func BenchmarkZigzagLevelOrderOptimized(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zigzagLevelOrderOptimized(root)
	}
}

func BenchmarkZigzagLevelOrderMain(b *testing.B) {
	vals := make([]*int, 511)
	for i := 0; i < 511; i++ {
		val := i
		vals[i] = &val
	}
	root := NewTreeFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zigzagLevelOrder(root)
	}
}