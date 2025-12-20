package leetcode

// IsMatch solves LeetCode problem 0010: Regular Expression Matching
// Difficulty: Hard
// Tags: String, Dynamic Programming, Recursion
//
// Given an input string s and a pattern p, implement regular expression matching
// with support for '.' and '*' where:
//   '.' Matches any single character.
//   '*' Matches zero or more of the preceding element.
//
// The matching should cover the entire input string (not partial).
//
// Time complexity: O(m*n) where m = len(s), n = len(p)
// Space complexity: O(m*n) for DP table
func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// dp[i][j] = true if s[0:i] matches p[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Empty string matches empty pattern
	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* etc. that can match empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			// Match zero of the preceding character
			if j >= 2 {
				dp[0][j] = dp[0][j-2]
			}
			// If star is at beginning (j=1), pattern is invalid, dp[0][1] remains false
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// Current characters match
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' && j >= 2 {
				// '*' case (only valid if there's a preceding character)
				// Option 1: Match zero of the preceding character
				dp[i][j] = dp[i][j-2]

				// Option 2: Match one or more of the preceding character
				// Check if preceding character matches current s character
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
			// If p[j-1] == '*' && j < 2, pattern is invalid, dp[i][j] remains false
		}
	}

	return dp[m][n]
}

// IsMatchRecursive recursive solution with memoization.
// Time complexity: O(m*n)
// Space complexity: O(m*n) for memoization
func IsMatchRecursive(s string, p string) bool {
	memo := make(map[[2]int]bool)
	return isMatchHelper(s, p, 0, 0, memo)
}

func isMatchHelper(s, p string, i, j int, memo map[[2]int]bool) bool {
	// Check memo
	key := [2]int{i, j}
	if val, ok := memo[key]; ok {
		return val
	}

	// Base cases
	if j == len(p) {
		result := i == len(s)
		memo[key] = result
		return result
	}

	// Check if current characters match
	firstMatch := i < len(s) && (p[j] == '.' || p[j] == s[i])

	var result bool
	if j+1 < len(p) && p[j+1] == '*' {
		// '*' case: either match zero or match one+
		result = isMatchHelper(s, p, i, j+2, memo) || // Match zero
			(firstMatch && isMatchHelper(s, p, i+1, j, memo)) // Match one+
	} else {
		// Normal match case
		result = firstMatch && isMatchHelper(s, p, i+1, j+1, memo)
	}

	memo[key] = result
	return result
}