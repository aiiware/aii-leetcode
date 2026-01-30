package main

import (
	"fmt"
	
	"leetcode/arrays"
	"leetcode/design"
	"leetcode/linkedlists"
	"leetcode/trees"
	"leetcode/utils"
)

func main() {
	fmt.Println("=== LeetCode Solutions Demo ===")

	// Problem 0141: Linked List Cycle
	fmt.Println("\n--- Problem 0141: Linked List Cycle ---")
	head1 := utils.IntsToList([]int{3, 2, 0, -4})
	// Create a cycle: -4 -> 2
	if head1 != nil && head1.Next != nil && head1.Next.Next != nil {
		head1.Next.Next.Next.Next = head1.Next
	}
	fmt.Printf("Has cycle: %v\n", linkedlists.HasCycle(head1))

	// Problem 0142: Linked List Cycle II
	fmt.Println("\n--- Problem 0142: Linked List Cycle II ---")
	head2 := utils.IntsToList([]int{3, 2, 0, -4})
	// Create a cycle: -4 -> 2
	if head2 != nil && head2.Next != nil && head2.Next.Next != nil {
		head2.Next.Next.Next.Next = head2.Next
	}
	cycleStart := linkedlists.DetectCycle(head2)
	if cycleStart != nil {
		fmt.Printf("Cycle starts at node with value: %d\n", cycleStart.Val)
	} else {
		fmt.Println("No cycle detected")
	}

	// Problem 0143: Reorder List
	fmt.Println("\n--- Problem 0143: Reorder List ---")
	head3 := utils.IntsToList([]int{1, 2, 3, 4})
	head3Copy := utils.CopyList(head3)
	linkedlists.ReorderList(head3Copy)
	fmt.Printf("Original: %s\n", utils.PrintList(head3))
	fmt.Printf("Reordered: %s\n", utils.PrintList(head3Copy))

	head3b := utils.IntsToList([]int{1, 2, 3, 4, 5})
	head3bCopy := utils.CopyList(head3b)
	linkedlists.ReorderListStack(head3bCopy)
	fmt.Printf("Original: %s\n", utils.PrintList(head3b))
	fmt.Printf("Reordered (stack): %s\n", utils.PrintList(head3bCopy))

	head3c := utils.IntsToList([]int{1, 2, 3, 4})
	head3cCopy := utils.CopyList(head3c)
	linkedlists.ReorderListArray(head3cCopy)
	fmt.Printf("Original: %s\n", utils.PrintList(head3c))
	fmt.Printf("Reordered (array): %s\n", utils.PrintList(head3cCopy))

	// Problem 0144: Binary Tree Preorder Traversal
	fmt.Println("\n--- Problem 0144: Binary Tree Preorder Traversal ---")
	root4 := &utils.TreeNode{
		Val: 1,
		Right: &utils.TreeNode{
			Val: 2,
			Left: &utils.TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Printf("Recursive: %v\n", trees.PreorderTraversal(root4))
	fmt.Printf("Iterative: %v\n", trees.PreorderTraversalIterative(root4))
	fmt.Printf("Morris: %v\n", trees.PreorderTraversalMorris(root4))

	// Problem 0145: Binary Tree Postorder Traversal
	fmt.Println("\n--- Problem 0145: Binary Tree Postorder Traversal ---")
	root5 := &utils.TreeNode{
		Val: 1,
		Right: &utils.TreeNode{
			Val: 2,
			Left: &utils.TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Printf("Recursive: %v\n", trees.PostorderTraversal(root5))
	fmt.Printf("Iterative: %v\n", trees.PostorderTraversalIterative(root5))
	fmt.Printf("Two Stacks: %v\n", trees.PostorderTraversalTwoStacks(root5))

	// Problem 0146: LRU Cache
	fmt.Println("\n--- Problem 0146: LRU Cache ---")
	lru := design.Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Printf("Get(1): %d\n", lru.Get(1))
	lru.Put(3, 3) // evicts key 2
	fmt.Printf("Get(2): %d\n", lru.Get(2))
	lru.Put(4, 4) // evicts key 1
	fmt.Printf("Get(1): %d\n", lru.Get(1))
	fmt.Printf("Get(3): %d\n", lru.Get(3))
	fmt.Printf("Get(4): %d\n", lru.Get(4))

	// Problem 0147: Insertion Sort List
	fmt.Println("\n--- Problem 0147: Insertion Sort List ---")
	head7 := utils.IntsToList([]int{4, 2, 1, 3})
	head7Copy := utils.CopyList(head7)
	sorted7 := linkedlists.InsertionSortList(head7Copy)
	fmt.Printf("Original: %s\n", utils.PrintList(head7))
	fmt.Printf("Sorted: %s\n", utils.PrintList(sorted7))

	// Problem 0148: Sort List
	fmt.Println("\n--- Problem 0148: Sort List ---")
	head8 := utils.IntsToList([]int{4, 2, 1, 3})
	head8Copy := utils.CopyList(head8)
	sorted8 := linkedlists.SortList(head8Copy)
	fmt.Printf("Original: %s\n", utils.PrintList(head8))
	fmt.Printf("Sorted: %s\n", utils.PrintList(sorted8))

	// Problem 0149: Max Points on a Line
	fmt.Println("\n--- Problem 0149: Max Points on a Line ---")
	points9 := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Printf("Points: %v\n", points9)
	fmt.Printf("Max points on a line: %d\n", arrays.MaxPoints(points9))

	points9b := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	fmt.Printf("Points: %v\n", points9b)
	fmt.Printf("Max points on a line: %d\n", arrays.MaxPoints(points9b))

	// Problem 0150: Evaluate Reverse Polish Notation
	fmt.Println("\n--- Problem 0150: Evaluate Reverse Polish Notation ---")
	tokens10 := []string{"2", "1", "+", "3", "*"}
	fmt.Printf("Tokens: %v\n", tokens10)
	fmt.Printf("Result: %d\n", arrays.EvalRPN(tokens10))

	tokens10b := []string{"4", "13", "5", "/", "+"}
	fmt.Printf("Tokens: %v\n", tokens10b)
	fmt.Printf("Result: %d\n", arrays.EvalRPN(tokens10b))

	tokens10c := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Printf("Tokens: %v\n", tokens10c)
	fmt.Printf("Result: %d\n", arrays.EvalRPN(tokens10c))

	fmt.Println("\n=== Demo Complete ===")
}