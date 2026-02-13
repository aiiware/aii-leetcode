package sliding_window

// 0438. Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
//
// Given two strings s and p, return an array of all the start indices of p's anagrams in s.
// You may return the answer in any order.
//
// Example 1:
// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
//
// Example 2:
// Input: s = "abab", p = "ab"
// Output: [0,1,2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".
//
// Constraints:
// - 1 <= s.length, p.length <= 3 * 10^4
// - s and p consist of lowercase English letters.
//
// Difficulty: Medium
// Tags: Hash Table, String, Sliding Window

// findAnagrams uses sliding window with frequency maps to find all anagram start indices.
// Time complexity: O(n + m), Space complexity: O(1) (fixed 26-element arrays)
func findAnagrams(s string, p string) []int {
	if len(p) == 0 {
		return []int{}
	}
	
	if len(s) < len(p) {
		return []int{}
	}

	result := []int{}
	pLen := len(p)
	sLen := len(s)

	// Frequency arrays for p and current window (26 lowercase letters)
	pFreq := [26]int{}
	windowFreq := [26]int{}

	// Initialize frequency for p and first window
	for i := 0; i < pLen; i++ {
		pFreq[p[i]-'a']++
		windowFreq[s[i]-'a']++
	}

	// Check first window
	if windowFreq == pFreq {
		result = append(result, 0)
	}

	// Slide the window
	for i := pLen; i < sLen; i++ {
		// Remove leftmost character
		windowFreq[s[i-pLen]-'a']--
		// Add new character
		windowFreq[s[i]-'a']++

		// Check if current window is an anagram
		if windowFreq == pFreq {
			result = append(result, i-pLen+1)
		}
	}

	return result
}

// findAnagramsMap uses maps instead of arrays for frequency counting.
// Time complexity: O(n + m), Space complexity: O(1) (max 26 entries)
func findAnagramsMap(s string, p string) []int {
	if len(p) == 0 {
		return []int{}
	}
	
	if len(s) < len(p) {
		return []int{}
	}

	result := []int{}
	pLen := len(p)
	sLen := len(s)

	// Frequency maps
	pFreq := make(map[byte]int)
	windowFreq := make(map[byte]int)

	// Initialize frequency for p
	for i := 0; i < pLen; i++ {
		pFreq[p[i]]++
	}

	// Initialize first window
	for i := 0; i < pLen; i++ {
		windowFreq[s[i]]++
	}

	// Helper function to compare maps
	mapsEqual := func(a, b map[byte]int) bool {
		if len(a) != len(b) {
			return false
		}
		for k, v := range a {
			if b[k] != v {
				return false
			}
		}
		return true
	}

	// Check first window
	if mapsEqual(windowFreq, pFreq) {
		result = append(result, 0)
	}

	// Slide the window
	for i := pLen; i < sLen; i++ {
		// Remove leftmost character
		leftChar := s[i-pLen]
		windowFreq[leftChar]--
		if windowFreq[leftChar] == 0 {
			delete(windowFreq, leftChar)
		}

		// Add new character
		windowFreq[s[i]]++

		// Check if current window is an anagram
		if mapsEqual(windowFreq, pFreq) {
			result = append(result, i-pLen+1)
		}
	}

	return result
}