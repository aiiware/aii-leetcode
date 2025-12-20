package leetcode

// LongestCommonPrefix solves LeetCode problem 0014: Longest Common Prefix
// Difficulty: Easy
// Tags: String, Trie
//
// Write a function to find the longest common prefix string amongst an array
// of strings. If there is no common prefix, return an empty string "".
//
// A common prefix is a prefix that is common to all strings in the array.
//
// Example:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Time complexity: O(n*m) where n is number of strings and m is length of shortest
// Space complexity: O(1)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	// Horizontal scanning approach
	// Start with first string as prefix and reduce it
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		// Reduce prefix length while it's not a prefix of current string
		for len(prefix) > 0 && !hasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		// If prefix becomes empty, no common prefix exists
		if len(prefix) == 0 {
			return ""
		}
	}

	return prefix
}

// Helper function to check if string has given prefix
func hasPrefix(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}

// Alternative implementation using vertical scanning
func LongestCommonPrefixVertical(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Check each character position
	for charIdx := 0; charIdx < len(strs[0]); charIdx++ {
		// Compare this character in all strings
		for strIdx := 1; strIdx < len(strs); strIdx++ {
			// If current string is shorter or character doesn't match
			if charIdx >= len(strs[strIdx]) || strs[strIdx][charIdx] != strs[0][charIdx] {
				return strs[0][:charIdx]
			}
		}
	}

	return strs[0]
}
