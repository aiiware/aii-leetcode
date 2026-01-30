package trees

import (
	"reflect"
	"testing"
    "leetcode/utils"
)

func TestSortedListToBSTSlowFast(t *testing.T) {
	tests := []struct {
		name     string
		head     *utils.ListNode
		validate func(*utils.TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			head: utils.NewListFromSlice([]int{-10, -3, 0, 5, 9}),
			validate: func(root *utils.TreeNode) bool {
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
			name: "Example 2: empty list",
			head: nil,
			validate: func(root *utils.TreeNode) bool {
				return root == nil
			},
		},
		{
			name: "Single element",
			head: utils.NewListFromSlice([]int{5}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{1, 2}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{1, 2, 3}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{1, 2, 3, 4}),
			validate: func(root *utils.TreeNode) bool {
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
			name: "Five elements",
			head: utils.NewListFromSlice([]int{1, 2, 3, 4, 5}),
			validate: func(root *utils.TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 2, 3, 4, 5})
			},
		},
		{
			name: "Negative values",
			head: utils.NewListFromSlice([]int{-5, -4, -3, -2, -1}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{-10, -5, 0, 5, 10}),
			validate: func(root *utils.TreeNode) bool {
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
			root := sortedListToBSTSlowFast(tt.head)
			if !tt.validate(root) {
				t.Errorf("sortedListToBSTSlowFast() failed validation")
			}
		})
	}
}

func TestSortedListToBSTArray(t *testing.T) {
	tests := []struct {
		name     string
		head     *utils.ListNode
		validate func(*utils.TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			head: utils.NewListFromSlice([]int{-10, -3, 0, 5, 9}),
			validate: func(root *utils.TreeNode) bool {
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
			name: "Example 2: empty list",
			head: nil,
			validate: func(root *utils.TreeNode) bool {
				return root == nil
			},
		},
		{
			name: "Single element",
			head: utils.NewListFromSlice([]int{5}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{1, 2, 3}),
			validate: func(root *utils.TreeNode) bool {
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
			name: "Five elements",
			head: utils.NewListFromSlice([]int{1, 2, 3, 4, 5}),
			validate: func(root *utils.TreeNode) bool {
				if !isValidBST(root) {
					return false
				}
				if !isBalancedBottomUp(root) {
					return false
				}
				inorder := inorderTraversal(root)
				return reflect.DeepEqual(inorder, []int{1, 2, 3, 4, 5})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedListToBSTArray(tt.head)
			if !tt.validate(root) {
				t.Errorf("sortedListToBSTArray() failed validation")
			}
		})
	}
}

func TestSortedListToBSTInorder(t *testing.T) {
	tests := []struct {
		name     string
		head     *utils.ListNode
		validate func(*utils.TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			head: utils.NewListFromSlice([]int{-10, -3, 0, 5, 9}),
			validate: func(root *utils.TreeNode) bool {
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
			name: "Example 2: empty list",
			head: nil,
			validate: func(root *utils.TreeNode) bool {
				return root == nil
			},
		},
		{
			name: "Single element",
			head: utils.NewListFromSlice([]int{5}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{1, 2, 3}),
			validate: func(root *utils.TreeNode) bool {
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
			head: utils.NewListFromSlice([]int{1, 2, 3, 4}),
			validate: func(root *utils.TreeNode) bool {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedListToBSTInorder(tt.head)
			if !tt.validate(root) {
				t.Errorf("sortedListToBSTInorder() failed validation")
			}
		})
	}
}

func TestSortedListToBST(t *testing.T) {
	tests := []struct {
		name     string
		head     *utils.ListNode
		validate func(*utils.TreeNode) bool
	}{
		{
			name: "Example 1: [-10,-3,0,5,9]",
			head: utils.NewListFromSlice([]int{-10, -3, 0, 5, 9}),
			validate: func(root *utils.TreeNode) bool {
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
			name: "Example 2: empty list",
			head: nil,
			validate: func(root *utils.TreeNode) bool {
				return root == nil
			},
		},
		{
			name: "Single element",
			head: utils.NewListFromSlice([]int{5}),
			validate: func(root *utils.TreeNode) bool {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedListToBST(tt.head)
			if !tt.validate(root) {
				t.Errorf("sortedListToBST() failed validation")
			}
		})
	}
}

func TestAllImplementationsProduceEquivalentResults(t *testing.T) {
	testCases := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{-5, -4, -3, -2, -1},
		{-10, -5, 0, 5, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	for _, nums := range testCases {
		t.Run("", func(t *testing.T) {
			head := utils.NewListFromSlice(nums)

			slowFastRoot := sortedListToBSTSlowFast(head)
			arrayRoot := sortedListToBSTArray(utils.NewListFromSlice(nums)) // Need fresh list
			inorderRoot := sortedListToBSTInorder(utils.NewListFromSlice(nums)) // Need fresh list

			// Get inorder traversals
			slowFastInorder := inorderTraversal(slowFastRoot)
			arrayInorder := inorderTraversal(arrayRoot)
			inorderInorder := inorderTraversal(inorderRoot)

			// All should produce the same inorder traversal
			if !reflect.DeepEqual(slowFastInorder, nums) {
				t.Errorf("SlowFast: expected inorder %v, got %v", nums, slowFastInorder)
			}
			if !reflect.DeepEqual(arrayInorder, nums) {
				t.Errorf("Array: expected inorder %v, got %v", nums, arrayInorder)
			}
			if !reflect.DeepEqual(inorderInorder, nums) {
				t.Errorf("Inorder: expected inorder %v, got %v", nums, inorderInorder)
			}

			// All should be valid BSTs
			if !isValidBST(slowFastRoot) {
				t.Errorf("SlowFast result not a valid BST for nums = %v", nums)
			}
			if !isValidBST(arrayRoot) {
				t.Errorf("Array result not a valid BST for nums = %v", nums)
			}
			if !isValidBST(inorderRoot) {
				t.Errorf("Inorder result not a valid BST for nums = %v", nums)
			}

			// All should be balanced
			if !isBalancedBottomUp(slowFastRoot) {
				t.Errorf("SlowFast result not balanced for nums = %v", nums)
			}
			if !isBalancedBottomUp(arrayRoot) {
				t.Errorf("Array result not balanced for nums = %v", nums)
			}
			if !isBalancedBottomUp(inorderRoot) {
				t.Errorf("Inorder result not balanced for nums = %v", nums)
			}
		})
	}
}

func BenchmarkSortedListToBSTSlowFast(b *testing.B) {
	// Create a sorted linked list with 1000 elements
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	head := utils.NewListFromSlice(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedListToBSTSlowFast(head)
	}
}

func BenchmarkSortedListToBSTArray(b *testing.B) {
	// Create a sorted linked list with 1000 elements
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	head := utils.NewListFromSlice(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedListToBSTArray(head)
	}
}

func BenchmarkSortedListToBSTInorder(b *testing.B) {
	// Create a sorted linked list with 1000 elements
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	head := utils.NewListFromSlice(nums)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortedListToBSTInorder(head)
	}
}