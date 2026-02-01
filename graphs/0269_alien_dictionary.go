package graphs

// 0269 - Alien Dictionary (Hard)
// Problem: Given a sorted dictionary of an alien language having N words and k starting alphabets of standard dictionary,
// find the order of characters in the alien language.
// Input: words = ["wrt","wrf","er","ett","rftt"], k = 4
// Output: "wertf"
// Explanation: From "wrt" and "wrf", we can get 't' < 'f'
// From "wrf" and "er", we can get 'w' < 'e'
// From "er" and "ett", we can get 'r' < 't'
// From "ett" and "rftt", we can get 'e' < 'r'
// So one possible order is "wertf"

func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Build adjacency list and indegree map
	adj := make(map[byte][]byte)
	indegree := make(map[byte]int)

	// Initialize indegree for all unique characters
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			ch := word[i]
			if _, exists := indegree[ch]; !exists {
				indegree[ch] = 0
			}
		}
	}

	// Build graph by comparing adjacent words
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		// Check for invalid case: word2 is prefix of word1
		if len(word1) > len(word2) && word1[:len(word2)] == word2 {
			return ""
		}

		// Find first differing character
		minLen := len(word1)
		if len(word2) < minLen {
			minLen = len(word2)
		}

		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				// Add edge from word1[j] to word2[j]
				u := word1[j]
				v := word2[j]

				// Check if edge already exists
				exists := false
				for _, neighbor := range adj[u] {
					if neighbor == v {
						exists = true
						break
					}
				}

				if !exists {
					adj[u] = append(adj[u], v)
					indegree[v]++
				}
				break
			}
		}
	}

	// If there's only one word, return all unique characters in any order
	if len(words) == 1 {
		result := make([]byte, 0, len(indegree))
		for ch := range indegree {
			result = append(result, ch)
		}
		return string(result)
	}

	// Perform topological sort using Kahn's algorithm
	queue := make([]byte, 0)
	for ch, deg := range indegree {
		if deg == 0 {
			queue = append(queue, ch)
		}
	}

	result := make([]byte, 0)
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		result = append(result, ch)

		for _, neighbor := range adj[ch] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check for cycle (if not all characters are in result)
	if len(result) != len(indegree) {
		return ""
	}

	return string(result)
}