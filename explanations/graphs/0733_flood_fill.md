# 0733 - Flood Fill

## Problem Statement
An image is represented by an m x n integer grid where `image[i][j]` represents the pixel value. Given starting pixel `(sr, sc)` and new color `color`, perform a flood fill operation.

**Flood Fill Rules:**
1. Start from pixel `(sr, sc)`
2. Change its color to `color`
3. For each pixel connected 4-directionally (up, down, left, right) that has the same original color, repeat the process

**Example:**
```
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
```

## Solution Approaches

### Approach 1: Breadth-First Search (BFS)
```go
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }
    
    originalColor := image[sr][sc]
    queue := [][2]int{{sr, sc}}
    directions := [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
    
    for len(queue) > 0 {
        r, c := queue[0][0], queue[0][1]
        queue = queue[1:]
        image[r][c] = color
        
        for _, dir := range directions {
            nr, nc := r+dir[0], c+dir[1]
            if nr >= 0 && nr < len(image) && nc >= 0 && nc < len(image[0]) && 
               image[nr][nc] == originalColor {
                queue = append(queue, [2]int{nr, nc})
            }
        }
    }
    return image
}
```

### Approach 2: Depth-First Search (DFS)
```go
func floodFillDFS(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color {
        return image
    }
    dfs(image, sr, sc, image[sr][sc], color)
    return image
}

func dfs(image [][]int, r, c, originalColor, newColor int) {
    if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || 
       image[r][c] != originalColor {
        return
    }
    
    image[r][c] = newColor
    dfs(image, r-1, c, originalColor, newColor) // up
    dfs(image, r+1, c, originalColor, newColor) // down
    dfs(image, r, c-1, originalColor, newColor) // left
    dfs(image, r, c+1, originalColor, newColor) // right
}
```

## Complexity Analysis

**Both BFS and DFS:**
- **Time Complexity:** O(m × n) in worst case where we fill the entire grid
- **Space Complexity:** 
  - BFS: O(min(m, n)) for queue (depends on implementation)
  - DFS: O(m × n) for recursion stack in worst case

## Key Insights

1. **Early Exit:** If starting color equals target color, return immediately
2. **Boundary Checking:** Always check array bounds before accessing
3. **Color Matching:** Only fill pixels with the original starting color
4. **4-directional Connectivity:** Only consider up, down, left, right (not diagonal)

## Edge Cases

1. **Single pixel image:** `[[5]]`, sr=0, sc=0, color=9 → `[[9]]`
2. **Same color:** If new color equals original color, do nothing
3. **Out of bounds:** Starting position outside image bounds
4. **Empty image:** Return empty image as-is
5. **Disconnected regions:** Only fill connected region with same color

## Real-World Applications

1. **Paint Bucket Tool:** In image editing software
2. **Region Detection:** In computer vision for object segmentation
3. **Game Development:** For filling connected areas in puzzle games
4. **Geographic Mapping:** For flood simulation or region coloring

## Optimization Considerations

1. **BFS vs DFS:** 
   - BFS uses queue, better for wide areas
   - DFS uses recursion/stack, simpler implementation
   - Both have same time complexity

2. **Memory Optimization:** 
   - For very large images, consider iterative DFS with explicit stack
   - Use bit manipulation for color representation if possible

3. **Parallel Processing:** 
   - Flood fill can be parallelized for large images
   - Divide image into regions and process concurrently

## Related Problems
- 0200 - Number of Islands (similar connectivity concept)
- 0130 - Surrounded Regions
- 0417 - Pacific Atlantic Water Flow
- 0695 - Max Area of Island

## Learning Points
- Graph traversal algorithms (BFS/DFS) on grid
- Connected component analysis
- Recursive vs iterative implementations
- Boundary condition handling