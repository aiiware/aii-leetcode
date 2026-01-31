package strings

/*
115. Distinct Subsequences
https://leetcode.com/problems/distinct-subsequences/

Given two strings s and t, return the number of distinct subsequences of s which equals t.

The test cases are generated so that the answer fits on a 32-bit signed integer.

Example 1:
Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
There are 3 ways you can generate "rabbit" from "rabbbit":
rabb b it
ra b bbit
rab b bit

Example 2:
Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
There are 5 ways you can generate "bag" from "babgbag":
ba b g bag
ba bgba g
b abgb ag
ba b gb ag
babg bag

Constraints:
- 1 <= s.length, t.length <= 1000
- s and t consist of English letters.

Difficulty: Hard
Tags: String, Dynamic Programming
Companies: Google, Amazon, Microsoft, Bloomberg
*/

// numDistinctDP is the standard dynamic programming solution.
// Uses a 2D DP table where dp[i][j] = number of distinct subsequences of s[0:i] that equal t[0:j].
// Time complexity: O(m*n), Space complexity: O(m*n)
func numDistinctDP(s string, t string) int {
	m, n := len(s), len(t)
	
	// If t is longer than s, no subsequences possible
	if n > m {
		return 0
	}
	
	// Create DP table with dimensions (m+1) x (n+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Base cases:
	// 1. Empty t can be formed from any prefix of s in exactly 1 way (by taking empty subsequence)
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}
	// 2. Non-empty t cannot be formed from empty s (except when t is also empty, handled above)
	// dp[0][j] = 0 for j > 0 (already initialized to 0)
	
	// Fill DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If characters match, we have two options:
			// 1. Use the current character: dp[i-1][j-1]
			// 2. Skip the current character: dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				// Characters don't match, can only skip current character
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	
	return dp[m][n]
}

// numDistinctDPSpaceOptimized is a space-optimized version using 1D DP array.
// Since we only need previous row, we can reduce space to O(n).
// Time complexity: O(m*n), Space complexity: O(n)
func numDistinctDPSpaceOptimized(s string, t string) int {
	m, n := len(s), len(t)
	
	if n > m {
		return 0
	}
	
	// dp[j] represents number of ways to form t[0:j] using current prefix of s
	dp := make([]int, n+1)
	dp[0] = 1 // Empty t can be formed in 1 way
	
	// Fill DP array
	for i := 1; i <= m; i++ {
		// Process from right to left to avoid overwriting values we need
		for j := n; j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
			// If characters don't match, dp[j] stays the same (skip current char)
		}
	}
	
	return dp[n]
}

// numDistinctMemoization is a memoized recursive solution.
// This approach uses top-down recursion with memoization.
// Time complexity: O(m*n), Space complexity: O(m*n) for memoization
func numDistinctMemoization(s string, t string) int {
	m, n := len(s), len(t)
	
	if n > m {
		return 0
	}
	
	// Create memoization table
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// Base cases
		if j == n { // All characters of t matched
			return 1
		}
		if i == m { // Reached end of s but not all of t matched
			return 0
		}
		
		// Check memo
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		
		// Recursive cases
		result := 0
		if s[i] == t[j] {
			// Option 1: Use current character (match)
			result += dfs(i+1, j+1)
		}
		// Option 2: Skip current character (always available)
		result += dfs(i+1, j)
		
		memo[i][j] = result
		return result
	}
	
	return dfs(0, 0)
}

// numDistinctDPWithLargeNumbers handles cases where result might exceed 32-bit.
// Uses int64 to avoid overflow.
// Time complexity: O(m*n), Space complexity: O(n)
func numDistinctDPWithLargeNumbers(s string, t string) int64 {
	m, n := len(s), len(t)
	
	if n > m {
		return 0
	}
	
	dp := make([]int64, n+1)
	dp[0] = 1
	
	for i := 1; i <= m; i++ {
		// Process from right to left
		for j := n; j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	
	return dp[n]
}