package graphs

/*
399. Evaluate Division

You are given an array of variable pairs equations and an array of real numbers values, 
where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. 
Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query 
where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not 
result in division by zero and that there is no contradiction.

Example 1:
Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], 
       queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation: 
Given: a / b = 2.0, b / c = 3.0
Queries are: 
a / c = (a / b) * (b / c) = 2.0 * 3.0 = 6.0
b / a = 1 / (a / b) = 1 / 2.0 = 0.5
a / e = -1.0 (no equation for e)
a / a = 1.0
x / x = -1.0 (no equation for x)

Constraints:
- 1 <= equations.length <= 20
- equations[i].length == 2
- 1 <= Ai.length, Bi.length <= 5
- values.length == equations.length
- 0.0 < values[i] <= 20.0
- 1 <= queries.length <= 20
- queries[i].length == 2
- 1 <= Cj.length, Dj.length <= 5
- Ai, Bi, Cj, Dj consist of lower case English letters and digits.
*/

/*
Difficulty: Medium
Tags: Array, Depth-First Search, Breadth-First Search, Union Find, Graph
Companies: Amazon, Facebook, Google, Bloomberg, Uber, LinkedIn
*/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    // Build graph: node -> map[neighbor]value
    graph := make(map[string]map[string]float64)
    
    // Initialize graph with equations
    for i, eq := range equations {
        a, b := eq[0], eq[1]
        val := values[i]
        
        if graph[a] == nil {
            graph[a] = make(map[string]float64)
        }
        if graph[b] == nil {
            graph[b] = make(map[string]float64)
        }
        
        graph[a][b] = val
        graph[b][a] = 1.0 / val
    }
    
    // DFS function to find path from start to end
    var dfs func(start, end string, visited map[string]bool) float64
    dfs = func(start, end string, visited map[string]bool) float64 {
        // If either node doesn't exist
        if _, ok := graph[start]; !ok {
            return -1.0
        }
        if _, ok := graph[end]; !ok {
            return -1.0
        }
        
        // If we found the end node
        if start == end {
            return 1.0
        }
        
        visited[start] = true
        
        // Explore all neighbors
        for neighbor, weight := range graph[start] {
            if !visited[neighbor] {
                result := dfs(neighbor, end, visited)
                if result != -1.0 {
                    return weight * result
                }
            }
        }
        
        return -1.0
    }
    
    // Process each query
    results := make([]float64, len(queries))
    for i, query := range queries {
        start, end := query[0], query[1]
        visited := make(map[string]bool)
        results[i] = dfs(start, end, visited)
    }
    
    return results
}

// Union-Find with weights version
func calcEquationUnionFind(equations [][]string, values []float64, queries [][]string) []float64 {
    parent := make(map[string]string)
    weight := make(map[string]float64)
    
    // Find with path compression
    var find func(x string) (string, float64)
    find = func(x string) (string, float64) {
        if _, exists := parent[x]; !exists {
            parent[x] = x
            weight[x] = 1.0
            return x, 1.0
        }
        
        if parent[x] == x {
            return x, weight[x]
        }
        
        root, rootWeight := find(parent[x])
        parent[x] = root
        weight[x] = weight[x] * rootWeight
        return root, weight[x]
    }
    
    // Union
    union := func(x, y string, value float64) {
        rootX, weightX := find(x)
        rootY, weightY := find(y)
        
        if rootX != rootY {
            parent[rootX] = rootY
            // weight[x] * weight[rootX] = value * weight[y]
            // So weight[rootX] = value * weight[y] / weight[x]
            weight[rootX] = value * weightY / weightX
        }
    }
    
    // Build Union-Find structure from equations
    for i, eq := range equations {
        a, b := eq[0], eq[1]
        union(a, b, values[i])
    }
    
    // Process queries
    results := make([]float64, len(queries))
    for i, query := range queries {
        a, b := query[0], query[1]
        
        // Check if both variables exist
        if _, existsA := parent[a]; !existsA {
            results[i] = -1.0
            continue
        }
        if _, existsB := parent[b]; !existsB {
            results[i] = -1.0
            continue
        }
        
        rootA, weightA := find(a)
        rootB, weightB := find(b)
        
        if rootA != rootB {
            results[i] = -1.0
        } else {
            // a / b = weightA / weightB
            results[i] = weightA / weightB
        }
    }
    
    return results
}