package math

// IsUgly solves LeetCode problem 0263: Ugly Number
// Difficulty: Easy
// Tags: Math
//
// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
// Given an integer n, return true if n is an ugly number.
//
// Example 1:
// Input: n = 6
// Output: true
// Explanation: 6 = 2 × 3
//
// Example 2:
// Input: n = 1
// Output: true
// Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
//
// Example 3:
// Input: n = 14
// Output: false
// Explanation: 14 is not ugly since it includes the prime factor 7.
//
// Constraints:
// - -2^31 <= n <= 2^31 - 1
//
// Time complexity: O(log n), Space complexity: O(1)
func IsUgly(n int) bool {
	// Edge cases
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	
	// Keep dividing by 2, 3, 5 while possible
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	
	// If after removing all factors of 2, 3, 5 we get 1, it's ugly
	return n == 1
}