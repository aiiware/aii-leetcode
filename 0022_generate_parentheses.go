package leetcode

// GenerateParenthesis solves LeetCode problem 0022: Generate Parentheses
// Difficulty: Medium
// Tags: String, Backtracking, Dynamic Programming
//
// Given n pairs of parentheses, write a function to generate all combinations
// of well-formed parentheses.
//
// Time complexity: O(4^n / sqrt(n)) - nth Catalan number
// Space complexity: O(4^n / sqrt(n)) for storing results
func GenerateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, "", 0, 0, n)
	return result
}

// backtrack is a helper function for recursive backtracking
func backtrack(result *[]string, current string, open, close, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}

// GenerateParenthesisDP uses dynamic programming approach
func GenerateParenthesisDP(n int) []string {
	if n == 0 {
		return []string{""}
	}

	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for i := 1; i <= n; i++ {
		dp[i] = []string{}
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-1-j] {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}
		}
	}

	return dp[n]
}