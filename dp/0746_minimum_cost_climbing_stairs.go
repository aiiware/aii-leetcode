package dp

// 746. Minimum Cost Climbing Stairs
//
// You are given an integer array cost where cost[i] is the cost of i-th step on a staircase.
// Once you pay the cost, you can either climb one or two steps.
// You can either start from the step with index 0, or the step with index 1.
// Return the minimum cost to reach the top of the floor (beyond the last stair).
//
// Example 1:
// Input: cost = [10,15,20]
// Output: 15
// Explanation: You will start at index 1.
// - Pay 15 and climb two steps to reach the top.
// The total cost is 15.
//
// Example 2:
// Input: cost = [1,100,1,1,1,100,1,1,100,1]
// Output: 6
// Explanation: You will start at index 0.
// - Pay 1 and climb two steps to reach index 2.
// - Pay 1 and climb two steps to reach index 4.
// - Pay 1 and climb two steps to reach index 6.
// - Pay 1 and climb one step to reach index 7.
// - Pay 1 and climb two steps to reach index 9.
// - Pay 1 and climb one step to reach the top.
// The total cost is 6.
//
// Constraints:
// - 2 <= cost.length <= 1000
// - 0 <= cost[i] <= 999

// minCostClimbingStairsDP solves using dynamic programming with O(n) time and O(n) space
func minCostClimbingStairsDP(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	// dp[i] represents the minimum cost to reach step i
	dp := make([]int, n+1)
	
	// Base cases: we can start from step 0 or step 1 with 0 cost
	dp[0] = 0
	dp[1] = 0
	
	// Recurrence relation: dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	
	return dp[n]
}

// minCostClimbingStairsOptimized solves using dynamic programming with O(n) time and O(1) space
func minCostClimbingStairsOptimized(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}
	
	// We only need to keep track of the last two values
	prev2 := 0 // dp[i-2]
	prev1 := 0 // dp[i-1]
	
	for i := 2; i <= n; i++ {
		current := min(prev1+cost[i-1], prev2+cost[i-2])
		prev2 = prev1
		prev1 = current
	}
	
	return prev1
}

// minCostClimbingStairs is the main function that uses the optimized solution
func minCostClimbingStairs(cost []int) int {
	return minCostClimbingStairsOptimized(cost)
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}