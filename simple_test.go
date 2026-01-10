package leetcode

import (
	"fmt"
	"testing"
)

func TestSimpleBuildTree(t *testing.T) {
	// Test a simple case
	preorder := []int{1, 2, 3, 4}
	inorder := []int{4, 3, 2, 1}
	
	fmt.Println("Testing buildTree with:")
	fmt.Printf("preorder: %v\n", preorder)
	fmt.Printf("inorder: %v\n", inorder)
	
	root := buildTreeRecursive(preorder, inorder)
	result1 := treeToSliceSimple(root)
	fmt.Printf("Recursive result: %v\n", result1)
	
	root = buildTreeOptimized(preorder, inorder)
	result2 := treeToSliceSimple(root)
	fmt.Printf("Optimized result: %v\n", result2)
	
	root = buildTreeIterative(preorder, inorder)
	result3 := treeToSliceSimple(root)
	fmt.Printf("Iterative result: %v\n", result3)
}

func treeToSliceSimple(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, -999) // Use -999 to represent nil
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Remove trailing -999 values
	for len(result) > 0 && result[len(result)-1] == -999 {
		result = result[:len(result)-1]
	}

	return result
}
