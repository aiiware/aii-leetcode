package leetcode

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int // -1 means no cycle
		expected *int // nil means no cycle, otherwise pointer to expected value
	}{
		{
			name:     "Empty list",
			values:   []int{},
			pos:      -1,
			expected: nil,
		},
		{
			name:     "Single node, no cycle",
			values:   []int{1},
			pos:      -1,
			expected: nil,
		},
		{
			name:     "Single node with cycle",
			values:   []int{1},
			pos:      0,
			expected: IntPtr(1),
		},
		{
			name:     "Two nodes, no cycle",
			values:   []int{1, 2},
			pos:      -1,
			expected: nil,
		},
		{
			name:     "Two nodes with cycle at head",
			values:   []int{1, 2},
			pos:      0,
			expected: IntPtr(1),
		},
		{
			name:     "Four nodes with cycle at position 1",
			values:   []int{3, 2, 0, -4},
			pos:      1,
			expected: IntPtr(2),
		},
		{
			name:     "Four nodes with cycle at position 2",
			values:   []int{3, 2, 0, -4},
			pos:      2,
			expected: IntPtr(0),
		},
		{
			name:     "Four nodes with cycle at tail",
			values:   []int{3, 2, 0, -4},
			pos:      3,
			expected: IntPtr(-4),
		},
		{
			name:     "Large list with cycle at middle",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:      5,
			expected: IntPtr(6),
		},
		{
			name:     "Large list without cycle",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:      -1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create linked list with optional cycle
			head := createLinkedListWithCycle(tt.values, tt.pos)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*ListNode) *ListNode
			}{
				{"DetectCycle", DetectCycle},
				{"DetectCycleHashTable", DetectCycleHashTable},
				{"DetectCycleOptimized", DetectCycleOptimized},
			}

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(head)

					if tt.expected == nil {
						if result != nil {
							t.Errorf("%s() = node with value %d, expected nil", impl.name, result.Val)
						}
					} else {
						if result == nil {
							t.Errorf("%s() = nil, expected node with value %d", impl.name, *tt.expected)
						} else if result.Val != *tt.expected {
							t.Errorf("%s() = node with value %d, expected %d", impl.name, result.Val, *tt.expected)
						}
					}
				})
			}
		})
	}
}

func BenchmarkDetectCycle(b *testing.B) {
	// Create a large linked list with cycle
	values := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		values[i] = i
	}
	head := createLinkedListWithCycle(values, 5000)

	b.Run("FloydAlgorithm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DetectCycle(head)
		}
	})

	b.Run("HashTable", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DetectCycleHashTable(head)
		}
	})

	b.Run("OptimizedFloyd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DetectCycleOptimized(head)
		}
	})
}

func TestDetectCycleEdgeCases(t *testing.T) {
	t.Run("Cycle at different positions", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}

		// Test cycle at each position
		for pos := 0; pos < len(values); pos++ {
			head := createLinkedListWithCycle(values, pos)
			result := DetectCycle(head)

			if result == nil {
				t.Errorf("Expected cycle at position %d, but got nil", pos)
			} else if result.Val != values[pos] {
				t.Errorf("Expected node with value %d at position %d, but got %d", values[pos], pos, result.Val)
			}
		}
	})

	t.Run("No cycle in various lists", func(t *testing.T) {
		testCases := [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 3},
			{1, 2, 3, 4, 5},
		}

		for _, values := range testCases {
			head := createLinkedListWithCycle(values, -1)
			result := DetectCycle(head)

			if result != nil {
				t.Errorf("Expected no cycle for list %v, but got node with value %d", values, result.Val)
			}
		}
	})
}