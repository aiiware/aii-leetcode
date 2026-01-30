# 📊 Dynamic Programming Solutions Index

**Last Updated**: January 28, 2026  
**Total DP Problems**: 24  
**Total Files**: 315+ (LeetCode problems in repository)

## 📈 Overview

This index catalogs all Dynamic Programming solutions in the LeetCode codebase, organized by pattern, difficulty, and complexity.

## 📋 Quick Statistics

| Category | Count | Percentage |
|----------|-------|------------|
| **Total DP Problems** | 24 | 100% |
| **Easy** | 0 | 0% |
| **Medium** | 16 | 66.7% |
| **Hard** | 8 | 33.3% |
| **1D DP** | 10 | 41.7% |
| **2D DP** | 14 | 58.3% |
| **Memoization** | 19 | 79.2% |
| **Space-Optimized** | 18 | 75.0% |

## 🎯 DP Categories & Patterns

### 1. **1D DP (Linear State)**
Problems where state depends only on previous states in one dimension.

| Problem | Difficulty | Pattern | Time | Space | Key Insight |
|---------|------------|---------|------|-------|-------------|
| [0022. Generate Parentheses](0022_generate_parentheses.go) | Medium | Catalan Numbers | O(4ⁿ/√n) | O(4ⁿ/√n) | dp[i] = all combinations for i pairs |
| [0032. Longest Valid Parentheses](0032_longest_valid_parentheses.go) | Hard | Parentheses Validation | O(n) | O(n) | dp[i] = length ending at i |
| [0039. Combination Sum](0039_combination_sum.go) | Medium | Unbounded Knapsack | O(n×target) | O(target) | dp[i] = combinations for sum i |
| [0040. Combination Sum II](0040_combination_sum_ii.go) | Medium | Bounded Knapsack | O(n×target) | O(target) | Handle duplicates, each num once |
| [0045. Jump Game II](0045_jump_game_ii.go) | Medium | Minimum Jumps | O(n²) | O(n) | dp[i] = min jumps to reach i |
| [0091. Decode Ways](0091_decode_ways.go) | Medium | String Decoding | O(n) | O(n) | dp[i] = ways to decode s[0:i] |
| [0096. Unique Binary Search Trees](0096_unique_binary_search_trees.go) | Medium | Catalan Numbers | O(n²) | O(n) | dp[i] = ∑ dp[j-1]×dp[i-j] |
| [0120. Triangle](0120_triangle.go) | Medium | Path Sum | O(n²) | O(n) | dp[j] = min(left, right) + triangle[i][j] |
| [0132. Palindrome Partitioning II](0132_palindrome_partitioning_ii.go) | Hard | Palindrome + Min Cuts | O(n²) | O(n²) | dp[i] = min cuts for s[0:i] |
| [0139. Word Break](0139_word_break.go) | Medium | String Segmentation | O(n×m) | O(n) | dp[i] = s[0:i] can be segmented |

### 2. **2D DP (Matrix/String Comparison)**
Problems with two-dimensional state (grid, string comparison).

| Problem | Difficulty | Pattern | Time | Space | Key Insight |
|---------|------------|---------|------|-------|-------------|
| [0005. Longest Palindromic Substring](0005_longest_palindromic_substring.go) | Medium | Palindrome Check | O(n²) | O(n²) | dp[i][j] = s[i:j+1] is palindrome |
| [0010. Regular Expression Matching](0010_regular_expression_matching.go) | Hard | Pattern Matching | O(m×n) | O(m×n) | Handle '.', '*', character matching |
| [0044. Wildcard Matching](0044_wildcard_matching.go) | Hard | Wildcard Pattern | O(m×n) | O(m×n) | Handle '?', '*' matching |
| [0062. Unique Paths](0062_unique_paths.go) | Medium | Grid Paths | O(m×n) | O(n) | dp[j] = dp[j] + dp[j-1] |
| [0063. Unique Paths II](0063_unique_paths_ii.go) | Medium | Grid with Obstacles | O(m×n) | O(n) | Handle obstacles in grid |
| [0064. Minimum Path Sum](0064_minimum_path_sum.go) | Medium | Grid Min Sum | O(m×n) | O(n) | dp[j] = min(above, left) + grid[i][j] |
| [0072. Edit Distance](0072_edit_distance.go) | Hard | String Edit | O(m×n) | O(n) | dp[j] = min(insert, delete, replace) |
| [0077. Combinations](0077_combinations.go) | Medium | Combination Generation | O(C(n,k)) | O(C(n,k)) | dp[i][j] = combinations of size j from 1..i |
| [0087. Scramble String](0087_scramble_string.go) | Hard | String Scramble | O(n⁴) | O(n³) | dp[i][j][k] = s1[i:i+k] scramble of s2[j:j+k] |
| [0095. Unique Binary Search Trees II](0095_unique_binary_search_trees_ii.go) | Medium | Tree Generation | O(Cₙ) | O(Cₙ) | dp[i] = all BSTs with i nodes |
| [0097. Interleaving String](0097_interleaving_string.go) | Medium | String Interleaving | O(m×n) | O(n) | dp[i][j] = s3[0:i+j] interleaving of s1[0:i], s2[0:j] |
| [0115. Distinct Subsequences](0115_distinct_subsequences.go) | Hard | Subsequence Count | O(m×n) | O(n) | dp[i][j] = subsequences of s[0:i] equal to t[0:j] |
| [0161. One Edit Distance](0161_one_edit_distance.go) | Medium | Edit Distance Variant | O(m×n) | O(n) | Check if exactly one edit away |
| [0174. Dungeon Game](0174_dungeon_game.go) | Hard | Reverse DP | O(m×n) | O(m×n) | dp[i][j] = min health needed at (i,j) |

## 🔍 DP Pattern Analysis

### Pattern 1: Fibonacci-like Recurrence
```go
// Examples: Climbing Stairs, Decode Ways
dp[i] = dp[i-1] + dp[i-2]  // or similar variations
```
**Problems**: 0091, 0096

### Pattern 2: String Comparison (Edit Distance Style)
```go
// Examples: Edit Distance, Interleaving String
if s1[i-1] == s2[j-1] {
    dp[i][j] = dp[i-1][j-1]
} else {
    dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
}
```
**Problems**: 0072, 0097, 0115, 0161

### Pattern 3: Grid Path Counting
```go
// Examples: Unique Paths, Minimum Path Sum
dp[i][j] = dp[i-1][j] + dp[i][j-1]  // or min(above, left) + grid[i][j]
```
**Problems**: 0062, 0063, 0064, 0174

### Pattern 4: Palindrome Checking
```go
// Examples: Longest Palindromic Substring
if s[i] == s[j] && dp[i+1][j-1] {
    dp[i][j] = true
}
```
**Problems**: 0005, 0132

### Pattern 5: Knapsack/Combination
```go
// Examples: Combination Sum, Word Break
for num in nums {
    for i := num; i <= target; i++ {
        dp[i] += dp[i-num]
    }
}
```
**Problems**: 0039, 0040, 0139

### Pattern 6: Pattern Matching
```go
// Examples: Regular Expression Matching, Wildcard Matching
if p[j-1] == '*' {
    dp[i][j] = dp[i][j-2] || (match && dp[i-1][j])
}
```
**Problems**: 0010, 0044

### Pattern 7: Parentheses Generation/Validation
```go
// Examples: Generate Parentheses, Longest Valid Parentheses
dp[i] = "(" + dp[j] + ")" + dp[i-1-j]  // or dp[i] = dp[i-2] + 2
```
**Problems**: 0022, 0032

## 🚀 Space Optimization Techniques

### 1. **1D Array Optimization** (from 2D)
```go
// Before: O(m×n) space
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
}

// After: O(n) space
dp := make([]int, n+1)
prev := make([]int, n+1)
```
**Used in**: 0072, 0097, 0115, 0161

### 2. **Two-Variable Approach** (Fibonacci-like)
```go
// Before: O(n) space
dp := make([]int, n+1)
dp[0], dp[1] = 0, 1
for i := 2; i <= n; i++ {
    dp[i] = dp[i-1] + dp[i-2]
}

// After: O(1) space
prev2, prev1 := 0, 1
for i := 2; i <= n; i++ {
    current := prev1 + prev2
    prev2, prev1 = prev1, current
}
```
**Used in**: 0091, 0096

### 3. **Rolling Array** (Grid Problems)
```go
// Update in-place
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
    }
}
```
**Used in**: 0062, 0063, 0064

### 4. **Memoization with Maps** (Top-Down)
```go
func solve(n int, memo map[int]int) int {
    if val, exists := memo[n]; exists {
        return val
    }
    // Compute and store
    memo[n] = solve(n-1, memo) + solve(n-2, memo)
    return memo[n]
}
```
**Used in**: 19 files with memoization patterns

## 📊 Complexity Analysis

### Time Complexity Distribution
- **O(n)**: 4 problems (16.7%)
- **O(n²)**: 8 problems (33.3%)
- **O(m×n)**: 8 problems (33.3%)
- **O(n³) or higher**: 4 problems (16.7%)

### Space Complexity Distribution
- **O(1)**: 0 problems (0%)
- **O(n)**: 10 problems (41.7%)
- **O(n²)**: 8 problems (33.3%)
- **O(m×n)**: 6 problems (25%)

## 🎯 Difficulty Progression

### Beginner (1D DP)
1. **0091. Decode Ways** - Simple 1D recurrence
2. **0096. Unique Binary Search Trees** - Catalan numbers
3. **0120. Triangle** - Simple path optimization

### Intermediate (2D DP)
1. **0062. Unique Paths** - Basic grid DP
2. **0064. Minimum Path Sum** - Grid with weights
3. **0139. Word Break** - String segmentation
4. **0039. Combination Sum** - Unbounded knapsack

### Advanced (Complex Patterns)
1. **0005. Longest Palindromic Substring** - Palindrome DP
2. **0072. Edit Distance** - Classic string edit
3. **0097. Interleaving String** - String interleaving
4. **0115. Distinct Subsequences** - Subsequence counting

### Expert (Hard Problems)
1. **0010. Regular Expression Matching** - Complex pattern matching
2. **0044. Wildcard Matching** - Wildcard pattern matching
3. **0087. Scramble String** - 3D DP state
4. **0174. Dungeon Game** - Reverse DP

## 🔧 Implementation Patterns

### Common DP Table Initialization
```go
// 1D DP
dp := make([]int, n+1)
dp[0] = baseCase

// 2D DP
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
}
dp[0][0] = baseCase
```

### State Transition Patterns
```go
// 1. Additive (Fibonacci, Paths)
dp[i] = dp[i-1] + dp[i-2]

// 2. Minimum/Maximum (Edit Distance, Path Sum)
dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + cost

// 3. Conditional (String Matching)
if condition {
    dp[i][j] = dp[i-1][j-1]
} else {
    dp[i][j] = dp[i-1][j] + dp[i][j-1]
}

// 4. Combination (Knapsack)
for each item {
    for each capacity {
        dp[cap] = max(dp[cap], dp[cap-weight] + value)
    }
}
```

## 📚 Learning Path Recommendations

### Phase 1: Foundation (Week 1-2)
1. Master 1D DP patterns (0091, 0096, 0120)
2. Understand space optimization techniques
3. Practice Fibonacci-like recurrences

### Phase 2: Grid & Strings (Week 3-4)
1. Learn 2D DP for grids (0062, 0063, 0064)
2. Practice string comparison DP (0072, 0097)
3. Understand edit distance variations

### Phase 3: Advanced Patterns (Week 5-6)
1. Master pattern matching (0010, 0044)
2. Learn knapsack variations (0039, 0040, 0139)
3. Practice palindrome DP (0005, 0132)

### Phase 4: Expert Level (Week 7-8)
1. Solve hard DP problems (0087, 0115, 0174)
2. Implement 3D DP states
3. Master reverse DP techniques

## 🎮 Interactive Exercises

### Exercise 1: Identify DP Pattern
Given a problem, identify which DP pattern applies:
1. "Count ways to decode a string" → **Fibonacci-like**
2. "Find minimum path sum in grid" → **Grid DP**
3. "Check if strings are interleaving" → **String Comparison**
4. "Generate all valid parentheses" → **Catalan Numbers**

### Exercise 2: Space Optimization
Take a 2D DP solution and optimize to 1D:
1. Edit Distance (0072) → Already optimized
2. Interleaving String (0097) → Already optimized
3. Unique Paths (0062) → Already optimized

### Exercise 3: Pattern Recognition
Match problems to their recurrence relations:
1. 0091. Decode Ways → `dp[i] = dp[i-1] + dp[i-2]` (with conditions)
2. 0072. Edit Distance → `dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1`
3. 0062. Unique Paths → `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

## 📈 Missing DP Categories

Based on standard DP curriculum, consider adding:

### 1. **Bitmask DP**
- Traveling Salesman Problem
- Subset DP problems
- State compression techniques

### 2. **Digit DP**
- Counting problems with digit constraints
- Number of integers in range with property
- Digit-by-digit DP

### 3. **Tree DP**
- Maximum path sum in binary tree
- House Robber III (tree version)
- Tree diameter with DP

### 4. **Probability DP**
- Expected value problems
- Markov chain DP
- Game theory with probabilities

### 5. **Flow DP**
- Maximum flow with DP
- Network flow problems
- Matching problems

## ✅ Best Practices Checklist

### For Each DP Solution:
- [ ] **Clear state definition**: What does dp[i] or dp[i][j] represent?
- [ ] **Proper base cases**: Handle smallest subproblems correctly
- [ ] **Correct recurrence**: State transition matches problem logic
- [ ] **Space optimization**: Consider if O(1) or O(n) space is possible
- [ ] **Edge cases**: Handle empty inputs, boundaries, constraints
- [ ] **Time complexity**: Analyze and document complexity
- [ ] **Test coverage**: Include LeetCode examples and edge cases
- [ ] **Alternative approaches**: Mention memoization vs tabulation

### Code Quality:
- [ ] **Meaningful variable names**: dp, memo, prev, curr
- [ ] **Comments for complex logic**: Explain state transitions
- [ ] **Error handling**: Check inputs, handle invalid cases
- [ ] **Performance considerations**: Mention optimization trade-offs

## 🔍 Quick Reference Table

| Problem | Pattern | Difficulty | Time | Space | Key File |
|---------|---------|------------|------|-------|----------|
| 0005 | Palindrome 2D | Medium | O(n²) | O(n²) | [0005_longest_palindromic_substring.go](0005_longest_palindromic_substring.go) |
| 0010 | Regex Matching | Hard | O(m×n) | O(m×n) | [0010_regular_expression_matching.go](0010_regular_expression_matching.go) |
| 0022 | Catalan Numbers | Medium | O(4ⁿ/√n) | O(4ⁿ/√n) | [0022_generate_parentheses.go](0022_generate_parentheses.go) |
| 0032 | Parentheses DP | Hard | O(n) | O(n) | [0032_longest_valid_parentheses.go](0032_longest_valid_parentheses.go) |
| 0039 | Unbounded Knapsack | Medium | O(n×target) | O(target) | [0039_combination_sum.go](0039_combination_sum.go) |
| 0040 | Bounded Knapsack | Medium | O(n×target) | O(target) | [0040_combination_sum_ii.go](0040_combination_sum_ii.go) |
| 0044 | Wildcard Matching | Hard | O(m×n) | O(m×n) | [0044_wildcard_matching.go](0044_wildcard_matching.go) |
| 0045 | Min Jumps | Medium | O(n²) | O(n) | [0045_jump_game_ii.go](0045_jump_game_ii.go) |
| 0062 | Grid Paths | Medium | O(m×n) | O(n) | [0062_unique_paths.go](0062_unique_paths.go) |
| 0063 | Grid with Obstacles | Medium | O(m×n) | O(n) | [0063_unique_paths_ii.go](0063_unique_paths_ii.go) |
| 0064 | Min Path Sum | Medium | O(m×n) | O(n) | [0064_minimum_path_sum.go](0064_minimum_path_sum.go) |
| 0072 | Edit Distance | Hard | O(m×n) | O(n) | [0072_edit_distance.go](0072_edit_distance.go) |
| 0077 | Combinations | Medium | O(C(n,k)) | O(C(n,k)) | [0077_combinations.go](0077_combinations.go) |
| 0087 | Scramble String | Hard | O(n⁴) | O(n³) | [0087_scramble_string.go](0087_scramble_string.go) |
| 0091 | Decode Ways | Medium | O(n) | O(n) | [0091_decode_ways.go](0091_decode_ways.go) |
| 0095 | BST Generation | Medium | O(Cₙ) | O(Cₙ) | [0095_unique_binary_search_trees_ii.go](0095_unique_binary_search_trees_ii.go) |
| 0096 | Catalan Numbers | Medium | O(n²) | O(n) | [0096_unique_binary_search_trees.go](0096_unique_binary_search_trees.go) |
| 0097 | Interleaving String | Medium | O(m×n) | O(n) | [0097_interleaving_string.go](0097_interleaving_string.go) |
| 0115 | Distinct Subseq | Hard | O(m×n) | O(n) | [0115_distinct_subsequences.go](0115_distinct_subsequences.go) |
| 0120 | Triangle Path | Medium | O(n²) | O(n) | [0120_triangle.go](0120_triangle.go) |
| 0132 | Palindrome Cuts | Hard | O(n²) | O(n²) | [0132_palindrome_partitioning_ii.go](0132_palindrome_partitioning_ii.go) |
| 0139 | Word Break | Medium | O(n×m) | O(n) | [0139_word_break.go](0139_word_break.go) |
| 0161 | One Edit Distance | Medium | O(m×n) | O(n) | [0161_one_edit_distance.go](0161_one_edit_distance.go) |
| 0174 | Reverse DP | Hard | O(m×n) | O(m×n) | [0174_dungeon_game.go](0174_dungeon_game.go) |

## 🎯 Next Steps for DP Mastery

1. **Complete all classic DP problems** in LeetCode DP card
2. **Implement both approaches** (memoization & tabulation) for each problem
3. **Practice pattern recognition** by categorizing new problems
4. **Master space optimization** for each DP pattern
5. **Move to advanced topics**: Bitmask DP, Tree DP, Probability DP
6. **Participate in contests** to apply DP under time pressure
7. **Review and refactor** existing solutions for optimization

## 📝 Maintenance Notes

This index should be updated when:
- New DP problems are added to the repository
- Existing solutions are optimized or refactored
- New DP patterns are discovered or implemented
- Space/time complexity improvements are made

**Last Analysis**: January 28, 2026  
**Total Problems Analyzed**: 24  
**Files Examined**: 26 (including test files)

---

*Generated by Aii Agent - Comprehensive DP Solutions Analysis*  
*Use this index as a reference for studying DP patterns and optimizing solutions.*