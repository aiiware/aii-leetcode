package design

import "container/heap"

// MaxStack implements a stack that supports push, pop, top, peekMax, and popMax
type MaxStack struct {
	stack    []int          // main stack
	maxStack []int          // stack of max values
	removed  map[int]bool   // track removed elements for lazy deletion
	index    int            // current index for unique IDs
	heap     *MaxStackHeap  // max heap for popMax
}

type heapNode struct {
	value int
	index int // unique index for tie-breaking
}

// MaxStackHeap implements heap.Interface for max heap
type MaxStackHeap []heapNode

func (h MaxStackHeap) Len() int           { return len(h) }
func (h MaxStackHeap) Less(i, j int) bool { 
	if h[i].value == h[j].value {
		return h[i].index > h[j].index // larger index = more recent
	}
	return h[i].value > h[j].value 
}
func (h MaxStackHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxStackHeap) Push(x interface{}) {
	*h = append(*h, x.(heapNode))
}

func (h *MaxStackHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ConstructorMaxStack initializes a MaxStack
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

// Push pushes element x onto stack
func (this *MaxStack) Push(x int) {
	this.stack = append(this.stack, x)
	
	// Update maxStack
	if len(this.maxStack) == 0 || x >= this.maxStack[len(this.maxStack)-1] {
		this.maxStack = append(this.maxStack, x)
	} else {
		this.maxStack = append(this.maxStack, this.maxStack[len(this.maxStack)-1])
	}
	
	// Add to heap with unique index
	heap.Push(this.heap, heapNode{value: x, index: this.index})
	this.index++
}

// Pop removes the element on top of the stack and returns it
func (this *MaxStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	
	// Get top element
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.maxStack = this.maxStack[:len(this.maxStack)-1]
	
	// Mark as removed for lazy deletion
	this.removed[this.index-1] = true
	
	return top
}

// Top returns the element on the top of the stack
func (this *MaxStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

// PeekMax returns the maximum element in the stack
func (this *MaxStack) PeekMax() int {
	if len(this.stack) == 0 {
		return -1
	}
	
	// Clean up removed elements from heap
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

// PopMax retrieves the maximum element in the stack and removes it
func (this *MaxStack) PopMax() int {
	if len(this.stack) == 0 {
		return -1
	}
	
	// Get max from heap (with lazy deletion cleanup)
	for this.heap.Len() > 0 {
		node := heap.Pop(this.heap).(heapNode)
		if this.removed[node.index] {
			delete(this.removed, node.index)
			continue
		}
		
		// Found the max element, need to remove it from stack
		maxVal := node.value
		
		// Find and remove from stack
		tempStack := make([]int, 0)
		found := false
		
		// Pop elements until we find the max
		for len(this.stack) > 0 {
			top := this.stack[len(this.stack)-1]
			this.stack = this.stack[:len(this.stack)-1]
			this.maxStack = this.maxStack[:len(this.maxStack)-1]
			
			if top == maxVal && !found {
				found = true
				// Mark this instance as removed
				this.removed[node.index] = true
				break
			} else {
				tempStack = append(tempStack, top)
			}
		}
		
		// Push elements back onto stack
		for i := len(tempStack) - 1; i >= 0; i-- {
			val := tempStack[i]
			this.stack = append(this.stack, val)
			
			// Rebuild maxStack
			if len(this.maxStack) == 0 || val >= this.maxStack[len(this.maxStack)-1] {
				this.maxStack = append(this.maxStack, val)
			} else {
				this.maxStack = append(this.maxStack, this.maxStack[len(this.maxStack)-1])
			}
		}
		
		return maxVal
	}
	
	return -1
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := ConstructorMaxStack();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */