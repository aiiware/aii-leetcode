package main

import (
	"fmt"
    "leetcode/utils"
)

type utils.TreeNode struct {
	Val   int
	Left  *utils.TreeNode
	Right *utils.TreeNode
}

func intPtr(x int) *int {
	return &x
}

func NewTreeFromSlice(vals []*int) *utils.TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &utils.TreeNode{Val: *vals[0]}
	queue := []*utils.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != nil {
			node.Left = &utils.TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &utils.TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func printTree(root *utils.TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}
	
	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Printf("├── ")
	} else {
		fmt.Printf("└── ")
	}
	fmt.Printf("%d\n", root.Val)
	
	// Compute new prefix
	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}
	
	// Print children
	printTree(root.Left, newPrefix, true)
	printTree(root.Right, newPrefix, false)
}

func main() {
	// Test the complex symmetric tree case that's failing
	vals := []*int{
		intPtr(1),
		intPtr(2), intPtr(2),
		intPtr(3), intPtr(4), intPtr(4), intPtr(3),
		intPtr(5), intPtr(6), intPtr(7), intPtr(7), intPtr(6), intPtr(5),
	}
	
	fmt.Println("Tree structure from level-order slice:")
	printTree(NewTreeFromSlice(vals), "", false)
	
	// Let's trace through the construction
	fmt.Println("\nConstruction trace:")
	root := NewTreeFromSlice(vals)
	fmt.Printf("Root: %d\n", root.Val)
	fmt.Printf("Root.Left: %d\n", root.Left.Val)
	fmt.Printf("Root.Right: %d\n", root.Right.Val)
	fmt.Printf("Root.Left.Left: %d\n", root.Left.Left.Val)
	fmt.Printf("Root.Left.Right: %d\n", root.Left.Right.Val)
	fmt.Printf("Root.Right.Left: %d\n", root.Right.Left.Val)
	fmt.Printf("Root.Right.Right: %d\n", root.Right.Right.Val)
	fmt.Printf("Root.Left.Left.Left: %d\n", root.Left.Left.Left.Val)
	fmt.Printf("Root.Left.Left.Right: %d\n", root.Left.Left.Right.Val)
	fmt.Printf("Root.Left.Right.Left: %d\n", root.Left.Right.Left.Val)
	fmt.Printf("Root.Left.Right.Right: %d\n", root.Left.Right.Right.Val)
	fmt.Printf("Root.Right.Left.Left: %d\n", root.Right.Left.Left.Val)
	fmt.Printf("Root.Right.Left.Right: %d\n", root.Right.Left.Right.Val)
	// Root.Right.Right has no children because we ran out of values
	if root.Right.Right.Left == nil {
		fmt.Println("Root.Right.Right.Left: nil")
	}
	if root.Right.Right.Right == nil {
		fmt.Println("Root.Right.Right.Right: nil")
	}
}