package leetcode

// Node represents a node in a perfect binary tree with next pointer.
// This is the definition used in LeetCode problem 116.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect connects each node's next pointer to its next right node in a perfect binary tree.
// Problem 116: Populating Next Right Pointers in Each Node
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
//
// Approach: Level order traversal using existing next pointers (O(1) space)
// Time complexity: O(n), Space complexity: O(1)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Start with the leftmost node at the current level
	leftmost := root

	// Process each level
	for leftmost.Left != nil {
		// Traverse the current level and connect the next level
		current := leftmost
		for current != nil {
			// Connect left child to right child
			current.Left.Next = current.Right

			// Connect right child to next node's left child (if exists)
			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}

			// Move to next node in current level
			current = current.Next
		}

		// Move to the next level
		leftmost = leftmost.Left
	}

	return root
}

// connectBFS is an alternative solution using BFS with queue.
// Time complexity: O(n), Space complexity: O(n)
func connectBFS(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// Process each level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Connect to next node in same level (except last node)
			if i < levelSize-1 {
				node.Next = queue[0]
			}

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

// NewPerfectTree creates a perfect binary tree with given number of levels.
// Returns a tree with 2^levels - 1 nodes.
func NewPerfectTree(levels int) *Node {
	if levels <= 0 {
		return nil
	}

	// Total nodes in a perfect binary tree = 2^levels - 1
	totalNodes := 1<<levels - 1

	// Create all nodes
	nodes := make([]*Node, totalNodes+1) // 1-indexed for easier calculation
	for i := 1; i <= totalNodes; i++ {
		nodes[i] = &Node{Val: i}
	}

	// Connect parent-child relationships
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

// ToSliceWithNext converts a tree to a slice of slices representing each level.
// Each level is represented as a slice of node values.
func (n *Node) ToSliceWithNext() [][]int {
	if n == nil {
		return [][]int{}
	}

	result := [][]int{}
	current := n

	// Traverse using leftmost pointer
	for current != nil {
		level := []int{}
		node := current

		// Traverse current level using next pointers
		for node != nil {
			level = append(level, node.Val)
			node = node.Next
		}

		result = append(result, level)
		current = current.Left
	}

	return result
}

// GetNextValues returns a slice of next pointers for each node in level order.
// Returns -1 if next is nil.
func (n *Node) GetNextValues() []int {
	if n == nil {
		return []int{}
	}

	result := []int{}
	queue := []*Node{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Add next value (or -1 if nil)
		if node.Next != nil {
			result = append(result, node.Next.Val)
		} else {
			result = append(result, -1)
		}

		// Add children to queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

// Clone creates a deep copy of the tree.
func (n *Node) Clone() *Node {
	if n == nil {
		return nil
	}

	newNode := &Node{Val: n.Val}
	newNode.Left = n.Left.Clone()
	newNode.Right = n.Right.Clone()
	// Note: Next pointer is not cloned as it will be set by connect()
	return newNode
}