# 0310 Minimum Height Trees

## Problem Statement
Given a tree (undirected acyclic connected graph) with `n` nodes labeled from `0` to `n-1`, and an array of `n-1` edges, find all **Minimum Height Trees (MHTs)**. An MHT is a rooted tree where the height (longest downward path from root to leaf) is minimized.

## Key Insights
1. **Center of Tree**: MHTs are essentially the centers of the tree
2. **Leaf Trimming**: We can find centers by repeatedly trimming leaves (nodes with degree 1)
3. **At Most 2 Centers**: A tree can have at most 2 centers

## Algorithm: Leaf Trimming (Topological Sort)
The optimal approach uses a BFS-like leaf trimming strategy:

### Step-by-Step
1. **Build adjacency list** and calculate degrees for each node
2. **Initialize queue** with all leaves (nodes with degree 1)
3. **Trim leaves layer by layer**:
   - Process all leaves in current layer
   - For each leaf, reduce degree of its neighbors
   - If neighbor becomes a leaf, add to next layer
4. **Stop when 1 or 2 nodes remain** - these are the centers/MHT roots

### Why This Works
- Leaves contribute to height, so removing them reduces height
- Centers are last nodes remaining after trimming all leaves
- Similar to finding diameter endpoints and their midpoint

## Complexity Analysis
- **Time**: O(n) - each node and edge processed once
- **Space**: O(n) - for adjacency list and queue

## Example Walkthrough
**Input**: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]

```
Initial tree:
    0
    |
    3
   /|\
  1 2 4
        \
         5

Step 1: Trim leaves 0, 1, 2, 5
Remaining: 3-4

Step 2: Trim leaves 3, 4 (both become leaves)
Remaining: 3, 4

Result: [3, 4]
```

## Edge Cases
1. **n = 1**: Single node is its own MHT
2. **n = 2**: Both nodes are MHTs
3. **Line/Path**: Middle node(s) are MHTs
4. **Star**: Center node is the only MHT

## Implementation Details
```go
func Solve0310(n int, edges [][]int) []int {
    if n == 1 { return []int{0} }
    
    // Build graph and degrees
    adj := make([][]int, n)
    degree := make([]int, n)
    
    // ... (implementation)
    
    // Trim leaves until 1-2 nodes remain
    leaves := []int{}
    for i := 0; i < n; i++ {
        if degree[i] == 1 {
            leaves = append(leaves, i)
        }
    }
    
    remaining := n
    for remaining > 2 {
        // Process current layer
        // ... (trim leaves, update degrees)
    }
    
    return leaves
}
```

## Related Problems
- **0543 Diameter of Binary Tree**: Finding longest path
- **1245 Tree Diameter**: Similar center-finding concept
- **3102 Minimum Height Trees**: Same problem

## Takeaways
1. Tree centers minimize height
2. Leaf trimming is efficient O(n) approach
3. Useful for network design where minimizing communication delay is important