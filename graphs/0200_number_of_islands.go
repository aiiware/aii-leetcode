package graphs

// 0200 - Number of Islands (Medium)
// Problem: Given an m x n 2D binary grid which represents a map of '1's (land) and '0's (water),
// return the number of islands. An island is surrounded by water and is formed by connecting
// adjacent lands horizontally or vertically. You may assume all four edges of the grid are
// surrounded by water.
//
// Example 1:
// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1
//
// Example 2:
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3
//
// Constraints:
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.

// numIslands solves using BFS approach
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	count := 0

	// Directions: up, down, left, right
	_ = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // used via directions variable

	// BFS helper function
	bfs := func(startRow, startCol int) {
		queue := [][2]int{{startRow, startCol}}
		grid[startRow][startCol] = '0' // Mark as visited

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			row, col := cell[0], cell[1]

			// Explore all 4 directions
			for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				newRow, newCol := row+dir[0], col+dir[1]

				// Check bounds and if it's land
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == '1' {
					grid[newRow][newCol] = '0' // Mark as visited
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}
	}

	// Iterate through all cells
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				bfs(i, j)
			}
		}
	}

	return count
}

// numIslandsBFS is an alias for numIslands (BFS approach)
func numIslandsBFS(grid [][]byte) int {
	return numIslands(grid)
}

// numIslandsDFS solves using DFS approach (recursive)
func numIslandsDFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	count := 0

	// DFS helper function
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// Check bounds and if it's land
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] != '1' {
			return
		}

		// Mark as visited
		grid[row][col] = '0'

		// Recursively visit all 4 directions
		dfs(row-1, col) // up
		dfs(row+1, col) // down
		dfs(row, col-1) // left
		dfs(row, col+1) // right
	}

	// Iterate through all cells
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}

// numIslandsUnionFind solves using Union-Find (Disjoint Set Union) approach
func numIslandsUnionFind(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// Initialize Union-Find structure
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	count := 0

	// Initialize: count all lands and set up Union-Find
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				idx := i*n + j
				parent[idx] = idx
				count++
			}
		}
	}

	// Find with path compression
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // Path compression
			x = parent[x]
		}
		return x
	}

	// Union by rank
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)

		if rootX != rootY {
			if rank[rootX] > rank[rootY] {
				parent[rootY] = rootX
			} else if rank[rootX] < rank[rootY] {
				parent[rootX] = rootY
			} else {
				parent[rootY] = rootX
				rank[rootX]++
			}
			count-- // Each union reduces island count by 1
		}
	}

	// Process the grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				idx := i*n + j

				// Check right neighbor
				if j+1 < n && grid[i][j+1] == '1' {
					union(idx, i*n+(j+1))
				}

				// Check down neighbor
				if i+1 < m && grid[i+1][j] == '1' {
					union(idx, (i+1)*n+j)
				}
			}
		}
	}

	return count
}
