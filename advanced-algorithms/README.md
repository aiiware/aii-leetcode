# Advanced Algorithms

This directory contains implementations of advanced algorithms and data structures that go beyond standard LeetCode problems.

## Algorithms Included

### 1. Union-Find (Disjoint Set Union)
- **Purpose**: Efficiently manage partitions of a set
- **Applications**: Connected components, Kruskal's algorithm, network connectivity
- **Complexity**: Near-constant time operations with path compression and union by rank

### 2. Segment Tree
- **Purpose**: Range queries and updates in O(log n) time
- **Applications**: Range sum/min/max queries, range updates
- **Variants**: Lazy propagation, Fenwick Tree (Binary Indexed Tree)

### 3. Fenwick Tree (Binary Indexed Tree)
- **Purpose**: Efficient prefix sum queries and point updates
- **Applications**: Frequency counting, inversion counting, cumulative frequency

### 4. Trie (Prefix Tree)
- **Purpose**: Efficient string prefix operations
- **Applications**: Autocomplete, spell checking, IP routing

### 5. Suffix Array & LCP Array
- **Purpose**: Efficient substring search and pattern matching
- **Applications**: String matching, longest common substring, pattern discovery

### 6. KMP Algorithm (Knuth-Morris-Pratt)
- **Purpose**: Linear-time string pattern matching
- **Applications**: Text search, DNA sequence matching

### 7. Aho-Corasick Algorithm
- **Purpose**: Multiple pattern string matching
- **Applications**: Virus scanning, text mining, keyword search

### 8. Manacher's Algorithm
- **Purpose**: Find all palindromic substrings in linear time
- **Applications**: Palindrome detection, string analysis

### 9. Z-Algorithm
- **Purpose**: Pattern matching and string preprocessing
- **Applications**: String similarity, pattern search

### 10. Mo's Algorithm
- **Purpose**: Answer offline range queries efficiently
- **Applications**: Range queries with updates, statistics on subarrays

## Usage

Each algorithm is implemented as a standalone Go package with:
- Clean, well-documented code
- Comprehensive test cases
- Example usage
- Time and space complexity analysis

## Adding New Algorithms

When adding a new advanced algorithm:
1. Create a new Go file with the algorithm implementation
2. Include comprehensive test cases
3. Add documentation explaining the algorithm
4. Update this README with the new algorithm
5. Add to the index in the main README.md

## Testing

Run tests for all advanced algorithms:
```bash
go test ./advanced-algorithms/...
```