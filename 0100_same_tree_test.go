package leetcode

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		p        []*int
		q        []*int
		expected bool
	}{
		{
			name:     "Example 1: Same trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			q:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			expected: true,
		},
		{
			name:     "Example 2: Different structure",
			p:        []*int{IntPtr(1), IntPtr(2)},
			q:        []*int{IntPtr(1), nil, IntPtr(2)},
			expected: false,
		},
		{
			name:     "Example 3: Same structure different values",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(1)},
			q:        []*int{IntPtr(1), IntPtr(1), IntPtr(2)},
			expected: false,
		},
		{
			name:     "Both empty trees",
			p:        []*int{},
			q:        []*int{},
			expected: true,
		},
		{
			name:     "One empty tree",
			p:        []*int{IntPtr(1)},
			q:        []*int{},
			expected: false,
		},
		{
			name:     "Single node same value",
			p:        []*int{IntPtr(5)},
			q:        []*int{IntPtr(5)},
			expected: true,
		},
		{
			name:     "Single node different value",
			p:        []*int{IntPtr(5)},
			q:        []*int{IntPtr(6)},
			expected: false,
		},
		{
			name:     "Complex same trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			q:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: true,
		},
		{
			name:     "Complex different trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			q:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), nil},
			expected: false,
		},
		{
			name:     "Same values different structure 1",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			q:        []*int{IntPtr(1), IntPtr(3), IntPtr(2)},
			expected: false,
		},
		{
			name:     "Same values different structure 2",
			p:        []*int{IntPtr(1), nil, IntPtr(2), nil, IntPtr(3)},
			q:        []*int{IntPtr(1), IntPtr(2), nil, IntPtr(3)},
			expected: false,
		},
		{
			name:     "Trees with negative values",
			p:        []*int{IntPtr(-1), IntPtr(-2), IntPtr(-3)},
			q:        []*int{IntPtr(-1), IntPtr(-2), IntPtr(-3)},
			expected: true,
		},
		{
			name:     "Trees with zero values",
			p:        []*int{IntPtr(0), IntPtr(0), IntPtr(0)},
			q:        []*int{IntPtr(0), IntPtr(0), IntPtr(0)},
			expected: true,
		},
		{
			name:     "Large same trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			q:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			expected: true,
		},
		{
			name:     "Mirror trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			q:        []*int{IntPtr(1), IntPtr(3), IntPtr(2)},
			expected: false,
		},
		{
			name:     "Subtree difference",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), nil, IntPtr(4)},
			q:        []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4)},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewTreeFromSlice(tt.p)
			q := NewTreeFromSlice(tt.q)
			
			result := IsSameTree(p, q)
			if result != tt.expected {
				t.Errorf("IsSameTree() = %v, expected %v", result, tt.expected)
				t.Logf("Tree p: %v", tt.p)
				t.Logf("Tree q: %v", tt.q)
			}
		})
	}
}

func TestAllIsSameTreeImplementations(t *testing.T) {
	testCases := []struct {
		name string
		p    []*int
		q    []*int
	}{
		{
			name: "Same trees",
			p:    []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			q:    []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
		},
		{
			name: "Different structure",
			p:    []*int{IntPtr(1), IntPtr(2)},
			q:    []*int{IntPtr(1), nil, IntPtr(2)},
		},
		{
			name: "Different values",
			p:    []*int{IntPtr(1), IntPtr(2), IntPtr(1)},
			q:    []*int{IntPtr(1), IntPtr(1), IntPtr(2)},
		},
		{
			name: "Both empty",
			p:    []*int{},
			q:    []*int{},
		},
		{
			name: "Complex same",
			p:    []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
			q:    []*int{IntPtr(1), IntPtr(2), IntPtr(3), IntPtr(4), IntPtr(5), IntPtr(6), IntPtr(7)},
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode, *TreeNode) bool
	}{
		{"isSameTree", isSameTree},
		{"isSameTreeIterative", isSameTreeIterative},
		{"isSameTreeBFS", isSameTreeBFS},
		{"isSameTreeDFS", isSameTreeDFS},
		{"isSameTreeSerialization", isSameTreeSerialization},
		{"isSameTreePreorder", isSameTreePreorder},
		{"isSameTreeOptimized", isSameTreeOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewTreeFromSlice(tc.p)
			q := NewTreeFromSlice(tc.q)
			expected := IsSameTree(p, q)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create fresh trees for each implementation
					p := NewTreeFromSlice(tc.p)
					q := NewTreeFromSlice(tc.q)
					
					result := impl.fn(p, q)
					if result != expected {
						t.Errorf("%s() = %v, expected %v",
							impl.name, result, expected)
					}
				})
			}
		})
	}
}

func TestIsSameTreeEdgeCases(t *testing.T) {
	t.Run("Both nil trees", func(t *testing.T) {
		if !IsSameTree(nil, nil) {
			t.Error("Two nil trees should be equal")
		}
	})

	t.Run("One nil tree", func(t *testing.T) {
		p := &TreeNode{Val: 1}
		if IsSameTree(p, nil) {
			t.Error("Non-nil tree should not equal nil tree")
		}
		if IsSameTree(nil, p) {
			t.Error("Nil tree should not equal non-nil tree")
		}
	})

	t.Run("Large trees (1000 nodes)", func(t *testing.T) {
		// Create two identical large trees
		p := CreateCompleteTree(1000)
		q := CloneTree(p)
		
		if !IsSameTree(p, q) {
			t.Error("Large identical trees should be equal")
		}
		
		// Modify one tree
		q.Left.Val = -1
		if IsSameTree(p, q) {
			t.Error("Modified large trees should not be equal")
		}
	})

	t.Run("Trees with all same values but different structure", func(t *testing.T) {
		// Both trees have all nodes with value 1, but different structure
		p := NewTreeFromSlice([]*int{
			IntPtr(1), IntPtr(1), IntPtr(1),
			IntPtr(1), IntPtr(1), IntPtr(1), IntPtr(1),
		})
		q := NewTreeFromSlice([]*int{
			IntPtr(1), IntPtr(1), IntPtr(1),
			nil, IntPtr(1), IntPtr(1), IntPtr(1),
		})
		
		if IsSameTree(p, q) {
			t.Error("Trees with different structure should not be equal even with same values")
		}
	})

	t.Run("Deeply nested trees", func(t *testing.T) {
		// Create deeply nested right-skewed trees
		p := CreateRightSkewedTree(100)
		q := CreateRightSkewedTree(100)
		
		if !IsSameTree(p, q) {
			t.Error("Identical deeply nested trees should be equal")
		}
		
		// Make them different
		q.Right.Right.Val = -1
		if IsSameTree(p, q) {
			t.Error("Different deeply nested trees should not be equal")
		}
	})

	t.Run("Trees with maximum depth", func(t *testing.T) {
		// Test with trees at maximum reasonable depth
		p := CreateRightSkewedTree(1000)
		q := CreateRightSkewedTree(1000)
		
		if !IsSameTree(p, q) {
			t.Error("Identical maximum depth trees should be equal")
		}
	})

	t.Run("Self comparison", func(t *testing.T) {
		p := NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3)})
		if !IsSameTree(p, p) {
			t.Error("Tree should equal itself")
		}
	})
}

func TestIsSameTreeProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(*TreeNode, *TreeNode) bool
	}{
		{"isSameTree", isSameTree},
		{"isSameTreeIterative", isSameTreeIterative},
		{"isSameTreeBFS", isSameTreeBFS},
		{"isSameTreeDFS", isSameTreeDFS},
		{"isSameTreePreorder", isSameTreePreorder},
		{"isSameTreeOptimized", isSameTreeOptimized},
	}

	testCases := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
	}{
		{
			name: "Same tree",
			p:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3)}),
			q:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3)}),
		},
		{
			name: "Different values",
			p:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3)}),
			q:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(4)}),
		},
		{
			name: "Different structure",
			p:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2)}),
			q:    NewTreeFromSlice([]*int{IntPtr(1), nil, IntPtr(2)}),
		},
		{
			name: "Both nil",
			p:    nil,
			q:    nil,
		},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			// Property 1: Reflexivity - tree equals itself
			for _, tc := range testCases {
				if tc.p != nil {
					if !impl.fn(tc.p, tc.p) {
						t.Errorf("Reflexivity failed: tree should equal itself")
					}
				}
			}

			// Property 2: Symmetry - if p equals q, then q equals p
			for _, tc := range testCases {
				result1 := impl.fn(tc.p, tc.q)
				result2 := impl.fn(tc.q, tc.p)
				if result1 != result2 {
					t.Errorf("Symmetry failed: %v != %v", result1, result2)
				}
			}

			// Property 3: Transitivity - if p equals q and q equals r, then p equals r
			// (harder to test without third tree)

			// Property 4: Nil trees
			if !impl.fn(nil, nil) {
				t.Errorf("Two nil trees should be equal")
			}
		})
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	// Create test trees of different sizes
	testCases := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
	}{
		{
			name: "Small trees same",
			p:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3)}),
			q:    NewTreeFromSlice([]*int{IntPtr(1), IntPtr(2), IntPtr(3)}),
		},
		{
			name: "Medium trees same",
			p:    CreateCompleteTree(100),
			q:    CloneTree(CreateCompleteTree(100)),
		},
		{
			name: "Large trees same",
			p:    CreateCompleteTree(1000),
			q:    CloneTree(CreateCompleteTree(1000)),
		},
		{
			name: "Skewed trees same",
			p:    CreateRightSkewedTree(1000),
			q:    CreateRightSkewedTree(1000),
		},
		{
			name: "Trees different at root",
			p:    CreateCompleteTree(100),
			q:    func() *TreeNode {
				t := CloneTree(CreateCompleteTree(100))
				t.Val = -1
				return t
			}(),
		},
	}

	implementations := []struct {
		name string
		fn   func(*TreeNode, *TreeNode) bool
	}{
		{"isSameTree", isSameTree},
		{"isSameTreeIterative", isSameTreeIterative},
		{"isSameTreeBFS", isSameTreeBFS},
		{"isSameTreeDFS", isSameTreeDFS},
		{"isSameTreeSerialization", isSameTreeSerialization},
		{"isSameTreePreorder", isSameTreePreorder},
		{"isSameTreeOptimized", isSameTreeOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.p, tc.q)
					}
				})
			}
		})
	}
}

func BenchmarkIsSameTreeWorstCase(b *testing.B) {
	// Worst case: large identical trees (must traverse all nodes)
	p := CreateCompleteTree(10000)
	q := CloneTree(p)

	b.ResetTimer()

	b.Run("isSameTree", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTree(p, q)
		}
	})

	b.Run("isSameTreeIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTreeIterative(p, q)
		}
	})

	b.Run("isSameTreeBFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTreeBFS(p, q)
		}
	})

	b.Run("isSameTreeOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isSameTreeOptimized(p, q)
		}
	})
}

// Test TreeNode.Equal method
func TestTreeNodeEqualMethod(t *testing.T) {
	tests := []struct {
		name     string
		p        []*int
		q        []*int
		expected bool
	}{
		{
			name:     "Same trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			q:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			expected: true,
		},
		{
			name:     "Different trees",
			p:        []*int{IntPtr(1), IntPtr(2), IntPtr(3)},
			q:        []*int{IntPtr(1), IntPtr(3), IntPtr(2)},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewTreeFromSlice(tt.p)
			q := NewTreeFromSlice(tt.q)
			
			result := p.Equal(q)
			if result != tt.expected {
				t.Errorf("TreeNode.Equal() = %v, expected %v", result, tt.expected)
			}
			
			// Should match IsSameTree
			isSame := IsSameTree(p, q)
			if result != isSame {
				t.Errorf("TreeNode.Equal() = %v, IsSameTree() = %v, should match",
					result, isSame)
			}
		})
	}
}