# 0269 - Alien Dictionary

## Problem Statement
Given a sorted dictionary of an alien language having N words and k starting alphabets of standard dictionary, find the order of characters in the alien language.

**Example:**
```
Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"
```

## Solution Approach

### Key Insight
This is a **topological sorting** problem. The words are sorted according to the alien language's alphabetical order. By comparing adjacent words, we can extract ordering relationships between characters.

### Algorithm Steps

1. **Build Graph:**
   - Create adjacency list and indegree map for all unique characters
   - Compare each pair of adjacent words to find the first differing character
   - Add a directed edge from `word1[j]` to `word2[j]` where `j` is the first position where they differ

2. **Detect Invalid Cases:**
   - If `word2` is a prefix of `word1` and `word1` is longer, the ordering is invalid
   - Example: `["abc", "ab"]` is invalid because "ab" should come before "abc"

3. **Topological Sort (Kahn's Algorithm):**
   - Initialize queue with all nodes having indegree 0
   - Process nodes in queue, adding their neighbors and decrementing indegrees
   - If we can process all nodes, we have a valid topological order
   - If not, there's a cycle and the ordering is impossible

### Complexity Analysis
- **Time Complexity:** O(C) where C is the total number of characters across all words
- **Space Complexity:** O(1) or O(U) where U is the number of unique characters (max 26 for lowercase English)

### Code Implementation
```go
func alienOrder(words []string) string {
    // Build graph from word comparisons
    // Perform topological sort
    // Return result or empty string if cycle detected
}
```

## Example Walkthrough

For `["wrt","wrf","er","ett","rftt"]`:

1. Compare "wrt" and "wrf": t → f
2. Compare "wrf" and "er": w → e  
3. Compare "er" and "ett": r → t
4. Compare "ett" and "rftt": e → r

Graph edges: t→f, w→e, r→t, e→r

Topological order: w → e → r → t → f → "wertf"

## Edge Cases

1. **Single word:** Return all unique characters in any order
2. **Empty input:** Return empty string
3. **Cycle detection:** Return empty string if cycle exists
4. **Invalid prefix:** Return empty string if word2 is prefix of word1

## Related Problems
- 0207 - Course Schedule (similar topological sort)
- 0210 - Course Schedule II
- 0444 - Sequence Reconstruction

## Learning Points
- Topological sort is useful for dependency resolution
- Graph construction from implicit relationships
- Cycle detection in directed graphs