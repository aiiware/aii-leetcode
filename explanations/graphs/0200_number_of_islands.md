---
problem: "0200 - Number of Islands"
category: "graphs"
difficulty: "medium"
tags: ["dfs", "bfs", "matrix", "grid", "connected-components"]
complexity: "O(m×n) time, O(min(m,n)) space for BFS, O(m×n) space for DFS worst case"
leetcode_url: "https://leetcode.com/problems/number-of-islands/"
solution_file: "../graphs/0200_number_of_islands.go"
---

# 0200 - Number of Islands

## Problem Statement

Given an `m x n` 2D binary grid representing a map where `'1'` represents land and `'0'` represents water, count the number of islands. An island is formed by connecting adjacent lands horizontally or vertically (not diagonally). All four edges of the grid are surrounded by water.

**Example 1:**
```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

**Example 2:**
```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

**Constraints:**
- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`

## Approach Overview

This is a classic **connected components** problem in a grid. The practical approach involves treating the grid as an implicit graph where:
- Each cell is a node
- Adjacent cells (up, down, left, right) are connected edges
- We need to find all connected components of `'1'` cells

**Practical Applications:**
- **Image Processing**: Counting connected regions in binary images (object detection)
- **GIS/Mapping**: Counting land masses in satellite imagery
- **Game Development**: Flood fill algorithms in puzzle games
- **Network Analysis**: Finding connected clusters in network graphs

## Algorithm Walkthrough

### DFS (Depth-First Search) Approach

The DFS approach uses recursion to explore all connected land cells:

```go
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
```

**Step-by-step:**
1. **Initialize**: Check if grid is empty
2. **Iterate**: Scan through every cell in the grid
3. **Discover**: When we find a `'1'` (unvisited land), increment island count
4. **Explore**: Use DFS to mark all connected land cells as visited (set to `'0'`)
5. **Continue**: Move to next unvisited land cell

### BFS (Breadth-First Search) Approach

The BFS approach uses a queue for iterative exploration, which is more memory-efficient for certain grid shapes:

```go
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
```

## Complexity Analysis

### Time Complexity: **O(m × n)**
- We visit each cell at most once
- Each cell's processing (checking neighbors) is O(1)
- Total: m × n cells × O(1) = O(m × n)

### Space Complexity:
- **DFS**: O(m × n) in worst case (when entire grid is land and recursion depth equals total cells)
- **BFS**: O(min(m, n)) because queue stores perimeter of current island

**Trade-offs:**
- **DFS**: Simpler code, but risk of stack overflow for large connected components
- **BFS**: More memory-efficient for "long and thin" islands, avoids recursion depth issues

## Edge Cases & Pitfalls

### Common Mistakes:
1. **Not checking empty grid**: Always check `if len(grid) == 0 || len(grid[0]) == 0`
2. **Diagonal connections**: Islands only connect horizontally/vertically, not diagonally
3. **Modifying input**: The algorithm modifies the input grid (marks visited cells as `'0'`)
4. **Bounds checking**: Always check array bounds before accessing `grid[i][j]`

### Edge Cases to Test:
- **Empty grid**: `[]` or `[[]]` should return 0
- **Single cell**: `[["1"]]` returns 1, `[["0"]]` returns 0
- **All land**: Grid of all `'1'` should return 1
- **All water**: Grid of all `'0'` should return 0
- **Maximum constraints**: 300×300 grid (test performance)

## Real-World Applications

### 1. **Image Processing & Computer Vision**
- **Object counting**: Count distinct objects in binary images
- **Region segmentation**: Identify connected regions for analysis
- **Noise removal**: Filter out small connected components as noise

### 2. **Geographic Information Systems (GIS)**
- **Land mass counting**: Count islands in satellite imagery
- **Habitat analysis**: Identify connected forest regions
- **Urban planning**: Count building clusters in city maps

### 3. **Game Development**
- **Puzzle games**: Flood fill mechanics (like Minesweeper)
- **Terrain generation**: Identify isolated land masses
- **Pathfinding**: Preprocess map into connected regions

### 4. **Network Analysis**
- **Social networks**: Find connected user clusters
- **Computer networks**: Identify subnet clusters
- **Disease spread**: Model infection clusters

## Implementation Patterns

### Pattern 1: In-place Modification
```go
// Mark visited cells by modifying grid in-place
grid[i][j] = '0'  // Instead of using separate visited array
```
**Pros**: Saves memory, simpler code
**Cons**: Destructive (modifies input), not thread-safe

### Pattern 2: Direction Arrays
```go
// Define directions as offsets
directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// Use in loops for cleaner code
for _, dir := range directions {
    newI, newJ := i+dir[0], j+dir[1]
}
```

### Pattern 3: BFS vs DFS Choice
- **Use DFS when**: Code simplicity is priority, grid is small/medium
- **Use BFS when**: Memory efficiency matters, islands can be large
- **Alternative**: Union-Find for dynamic connectivity queries

## Related Problems

### Similar Grid Problems:
1. **695 - Max Area of Island**: Find largest connected component
2. **463 - Island Perimeter**: Calculate perimeter of islands
3. **130 - Surrounded Regions**: Mark regions surrounded by 'X'
4. **286 - Walls and Gates**: Multi-source BFS

### Graph Connectivity Problems:
1. **323 - Number of Connected Components in an Undirected Graph**
2. **547 - Number of Provinces** (Friend Circles)
3. **684 - Redundant Connection**
4. **1319 - Number of Operations to Make Network Connected**

### Advanced Variations:
1. **305 - Number of Islands II**: Dynamic islands (add land operations)
2. **694 - Number of Distinct Islands**: Consider island shapes
3. **711 - Number of Distinct Islands II**: With rotation/reflection

---

**Key Insight**: This problem teaches fundamental graph traversal on implicit graphs. Mastering it provides the foundation for solving many grid-based and connectivity problems in technical interviews and real-world applications.