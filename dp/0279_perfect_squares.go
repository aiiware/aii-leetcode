package dp

import "math"

// NumSquares solves LeetCode problem 0279: Perfect Squares
// Difficulty: Medium
// Tags: Dynamic Programming, Math, BFS
//
// Given an integer n, return the least number of perfect square numbers that sum to n.
// A perfect square is an integer that is the square of an integer.
//
// Example:
// Input: n = 12
// Output: 3
// Explanation: 12 = 4 + 4 + 4
//
// Input: n = 13
// Output: 2
// Explanation: 13 = 4 + 9
//
// Time complexity: O(n*sqrt(n)), Space complexity: O(n)
func NumSquares(n int) int {
	if n <= 0 {
		return 0
	}

	// dp[i] represents the least number of perfect squares that sum to i
	dp := make([]int, n+1)
	
	// Initialize dp array with maximum values
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	
	// Base case: 0 requires 0 perfect squares
	dp[0] = 0
	
	// Generate all perfect squares up to n
	maxSquare := int(math.Sqrt(float64(n)))
	squares := make([]int, maxSquare)
	for i := 1; i <= maxSquare; i++ {
		squares[i-1] = i * i
	}
	
	// Fill dp array
	for i := 1; i <= n; i++ {
		for _, square := range squares {
			if square > i {
				break
			}
			if dp[i-square]+1 < dp[i] {
				dp[i] = dp[i-square] + 1
			}
		}
	}
	
	return dp[n]
}