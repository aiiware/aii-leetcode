package strings

import "sort"

// IsAnagram solves LeetCode problem 0242: Valid Anagram
// Difficulty: Easy
// Tags: Hash Table, String, Sorting
//
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different
// word or phrase, typically using all the original letters exactly once.
//
// Time complexity: O(n), Space complexity: O(1) (fixed 52 characters for A-Z and a-z)
func IsAnagram(s string, t string) bool {
	// Quick check: different lengths cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Count characters in both strings (52 slots: 26 for A-Z, 26 for a-z)
	var countS, countT [52]int

	for i := 0; i < len(s); i++ {
		// Handle uppercase A-Z
		if s[i] >= 'A' && s[i] <= 'Z' {
			countS[s[i]-'A']++
		} else if s[i] >= 'a' && s[i] <= 'z' {
			// Handle lowercase a-z (offset by 26)
			countS[26+s[i]-'a']++
		}
		// Skip other characters
		
		if t[i] >= 'A' && t[i] <= 'Z' {
			countT[t[i]-'A']++
		} else if t[i] >= 'a' && t[i] <= 'z' {
			countT[26+t[i]-'a']++
		}
	}

	// Compare character counts
	for i := 0; i < 52; i++ {
		if countS[i] != countT[i] {
			return false
		}
	}

	return true
}

// IsAnagramSorting solves the same problem using sorting
// Time complexity: O(n log n), Space complexity: O(n)
func IsAnagramSorting(s string, t string) bool {
	// Quick check: different lengths cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Convert strings to rune slices for sorting
	sRunes := []rune(s)
	tRunes := []rune(t)

	// Sort both slices
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	// Compare sorted strings
	return string(sRunes) == string(tRunes)
}

// IsAnagramUnicode handles Unicode characters (beyond a-z)
// Time complexity: O(n), Space complexity: O(n)
func IsAnagramUnicode(s string, t string) bool {
	// Quick check: different lengths cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Use map to count Unicode characters
	count := make(map[rune]int)

	// Count characters in s
	for _, ch := range s {
		count[ch]++
	}

	// Subtract counts using t
	for _, ch := range t {
		count[ch]--
		// If count goes negative, t has more of this character than s
		if count[ch] < 0 {
			return false
		}
	}

	// Check all counts are zero
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}