# Advanced Graph Algorithms Reference Library Design

**Date**: January 30, 2026  
**Status**: Validated ✅  
**Author**: Aii Agent  
**Project**: aii-leetcode

## Overview

A comprehensive reference library of advanced graph algorithms in Go, designed as both a learning resource and production-ready toolkit for LeetCode problems and technical interviews.

## Goals

1. **Comprehensive Coverage**: 15+ advanced graph algorithms with complete implementations
2. **Educational Value**: Detailed documentation with complexity analysis and examples
3. **Production Ready**: Idiomatic Go code with comprehensive tests and benchmarks
4. **LeetCode Integration**: Direct mapping to real LeetCode problems and solution templates

## Architecture

### Core Components

#### 1. Graph Representation System
- **Primary Graph Struct** (`graph/graph.go`): Flexible adjacency list representation supporting directed/undirected, weighted/unweighted graphs
- **Edge Struct** (`graph/types.go`): Weighted edges with optional capacity/flow for network algorithms
- **Algorithm-Specific Interfaces**: `WeightedGraph`, `FlowNetwork` for type-safe algorithm composition

#### 2. Algorithm Categories (15+ implementations)

**A. Shortest Path Algorithms**
1. Dijkstra - Single-source, non-negative weights
2. Bellman-Ford - Single-source, handles negative weights  
3. Floyd-Warshall - All-pairs shortest path
4. A* - Heuristic search with custom cost functions

**B. Connectivity & Components**
5. Kosaraju's Algorithm - Strongly connected components (directed)
6. Tarjan's Algorithm - Bridges and articulation points
7. Union-Find - Connected components (undirected)

**C. Minimum Spanning Tree**
8. Kruskal's Algorithm - Union-Find based
9. Prim's Algorithm - Priority queue based

**D. Flow Networks**
10. Ford-Fulkerson - Maximum flow
11. Edmonds-Karp - BFS-based Ford-Fulkerson
12. Dinic's Algorithm - Layered networks

**E. Topological Sorting**
13. Kahn's Algorithm - BFS-based
14. DFS-based - Recursive with cycle detection

**F. Specialized Algorithms**
15. Eulerian Path/Circuit - Hierholzer's algorithm
16. Hamiltonian Path - Backtracking with pruning
17. Bipartite Matching - Hopcroft-Karp algorithm

## API Design

### Three API Patterns

1. **Functional API** (Pure functions):
```go
distances, err := graph.Dijkstra(graph, source)
```

2. **Method-based API** (Object-oriented):
```go
distances, err := g.Dijkstra(source)
```

3. **Builder Pattern** (For complex configurations):
```go
result, err := NewDijkstra(g, source).
    WithHeuristic(manhattanDistance).
    Run()
```

### Error Handling
Comprehensive `GraphError` type with error codes:
- `ErrNegativeCycle`, `ErrDisconnectedGraph`, `ErrCyclicGraph`
- `ErrInvalidVertex`, `ErrInvalidEdge`, `ErrNegativeWeight`

## Directory Structure

```
data_structures/graph/
├── graph.go              # Core Graph struct and basic operations
├── types.go              # Edge, Vertex, Error types
├── algorithms/           # Algorithm implementations
│   ├── dijkstra.go
│   ├── bellman_ford.go
│   ├── floyd_warshall.go
│   ├── union_find.go
│   ├── topological_sort.go
│   ├── kruskal.go
│   ├── prim.go
│   ├── tarjan.go
│   └── kosaraju.go
├── leetcode/             # LeetCode-specific utilities
│   ├── leetcode_parser.go
│   ├── solution_templates.go
│   └── problem_mapping.md
├── examples/             # Usage examples
│   ├── dijkstra_demo.go
│   ├── topological_sort_demo.go
│   └── leetcode_solutions.go
└── test/                 # Comprehensive tests
    ├── algorithms_test.go
    ├── leetcode_integration_test.go
    └── benchmarks_test.go
```

## LeetCode Integration

### Problem Mapping
- Dijkstra → #743 Network Delay Time, #1514 Path with Maximum Probability
- Bellman-Ford → #787 Cheapest Flights Within K Stops  
- Floyd-Warshall → #1334 Find the City With Smallest Number of Neighbors
- Union-Find → #547 Number of Provinces, #684 Redundant Connection
- Topological Sort → #207 Course Schedule, #210 Course Schedule II
- Kruskal/Prim → #1584 Min Cost to Connect All Points
- Tarjan's Algorithm → #1192 Critical Connections in a Network
- Kosaraju's Algorithm → #1557 Minimum Number of Vertices to Reach All Nodes

### Helper Utilities
- `ParseLeetCodeGraph()`: Convert LeetCode edge formats to Graph struct
- `ParseLeetCodeAdjList()`: Convert adjacency list format
- Solution templates for common problem patterns

## Implementation Plan

### Phase 1: Foundation (Week 1)
1. Core Graph struct with adjacency list representation
2. Dijkstra algorithm implementation
3. Union-Find data structure
4. Basic test infrastructure

### Phase 2: Core Algorithms (Week 2)
5. Bellman-Ford and Floyd-Warshall algorithms
6. Topological sort (Kahn's and DFS variants)
7. Kruskal and Prim's MST algorithms
8. LeetCode parser utilities

### Phase 3: Advanced Algorithms (Week 3)
9. Tarjan's algorithm for bridges/articulation points
10. Kosaraju's algorithm for strongly connected components
11. Flow algorithms (Ford-Fulkerson, Edmonds-Karp)
12. Comprehensive benchmark tests

### Phase 4: Polish & Integration (Week 4)
13. LeetCode solution templates and examples
14. Performance optimizations and memory pooling
15. Complete documentation with usage examples
16. Integration with existing test suite

## Success Metrics

1. **Completeness**: All 15+ algorithms implemented and tested
2. **Correctness**: 100% test coverage including edge cases
3. **Performance**: Benchmarks confirm optimal time/space complexity
4. **Usability**: Clear API documentation with practical examples
5. **Integration**: Seamless use with existing LeetCode solutions

## Technical Considerations

### Performance Optimizations
- Lazy initialization of adjacency matrix
- Memory pooling for BFS/DFS operations
- Custom heap implementations for priority queues
- Bitmask operations for small graphs (n ≤ 64)

### Testing Strategy
- Unit tests for each algorithm with multiple test cases
- Property-based testing using `testing/quick`
- Benchmark tests comparing against naive implementations
- Integration tests using real LeetCode problems

### Error Handling
- Comprehensive input validation
- Meaningful error messages with context
- Graceful degradation where appropriate

## Next Steps

1. **Ready for Implementation**: This design is validated and ready for implementation
2. **Git Worktree**: Use `git worktree` for isolated development
3. **Incremental Delivery**: Follow the phased implementation plan
4. **Continuous Validation**: Regular checkpoints with stakeholders

---
*Design validated through collaborative brainstorming session on January 30, 2026*