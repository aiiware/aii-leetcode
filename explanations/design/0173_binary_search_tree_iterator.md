# 0173 - Binary Search Tree Iterator

## Problem Statement
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST. Calling `next()` will return the next smallest number in the BST, and `hasNext()` should return whether there exists a next smallest number.

**Example:**
```
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return false
```

## Solution Approach
The BST Iterator problem can be solved using an **in-order traversal** approach. Since we need to return elements in ascending order (smallest to largest), we can use a stack to simulate the in-order traversal.

### Key Insight
In-order traversal of a BST yields elements in sorted order. We can implement this iteratively using a stack rather than recursively to achieve O(h) memory complexity (where h is the height of the tree).

## Algorithm Design

### Data Structures
- **Stack**: Stores nodes along the leftmost path of the BST
- **TreeNode**: Standard binary tree node with Val, Left, and Right pointers

### Constructor (`BSTIterator`)
1. Initialize an empty stack
2. Push all left children starting from the root (this gives us the smallest element at the top of the stack)

### `hasNext()`
- Return `true` if the stack is not empty (there are more elements to process)

### `next()`
1. Pop the top node from the stack (this is the next smallest element)
2. If the popped node has a right child, push all left children of that right child onto the stack
3. Return the value of the popped node

## Complexity Analysis

### Time Complexity
- **Constructor**: O(h) where h is the height of the tree (we traverse the leftmost path)
- **`next()`**: Amortized O(1) - Each node is pushed and popped exactly once
- **`hasNext()`**: O(1) - Simple stack check

### Space Complexity
- **Overall**: O(h) where h is the height of the tree (stack stores at most h nodes)

## Implementation Details

### Go Implementation
```go
type BSTIterator struct {
    stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    iterator := BSTIterator{}
    iterator.pushLeft(root)
    return iterator
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
    for node != nil {
        this.stack = append(this.stack, node)
        node = node.Left
    }
}

func (this *BSTIterator) Next() int {
    node := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    this.pushLeft(node.Right)
    return node.Val
}

func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0
}
```

### Key Points
1. **Amortized O(1) time**: Each node is visited exactly once across all `next()` calls
2. **O(h) memory**: Only stores the current path in the stack
3. **Lazy evaluation**: Nodes are processed on-demand rather than all at once

## Test Cases

### Example 1
```
Input: [7, 3, 15, null, null, 9, 20]
Operations:
- iterator.next()    → 3
- iterator.next()    → 7  
- iterator.hasNext() → true
- iterator.next()    → 9
- iterator.hasNext() → true
- iterator.next()    → 15
- iterator.hasNext() → true
- iterator.next()    → 20
- iterator.hasNext() → false
```

### Edge Cases
1. **Empty tree**: `hasNext()` returns false immediately
2. **Single node tree**: One `next()` call returns the value, then `hasNext()` returns false
3. **Left-skewed tree**: All nodes along the left path
4. **Right-skewed tree**: Still works correctly with O(n) memory in worst case

## Related Problems
- **0094 - Binary Tree Inorder Traversal**: Similar traversal pattern
- **0230 - Kth Smallest Element in BST**: Finding specific elements in sorted order
- **0285 - Inorder Successor in BST**: Finding next node in in-order traversal

## Learning Points
1. **Iterative tree traversal**: Converting recursive algorithms to iterative ones using stacks
2. **Amortized analysis**: Understanding why `next()` is O(1) amortized
3. **Lazy evaluation**: Processing elements on-demand rather than precomputing everything
4. **BST properties**: Leveraging the sorted nature of BSTs for efficient iteration

## Real-World Applications
- **Database indexing**: Iterating through B-tree indexes
- **File system traversal**: Navigating directory structures
- **Memory-efficient iteration**: When you can't store the entire traversal in memory
- **Stream processing**: Processing sorted data streams incrementally