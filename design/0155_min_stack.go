package design

/*
155. Min Stack
https://leetcode.com/problems/min-stack/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:
- MinStack() initializes the stack object.
- void push(int val) pushes the element val onto the stack.
- void pop() removes the element on the top of the stack.
- int top() gets the top element of the stack.
- int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.

Example 1:
Input:
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output:
[null,null,null,null,-3,null,0,-2]

Explanation:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2

Constraints:
- -2^31 <= val <= 2^31 - 1
- Methods pop, top and getMin operations will always be called on non-empty stacks.
- At most 3 * 10^4 calls will be made to push, pop, top, and getMin.

Difficulty: Easy
Tags: Stack, Design
Companies: Amazon, Microsoft, Google, Bloomberg, Adobe, Apple, Uber, Oracle
*/

// MinStack is a stack that supports push, pop, top, and retrieving the minimum element in constant time.
type MinStack struct {
	stack    []int // Main stack for values
	minStack []int // Auxiliary stack for minimum values
}

// MinStackConstructor initializes your data structure.
func MinStackConstructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

// Push pushes element val onto the stack.
// Time Complexity: O(1)
func (this *MinStack) Push(val int) {
	// Push to main stack
	this.stack = append(this.stack, val)

	// Push to minStack: if minStack is empty or val <= current min, push val
	// Otherwise push current min again
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

// Pop removes the element on the top of the stack.
// Time Complexity: O(1)
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	// Pop from both stacks
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

// Top gets the top element of the stack.
// Time Complexity: O(1)
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0 // According to problem constraints, stack won't be empty when Top() is called
	}
	return this.stack[len(this.stack)-1]
}

// GetMin retrieves the minimum element in the stack.
// Time Complexity: O(1)
func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return 0 // According to problem constraints, stack won't be empty when GetMin() is called
	}
	return this.minStack[len(this.minStack)-1]
}

// MinStackSingleStack is an alternative implementation using a single stack with tuples
// Each element in the stack is a pair: (value, current_min)
type MinStackSingleStack struct {
	stack [][2]int // Each element: [0] = value, [1] = min at this point
}

// MinStackSingleStackConstructor initializes the single-stack version.
func MinStackSingleStackConstructor() MinStackSingleStack {
	return MinStackSingleStack{
		stack: make([][2]int, 0),
	}
}

// Push pushes element val onto the stack.
func (this *MinStackSingleStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, [2]int{val, val})
	} else {
		currentMin := this.stack[len(this.stack)-1][1]
		if val < currentMin {
			this.stack = append(this.stack, [2]int{val, val})
		} else {
			this.stack = append(this.stack, [2]int{val, currentMin})
		}
	}
}

// Pop removes the element on the top of the stack.
func (this *MinStackSingleStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
}

// Top gets the top element of the stack.
func (this *MinStackSingleStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1][0]
}

// GetMin retrieves the minimum element in the stack.
func (this *MinStackSingleStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1][1]
}

// MinStackLinkedList is a linked list implementation
type MinStackNode struct {
	val    int
	min    int
	next   *MinStackNode
}

type MinStackLinkedList struct {
	head *MinStackNode
}

// MinStackLinkedListConstructor initializes the linked list version.
func MinStackLinkedListConstructor() MinStackLinkedList {
	return MinStackLinkedList{
		head: nil,
	}
}

// Push pushes element val onto the stack.
func (this *MinStackLinkedList) Push(val int) {
	if this.head == nil {
		this.head = &MinStackNode{
			val:  val,
			min:  val,
			next: nil,
		}
	} else {
		currentMin := this.head.min
		if val < currentMin {
			currentMin = val
		}
		newNode := &MinStackNode{
			val:  val,
			min:  currentMin,
			next: this.head,
		}
		this.head = newNode
	}
}

// Pop removes the element on the top of the stack.
func (this *MinStackLinkedList) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.next
}

// Top gets the top element of the stack.
func (this *MinStackLinkedList) Top() int {
	if this.head == nil {
		return 0
	}
	return this.head.val
}

// GetMin retrieves the minimum element in the stack.
func (this *MinStackLinkedList) GetMin() int {
	if this.head == nil {
		return 0
	}
	return this.head.min
}

// MinStackOptimized is an optimized version that uses less space
// It only pushes to minStack when the new value is <= current min
type MinStackOptimized struct {
	stack    []int
	minStack []int
}

// MinStackOptimizedConstructor initializes the optimized version.
func MinStackOptimizedConstructor() MinStackOptimized {
	return MinStackOptimized{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

// Push pushes element val onto the stack.
func (this *MinStackOptimized) Push(val int) {
	this.stack = append(this.stack, val)
	
	// Only push to minStack if it's empty or val <= current min
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

// Pop removes the element on the top of the stack.
func (this *MinStackOptimized) Pop() {
	if len(this.stack) == 0 {
		return
	}
	
	// Pop from main stack
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	
	// If the popped value equals the current min, pop from minStack too
	if val == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

// Top gets the top element of the stack.
func (this *MinStackOptimized) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

// GetMin retrieves the minimum element in the stack.
func (this *MinStackOptimized) GetMin() int {
	if len(this.minStack) == 0 {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}