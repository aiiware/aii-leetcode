package dp

// isMatchWildcard returns true if string s matches pattern p.
// Pattern p may contain:
//   - '?' which matches any single character
//   - '*' which matches any sequence of characters (including empty sequence)
//
// This is a classic 2D dynamic programming problem similar to regular expression matching.
// Time complexity: O(m*n) where m = len(s), n = len(p)
// Space complexity: O(m*n) for the DP table
func isMatchWildcard(s string, p string) bool {
	m, n := len(s), len(p)
	
	// dp[i][j] = true if s[0:i] matches p[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	
	// Empty string matches empty pattern
	dp[0][0] = true
	
	// Handle patterns starting with '*' that can match empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			// '*' can match empty sequence
			dp[0][j] = dp[0][j-1]
		}
		// For non-'*' patterns, empty string doesn't match (dp[0][j] remains false)
	}
	
	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				// '?' matches any single character, or characters match exactly
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' can match:
				// 1. Empty sequence: dp[i][j-1] (skip the '*')
				// 2. One or more characters: dp[i-1][j] (use '*' to match current char)
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
			// If characters don't match and pattern is not '?' or '*', dp[i][j] remains false
		}
	}
	
	return dp[m][n]
}

// isMatchWildcardOptimized is a space-optimized version using only 2 rows
// Time complexity: O(m*n)
// Space complexity: O(n)
func isMatchWildcardOptimized(s string, p string) bool {
	m, n := len(s), len(p)
	
	// Use only 2 rows to save space
	prev := make([]bool, n+1)
	curr := make([]bool, n+1)
	
	// Empty string matches empty pattern
	prev[0] = true
	
	// Handle patterns starting with '*' that can match empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			prev[j] = prev[j-1]
		}
	}
	
	// Fill the DP table row by row
	for i := 1; i <= m; i++ {
		// Initialize current row
		curr[0] = false // Non-empty string doesn't match empty pattern
		
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				// '?' matches any single character, or characters match exactly
				curr[j] = prev[j-1]
			} else if p[j-1] == '*' {
				// '*' can match empty sequence or one/more characters
				curr[j] = curr[j-1] || prev[j]
			} else {
				curr[j] = false
			}
		}
		
		// Swap rows for next iteration
		prev, curr = curr, prev
	}
	
	return prev[n]
}