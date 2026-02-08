package dp

import "leetcode/utils"

// MaxProfitWithCooldown solves LeetCode problem 0309: Best Time to Buy and Sell Stock with Cooldown
// Difficulty: Medium
// Tags: Array, Dynamic Programming
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// Find the maximum profit you can achieve. You may complete as many transactions as you like
// with the following restrictions:
// - After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day)
// - You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again)
//
// Time complexity: O(n), Space complexity: O(1)
func MaxProfitWithCooldown(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// hold: max profit when holding a stock
	// sold: max profit when just sold a stock (in cooldown)
	// rest: max profit when in cooldown or can buy
	hold := -prices[0]
	sold := 0
	rest := 0

	for i := 1; i < len(prices); i++ {
		prevHold := hold
		prevSold := sold
		prevRest := rest

		// Either keep holding or buy today (from rest state)
		hold = utils.Max(prevHold, prevRest-prices[i])

		// Sell today (from hold state)
		sold = prevHold + prices[i]

		// Either stay in rest or come from sold state (cooldown)
		rest = utils.Max(prevRest, prevSold)
	}

	return utils.Max(sold, rest)
}
