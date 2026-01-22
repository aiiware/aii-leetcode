package leetcode

/*
172. Factorial Trailing Zeroes

Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

Example 1:
Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.

Example 2:
Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.

Example 3:
Input: n = 0
Output: 0
Explanation: 0! = 1, no trailing zero.

Constraints:
- 0 <= n <= 10^4

Difficulty: Medium
Tags: Math
Companies: Microsoft, Google, Amazon, Bloomberg, Apple
*/

// trailingZeroes returns the number of trailing zeroes in n!.
// Trailing zeroes are created by factors of 10, and 10 = 2 * 5.
// Since there are always more factors of 2 than 5 in factorial,
// we only need to count the number of factors of 5.
func trailingZeroes(n int) int {
	count := 0
	// Count factors of 5, 25, 125, etc.
	// Each multiple of 5 contributes at least one factor of 5
	// Each multiple of 25 contributes an additional factor of 5
	// Each multiple of 125 contributes another factor of 5, and so on
	for i := 5; i <= n; i *= 5 {
		count += n / i
	}
	return count
}