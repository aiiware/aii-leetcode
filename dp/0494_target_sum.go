package dp

// FindTargetSumWays solves LeetCode problem 0494: Target Sum
// Difficulty: Medium
// Tags: Array, Dynamic Programming
//
// You are given an integer array nums and an integer target.
// You want to build an expression out of nums by adding one of the symbols '+' and '-' 
// before each integer and then concatenate all the integers.
// Return the number of different expressions that evaluate to target.
//
// Time complexity: O(n * sum), Space complexity: O(sum)
func FindTargetSumWays(nums []int, target int) int {
	// Calculate the total sum
	sum := 0
	for _, num := range nums {
		sum += num
	}
	
	// If target is greater than sum, it's impossible
	if target > sum || target < -sum {
		return 0
	}
	
	// We can convert this to subset sum problem
	// If we partition nums into two subsets S1 and S2 such that:
	// S1 - S2 = target and S1 + S2 = sum
	// Then: S1 = (target + sum) / 2
	// So we need to find number of ways to make subset with sum = (target + sum) / 2
	
	// New target for subset sum
	newTarget := target + sum
	if newTarget%2 != 0 {
		return 0
	}
	newTarget = newTarget / 2
	
	// dp[i] represents number of ways to achieve sum i
	dp := make([]int, newTarget+1)
	dp[0] = 1 // One way to get sum 0 (empty subset)
	
	// For each number, update dp array
	for _, num := range nums {
		// Traverse backwards to avoid using updated values in same iteration
		for i := newTarget; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	
	return dp[newTarget]
}