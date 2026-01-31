package sorting

/*
347. Top K Frequent Elements

Given an integer array nums and an integer k, return the k most frequent elements. 
You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

Constraints:
- 1 <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4
- k is in the range [1, the number of unique elements in the array].
- It is guaranteed that the answer is unique.
*/

/*
Difficulty: Medium
Tags: Array, Hash Table, Divide and Conquer, Sorting, Heap (Priority Queue), Bucket Sort, Counting, Quickselect
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

import (
    "container/heap"
    "math/rand"
)

// Element represents a number with its frequency
type Element struct {
    num   int
    count int
}

// topKFrequentHeap uses a min-heap to track top k frequent elements
func topKFrequentHeap(nums []int, k int) []int {
    // Count frequencies
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }
    
    // Define a min-heap based on frequency
    minHeap := &MinHeapFreq{}
    heap.Init(minHeap)
    
    // Push elements to heap, maintaining size k
    for num, count := range freqMap {
        heap.Push(minHeap, Element{num, count})
        if minHeap.Len() > k {
            heap.Pop(minHeap)
        }
    }
    
    // Extract results from heap
    result := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        result[i] = heap.Pop(minHeap).(Element).num
    }
    
    return result
}

// MinHeapFreq implements a min-heap based on frequency
type MinHeapFreq []Element

func (h MinHeapFreq) Len() int           { return len(h) }
func (h MinHeapFreq) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeapFreq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapFreq) Push(x interface{}) {
    *h = append(*h, x.(Element))
}

func (h *MinHeapFreq) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// topKFrequentBucketSort uses bucket sort for O(n) time
func topKFrequentBucketSort(nums []int, k int) []int {
    // Count frequencies
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }
    
    // Create buckets: index = frequency, value = list of numbers with that frequency
    maxFreq := 0
    for _, count := range freqMap {
        if count > maxFreq {
            maxFreq = count
        }
    }
    
    buckets := make([][]int, maxFreq+1)
    for num, count := range freqMap {
        buckets[count] = append(buckets[count], num)
    }
    
    // Collect top k frequent elements from buckets
    result := make([]int, 0, k)
    for i := maxFreq; i >= 0 && len(result) < k; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)
        }
    }
    
    // If we have more than k (due to multiple numbers with same frequency), trim
    if len(result) > k {
        result = result[:k]
    }
    
    return result
}

// topKFrequentQuickSelect uses quickselect on frequencies
func topKFrequentQuickSelect(nums []int, k int) []int {
    // Count frequencies
    freqMap := make(map[int]int)
    for _, num := range nums {
        freqMap[num]++
    }
    
    // Convert to slice of elements
    elements := make([]Element, 0, len(freqMap))
    for num, count := range freqMap {
        elements = append(elements, Element{num, count})
    }
    
    // Quickselect to find the kth largest frequency
    targetIdx := len(elements) - k
    
    var quickselect func(left, right int)
    quickselect = func(left, right int) {
        if left >= right {
            return
        }
        
        // Random pivot
        pivotIdx := left + rand.Intn(right-left+1)
        elements[pivotIdx], elements[right] = elements[right], elements[pivotIdx]
        
        pivot := elements[right].count
        storeIdx := left
        
        for i := left; i < right; i++ {
            if elements[i].count < pivot {
                elements[storeIdx], elements[i] = elements[i], elements[storeIdx]
                storeIdx++
            }
        }
        
        elements[storeIdx], elements[right] = elements[right], elements[storeIdx]
        
        if storeIdx == targetIdx {
            return
        } else if storeIdx < targetIdx {
            quickselect(storeIdx+1, right)
        } else {
            quickselect(left, storeIdx-1)
        }
    }
    
    quickselect(0, len(elements)-1)
    
    // Collect top k elements
    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = elements[targetIdx+i].num
    }
    
    return result
}