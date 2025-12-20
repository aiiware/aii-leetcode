package leetcode

import "fmt"

// IsPalindrome solves LeetCode problem 0009: Palindrome Number
// Difficulty: Easy
// Tags: Math
//
// Given an integer x, return true if x is a palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// Negative numbers are not palindromes.
//
// Time complexity: O(log10(x)), Space complexity: O(1)
func IsPalindrome(x int) bool {
	return isPalindromeMath(x)
}

// isPalindromeMath mathematical approach (reverse half of the number).
// Time complexity: O(log10(x)), Space complexity: O(1)
func isPalindromeMath(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Numbers ending with 0 (except 0 itself) are not palindromes
	if x%10 == 0 && x != 0 {
		return false
	}

	reversedHalf := 0
	// Reverse only half of the number
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// Check if original number equals reversed half (odd length case)
	// or original number equals reversed half / 10 (even length case)
	return x == reversedHalf || x == reversedHalf/10
}

// isPalindromeString string conversion approach.
// Time complexity: O(log10(x)), Space complexity: O(log10(x))
func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	s := fmt.Sprintf("%d", x)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}