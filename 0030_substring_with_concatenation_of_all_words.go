package leetcode

// FindSubstring solves LeetCode problem 0030: Substring with Concatenation of All Words
// Difficulty: Hard
// Tags: Hash Table, String, Sliding Window
//
// You are given a string s and an array of strings words. All the strings of words
// are of the same length. A concatenated substring in s is a substring that contains
// all the strings from words concatenated in any order, without any intervening characters.
// Return the starting indices of all the concatenated substrings in s.
//
// Time complexity: O(n * m * k) where n is length of s, m is number of words, k is word length
// Space complexity: O(m) for storing word counts
func FindSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	result := []int{}

	// Count frequency of each word
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Try each starting position
	for i := 0; i <= len(s)-len(words)*wordLen; i++ {
		seen := make(map[string]int)
		j := 0

		// Check each word in the substring
		for j < len(words) {
			start := i + j*wordLen
			end := start + wordLen
			if end > len(s) {
				break
			}

			word := s[start:end]
			if count, exists := wordCount[word]; exists {
				seen[word]++
				if seen[word] > count {
					break // Too many occurrences of this word
				}
			} else {
				break // Word not in dictionary
			}
			j++
		}

		// If we found all words
		if j == len(words) {
			result = append(result, i)
		}
	}

	return result
}

// FindSubstringOptimized uses sliding window with optimization
func FindSubstringOptimized(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	numWords := len(words)
	result := []int{}

	// Count frequency of each word
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// We only need to check wordLen different starting positions
	// because the pattern repeats every wordLen characters
	for start := 0; start < wordLen; start++ {
		left := start
		right := start
		seen := make(map[string]int)
		count := 0

		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			if _, exists := wordCount[word]; exists {
				seen[word]++
				count++

				// If we have too many of this word, move left pointer
				for seen[word] > wordCount[word] {
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					count--
					left += wordLen
				}

				// If we found all words
				if count == numWords {
					result = append(result, left)
					// Move left pointer one word forward
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					count--
					left += wordLen
				}
			} else {
				// Reset if we encounter a word not in dictionary
				seen = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}