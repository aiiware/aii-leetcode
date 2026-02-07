// 0212 - Word Search II
// https://leetcode.com/problems/word-search-ii/
// Hard - Array, Backtracking, Trie

package arrays

// TrieNode represents a node in the Trie data structure
type TrieNode struct {
	children [26]*TrieNode
	word     string // Store the complete word at the end node
}

// WordSearchII finds all words from the given list that exist on the board
func WordSearchII(board [][]byte, words []string) []string {
	// Build Trie from words
	root := &TrieNode{}
	for _, word := range words {
		node := root
		for _, ch := range word {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = word
	}

	result := []string{}
	m, n := len(board), len(board[0])

	// DFS function to search words
	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		// Boundary check
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		// Get current character
		ch := board[i][j]
		if ch == '#' {
			return // Already visited in current path
		}

		// Check if character exists in Trie
		idx := ch - 'a'
		nextNode := node.children[idx]
		if nextNode == nil {
			return // No word with this prefix
		}

		// If we found a word, add to result
		if nextNode.word != "" {
			result = append(result, nextNode.word)
			nextNode.word = "" // Mark as found to avoid duplicates
		}

		// Mark as visited
		board[i][j] = '#'

		// Explore all four directions
		dfs(i+1, j, nextNode)
		dfs(i-1, j, nextNode)
		dfs(i, j+1, nextNode)
		dfs(i, j-1, nextNode)

		// Backtrack
		board[i][j] = ch
	}

	// Start DFS from every cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}

	return result
}

// WordSearchIIOptimized is an optimized version with pruning
func WordSearchIIOptimized(board [][]byte, words []string) []string {
	// Build Trie
	root := &TrieNode{}
	for _, word := range words {
		node := root
		for _, ch := range word {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = word
	}

	result := []string{}
	m, n := len(board), len(board[0])

	// Directions: down, up, right, left
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var backtrack func(i, j int, node *TrieNode)
	backtrack = func(i, j int, node *TrieNode) {
		ch := board[i][j]
		idx := ch - 'a'

		// Move to next node
		nextNode := node.children[idx]
		if nextNode == nil {
			return
		}

		// Check if we found a word
		if nextNode.word != "" {
			result = append(result, nextNode.word)
			nextNode.word = "" // Avoid duplicates
		}

		// Mark as visited
		board[i][j] = '#'

		// Explore neighbors
		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && board[ni][nj] != '#' {
				backtrack(ni, nj, nextNode)
			}
		}

		// Backtrack
		board[i][j] = ch

		// Optimization: prune empty leaf nodes
		hasChildren := false
		for _, child := range nextNode.children {
			if child != nil {
				hasChildren = true
				break
			}
		}
		if !hasChildren && nextNode.word == "" {
			node.children[idx] = nil
		}
	}

	// Start from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(i, j, root)
		}
	}

	return result
}