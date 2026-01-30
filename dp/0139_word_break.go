package dp

/*
# 0139 - Word Break
## Problem Description
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

## Examples
Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false

## Constraints
- 1 <= s.length <= 300
- 1 <= wordDict.length <= 1000
- 1 <= wordDict[i].length <= 20
- s and wordDict[i] consist of only lowercase English letters.
- All the strings of wordDict are unique.

## Solution Approach
This problem can be solved using dynamic programming:
- dp[i] = true if s[0:i] can be segmented into dictionary words
- dp[0] = true (empty string can be segmented)
- For each position i from 1 to n:
  - For each word in wordDict:
    - If word matches the substring ending at i and dp[i-len(word)] is true,
      then dp[i] = true

Time Complexity: O(N * M * L) where N is length of s, M is number of words, L is average word length
Space Complexity: O(N) for DP array
*/

// WordBreak returns true if s can be segmented into dictionary words
func WordBreak(s string, wordDict []string) bool {
	n := len(s)
	
	// Convert wordDict to a set for O(1) lookups
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	
	// dp[i] = true if s[0:i] can be segmented
	dp := make([]bool, n+1)
	dp[0] = true // Empty string can be segmented
	
	// For each position in the string
	for i := 1; i <= n; i++ {
		// Check all possible substrings ending at i
		for j := 0; j < i; j++ {
			// If prefix s[0:j] can be segmented and s[j:i] is in dictionary
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // No need to check other j values
			}
		}
	}
	
	return dp[n]
}

// WordBreakOptimized is optimized by checking only up to max word length
func WordBreakOptimized(s string, wordDict []string) bool {
	n := len(s)
	
	// Convert wordDict to a set and find max word length
	wordSet := make(map[string]bool)
	maxWordLen := 0
	for _, word := range wordDict {
		wordSet[word] = true
		if len(word) > maxWordLen {
			maxWordLen = len(word)
		}
	}
	
	// dp[i] = true if s[0:i] can be segmented
	dp := make([]bool, n+1)
	dp[0] = true
	
	// For each position in the string
	for i := 1; i <= n; i++ {
		// Only check substrings up to max word length
		start := max(0, i-maxWordLen)
		for j := start; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	
	return dp[n]
}