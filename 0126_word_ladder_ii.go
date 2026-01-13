package leetcode

// FindLadders solves LeetCode problem 0126: Word Ladder II
// Difficulty: Hard
// Tags: Breadth-First Search, Hash Table, String, Backtracking
//
// A transformation sequence from word beginWord to word endWord using a dictionary wordList
// is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
// - Every adjacent pair of words differs by a single letter.
// - Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
// - sk == endWord
//
// Given two words, beginWord and endWord, and a dictionary wordList, return all the shortest
// transformation sequences from beginWord to endWord, or an empty list if no such sequence exists.
// Each sequence should be returned as a list of the words [beginWord, s1, s2, ..., sk].
//
// Example 1:
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// Output: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
//
// Example 2:
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// Output: []
// Explanation: The endWord "cog" is not in wordList, therefore no valid transformation sequence exists.
//
// Constraints:
// 1 <= beginWord.length <= 5
// beginWord != endWord
// 1 <= wordList.length <= 500
// wordList[i].length == beginWord.length
// beginWord, endWord, and wordList[i] consist of lowercase English letters.
// All the words in wordList are unique.
//
// Time complexity: O(N * L^2) where N is number of words, L is word length
// Space complexity: O(N * L)
func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	// Convert wordList to set for O(1) lookups
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// If endWord is not in wordList, return empty result
	if !wordSet[endWord] {
		return [][]string{}
	}

	// Remove beginWord from set if present
	delete(wordSet, beginWord)

	// BFS to build graph and find shortest distances
	graph := make(map[string][]string)
	distance := make(map[string]int)
	queue := []string{beginWord}
	distance[beginWord] = 0

	found := false
	level := 0

	for len(queue) > 0 && !found {
		level++
		nextLevel := []string{}

		for _, word := range queue {
			chars := []byte(word)
			// Try changing each position
			for i := 0; i < len(chars); i++ {
				original := chars[i]
				// Try all possible characters
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == original {
						continue
					}
					chars[i] = byte(c)
					nextWord := string(chars)

					// If nextWord is the endWord, we found a path
					if nextWord == endWord {
						found = true
					}

					// If nextWord is in wordSet and (not visited or visited at same level)
					if wordSet[nextWord] {
						if _, visited := distance[nextWord]; !visited {
							distance[nextWord] = level
							nextLevel = append(nextLevel, nextWord)
							graph[word] = append(graph[word], nextWord)
						} else if distance[nextWord] == level {
							// If visited at same level, add to graph but not to nextLevel
							graph[word] = append(graph[word], nextWord)
						}
					}
				}
				chars[i] = original
			}
		}
		queue = nextLevel
	}

	// If endWord not found, return empty
	if !found {
		return [][]string{}
	}

	// DFS to find all shortest paths
	result := [][]string{}
	var dfs func(current string, path []string)
	dfs = func(current string, path []string) {
		path = append(path, current)

		if current == endWord {
			// Make a copy of the path
			newPath := make([]string, len(path))
			copy(newPath, path)
			result = append(result, newPath)
			return
		}

		for _, neighbor := range graph[current] {
			// Only explore if distance increases by 1 (ensures shortest path)
			if distance[neighbor] == distance[current]+1 {
				dfs(neighbor, path)
			}
		}
	}

	dfs(beginWord, []string{})
	return result
}