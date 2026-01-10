package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBSTRecursive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(*TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			nums: []int{-10, -3, 0, 5, 9},
			validate: func(root *TreeNode) bool {
				// Check if it's a valid BST
				if !isValidBST(root) {
					return false
				}
				// Check if it's height-balanced
				if !isBalancedBottomUp(root) {
					return false
				}
				// Check if inorder traversal gives sorted array
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{-10, -3, 0, 5, 9})
			},
		},
		{
			name: "Example 2: [1,3]",
			nums: []int{1, 3},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 3})
			},
		},
		{
			name: "Empty array",
			nums: []int{},
			validate: func(root *TreeNode) bool {
				return root == nil
			},
		},
		{
			name: "Single element",
			nums: []int{5},
			validate: func(root *TreeNode) bool {
				if root == nil {
					return false
				}
				if root.Val != 5 {
					return false
				}
				if root.Left != nil || root.Right != nil {
					return false
				}
				return true
			},
		},
		{
			name: "Two elements",
			nums: []int{1, 2},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 2})
			},
		},
		{
			name: "Three elements",
			nums: []int{1, 2, 3},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 2, 3})
			},
		},
		{
			name: "Four elements",
			nums: []int{1, 2, 3, 4},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 2, 3, 4})
			},
		},
		{
			name: "Negative values",
			nums: []int{-5, -4, -3, -2, -1},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{-5, -4, -3, -2, -1})
			},
		},
		{
			name: "Mixed positive and negative",
			nums: []int{-10, -5, 0, 5, 10},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{-10, -5, 0, 5, 10})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBSTRecursive(tt.nums)
			if !tt.validate(root) {
				t.Errorf("sortedArrayToBSTRecursive() failed validation for nums = %v", tt.nums)
			}
		})
	}
}

func TestSortedArrayToBSTIterative(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(*TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			nums: []int{-10, -3, 0, 5, 9},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{-10, -3, 0, 5, 9})
			},
		},
		{
			name: "Example 2: [1,3]",
			nums: []int{1, 3},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 3})
			},
		},
		{
			name: "Empty array",
			nums: []int{},
			validate: func(root *TreeNode) bool {
				return root == nil
			},
		},
		{
			name: "Single element",
			nums: []int{5},
			validate: func(root *TreeNode) bool {
				if root == nil {
					return false
				}
				if root.Val != 5 {
					return false
				}
				if root.Left != nil || root.Right != nil {
					return false
				}
				return true
			},
		},
		{
			name: "Three elements",
			nums: []int{1, 2, 3},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 2, 3})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBSTIterative(tt.nums)
			if !tt.validate(root) {
				t.Errorf("sortedArrayToBSTIterative() failed validation for nums = %v", tt.nums)
			}
		})
	}
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(*TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			nums: []int{-10, -3, 0, 5, 9},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{-10, -3, 0, 5, 9})
			},
		},
		{
			name: "Example 2: [1,3]",
			nums: []int{1, 3},
			validate: func(root *TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 3})
			},
		},
		{
			name: "Empty array",
			nums: []int{},
			validate: func(root *TreeNode) bool {
				return root == nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)
			if !tt.validate(root) {
				t.Errorf("sortedArrayToBST() failed validation for nums = %v", tt.nums)
			}
		})
	}
}

func TestRecursiveAndIterativeProduceSameStructure(t *testing.T) {
	testCases := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{-5, -4, -3, -2, -1},
		{-10, -5, 0, 5, 10},
	}

	for _, nums := range testCases {
		t.Run("", func(t *testing.T) {
			recursiveRoot := sortedArrayToBSTRecursive(nums)
			iterativeRoot := sortedArrayToBSTIterative(nums)

			// Compare inorder traversals
			recursiveInorder := inorderTraversal(recursiveRoot)
			iterativeInorder := inorderTraversal(iterativeRoot)

			if !reflect.DeepEqual(recursiveInorder, iterativeInorder) {
				t.Errorf("For nums = %v: recursive inorder = %v, iterative inorder = %v",
					nums, recursiveInorder, iterativeInorder)
			}

			// Both should be valid BSTs
			if !isValidBST(recursiveRoot) {
				t.Errorf("Recursive result not a valid BST for nums = %v", nums)
			}
			if !isValidBST(iterativeRoot) {
				t.Errorf("Iterative result not a valid BST for nums = %v", nums)
			}

			// Both should be balanced
			if !isBalancedBottomUp(recursiveRoot) {
				t.Errorf("Recursive result not balanced for nums = %v", nums)
			}
			if !isBalancedBottomUp(iterativeRoot) {
				t.Errorf("Iterative result not balanced for nums = %v", nums)
			}
		})
	}
}

func BenchmarkSortedArrayToBSTRecursive(b *testing.B) {
	// Create a sorted array with 1000 elements
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedArrayToBSTRecursive(nums)
	}
}

func BenchmarkSortedArrayToBSTIterative(b *testing.B) {
	// Create a sorted array with 1000 elements
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedArrayToBSTIterative(nums)
	}
}