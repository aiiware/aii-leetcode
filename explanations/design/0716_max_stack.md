# 716. Max Stack

## Problem Statement
Design a max stack that supports push, pop, top, peekMax, and popMax operations:
- `push(x)`: Push element x onto stack
- `pop()`: Remove the element on top of the stack and return it
- `top()`: Get the element on the top
- `peekMax()`: Retrieve the maximum element in the stack
- `popMax()`: Retrieve the maximum element and remove it

If there are multiple maximum elements, `popMax()` should remove the one closest to the top.

## Example
```
Input: ["MaxStack", "push", "push", "push", "top", "popMax", "top", "peekMax", "pop", "top"]
       [[], [5], [1], [5], [], [], [], [], [], []]
Output: [null, null, null, null, 5, 5, 1, 5, 1, 5]

Explanation:
- After pushes: [5, 1, 5]
- top() returns 5
- popMax() removes the top 5 (closest to top), returns 5
- Stack becomes: [5, 1]
- top() returns 1
- peekMax() returns 5
- pop() returns 1
- top() returns 5
```

## Solution Approach: Two Stacks + Max Heap

### Intuition
We need to support:
- Regular stack operations: O(1)
- Max operations: O(log n) or better
- `popMax()`: Remove max element from anywhere in stack

This suggests using:
- **Main Stack**: Regular stack operations
- **Max Stack**: Track current max at each position
- **Max Heap**: For efficient max operations with lazy deletion

### Data Structure Design
```
MaxStack {
    stack:    []int          // main stack
    maxStack: []int          // max at each position
    removed:  map[int]bool   // lazy deletion markers
    index:    int            // unique IDs for heap nodes
    heap:     *MaxStackHeap  // max heap
}

heapNode {
    value: int  // stack value
    index: int  // unique ID for tie-breaking
}
```

### Algorithm
1. **push(x)**:
   - Add to main stack
   - Update maxStack: max(x, current max)
   - Add to heap with unique index

2. **pop()**:
   - Remove from top of both stacks
   - Mark corresponding heap node as removed (lazy deletion)
   - Return popped value

3. **top()**: Return top of main stack

4. **peekMax()**:
   - Clean up removed nodes from heap
   - Return value at heap root

5. **popMax()**:
   - Get max from heap (with cleanup)
   - Find and remove from stack (may require temporary stack)
   - Rebuild maxStack for remaining elements

### Complexity Analysis
- **push, pop, top**: O(1)
- **peekMax**: O(log n) amortized (with lazy deletion)
- **popMax**: O(n) worst case (searching stack), O(log n) amortized for heap

### Code Implementation
```go
type MaxStack struct {
    stack    []int
    maxStack []int
    removed  map[int]bool
    index    int
    heap     *MaxStackHeap
}

func ConstructorMaxStack() MaxStack {
    h := &MaxStackHeap{}
    heap.Init(h)
    return MaxStack{
        stack:    make([]int, 0),
        maxStack: make([]int, 0),
        removed:  make(map[int]bool),
        index:    0,
        heap:     h,
    }
}

func (this *MaxStack) Push(x int) {
    this.stack = append(this.stack, x)
    
    // Update maxStack
    if len(this.maxStack) == 0 || x >= this.maxStack[len(this.maxStack)-1] {
        this.maxStack = append(this.maxStack, x)
    } else {
        this.maxStack = append(this.maxStack, this.maxStack[len(this.maxStack)-1])
    }
    
    // Add to heap
    heap.Push(this.heap, heapNode{value: x, index: this.index})
    this.index++
}

func (this *MaxStack) Pop() int {
    top := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    this.maxStack = this.maxStack[:len(this.maxStack)-1]
    
    // Lazy deletion
    this.removed[this.index-1] = true
    return top
}

func (this *MaxStack) PeekMax() int {
    // Clean up removed nodes
    for this.heap.Len() > 0 {
        node := (*this.heap)[0]
        if this.removed[node.index] {
            heap.Pop(this.heap)
            delete(this.removed, node.index)
        } else {
            return node.value
        }
    }
    return -1
}
```

### Alternative Approaches
1. **Double Linked List + TreeMap**: O(log n) all operations
2. **Two Stacks Only**: O(n) for popMax by searching
3. **Stack of Nodes**: Each node stores value and current max

### Key Insights
1. **Lazy Deletion**: Mark heap nodes as removed instead of deleting immediately
2. **Unique Indices**: Handle duplicate values in heap
3. **Max Stack Optimization**: O(1) peekMax without heap
4. **Tie-breaking**: For duplicate max values, remove most recent (closest to top)

### Edge Cases
1. **Empty Stack**: Return -1 or appropriate sentinel
2. **Duplicate Max Values**: popMax removes most recent
3. **All Elements Same**: All operations should work correctly
4. **Large Number of Operations**: Efficient lazy deletion crucial

### Optimization Details
1. **Amortized Complexity**: Lazy deletion makes peekMax O(log n) amortized
2. **Memory Efficiency**: Only store necessary metadata
3. **Concurrent Operations**: Design handles interleaved push/pop/popMax

### Real-World Applications
- Undo/redo systems with max operation tracking
- Game score tracking with max score retrieval
- Transaction systems needing max value operations
- Monitoring systems tracking peak values

### Related Problems
- 155. Min Stack (simpler, no popMin)
- 895. Maximum Frequency Stack
- 239. Sliding Window Maximum
- 496. Next Greater Element I