package design

/*
232. Implement Queue using Stacks

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue 
should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:
- void push(int x) Pushes element x to the back of the queue.
- int pop() Removes the element from the front of the queue and returns it.
- int peek() Returns the element at the front of the queue.
- boolean empty() Returns true if the queue is empty, false otherwise.

Notes:
- You must use only standard operations of a stack, which means only push to top, 
  peek/pop from top, size, and is empty operations are valid.
- Depending on your language, the stack may not be supported natively. You may simulate 
  a stack using a list or deque (double-ended queue) as long as you use only a stack's 
  standard operations.

Example 1:
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek();  // return 1
myQueue.pop();   // return 1, queue is [2]
myQueue.empty(); // return false

Constraints:
- 1 <= x <= 9
- At most 100 calls will be made to push, pop, peek, and empty.
- All the calls to pop and peek are valid.
*/

/*
Difficulty: Easy
Tags: Stack, Design, Queue
Companies: Amazon, Facebook, Google, Microsoft, Bloomberg, Apple
*/

// MyQueue implements a queue using two stacks
type MyQueue struct {
    input  []int // stack for push operations
    output []int // stack for pop/peek operations
}

func ConstructorMyQueue() MyQueue {
    return MyQueue{
        input:  make([]int, 0),
        output: make([]int, 0),
    }
}

func (this *MyQueue) Push(x int) {
    this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
    if this.Empty() {
        return -1
    }
    
    // If output stack is empty, transfer all elements from input stack
    if len(this.output) == 0 {
        for len(this.input) > 0 {
            // Pop from input
            lastIdx := len(this.input) - 1
            val := this.input[lastIdx]
            this.input = this.input[:lastIdx]
            
            // Push to output
            this.output = append(this.output, val)
        }
    }
    
    // Pop from output
    lastIdx := len(this.output) - 1
    val := this.output[lastIdx]
    this.output = this.output[:lastIdx]
    return val
}

func (this *MyQueue) Peek() int {
    if this.Empty() {
        return -1
    }
    
    // If output stack is empty, transfer all elements from input stack
    if len(this.output) == 0 {
        for len(this.input) > 0 {
            // Pop from input
            lastIdx := len(this.input) - 1
            val := this.input[lastIdx]
            this.input = this.input[:lastIdx]
            
            // Push to output
            this.output = append(this.output, val)
        }
    }
    
    // Peek from output
    lastIdx := len(this.output) - 1
    return this.output[lastIdx]
}

func (this *MyQueue) Empty() bool {
    return len(this.input) == 0 && len(this.output) == 0
}

// Alternative implementation with amortized O(1) operations
type MyQueueAmortized struct {
    pushStack []int
    popStack  []int
}

func ConstructorMyQueueAmortized() MyQueueAmortized {
    return MyQueueAmortized{
        pushStack: make([]int, 0),
        popStack:  make([]int, 0),
    }
}

func (this *MyQueueAmortized) Push(x int) {
    this.pushStack = append(this.pushStack, x)
}

func (this *MyQueueAmortized) transfer() {
    if len(this.popStack) == 0 {
        for len(this.pushStack) > 0 {
            // Pop from pushStack
            lastIdx := len(this.pushStack) - 1
            val := this.pushStack[lastIdx]
            this.pushStack = this.pushStack[:lastIdx]
            
            // Push to popStack
            this.popStack = append(this.popStack, val)
        }
    }
}

func (this *MyQueueAmortized) Pop() int {
    this.transfer()
    if len(this.popStack) == 0 {
        return -1
    }
    
    lastIdx := len(this.popStack) - 1
    val := this.popStack[lastIdx]
    this.popStack = this.popStack[:lastIdx]
    return val
}

func (this *MyQueueAmortized) Peek() int {
    this.transfer()
    if len(this.popStack) == 0 {
        return -1
    }
    
    lastIdx := len(this.popStack) - 1
    return this.popStack[lastIdx]
}

func (this *MyQueueAmortized) Empty() bool {
    return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */