package graphs

/*
261. Graph Valid Tree

Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), 
write a function to check whether these edges make up a valid tree.

Example 1:
Input: n = 5, edges = [[0,1], [0,2], [0,3], [1,4]]
Output: true

Example 2:
Input: n = 5, edges = [[0,1], [1,2], [2,3], [1,3], [1,4]]
Output: false

Constraints:
- 1 <= n <= 2000
- 0 <= edges.length <= 5000
- edges[i].length == 2
- 0 <= ai, bi < n
- ai != bi
- There are no self-loops or repeated edges.
*/

/*
Difficulty: Medium (LeetCode Premium)
Tags: Depth-First Search, Breadth-First Search, Union Find, Graph
Companies: Amazon, Facebook, Google, LinkedIn
*/

// validTreeUnionFind uses Union-Find to check for cycles and connectivity
func validTreeUnionFind(n int, edges [][]int) bool {
    // For a graph to be a valid tree:
    // 1. It must have exactly n-1 edges
    // 2. It must be connected (no isolated components)
    // 3. It must have no cycles
    
    if len(edges) != n-1 {
        return false
    }
    
    // Initialize Union-Find
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
    union := func(x, y int) bool {
        rootX, rootY := find(x), find(y)
        if rootX == rootY {
            return false // Cycle detected
        }
        parent[rootX] = rootY
        return true
    }
    
    // Process all edges
    for _, edge := range edges {
        if !union(edge[0], edge[1]) {
            return false // Cycle found
        }
    }
    
    return true
}

// validTreeDFS uses DFS to check for cycles and connectivity
func validTreeDFS(n int, edges [][]int) bool {
    if len(edges) != n-1 {
        return false
    }
    
    // Build adjacency list
    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    
    visited := make([]bool, n)
    
    var hasCycle func(node, parent int) bool
    hasCycle = func(node, parent int) bool {
        visited[node] = true
        
        for _, neighbor := range adj[node] {
            if !visited[neighbor] {
                if hasCycle(neighbor, node) {
                    return true
                }
            } else if neighbor != parent {
                // If neighbor is visited and not parent, we have a cycle
                return true
            }
        }
        
        return false
    }
    
    // Check for cycles starting from node 0
    if hasCycle(0, -1) {
        return false
    }
    
    // Check if all nodes are connected
    for i := 0; i < n; i++ {
        if !visited[i] {
            return false
        }
    }
    
    return true
}