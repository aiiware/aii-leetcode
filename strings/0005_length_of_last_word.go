package strings

// LengthOfLastWord solves LeetCode problem 0005: Length of Last Word
// Difficulty: Easy
// Tags: String
//
// Given a string s consisting of words and spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.
//
// Time complexity: O(n), Space complexity: O(1)
func LengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	// Find the last word by scanning from right to left
	start := len(s) - 1
	
	// Skip trailing spaces
	for start >= 0 && s[start] == ' ' {
		start--
	}
	
	if start < 0 {
		return 0
	}
	
	// Count the length of the last word
	end := start
	for end >= 0 && s[end] != ' ' {
		end--
	}
	
	return start - end
}