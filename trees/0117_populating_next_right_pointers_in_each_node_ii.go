package trees


/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

// Node117 definition for problem 117 (Populating Next Right Pointers in Each Node II)
type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

// connect117 connects each node to its next right node in a binary tree (any binary tree)
// This solution works for any binary tree, not just perfect trees
func connect117(root *Node117) *Node117 {
	if root == nil {
		return nil
	}

	// Use a dummy head for each level
	dummy := &Node117{}
	// prev tracks the previous node in the current level
	prev := dummy
	// current starts at root
	current := root

	for current != nil {
		// Process current level
		for current != nil {
			// Connect left child if exists
			if current.Left != nil {
				prev.Next = current.Left
				prev = prev.Next
			}
			// Connect right child if exists
			if current.Right != nil {
				prev.Next = current.Right
				prev = prev.Next
			}
			// Move to next node in current level
			current = current.Next
		}

		// Move to next level
		current = dummy.Next
		// Reset dummy and prev for next level
		dummy.Next = nil
		prev = dummy
	}

	return root
}

// connectBFS117 is an alternative BFS solution using level-order traversal
func connectBFS117(root *Node117) *Node117 {
	if root == nil {
		return nil
	}

	queue := []*Node117{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var prev *Node117

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Connect to previous node in same level
			if prev != nil {
				prev.Next = node
			}
			prev = node

			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

// Helper function to create a binary tree from slice representation
// nil values represent null nodes
// This implementation properly handles the LeetCode tree serialization format
func createTreeFromSlice117(values []interface{}) *Node117 {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &Node117{Val: values[0].(int)}
	// Use a queue to process nodes in level order
	nodes := make([]*Node117, len(values))
	nodes[0] = root

	for i := 0; i < len(values); i++ {
		if nodes[i] == nil {
			continue
		}

		// Calculate indices for left and right children
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2

		// Create left child if exists and within bounds
		if leftIdx < len(values) && values[leftIdx] != nil {
			nodes[leftIdx] = &Node117{Val: values[leftIdx].(int)}
			nodes[i].Left = nodes[leftIdx]
		}

		// Create right child if exists and within bounds
		if rightIdx < len(values) && values[rightIdx] != nil {
			nodes[rightIdx] = &Node117{Val: values[rightIdx].(int)}
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return root
}

// Helper function to traverse tree via next pointers (level by level)
func traverseByLevel117(root *Node117) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	current := root

	for current != nil {
		levelStart := current
		level := []int{}
		for levelStart != nil {
			level = append(level, levelStart.Val)
			levelStart = levelStart.Next
		}
		result = append(result, level)

		// Find first node of next level
		current = findFirstChild117(current)
	}

	return result
}

// Helper to find first child in next level
func findFirstChild117(node *Node117) *Node117 {
	for node != nil {
		if node.Left != nil {
			return node.Left
		}
		if node.Right != nil {
			return node.Right
		}
		node = node.Next
	}
	return nil
}