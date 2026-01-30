package strings

// Exist solves LeetCode problem 0079: Word Search
// Difficulty: Medium
// Tags: Array, Backtracking, Depth-First Search, Matrix
//
// Given an m x n grid of characters board and a string word, return true if word
// exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cells,
// where adjacent cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
//
// Example 1:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true
//
// Example 2:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// Output: true
//
// Example 3:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// Output: false
//
// Constraints:
// m == board.length
// n == board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board and word consists of only lowercase and uppercase English letters.
//
// Time complexity: O(m * n * 4^L) where L is length of word
// Space complexity: O(L) for recursion stack
func Exist(board [][]byte, word string) bool {
	// Empty word always exists
	if len(word) == 0 {
		return true
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrackWordSearch(board, word, visited, i, j, 0) {
				return true
			}
		}
	}

	return false
}

// backtrackWordSearch is a helper function for DFS search
func backtrackWordSearch(board [][]byte, word string, visited [][]bool, i, j, index int) bool {
	// If we've matched all characters
	if index == len(word) {
		return true
	}

	// Check bounds and if cell is already visited
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
		return false
	}

	// Check if current cell matches current character
	if board[i][j] != word[index] {
		return false
	}

	// Mark cell as visited
	visited[i][j] = true

	// Explore all four directions
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if backtrackWordSearch(board, word, visited, ni, nj, index+1) {
			return true
		}
	}

	// Backtrack: unmark cell
	visited[i][j] = false
	return false
}

// ExistOptimized is an optimized version that uses board modification instead of visited matrix
func ExistOptimized(board [][]byte, word string) bool {
	// Empty word always exists
	if len(word) == 0 {
		return true
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrackOptimized(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

// backtrackOptimized modifies board in place to mark visited cells
func backtrackOptimized(board [][]byte, word string, i, j, index int) bool {
	// If we've matched all characters
	if index == len(word) {
		return true
	}

	// Check bounds
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}

	// Check if current cell matches current character
	if board[i][j] != word[index] {
		return false
	}

	// Temporarily mark cell as visited by changing its value
	// We use a character that won't appear in the word
	temp := board[i][j]
	board[i][j] = '#'

	// Explore all four directions
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if backtrackOptimized(board, word, ni, nj, index+1) {
			// Restore board before returning
			board[i][j] = temp
			return true
		}
	}

	// Backtrack: restore cell value
	board[i][j] = temp
	return false
}

// ExistBFS uses BFS approach (less efficient for this problem but included for completeness)
// Note: BFS is not ideal for path finding with visited constraints
func ExistBFS(board [][]byte, word string) bool {
	// Empty word always exists
	if len(word) == 0 {
		return true
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if bfs(board, word, i, j) {
					return true
				}
			}
		}
	}

	return false
}

// bfs performs breadth-first search
func bfs(board [][]byte, word string, startI, startJ int) bool {
	m, n := len(board), len(board[0])
	
	// Queue stores (i, j, index, visitedMask)
	// Since m*n <= 36, we can use bitmask to track visited cells
	queue := [][]int{{startI, startJ, 1, 1 << uint(startI*n+startJ)}}
	
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		
		i, j, index, visitedMask := curr[0], curr[1], curr[2], curr[3]
		
		// If we've matched all characters
		if index == len(word) {
			return true
		}
		
		// Explore neighbors
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range directions {
			ni, nj := i+dir[0], j+dir[1]
			
			// Check bounds
			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}
			
			// Check if cell matches next character and not visited
			pos := ni*n + nj
			if board[ni][nj] == word[index] && (visitedMask&(1<<uint(pos))) == 0 {
				newMask := visitedMask | (1 << uint(pos))
				queue = append(queue, []int{ni, nj, index + 1, newMask})
			}
		}
	}
	
	return false
}

// ExistEarlyPruning adds early pruning based on character frequency
func ExistEarlyPruning(board [][]byte, word string) bool {
	// Empty word always exists
	if len(word) == 0 {
		return true
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	
	// Early pruning: check if board has all characters in word
	boardCount := make(map[byte]int)
	wordCount := make(map[byte]int)
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			boardCount[board[i][j]]++
		}
	}
	
	for i := 0; i < len(word); i++ {
		wordCount[word[i]]++
	}
	
	// If board doesn't have enough of any character, return false early
	for ch, count := range wordCount {
		if boardCount[ch] < count {
			return false
		}
	}
	
	// Try starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrackOptimized(board, word, i, j, 0) {
				return true
			}
		}
	}
	
	return false
}

// ExistDirectional tries to optimize search direction
func ExistDirectional(board [][]byte, word string) bool {
	// Empty word always exists
	if len(word) == 0 {
		return true
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	
	// Find all possible starting positions
	starts := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				starts = append(starts, []int{i, j})
			}
		}
	}
	
	// Try each starting position
	for _, start := range starts {
		if backtrackOptimized(board, word, start[0], start[1], 0) {
			return true
		}
	}
	
	return false
}