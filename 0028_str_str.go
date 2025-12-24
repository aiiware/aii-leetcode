package leetcode

// StrStr solves LeetCode problem 0028: Find the Index of the First Occurrence in a String
// Difficulty: Easy
// Tags: String, Two Pointers, String Matching
//
// Given two strings needle and haystack, return the index of the first occurrence
// of needle in haystack, or -1 if needle is not part of haystack.
//
// Time complexity: O(n*m) worst case, Space complexity: O(1)
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	n := len(haystack)
	m := len(needle)

	if m > n {
		return -1
	}

	// Simple brute force approach
	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && haystack[i+j] == needle[j] {
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}

// StrStrKMP uses Knuth-Morris-Pratt algorithm for efficient string matching
func StrStrKMP(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	n := len(haystack)
	m := len(needle)

	if m > n {
		return -1
	}

	// Build the prefix table (also called failure function or LPS - Longest Prefix Suffix)
	lps := buildLPS(needle)

	i := 0 // index for haystack
	j := 0 // index for needle

	for i < n {
		if haystack[i] == needle[j] {
			i++
			j++
		}

		if j == m {
			return i - j // found match
		} else if i < n && haystack[i] != needle[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

// buildLPS builds the Longest Prefix Suffix table for KMP algorithm
func buildLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0 // length of the previous longest prefix suffix
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}