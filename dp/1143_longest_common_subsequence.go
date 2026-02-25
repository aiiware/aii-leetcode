package dp

// LongestCommonSubsequence solves LeetCode problem 1143: Longest Common Subsequence
// Difficulty: Medium
// Tags: String, Dynamic Programming
//
// Given two strings text1 and text2, return the length of their longest common subsequence.
// If there is no common subsequence, return 0.
//
// A subsequence of a string is a new string generated from the original string with some
// characters (can be none) deleted without changing the relative order of the remaining characters.
//
// Example 1:
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
//
// Example 2:
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
//
// Example 3:
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.
//
// Constraints:
// 1 <= text1.length, text2.length <= 1000
// text1 and text2 consist of only lowercase English letters.
//
// Time complexity: O(m*n) where m = len(text1), n = len(text2)
// Space complexity: O(m*n) for the DP table
func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	
	// Create DP table with dimensions (m+1) x (n+1)
	// dp[i][j] represents LCS length of text1[0:i] and text2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Fill DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// Characters match: add 1 to the LCS from previous characters
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// Characters don't match: take max of LCS without current char from text1 or text2
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	return dp[m][n]
}

// LongestCommonSubsequenceOptimized solves the same problem with optimized space complexity
// Space complexity: O(min(m, n))
func LongestCommonSubsequenceOptimized(text1 string, text2 string) int {
	// Ensure text1 is the shorter string for space optimization
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}
	
	m, n := len(text1), len(text2)
	
	// Use two rows for DP to reduce space complexity
	prev := make([]int, n+1)
	curr := make([]int, n+1)
	
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		// Swap rows for next iteration
		prev, curr = curr, prev
		// Reset current row (not strictly necessary but good practice)
		for j := 0; j <= n; j++ {
			curr[j] = 0
		}
	}
	
	return prev[n]
}

// LongestCommonSubsequenceRecursive solves the problem using recursion with memoization
// Time complexity: O(m*n), Space complexity: O(m*n) for memoization
func LongestCommonSubsequenceRecursive(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	
	// Create memoization table
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	
	return lcsRecursiveHelper(text1, text2, 0, 0, memo)
}

func lcsRecursiveHelper(text1, text2 string, i, j int, memo [][]int) int {
	// Base case: reached end of either string
	if i == len(text1) || j == len(text2) {
		return 0
	}
	
	// Check memoization
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	
	var result int
	if text1[i] == text2[j] {
		// Characters match: include in LCS
		result = 1 + lcsRecursiveHelper(text1, text2, i+1, j+1, memo)
	} else {
		// Characters don't match: try skipping char from text1 or text2
		skipText1 := lcsRecursiveHelper(text1, text2, i+1, j, memo)
		skipText2 := lcsRecursiveHelper(text1, text2, i, j+1, memo)
		result = max(skipText1, skipText2)
	}
	
	// Memoize result
	memo[i][j] = result
	return result
}

// LongestCommonSubsequenceWithReconstruction returns the actual LCS string along with its length
func LongestCommonSubsequenceWithReconstruction(text1 string, text2 string) (string, int) {
	m, n := len(text1), len(text2)
	
	// Create DP table
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Fill DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	// Reconstruct LCS string
	lcs := make([]byte, dp[m][n])
	i, j, k := m, n, dp[m][n]-1
	
	for i > 0 && j > 0 && k >= 0 {
		if text1[i-1] == text2[j-1] {
			// Characters match: part of LCS
			lcs[k] = text1[i-1]
			i--
			j--
			k--
		} else if dp[i-1][j] > dp[i][j-1] {
			// Move in direction of larger LCS
			i--
		} else {
			j--
		}
	}
	
	return string(lcs), dp[m][n]
}