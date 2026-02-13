package sliding_window

import "math"

/*
76. Minimum Window Substring
https://leetcode.com/problems/minimum-window-substring/

Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

Example 3:
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

Constraints:
- m == s.length
- n == t.length
- 1 <= m, n <= 10^5
- s and t consist of uppercase and lowercase English letters.

Follow up: Could you find an algorithm that runs in O(m + n) time?

Difficulty: Hard
Tags: Hash Table, String, Sliding Window
Companies: Facebook, Amazon, Microsoft, Google, Bloomberg, Uber, Apple, LinkedIn, Oracle, TikTok
*/

// MinWindow solves LeetCode problem 0076: Minimum Window Substring
func MinWindow(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	windowFreq := make(map[byte]int)
	required := len(tFreq)
	formed := 0
	left := 0
	minLength := math.MaxInt32
	minStart := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		windowFreq[char]++

		if _, ok := tFreq[char]; ok && windowFreq[char] == tFreq[char] {
			formed++
		}

		for left <= right && formed == required {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
				minStart = left
			}

			leftChar := s[left]
			windowFreq[leftChar]--
			if _, ok := tFreq[leftChar]; ok && windowFreq[leftChar] < tFreq[leftChar] {
				formed--
			}
			left++
		}
	}

	if minLength == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLength]
}


// MinWindowOptimized calls the main MinWindow function.
func MinWindowOptimized(s string, t string) string {
	return MinWindow(s, t)
}

// MinWindowSimplified calls the main MinWindow function.
func MinWindowSimplified(s string, t string) string {
	return MinWindow(s, t)
}