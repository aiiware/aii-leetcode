package leetcode

// Problem 0099: Recover Binary Search Tree
//
// You are given the root of a binary search tree (BST), where the values of exactly two nodes 
// of the tree were swapped by mistake. Recover the tree without changing its structure.
//
// Example 1:
// Input: root = [1,3,null,null,2]
// Output: [3,1,null,null,2]
// Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.
//
// Example 2:
// Input: root = [3,1,4,null,null,2]
// Output: [2,1,4,null,null,3]
// Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.
//
// Constraints:
// - The number of nodes in the tree is in the range [2, 1000].
// - -2^31 <= Node.val <= 2^31 - 1

// recoverTree is the main solution function using inorder traversal.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	
	// Inorder traversal to find the two swapped nodes
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		inorder(node.Left)
		
		// Check for swapped nodes
		if prev != nil && node.Val < prev.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		
		inorder(node.Right)
	}
	
	inorder(root)
	
	// Swap the values of the two nodes
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeIterative uses iterative inorder traversal.
func recoverTreeIterative(root *TreeNode) {
	var first, second, prev *TreeNode
	stack := []*TreeNode{}
	current := root
	
	for current != nil || len(stack) > 0 {
		// Go to leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Check for swapped nodes
		if prev != nil && current.Val < prev.Val {
			if first == nil {
				first = prev
			}
			second = current
		}
		prev = current
		
		// Go to right subtree
		current = current.Right
	}
	
	// Swap values
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeMorris uses Morris traversal (threaded binary tree).
// Time complexity: O(n), Space complexity: O(1)
func recoverTreeMorris(root *TreeNode) {
	var first, second, prev *TreeNode
	current := root
	
	for current != nil {
		if current.Left == nil {
			// Visit current node
			if prev != nil && current.Val < prev.Val {
				if first == nil {
					first = prev
				}
				second = current
			}
			prev = current
			current = current.Right
		} else {
			// Find inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Create thread to current
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove thread and visit current
				predecessor.Right = nil
				if prev != nil && current.Val < prev.Val {
					if first == nil {
						first = prev
					}
					second = current
				}
				prev = current
				current = current.Right
			}
		}
	}
	
	// Swap values
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeDFS uses DFS with explicit stack.
func recoverTreeDFS(root *TreeNode) {
	var first, second, prev *TreeNode
	
	type stackItem struct {
		node *TreeNode
		visited bool
	}
	
	stack := []stackItem{{node: root, visited: false}}
	
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if item.node == nil {
			continue
		}
		
		if item.visited {
			// Process node
			if prev != nil && item.node.Val < prev.Val {
				if first == nil {
					first = prev
				}
				second = item.node
			}
			prev = item.node
		} else {
			// Push in reverse order: right, node (visited), left
			stack = append(stack, stackItem{node: item.node.Right, visited: false})
			stack = append(stack, stackItem{node: item.node, visited: true})
			stack = append(stack, stackItem{node: item.node.Left, visited: false})
		}
	}
	
	// Swap values
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeSimple uses simple approach with inorder traversal array.
func recoverTreeSimple(root *TreeNode) {
	// Get inorder traversal
	nodes := []*TreeNode{}
	values := []int{}
	
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nodes = append(nodes, node)
		values = append(values, node.Val)
		inorder(node.Right)
	}
	
	inorder(root)
	
	// Sort values to find correct order
	sortedValues := make([]int, len(values))
	copy(sortedValues, values)
	
	// Simple bubble sort (could use sort.Ints, but implementing for clarity)
	for i := 0; i < len(sortedValues)-1; i++ {
		for j := 0; j < len(sortedValues)-i-1; j++ {
			if sortedValues[j] > sortedValues[j+1] {
				sortedValues[j], sortedValues[j+1] = sortedValues[j+1], sortedValues[j]
			}
		}
	}
	
	// Find nodes that need to be swapped
	var first, second *TreeNode
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Val != sortedValues[i] {
			if first == nil {
				first = nodes[i]
			} else {
				second = nodes[i]
				break
			}
		}
	}
	
	// Swap values
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeOptimized is an optimized version.
func recoverTreeOptimized(root *TreeNode) {
	var first, second, prev *TreeNode
	
	// Helper function for inorder traversal
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		dfs(node.Left)
		
		// Check for swapped nodes
		if prev != nil && node.Val < prev.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		
		dfs(node.Right)
	}
	
	dfs(root)
	
	// Swap the two nodes
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// recoverTreeTwoPass uses two-pass approach.
func recoverTreeTwoPass(root *TreeNode) {
	// First pass: find the two nodes
	var first, second *TreeNode
	var prev *TreeNode
	
	// Find first violation (larger node)
	var findFirst func(*TreeNode)
	findFirst = func(node *TreeNode) {
		if node == nil || first != nil {
			return
		}
		
		findFirst(node.Left)
		
		if prev != nil && node.Val < prev.Val && first == nil {
			first = prev
		}
		prev = node
		
		findFirst(node.Right)
	}
	
	findFirst(root)
	
	// Second pass: find second violation (smaller node)
	prev = nil
	var findSecond func(*TreeNode)
	findSecond = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		findSecond(node.Left)
		
		if prev != nil && node.Val < prev.Val {
			second = node
		}
		prev = node
		
		findSecond(node.Right)
	}
	
	findSecond(root)
	
	// Swap values
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// RecoverTree is the public interface function.
// It uses the optimized recursive solution by default.
func RecoverTree(root *TreeNode) {
	recoverTreeOptimized(root)
}