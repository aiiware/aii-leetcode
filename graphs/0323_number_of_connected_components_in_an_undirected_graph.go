package graphs

/*
323. Number of Connected Components in an Undirected Graph

Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), 
write a function to find the number of connected components in an undirected graph.

Example 1:
Input: n = 5, edges = [[0,1], [1,2], [3,4]]
Output: 2

Example 2:
Input: n = 5, edges = [[0,1], [1,2], [2,3], [3,4]]
Output: 1

Constraints:
- 1 <= n <= 2000
- 1 <= edges.length <= 5000
- edges[i].length == 2
- 0 <= ai, bi < n
- ai != bi
- There are no self-loops or repeated edges.
*/

/*
Difficulty: Medium (LeetCode Premium)
Tags: Depth-First Search, Breadth-First Search, Union Find, Graph
Companies: Amazon, Facebook, Google, Microsoft, LinkedIn
*/

// countComponentsUnionFind uses Union-Find to count connected components
func countComponentsUnionFind(n int, edges [][]int) int {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
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
            parent[rootX] = rootY
        }
    }
    
    // Union all edges
    for _, edge := range edges {
        union(edge[0], edge[1])
    }
    
    // Count unique roots
    components := 0
    for i := 0; i < n; i++ {
        if find(i) == i {
            components++
        }
    }
    
    return components
}

// countComponentsDFS uses DFS to count connected components
func countComponentsDFS(n int, edges [][]int) int {
    // Build adjacency list
    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    
    visited := make([]bool, n)
    components := 0
    
    var dfs func(node int)
    dfs = func(node int) {
        visited[node] = true
        for _, neighbor := range adj[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }
    }
    
    for i := 0; i < n; i++ {
        if !visited[i] {
            components++
            dfs(i)
        }
    }
    
    return components
}

// countComponentsBFS uses BFS to count connected components
func countComponentsBFS(n int, edges [][]int) int {
    // Build adjacency list
    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    
    visited := make([]bool, n)
    components := 0
    
    for i := 0; i < n; i++ {
        if !visited[i] {
            components++
            
            // BFS
            queue := []int{i}
            visited[i] = true
            
            for len(queue) > 0 {
                current := queue[0]
                queue = queue[1:]
                
                for _, neighbor := range adj[current] {
                    if !visited[neighbor] {
                        visited[neighbor] = true
                        queue = append(queue, neighbor)
                    }
                }
            }
        }
    }
    
    return components
}