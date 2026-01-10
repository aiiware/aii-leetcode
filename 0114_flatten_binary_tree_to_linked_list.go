package leetcode

// Problem 0114: Flatten Binary Tree to Linked List
//
// Given the root of a binary tree, flatten the tree into a "linked list":
// - The "linked list" should use the same TreeNode class where the right child
//   pointer points to the next node in the list and the left child pointer is
//   always nil.
// - The "linked list" should be in the same order as a pre-order traversal of
//   the binary tree.
//
// Example 1:
// Input: root = [1,2,5,3,4,null,6]
// Output: [1,null,2,null,3,null,4,null,5,null,6]
// Explanation:
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// 
// Flattened:
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [0]
// Output: [0]
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 2000].
// - -100 <= Node.val <= 100

// flatten flattens a binary tree to a linked list in-place using recursion.
// This approach uses a modified post-order traversal (right-left-root).
// Time complexity: O(n), Space complexity: O(h) where h is tree height
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	
	// Flatten left and right subtrees first
	flatten(root.Left)
	flatten(root.Right)
	
	// Store the original right subtree
	right := root.Right
	
	// Move left subtree to right
	root.Right = root.Left
	root.Left = nil
	
	// Find the end of the new right subtree (which was left subtree)
	current := root
	for current.Right != nil {
		current = current.Right
	}
	
	// Append the original right subtree
	current.Right = right
}

// flattenIterative flattens a binary tree to a linked list using iterative approach.
// This approach uses a stack to simulate recursion.
// Time complexity: O(n), Space complexity: O(n) for stack
func flattenIterative(root *TreeNode) {
	if root == nil {
		return
	}
	
	stack := []*TreeNode{root}
	var prev *TreeNode
	
	for len(stack) > 0 {
		// Pop from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Push right then left to stack (pre-order: root, left, right)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		
		// Link nodes
		if prev != nil {
			prev.Left = nil
			prev.Right = node
		}
		prev = node
	}
}

// flattenMorris flattens a binary tree to a linked list using Morris traversal.
// This approach modifies the tree structure in-place without using extra space.
// Time complexity: O(n), Space complexity: O(1)
func flattenMorris(root *TreeNode) {
	current := root
	
	for current != nil {
		if current.Left != nil {
			// Find the rightmost node in the left subtree
			predecessor := current.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			
			// Rewire connections
			predecessor.Right = current.Right
			current.Right = current.Left
			current.Left = nil
		}
		
		// Move to the next node
		current = current.Right
	}
}

// flattenReversePostOrder flattens using reverse post-order traversal.
// This approach processes nodes in reverse order (right-left-root).
// Time complexity: O(n), Space complexity: O(h) for recursion
func flattenReversePostOrder(root *TreeNode) {
	var prev *TreeNode
	
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		// Process right subtree first, then left, then root
		dfs(node.Right)
		dfs(node.Left)
		
		// Link nodes
		node.Right = prev
		node.Left = nil
		prev = node
	}
	
	dfs(root)
}