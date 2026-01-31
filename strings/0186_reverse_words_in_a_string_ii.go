package strings

// 0186 - Reverse Words in a String II
// https://leetcode.com/problems/reverse-words-in-a-string-ii/
// Difficulty: Medium
// Topics: Two Pointers, String
// Companies: Facebook, Amazon, Microsoft, Apple, Google, Bloomberg, Uber, Oracle, Adobe

/*
Description:
Given a character array s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by a single space.

Your code must solve the problem in-place, i.e. without allocating extra space.

Example 1:
Input: s = ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
Output: ["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]

Example 2:
Input: s = ["a"]
Output: ["a"]

Constraints:
- 1 <= s.length <= 10^5
- s[i] is an English letter (uppercase or lowercase), digit, or space ' '.
- There is at least one word in s.
- s does not contain leading or trailing spaces.
- All the words in s are guaranteed to be separated by a single space.
*/

// ReverseWordsII reverses the order of words in a character array in-place
func ReverseWordsII(s []byte) {
	n := len(s)
	if n <= 1 {
		return
	}

	// Step 1: Reverse the entire array
	reverseRange(s, 0, n-1)

	// Step 2: Reverse each word individually
	start := 0
	for i := 0; i <= n; i++ {
		// When we reach a space or the end of array, reverse the word
		if i == n || s[i] == ' ' {
			reverseRange(s, start, i-1)
			start = i + 1
		}
	}
}

// ReverseWordsIIAlternative is an alternative implementation that reverses words first, then the whole array
func ReverseWordsIIAlternative(s []byte) {
	n := len(s)
	if n <= 1 {
		return
	}

	// Step 1: Reverse each word individually
	start := 0
	for i := 0; i <= n; i++ {
		// When we reach a space or the end of array, reverse the word
		if i == n || s[i] == ' ' {
			reverseRange(s, start, i-1)
			start = i + 1
		}
	}

	// Step 2: Reverse the entire array
	reverseRange(s, 0, n-1)
}

// reverseRange reverses a portion of a byte slice from start to end (inclusive)
func reverseRange(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// ReverseWordsIIString is a helper function that takes a string and returns the reversed words string
// This is not in-place but useful for testing and comparison
func ReverseWordsIIString(str string) string {
	// Convert string to byte slice
	s := []byte(str)
	ReverseWordsII(s)
	return string(s)
}