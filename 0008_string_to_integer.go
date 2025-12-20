package leetcode

import (
	"math"
)

// MyAtoi solves LeetCode problem 0008: String to Integer (atoi)
// Difficulty: Medium
// Tags: String
//
// Implement the myAtoi(string s) function, which converts a string to a 32-bit
// signed integer (similar to C/C++'s atoi function).
//
// The algorithm:
// 1. Read in and ignore any leading whitespace.
// 2. Check if the next character is '-' or '+'.
// 3. Read in next characters until a non-digit or end of string.
// 4. Convert digits to integer, clamp to 32-bit range.
//
// Time complexity: O(n), Space complexity: O(1)
func MyAtoi(s string) int {
	i := 0
	n := len(s)

	// 1. Skip leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	if i >= n {
		return 0
	}

	// 2. Handle sign
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	// 3. Convert digits
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Check for overflow before multiplying by 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}