package dp


/*
Difficulty: Hard
Tags: [Add relevant tags]
Companies: [Add company names]
*/

/*
# 0131 - Palindrome Partitioning
## Problem Description
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.

## Examples
Example 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:
Input: s = "a"
Output: [["a"]]

## Constraints
- 1 <= s.length <= 16
- s contains only lowercase English letters.

## Solution Approach
This problem can be solved using backtracking with memoization:
1. Use DFS to explore all possible partitions
2. At each step, check if the current substring is a palindrome
3. If it is, add it to the current partition and continue with the rest of the string
4. When we reach the end of the string, add the current partition to the result

Time Complexity: O(N * 2^N) where N is the length of the string
Space Complexity: O(N) for the recursion stack
*/

// Partition returns all possible palindrome partitions of s
func Partition(s string) [][]string {
	var result [][]string
	var current []string
	
	// Precompute palindrome table for O(1) palindrome checks
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}
	
	// Fill palindrome table using dynamic programming
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				isPalindrome[i][j] = true
			} else if i+1 == j {
				isPalindrome[i][j] = s[i] == s[j]
			} else {
				isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i+1][j-1]
			}
		}
	}
	
	var backtrack func(start int)
	backtrack = func(start int) {
		// If we've reached the end of the string, add current partition to result
		if start == n {
			// Make a copy of current partition
			temp := make([]string, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		
		// Try all possible end positions for the current substring
		for end := start; end < n; end++ {
			// If substring s[start:end+1] is a palindrome
			if isPalindrome[start][end] {
				// Add this palindrome to current partition
				current = append(current, s[start:end+1])
				// Continue with the rest of the string
				backtrack(end + 1)
				// Backtrack: remove last added palindrome
				current = current[:len(current)-1]
			}
		}
	}
	
	backtrack(0)
	return result
}