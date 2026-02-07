package dp

// 322. Coin Change
// https://leetcode.com/problems/coin-change/
//
// You are given an integer array coins representing coins of different denominations
// and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
// Example 1:
// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
//
// Example 2:
// Input: coins = [2], amount = 3
// Output: -1
//
// Example 3:
// Input: coins = [1], amount = 0
// Output: 0

// coinChangeDP uses dynamic programming to find the minimum number of coins
// Time complexity: O(amount * n) where n is the number of coin denominations
// Space complexity: O(amount)
func coinChangeDP(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i] represents the minimum number of coins to make amount i
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // Initialize with a value larger than any possible solution
	}

	// Base case: 0 coins needed to make amount 0
	dp[0] = 0

	// For each amount from 1 to target amount
	for i := 1; i <= amount; i++ {
		// Try each coin denomination
		for _, coin := range coins {
			if coin <= i {
				// If we can use this coin, check if it gives us a better solution
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	// If dp[amount] is still the initial value, no solution exists
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// coinChangeBFS uses BFS to find the minimum number of coins
// This approach is often faster for some inputs
// Time complexity: O(amount * n) in worst case
// Space complexity: O(amount)
func coinChangeBFS(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	visited := make([]bool, amount+1)
	queue := []int{0}
	visited[0] = true
	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		steps++

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			for _, coin := range coins {
				next := current + coin

				if next == amount {
					return steps
				}

				if next < amount && !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

// CoinChange is the main function that chooses the appropriate algorithm
// It defaults to DP which is more space-efficient for large amounts
func CoinChange(coins []int, amount int) int {
	return coinChangeDP(coins, amount)
}