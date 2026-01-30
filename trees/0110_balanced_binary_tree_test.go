package trees

import (
	"testing"
    "leetcode/utils"
)

func TestIsBalancedTopDown(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected bool
	}{
		{
			name:     "Example 1: Balanced tree [3,9,20,null,null,15,7]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Example 2: Unbalanced tree [1,2,2,3,3,null,null,4,4]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Left-skewed tree (unbalanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil}),
			expected: false,
		},
		{
			name:     "Right-skewed tree (unbalanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Complete binary tree (balanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Almost balanced tree (difference of 1)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil}),
			expected: true,
		},
		{
			name:     "Unbalanced at root",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5), nil}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalancedTopDown(tt.root)
			if result != tt.expected {
				t.Errorf("isBalancedTopDown() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsBalancedBottomUp(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected bool
	}{
		{
			name:     "Example 1: Balanced tree [3,9,20,null,null,15,7]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Example 2: Unbalanced tree [1,2,2,3,3,null,null,4,4]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Left-skewed tree (unbalanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil}),
			expected: false,
		},
		{
			name:     "Right-skewed tree (unbalanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Complete binary tree (balanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Almost balanced tree (difference of 1)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil}),
			expected: true,
		},
		{
			name:     "Unbalanced at root",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5), nil}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalancedBottomUp(tt.root)
			if result != tt.expected {
				t.Errorf("isBalancedBottomUp() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsBalancedIterative(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected bool
	}{
		{
			name:     "Example 1: Balanced tree [3,9,20,null,null,15,7]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Example 2: Unbalanced tree [1,2,2,3,3,null,null,4,4]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: true,
		},
		{
			name:     "Left-skewed tree (unbalanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil}),
			expected: false,
		},
		{
			name:     "Right-skewed tree (unbalanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Complete binary tree (balanced)",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalancedIterative(tt.root)
			if result != tt.expected {
				t.Errorf("isBalancedIterative() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected bool
	}{
		{
			name:     "Example 1: Balanced tree [3,9,20,null,null,15,7]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}),
			expected: true,
		},
		{
			name:     "Example 2: Unbalanced tree [1,2,2,3,3,null,null,4,4]",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(4)}),
			expected: false,
		},
		{
			name:     "Example 3: Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isBalanced(tt.root)
			if result != tt.expected {
				t.Errorf("isBalanced() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	tests := []struct {
		name     string
		root     *utils.TreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expected: 1,
		},
		{
			name:     "Tree with left child only",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil}),
			expected: 2,
		},
		{
			name:     "Tree with right child only",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2)}),
			expected: 2,
		},
		{
			name:     "Balanced tree height 3",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
			expected: 3,
		},
		{
			name:     "Left-skewed tree height 4",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil}),
			expected: 4,
		},
		{
			name:     "Right-skewed tree height 4",
			root:     utils.NewTreeFromSlice([]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := height(tt.root)
			if result != tt.expected {
				t.Errorf("height() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCheckBalance(t *testing.T) {
	tests := []struct {
		name          string
		root          *utils.TreeNode
		expectedHeight int
		expectedBalanced bool
	}{
		{
			name:          "Empty tree",
			root:          nil,
			expectedHeight: 0,
			expectedBalanced: true,
		},
		{
			name:          "Single node",
			root:          utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
			expectedHeight: 1,
			expectedBalanced: true,
		},
		{
			name:          "Balanced tree",
			root:          utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}),
			expectedHeight: 2,
			expectedBalanced: true,
		},
		{
			name:          "Unbalanced tree",
			root:          utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil}),
			expectedHeight: 0, // Returns 0 when unbalanced
			expectedBalanced: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			height, balanced := checkBalance(tt.root)
			if height != tt.expectedHeight {
				t.Errorf("checkBalance() height = %v, expected %v", height, tt.expectedHeight)
			}
			if balanced != tt.expectedBalanced {
				t.Errorf("checkBalance() balanced = %v, expected %v", balanced, tt.expectedBalanced)
			}
		})
	}
}

func TestAllImplementationsAgree(t *testing.T) {
	testCases := []struct {
		name string
		root *utils.TreeNode
	}{
		{
			name: "Empty tree",
			root: nil,
		},
		{
			name: "Single node",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1)}),
		},
		{
			name: "Balanced tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}),
		},
		{
			name: "Unbalanced tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil}),
		},
		{
			name: "Complex balanced tree",
			root: utils.NewTreeFromSlice([]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			topDown := isBalancedTopDown(tc.root)
			bottomUp := isBalancedBottomUp(tc.root)
			iterative := isBalancedIterative(tc.root)
			main := isBalanced(tc.root)

			if !(topDown == bottomUp && bottomUp == iterative && iterative == main) {
				t.Errorf("Implementations disagree: TopDown=%v, BottomUp=%v, Iterative=%v, Main=%v",
					topDown, bottomUp, iterative, main)
			}
		})
	}
}

func BenchmarkIsBalancedTopDown(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := utils.CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isBalancedTopDown(root)
	}
}

func BenchmarkIsBalancedBottomUp(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := utils.CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isBalancedBottomUp(root)
	}
}

func BenchmarkIsBalancedIterative(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := utils.CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isBalancedIterative(root)
	}
}

func BenchmarkIsBalanced(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	root := utils.CreateCompleteTree(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isBalanced(root)
	}
}