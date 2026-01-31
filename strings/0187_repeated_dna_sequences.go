package strings

// 0187 - Repeated DNA Sequences
// https://leetcode.com/problems/repeated-dna-sequences/
// Difficulty: Medium
// Topics: Hash Table, String, Bit Manipulation, Sliding Window, Rolling Hash
// Companies: Amazon, Google, LinkedIn

/*
Description:
All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T', for example: "ACGAATTCCG". 
When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

Example 1:
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]

Example 2:
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]

Constraints:
- 1 <= s.length <= 10^5
- s[i] is either 'A', 'C', 'G', or 'T'.
*/

// FindRepeatedDnaSequences finds all 10-letter-long sequences that occur more than once
func FindRepeatedDnaSequences(s string) []string {
	// Solution 1: Hash map with sliding window (most straightforward)
	return findRepeatedDnaSequencesHashMap(s)
}

// ===== Solution 1: Hash Map with Sliding Window =====
// Time complexity: O(n) where n is the length of the string
// Space complexity: O(n) for storing substrings in the map

func findRepeatedDnaSequencesHashMap(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}

	// Use a map to count occurrences of each 10-letter sequence
	sequenceCount := make(map[string]int)
	result := make([]string, 0)

	// Slide a window of size 10 through the string
	for i := 0; i <= n-10; i++ {
		substr := s[i : i+10]
		sequenceCount[substr]++

		// Add to result only when we see it for the second time
		// This avoids duplicates in the result
		if sequenceCount[substr] == 2 {
			result = append(result, substr)
		}
	}

	return result
}

// ===== Solution 2: Optimized with Two Hash Sets =====
// Time complexity: O(n)
// Space complexity: O(n)
// More memory efficient for very large inputs

func findRepeatedDnaSequencesTwoSets(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}

	seen := make(map[string]bool)
	repeated := make(map[string]bool)
	result := make([]string, 0)

	for i := 0; i <= n-10; i++ {
		substr := s[i : i+10]

		// If we've seen it before, add to repeated set
		if seen[substr] {
			repeated[substr] = true
		} else {
			seen[substr] = true
		}
	}

	// Convert repeated set to slice
	for seq := range repeated {
		result = append(result, seq)
	}

	return result
}

// ===== Solution 3: Bit Manipulation (Optimized Space) =====
// Time complexity: O(n)
// Space complexity: O(n) but more efficient since we store integers instead of strings
// This solution encodes DNA sequences as 2-bit values:
// A = 00, C = 01, G = 10, T = 11
// A 10-letter sequence can be stored in 20 bits (fits in a 32-bit integer)

func findRepeatedDnaSequencesBitManipulation(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}

	// Map characters to 2-bit values
	charToBits := map[byte]int{
		'A': 0, // 00
		'C': 1, // 01
		'G': 2, // 10
		'T': 3, // 11
	}

	// First, build the initial 10-letter sequence hash
	var hash int
	for i := 0; i < 10; i++ {
		hash = (hash << 2) | charToBits[s[i]]
	}

	// Use maps to track seen and repeated sequences
	seen := make(map[int]bool)
	repeated := make(map[int]bool)
	sequenceMap := make(map[int]string) // Map hash back to string for result

	// Store the first sequence
	seen[hash] = true
	sequenceMap[hash] = s[:10]

	// Slide the window
	for i := 10; i < n; i++ {
		// Remove the leftmost 2 bits (20 bits total, we keep 18 bits from previous)
		hash = (hash << 2) & 0xFFFFF // 0xFFFFF = 20 bits mask (2^20 - 1)
		// Add the new character
		hash |= charToBits[s[i]]

		seq := s[i-9 : i+1] // Current 10-letter sequence

		if seen[hash] {
			repeated[hash] = true
		} else {
			seen[hash] = true
			sequenceMap[hash] = seq
		}
	}

	// Convert repeated hashes to strings
	result := make([]string, 0, len(repeated))
	for hash := range repeated {
		result = append(result, sequenceMap[hash])
	}

	return result
}

// ===== Solution 4: Rolling Hash with Large Prime =====
// Time complexity: O(n)
// Space complexity: O(n)
// Uses polynomial rolling hash for better distribution

func findRepeatedDnaSequencesRollingHash(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}

	// Map characters to values 1-4 (avoid 0 for better hash distribution)
	charToVal := map[byte]int{
		'A': 1,
		'C': 2,
		'G': 3,
		'T': 4,
	}

	// Constants for rolling hash
	const base = 5    // base > number of characters (4)
	const mod = 1e9 + 7 // large prime

	// Precompute powers of base
	pow := make([]int, 11)
	pow[0] = 1
	for i := 1; i <= 10; i++ {
		pow[i] = (pow[i-1] * base) % mod
	}

	// Compute hash of first 10-letter sequence
	var hash int
	for i := 0; i < 10; i++ {
		hash = (hash*base + charToVal[s[i]]) % mod
	}

	// Track seen and repeated sequences
	seen := make(map[int]bool)
	repeated := make(map[int]bool)
	sequenceMap := make(map[int]string)

	seen[hash] = true
	sequenceMap[hash] = s[:10]

	// Slide the window using rolling hash
	for i := 10; i < n; i++ {
		// Remove leftmost character
		leftCharVal := charToVal[s[i-10]]
		hash = (hash - leftCharVal*pow[9]%mod + mod) % mod
		// Add new character
		hash = (hash*base + charToVal[s[i]]) % mod

		seq := s[i-9 : i+1]

		if seen[hash] {
			repeated[hash] = true
		} else {
			seen[hash] = true
			sequenceMap[hash] = seq
		}
	}

	// Convert to result
	result := make([]string, 0, len(repeated))
	for hash := range repeated {
		result = append(result, sequenceMap[hash])
	}

	return result
}

// Helper function to validate DNA string
func isValidDnaString(s string) bool {
	for _, ch := range s {
		if ch != 'A' && ch != 'C' && ch != 'G' && ch != 'T' {
			return false
		}
	}
	return true
}