package strings

import (
	"unicode"
)

// 0151 - Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/
// Difficulty: Medium
// Topics: Two Pointers, String
// Companies: Google, Facebook, Amazon, Microsoft, Apple, Bloomberg, Uber, Oracle, Adobe, TikTok

/*
Description:
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. 
The returned string should only have a single space separating the words. 
Do not include any extra spaces.

Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Constraints:
- 1 <= s.length <= 10^4
- s contains English letters (upper-case and lower-case), digits, and spaces ' '.
- There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

// ReverseWords reverses the order of words in a string
func ReverseWords(s string) string {
	// Convert string to rune slice for easier manipulation
	runes := []rune(s)
	n := len(runes)
	
	// Step 1: Reverse the entire string
	reverseRunes(runes, 0, n-1)
	
	// Step 2: Reverse each word individually
	start := 0
	for start < n {
		// Skip whitespace (spaces, tabs, newlines, etc.)
		for start < n && unicode.IsSpace(runes[start]) {
			start++
		}
		
		if start >= n {
			break
		}
		
		// Find end of current word
		end := start
		for end < n && !unicode.IsSpace(runes[end]) {
			end++
		}
		
		// Reverse the current word
		reverseRunes(runes, start, end-1)
		
		// Move to next word
		start = end
	}
	
	// Step 3: Clean up extra spaces
	return cleanSpaces(runes)
}

// reverseRunes reverses a portion of a rune slice from start to end (inclusive)
func reverseRunes(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

// cleanSpaces removes extra spaces from the rune slice and returns a string
func cleanSpaces(runes []rune) string {
	n := len(runes)
	
	// Use two pointers: i for reading, j for writing
	i, j := 0, 0
	
	// Skip leading whitespace
	for i < n && unicode.IsSpace(runes[i]) {
		i++
	}
	
	// Process the rest of the string
	for i < n {
		// Copy non-space characters
		for i < n && !unicode.IsSpace(runes[i]) {
			runes[j] = runes[i]
			i++
			j++
		}
		
		// Skip extra whitespace
		for i < n && unicode.IsSpace(runes[i]) {
			i++
		}
		
		// Add a single space if there are more words
		if i < n {
			runes[j] = ' '
			j++
		}
	}
	
	// Return the cleaned string
	return string(runes[:j])
}

// ReverseWordsSimple is a simpler solution using built-in functions
func ReverseWordsSimple(s string) string {
	// Split the string into words (handles multiple spaces)
	words := make([]string, 0)
	
	start := 0
	n := len(s)
	
	for start < n {
		// Skip whitespace (spaces, tabs, newlines, etc.)
		for start < n && unicode.IsSpace(rune(s[start])) {
			start++
		}
		
		if start >= n {
			break
		}
		
		// Find end of current word
		end := start
		for end < n && !unicode.IsSpace(rune(s[end])) {
			end++
		}
		
		// Add word to list
		words = append(words, s[start:end])
		
		// Move to next word
		start = end
	}
	
	// Reverse the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	
	// Join words with single space
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	
	return result
}