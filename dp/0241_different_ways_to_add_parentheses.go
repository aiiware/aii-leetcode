package dp

import (
	"strconv"
)

// DiffWaysToCompute solves LeetCode problem 0241: Different Ways to Add Parentheses
// Difficulty: Medium
// Tags: Divide and Conquer, Memoization, Dynamic Programming
//
// Given a string expression of numbers and operators, return all possible results
// from computing all the different possible ways to group numbers and operators.
// You may return the answer in any order.
//
// The test cases are generated such that the output values fit in a 32-bit integer
// and the number of different results does not exceed 10^4.
//
// Example 1:
// Input: expression = "2-1-1"
// Output: [0,2]
// Explanation:
// ((2-1)-1) = 0
// (2-(1-1)) = 2
//
// Example 2:
// Input: expression = "2*3-4*5"
// Output: [-34,-14,-10,-10,10]
// Explanation:
// (2*(3-(4*5))) = -34
// ((2*3)-(4*5)) = -14
// ((2*(3-4))*5) = -10
// (2*((3-4)*5)) = -10
// (((2*3)-4)*5) = 10
//
// Constraints:
// - 1 <= expression.length <= 20
// - expression consists of digits and the operators '+', '-', and '*'.
// - All the integer values in the expression are in the range [0, 99].
//
// Time complexity: O(Catalan numbers) ~ O(4^n / n^(3/2)), Space complexity: O(4^n / n^(3/2))
func DiffWaysToCompute(expression string) []int {
	// Memoization map
	memo := make(map[string][]int)
	return diffWaysToComputeMemo(expression, memo)
}

func diffWaysToComputeMemo(expression string, memo map[string][]int) []int {
	// Check memo
	if result, exists := memo[expression]; exists {
		return result
	}

	// Try to parse as a single number
	if num, err := strconv.Atoi(expression); err == nil {
		result := []int{num}
		memo[expression] = result
		return result
	}

	var results []int

	// Try every operator as the last operation
	for i := 0; i < len(expression); i++ {
		char := expression[i]
		if char == '+' || char == '-' || char == '*' {
			// Split into left and right parts
			left := expression[:i]
			right := expression[i+1:]

			// Compute all possible results for left and right
			leftResults := diffWaysToComputeMemo(left, memo)
			rightResults := diffWaysToComputeMemo(right, memo)

			// Combine results
			for _, leftVal := range leftResults {
				for _, rightVal := range rightResults {
					var result int
					switch char {
					case '+':
						result = leftVal + rightVal
					case '-':
						result = leftVal - rightVal
					case '*':
						result = leftVal * rightVal
					}
					results = append(results, result)
				}
			}
		}
	}

	memo[expression] = results
	return results
}