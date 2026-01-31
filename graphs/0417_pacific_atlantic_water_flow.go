package graphs

// 0417. Pacific Atlantic Water Flow
// https://leetcode.com/problems/pacific-atlantic-water-flow/

// PacificAtlantic finds cells where water can flow to both Pacific and Atlantic oceans
// Pacific ocean touches left and top edges, Atlantic ocean touches right and bottom edges
func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])
	
	// Create visited matrices for Pacific and Atlantic
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// Directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// DFS function to mark reachable cells
	var dfs func(i, j int, visited [][]bool, prevHeight int)
	dfs = func(i, j int, visited [][]bool, prevHeight int) {
		// Check bounds and if already visited or height is lower (water can't flow uphill)
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || heights[i][j] < prevHeight {
			return
		}
		
		visited[i][j] = true
		
		// Explore all four directions
		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1], visited, heights[i][j])
		}
	}

	// Start DFS from Pacific ocean edges (top and left)
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific, heights[i][0])
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific, heights[0][j])
	}

	// Start DFS from Atlantic ocean edges (bottom and right)
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic, heights[i][n-1])
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atlantic, heights[m-1][j])
	}

	// Find cells that can reach both oceans
	result := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// PacificAtlanticBFS finds cells using BFS approach
func PacificAtlanticBFS(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])
	
	// Create visited matrices
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// Directions
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// BFS function
	bfs := func(queue [][]int, visited [][]bool) {
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			i, j := cell[0], cell[1]
			
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && 
				   !visited[ni][nj] && heights[ni][nj] >= heights[i][j] {
					visited[ni][nj] = true
					queue = append(queue, []int{ni, nj})
				}
			}
		}
	}

	// Initialize queues for Pacific and Atlantic
	pacificQueue := [][]int{}
	atlanticQueue := [][]int{}

	// Add Pacific edges
	for i := 0; i < m; i++ {
		pacificQueue = append(pacificQueue, []int{i, 0})
		pacific[i][0] = true
	}
	for j := 0; j < n; j++ {
		pacificQueue = append(pacificQueue, []int{0, j})
		pacific[0][j] = true
	}

	// Add Atlantic edges
	for i := 0; i < m; i++ {
		atlanticQueue = append(atlanticQueue, []int{i, n-1})
		atlantic[i][n-1] = true
	}
	for j := 0; j < n; j++ {
		atlanticQueue = append(atlanticQueue, []int{m-1, j})
		atlantic[m-1][j] = true
	}

	// Run BFS
	bfs(pacificQueue, pacific)
	bfs(atlanticQueue, atlantic)

	// Find intersection
	result := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}