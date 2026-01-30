package trees

import (
	"testing"
    "leetcode/utils"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected bool
	}{
		{
			name:     "Example 1: Symmetric tree",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
			expected: true,
		},
		{
			name:     "Example 2: Asymmetric tree",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(3)},
			expected: false,
		},
		{
			name:     "Empty tree",
			root:     []*int{},
			expected: true,
		},
		{
			name:     "Single node",
			root:     []*int{utils.IntPtr(1)},
			expected: true,
		},
		{
			name:     "Two nodes symmetric",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2)},
			expected: true,
		},
		{
			name:     "Two nodes asymmetric",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
			expected: false,
		},
		{
			name:     "Three levels symmetric",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
			expected: true,
		},
		{
			name:     "Three levels asymmetric",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(4)},
			expected: false,
		},
		{
			name:     "Tree with nil values symmetric",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(3), nil},
			expected: true,
		},
		{
			name:     "Tree with nil values asymmetric",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(3)},
			expected: false,
		},
		{
			name:     "Complex symmetric tree",
			// Fixed: Correct 15-node symmetric tree
			// Level 0: [1]
			// Level 1: [2, 2]
			// Level 2: [3, 4, 4, 3]
			// Level 3: [5, 6, 7, 8, 8, 7, 6, 5] - This is symmetric!
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(8), utils.IntPtr(7), utils.IntPtr(6), utils.IntPtr(5)},
			expected: true,
		},
		{
			name:     "Complex asymmetric tree",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(7), utils.IntPtr(6), utils.IntPtr(8)},
			expected: false,
		},
		{
			name:     "All same values symmetric structure",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1)},
			expected: true,
		},
		{
			name:     "All same values asymmetric structure",
			root:     []*int{utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1), nil, utils.IntPtr(1), utils.IntPtr(1)},
			expected: false,
		},
		{
			name:     "Negative values symmetric",
			root:     []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-4), utils.IntPtr(-4), utils.IntPtr(-3)},
			expected: true,
		},
		{
			name:     "Mixed positive negative symmetric",
			root:     []*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(-1), utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(1)},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := utils.NewTreeFromSlice(tt.root)
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
			root: []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
		},
		{
			name: "Asymmetric tree",
			root: []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(3)},
		},
		{
			name: "Empty tree",
			root: []*int{},
		},
		{
			name: "Single node",
			root: []*int{utils.IntPtr(1)},
		},
		{
			name: "Complex symmetric",
			// Fixed: Correct 15-node symmetric tree
			root: []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(8), utils.IntPtr(7), utils.IntPtr(6), utils.IntPtr(5)},
		},
		{
			name: "Complex asymmetric",
			root: []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(7), utils.IntPtr(6), utils.IntPtr(8)},
		},
	}

	implementations := []struct {
		name string
		fn   func(*utils.TreeNode) bool
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
			root := utils.NewTreeFromSlice(tc.root)
			expected := IsSymmetric(root)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create fresh tree for each implementation
					root := utils.NewTreeFromSlice(tc.root)
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
		// Create a large symmetric tree (complete binary tree with symmetric values)
		// Use CreateSymmetricTree instead of createCompleteTree
		// 10 levels gives us 2^10 - 1 = 1023 nodes
		root := utils.CreateSymmetricTree(10) // 2^10 - 1 nodes
		if !IsSymmetric(root) {
			t.Error("Large symmetric tree should be symmetric")
		}
	})

	t.Run("Large asymmetric tree", func(t *testing.T) {
		// Create a right-skewed tree (definitely not symmetric)
		root := utils.CreateRightSkewedTree(100)
		if IsSymmetric(root) {
			t.Error("Right-skewed tree should not be symmetric")
		}
	})

	t.Run("Left-skewed tree", func(t *testing.T) {
		root := utils.CreateLeftSkewedTree(100)
		if IsSymmetric(root) {
			t.Error("Left-skewed tree should not be symmetric")
		}
	})

	t.Run("Tree with only left children", func(t *testing.T) {
		root := &utils.TreeNode{Val: 1}
		root.Left = &utils.TreeNode{Val: 2}
		root.Left.Left = &utils.TreeNode{Val: 3}
		if IsSymmetric(root) {
			t.Error("Tree with only left children should not be symmetric")
		}
	})

	t.Run("Tree with only right children", func(t *testing.T) {
		root := &utils.TreeNode{Val: 1}
		root.Right = &utils.TreeNode{Val: 2}
		root.Right.Right = &utils.TreeNode{Val: 3}
		if IsSymmetric(root) {
			t.Error("Tree with only right children should not be symmetric")
		}
	})

	t.Run("Mirror values but different structure", func(t *testing.T) {
		// Values are mirror but structure isn't
		root := utils.NewTreeFromSlice([]*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(2),
			utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(4),
		})
		if IsSymmetric(root) {
			t.Error("Tree with mirror values but different structure should not be symmetric")
		}
	})
}

func TestIsSymmetricProperties(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*utils.TreeNode) bool
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
			singleNode := utils.NewTreeFromSlice([]*int{utils.IntPtr(1)})
			if !impl.fn(singleNode) {
				t.Error("Single node tree should be symmetric")
			}

			// Property 3: Mirror of a symmetric tree is symmetric
			symmetricTree := utils.NewTreeFromSlice([]*int{
				utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2),
				utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3),
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
		root *utils.TreeNode
	}{
		{
			name: "Small symmetric",
			root: utils.NewTreeFromSlice([]*int{
				utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2),
				utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3),
			}),
		},
		{
			name: "Small asymmetric",
			root: utils.NewTreeFromSlice([]*int{
				utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2),
				nil, utils.IntPtr(3), nil, utils.IntPtr(3),
			}),
		},
		{
			name: "Medium symmetric",
			root: utils.CreateSymmetricTree(9), // 2^9 - 1 = 511 nodes
		},
		{
			name: "Medium asymmetric",
			root: utils.CreateRightSkewedTree(500),
		},
		{
			name: "Large symmetric",
			root: utils.CreateSymmetricTree(12), // 2^12 - 1 = 4095 nodes
		},
		{
			name: "Large asymmetric",
			root: utils.CreateRightSkewedTree(4000),
		},
	}

	implementations := []struct {
		name string
		fn   func(*utils.TreeNode) bool
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
	root := utils.CreateSymmetricTree(14) // 2^14 - 1 = 16383 nodes

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