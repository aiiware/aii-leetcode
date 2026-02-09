package arrays

// LongestPalindromicSubstring solves LeetCode problem 0005: Longest Palindromic Substring
// Difficulty: Hard
// Tags: String, Dynamic Programming
//
// Given a string s, return the longest palindromic substring in s.
//
// Time complexity: O(n^2), Space complexity: O(n^2)
func LongestPalindromicSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		// Odd length palindromes
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}

		// Even length palindromes
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	return s[start : start+maxLen]
}
