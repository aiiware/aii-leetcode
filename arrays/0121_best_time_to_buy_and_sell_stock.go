package arrays

// MaxProfit solves LeetCode problem 0121: Best Time to Buy and Sell Stock
// Difficulty: Easy
// Tags: Array, Dynamic Programming
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and 
// choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. 
// If you cannot achieve any profit, return 0.
//
// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
//
// Example 2:
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.
//
// Time complexity: O(n), Space complexity: O(1)
func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// Update minimum price seen so far
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		
		// Calculate potential profit if selling at current price
		profit := prices[i] - minPrice
		
		// Update maximum profit if current profit is better
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}