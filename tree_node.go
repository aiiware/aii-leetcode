package leetcode

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeFromSlice creates a binary tree from a level-order slice representation.
// nil values in the slice represent missing nodes.
// This handles the compact LeetCode representation where nil nodes don't have their children represented.
func NewTreeFromSlice(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	
	// Use a queue to process nodes in level order
	queue := []*TreeNode{root}
	
	// We'll use i to track our position in the vals array
	i := 1
	
	for len(queue) > 0 && i < len(vals) {
		// Get the next node to process
		node := queue[0]
		queue = queue[1:]
		
		// Process left child
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Left)
			}
			// If vals[i] is nil, we don't add anything to the queue
			i++
		}
		
		// Process right child
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Right)
			}
			// If vals[i] is nil, we don't add anything to the queue
			i++
		}
	}

	return root
}

// ToSlice converts a binary tree to a level-order slice representation.
func (t *TreeNode) ToSlice() []*int {
	if t == nil {
		return []*int{}
	}

	var result []*int
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, &node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Remove trailing nil values
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// Equal compares two trees for equality
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

// Helper functions for testing

// CloneTree creates a deep copy of a binary tree
func CloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  CloneTree(root.Left),
		Right: CloneTree(root.Right),
	}
}

// CreateCompleteTree creates a complete binary tree with n nodes
func CreateCompleteTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	// Create nodes
	nodes := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &TreeNode{Val: i + 1}
	}
	
	// Connect parent-child relationships
	for i := 0; i < n; i++ {
		leftIdx := 2*i + 1
		if leftIdx < n {
			nodes[i].Left = nodes[leftIdx]
		}
		
		rightIdx := 2*i + 2
		if rightIdx < n {
			nodes[i].Right = nodes[rightIdx]
		}
	}
	
	return nodes[0]
}

// CreateRightSkewedTree creates a right-skewed tree with n nodes
func CreateRightSkewedTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	root := &TreeNode{Val: 1}
	current := root
	
	for i := 2; i <= n; i++ {
		current.Right = &TreeNode{Val: i}
		current = current.Right
	}
	
	return root
}

// CreateLeftSkewedTree creates a left-skewed tree with n nodes
func CreateLeftSkewedTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	root := &TreeNode{Val: 1}
	current := root
	
	for i := 2; i <= n; i++ {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}
	
	return root
}

// CreateSymmetricTree creates a symmetric binary tree with specified number of levels
// levels=1: single node, levels=2: 3 nodes, levels=3: 7 nodes, etc.
// The tree will have symmetric values: for 3 levels: [1, 2, 2, 3, 4, 4, 3]
func CreateSymmetricTree(levels int) *TreeNode {
	if levels <= 0 {
		return nil
	}
	
	// Actually, let's use a simpler approach: build the tree level by level
	// with symmetric values
	
	// We'll use a queue to build the tree level by level
	if levels == 1 {
		return &TreeNode{Val: 1}
	}
	
	// Create root
	root := &TreeNode{Val: 1}
	
	// We'll use a slice to track nodes at each level
	currentLevel := []*TreeNode{root}
	nextVal := 2  // Next value to assign
	
	for level := 1; level < levels; level++ {
		nextLevel := make([]*TreeNode, 0, 1<<level)
		
		// For a symmetric tree, values at this level should be symmetric
		// Number of nodes at this level: 2^level
		nodesThisLevel := 1 << level
		
		// Create values for this level
		// First half: increasing values
		// Second half: mirror of first half
		half := nodesThisLevel / 2
		values := make([]int, nodesThisLevel)
		
		for i := 0; i < half; i++ {
			values[i] = nextVal
			nextVal++
		}
		
		// Mirror the values for the second half
		for i := 0; i < half; i++ {
			values[nodesThisLevel-1-i] = values[i]
		}
		
		// Create nodes for this level
		for i := 0; i < nodesThisLevel; i++ {
			nextLevel = append(nextLevel, &TreeNode{Val: values[i]})
		}
		
		// Connect parent-child relationships
		for i := 0; i < len(currentLevel); i++ {
			parent := currentLevel[i]
			leftIdx := 2 * i
			rightIdx := 2*i + 1
			
			if leftIdx < len(nextLevel) {
				parent.Left = nextLevel[leftIdx]
			}
			if rightIdx < len(nextLevel) {
				parent.Right = nextLevel[rightIdx]
			}
		}
		
		// Move to next level
		currentLevel = nextLevel
	}
	
	return root
}