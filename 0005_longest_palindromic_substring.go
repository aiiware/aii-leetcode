package leetcode

// LongestPalindrome solves LeetCode problem 0005: Longest Palindromic Substring
// Difficulty: Medium
// Tags: String, Dynamic Programming
//
// Given a string s, return the longest palindromic substring in s.
//
// Time complexity: O(n²), Space complexity: O(1) for expand around center
func LongestPalindrome(s string) string {
	return longestPalindromeExpand(s)
}

// longestPalindromeExpand expands around center approach.
// Time complexity: O(n²), Space complexity: O(1)
func longestPalindromeExpand(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindrome (single character center)
		len1 := expandAroundCenter(s, i, i)
		// Check for even-length palindrome (two character center)
		len2 := expandAroundCenter(s, i, i+1)

		// Take the longer palindrome
		maxLen := maxInt(len1, len2)

		// Update start and end if we found a longer palindrome
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// expandAroundCenter expands around the given center(s) and returns palindrome length
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// Return length of palindrome (right - left - 1)
	return right - left - 1
}

// longestPalindromeDP dynamic programming approach.
// Time complexity: O(n²), Space complexity: O(n²)
func longestPalindromeDP(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	// dp[i][j] = true if s[i:j+1] is palindrome
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// All single characters are palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	start, maxLen := 0, 1

	// Check for palindromes of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// Check for palindromes of length >= 3
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			// s[i:j+1] is palindrome if s[i] == s[j] and s[i+1:j] is palindrome
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLen = length
			}
		}
	}

	return s[start : start+maxLen]
}