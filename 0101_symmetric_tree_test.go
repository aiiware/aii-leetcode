package leetcode

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected bool
	}{
		{
			name:     "Example 1: Symmetric tree",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3)},
			expected: true,
		},
		{
			name:     "Example 2: Asymmetric tree",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), nil, intPtr(3), nil, intPtr(3)},
			expected: false,
		},
		{
			name:     "Empty tree",
			root:     []*int{},
			expected: true,
		},
		{
			name:     "Single node",
			root:     []*int{intPtr(1)},
			expected: true,
		},
		{
			name:     "Two nodes symmetric",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2)},
			expected: true,
		},
		{
			name:     "Two nodes asymmetric",
			root:     []*int{intPtr(1), intPtr(2), intPtr(3)},
			expected: false,
		},
		{
			name:     "Three levels symmetric",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3)},
			expected: true,
		},
		{
			name:     "Three levels asymmetric",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(3), intPtr(4)},
			expected: false,
		},
		{
			name:     "Tree with nil values symmetric",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), nil, intPtr(3), intPtr(3), nil},
			expected: true,
		},
		{
			name:     "Tree with nil values asymmetric",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), nil, intPtr(3), nil, intPtr(3)},
			expected: false,
		},
		{
			name:     "Complex symmetric tree",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3), intPtr(5), intPtr(6), intPtr(7), intPtr(7), intPtr(6), intPtr(5)},
			expected: true,
		},
		{
			name:     "Complex asymmetric tree",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3), intPtr(5), intPtr(6), intPtr(7), intPtr(7), intPtr(6), intPtr(8)},
			expected: false,
		},
		{
			name:     "All same values symmetric structure",
			root:     []*int{intPtr(1), intPtr(1), intPtr(1), intPtr(1), intPtr(1), intPtr(1), intPtr(1)},
			expected: true,
		},
		{
			name:     "All same values asymmetric structure",
			root:     []*int{intPtr(1), intPtr(1), intPtr(1), intPtr(1), nil, intPtr(1), intPtr(1)},
			expected: false,
		},
		{
			name:     "Negative values symmetric",
			root:     []*int{intPtr(-1), intPtr(-2), intPtr(-2), intPtr(-3), intPtr(-4), intPtr(-4), intPtr(-3)},
			expected: true,
		},
		{
			name:     "Mixed positive negative symmetric",
			root:     []*int{intPtr(0), intPtr(-1), intPtr(-1), intPtr(1), intPtr(2), intPtr(2), intPtr(1)},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewTreeFromSlice(tt.root)
			result := IsSymmetric(root)
			if result != tt.expected {
				t.Errorf("IsSymmetric() = %v, expected %v", result, tt.expected)
				t.Logf("Tree: %v", tt.root)
			}
		})
	}
}

func TestAllIsSymmetricImplementations(t *testing.T) {
	testCases := []struct {
		name string
		root []*int
	}{
		{
			name: "Symmetric tree",
			root: []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3)},
		},
		{
			name: "Asymmetric tree",
			root: []*int{intPtr(1), intPtr(2), intPtr(2), nil, intPtr(3), nil, intPtr(3)},
		},
		{
			name: "Empty tree",
			root: []*int{},
		},
		{
			name: "Single node",
			root: []*int{intPtr(1)},
		},
		{
			name: "Complex symmetric",
			root: []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3), intPtr(5), intPtr(6), intPtr(7), intPtr(7), intPtr(6), intPtr(5)},
		},
		{
			name: "Complex asymmetric",
			root: []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(4), intPtr(4), intPtr(3), intPtr(5), intPtr(6), intPtr(7), intPtr(7), intPtr(6), intPtr(8)},
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{"isSymmetricRecursive", isSymmetricRecursive},
		{"isSymmetricIterative", isSymmetricIterative},
		{"isSymmetricStack", isSymmetricStack},
		{"isSymmetricLevelOrder", isSymmetricLevelOrder},
		{"isSymmetricDFS", isSymmetricDFS},
		{"isSymmetricOptimized", isSymmetricOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			root := NewTreeFromSlice(tc.root)
			expected := IsSymmetric(root)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create fresh tree for each implementation
					root := NewTreeFromSlice(tc.root)
					result := impl.fn(root)
					if result != expected {
						t.Errorf("%s() = %v, expected %v", impl.name, result, expected)
					}
				})
			}
		})
	}
}

func TestIsSymmetricEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		if !IsSymmetric(nil) {
			t.Error("Nil tree should be symmetric")
		}
	})

	t.Run("Large symmetric tree", func(t *testing.T) {
		// Create a large symmetric tree (complete binary tree)
		root := createCompleteTree(1023) // 2^10 - 1 nodes
		if !IsSymmetric(root) {
			t.Error("Large complete tree should be symmetric")
		}
	})

	t.Run("Large asymmetric tree", func(t *testing.T) {
		// Create a right-skewed tree (definitely not symmetric)
		root := createRightSkewedTree(100)
		if IsSymmetric(root) {
			t.Error("Right-skewed tree should not be symmetric")
		}
	})

	t.Run("Left-skewed tree", func(t *testing.T) {
		root := createLeftSkewedTree(100)
		if IsSymmetric(root) {
			t.Error("Left-skewed tree should not be symmetric")
		}
	})

	t.Run("Tree with only left children", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Left.Left = &TreeNode{Val: 3}
		if IsSymmetric(root) {
			t.Error("Tree with only left children should not be symmetric")
		}
	})

	t.Run("Tree with only right children", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Right = &TreeNode{Val: 3}
		if IsSymmetric(root) {
			t.Error("Tree with only right children should not be symmetric")
		}
	})

	t.Run("Mirror values but different structure", func(t *testing.T) {
		// Values are mirror but structure isn't
		root := NewTreeFromSlice([]*int{
			intPtr(1),
			intPtr(2), intPtr(2),
			intPtr(3), intPtr(4), intPtr(3), intPtr(4),
		})
		if IsSymmetric(root) {
			t.Error("Tree with mirror values but different structure should not be symmetric")
		}
	})
}

func TestIsSymmetricProperties(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{"isSymmetricRecursive", isSymmetricRecursive},
		{"isSymmetricIterative", isSymmetricIterative},
		{"isSymmetricOptimized", isSymmetricOptimized},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			// Property 1: Nil tree is symmetric
			if !impl.fn(nil) {
				t.Error("Nil tree should be symmetric")
			}

			// Property 2: Single node tree is symmetric
			singleNode := NewTreeFromSlice([]*int{intPtr(1)})
			if !impl.fn(singleNode) {
				t.Error("Single node tree should be symmetric")
			}

			// Property 3: Mirror of a symmetric tree is symmetric
			symmetricTree := NewTreeFromSlice([]*int{
				intPtr(1), intPtr(2), intPtr(2),
				intPtr(3), intPtr(4), intPtr(4), intPtr(3),
			})
			if !impl.fn(symmetricTree) {
				t.Error("Symmetric tree should be symmetric")
			}

			// Property 4: Tree and its mirror should both give same result
			// (trivially true since mirror of asymmetric tree is still asymmetric)
		})
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	// Create test trees of different sizes and types
	testCases := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Small symmetric",
			root: NewTreeFromSlice([]*int{
				intPtr(1), intPtr(2), intPtr(2),
				intPtr(3), intPtr(4), intPtr(4), intPtr(3),
			}),
		},
		{
			name: "Small asymmetric",
			root: NewTreeFromSlice([]*int{
				intPtr(1), intPtr(2), intPtr(2),
				nil, intPtr(3), nil, intPtr(3),
			}),
		},
		{
			name: "Medium symmetric",
			root: createCompleteTree(511), // 2^9 - 1 nodes
		},
		{
			name: "Medium asymmetric",
			root: createRightSkewedTree(500),
		},
		{
			name: "Large symmetric",
			root: createCompleteTree(4095), // 2^12 - 1 nodes
		},
		{
			name: "Large asymmetric",
			root: createRightSkewedTree(4000),
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{"isSymmetricRecursive", isSymmetricRecursive},
		{"isSymmetricIterative", isSymmetricIterative},
		{"isSymmetricStack", isSymmetricStack},
		{"isSymmetricLevelOrder", isSymmetricLevelOrder},
		{"isSymmetricDFS", isSymmetricDFS},
		{"isSymmetricOptimized", isSymmetricOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.root)
					}
				})
			}
		})
	}
}

func BenchmarkIsSymmetricWorstCase(b *testing.B) {
	// Worst case: large symmetric tree (must check all nodes)
	root := createCompleteTree(16383) // 2^14 - 1 nodes

	b.ResetTimer()

	b.Run("isSymmetricRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSymmetricRecursive(root)
		}
	})

	b.Run("isSymmetricIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSymmetricIterative(root)
		}
	})

	b.Run("isSymmetricOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSymmetricOptimized(root)
		}
	})
}