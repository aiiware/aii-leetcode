package dp

// ClimbStairs solves LeetCode problem 0070: Climbing Stairs
// Difficulty: Easy
// Tags: Math, Dynamic Programming, Memoization
//
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// Example 2:
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
//
// Time complexity: O(n), Space complexity: O(1)
func ClimbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// Fibonacci-like sequence: f(n) = f(n-1) + f(n-2)
	// Use two variables to store previous results
	prev2 := 1 // f(1)
	prev1 := 2 // f(2)

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}