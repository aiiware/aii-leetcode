package arrays

// LengthOfLongestSubstring solves LeetCode problem 0003: Longest Substring Without Repeating Characters
// Difficulty: Medium
// Tags: Hash Table, String, Sliding Window
//
// Given a string s, find the length of the longest substring without repeating characters.
//
// Time complexity: O(n), Space complexity: O(min(m,n)) where m is the size of the character set
func LengthOfLongestSubstring(s string) int {
	// Use a map to store the last seen index (rune index) of each character
	charIndex := make(map[rune]int)

	// Sliding window approach - left is the start of current window (rune index)
	left := 0
	maxLength := 0

	// Track rune index separately since range gives byte positions
	runeIdx := 0
	for _, char := range s {
		// If we've seen this character before and it's within our current window
		if prevIndex, exists := charIndex[char]; exists && prevIndex >= left {
			// Move the left pointer to the right of the previous occurrence
			left = prevIndex + 1
		}

		// Update the last seen index of the current character (store rune index)
		charIndex[char] = runeIdx

		// Update the maximum length
		currentLength := runeIdx - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		runeIdx++
	}

	return maxLength
}
