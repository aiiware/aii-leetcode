package leetcode

import (
	"sort"
)

// FourSum solves LeetCode problem 0018: 4Sum
// Difficulty: Medium
// Tags: Array, Two Pointers, Sorting
//
// Given an array nums of n integers, return all unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
// - 0 <= a, b, c, d < n
// - a, b, c, and d are distinct
// - nums[a] + nums[b] + nums[c] + nums[d] == target
//
// You may return the answer in any order.
//
// Time complexity: O(n³), Space complexity: O(1) (not counting output)
func FourSum(nums []int, target int) [][]int {
	result := [][]int{}

	if len(nums) < 4 {
		return result
	}

	// Sort the array first
	sort.Ints(nums)

	// Outer loop for first number
	for i := 0; i < len(nums)-3; i++ {
		// Skip duplicate numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early termination: if minimum possible sum is too large
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		// Skip if maximum possible sum is too small
		if nums[i]+nums[len(nums)-3]+nums[len(nums)-2]+nums[len(nums)-1] < target {
			continue
		}

		// Inner loop for second number
		for j := i + 1; j < len(nums)-2; j++ {
			// Skip duplicate numbers
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Early termination
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			// Skip if maximum possible sum is too small
			if nums[i]+nums[j]+nums[len(nums)-2]+nums[len(nums)-1] < target {
				continue
			}

			// Use two pointers for remaining two numbers
			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

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
	}

	return result
}
