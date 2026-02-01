# 0155 - Min Stack

## Problem Statement
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:
- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element val onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

**Example:**
```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]
```

## Solution Approaches

### Approach 1: Two Stacks
Maintain two stacks:
1. **Main stack:** Stores all elements
2. **Min stack:** Stores minimum values

```go
type MinStack struct {
    stack    []int
    minStack []int
}

func Constructor() MinStack {
    return MinStack{
        stack:    make([]int, 0),
        minStack: make([]int, 0),
    }
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    
    // Push to minStack if it's empty or val <= current min
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    if len(this.stack) == 0 {
        return
    }
    
    val := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    
    // Pop from minStack if it matches the popped value
    if val == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
```

### Approach 2: Single Stack with Pairs
Store pairs of (value, current_min) in a single stack.

```go
type MinStack struct {
    stack [][2]int  // [0] = value, [1] = min_so_far
}

func Constructor() MinStack {
    return MinStack{stack: make([][2]int, 0)}
}

func (this *MinStack) Push(val int) {
    minVal := val
    if len(this.stack) > 0 && this.stack[len(this.stack)-1][1] < minVal {
        minVal = this.stack[len(this.stack)-1][1]
    }
    this.stack = append(this.stack, [2]int{val, minVal})
}

func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1][1]
}
```

## Complexity Analysis

**Both Approaches:**
- **Time Complexity:** O(1) for all operations
- **Space Complexity:** O(n) where n is number of elements

## Key Insights

1. **Constant Time Minimum:** The challenge is maintaining minimum in O(1) time
2. **Duplicate Minimum Values:** Must handle cases where same minimum appears multiple times
3. **Pop Operation:** When popping, need to update min if popped value was current minimum

## Edge Cases

1. **Empty stack:** Calling pop/top/getMin on empty stack
2. **Duplicate values:** Multiple pushes of same minimum value
3. **Negative numbers:** Minimum could be negative
4. **Large numbers:** Integer overflow considerations

## Real-World Applications

1. **Undo/Redo Operations:** In text editors or graphic design software
2. **Transaction Systems:** Maintaining minimum balance
3. **Game Development:** Tracking minimum scores or resources
4. **Financial Systems:** Monitoring minimum stock prices

## Common Mistakes

1. **Not handling duplicates:** Forgetting to push duplicate mins to minStack
2. **Incorrect pop logic:** Not checking if popped value equals current min
3. **Empty stack access:** Not checking stack bounds
4. **Space inefficiency:** Storing redundant information

## Optimization Considerations

1. **Two stacks vs pairs:** 
   - Two stacks: Simpler logic, potentially less space if few duplicates
   - Pairs: Single data structure, easier to manage

2. **Memory optimization:** 
   - Store deltas instead of absolute values
   - Use bit manipulation for special cases

3. **Thread safety:** 
   - Add mutex locks for concurrent access
   - Consider lock-free implementations

## Related Problems
- 0716 - Max Stack (similar but tracks maximum)
- 0232 - Implement Queue using Stacks
- 0225 - Implement Stack using Queues
- 0146 - LRU Cache (another design problem)

## Learning Points
- Designing data structures with specific constraints
- Trade-offs between time and space complexity
- Handling edge cases in stack operations
- Multiple implementation approaches for same problem