package leetcode

import "fmt"

/*
163. Missing Ranges

You are given an inclusive range [lower, upper] and a sorted unique integer array nums,
where all elements are within the inclusive range.

Return the smallest sorted list of ranges that cover every missing number exactly.
That is, no element of nums is in any of the ranges, and each missing number is in one of the ranges.

Each range [a,b] in the list should be output as:
- "a->b" if a != b
- "a" if a == b

Example 1:
Input: nums = [0,1,3,50,75], lower = 0, upper = 99
Output: ["2","4->49","51->74","76->99"]
Explanation: The ranges are:
[2,2] --> "2"
[4,49] --> "4->49"
[51,74] --> "51->74"
[76,99] --> "76->99"

Example 2:
Input: nums = [-1], lower = -1, upper = -1
Output: []
Explanation: There are no missing numbers since all numbers in the range are covered.

Constraints:
- -10^9 <= lower <= upper <= 10^9
- 0 <= nums.length <= 100
- lower <= nums[i] <= upper
- All the values of nums are unique.
- nums is sorted in ascending order.

Difficulty: Medium
Tags: Array
Companies: Amazon, Facebook, Google, Oracle
*/

func findMissingRanges(nums []int, lower int, upper int) []string {
	result := []string{}
	
	// Helper function to add range to result
	addRange := func(start, end int) {
		if start > end {
			return
		}
		if start == end {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}
	}
	
	// Handle empty array case
	if len(nums) == 0 {
		addRange(lower, upper)
		return result
	}
	
	// Check for missing numbers before the first element
	if nums[0] > lower {
		addRange(lower, nums[0]-1)
	}
	
	// Check for missing numbers between elements
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1]+1 {
			addRange(nums[i-1]+1, nums[i]-1)
		}
	}
	
	// Check for missing numbers after the last element
	if nums[len(nums)-1] < upper {
		addRange(nums[len(nums)-1]+1, upper)
	}
	
	return result
}