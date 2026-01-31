package graphs

/*
695. Max Area of Island

You are given an m x n binary matrix grid. An island is a group of 1's (representing land) 
connected 4-directionally (horizontal or vertical). You may assume all four edges of the 
grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.

Example 1:
Input: grid = [
  [0,0,1,0,0,0,0,1,0,0,0,0,0],
  [0,0,0,0,0,0,0,1,1,1,0,0,0],
  [0,1,1,0,1,0,0,0,0,0,0,0,0],
  [0,1,0,0,1,1,0,0,1,0,1,0,0],
  [0,1,0,0,1,1,0,0,1,1,1,0,0],
  [0,0,0,0,0,0,0,0,0,0,1,0,0],
  [0,0,0,0,0,0,0,1,1,1,0,0,0],
  [0,0,0,0,0,0,0,1,1,0,0,0,0]
]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.

Constraints:
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 50
- grid[i][j] is either 0 or 1.
*/

/*
Difficulty: Medium
Tags: Array, Depth-First Search, Breadth-First Search, Union Find, Matrix
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber
*/

func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    maxArea := 0
    
    // DFS function to calculate area of an island
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        // Check bounds and if current cell is land
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
            return 0
        }
        
        // Mark current cell as visited by setting it to 0
        grid[i][j] = 0
        
        // Start with current cell
        area := 1
        
        // Visit all adjacent cells (up, down, left, right)
        area += dfs(i-1, j) // up
        area += dfs(i+1, j) // down
        area += dfs(i, j-1) // left
        area += dfs(i, j+1) // right
        
        return area
    }
    
    // Iterate through all cells
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                area := dfs(i, j)
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }
    
    return maxArea
}

// BFS version
func maxAreaOfIslandBFS(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    maxArea := 0
    
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                area := 0
                
                // BFS queue
                queue := [][]int{{i, j}}
                grid[i][j] = 0 // Mark as visited
                
                for len(queue) > 0 {
                    cell := queue[0]
                    queue = queue[1:]
                    area++
                    
                    // Check all four directions
                    for _, dir := range directions {
                        newI, newJ := cell[0]+dir[0], cell[1]+dir[1]
                        
                        if newI >= 0 && newI < m && newJ >= 0 && newJ < n && grid[newI][newJ] == 1 {
                            queue = append(queue, []int{newI, newJ})
                            grid[newI][newJ] = 0 // Mark as visited
                        }
                    }
                }
                
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }
    
    return maxArea
}

// Union-Find version
func maxAreaOfIslandUnionFind(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    
    // Initialize Union-Find
    parent := make([]int, m*n)
    size := make([]int, m*n)
    
    for i := 0; i < m*n; i++ {
        parent[i] = i
        size[i] = 1
    }
    
    // Find with path compression
    var find func(x int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    
    // Union
    union := func(x, y int) {
        rootX, rootY := find(x), find(y)
        if rootX != rootY {
            if size[rootX] < size[rootY] {
                parent[rootX] = rootY
                size[rootY] += size[rootX]
            } else {
                parent[rootY] = rootX
                size[rootX] += size[rootY]
            }
        }
    }
    
    // Process grid
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idx := i*n + j
                
                // Check right neighbor
                if j+1 < n && grid[i][j+1] == 1 {
                    union(idx, i*n+(j+1))
                }
                
                // Check down neighbor
                if i+1 < m && grid[i+1][j] == 1 {
                    union(idx, (i+1)*n+j)
                }
            }
        }
    }
    
    // Find maximum size
    maxArea := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                idx := i*n + j
                root := find(idx)
                if size[root] > maxArea {
                    maxArea = size[root]
                }
            }
        }
    }
    
    return maxArea
}