package leetcode

/*
173. Binary Search Tree Iterator

Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):

- BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.
- boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
- int next() Moves the pointer to the right, then returns the number at the pointer.

Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.

You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.

Example 1:
Input
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]

Explanation
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // return 3
bSTIterator.next();    // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 20
bSTIterator.hasNext(); // return False

Constraints:
- The number of nodes in the tree is in the range [1, 10^5].
- 0 <= Node.val <= 10^6
- At most 10^5 calls will be made to hasNext, and next.

Difficulty: Medium
Tags: Stack, Tree, Design, Binary Search Tree, Binary Tree, Iterator
Companies: Facebook, Google, Amazon, Microsoft, Bloomberg
*/

// BSTIterator implements an iterator over a binary search tree (BST) that returns elements in ascending order
type BSTIterator struct {
	stack []*TreeNode
}

// BSTIteratorConstructor initializes a BSTIterator with the given root
func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{
		stack: make([]*TreeNode, 0),
	}
	// Push all left nodes to initialize the stack
	iterator.pushAllLeft(root)
	return iterator
}

// pushAllLeft pushes all left nodes from the given node onto the stack
func (this *BSTIterator) pushAllLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

// Next returns the next smallest number in the BST
func (this *BSTIterator) Next() int {
	// Pop the top node from stack
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	
	// If the popped node has a right child, push all left nodes of the right child
	if top.Right != nil {
		this.pushAllLeft(top.Right)
	}
	
	return top.Val
}

// HasNext returns whether we have a next smallest number
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */