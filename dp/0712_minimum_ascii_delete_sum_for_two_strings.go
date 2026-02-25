package dp

// minimumDeleteSum returns the minimum ASCII sum of deleted characters to make two strings equal.
// Given two strings s1 and s2, we need to delete characters from both strings so that
// the resulting strings become equal. The cost of deleting a character is its ASCII value.
// We need to find the minimum total ASCII sum of deleted characters.
//
// This is a classic 2D dynamic programming problem similar to Longest Common Subsequence (LCS),
// but instead of maximizing the length of common subsequence, we minimize the ASCII cost
// of characters not in the common subsequence.
//
// Approach:
// 1. DP[i][j] = minimum ASCII delete sum for s1[0:i] and s2[0:j]
// 2. Base cases:
//    - DP[i][0] = sum of ASCII values of first i characters of s1 (delete all from s1)
//    - DP[0][j] = sum of ASCII values of first j characters of s2 (delete all from s2)
// 3. Recurrence:
//    - If s1[i-1] == s2[j-1]: DP[i][j] = DP[i-1][j-1] (no deletion needed)
//    - Else: DP[i][j] = min(DP[i-1][j] + ascii(s1[i-1]), DP[i][j-1] + ascii(s2[j-1]))
//
// Time Complexity: O(m*n) where m = len(s1), n = len(s2)
// Space Complexity: O(m*n) for the DP table
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	
	// dp[i][j] = minimum ASCII delete sum for s1[0:i] and s2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Base case: delete all characters from s1 when s2 is empty
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	
	// Base case: delete all characters from s2 when s1 is empty
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	
	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				// Characters match, no deletion needed
				dp[i][j] = dp[i-1][j-1]
			} else {
				// Characters don't match, delete from either s1 or s2
				deleteFromS1 := dp[i-1][j] + int(s1[i-1]) // Delete s1[i-1]
				deleteFromS2 := dp[i][j-1] + int(s2[j-1]) // Delete s2[j-1]
				
				// Take the minimum cost
				if deleteFromS1 < deleteFromS2 {
					dp[i][j] = deleteFromS1
				} else {
					dp[i][j] = deleteFromS2
				}
			}
		}
	}
	
	return dp[m][n]
}

// minimumDeleteSumOptimized is a space-optimized version using only 2 rows
// Time Complexity: O(m*n)
// Space Complexity: O(n)
func minimumDeleteSumOptimized(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	
	// Use only 2 rows to save space
	prev := make([]int, n+1)
	curr := make([]int, n+1)
	
	// Base case: delete all characters from s2 when s1 is empty
	for j := 1; j <= n; j++ {
		prev[j] = prev[j-1] + int(s2[j-1])
	}
	
	// Fill the DP table row by row
	for i := 1; i <= m; i++ {
		// Initialize current row: delete all characters from s1 up to i when s2 is empty
		curr[0] = prev[0] + int(s1[i-1])
		
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				// Characters match, no deletion needed
				curr[j] = prev[j-1]
			} else {
				// Characters don't match, delete from either s1 or s2
				deleteFromS1 := prev[j] + int(s1[i-1]) // Delete s1[i-1]
				deleteFromS2 := curr[j-1] + int(s2[j-1]) // Delete s2[j-1]
				
				// Take the minimum cost
				if deleteFromS1 < deleteFromS2 {
					curr[j] = deleteFromS1
				} else {
					curr[j] = deleteFromS2
				}
			}
		}
		
		// Swap rows for next iteration
		prev, curr = curr, prev
	}
	
	return prev[n]
}