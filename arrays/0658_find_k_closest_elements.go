package arrays

import "sort"

// FindClosestElements solves LeetCode problem 658: Find K Closest Elements
// Difficulty: Medium
// Tags: Array, Two Pointers, Binary Search, Sliding Window, Sorting
//
// Given a sorted integer array arr, two integers k and x, return the k closest
// integers to x in the array. The result should also be sorted in ascending order.
//
// An integer a is closer to x than an integer b if:
// |a - x| < |b - x|, or
// |a - x| == |b - x| and a < b
//
// Time complexity: O(log n + k), Space complexity: O(1) excluding result
func FindClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	if k >= n {
		// If k is greater than or equal to array length, return the whole array
		return arr
	}

	// Binary search to find the position where x would be inserted
	left := 0
	right := n - k // We only need to search up to n-k because window size is k

	// Binary search for the left boundary of the window
	for left < right {
		mid := left + (right-left)/2

		// Compare distances from x to arr[mid] and arr[mid+k]
		// We want to find the starting point where the window is optimal
		if x-arr[mid] > arr[mid+k]-x {
			// If x is closer to arr[mid+k] than to arr[mid],
			// move window to the right
			left = mid + 1
		} else {
			// Otherwise, keep searching in the left half
			right = mid
		}
	}

	// Return k elements starting from left
	return arr[left : left+k]
}

// FindClosestElementsTwoPointers solves the same problem using two pointers
// This approach is more intuitive but has O(n) time complexity
func FindClosestElementsTwoPointers(arr []int, k int, x int) []int {
	n := len(arr)
	if k >= n {
		return arr
	}

	// Find the position where x would be inserted
	pos := sort.SearchInts(arr, x)

	// Initialize two pointers
	left := pos - 1
	right := pos

	// Expand window by comparing distances
	for right-left-1 < k {
		if left < 0 {
			// No more elements on the left, take from right
			right++
		} else if right >= n {
			// No more elements on the right, take from left
			left--
		} else {
			// Compare distances
			leftDist := absInt(x - arr[left])
			rightDist := absInt(x - arr[right])

			if leftDist <= rightDist {
				// Prefer left element if distances are equal (because a < b)
				left--
			} else {
				right++
			}
		}
	}

	// Return the window
	return arr[left+1 : right]
}

// FindClosestElementsSorting solves the problem by sorting based on distance
// This is the simplest approach but has O(n log n) time complexity
func FindClosestElementsSorting(arr []int, k int, x int) []int {
	// Create a copy to avoid modifying original
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	// Sort by distance to x, with tie-breaker by value
	sort.Slice(sorted, func(i, j int) bool {
		distI := absInt(sorted[i] - x)
		distJ := absInt(sorted[j] - x)

		if distI == distJ {
			return sorted[i] < sorted[j]
		}
		return distI < distJ
	})

	// Take first k elements
	result := sorted[:k]

	// Sort result in ascending order
	sort.Ints(result)

	return result
}