package arrays

import (
	"container/heap"
)

// Item represents an element in the min-heap for kthSmallest
type Item struct {
	value int
	row   int
	col   int
}

// kthSmallest solves LeetCode problem 0378: Kth Smallest Element in a Sorted Matrix
// Difficulty: Medium
// Tags: Array, Binary Search, Heap (Priority Queue), Matrix
//
// Given an n x n matrix where each of the rows and columns is sorted in ascending order,
// return the kth smallest element in the matrix.
// Note that it is the kth smallest element in the sorted order, not the kth distinct element.
//
// Example 1:
// Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// Output: 13
// Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13.
//
// Example 2:
// Input: matrix = [[-5]], k = 1
// Output: -5
//
// Constraints:
// n == matrix.length == matrix[i].length
// 1 <= n <= 300
// -10^9 <= matrix[i][j] <= 10^9
// All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
// 1 <= k <= n^2
//
// Time complexity: O(k log n), Space complexity: O(n)
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	// Min-heap approach
	// We'll use a min-heap to always get the smallest element
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Push the first element of each row into the heap
	for i := 0; i < n && i < k; i++ {
		heap.Push(minHeap, Item{
			value: matrix[i][0],
			row:   i,
			col:   0,
		})
	}

	// Extract the smallest element k-1 times
	for i := 0; i < k-1; i++ {
		item := heap.Pop(minHeap).(Item)

		// If there's a next element in the same row, push it
		if item.col+1 < n {
			heap.Push(minHeap, Item{
				value: matrix[item.row][item.col+1],
				row:   item.row,
				col:   item.col + 1,
			})
		}
	}

	// The kth smallest is at the top of the heap
	return (*minHeap)[0].value
}

// MinHeap implements heap.Interface
type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// kthSmallestBinarySearch solves using binary search approach
// Time complexity: O(n log(max-min)), Space complexity: O(1)
func kthSmallestBinarySearch(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	// Binary search on value range
	left, right := matrix[0][0], matrix[n-1][n-1]

	for left < right {
		mid := left + (right-left)/2

		// Count how many elements are <= mid
		count := 0
		col := n - 1

		// Start from top-right corner
		for row := 0; row < n; row++ {
			// Move left until we find element <= mid
			for col >= 0 && matrix[row][col] > mid {
				col--
			}
			count += (col + 1)
		}

		if count < k {
			// Need more elements, search in right half
			left = mid + 1
		} else {
			// Too many elements, search in left half
			right = mid
		}
	}

	return left
}
