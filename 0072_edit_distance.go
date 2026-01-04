package leetcode

// MinDistance solves LeetCode problem 0072: Edit Distance
// Difficulty: Hard
// Tags: String, Dynamic Programming
//
// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
//
// You have the following three operations permitted on a word:
// - Insert a character
// - Delete a character
// - Replace a character
//
// Example 1:
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
//
// Example 2:
// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')
//
// Constraints:
// 0 <= word1.length, word2.length <= 500
// word1 and word2 consist of lowercase English letters.
//
// Time complexity: O(m*n), Space complexity: O(min(m,n))
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// Use the shorter string for space optimization
	if m < n {
		return MinDistance(word2, word1)
	}

	// Create DP array with size n+1
	dp := make([]int, n+1)

	// Initialize first row (empty word1 to word2)
	for j := 0; j <= n; j++ {
		dp[j] = j // j insertions
	}

	// Fill DP table
	for i := 1; i <= m; i++ {
		prev := dp[0] // dp[i-1][0]
		dp[0] = i     // i deletions

		for j := 1; j <= n; j++ {
			temp := dp[j] // dp[i-1][j]

			if word1[i-1] == word2[j-1] {
				// Characters match, no operation needed
				dp[j] = prev // dp[i][j] = dp[i-1][j-1]
			} else {
				// Take minimum of three operations:
				// 1. Delete: dp[i-1][j] + 1
				// 2. Insert: dp[i][j-1] + 1
				// 3. Replace: dp[i-1][j-1] + 1
				dp[j] = min(temp, min(dp[j-1], prev)) + 1
			}

			prev = temp // Update prev for next iteration
		}
	}

	return dp[n]
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}