package strings

// ShortestPalindrome solves LeetCode problem 0214: Shortest Palindrome
// Difficulty: Hard
// Tags: String, KMP
//
// You are given a string s. You can convert s to a palindrome by adding characters in front of it.
// Return the shortest palindrome you can find by performing this transformation.
//
// Time complexity: O(n), Space complexity: O(n)
func ShortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Find the longest prefix that is also a suffix of reversed string
	// This is the longest palindromic prefix
	reversed := reverse(s)
	
	for i := 0; i < len(s); i++ {
		if s[:len(s)-i] == reversed[i:] {
			return reversed[:i] + s
		}
	}
	
	return reversed + s
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
