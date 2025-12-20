package leetcode

import (
	"sort"
)

// ThreeSum solves LeetCode problem 0015: 3Sum
// Difficulty: Medium
// Tags: Array, Two Pointers, Sorting
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
// Time complexity: O(n²), Space complexity: O(1) (not counting output)
func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	if len(nums) < 3 {
		return result
	}

	// Sort the array first
	sort.Ints(nums)

	// Iterate through array, using two pointers for remaining two numbers
	for i := 0; i < len(nums)-2; i++ {
		// If current number is positive, no triplet can sum to zero
		if nums[i] > 0 {
			break
		}

		// Skip duplicate numbers to avoid duplicate triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two pointers to find pairs that sum to -nums[i]
		left := i + 1
		right := len(nums) - 1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate numbers
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
