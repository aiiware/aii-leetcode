package sorting

/*
215. Kth Largest Element in an Array

Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

Constraints:
- 1 <= k <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4
*/

/*
Difficulty: Medium
Tags: Array, Divide and Conquer, Sorting, Heap (Priority Queue), Quickselect
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

import (
    "container/heap"
    "math/rand"
)

// findKthLargestQuickSelect uses Quickselect algorithm (Hoare's selection algorithm)
func findKthLargestQuickSelect(nums []int, k int) int {
    // Convert kth largest to kth smallest (0-indexed)
    k = len(nums) - k
    
    // Helper function for quickselect
    var quickselect func(left, right int) int
    quickselect = func(left, right int) int {
        if left == right {
            return nums[left]
        }
        
        // Random pivot for better average performance
        pivotIndex := left + rand.Intn(right-left+1)
        pivotValue := nums[pivotIndex]
        
        // Move pivot to the end
        nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
        
        // Partition
        storeIndex := left
        for i := left; i < right; i++ {
            if nums[i] < pivotValue {
                nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
                storeIndex++
            }
        }
        
        // Move pivot to its final place
        nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
        
        if storeIndex == k {
            return nums[storeIndex]
        } else if storeIndex < k {
            return quickselect(storeIndex+1, right)
        } else {
            return quickselect(left, storeIndex-1)
        }
    }
    
    return quickselect(0, len(nums)-1)
}

// MinHeap for heap solution
type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeapInt) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeapInt) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

// findKthLargestHeap uses a min-heap of size k
func findKthLargestHeap(nums []int, k int) int {
    h := &MinHeapInt{}
    heap.Init(h)
    
    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    return (*h)[0]
}

// findKthLargestSort uses sorting (simplest but O(n log n))
func findKthLargestSort(nums []int, k int) int {
    // Sort in descending order
    sortDesc := func(nums []int) {
        sort := func(arr []int) {
            for i := 0; i < len(arr); i++ {
                for j := i + 1; j < len(arr); j++ {
                    if arr[i] < arr[j] {
                        arr[i], arr[j] = arr[j], arr[i]
                    }
                }
            }
        }
        sort(nums)
    }
    
    sortDesc(nums)
    return nums[k-1]
}

// Counting sort version for limited range
func findKthLargestCounting(nums []int, k int) int {
    // Since nums[i] range is [-10^4, 10^4]
    offset := 10000
    count := make([]int, 20001) // -10000 to 10000
    
    for _, num := range nums {
        count[num+offset]++
    }
    
    // Find kth largest by counting from the end
    for i := len(count) - 1; i >= 0; i-- {
        k -= count[i]
        if k <= 0 {
            return i - offset
        }
    }
    
    return 0
}