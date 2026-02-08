package arrays

// PreorderTraversal returns the preorder traversal of a binary tree's nodes' values
func PreorderTraversal(root *TreeNode) []int {
	result := []int{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}