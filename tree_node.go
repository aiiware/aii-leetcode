package leetcode

// TreeNode represents a node in a binary tree.
// This is the standard definition used in LeetCode problems.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new TreeNode with the given value.
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// NewTreeFromSlice creates a binary tree from a slice of integers.
// Uses LeetCode's level-order traversal representation where nil values
// are represented as null in JSON but we'll use a slice of *int.
// Returns the root of the binary tree.
func NewTreeFromSlice(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// ToSlice converts a binary tree to a slice of integers in level-order.
// Nil children are represented as nil pointers in the slice.
func (t *TreeNode) ToSlice() []*int {
	if t == nil {
		return []*int{}
	}

	result := []*int{}
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		val := node.Val
		result = append(result, &val)

		// Add children to queue
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// Equal compares two binary trees for equality.
func (t *TreeNode) Equal(other *TreeNode) bool {
	if t == nil && other == nil {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	if t.Val != other.Val {
		return false
	}
	return t.Left.Equal(other.Left) && t.Right.Equal(other.Right)
}

// InorderTraversal returns the inorder traversal of the tree.
func (t *TreeNode) InorderTraversal() []int {
	result := []int{}
	t.inorder(&result)
	return result
}

func (t *TreeNode) inorder(result *[]int) {
	if t == nil {
		return
	}
	t.Left.inorder(result)
	*result = append(*result, t.Val)
	t.Right.inorder(result)
}

// PreorderTraversal returns the preorder traversal of the tree.
func (t *TreeNode) PreorderTraversal() []int {
	result := []int{}
	t.preorder(&result)
	return result
}

func (t *TreeNode) preorder(result *[]int) {
	if t == nil {
		return
	}
	*result = append(*result, t.Val)
	t.Left.preorder(result)
	t.Right.preorder(result)
}

// PostorderTraversal returns the postorder traversal of the tree.
func (t *TreeNode) PostorderTraversal() []int {
	result := []int{}
	t.postorder(&result)
	return result
}

func (t *TreeNode) postorder(result *[]int) {
	if t == nil {
		return
	}
	t.Left.postorder(result)
	t.Right.postorder(result)
	*result = append(*result, t.Val)
}