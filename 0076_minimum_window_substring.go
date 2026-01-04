package leetcode

// MinWindow solves LeetCode problem 0076: Minimum Window Substring
// Difficulty: Hard
// Tags: Hash Table, String, Sliding Window
//
// Given two strings s and t, return the minimum window substring of s such that
// every character in t (including duplicates) is included in the window.
// If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.
//
// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
//
// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.
//
// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.
//
// Constraints:
// m == s.length
// n == t.length
// 1 <= m, n <= 10^5
// s and t consist of uppercase and lowercase English letters.
//
// Time complexity: O(m + n), Space complexity: O(1) (fixed size character set)
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Frequency map for characters in t
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Frequency map for current window in s
	windowFreq := make(map[byte]int)

	// Variables to track the sliding window
	left, right := 0, 0
	required := len(tFreq) // number of unique characters we need to match
	formed := 0            // number of unique characters currently matched
	minLength := len(s) + 1
	minLeft := 0

	// Expand the window by moving right pointer
	for right < len(s) {
		char := s[right]
		windowFreq[char]++

		// Check if this character completes a requirement from t
		if count, needed := tFreq[char]; needed && windowFreq[char] == count {
			formed++
		}

		// Try to contract the window from the left while it's valid
		for left <= right && formed == required {
			char = s[left]

			// Update minimum window
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
				minLeft = left
			}

			// Remove left character from window
			windowFreq[char]--
			if count, needed := tFreq[char]; needed && windowFreq[char] < count {
				formed--
			}

			left++
		}

		right++
	}

	if minLength == len(s)+1 {
		return ""
	}

	return s[minLeft : minLeft+minLength]
}

// MinWindowOptimized is an optimized version using arrays instead of maps
// for better performance with ASCII characters
func MinWindowOptimized(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Use arrays for frequency counting (ASCII characters)
	tFreq := [128]int{}
	windowFreq := [128]int{}

	// Count frequencies in t
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Count required unique characters
	required := 0
	for i := 0; i < 128; i++ {
		if tFreq[i] > 0 {
			required++
		}
	}

	left, right := 0, 0
	formed := 0
	minLength := len(s) + 1
	minLeft := 0

	// Expand window
	for right < len(s) {
		char := s[right]
		windowFreq[char]++

		// Check if this character completes a requirement
		if tFreq[char] > 0 && windowFreq[char] == tFreq[char] {
			formed++
		}

		// Contract window while valid
		for left <= right && formed == required {
			char = s[left]

			// Update minimum window
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
				minLeft = left
			}

			// Remove left character
			windowFreq[char]--
			if tFreq[char] > 0 && windowFreq[char] < tFreq[char] {
				formed--
			}

			left++
		}

		right++
	}

	if minLength == len(s)+1 {
		return ""
	}

	return s[minLeft : minLeft+minLength]
}

// MinWindowSimplified is a simplified version that's easier to understand
// but may be slightly less efficient
func MinWindowSimplified(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Create frequency map for t
	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	// Variables for sliding window
	left, right := 0, 0
	minStart, minLen := 0, len(s)+1
	required := len(t) // total characters needed (including duplicates)

	// Expand window
	for right < len(s) {
		// Add current character to window
		char := s[right]

		// If this character is in t, decrement required count
		if tFreq[char] > 0 {
			required--
		}
		tFreq[char]--

		// When we have all required characters, try to shrink window
		for required == 0 {
			// Update minimum window
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
				minStart = left
			}

			// Remove left character from window
			leftChar := s[left]
			tFreq[leftChar]++
			if tFreq[leftChar] > 0 {
				required++
			}
			left++
		}

		right++
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[minStart : minStart+minLen]
}