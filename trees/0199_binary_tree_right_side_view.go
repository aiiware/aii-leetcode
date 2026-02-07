package trees

import "leetcode/utils"

// RightSideView solves LeetCode problem 0199: Binary Tree Right Side View
// Difficulty: Medium
// Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
//
// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.
//
// Example:
// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]
//
// Time complexity: O(n), Space complexity: O(n)
func RightSideView(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		
		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// If this is the last node at current level, add to result
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}