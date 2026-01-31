package design

/*
225. Implement Stack using Queues

Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack 
should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:
- void push(int x) Pushes element x to the top of the stack.
- int pop() Removes the element on the top of the stack and returns it.
- int top() Returns the element on the top of the stack.
- boolean empty() Returns true if the stack is empty, false otherwise.

Notes:
- You must use only standard operations of a queue, which means that only push to back, 
  peek/pop from front, size and is empty operations are valid.
- Depending on your language, the queue may not be supported natively. You may simulate 
  a queue using a list or deque (double-ended queue) as long as you use only a queue's 
  standard operations.

Example 1:
Input
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 2, 2, false]

Explanation
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // return 2
myStack.pop(); // return 2
myStack.empty(); // return False

Constraints:
- 1 <= x <= 9
- At most 100 calls will be made to push, pop, top, and empty.
- All the calls to pop and top are valid.
*/

/*
Difficulty: Easy
Tags: Stack, Design, Queue
Companies: Amazon, Facebook, Google, Microsoft, Bloomberg, Apple
*/

// MyStack implements a stack using two queues
type MyStack struct {
    q1 []int // main queue
    q2 []int // auxiliary queue
}

func ConstructorMyStack() MyStack {
    return MyStack{
        q1: make([]int, 0),
        q2: make([]int, 0),
    }
}

func (this *MyStack) Push(x int) {
    // Push to q2
    this.q2 = append(this.q2, x)
    
    // Move all elements from q1 to q2
    for len(this.q1) > 0 {
        this.q2 = append(this.q2, this.q1[0])
        this.q1 = this.q1[1:]
    }
    
    // Swap q1 and q2
    this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
    if this.Empty() {
        return -1
    }
    val := this.q1[0]
    this.q1 = this.q1[1:]
    return val
}

func (this *MyStack) Top() int {
    if this.Empty() {
        return -1
    }
    return this.q1[0]
}

func (this *MyStack) Empty() bool {
    return len(this.q1) == 0
}

// Alternative implementation using single queue
type MyStackSingleQueue struct {
    q []int
}

func ConstructorMyStackSingleQueue() MyStackSingleQueue {
    return MyStackSingleQueue{
        q: make([]int, 0),
    }
}

func (this *MyStackSingleQueue) Push(x int) {
    n := len(this.q)
    this.q = append(this.q, x)
    // Rotate the queue to make the new element at the front
    for i := 0; i < n; i++ {
        this.q = append(this.q, this.q[0])
        this.q = this.q[1:]
    }
}

func (this *MyStackSingleQueue) Pop() int {
    if this.Empty() {
        return -1
    }
    val := this.q[0]
    this.q = this.q[1:]
    return val
}

func (this *MyStackSingleQueue) Top() int {
    if this.Empty() {
        return -1
    }
    return this.q[0]
}

func (this *MyStackSingleQueue) Empty() bool {
    return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */