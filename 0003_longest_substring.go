package leetcode

// LengthOfLongestSubstring solves LeetCode problem 0003: Longest Substring Without Repeating Characters
// Difficulty: Medium
// Tags: Hash Table, String, Sliding Window
//
// Given a string s, find the length of the longest substring without repeating characters.
//
// Time complexity: O(n), Space complexity: O(min(m, n)) where m is character set size
func LengthOfLongestSubstring(s string) int {
	return LengthOfLongestSubstringRune(s)
}

// LengthOfLongestSubstringRune Unicode-aware implementation using runes.
// Time complexity: O(n), Space complexity: O(min(m, n))
func LengthOfLongestSubstringRune(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Map to store last seen index of each character
	charIndex := make(map[rune]int)
	maxLength := 0
	start := 0
	runeIndex := 0 // Track rune index, not byte index

	// Iterate through string as runes
	for _, ch := range s {
		// If character seen before and within current window, move start
		if lastIdx, found := charIndex[ch]; found && lastIdx >= start {
			start = lastIdx + 1
		}

		// Update character's last seen index (using rune index)
		charIndex[ch] = runeIndex

		// Update max length
		currentLength := runeIndex - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		runeIndex++
	}

	return maxLength
}

// LengthOfLongestSubstringOptimized ASCII-optimized implementation using fixed array.
// Assumes string contains only ASCII characters (0-127).
// Time complexity: O(n), Space complexity: O(1) (fixed 128-element array)
func LengthOfLongestSubstringOptimized(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Array for ASCII characters (0-127)
	var charIndex [128]int
	for i := range charIndex {
		charIndex[i] = -1
	}

	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		// If character seen before and within current window, move start
		if charIndex[ch] >= start {
			start = charIndex[ch] + 1
		}

		// Update character's last seen index
		charIndex[ch] = i

		// Update max length
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}