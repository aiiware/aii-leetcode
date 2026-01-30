package math

// MySqrt solves LeetCode problem 0069: Sqrt(x)
// Difficulty: Easy
// Tags: Math, Binary Search
//
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
// The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
//
// Example 1:
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.
//
// Example 2:
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
//
// Time complexity: O(log x), Space complexity: O(1)
func MySqrt(x int) int {
	if x < 0 {
		return -1 // Invalid input, but problem says non-negative
	}
	if x < 2 {
		return x // sqrt(0)=0, sqrt(1)=1
	}

	left, right := 1, x/2
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		
		// Use division to avoid overflow
		if mid <= x/mid {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}