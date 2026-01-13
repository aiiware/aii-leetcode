package leetcode

import (
	"strings"
	"unicode"
)

// IsValidPalindrome solves LeetCode problem 0125: Valid Palindrome
// Difficulty: Easy
// Tags: Two Pointers, String
//
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward.
// Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
//
// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
//
// Constraints:
// 1 <= s.length <= 2 * 10^5
// s consists only of printable ASCII characters.
//
// Time complexity: O(n), Space complexity: O(1)
func IsValidPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	// Convert to lowercase for case-insensitive comparison
	s = strings.ToLower(s)

	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric characters from left
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}

		// Skip non-alphanumeric characters from right
		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}

		// Compare characters
		if left < right && s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// isAlphanumeric checks if a character is a letter or digit
func isAlphanumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}