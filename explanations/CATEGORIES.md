# LeetCode Solution Explanations - Category Index

This file serves as a master index for all solution explanations organized by category.

## 📊 Progress Tracking

**Total Target: 20 explanations** (7 DP, 7 Graphs, 6 Design)
**Current Status: 5/20 completed** (25%)

### ✅ Completed Explanations

#### Arrays (1)
- ✅ **0004 - Median of Two Sorted Arrays** - `arrays/0004_median_of_two_sorted_arrays.md`
  - *Technical style: Binary search on two sorted arrays*
  - Difficulty: Hard
  - Tags: binary-search, arrays, median, partition

#### Dynamic Programming (2/7)
- ✅ **0010 - Regular Expression Matching** - `dp/0010_regular_expression_matching.md`
  - *Educational style: DP with string matching patterns*
  - Difficulty: Hard
  - Tags: dynamic-programming, strings, regex, memoization

- ✅ **0070 - Climbing Stairs** - `dp/0070_climbing_stairs.md`
  - *Educational style: Step-by-step Fibonacci pattern explanation*
  - Difficulty: Easy
  - Tags: fibonacci, memoization, dynamic-programming

#### Graphs (1/7)
- ✅ **0200 - Number of Islands** - `graphs/0200_number_of_islands.md`
  - *Practical style: Real-world applications and implementation patterns*
  - Difficulty: Medium
  - Tags: dfs, bfs, matrix, grid, connected-components

#### Design (1/6)
- ✅ **0146 - LRU Cache** - `design/0146_lru_cache.md`
  - *Technical style: Data structure design with O(1) operations*
  - Difficulty: Medium
  - Tags: design, cache, linked-list, hash-map

### 📋 Planned Explanations

#### Dynamic Programming (5 remaining)
- [ ] **0198 - House Robber** - *1D DP*
- [ ] **0322 - Coin Change** - *Unbounded knapsack*
- [ ] **0416 - Partition Equal Subset Sum** - *Subset sum*
- [ ] **1143 - Longest Common Subsequence** - *2D DP*
- [ ] **0072 - Edit Distance** - *String DP*
- [ ] **0312 - Burst Balloons** - *Interval DP*

#### Graphs (6 remaining)
- [ ] **0207 - Course Schedule** - *Topological sort*
- [ ] **0210 - Course Schedule II** - *Topological sort with ordering*
- [ ] **0399 - Evaluate Division** - *Graph + Union Find*
- [ ] **0785 - Is Graph Bipartite?** - *Graph coloring*
- [ ] **0743 - Network Delay Time** - *Dijkstra's*
- [ ] **0323 - Number of Connected Components** - *Union Find*

#### Design (5 remaining)
- [ ] **0155 - Min Stack** - *Stack design*
- [ ] **0297 - Serialize and Deserialize Binary Tree** - *Tree serialization*
- [ ] **0380 - Insert Delete GetRandom O(1)** - *Randomized data structure*
- [ ] **0295 - Find Median from Data Stream** - *Heap design*
- [ ] **0355 - Design Twitter** - *System design*

### 📁 Directory Structure

```
explanations/
├── dp/                          # Dynamic Programming
│   ├── 0010_regular_expression_matching.md ✅
│   ├── 0070_climbing_stairs.md ✅
│   └── ... (5 more planned)
├── graphs/                      # Graph Algorithms
│   ├── 0200_number_of_islands.md ✅
│   └── ... (6 more planned)
├── design/                      # Design Problems
│   ├── 0146_lru_cache.md ✅
│   └── ... (5 more planned)
├── arrays/                      # Array Problems
│   ├── 0004_median_of_two_sorted_arrays.md ✅
│   └── (additional array problems)
├── linked-lists/                # Linked List Problems
│   └── (empty - placeholder)
├── math/                        # Math Problems
│   └── (empty - placeholder)
├── sorting/                     # Sorting Problems
│   └── (empty - placeholder)
├── data_structures/             # Data Structures
│   └── (empty - placeholder)
├── indexes/                     # Index Problems
│   └── (empty - placeholder)
├── trees/                       # Tree Problems
│   └── (empty - placeholder)
├── TEMPLATE.md                  # Explanation template ✅
├── CATEGORIES.md                # This file (master index) ✅
└── README.md                    # Directory overview
```

### 🎯 Explanation Styles

Each explanation follows one of three styles based on problem type:

1. **Educational** (DP problems): Step-by-step reasoning, visual diagrams, learning objectives
2. **Practical** (Graph problems): Real-world applications, implementation patterns, common pitfalls
3. **Technical** (Design problems): Algorithm analysis, complexity proofs, data structure design

### 🔄 Workflow

1. **Select problem** from planned list above
2. **Choose style** based on problem category
3. **Create file** using `TEMPLATE.md` as guide
4. **Update this index** to mark as completed
5. **Link to solution** in corresponding Go file

### 📈 Statistics

- **Total problems in repository**: 150+ (estimated)
- **Target explanations**: 20 (13% coverage)
- **Current completion**: 5 (25% of target)
- **Categories covered**: 4/10 (Arrays, DP, Graphs, Design)
- **Styles demonstrated**: 3/3 (Educational, Practical, Technical)

---

*Last updated: 2026-01-31*
*Next priority: Complete more DP and Graph explanations to reach 10 total*