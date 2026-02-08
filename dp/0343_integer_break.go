package dp

// IntegerBreak solves LeetCode problem 0343: Integer Break
// Difficulty: Medium
// Tags: Math, Dynamic Programming
//
// Given a positive integer n, break it into the sum of k positive integers (k >= 2) 
// such that the product of those integers is maximized.
// Return the maximum product you can get.
//
// Time complexity: O(n), Space complexity: O(1)
func IntegerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	
	// For n > 3, the optimal strategy is to break into as many 3s as possible
	// If remainder is 1, it's better to use one 3 and make it 2+2 (instead of 3+1)
	// If remainder is 2, just multiply by 2
	remainder := n % 3
	numberOfThrees := n / 3
	
	if remainder == 0 {
		// n is divisible by 3
		return pow3(numberOfThrees)
	} else if remainder == 1 {
		// Better to use 2*2 instead of 3*1
		return pow3(numberOfThrees-1) * 4
	} else {
		// remainder == 2
		// n = 3*numberOfThrees + 2
		return pow3(numberOfThrees) * 2
	}
}

// pow3 calculates 3 to the power of n efficiently
func pow3(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 3
	}
	return result
}