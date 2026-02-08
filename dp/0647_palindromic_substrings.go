package dp

// CountSubstrings solves LeetCode problem 0647: Palindromic Substrings
// Difficulty: Medium
// Tags: String, Dynamic Programming
//
// Given a string s, return the number of palindromic substrings in it.
// A string is a palindrome when it reads the same backward as forward.
// A substring is a contiguous sequence of characters within the string.
//
// Time complexity: O(n^2), Space complexity: O(1)
func CountSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	count := 0
	
	// For each possible center of palindrome (both odd and even length)
	for i := 0; i < len(s); i++ {
		// Odd length palindromes (center is at i)
		count += expandAroundCenter(s, i, i)
		// Even length palindromes (center is between i and i+1)
		count += expandAroundCenter(s, i, i+1)
	}
	
	return count
}

// expandAroundCenter expands around the center and counts palindromes
func expandAroundCenter(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}