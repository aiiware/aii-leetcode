package linkedlists

import (
	"testing"
    "leetcode/utils"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		pos      int // -1 means no cycle
		expected bool
	}{
		{
			name:     "Empty list",
			values:   []int{},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Single node, no cycle",
			values:   []int{1},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Single node with cycle",
			values:   []int{1},
			pos:      0,
			expected: true,
		},
		{
			name:     "Two nodes, no cycle",
			values:   []int{1, 2},
			pos:      -1,
			expected: false,
		},
		{
			name:     "Two nodes with cycle at head",
			values:   []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			name:     "Four nodes with cycle at position 1",
			values:   []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			name:     "Four nodes with cycle at tail",
			values:   []int{3, 2, 0, -4},
			pos:      3,
			expected: true,
		},
		{
			name:     "Large list with cycle",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:      5,
			expected: true,
		},
		{
			name:     "Large list without cycle",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:      -1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create linked list with optional cycle
			head := createLinkedListWithCycle(tt.values, tt.pos)

			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*utils.ListNode) bool
			}{
				{"HasCycle", HasCycle},
				{"HasCycleHashTable", HasCycleHashTable},
				{"HasCycleOptimized", HasCycleOptimized},
			}

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(head)
					if result != tt.expected {
						t.Errorf("%s() = %v, expected %v", impl.name, result, tt.expected)
					}
				})
			}
		})
	}
}

func BenchmarkHasCycle(b *testing.B) {
	// Create a large linked list with cycle
	values := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		values[i] = i
	}
	head := createLinkedListWithCycle(values, 5000)

	b.Run("FloydAlgorithm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HasCycle(head)
		}
	})

	b.Run("HashTable", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HasCycleHashTable(head)
		}
	})

	b.Run("OptimizedFloyd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HasCycleOptimized(head)
		}
	})
}

// Helper function to create a linked list with optional cycle
func createLinkedListWithCycle(values []int, pos int) *utils.ListNode {
	if len(values) == 0 {
		return nil
	}

	// Create all nodes
	nodes := make([]*utils.ListNode, len(values))
	for i, val := range values {
		nodes[i] = &utils.ListNode{Val: val}
	}

	// Connect nodes
	for i := 0; i < len(values)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Create cycle if pos is valid
	if pos >= 0 && pos < len(values) {
		nodes[len(values)-1].Next = nodes[pos]
	}

	return nodes[0]
}

func TestCreateLinkedListWithCycle(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		pos    int
	}{
		{
			name:   "No cycle",
			values: []int{1, 2, 3},
			pos:    -1,
		},
		{
			name:   "Cycle at head",
			values: []int{1, 2, 3},
			pos:    0,
		},
		{
			name:   "Cycle at middle",
			values: []int{1, 2, 3, 4, 5},
			pos:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedListWithCycle(tt.values, tt.pos)

			// Verify the list structure
			if tt.pos == -1 {
				// No cycle - should reach nil
				current := head
				count := 0
				for current != nil && count <= len(tt.values) {
					current = current.Next
					count++
				}
				if count != len(tt.values) {
					t.Errorf("Expected to traverse %d nodes, but traversed %d", len(tt.values), count)
				}
			} else {
				// Has cycle - should detect cycle
				if !HasCycle(head) {
					t.Errorf("Expected cycle at position %d, but no cycle detected", tt.pos)
				}
			}
		})
	}
}