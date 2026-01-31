package strings

/*
159. Longest Substring with At Most Two Distinct Characters
https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/

Given a string s, return the length of the longest substring that contains at most two distinct characters.

Example 1:
Input: s = "eceba"
Output: 3
Explanation: The substring is "ece" which its length is 3.

Example 2:
Input: s = "ccaabbb"
Output: 5
Explanation: The substring is "aabbb" which its length is 5.

Constraints:
- 1 <= s.length <= 10^5
- s consists of English letters.

Difficulty: Medium
Tags: Hash Table, String, Sliding Window
Companies: Google, Amazon, Facebook, Microsoft, Apple
*/

// lengthOfLongestSubstringTwoDistinct is the main solution function
func lengthOfLongestSubstringTwoDistinct(s string) int {
	// Solution 1: Sliding window with hash map
	return lengthOfLongestSubstringTwoDistinctSlidingWindow(s)
}

// ===== Solution 1: Sliding window with hash map =====
// Time complexity: O(n) where n is the length of the string
// Space complexity: O(1) since we store at most 3 characters in the map

func lengthOfLongestSubstringTwoDistinctSlidingWindow(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	// Use a map to track character frequencies in the current window
	charCount := make(map[byte]int)
	left := 0
	maxLength := 0

	for right := 0; right < n; right++ {
		// Add current character to the window
		charCount[s[right]]++

		// If we have more than 2 distinct characters, shrink the window
		for len(charCount) > 2 {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		// Update max length
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// ===== Solution 2: Sliding window with optimized array (for ASCII) =====
// Time complexity: O(n)
// Space complexity: O(1) - fixed size array of 128

func lengthOfLongestSubstringTwoDistinctArray(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	// Use an array for ASCII characters (0-127)
	charCount := make([]int, 128)
	distinctCount := 0
	left := 0
	maxLength := 0

	for right := 0; right < n; right++ {
		// Add current character to the window
		if charCount[s[right]] == 0 {
			distinctCount++
		}
		charCount[s[right]]++

		// If we have more than 2 distinct characters, shrink the window
		for distinctCount > 2 {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				distinctCount--
			}
			left++
		}

		// Update max length
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// ===== Solution 3: Sliding window tracking last positions =====
// Time complexity: O(n)
// Space complexity: O(1) - stores at most 3 characters

func lengthOfLongestSubstringTwoDistinctLastPos(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}

	// Track the two most recent distinct characters and their last positions
	char1, char2 := byte(0), byte(0)
	lastPos1, lastPos2 := -1, -1
	left := 0
	maxLength := 0

	for i := 0; i < n; i++ {
		c := s[i]
		
		if char1 == 0 || c == char1 {
			char1 = c
			lastPos1 = i
		} else if char2 == 0 || c == char2 {
			char2 = c
			lastPos2 = i
		} else {
			// We found a third distinct character
			// Move left pointer to min(lastPos1, lastPos2) + 1
			if lastPos1 < lastPos2 {
				left = lastPos1 + 1
				char1 = char2
				lastPos1 = lastPos2
			} else {
				left = lastPos2 + 1
			}
			// The new character becomes the second character
			char2 = c
			lastPos2 = i
		}

		// Update max length
		currentLength := i - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}