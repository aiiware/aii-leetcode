# Graph Algorithms Category Overview

## Introduction to Graph Algorithms

Graph algorithms are fundamental to computer science, solving problems involving relationships between entities. A graph consists of vertices (nodes) and edges (connections) that can be directed or undirected, weighted or unweighted.

### Graph Representations
1. **Adjacency List**: Map from vertex to list of neighbors (space: O(V+E))
2. **Adjacency Matrix**: V×V matrix where entry (i,j) indicates edge (space: O(V²))
3. **Edge List**: Simple list of all edges (space: O(E))

## Graph Problem Patterns

### 1. Traversal Problems
**Characteristics**: Visit all vertices/nodes in some order
**Algorithms**: BFS, DFS
**Examples**:
- **0200 - Number of Islands**: Count connected components
- **0133 - Clone Graph**: Deep copy with traversal
- **0797 - All Paths From Source to Target**: Find all paths in DAG

### 2. Shortest Path Problems
**Characteristics**: Find minimum distance between vertices
**Algorithms**: BFS (unweighted), Dijkstra (non-negative weights), Bellman-Ford (negative weights)
**Examples**:
- **0743 - Network Delay Time**: Single-source shortest path
- **0399 - Evaluate Division**: Weighted path finding
- **0787 - Cheapest Flights Within K Stops**: Constrained shortest path

### 3. Cycle Detection
**Characteristics**: Detect cycles in directed/undirected graphs
**Algorithms**: DFS with coloring, Union-Find (undirected), Topological Sort (directed)
**Examples**:
- **0207 - Course Schedule**: Detect cycle in directed graph
- **0684 - Redundant Connection**: Find edge creating cycle

### 4. Connectivity Problems
**Characteristics**: Determine if graph is connected or find components
**Algorithms**: DFS, BFS, Union-Find
**Examples**:
- **0323 - Number of Connected Components**: Count components
- **0547 - Number of Provinces**: Find connected components
- **0695 - Max Area of Island**: Largest connected component

### 5. Topological Sorting
**Characteristics**: Linear ordering of vertices in DAG
**Algorithms**: Kahn's algorithm, DFS-based
**Examples**:
- **0210 - Course Schedule II**: Find valid course order
- **0269 - Alien Dictionary**: Reconstruct order from constraints

## Graph Algorithm Strategies

### Breadth-First Search (BFS)
```go
func BFS(start Node) {
    queue := []Node{start}
    visited := make(map[Node]bool)
    visited[start] = true
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        // Process current node
        for _, neighbor := range current.neighbors {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}
```

### Depth-First Search (DFS)
```go
func DFS(node Node, visited map[Node]bool) {
    visited[node] = true
    
    // Process node
    for _, neighbor := range node.neighbors {
        if !visited[neighbor] {
            DFS(neighbor, visited)
        }
    }
}
```

### Union-Find (Disjoint Set Union)
```go
type UnionFind struct {
    parent []int
    rank   []int
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX == rootY {
        return
    }
    
    // Union by rank
    if uf.rank[rootX] < uf.rank[rootY] {
        uf.parent[rootX] = rootY
    } else if uf.rank[rootX] > uf.rank[rootY] {
        uf.parent[rootY] = rootX
    } else {
        uf.parent[rootY] = rootX
        uf.rank[rootX]++
    }
}
```

## Complexity Analysis

### Time Complexity
- **BFS/DFS**: O(V + E) for adjacency list, O(V²) for adjacency matrix
- **Dijkstra**: O((V+E) log V) with binary heap, O(V²) with array
- **Topological Sort**: O(V + E)
- **Union-Find**: O(α(V)) amortized per operation (inverse Ackermann)

### Space Complexity
- **BFS/DFS**: O(V) for visited set and queue/stack
- **Dijkstra**: O(V) for distance array and priority queue
- **Union-Find**: O(V) for parent and rank arrays

## Common Graph Techniques

### 1. Multi-source BFS
Start BFS from multiple sources simultaneously.

### 2. Bidirectional BFS
Search from both start and goal simultaneously.

### 3. 0-1 BFS
For graphs with weights 0 or 1, use deque instead of priority queue.

### 4. Flood Fill
DFS/BFS for connected component labeling.

### 5. Eulerian Path/Circuit
Hierholzer's algorithm for finding Eulerian paths.

## Graph vs Other Approaches

### When to Use Graph Algorithms
- Problem involves relationships/connections between entities
- Need to find paths, cycles, or connectivity
- Data naturally forms a network structure
- Problem reduces to known graph problem

### When NOT to Use Graph Algorithms
- Data has tree structure (use tree algorithms)
- Problem is purely sequential (use DP or greedy)
- Graph would be too large to construct
- Simpler approach exists

## Learning Path

### Beginner Level
1. **0200 - Number of Islands**: Basic BFS/DFS traversal
2. **0695 - Max Area of Island**: Connected components
3. **0733 - Flood Fill**: Simple graph coloring

### Intermediate Level
1. **0207 - Course Schedule**: Cycle detection in directed graphs
2. **0210 - Course Schedule II**: Topological sort
3. **0323 - Number of Connected Components**: Union-Find
4. **0785 - Is Graph Bipartite?**: Graph coloring

### Advanced Level
1. **0743 - Network Delay Time**: Dijkstra's algorithm
2. **0399 - Evaluate Division**: Weighted graph traversal
3. **0329 - Longest Increasing Path in Matrix**: DP on DAG
4. **0684 - Redundant Connection**: Union-Find applications

## Practice Problems by Difficulty

### Easy
- 0733 - Flood Fill
- 0997 - Find the Town Judge
- 1971 - Find if Path Exists in Graph

### Medium
- 0200 - Number of Islands
- 0207 - Course Schedule
- 0210 - Course Schedule II
- 0323 - Number of Connected Components
- 0785 - Is Graph Bipartite?
- 0399 - Evaluate Division

### Hard
- 0743 - Network Delay Time
- 0329 - Longest Increasing Path in Matrix
- 0684 - Redundant Connection
- 0787 - Cheapest Flights Within K Stops

## Optimization Tips

### Graph Construction
- Use adjacency list for sparse graphs (most LeetCode problems)
- Consider implicit graph representation for grid problems
- Use integer indices instead of objects for performance

### Algorithm Choice
- BFS for shortest path in unweighted graphs
- DFS for connectivity and cycle detection
- Union-Find for dynamic connectivity
- Topological sort for dependency resolution

### Memory Optimization
- Use bitmask for visited set when V ≤ 64
- Reuse arrays instead of allocating new ones
- Use iterative DFS with stack to avoid recursion depth limits

## Common Mistakes

1. **Forgetting Visited Set**: Infinite loops in cycles
2. **Wrong Graph Representation**: Using matrix for sparse graphs
3. **Incorrect BFS/DFS Implementation**: Queue vs stack confusion
4. **Not Handling Disconnected Graphs**: Assuming single component
5. **Integer Overflow**: Large distances in weighted graphs

## Real-World Applications

1. **Social Networks**: Friend recommendations, community detection
2. **Transportation Networks**: Route planning, traffic optimization
3. **Web Crawling**: Link analysis, page ranking
4. **Circuit Design**: Signal routing, timing analysis
5. **Biology**: Protein interaction networks, phylogenetic trees
6. **Recommendation Systems**: Collaborative filtering, item similarity

## Special Graph Types

### 1. Trees
- Special case of graphs (acyclic, connected)
- Many optimized algorithms available
- Examples: Binary trees, n-ary trees, tries

### 2. Bipartite Graphs
- Vertices can be divided into two disjoint sets
- No edges within the same set
- Applications: Matching problems, scheduling

### 3. Directed Acyclic Graphs (DAGs)
- No directed cycles
- Can be topologically sorted
- Applications: Task scheduling, dependency resolution

### 4. Weighted Graphs
- Edges have weights/costs
- Algorithms: Dijkstra, Bellman-Ford, Floyd-Warshall
- Applications: Network routing, resource allocation

## Additional Resources

### Books
- "Introduction to Algorithms" (CLRS) - Graph algorithms
- "Algorithm Design Manual" - Practical graph algorithms
- "Graph Algorithms" by Neo4j - Applied graph theory

### Online Courses
- Stanford Algorithms Specialization (Coursera)
- MIT OpenCourseWare: Graph Theory
- Princeton Algorithms (Coursera)

### Visualization Tools
- VisuAlgo: Interactive algorithm visualizations
- Graph Online: Graph drawing and algorithm simulation
- NetworkX (Python): Graph analysis and visualization

### Practice Platforms
- LeetCode: Graph problems by frequency
- Codeforces: Graph theory contests
- TopCoder: Algorithm competitions with graph problems

---

*Last Updated: 2026-01-31*  
*Next: Create Design Patterns overview*