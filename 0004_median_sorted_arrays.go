package leetcode

import (
	"math"
)

// FindMedianSortedArrays solves LeetCode problem 0004: Median of Two Sorted Arrays
// Difficulty: Hard
// Tags: Array, Binary Search, Divide and Conquer
//
// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log(min(m, n))).
//
// Time complexity: O(log(min(m, n)))
// Space complexity: O(1)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array for binary search optimization
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := total / 2

	// Binary search on the smaller array
	left, right := 0, m
	for left <= right {
		// Partition nums1
		mid1 := (left + right) / 2
		// Partition nums2 (half - mid1 elements from nums2)
		mid2 := half - mid1

		// Get the four boundary values
		var left1, right1, left2, right2 int

		if mid1 > 0 {
			left1 = nums1[mid1-1]
		} else {
			left1 = math.MinInt32
		}

		if mid1 < m {
			right1 = nums1[mid1]
		} else {
			right1 = math.MaxInt32
		}

		if mid2 > 0 {
			left2 = nums2[mid2-1]
		} else {
			left2 = math.MinInt32
		}

		if mid2 < n {
			right2 = nums2[mid2]
		} else {
			right2 = math.MaxInt32
		}

		// Check if partition is correct
		if left1 <= right2 && left2 <= right1 {
			// Found correct partition
			if total%2 == 0 {
				// Even total length: median is average of max(left) and min(right)
				maxLeft := maxInt(left1, left2)
				minRight := minInt(right1, right2)
				return float64(maxLeft+minRight) / 2.0
			} else {
				// Odd total length: median is min(right)
				return float64(minInt(right1, right2))
			}
		} else if left1 > right2 {
			// Too far right in nums1, move left
			right = mid1 - 1
		} else {
			// Too far left in nums1, move right
			left = mid1 + 1
		}
	}

	return 0.0
}

// maxInt returns the maximum of two integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// minInt returns the minimum of two integers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}