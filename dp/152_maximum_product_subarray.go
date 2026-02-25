package dp

import "leetcode/utils"

// 152. Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/
//
// Problem: Given an integer array nums, find a subarray that has the largest product, and return the product.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
// The test cases are generated so that the answer will fit in a 32-bit integer.
//
// Example 1:
// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
//
// Example 2:
// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
//
// Constraints:
// - 1 <= nums.length <= 2 * 10^4
// - -10 <= nums[i] <= 10
// - The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
//
// Approach: Dynamic Programming with state tracking
// We need to track both maximum and minimum products at each position because
// a negative number can turn a minimum product into a maximum product.
//
// Time Complexity: O(n)
// Space Complexity: O(1) with optimization, O(n) for explicit DP arrays

// MaxProductDPTabulation implements the DP solution with explicit tabulation
// This version uses O(n) space for clarity of the DP approach
func MaxProductDPTabulation(nums []int) int {
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
		// 2. Extend previous max product (nums[i] * dpMax[i-1])
		// 3. Previous min product * nums[i] (negative * negative = positive)
		dpMax[i] = utils.Max(num, utils.Max(num*dpMax[i-1], num*dpMin[i-1]))

		// Three possibilities for dpMin[i]:
		// 1. Start new subarray with nums[i]
		// 2. Extend previous min product (nums[i] * dpMin[i-1])
		// 3. Previous max product * nums[i] (positive * negative = negative)
		dpMin[i] = utils.Min(num, utils.Min(num*dpMax[i-1], num*dpMin[i-1]))

		result = utils.Max(result, dpMax[i])
	}

	return result
}

// MaxProductDPOptimized implements the DP solution with O(1) space
// This is the space-optimized version of the tabulation approach
func MaxProductDPOptimized(nums []int) int {
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
		curMax = utils.Max(num, utils.Max(num*curMax, num*curMin))

		// Update current min: min of (num, num * tempMax, num * curMin)
		curMin = utils.Min(num, utils.Min(num*tempMax, num*curMin))

		// Update global result
		result = utils.Max(result, curMax)
	}

	return result
}

// MaxProductDPRecursive implements the DP solution with memoization (top-down)
// This version uses recursion with memoization for educational purposes
func MaxProductDPRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	memoMax := make([]int, n)
	memoMin := make([]int, n)
	for i := range memoMax {
		memoMax[i] = -1
		memoMin[i] = -1
	}

	var dfs func(i int) (int, int)
	dfs = func(i int) (int, int) {
		if i == 0 {
			return nums[0], nums[0]
		}

		if memoMax[i] != -1 && memoMin[i] != -1 {
			return memoMax[i], memoMin[i]
		}

		prevMax, prevMin := dfs(i - 1)
		num := nums[i]

		// Calculate current max and min
		curMax := utils.Max(num, utils.Max(num*prevMax, num*prevMin))
		curMin := utils.Min(num, utils.Min(num*prevMax, num*prevMin))

		memoMax[i] = curMax
		memoMin[i] = curMin

		return curMax, curMin
	}

	// Compute results for all positions
	result := nums[0]
	for i := 0; i < n; i++ {
		curMax, _ := dfs(i)
		result = utils.Max(result, curMax)
	}

	return result
}

// MaxProduct is the main function that uses the optimal DP solution
func MaxProduct(nums []int) int {
	return MaxProductDPOptimized(nums)
}