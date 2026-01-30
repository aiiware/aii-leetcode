package sorting

/*
179. Largest Number

Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

Example 1:
Input: nums = [10,2]
Output: "210"

Example 2:
Input: nums = [3,30,34,5,9]
Output: "9534330"

Example 3:
Input: nums = [0,0]
Output: "0"

Constraints:
- 1 <= nums.length <= 100
- 0 <= nums[i] <= 10^9

Difficulty: Medium
Tags: Sorting, Greedy
Companies: Amazon, Microsoft, Google, Facebook, Apple, Uber, Bloomberg, Adobe, Oracle, Twitter
*/

import (
	"sort"
	"strconv"
	"strings"
)

// largestNumber returns the largest possible number as a string by concatenating the given numbers
func largestNumber(nums []int) string {
	// Handle empty slice (though constraints say length >= 1)
	if len(nums) == 0 {
		return ""
	}

	// Convert integers to strings for custom sorting
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	// Custom sort: compare concatenations a+b vs b+a
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// Handle edge case where all numbers are zero
	if strs[0] == "0" {
		return "0"
	}

	// Join all strings to form the result
	return strings.Join(strs, "")
}