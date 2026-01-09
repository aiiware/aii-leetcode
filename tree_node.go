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

// CreateCompleteTree creates a complete binary tree with n nodes.
// Nodes are numbered from 1 to n in level order.
func CreateCompleteTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}

	// Create nodes
	nodes := make([]*TreeNode, n+1) // 1-indexed for easier calculation
	for i := 1; i <= n; i++ {
		nodes[i] = &TreeNode{Val: i}
	}

	// Connect nodes in level order
	for i := 1; i <= n; i++ {
		leftIdx := 2 * i
		rightIdx := 2*i + 1

		if leftIdx <= n {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx <= n {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[1]
}

// createCompleteTree is a lowercase version for internal use (to match test expectations)
func createCompleteTree(n int) *TreeNode {
	return CreateCompleteTree(n)
}

// createRightSkewedTree creates a right-skewed tree (linked list to the right)
func createRightSkewedTree(n int) *TreeNode {
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

// createLeftSkewedTree creates a left-skewed tree (linked list to the left)
func createLeftSkewedTree(n int) *TreeNode {
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

// cloneTree creates a deep copy of a binary tree
func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = cloneTree(root.Left)
	newRoot.Right = cloneTree(root.Right)
	return newRoot
}

// createSymmetricTree creates a symmetric binary tree with given number of levels
// Returns a tree with 2^levels - 1 nodes (perfect binary tree)
func createSymmetricTree(levels int) *TreeNode {
	if levels <= 0 {
		return nil
	}
	
	// Total nodes in a perfect binary tree = 2^levels - 1
	totalNodes := 1<<levels - 1
	
	// We'll build the tree level by level with symmetric values
	// Create all nodes first
	nodes := make([]*TreeNode, totalNodes+1) // 1-indexed
	
	// Assign values in a symmetric pattern
	// For each level, assign increasing values that are mirrored
	value := 1
	for level := 1; level <= levels; level++ {
		nodesInLevel := 1 << (level - 1) // 2^(level-1) nodes at this level
		startIdx := 1 << (level - 1)     // Starting index for this level
		
		// Assign values symmetrically
		for i := 0; i < nodesInLevel/2; i++ {
			// Left position in level
			leftPos := startIdx + i
			// Right position (mirror)
			rightPos := startIdx + nodesInLevel - 1 - i
			
			nodes[leftPos] = &TreeNode{Val: value}
			nodes[rightPos] = &TreeNode{Val: value}
			value++
		}
		
		// If odd number of nodes in level, assign middle value
		if nodesInLevel%2 == 1 {
			middlePos := startIdx + nodesInLevel/2
			nodes[middlePos] = &TreeNode{Val: value}
			value++
		}
	}
	
	// Now connect nodes to form a perfect binary tree
	for i := 1; i <= totalNodes/2; i++ {
		leftChild := 2 * i
		rightChild := 2*i + 1
		
		if leftChild <= totalNodes {
			nodes[i].Left = nodes[leftChild]
		}
		if rightChild <= totalNodes {
			nodes[i].Right = nodes[rightChild]
		}
	}
	
	return nodes[1]
}

// countNodes counts the number of nodes in a binary tree
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// isBST checks if a binary tree is a valid Binary Search Tree
func isBST(root *TreeNode) bool {
	var validate func(node *TreeNode, min, max int) bool
	validate = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
	}
	return validate(root, -1<<31, 1<<31-1)
}

// getAllValues collects all values from a binary tree
func getAllValues(root *TreeNode) []int {
	values := []int{}
	var collect func(node *TreeNode)
	collect = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		collect(node.Left)
		collect(node.Right)
	}
	collect(root)
	return values
}

// contains checks if a value exists in a slice
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// countValues counts the frequency of each value in a binary tree
func countValues(root *TreeNode, counts map[int]int) {
	if root == nil {
		return
	}
	counts[root.Val]++
	countValues(root.Left, counts)
	countValues(root.Right, counts)
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