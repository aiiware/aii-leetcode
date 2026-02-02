package math

// 0202 - Happy Number (Easy)
// https://leetcode.com/problems/happy-number/

// IsHappy determines if a number is a happy number
// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Time Complexity: O(log n) - each step reduces the number significantly
// Space Complexity: O(log n) - for the hash set
func IsHappy(n int) bool {
	seen := make(map[int]bool)
	
	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumOfSquares(n)
	}
	
	return n == 1
}

// IsHappyFloyd uses Floyd's Cycle Detection algorithm (tortoise and hare)
// This approach uses O(1) space
func IsHappyFloyd(n int) bool {
	slow := n
	fast := sumOfSquares(n)
	
	for fast != 1 && slow != fast {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))
	}
	
	return fast == 1
}

// sumOfSquares calculates the sum of squares of digits of a number
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// IsHappyMath uses mathematical properties of happy numbers
// All numbers in [1, 9] are either happy or lead to a cycle
func IsHappyMath(n int) bool {
	for n > 9 {
		n = sumOfSquares(n)
	}
	
	// Single digit happy numbers are 1 and 7
	return n == 1 || n == 7
}