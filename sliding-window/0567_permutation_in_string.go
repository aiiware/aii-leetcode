package sliding_window

// 0567. Permutation in String
// https://leetcode.com/problems/permutation-in-string/
//
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
//
// Example 1:
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
//
// Example 2:
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false
//
// Constraints:
// - 1 <= s1.length, s2.length <= 10^4
// - s1 and s2 consist of lowercase English letters.
//
// Difficulty: Medium
// Tags: Hash Table, Two Pointers, String, Sliding Window

// checkInclusion uses sliding window with frequency arrays to check if s2 contains a permutation of s1.
// Time complexity: O(n + m), Space complexity: O(1) (fixed 26-element arrays)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Len := len(s1)
	s2Len := len(s2)

	// Frequency arrays for s1 and current window (26 lowercase letters)
	s1Freq := [26]int{}
	windowFreq := [26]int{}

	// Initialize frequency for s1 and first window
	for i := 0; i < s1Len; i++ {
		s1Freq[s1[i]-'a']++
		windowFreq[s2[i]-'a']++
	}

	// Check first window
	if windowFreq == s1Freq {
		return true
	}

	// Slide the window
	for i := s1Len; i < s2Len; i++ {
		// Remove leftmost character
		windowFreq[s2[i-s1Len]-'a']--
		// Add new character
		windowFreq[s2[i]-'a']++

		// Check if current window is a permutation
		if windowFreq == s1Freq {
			return true
		}
	}

	return false
}

// checkInclusionOptimized uses optimized sliding window with match count.
// Time complexity: O(n + m), Space complexity: O(1) (fixed 26-element arrays)
func checkInclusionOptimized(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Len := len(s1)
	s2Len := len(s2)

	// Frequency arrays
	s1Freq := [26]int{}
	s2Freq := [26]int{}

	// Initialize frequencies for first window
	for i := 0; i < s1Len; i++ {
		s1Freq[s1[i]-'a']++
		s2Freq[s2[i]-'a']++
	}

	// Count matching frequencies
	matches := 0
	for i := 0; i < 26; i++ {
		if s1Freq[i] == s2Freq[i] {
			matches++
		}
	}

	// Check first window
	if matches == 26 {
		return true
	}

	// Slide the window
	for i := s1Len; i < s2Len; i++ {
		// Remove leftmost character
		leftChar := s2[i-s1Len] - 'a'
		s2Freq[leftChar]--
		if s2Freq[leftChar] == s1Freq[leftChar] {
			matches++
		} else if s2Freq[leftChar]+1 == s1Freq[leftChar] {
			matches--
		}

		// Add new character
		rightChar := s2[i] - 'a'
		s2Freq[rightChar]++
		if s2Freq[rightChar] == s1Freq[rightChar] {
			matches++
		} else if s2Freq[rightChar]-1 == s1Freq[rightChar] {
			matches--
		}

		// Check if all frequencies match
		if matches == 26 {
			return true
		}
	}

	return false
}