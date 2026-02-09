package arrays

// 0150 - Evaluate Reverse Polish Notation
// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Difficulty: Medium
// Topics: Array, Math, Stack
// Companies: Amazon, Microsoft, Google, Apple, Facebook, Bloomberg, Uber, Oracle, Adobe, TikTok

/*
Description:
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression and return an integer that represents the value of the expression.

Note that:
- The valid operators are '+', '-', '*', and '/'.
- Each operand may be an integer or another expression.
- The division between two integers always truncates toward zero.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in a reverse polish notation.
- The answer and all the intermediate calculations can be represented in a 32-bit integer.

Example 1:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Constraints:
- 1 <= tokens.length <= 10^4
- tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
*/

import (
	"strconv"
)

// EvalRPN evaluates an arithmetic expression in Reverse Polish Notation
func EvalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		// Check if token is an operator
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// Pop two operands
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Perform operation
			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}

			// Push result back to stack
			stack = append(stack, result)
		} else {
			// Token is an operand, convert to integer
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	// Final result is the only element in stack
	return stack[0]
}
