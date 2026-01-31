package design

/*
295. Find Median from Data Stream

The median is the middle value in an ordered integer list. If the size of the list is even, 
there is no middle value, and the median is the mean of the two middle values.

For example, for arr = [2,3,4], the median is 3.
For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.

Implement the MedianFinder class:
- MedianFinder() initializes the MedianFinder object.
- void addNum(int num) adds the integer num from the data stream to the data structure.
- double findMedian() returns the median of all elements so far. Answers within 10^-5 of 
  the actual answer will be accepted.

Example 1:
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0

Constraints:
- -10^5 <= num <= 10^5
- There will be at least one element in the data structure before calling findMedian.
- At most 5 * 10^4 calls will be made to addNum and findMedian.
*/

/*
Difficulty: Hard
Tags: Two Pointers, Design, Sorting, Heap (Priority Queue), Data Stream
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

import (
    "container/heap"
)

// MaxHeap implements a max heap using container/heap
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func (h *MaxHeap) Peek() int {
    return (*h)[0]
}

// MinHeap implements a min heap using container/heap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func (h *MinHeap) Peek() int {
    return (*h)[0]
}

// MedianFinder maintains two heaps to find median in O(log n) time
type MedianFinder struct {
    maxHeap *MaxHeap // stores the smaller half (max at root)
    minHeap *MinHeap // stores the larger half (min at root)
}

func ConstructorMedianFinder() MedianFinder {
    maxHeap := &MaxHeap{}
    minHeap := &MinHeap{}
    heap.Init(maxHeap)
    heap.Init(minHeap)
    
    return MedianFinder{
        maxHeap: maxHeap,
        minHeap: minHeap,
    }
}

func (this *MedianFinder) AddNum(num int) {
    // Add to maxHeap (smaller half) first
    heap.Push(this.maxHeap, num)
    
    // Balance: move the largest from maxHeap to minHeap
    heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
    
    // If minHeap has more elements, move one back to maxHeap
    if this.minHeap.Len() > this.maxHeap.Len() {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() > this.minHeap.Len() {
        // Odd number of elements, median is the root of maxHeap
        return float64(this.maxHeap.Peek())
    } else {
        // Even number of elements, median is average of two roots
        return float64(this.maxHeap.Peek()+this.minHeap.Peek()) / 2.0
    }
}

// Alternative implementation with simpler balancing logic
type MedianFinderAlt struct {
    maxHeap *MaxHeap
    minHeap *MinHeap
}

func ConstructorMedianFinderAlt() MedianFinderAlt {
    maxHeap := &MaxHeap{}
    minHeap := &MinHeap{}
    heap.Init(maxHeap)
    heap.Init(minHeap)
    
    return MedianFinderAlt{
        maxHeap: maxHeap,
        minHeap: minHeap,
    }
}

func (this *MedianFinderAlt) AddNum(num int) {
    // Always add to maxHeap first
    heap.Push(this.maxHeap, num)
    
    // Ensure all elements in maxHeap <= all elements in minHeap
    if this.maxHeap.Len() > 0 && this.minHeap.Len() > 0 && 
       this.maxHeap.Peek() > this.minHeap.Peek() {
        heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
    }
    
    // Balance sizes: maxHeap should have at most one more element than minHeap
    if this.maxHeap.Len() > this.minHeap.Len()+1 {
        heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
    } else if this.minHeap.Len() > this.maxHeap.Len() {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
    }
}

func (this *MedianFinderAlt) FindMedian() float64 {
    if this.maxHeap.Len() == this.minHeap.Len() {
        return float64(this.maxHeap.Peek()+this.minHeap.Peek()) / 2.0
    }
    return float64(this.maxHeap.Peek())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */