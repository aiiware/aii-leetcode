package leetcode

import (
	"math"
)

// Reverse solves LeetCode problem 0007: Reverse Integer
// Difficulty: Medium
// Tags: Math
//
// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer
// range [-2^31, 2^31 - 1], then return 0.
//
// Time complexity: O(log10(x)), Space complexity: O(1)
func Reverse(x int) int {
	return reverseInt32(x)
}

// reverseInt32 handles 32-bit integer overflow checking.
// Time complexity: O(log10(x)), Space complexity: O(1)
func reverseInt32(x int) int {
	result := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before multiplying by 10
		if result > math.MaxInt32/10 {
			return 0
		}
		result *= 10

		// Check for overflow before adding digit
		if result > math.MaxInt32-digit {
			return 0
		}
		result += digit
	}

	return result * sign
}

// reverseInt64 uses int64 to avoid overflow issues.
// Time complexity: O(log10(x)), Space complexity: O(1)
func reverseInt64(x int) int {
	var result int64 = 0
	original := int64(x)

	for original != 0 {
		digit := original % 10
		original /= 10
		result = result*10 + digit
	}

	// Check for 32-bit overflow
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	return int(result)
}