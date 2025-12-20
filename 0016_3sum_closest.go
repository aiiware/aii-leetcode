package leetcode

import (
	"sort"
)

// ThreeSumClosest solves LeetCode problem 0016: 3Sum Closest
// Difficulty: Medium
// Tags: Array, Two Pointers, Sorting
//
// Given an array nums of n integers and an integer target, find three integers
// in nums such that the sum is closest to target.
//
// Return the sum of the three integers.
//
// You may assume that each input would have exactly one solution.
//
// Time complexity: O(n²), Space complexity: O(1)
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	// Iterate through array using two pointers for remaining two numbers
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			diff := abs(currentSum - target)

			// Update closest sum if current sum is closer to target
			if diff < minDiff {
				minDiff = diff
				closestSum = currentSum
			}

			// If we found exact match, return immediately
			if currentSum == target {
				return currentSum
			}

			// Move pointers based on comparison with target
			if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}

// Helper function to get absolute value
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
