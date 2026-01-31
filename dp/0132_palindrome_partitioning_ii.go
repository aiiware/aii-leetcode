package dp


/*
Difficulty: Hard
Tags: [Add relevant tags]
Companies: [Add company names]
*/

/*
# 0132 - Palindrome Partitioning II
## Problem Description
Given a string s, partition s such that every substring of the partition is a palindrome.
Return the minimum cuts needed for a palindrome partitioning of s.

## Examples
Example 1:
Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

Example 2:
Input: s = "a"
Output: 0

Example 3:
Input: s = "ab"
Output: 1

## Constraints
- 1 <= s.length <= 2000
- s consists of lowercase English letters only.

## Solution Approach
This problem can be solved using dynamic programming:
1. First, precompute a palindrome table to check if any substring is a palindrome in O(1)
2. Then use DP to find the minimum cuts:
   - dp[i] = minimum cuts needed for substring s[0:i]
   - For each position i, check all j <= i where s[j:i] is a palindrome
   - Update dp[i] = min(dp[i], dp[j-1] + 1) for j > 0, or 0 if j == 0

Time Complexity: O(N^2) where N is the length of the string
Space Complexity: O(N^2) for the palindrome table
*/

// MinCut returns the minimum cuts needed for palindrome partitioning
func MinCut(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	
	// Precompute palindrome table
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
	
	// dp[i] = minimum cuts needed for substring s[0:i+1]
	dp := make([]int, n)
	
	for i := 0; i < n; i++ {
		// Initialize with maximum possible cuts (cut after each character)
		dp[i] = i
		
		// Check if the entire substring s[0:i+1] is a palindrome
		if isPalindrome[0][i] {
			dp[i] = 0
			continue
		}
		
		// Try all possible partition points
		for j := 1; j <= i; j++ {
			// If substring s[j:i+1] is a palindrome
			if isPalindrome[j][i] {
				// Update minimum cuts
				if dp[j-1]+1 < dp[i] {
					dp[i] = dp[j-1] + 1
				}
			}
		}
	}
	
	return dp[n-1]
}