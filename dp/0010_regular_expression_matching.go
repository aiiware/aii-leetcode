package dp

// isMatch returns true if string s matches pattern p.
// Pattern p may contain '.' which matches any single character,
// and '*' which matches zero or more of the preceding element.
//
// This is a classic 2D dynamic programming problem.
// Time complexity: O(m*n) where m = len(s), n = len(p)
// Space complexity: O(m*n) for the DP table
func isMatch(s string, p string) bool {
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
			// '*' can match zero of the preceding element
			dp[0][j] = dp[0][j-2]
		}
	}
	
	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// Current characters match ('.' matches any char)
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' can match zero or more of the preceding element
				// Case 1: Match zero of the preceding element
				dp[i][j] = dp[i][j-2]
				
				// Case 2: Match one or more of the preceding element
				// Check if preceding character matches current s character
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	
	return dp[m][n]
}

// isMatchOptimized is a space-optimized version using only 2 rows
// Time complexity: O(m*n)
// Space complexity: O(n)
func isMatchOptimized(s string, p string) bool {
	m, n := len(s), len(p)
	
	// Use only 2 rows to save space
	prev := make([]bool, n+1)
	curr := make([]bool, n+1)
	
	// Empty string matches empty pattern
	prev[0] = true
	
	// Handle patterns that can match empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			prev[j] = prev[j-2]
		}
	}
	
	// Fill the DP table row by row
	for i := 1; i <= m; i++ {
		// Initialize current row
		curr[0] = false // Non-empty string doesn't match empty pattern
		
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// Current characters match
				curr[j] = prev[j-1]
			} else if p[j-1] == '*' {
				// '*' can match zero or more of the preceding element
				// Case 1: Match zero of the preceding element
				curr[j] = curr[j-2]
				
				// Case 2: Match one or more of the preceding element
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					curr[j] = curr[j] || prev[j]
				}
			} else {
				curr[j] = false
			}
		}
		
		// Swap rows for next iteration
		prev, curr = curr, prev
	}
	
	return prev[n]
}