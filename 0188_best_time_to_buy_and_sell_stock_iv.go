package leetcode

// MaxProfitIV solves LeetCode problem 0188: Best Time to Buy and Sell Stock IV
// Difficulty: Hard
// Tags: Array, Dynamic Programming
//
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day,
// and an integer k.
//
// Find the maximum profit you can achieve. You may complete at most k transactions.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
// Example 1:
// Input: k = 2, prices = [2,4,1]
// Output: 2
// Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
//
// Example 2:
// Input: k = 2, prices = [3,2,6,5,0,3]
// Output: 7
// Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
// Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//
// Constraints:
// 1 <= k <= 100
// 1 <= prices.length <= 1000
// 0 <= prices[i] <= 1000
//
// Time complexity: O(k*n), Space complexity: O(k)
func MaxProfitIV(k int, prices []int) int {
	n := len(prices)
	if n < 2 || k < 1 {
		return 0
	}

	// If k >= n/2, we can make as many transactions as we want
	// This becomes the same as problem 122 (unlimited transactions)
	if k >= n/2 {
		return maxProfitUnlimited(prices)
	}

	// DP approach for limited transactions
	// dp[i][j] = max profit with at most i transactions up to day j
	// We can optimize space to O(k) by using rolling arrays
	buy := make([]int, k+1)
	sell := make([]int, k+1)

	// Initialize buy array with minimum possible value
	for i := 0; i <= k; i++ {
		buy[i] = -prices[0]
	}

	// Iterate through prices
	for i := 1; i < n; i++ {
		// Update for each transaction count
		for j := 1; j <= k; j++ {
			// Option 1: Don't buy on day i (carry over from previous day)
			// Option 2: Buy on day i (profit from previous sell minus current price)
			buy[j] = max(buy[j], sell[j-1]-prices[i])

			// Option 1: Don't sell on day i (carry over from previous day)
			// Option 2: Sell on day i (profit from previous buy plus current price)
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	return sell[k]
}

// Helper function for unlimited transactions (k >= n/2 case)
func maxProfitUnlimited(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}