package math

// AddDigits solves LeetCode problem 0258: Add Digits
// Difficulty: Easy
// Tags: Math, Simulation, Number Theory
//
// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
//
// Example 1:
// Input: num = 38
// Output: 2
// Explanation: The process is
// 38 --> 3 + 8 --> 11
// 11 --> 1 + 1 --> 2
// Since 2 has only one digit, return it.
//
// Example 2:
// Input: num = 0
// Output: 0
//
// Constraints:
// - 0 <= num <= 2^31 - 1
//
// Follow up: Could you do it without any loop/recursion in O(1) runtime?
//
// Time complexity: O(1) with mathematical formula, O(log n) with loop
// Space complexity: O(1)
func AddDigits(num int) int {
	// Mathematical solution using digital root formula
	// Digital root of n = 0 if n == 0
	// Digital root of n = 9 if n % 9 == 0 and n != 0
	// Digital root of n = n % 9 otherwise
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

// AddDigitsLoop is the iterative solution
func AddDigitsLoop(num int) int {
	for num >= 10 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	return num
}