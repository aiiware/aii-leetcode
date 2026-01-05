package leetcode

// Problem 0100: Same Tree
//
// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical, 
// and the nodes have the same value.
//
// Example 1:
// Input: p = [1,2,3], q = [1,2,3]
// Output: true
//
// Example 2:
// Input: p = [1,2], q = [1,null,2]
// Output: false
//
// Example 3:
// Input: p = [1,2,1], q = [1,1,2]
// Output: false
//
// Constraints:
// - The number of nodes in both trees is in the range [0, 100].
// - -10^4 <= Node.val <= 10^4

// isSameTree is the main solution function using recursion.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base cases
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	
	// Recursively check left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// isSameTreeIterative uses iterative approach with stacks.
func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	// Use two stacks for simultaneous traversal
	stackP := []*TreeNode{p}
	stackQ := []*TreeNode{q}
	
	for len(stackP) > 0 && len(stackQ) > 0 {
		// Pop from both stacks
		nodeP := stackP[len(stackP)-1]
		stackP = stackP[:len(stackP)-1]
		
		nodeQ := stackQ[len(stackQ)-1]
		stackQ = stackQ[:len(stackQ)-1]
		
		// Check if both are nil
		if nodeP == nil && nodeQ == nil {
			continue
		}
		
		// Check if one is nil
		if nodeP == nil || nodeQ == nil {
			return false
		}
		
		// Check values
		if nodeP.Val != nodeQ.Val {
			return false
		}
		
		// Push children (push right first for preorder)
		stackP = append(stackP, nodeP.Right, nodeP.Left)
		stackQ = append(stackQ, nodeQ.Right, nodeQ.Left)
	}
	
	// Both stacks should be empty
	return len(stackP) == 0 && len(stackQ) == 0
}

// isSameTreeBFS uses BFS (level order traversal).
func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	// Use two queues for simultaneous BFS
	queueP := []*TreeNode{p}
	queueQ := []*TreeNode{q}
	
	for len(queueP) > 0 && len(queueQ) > 0 {
		// Dequeue from both queues
		nodeP := queueP[0]
		queueP = queueP[1:]
		
		nodeQ := queueQ[0]
		queueQ = queueQ[1:]
		
		// Check if both are nil
		if nodeP == nil && nodeQ == nil {
			continue
		}
		
		// Check if one is nil
		if nodeP == nil || nodeQ == nil {
			return false
		}
		
		// Check values
		if nodeP.Val != nodeQ.Val {
			return false
		}
		
		// Enqueue children
		queueP = append(queueP, nodeP.Left, nodeP.Right)
		queueQ = append(queueQ, nodeQ.Left, nodeQ.Right)
	}
	
	// Both queues should be empty
	return len(queueP) == 0 && len(queueQ) == 0
}

// isSameTreeDFS uses DFS with explicit stack.
func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	type stackItem struct {
		p *TreeNode
		q *TreeNode
	}
	
	stack := []stackItem{{p, q}}
	
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Check if both are nil
		if item.p == nil && item.q == nil {
			continue
		}
		
		// Check if one is nil
		if item.p == nil || item.q == nil {
			return false
		}
		
		// Check values
		if item.p.Val != item.q.Val {
			return false
		}
		
		// Push children
		stack = append(stack, stackItem{item.p.Right, item.q.Right})
		stack = append(stack, stackItem{item.p.Left, item.q.Left})
	}
	
	return true
}

// isSameTreeMorris uses Morris traversal (not ideal for comparison).
func isSameTreeMorris(p *TreeNode, q *TreeNode) bool {
	// Morris traversal modifies the tree, so not suitable for comparison
	// We'll implement a simplified version that doesn't modify trees
	return isSameTree(p, q) // Fall back to recursive
}

// isSameTreeSerialization uses tree serialization.
func isSameTreeSerialization(p *TreeNode, q *TreeNode) bool {
	// Serialize both trees and compare
	return serializeTree(p) == serializeTree(q)
}

func serializeTree(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return fmt.Sprintf("(%d %s %s)", root.Val, 
		serializeTree(root.Left), serializeTree(root.Right))
}

// isSameTreePreorder uses preorder traversal comparison.
func isSameTreePreorder(p *TreeNode, q *TreeNode) bool {
	var preorder func(*TreeNode, *[]interface{})
	preorder = func(node *TreeNode, result *[]interface{}) {
		if node == nil {
			*result = append(*result, nil)
			return
		}
		*result = append(*result, node.Val)
		preorder(node.Left, result)
		preorder(node.Right, result)
	}
	
	var preorderP []interface{}
	var preorderQ []interface{}
	
	preorder(p, &preorderP)
	preorder(q, &preorderQ)
	
	if len(preorderP) != len(preorderQ) {
		return false
	}
	
	for i := 0; i < len(preorderP); i++ {
		if preorderP[i] != preorderQ[i] {
			return false
		}
	}
	
	return true
}

// isSameTreeOptimized is an optimized version.
func isSameTreeOptimized(p *TreeNode, q *TreeNode) bool {
	// Short circuit: both nil
	if p == nil && q == nil {
		return true
	}
	
	// Short circuit: one nil
	if p == nil || q == nil {
		return false
	}
	
	// Use a stack for iterative comparison
	type pair struct {
		p *TreeNode
		q *TreeNode
	}
	
	stack := []pair{{p, q}}
	
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Check values
		if current.p.Val != current.q.Val {
			return false
		}
		
		// Check left children
		if current.p.Left != nil && current.q.Left != nil {
			stack = append(stack, pair{current.p.Left, current.q.Left})
		} else if current.p.Left != nil || current.q.Left != nil {
			return false
		}
		
		// Check right children
		if current.p.Right != nil && current.q.Right != nil {
			stack = append(stack, pair{current.p.Right, current.q.Right})
		} else if current.p.Right != nil || current.q.Right != nil {
			return false
		}
	}
	
	return true
}

// IsSameTree is the public interface function.
// It uses the optimized iterative solution by default.
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeOptimized(p, q)
}