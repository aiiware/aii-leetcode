package graphs

/*
200. Number of Islands

Given an m x n 2D binary grid which represents a map of '1's (land) and '0's (water), 
return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally 
or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 300
- grid[i][j] is '0' or '1'.
*/

/*
Difficulty: Medium
Tags: Array, Depth-First Search, Breadth-First Search, Union Find, Matrix
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    count := 0
    
    // DFS function to mark all connected land cells
    var dfs func(i, j int)
    dfs = func(i, j int) {
        // Check bounds and if current cell is land
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
            return
        }
        
        // Mark current cell as visited by setting it to '0'
        grid[i][j] = '0'
        
        // Visit all adjacent cells (up, down, left, right)
        dfs(i-1, j) // up
        dfs(i+1, j) // down
        dfs(i, j-1) // left
        dfs(i, j+1) // right
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

// Alternative BFS solution
func numIslandsBFS(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    count := 0
    
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                count++
                
                // BFS queue
                queue := [][]int{{i, j}}
                grid[i][j] = '0' // Mark as visited
                
                for len(queue) > 0 {
                    cell := queue[0]
                    queue = queue[1:]
                    
                    // Check all four directions
                    for _, dir := range directions {
                        newI, newJ := cell[0]+dir[0], cell[1]+dir[1]
                        
                        if newI >= 0 && newI < m && newJ >= 0 && newJ < n && grid[newI][newJ] == '1' {
                            queue = append(queue, []int{newI, newJ})
                            grid[newI][newJ] = '0' // Mark as visited
                        }
                    }
                }
            }
        }
    }
    
    return count
}