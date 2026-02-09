package arrays

// LengthOfLastWord solves LeetCode problem 0058: Length of Last Word
// Difficulty: Easy
// Tags: String
//
// Given a string s consisting of words and spaces, return the length of the last word in the string.
//
// A word is a maximal substring consisting of non-space characters only.
//
// Example 1:
// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
//
// Example 2:
// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
//
// Example 3:
// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.
//
// Time complexity: O(n), Space complexity: O(1)
func LengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	// Skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Count characters of the last word
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}
