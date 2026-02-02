package strings

// IsIsomorphic solves LeetCode problem 0205: Isomorphic Strings
// Difficulty: Easy
// Tags: Hash Table, String
//
// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving
// the order of characters. No two characters may map to the same character, but a character
// may map to itself.
//
// Example 1:
// Input: s = "egg", t = "add"
// Output: true
// Explanation: e -> a, g -> d
//
// Example 2:
// Input: s = "foo", t = "bar"
// Output: false
// Explanation: f -> b, o -> a, o -> r (o maps to both a and r)
//
// Example 3:
// Input: s = "paper", t = "title"
// Output: true
// Explanation: p -> t, a -> i, p -> t, e -> l, r -> e
//
// Time complexity: O(n), Space complexity: O(1) (fixed 256 ASCII characters)
func IsIsomorphic(s string, t string) bool {
	// Quick check: different lengths cannot be isomorphic
	if len(s) != len(t) {
		return false
	}

	// Create two mapping arrays (ASCII has 256 characters)
	sToT := make([]byte, 256)
	tToS := make([]byte, 256)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		// Check if sChar already maps to a different tChar
		if sToT[sChar] != 0 && sToT[sChar] != tChar {
			return false
		}

		// Check if tChar already maps to a different sChar
		if tToS[tChar] != 0 && tToS[tChar] != sChar {
			return false
		}

		// Establish the mapping in both directions
		sToT[sChar] = tChar
		tToS[tChar] = sChar
	}

	return true
}

// IsIsomorphicMap solves the same problem using maps for clarity
// Time complexity: O(n), Space complexity: O(n)
func IsIsomorphicMap(s string, t string) bool {
	// Quick check: different lengths cannot be isomorphic
	if len(s) != len(t) {
		return false
	}

	// Create two maps for bidirectional mapping
	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		// Check if sChar already maps to a different tChar
		if mappedT, exists := sToT[sChar]; exists && mappedT != tChar {
			return false
		}

		// Check if tChar already maps to a different sChar
		if mappedS, exists := tToS[tChar]; exists && mappedS != sChar {
			return false
		}

		// Establish the mapping in both directions
		sToT[sChar] = tChar
		tToS[tChar] = sChar
	}

	return true
}

// IsIsomorphicPattern uses pattern matching approach
// Time complexity: O(n), Space complexity: O(n)
func IsIsomorphicPattern(s string, t string) bool {
	// Quick check: different lengths cannot be isomorphic
	if len(s) != len(t) {
		return false
	}

	// Create maps to store the first occurrence index of each character
	sFirstIndex := make(map[byte]int)
	tFirstIndex := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		// Get the first occurrence index of each character
		sIdx, sExists := sFirstIndex[sChar]
		tIdx, tExists := tFirstIndex[tChar]

		// If both characters have been seen before, check if their patterns match
		if sExists && tExists {
			if sIdx != tIdx {
				return false
			}
		} else if !sExists && !tExists {
			// If both characters are new, record their first occurrence
			sFirstIndex[sChar] = i
			tFirstIndex[tChar] = i
		} else {
			// One character is new while the other has been seen before
			return false
		}
	}

	return true
}