# Dynamic Programming Category Overview

## Introduction to Dynamic Programming

Dynamic Programming (DP) is a method for solving complex problems by breaking them down into simpler subproblems. It's particularly useful for optimization problems where the solution can be constructed from solutions to overlapping subproblems.

### Key Characteristics of DP Problems
1. **Optimal Substructure**: Optimal solution can be constructed from optimal solutions of subproblems
2. **Overlapping Subproblems**: The same subproblems are solved multiple times
3. **Memoization/Tabulation**: Store solutions to subproblems to avoid recomputation

## DP Problem Patterns

### 1. 1D DP (Linear)
**Characteristics**: Single dimension state, linear recurrence
**Examples**:
- **0070 - Climbing Stairs**: `dp[i] = dp[i-1] + dp[i-2]`
- **0198 - House Robber**: `dp[i] = max(dp[i-1], dp[i-2] + nums[i])`
- **0322 - Coin Change**: `dp[amount] = min(dp[amount], 1 + dp[amount - coin])`

### 2. 2D DP (Matrix)
**Characteristics**: Two-dimensional state, often for sequences or grids
**Examples**:
- **1143 - Longest Common Subsequence**: `dp[i][j]` for strings of length i and j
- **0072 - Edit Distance**: `dp[i][j]` for transforming one string to another
- **0064 - Minimum Path Sum**: `dp[i][j]` for grid traversal

### 3. Interval DP
**Characteristics**: Solve problems on intervals or subsequences
**Examples**:
- **0312 - Burst Balloons**: `dp[i][j]` for maximum coins from balloons i to j
- **0516 - Longest Palindromic Subsequence**: `dp[i][j]` for substring i to j

### 4. Knapsack DP
**Characteristics**: Resource allocation with constraints
**Examples**:
- **0416 - Partition Equal Subset Sum**: 0/1 knapsack variation
- **0494 - Target Sum**: Subset sum with +/- operations

## DP Solution Strategies

### Top-Down (Memoization)
```go
func solve(n int, memo map[int]int) int {
    if n <= baseCase { return baseValue }
    if val, exists := memo[n]; exists { return val }
    
    result = combine(solve(sub1, memo), solve(sub2, memo))
    memo[n] = result
    return result
}
```

### Bottom-Up (Tabulation)
```go
func solve(n int) int {
    dp := make([]int, n+1)
    dp[0] = baseValue
    
    for i := 1; i <= n; i++ {
        dp[i] = combine(dp[sub1], dp[sub2])
    }
    
    return dp[n]
}
```

## Complexity Analysis

### Time Complexity
- **1D DP**: Typically O(n) or O(n × k) where k is a constant
- **2D DP**: Typically O(m × n) for m × n states
- **State Transition**: Multiply by cost of computing each state

### Space Complexity
- **Naive**: O(n) or O(m × n) for storing all states
- **Optimized**: Often reducible to O(n) or O(1) with careful implementation
- **Rolling Arrays**: Use only needed previous states

## Common DP Techniques

### 1. State Compression
Reduce multi-dimensional DP to fewer dimensions when possible.

### 2. Prefix/Suffix DP
Compute prefix or suffix results to optimize transitions.

### 3. Bitmask DP
Use bitmasks to represent subsets (for problems up to n ≤ 20).

### 4. Digit DP
Solve problems involving digits of numbers.

## DP vs Other Approaches

### When to Use DP
- Problem has optimal substructure
- Subproblems overlap significantly
- Brute force would be exponential
- Greedy doesn't work (no optimal local choice)

### When NOT to Use DP
- Subproblems don't overlap (use divide and conquer)
- Greedy approach works (simpler and faster)
- Problem size too large for DP table
- State space is too complex to define

## Learning Path

### Beginner Level
1. **0070 - Climbing Stairs**: Introduction to recurrence relations
2. **0198 - House Robber**: Simple 1D DP with constraints
3. **0322 - Coin Change**: Unbounded knapsack pattern

### Intermediate Level
1. **0416 - Partition Equal Subset Sum**: 0/1 knapsack
2. **1143 - Longest Common Subsequence**: Classic 2D DP
3. **0072 - Edit Distance**: String transformation DP

### Advanced Level
1. **0312 - Burst Balloons**: Interval DP
2. **0516 - Longest Palindromic Subsequence**: Palindrome DP
3. **0329 - Longest Increasing Path in Matrix**: DP on DAG

## Practice Problems by Difficulty

### Easy
- 0070 - Climbing Stairs
- 0198 - House Robber
- 0303 - Range Sum Query (Prefix Sum)

### Medium
- 0322 - Coin Change
- 0416 - Partition Equal Subset Sum
- 1143 - Longest Common Subsequence
- 0072 - Edit Distance
- 0300 - Longest Increasing Subsequence

### Hard
- 0312 - Burst Balloons
- 0516 - Longest Palindromic Subsequence
- 0329 - Longest Increasing Path in Matrix
- 0877 - Stone Game

## Optimization Tips

### Space Optimization
```go
// Before: O(n) space
dp := make([]int, n+1)

// After: O(1) space for Fibonacci-like
prev2, prev1 := 1, 1
for i := 2; i <= n; i++ {
    current := prev1 + prev2
    prev2, prev1 = prev1, current
}
```

### Time Optimization
- Use prefix sums to compute range queries in O(1)
- Precompute frequently used values
- Use binary search in DP transitions when applicable

## Common Mistakes

1. **Incorrect State Definition**: DP state should capture all needed information
2. **Missing Base Cases**: Forgot to initialize dp[0] or handle edge cases
3. **Wrong Transition Order**: Computing dp[i] before needed dp[j]
4. **Integer Overflow**: Large numbers in DP tables
5. **Memory Limit Exceeded**: Not optimizing space when possible

## Real-World Applications

1. **Bioinformatics**: Sequence alignment (Edit Distance, LCS)
2. **Finance**: Portfolio optimization (Knapsack variations)
3. **Natural Language Processing**: Word segmentation, parsing
4. **Computer Graphics**: Image processing, seam carving
5. **Operations Research**: Resource allocation, scheduling

## Additional Resources

### Books
- "Introduction to Algorithms" (CLRS) - DP chapter
- "Algorithm Design Manual" - DP techniques
- "Competitive Programming" - DP patterns

### Online Courses
- MIT OpenCourseWare: Dynamic Programming
- Coursera: Algorithms Specialization
- LeetCode: Dynamic Programming card

### Practice Platforms
- LeetCode: DP problems sorted by frequency
- AtCoder: Educational DP contest
- Codeforces: DP problems in contests

---

*Last Updated: 2026-01-31*  
*Next: Create Graph Algorithms overview*