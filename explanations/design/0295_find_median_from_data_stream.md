# 295. Find Median from Data Stream

## Problem Statement
Design a data structure that supports adding numbers from a data stream and finding the median of all numbers added so far.

The median is the middle value in an ordered list:
- If the list has odd length: median is the middle element
- If the list has even length: median is the average of the two middle elements

## Example
```
Input: ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
       [[], [1], [2], [], [3], []]
Output: [null, null, null, 1.5, null, 2.0]

Explanation:
- After addNum(1): [1] → median = 1
- After addNum(2): [1, 2] → median = (1 + 2) / 2 = 1.5
- After addNum(3): [1, 2, 3] → median = 2
```

## Solution Approach: Two Heaps

### Intuition
Maintaining a sorted list for each insertion would be O(n log n). We can do better using two heaps:
- **Max Heap**: Stores the smaller half of numbers (max at root)
- **Min Heap**: Stores the larger half of numbers (min at root)

This allows us to:
- Add numbers in O(log n) time
- Find median in O(1) time

### Algorithm
1. **Data Structure**:
   - `maxHeap`: Max heap containing the smaller half
   - `minHeap`: Min heap containing the larger half

2. **Invariants**:
   - All elements in `maxHeap` ≤ all elements in `minHeap`
   - `maxHeap` can have at most one more element than `minHeap`

3. **AddNum(num)**:
   - Always add to `maxHeap` first
   - Move the largest element from `maxHeap` to `minHeap` to maintain ordering
   - Balance sizes: if `minHeap` has more elements, move one back to `maxHeap`

4. **FindMedian()**:
   - If `maxHeap` has more elements: return root of `maxHeap`
   - If equal sizes: return average of both roots

### Complexity Analysis
- **Time Complexity**:
  - `AddNum`: O(log n) for heap operations
  - `FindMedian`: O(1) for accessing heap roots
- **Space Complexity**: O(n) to store all numbers

### Code Implementation
```go
type MedianFinder struct {
    maxHeap *MaxHeap // stores smaller half
    minHeap *MinHeap // stores larger half
}

func Constructor() MedianFinder {
    maxHeap := &MaxHeap{}
    minHeap := &MinHeap{}
    heap.Init(maxHeap)
    heap.Init(minHeap)
    
    return MedianFinder{maxHeap, minHeap}
}

func (this *MedianFinder) AddNum(num int) {
    heap.Push(this.maxHeap, num)
    heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
    
    if this.minHeap.Len() > this.maxHeap.Len() {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() > this.minHeap.Len() {
        return float64((*this.maxHeap)[0])
    }
    return float64((*this.maxHeap)[0] + (*this.minHeap)[0]) / 2.0
}
```

### Alternative Implementation
The solution also includes an alternative implementation with different balancing logic that's more explicit about the invariants.

### Key Insights
1. **Heap Choice**: Max heap for smaller half, min heap for larger half
2. **Balancing**: Maintain size difference ≤ 1
3. **Ordering**: Ensure all elements in max heap ≤ all elements in min heap
4. **Edge Cases**: Handle empty streams, single elements, duplicates

### Common Pitfalls
1. Forgetting to initialize heaps with `heap.Init()`
2. Incorrect size balancing logic
3. Not handling the case when both heaps are empty
4. Integer division when calculating average (use float64)

### Real-World Applications
- Real-time analytics for streaming data
- Financial systems tracking median prices
- Network monitoring for median latency
- Gaming systems for player skill ratings

### Related Problems
- 480. Sliding Window Median
- 4. Median of Two Sorted Arrays
- 703. Kth Largest Element in a Stream