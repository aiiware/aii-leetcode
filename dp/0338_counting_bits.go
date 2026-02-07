package dp

// CountBits solves LeetCode problem 0338: Counting Bits
// Difficulty: Easy
// Tags: Dynamic Programming, Bit Manipulation
//
// Given an integer n, return an array ans of length n + 1 such that
// for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
//
// Example:
// Input: n = 2
// Output: [0,1,1]
// Explanation:
// 0 --> 0 (0 ones)
// 1 --> 1 (1 one)
// 2 --> 10 (1 one)
//
// Input: n = 5
// Output: [0,1,1,2,1,2]
//
// Time complexity: O(n), Space complexity: O(n)
func CountBits(n int) []int {
	if n < 0 {
		return []int{}
	}

	// dp[i] represents the number of 1's in binary representation of i
	dp := make([]int, n+1)
	
	// Base case: 0 has 0 ones
	dp[0] = 0
	
	// Fill dp array using the recurrence relation:
	// dp[i] = dp[i >> 1] + (i & 1)
	// This works because:
	// - i >> 1 is i/2 (removes the least significant bit)
	// - i & 1 gives 1 if the least significant bit is 1, 0 otherwise
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	
	return dp
}

// Alternative solution using another recurrence relation:
// dp[i] = dp[i & (i-1)] + 1
// This works because i & (i-1) removes the least significant 1-bit
func CountBits2(n int) []int {
	if n < 0 {
		return []int{}
	}

	dp := make([]int, n+1)
	dp[0] = 0
	
	for i := 1; i <= n; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	
	return dp
}