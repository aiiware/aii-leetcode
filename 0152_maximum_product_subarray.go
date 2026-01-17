package leetcode

// 0152 - Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/
// Difficulty: Medium
// Topics: Array, Dynamic Programming, Kadane's Algorithm
// Companies: Google, Facebook, Amazon, Microsoft, Apple, Bloomberg, Uber, Oracle, Adobe, TikTok

/*
Description:
Given an integer array nums, find a subarray that has the largest product, and return the product.

A subarray is a contiguous non-empty sequence of elements within an array.
The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints:
- 1 <= nums.length <= 2 * 10^4
- -10 <= nums[i] <= 10
- The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

Follow-up: Can you solve it in O(n) time and O(1) space?
*/

// MaxProductKadane implements the solution using Kadane's algorithm adapted for products
// Time: O(n), Space: O(1)
func MaxProductKadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize with first element
	result := nums[0]
	curMax, curMin := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Store current max before updating (needed for min calculation)
		tempMax := curMax

		// Update current max: max of (num, num * curMax, num * curMin)
		curMax = Max(num, Max(num*curMax, num*curMin))

		// Update current min: min of (num, num * tempMax, num * curMin)
		curMin = Min(num, Min(num*tempMax, num*curMin))

		// Update global result
		result = Max(result, curMax)
	}

	return result
}

// MaxProductPrefixSuffix implements the solution using prefix and suffix products
// Time: O(n), Space: O(1)
func MaxProductPrefixSuffix(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	result := nums[0]
	prefix, suffix := 0, 0

	for i := 0; i < n; i++ {
		// Reset to 1 if we encounter 0 (not to 0, because we want to start fresh)
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}

		// Update prefix product (left to right)
		prefix *= nums[i]

		// Update suffix product (right to left)
		suffix *= nums[n-1-i]

		// Update result with max of current result, prefix, and suffix
		result = Max(result, Max(prefix, suffix))
	}

	return result
}

// MaxProductBruteForce implements the brute force solution
// Time: O(n^2), Space: O(1)
func MaxProductBruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for i := 0; i < len(nums); i++ {
		product := 1
		for j := i; j < len(nums); j++ {
			product *= nums[j]
			result = Max(result, product)
		}
	}

	return result
}

// MaxProductDP implements a dynamic programming solution with explicit DP arrays
// Time: O(n), Space: O(n)
func MaxProductDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	// dpMax[i] = maximum product ending at position i
	// dpMin[i] = minimum product ending at position i
	dpMax := make([]int, n)
	dpMin := make([]int, n)

	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	result := nums[0]

	for i := 1; i < n; i++ {
		num := nums[i]

		// Three possibilities for dpMax[i]:
		// 1. Start new subarray with nums[i]
		// 2. Extend previous max product
		// 3. Previous min product * nums[i] (negative * negative = positive)
		dpMax[i] = Max(num, Max(num*dpMax[i-1], num*dpMin[i-1]))

		// Three possibilities for dpMin[i]:
		// 1. Start new subarray with nums[i]
		// 2. Extend previous min product
		// 3. Previous max product * nums[i] (positive * negative = negative)
		dpMin[i] = Min(num, Min(num*dpMax[i-1], num*dpMin[i-1]))

		result = Max(result, dpMax[i])
	}

	return result
}

// MaxProduct is the main function that uses the optimal solution (Kadane's algorithm)
func MaxProduct(nums []int) int {
	return MaxProductKadane(nums)
}