package dp

/*
0140. Word Break II
https://leetcode.com/problems/word-break-ii/
Difficulty: Hard

Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is a valid dictionary word.
Return all such possible sentences in any order.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
Output: ["cats and dog","cat sand dog"]

Example 2:
Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: []

Constraints:
1 <= s.length <= 20
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 10
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.
*/

func wordBreakII(s string, wordDict []string) []string {
	// Convert wordDict to a set for O(1) lookups
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// Memoization map: index -> list of sentences that can be formed from s[index:]
	memo := make(map[int][]string)

	// DFS with memoization
	var dfs func(int) []string
	dfs = func(start int) []string {
		// Check memo
		if sentences, exists := memo[start]; exists {
			return sentences
		}

		// Base case: reached end of string
		if start == len(s) {
			return []string{""}
		}

		result := []string{}
		// Try all possible end positions
		for end := start + 1; end <= len(s); end++ {
			word := s[start:end]
			if wordSet[word] {
				// Get all sentences from the remaining substring
				remainingSentences := dfs(end)
				for _, sentence := range remainingSentences {
					if sentence == "" {
						result = append(result, word)
					} else {
						result = append(result, word+" "+sentence)
					}
				}
			}
		}

		// Store in memo
		memo[start] = result
		return result
	}

	return dfs(0)
}