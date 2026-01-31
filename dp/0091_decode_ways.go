package dp


/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

// Problem 0091: Decode Ways
//
// A message containing letters from A-Z can be encoded into numbers using the following mapping:
// 'A' -> "1"
// 'B' -> "2"
// ...
// 'Z' -> "26"
//
// To decode an encoded message, all the digits must be grouped then mapped back into letters 
// using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:
// - "AAJF" with the grouping (1 1 10 6)
// - "KJF" with the grouping (11 10 6)
// Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".
//
// Given a string s containing only digits, return the number of ways to decode it.
// The test cases are generated so that the answer fits in a 32-bit integer.
//
// Example 1:
// Input: s = "12"
// Output: 2
// Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
//
// Example 2:
// Input: s = "226"
// Output: 3
// Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
//
// Example 3:
// Input: s = "06"
// Output: 0
// Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").
//
// Constraints:
// - 1 <= s.length <= 100
// - s contains only digits and may contain leading zero(s).

// numDecodings is the main solution function using dynamic programming.
// Time complexity: O(n), Space complexity: O(n)
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1 // Empty string has 1 way to decode
	dp[1] = 1 // First character (non-zero) has 1 way

	for i := 2; i <= n; i++ {
		// Check single digit
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		// Check two digits
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// numDecodingsOptimized uses optimized space O(1).
func numDecodingsOptimized(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	if n == 1 {
		return 1
	}

	// dp[i-2], dp[i-1], dp[i]
	// dp[0] = 1 (empty string)
	// dp[1] = 1 (first character, which we know is not '0')
	prev2, prev1 := 1, 1
	curr := 0

	for i := 2; i <= n; i++ {
		curr = 0

		// Check single digit
		if s[i-1] != '0' {
			curr += prev1
		}

		// Check two digits
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			curr += prev2
		}

		// Update for next iteration
		prev2, prev1 = prev1, curr
	}

	return curr
}

// numDecodingsDFS uses DFS with memoization.
func numDecodingsDFS(s string) int {
	if len(s) == 0 {
		return 0
	}

	memo := make(map[int]int)
	return dfsDecodings(s, 0, memo)
}

func dfsDecodings(s string, index int, memo map[int]int) int {
	// If we've reached the end, found a valid decoding
	if index == len(s) {
		return 1
	}

	// Check memo
	if val, exists := memo[index]; exists {
		return val
	}

	// If current digit is '0', cannot decode
	if s[index] == '0' {
		memo[index] = 0
		return 0
	}

	// Decode single digit
	ways := dfsDecodings(s, index+1, memo)

	// Decode two digits if possible
	if index+1 < len(s) {
		twoDigit := (s[index]-'0')*10 + (s[index+1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			ways += dfsDecodings(s, index+2, memo)
		}
	}

	memo[index] = ways
	return ways
}

// numDecodingsIterative uses iterative approach without full DP array.
func numDecodingsIterative(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	ways := make([]int, n)
	
	// Initialize first position
	if s[0] != '0' {
		ways[0] = 1
	}

	// Handle second position
	if n > 1 {
		// Single digit at position 1
		if s[1] != '0' {
			ways[1] += ways[0]
		}
		// Two digits from position 0
		twoDigit := (s[0]-'0')*10 + (s[1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			ways[1] += 1
		}
	}

	// Process rest of the string
	for i := 2; i < n; i++ {
		// Single digit
		if s[i] != '0' {
			ways[i] += ways[i-1]
		}

		// Two digits
		twoDigit := (s[i-1]-'0')*10 + (s[i] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			ways[i] += ways[i-2]
		}
	}

	return ways[n-1]
}

// numDecodingsRecursive is a pure recursive solution (exponential time).
func numDecodingsRecursive(s string) int {
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}

	// Decode single digit
	ways := numDecodingsRecursive(s[1:])

	// Decode two digits if possible
	if len(s) >= 2 {
		twoDigit := (s[0]-'0')*10 + (s[1] - '0')
		if twoDigit >= 10 && twoDigit <= 26 {
			ways += numDecodingsRecursive(s[2:])
		}
	}

	return ways
}

// numDecodingsDP2 is another DP formulation.
func numDecodingsDP2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[n] = 1 // Base case: empty suffix

	// Handle last character
	if s[n-1] != '0' {
		dp[n-1] = 1
	}

	// Process from right to left
	for i := n - 2; i >= 0; i-- {
		// Skip leading zeros
		if s[i] == '0' {
			continue
		}

		// Single digit
		dp[i] += dp[i+1]

		// Two digits
		twoDigit := (s[i]-'0')*10 + (s[i+1] - '0')
		if twoDigit <= 26 {
			dp[i] += dp[i+2]
		}
	}

	return dp[0]
}

// NumDecodings is the public interface function.
// It uses the optimized DP solution by default.
func NumDecodings(s string) int {
	return numDecodingsOptimized(s)
}