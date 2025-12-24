// 0032 - Longest Valid Parentheses
// https://leetcode.com/problems/longest-valid-parentheses/
// Hard - String, Dynamic Programming, Stack

package leetcode

// LongestValidParentheses finds the length of the longest valid parentheses substring.
// Time Complexity: O(n) where n is length of s
// Space Complexity: O(n) for stack or DP array
func LongestValidParentheses(s string) int {
	maxLen := 0
	stack := []int{-1} // Initialize with -1 as base

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i)
		} else {
			// Pop the last element
			stack = stack[:len(stack)-1]
			
			if len(stack) == 0 {
				// No matching '(' found, push current index as new base
				stack = append(stack, i)
			} else {
				// Calculate length of valid parentheses
				length := i - stack[len(stack)-1]
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}

// LongestValidParenthesesDP uses dynamic programming approach
func LongestValidParenthesesDP(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}

	dp := make([]int, n) // dp[i] = length of longest valid parentheses ending at i
	maxLen := 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// Case: "()"
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// Case: "))" with valid substring before
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}

	return maxLen
}

// LongestValidParenthesesTwoPass uses two-pointer approach
func LongestValidParenthesesTwoPass(s string) int {
	maxLen := 0
	
	// Left to right scan
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		
		if left == right {
			if 2*right > maxLen {
				maxLen = 2 * right
			}
		} else if right > left {
			left, right = 0, 0
		}
	}
	
	// Right to left scan
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		
		if left == right {
			if 2*left > maxLen {
				maxLen = 2 * left
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	
	return maxLen
}