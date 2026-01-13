package leetcode

// MaxProfitIII solves LeetCode problem 0123: Best Time to Buy and Sell Stock III
// Difficulty: Hard
// Tags: Array, Dynamic Programming
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// Find the maximum profit you can achieve. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
// Example 1:
// Input: prices = [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
// Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
//
// Example 2:
// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//
// Example 3:
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e., max profit = 0.
//
// Constraints:
// 1 <= prices.length <= 10^5
// 0 <= prices[i] <= 10^5
//
// Time complexity: O(n), Space complexity: O(n)
func MaxProfitIII(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	n := len(prices)

	// leftProfits[i] = max profit with one transaction in prices[0..i]
	leftProfits := make([]int, n)
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		leftProfits[i] = max(leftProfits[i-1], prices[i]-minPrice)
	}

	// rightProfits[i] = max profit with one transaction in prices[i..n-1]
	rightProfits := make([]int, n)
	maxPrice := prices[n-1]

	for i := n - 2; i >= 0; i-- {
		maxPrice = max(maxPrice, prices[i])
		rightProfits[i] = max(rightProfits[i+1], maxPrice-prices[i])
	}

	// Find the best combination of two transactions
	maxProfit := 0
	for i := 0; i < n; i++ {
		// Profit from left segment + profit from right segment
		profit := leftProfits[i]
		if i+1 < n {
			profit += rightProfits[i+1]
		}
		maxProfit = max(maxProfit, profit)
	}

	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}