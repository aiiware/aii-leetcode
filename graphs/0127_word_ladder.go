package graphs

// LadderLength solves LeetCode problem 0127: Word Ladder
// Difficulty: Hard
// Tags: Breadth-First Search, Hash Table, String
//
// A transformation sequence from word beginWord to word endWord using a dictionary wordList
// is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
// - Every adjacent pair of words differs by a single letter.
// - Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
// - sk == endWord
//
// Given two words, beginWord and endWord, and a dictionary wordList, return the number of words
// in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.
//
// Example 1:
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// Output: 5
// Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> "cog", which is 5 words long.
//
// Example 2:
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// Output: 0
// Explanation: The endWord "cog" is not in wordList, therefore no valid transformation sequence exists.
//
// Constraints:
// 1 <= beginWord.length <= 10
// beginWord != endWord
// 1 <= wordList.length <= 5000
// wordList[i].length == beginWord.length
// beginWord, endWord, and wordList[i] consist of lowercase English letters.
// All the words in wordList are unique.
//
// Time complexity: O(N * L^2) where N is number of words, L is word length
// Space complexity: O(N * L)
func LadderLength(beginWord string, endWord string, wordList []string) int {
	// Special case: beginWord equals endWord
	if beginWord == endWord {
		return 1
	}

	// Convert wordList to set for O(1) lookups
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// If endWord is not in wordList, return 0
	if !wordSet[endWord] {
		return 0
	}

	// Remove beginWord from set if present
	delete(wordSet, beginWord)

	// BFS queue
	queue := []string{beginWord}
	level := 1 // Start with beginWord

	for len(queue) > 0 {
		levelSize := len(queue)
		level++

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			chars := []byte(current)
			// Try changing each position
			for j := 0; j < len(chars); j++ {
				original := chars[j]
				// Try all possible characters
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == original {
						continue
					}
					chars[j] = byte(c)
					nextWord := string(chars)

					// If we found the endWord
					if nextWord == endWord {
						return level
					}

					// If nextWord is in wordSet
					if wordSet[nextWord] {
						queue = append(queue, nextWord)
						delete(wordSet, nextWord) // Mark as visited
					}
				}
				chars[j] = original
			}
		}
	}

	return 0
}