package dp

// 416. Partition Equal Subset Sum
// https://leetcode.com/problems/partition-equal-subset-sum/
//
// Given an integer array nums, return true if you can partition the array into
// two subsets such that the sum of the elements in both subsets is equal or false otherwise.
//
// Example 1:
// Input: nums = [1,5,11,5]
// Output: true
// Explanation: The array can be partitioned as [1, 5, 5] and [11].
//
// Example 2:
// Input: nums = [1,2,3,5]
// Output: false
// Explanation: The array cannot be partitioned into equal sum subsets.

// canPartitionDP uses dynamic programming (subset sum problem)
// Time complexity: O(n * sum) where n is the length of nums and sum is total sum/2
// Space complexity: O(sum)
func canPartitionDP(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	// Calculate total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is odd, cannot partition into equal sum subsets
	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2

	// dp[i] represents whether we can achieve sum i using some subset of nums
	dp := make([]bool, target+1)
	dp[0] = true // Base case: empty subset has sum 0

	// For each number in nums
	for _, num := range nums {
		// Iterate backwards to avoid reusing the same element multiple times
		for i := target; i >= num; i-- {
			// If we can achieve sum i-num, then we can achieve sum i by adding num
			if dp[i-num] {
				dp[i] = true
			}
		}
		// Early exit if we found the target
		if dp[target] {
			return true
		}
	}

	return dp[target]
}

// canPartitionMemo uses memoization (top-down DP)
// Time complexity: O(n * sum)
// Space complexity: O(n * sum) for memoization
func canPartitionMemo(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	// Calculate total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is odd, cannot partition into equal sum subsets
	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2

	// Create memoization table
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 = uncomputed, 0 = false, 1 = true
		}
	}

	var dfs func(idx, remaining int) bool
	dfs = func(idx, remaining int) bool {
		// Base cases
		if remaining == 0 {
			return true
		}
		if idx >= n || remaining < 0 {
			return false
		}

		// Check memo
		if memo[idx][remaining] != -1 {
			return memo[idx][remaining] == 1
		}

		// Two choices: include nums[idx] or skip it
		include := dfs(idx+1, remaining-nums[idx])
		exclude := dfs(idx+1, remaining)

		result := include || exclude
		if result {
			memo[idx][remaining] = 1
		} else {
			memo[idx][remaining] = 0
		}

		return result
	}

	return dfs(0, target)
}

// CanPartition is the main function that chooses the appropriate algorithm
// It defaults to the bottom-up DP approach which is more space-efficient
func CanPartition(nums []int) bool {
	return canPartitionDP(nums)
}