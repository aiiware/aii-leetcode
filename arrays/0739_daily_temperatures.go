package arrays

// DailyTemperatures solves LeetCode problem 0739: Daily Temperatures
// Difficulty: Medium
// Tags: Array, Stack, Monotonic Stack
//
// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have to
// wait after the ith day to get a warmer temperature. If there is no future day
// for which this is possible, put 0 instead.
//
// Time complexity: O(n), Space complexity: O(n)
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	if n == 0 {
		return []int{}
	}

	// Initialize result array with zeros
	result := make([]int, n)

	// Use a stack to store indices of temperatures
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		// While stack is not empty and current temperature is greater than
		// temperature at index on top of stack
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// Pop the index from stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Calculate the number of days to wait
			result[top] = i - top
		}

		// Push current index to stack
		stack = append(stack, i)
	}

	// Remaining indices in stack have no warmer temperature ahead
	// They already have 0 in result array

	return result
}
