package dp

// IsSubsequence solves LeetCode problem 0392: Is Subsequence
// Difficulty: Easy
// Tags: Two Pointers, Dynamic Programming
//
// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by 
// deleting some (can be none) of the characters without disturbing the relative positions 
// of the remaining characters.
//
// Time complexity: O(n) where n is the length of t, Space complexity: O(1)
func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	
	// Two pointers approach
	sIndex := 0
	tIndex := 0
	
	for tIndex < len(t) && sIndex < len(s) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}
	
	// If we've matched all characters in s, then it's a subsequence
	return sIndex == len(s)
}