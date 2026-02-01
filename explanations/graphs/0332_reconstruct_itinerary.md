# 0332 Reconstruct Itinerary

## Problem Statement
Given a list of airline tickets where `tickets[i] = [from_i, to_i]`, reconstruct the itinerary in order starting from "JFK". Use all tickets exactly once. If multiple valid itineraries exist, return the one with smallest lexical order.

## Key Insights
1. **Eulerian Path**: This is finding an Eulerian path in a directed graph
2. **Hierholzer's Algorithm**: Efficient O(E) algorithm for Eulerian paths
3. **Lexical Order**: Need to visit destinations in sorted order
4. **All Tickets Used**: Must be a valid Eulerian path using all edges

## Algorithm: Hierholzer's Algorithm with Lexical Order

### Step-by-Step
1. **Build adjacency list** with destinations sorted lexicographically
2. **Use DFS (post-order traversal)**:
   - From current airport, visit smallest lexical destination first
   - Recursively explore that path
   - Add airport to itinerary after exploring all outgoing edges (post-order)
3. **Reverse result** (post-order gives reverse Eulerian path)

### Why Post-Order Works
- Ensures we only add node to path after exploring all outgoing edges
- Guarantees we find valid Eulerian path
- Handles cycles and dead ends correctly

## Complexity Analysis
- **Time**: O(E log E) for sorting destinations, O(E) for DFS
- **Space**: O(V + E) for graph representation

## Example Walkthrough
**Input**: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]

```
Graph:
JFK -> [ATL, SFO]  (sorted: ATL < SFO)
ATL -> [JFK, SFO]
SFO -> [ATL]

DFS traversal:
1. Start JFK, go to ATL (smallest)
2. From ATL, go to JFK (only option)
3. From JFK, go to SFO (only remaining)
4. From SFO, go to ATL
5. From ATL, go to SFO

Post-order: [SFO, ATL, SFO, JFK, ATL, JFK]
Reverse: [JFK, ATL, JFK, SFO, ATL, SFO]
```

## Edge Cases
1. **Single ticket**: Simple JFK → DEST
2. **Cycle**: JFK ↔ ATL (both directions)
3. **Dead end**: Need to ensure all tickets used
4. **Multiple starts from same airport**: Choose lexical smallest

## Implementation Details
```go
func Solve0332(tickets [][]string) []string {
    // Build graph with sorted destinations
    graph := make(map[string][]string)
    for _, ticket := range tickets {
        from, to := ticket[0], ticket[1]
        graph[from] = append(graph[from], to)
    }
    
    // Sort for lexical order
    for airport := range graph {
        sort.Strings(graph[airport])
    }
    
    // Hierholzer's algorithm
    itinerary := []string{}
    var dfs func(airport string)
    dfs = func(airport string) {
        for len(graph[airport]) > 0 {
            next := graph[airport][0]
            graph[airport] = graph[airport][1:]  // Remove edge
            dfs(next)
        }
        itinerary = append(itinerary, airport)  // Post-order
    }
    
    dfs("JFK")
    
    // Reverse for correct order
    for i, j := 0, len(itinerary)-1; i < j; i, j = i+1, j-1 {
        itinerary[i], itinerary[j] = itinerary[j], itinerary[i]
    }
    
    return itinerary
}
```

## Related Problems
- **2097 Valid Arrangement of Pairs**: General Eulerian path
- **332 Reconstruct Itinerary**: Same problem
- **753 Cracking the Safe**: De Bruijn sequence (related concept)

## Takeaways
1. Eulerian path problems can be solved with Hierholzer's algorithm
2. Post-order DFS ensures valid path construction
3. Lexical order requires sorting destinations
4. Real-world application: flight itinerary reconstruction