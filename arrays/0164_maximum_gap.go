package arrays

import (
	"math"
)

/*
164. Maximum Gap

Given an integer array nums, return the maximum difference between two successive elements
in its sorted form. If the array contains less than two elements, return 0.

You must write an algorithm that runs in linear time and uses linear extra space.

Example 1:
Input: nums = [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.

Example 2:
Input: nums = [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.

Constraints:
- 1 <= nums.length <= 10^5
- 0 <= nums[i] <= 10^9

Difficulty: Hard
Tags: Array, Bucket Sort, Radix Sort, Sorting
Companies: Amazon
*/

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	
	// Find min and max values
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	
	// If all elements are the same
	if minVal == maxVal {
		return 0
	}
	
	// Calculate bucket size and number of buckets
	bucketSize := maxIntGap(1, (maxVal-minVal)/(n-1))
	bucketCount := (maxVal-minVal)/bucketSize + 1
	
	// Initialize buckets
	bucketMin := make([]int, bucketCount)
	bucketMax := make([]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		bucketMin[i] = math.MaxInt32
		bucketMax[i] = math.MinInt32
	}
	
	// Put numbers into buckets
	for _, num := range nums {
		bucketIdx := (num - minVal) / bucketSize
		if num < bucketMin[bucketIdx] {
			bucketMin[bucketIdx] = num
		}
		if num > bucketMax[bucketIdx] {
			bucketMax[bucketIdx] = num
		}
	}
	
	// Calculate maximum gap
	maxGap := 0
	prevMax := minVal
	for i := 0; i < bucketCount; i++ {
		// Skip empty buckets
		if bucketMin[i] == math.MaxInt32 {
			continue
		}
		// Calculate gap between current bucket min and previous bucket max
		maxGap = maxIntGap(maxGap, bucketMin[i]-prevMax)
		prevMax = bucketMax[i]
	}
	
	return maxGap
}

func maxIntGap(a, b int) int {
	if a > b {
		return a
	}
	return b
}