package leetcode

// IsValid solves LeetCode problem 0020: Valid Parentheses
// Difficulty: Easy
// Tags: String, Stack
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// 3. Every close bracket has a corresponding open bracket of the same type.
//
// Time complexity: O(n), Space complexity: O(n)
func IsValid(s string) bool {
	// Stack to keep track of opening brackets
	stack := []rune{}

	// Map to match closing brackets with opening brackets
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if _, isClosing := bracketMap[char]; isClosing {
			// Current character is a closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		} else {
			// Current character is an opening bracket
			stack = append(stack, char)
		}
	}

	// If stack is not empty, there are unmatched opening brackets
	return len(stack) == 0
}
