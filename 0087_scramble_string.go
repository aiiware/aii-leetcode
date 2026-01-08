package leetcode

// Problem 0087: Scramble String
//
// We can scramble a string s to get a string t using the following algorithm:
// 1. If the length of the string is 1, stop.
// 2. If the length of the string is > 1, do the following:
//    - Split the string into two non-empty substrings at a random index, i.e., 
//      if the string is s, divide it into x and y where s = x + y.
//    - Randomly decide to swap the two substrings or to keep them in the same order. 
//      i.e., after this step, s may become s = x + y or s = y + x.
//    - Apply step 1 recursively on each of the two substrings x and y.
//
// Given two strings s1 and s2 of the same length, return true if s2 is a scrambled 
// string of s1, otherwise return false.
//
// Example 1:
// Input: s1 = "great", s2 = "rgeat"
// Output: true
// Explanation: One possible scenario applied on s1 is:
// "great" --> "gr/eat" // divide at random index.
// "gr/eat" --> "gr/eat" // random decision is not to swap the two substrings.
// "gr/eat" --> "g/r" + "e/at" // apply recursively on both substrings. divide at random index each.
// "g/r" --> "r/g" // random decision was to swap the first substring and to keep the second substring in the same order.
// "e/at" --> "e/at" // random decision is not to swap the two substrings.
// "r/g" + "e/at" --> "rgeat" // after merging: "rgeat".
// "rgeat" is scrambled string of "great".
//
// Example 2:
// Input: s1 = "abcde", s2 = "caebd"
// Output: false
//
// Example 3:
// Input: s1 = "a", s2 = "a"
// Output: true
//
// Constraints:
// - s1.length == s2.length
// - 1 <= s1.length <= 30
// - s1 and s2 consist of lowercase English letters.

// isScramble is the main solution function using memoization.
// Time complexity: O(n^4), Space complexity: O(n^3)
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	if len(s1) == 1 {
		return s1 == s2
	}

	// Memoization map: key = s1_start|s2_start|length
	memo := make(map[[3]int]bool)
	return isScrambleMemo(s1, s2, 0, 0, len(s1), memo)
}

func isScrambleMemo(s1, s2 string, start1, start2, length int, memo map[[3]int]bool) bool {
	// Check memo
	key := [3]int{start1, start2, length}
	if val, exists := memo[key]; exists {
		return val
	}

	// Base case: strings of length 1
	if length == 1 {
		result := s1[start1] == s2[start2]
		memo[key] = result
		return result
	}

	// Check if character counts match (pruning)
	count1 := make([]int, 26)
	count2 := make([]int, 26)
	for i := 0; i < length; i++ {
		count1[s1[start1+i]-'a']++
		count2[s2[start2+i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if count1[i] != count2[i] {
			memo[key] = false
			return false
		}
	}

	// Try all possible split positions
	for i := 1; i < length; i++ {
		// Case 1: No swap - first i chars of s1 match first i chars of s2
		// and remaining chars match
		if isScrambleMemo(s1, s2, start1, start2, i, memo) &&
			isScrambleMemo(s1, s2, start1+i, start2+i, length-i, memo) {
			memo[key] = true
			return true
		}

		// Case 2: Swap - first i chars of s1 match last i chars of s2
		// and remaining chars match
		if isScrambleMemo(s1, s2, start1, start2+length-i, i, memo) &&
			isScrambleMemo(s1, s2, start1+i, start2, length-i, memo) {
			memo[key] = true
			return true
		}
	}

	memo[key] = false
	return false
}

// isScrambleDP is a dynamic programming solution.
// Time complexity: O(n^4), Space complexity: O(n^3)
func isScrambleDP(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	// dp[i][j][k] = whether s1[i:i+k] is a scramble of s2[j:j+k]
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
		}
	}

	// Initialize for length 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// Build up for longer lengths
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			for j := 0; j <= n-length; j++ {
				// Try all possible split positions
				for split := 1; split < length; split++ {
					// Case 1: No swap
					if dp[i][j][split] && dp[i+split][j+split][length-split] {
						dp[i][j][length] = true
						break
					}
					// Case 2: Swap
					if dp[i][j+length-split][split] && dp[i+split][j][length-split] {
						dp[i][j][length] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

// isScrambleOptimized is an optimized version with pruning.
// Time complexity: O(n^4) worst case, but much faster with pruning.
func isScrambleOptimized(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	memo := make(map[string]bool)
	return isScrambleHelper(s1, s2, memo)
}

func isScrambleHelper(s1, s2 string, memo map[string]bool) bool {
	// Check memo
	key := s1 + "|" + s2
	if _, exists := memo[key]; exists {
		return memo[key]
	}

	// Base cases
	if s1 == s2 {
		memo[key] = true
		return true
	}
	if len(s1) == 1 {
		result := s1 == s2
		memo[key] = result
		return result
	}

	// Quick check: character counts must match
	count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			memo[key] = false
			return false
		}
	}

	// Try all possible splits
	n := len(s1)
	for i := 1; i < n; i++ {
		// Case 1: No swap
		if isScrambleHelper(s1[:i], s2[:i], memo) &&
			isScrambleHelper(s1[i:], s2[i:], memo) {
			memo[key] = true
			return true
		}

		// Case 2: Swap
		if isScrambleHelper(s1[:i], s2[n-i:], memo) &&
			isScrambleHelper(s1[i:], s2[:n-i], memo) {
			memo[key] = true
			return true
		}
	}

	memo[key] = false
	return false
}

// isScrambleIterative is an iterative solution using stack.
// This is a simpler implementation that just uses the optimized version
// since implementing a proper iterative solution is complex.
func isScrambleIterative(s1 string, s2 string) bool {
	// For simplicity, just use the optimized version
	// A true iterative implementation would be complex and error-prone
	return isScrambleOptimized(s1, s2)
}

// IsScramble is the public interface function.
// It uses the optimized memoization solution by default.
func IsScramble(s1 string, s2 string) bool {
	return isScrambleOptimized(s1, s2)
}