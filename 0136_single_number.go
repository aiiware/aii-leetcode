package leetcode

/*
# 0136 - Single Number
## Problem Description
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

## Examples
Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1

## Constraints
- 1 <= nums.length <= 3 * 10^4
- -3 * 10^4 <= nums[i] <= 3 * 10^4
- Each element in the array appears twice except for one element which appears only once.

## Solution Approach
This problem can be solved using XOR bitwise operation:
- XOR of a number with itself is 0
- XOR of a number with 0 is the number itself
- XOR is commutative and associative

So if we XOR all numbers in the array:
- Numbers appearing twice will cancel each other (a XOR a = 0)
- The single number will remain (0 XOR b = b)

Time Complexity: O(N) where N is the length of nums
Space Complexity: O(1)
*/

// SingleNumber returns the element that appears only once
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}