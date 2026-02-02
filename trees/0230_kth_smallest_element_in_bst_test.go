package trees

import (
	"testing"

	"leetcode/utils"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(1), utils.IntPtr(4), nil, utils.IntPtr(2)}),
			k:        1,
			expected: 1,
		},
		{
			name:     "Example 2",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(6), utils.IntPtr(2), utils.IntPtr(4), nil, nil, utils.IntPtr(1)}),
			k:        3,
			expected: 3,
		},
		{
			name:     "Single node tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			k:        1,
			expected: 1,
		},
		{
			name:     "Left-skewed tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(5), utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}),
			k:        2,
			expected: 2,
		},
		{
			name:     "Right-skewed tree",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			k:        4,
			expected: 4,
		},
		{
			name:     "Complete BST",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(7)}),
			k:        5,
			expected: 5,
		},
		{
			name:     "k equals number of nodes",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3)}),
			k:        5,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test iterative solution
			result := KthSmallest(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("KthSmallest() = %v, expected %v", result, tt.expected)
			}

			// Test recursive solution
			result = KthSmallestRecursive(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("KthSmallestRecursive() = %v, expected %v", result, tt.expected)
			}

			// Test Morris traversal solution
			result = KthSmallestMorris(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("KthSmallestMorris() = %v, expected %v", result, tt.expected)
			}

			// Test counter-based solution
			result = KthSmallestWithCounter(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("KthSmallestWithCounter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKthSmallestEdgeCases(t *testing.T) {
	// Test with nil root (should handle gracefully)
	if result := KthSmallest(nil, 1); result != -1 {
		t.Errorf("KthSmallest(nil, 1) = %v, expected -1", result)
	}

	// Test with k larger than tree size
	root := utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)})
	if result := KthSmallest(root, 10); result != -1 {
		t.Errorf("KthSmallest(root, 10) = %v, expected -1", result)
	}
}

func BenchmarkKthSmallest(b *testing.B) {
	// Create a large BST for benchmarking
	root := utils.CreateCompleteTree(10000)

	b.Run("Iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KthSmallest(root, 5000)
		}
	})

	b.Run("Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KthSmallestRecursive(root, 5000)
		}
	})

	b.Run("Morris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KthSmallestMorris(root, 5000)
		}
	})

	b.Run("WithCounter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			KthSmallestWithCounter(root, 5000)
		}
	})
}