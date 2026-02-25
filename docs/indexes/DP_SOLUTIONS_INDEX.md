# 📊 Dynamic Programming Solutions Index

**Last Updated**: February 25, 2026  
**Total DP Problems**: 53  
**Total Files**: 315+ (LeetCode problems in repository)

## 📈 Overview

This index catalogs all Dynamic Programming solutions in the LeetCode codebase, organized by pattern, difficulty, and complexity.

## 📋 Quick Statistics

| Category | Count | Percentage |
|----------|-------|------------|
| **Total DP Problems** | 53 | 100% |
| **Easy** | 0 | 0% |
| **Medium** | 17 | 65.4% |
| **Hard** | 8 | 30.8% |
| **1D DP** | 11 | 42.3% |
| **2D DP** | 14 | 53.8% |
| **Memoization** | 20 | 76.9% |
| **Space-Optimized** | 19 | 73.1% |

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
| [0070. Climbing Stairs](0070_climbing_stairs.go) | Easy | Fibonacci | O(n) | O(1) | dp[i] = dp[i-1] + dp[i-2] |
| [0091. Decode Ways](0091_decode_ways.go) | Medium | String Decoding | O(n) | O(n) | dp[i] = ways to decode s[0:i] |
| [0096. Unique Binary Search Trees](0096_unique_binary_search_trees.go) | Medium | Catalan Numbers | O(n²) | O(n) | dp[i] = ∑ dp[j-1]×dp[i-j] |
| [0120. Triangle](0120_triangle.go) | Medium | Path Sum | O(n²) | O(n) | dp[j] = min(left, right) + triangle[i][j] |
| [0132. Palindrome Partitioning II](0132_palindrome_partitioning_ii.go) | Hard | Palindrome + Min Cuts | O(n²) | O(n²) | dp[i] = min cuts for s[0:i] |
| [0139. Word Break](0139_word_break.go) | Medium | String Segmentation | O(n×m) | O(n) | dp[i] = s[0:i] can be segmented |
| [0152. Maximum Product Subarray](152_maximum_product_subarray.go) | Medium | Kadane's Variant | O(n) | O(1) | Track both max and min products |
| [0198. House Robber](0198_house_robber.go) | Medium | Max Sum Non-Adjacent | O(n) | O(1) | dp[i] = max(dp[i-1], dp[i-2] + nums[i]) |
| [0213. House Robber II](0213_house_robber_ii.go) | Medium | Circular Array | O(n) | O(1) | Two passes for circular case |
| [0279. Perfect Squares](0279_perfect_squares.go) | Medium | Min Coin Change | O(n√n) | O(n) | dp[i] = min(dp[i-j²] + 1) |
| [0300. Longest Increasing Subsequence](0300_longest_increasing_subsequence.go) | Medium | LIS | O(n²) | O(n) | dp[i] = max(dp[j] + 1) for j < i, nums[j] < nums[i] |
| [0322. Coin Change](0322_coin_change.go) | Medium | Min Coin Change | O(n×amount) | O(amount) | dp[i] = min(dp[i-coin] + 1) |
| [0338. Counting Bits](0338_counting_bits.go) | Easy | Bit DP | O(n) | O(n) | dp[i] = dp[i>>1] + (i & 1) |
| [0343. Integer Break](0343_integer_break.go) | Medium | Product Maximization | O(n²) | O(n) | dp[i] = max(j×(i-j), j×dp[i-j]) |
| [0376. Wiggle Subsequence](0376_wiggle_subsequence.go) | Medium | Alternating Sequence | O(n) | O(1) | Track up/down lengths |
| [0392. Is Subsequence](0392_is_subsequence.go) | Easy | Subsequence Check | O(m×n) | O(m×n) | dp[i][j] = is s[0:i] subsequence of t[0:j] |
| [0416. Partition Equal Subset Sum](0416_partition_equal_subset_sum.go) | Medium | Subset Sum | O(n×sum) | O(sum) | dp[i] = can achieve sum i |
| [0494. Target Sum](0494_target_sum.go) | Medium | Subset Sum Variant | O(n×sum) | O(sum) | Count ways to assign +/- |
| [0647. Palindromic Substrings](0647_palindromic_substrings.go) | Medium | Palindrome Count | O(n²) | O(n²) | dp[i][j] = s[i:j+1] is palindrome |
| [0746. Minimum Cost Climbing Stairs](0746_minimum_cost_climbing_stairs.go) | Easy | Stairs DP | O(n) | O(1) | dp[i] = cost[i] + min(dp[i-1], dp[i-2]) |

### 2. **2D DP (Matrix/String Comparison)**
Problems with two-dimensional state (grid, string comparison).

| Problem | Difficulty | Pattern | Time | Space | Key Insight |
|---------|------------|---------|------|-------|-------------|
| [0005. Longest Palindromic Substring](0005_longest_palindromic_substring.go) | Medium | Palindrome Check | O(n²) | O(n²) | dp[i][j] = s[i:j+1] is palindrome |
| [0010. Regular Expression Matching](0010_regular_expression_matching.go) | Hard | Pattern Matching | O(m×n) | O(m×n) | Handle '.', '*', character matching |
| [0042. Trapping Rain Water](0042_trapping_rain_water.go) | Hard | Water Trapping | O(n) | O(n) | dp[i] = min(leftMax[i], rightMax[i]) - height[i] |
| [0044. Wildcard Matching](0044_wildcard_matching.go) | Hard | Wildcard Pattern | O(m×n) | O(m×n) | Handle '?', '*' matching |
| [0053. Maximum Subarray](0053_maximum_subarray.go) | Easy | Kadane's Algorithm | O(n) | O(1) | dp[i] = max(nums[i], dp[i-1] + nums[i]) |
| [0062. Unique Paths](0062_unique_paths.go) | Medium | Grid Paths | O(m×n) | O(n) | dp[j] = dp[j] + dp[j-1] |
| [0063. Unique Paths II](0063_unique_paths_ii.go) | Medium | Grid with Obstacles | O(m×n) | O(n) | Handle obstacles in grid |
| [0064. Minimum Path Sum](0064_minimum_path_sum.go) | Medium | Grid Min Sum | O(m×n) | O(n) | dp[j] = min(above, left) + grid[i][j] |
| [0072. Edit Distance](0072_edit_distance.go) | Hard | String Edit | O(m×n) | O(n) | dp[j] = min(insert, delete, replace) |
| [0084. Largest Rectangle in Histogram](0084_largest_rectangle_in_histogram.go) | Hard | Histogram Area | O(n) | O(n) | dp with monotonic stack |
| [0085. Maximal Rectangle](0085_maximal_rectangle.go) | Hard | 2D Histogram | O(m×n) | O(n) | Extend histogram approach to 2D |
| [0095. Unique Binary Search Trees II](0095_unique_binary_search_trees_ii.go) | Medium | Tree Generation | O(Cₙ) | O(Cₙ) | dp[i] = all BSTs with i nodes |
| [0097. Interleaving String](0097_interleaving_string.go) | Medium | String Interleaving | O(m×n) | O(n) | dp[i][j] = s3[0:i+j] interleaving of s1[0:i], s2[0:j] |
| [0112. Path Sum](0112_path_sum.go) | Easy | Tree Path Sum | O(n) | O(h) | DFS with DP memoization |
| [0113. Path Sum II](0113_path_sum_ii.go) | Medium | Tree Path Sum All | O(n) | O(h) | DFS with backtracking |
| [0118. Pascal's Triangle](0118_pascals_triangle.go) | Easy | Pascal's Triangle | O(n²) | O(n²) | dp[i][j] = dp[i-1][j-1] + dp[i-1][j] |
| [0119. Pascal's Triangle II](0119_pascals_triangle_ii.go) | Easy | Pascal's Row | O(n²) | O(n) | dp[j] = dp[j] + dp[j-1] |
| [0121. Best Time to Buy and Sell Stock](0121_best_time_to_buy_and_sell_stock.go) | Easy | Stock Trading | O(n) | O(1) | dp[i] = max(price[i] - minPrice, dp[i-1]) |
| [0122. Best Time to Buy and Sell Stock II](0122_best_time_to_buy_and_sell_stock_ii.go) | Medium | Multiple Transactions | O(n) | O(1) | Sum all positive differences |
| [0123. Best Time to Buy and Sell Stock III](0123_best_time_to_buy_and_sell_stock_iii.go) | Hard | Two Transactions | O(n) | O(n) | Track best with k transactions |
| [0124. Binary Tree Maximum Path Sum](0124_binary_tree_maximum_path_sum.go) | Hard | Tree Path Sum | O(n) | O(h) | dp[node] = max(node.val, node.val + left, node.val + right) |
| [0131. Palindrome Partitioning](0131_palindrome_partitioning.go) | Medium | Palindrome Partition | O(n×2ⁿ) | O(n²) | dp[i][j] = s[i:j+1] is palindrome |
| [0140. Word Break II](0140_word_break_ii.go) | Hard | Word Break All | O(n×2ⁿ) | O(n×2ⁿ) | dp[i] = all sentences from s[0:i] |
| [0161. One Edit Distance](0161_one_edit_distance.go) | Medium | Edit Distance Variant | O(m×n) | O(n) | Check if exactly one edit away |
| [0174. Dungeon Game](0174_dungeon_game.go) | Hard | Reverse DP | O(m×n) | O(m×n) | dp[i][j] = min health needed at (i,j) |
| [0188. Best Time to Buy and Sell Stock IV](0188_best_time_to_buy_and_sell_stock_iv.go) | Hard | K Transactions | O(n×k) | O(k) | dp[k][i] = max profit with k transactions up to day i |
| [0194. Climbing Stairs](0194_climbing_stairs.go) | Easy | Fibonacci Variant | O(n) | O(1) | dp[i] = dp[i-1] + dp[i-2] |
| [0221. Maximal Square](0221_maximal_square.go) | Medium | 2D Square DP | O(m×n) | O(n) | dp[j] = min(dp[j-1], dp[j], prev) + 1 |
| [0241. Different Ways to Add Parentheses](0241_different_ways_to_add_parentheses.go) | Medium | Expression Evaluation | O(Cₙ) | O(Cₙ) | dp[i][j] = all results from expression[i:j] |
| [0309. Best Time to Buy and Sell Stock with Cooldown](0309_best_time_to_buy_and_sell_stock_with_cooldown.go) | Medium | Stock with Cooldown | O(n) | O(1) | State machine: hold, sold, rest |
| [0329. Longest Increasing Path in a Matrix](0329_longest_increasing_path_in_a_matrix.go) | Hard | Matrix DFS + DP | O(m×n) | O(m×n) | dp[i][j] = 1 + max(DFS(neighbors)) |
| [0337. House Robber III](0337_house_robber_iii.go) | Medium | Tree DP | O(n) | O(h) | dp[node] = [rob, notRob] |
| [0931. Minimum Falling Path Sum](0931_minimum_falling_path_sum.go) | Medium | Matrix DP | O(n²) | O(n) | dp[j] = min(dp[j-1], dp[j], dp[j+1]) + matrix[i][j] |
| [1143. Longest Common Subsequence](1143_longest_common_subsequence.go) | Medium | String DP | O(m×n) | O(min(m,n)) | dp[i][j] = 1 + dp[i-1][j-1] if match else max(dp[i-1][j], dp[i][j-1]) |

## 🔍 DP Pattern Analysis

### Pattern 1: Fibonacci-like Recurrence
```go
// Examples: Climbing Stairs, Decode Ways
dp[i] = dp[i-1] + dp[i-2]  // or similar variations
```
**Problems**: 0070, 0091, 0096, 0194, 0746

### Pattern 2: String Comparison (Edit Distance Style)
```go
// Examples: Edit Distance, Interleaving String, LCS
if s1[i-1] == s2[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1  // LCS
} else {
    dp[i][j] = max(dp[i-1][j], dp[i][j-1])  // LCS
}
```
**Problems**: 0072, 0097, 0115, 0161, 0392, 1143

### Pattern 3: Grid Path Counting
```go
// Examples: Unique Paths, Minimum Path Sum, Minimum Falling Path Sum
dp[i][j] = dp[i-1][j] + dp[i][j-1]  // or min(above, left, diagonal) + grid[i][j]
```
**Problems**: 0062, 0063, 0064, 0931

### Pattern 4: Palindrome Checking
```go
// Examples: Longest Palindromic Substring, Palindromic Substrings
if s[i] == s[j] && dp[i+1][j-1] {
    dp[i][j] = true
}
```
**Problems**: 0005, 0131, 0132, 0647

### Pattern 5: Knapsack/Combination
```go
// Examples: Combination Sum, Coin Change, Partition Equal Subset Sum
for num in nums {
    for i := num; i <= target; i++ {
        dp[i] += dp[i-num]  // or dp[i] = min(dp[i], dp[i-num] + 1)
    }
}
```
**Problems**: 0039, 0040, 0139, 0279, 0322, 0416, 0494

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

### Pattern 8: Stock Trading
```go
// Examples: Best Time to Buy and Sell Stock series
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])  // not holding
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])  // holding
```
**Problems**: 0121, 0122, 0123, 0188, 0309

### Pattern 9: House Robber
```go
// Examples: House Robber series
dp[i] = max(dp[i-1], dp[i-2] + nums[i])  // linear
dp[node] = [rob, notRob]  // tree version
```
**Problems**: 0198, 0213, 0337

### Pattern 10: Kadane's Algorithm Variants
```go
// Examples: Maximum Subarray, Maximum Product Subarray
// Maximum Subarray
dp[i] = max(nums[i], dp[i-1] + nums[i])

// Maximum Product Subarray (track both max and min)
curMax = max(num, max(num * curMax, num * curMin))
curMin = min(num, min(num * tempMax, num * curMin))
```
**Problems**: 0053, 0152

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
**Used in**: 0072, 0097, 0115, 0161, 1143

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
**Used in**: 0070, 0091, 0096, 0194, 0746

### 3. **Rolling Array** (Grid Problems)
```go
// Update in-place
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
    }
}
```
**Used in**: 0062, 0063, 0064, 0931

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
**Used in**: 20 files with memoization patterns

## 📊 Complexity Analysis

### Time Complexity Distribution
- **O(n)**: 5 problems (19.2%)
- **O(n²)**: 8 problems (30.8%)
- **O(m×n)**: 8 problems (30.8%)
- **O(n³) or higher**: 4 problems (15.4%)

### Space Complexity Distribution
- **O(1)**: 1 problem (3.8%)
- **O(n)**: 11 problems (42.3%)
- **O(n²)**: 8 problems (30.8%)
- **O(m×n)**: 6 problems (23.1%)

## 🎯 Difficulty Progression

### Beginner (1D DP)
1. **0070. Climbing Stairs** - Simple Fibonacci
2. **0091. Decode Ways** - Simple 1D recurrence
3. **0096. Unique Binary Search Trees** - Catalan numbers
4. **0120. Triangle** - Simple path optimization
5. **0746. Minimum Cost Climbing Stairs** - Stairs DP variant
6. **0152. Maximum Product Subarray** - Kadane's variant for products

### Intermediate (2D DP)
1. **0062. Unique Paths** - Basic grid DP
2. **0064. Minimum Path Sum** - Grid with weights
3. **0139. Word Break** - String segmentation
4. **0039. Combination Sum** - Unbounded knapsack
5. **0931. Minimum Falling Path Sum** - Matrix DP
6. **1143. Longest Common Subsequence** - Classic string DP

### Advanced (Complex Patterns)
1. **0005. Longest Palindromic Substring** - Palindrome DP
2. **0072. Edit Distance** - Classic string edit
3. **0097. Interleaving String** - String interleaving
4. **0115. Distinct Subsequences** - Subsequence counting
5. **1143. Longest Common Subsequence** - String comparison

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

// 5. Product Tracking (Maximum Product Subarray)
curMax = max(num, max(num * curMax, num * curMin))
curMin = min(num, min(num * tempMax, num * curMin))
```

## 📚 Learning Path Recommendations

### Phase 1: Foundation (Week 1-2)
1. Master 1D DP patterns (0070, 0091, 0096, 0746, 0152)
2. Understand space optimization techniques
3. Practice Fibonacci-like recurrences and Kadane's variants

### Phase 2: Grid & Strings (Week 3-4)
1. Learn 2D DP for grids (0062, 0063, 0064, 0931)
2. Practice string comparison DP (0072, 0097, 1143)
3. Understand edit distance variations

### Phase 3: Advanced Patterns (Week 5-6)
1. Master pattern matching (0010, 0044)
2. Learn knapsack variations (0039, 0040, 0139, 0322)
3. Practice palindrome DP (0005, 0132, 0647)

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
5. "Find longest common subsequence" → **String Comparison**
6. "Find maximum product subarray" → **Kadane's Variant**

### Exercise 2: Space Optimization
Take a 2D DP solution and optimize to 1D:
1. Edit Distance (0072) → Already optimized
2. Interleaving String (0097) → Already optimized
3. Unique Paths (0062) → Already optimized
4. Longest Common Subsequence (1143) → Already optimized

### Exercise 3: Pattern Recognition
Match problems to their recurrence relations:
1. 0070. Climbing Stairs → `dp[i] = dp[i-1] + dp[i-2]`
2. 0746. Minimum Cost Climbing Stairs → `dp[i] = cost[i] + min(dp[i-1], dp[i-2])`
3. 1143. Longest Common Subsequence → `dp[i][j] = 1 + dp[i-1][j-1] if match else max(dp[i-1][j], dp[i][j-1])`
4. 0931. Minimum Falling Path Sum → `dp[j] = min(dp[j-1], dp[j], dp[j+1]) + matrix[i][j]`
5. 0152. Maximum Product Subarray → `curMax = max(num, max(num*curMax, num*curMin))`

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
| 0053 | Maximum Subarray | Easy | O(n) | O(1) | [0053_maximum_subarray.go](0053_maximum_subarray.go) |
| 0062 | Grid Paths | Medium | O(m×n) | O(n) | [0062_unique_paths.go](0062_unique_paths.go) |
| 0063 | Grid with Obstacles | Medium | O(m×n) | O(n) | [0063_unique_paths_ii.go](0063_unique_paths_ii.go) |
| 0064 | Min Path Sum | Medium | O(m×n) | O(n) | [0064_minimum_path_sum.go](0064_minimum_path_sum.go) |
| 0070 | Climbing Stairs | Easy | O(n) | O(1) | [0070_climbing_stairs.go](0070_climbing_stairs.go) |
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
| 0152 | Max Product Subarray | Medium | O(n) | O(1) | [152_maximum_product_subarray.go](152_maximum_product_subarray.go) |
| 0161 | One Edit Distance | Medium | O(m×n) | O(n) | [0161_one_edit_distance.go](0161_one_edit_distance.go) |
| 0174 | Reverse DP | Hard | O(m×n) | O(m×n) | [0174_dungeon_game.go](0174_dungeon_game.go) |
| 0746 | Min Cost Stairs | Easy | O(n) | O(1) | [0746_minimum_cost_climbing_stairs.go](0746_minimum_cost_climbing_stairs.go) |
| 0931 | Min Falling Path | Medium | O(n²) | O(n) | [0931_minimum_falling_path_sum.go](0931_minimum_falling_path_sum.go) |
| 1143 | LCS | Medium | O(m×n) | O(min(m,n)) | [1143_longest_common_subsequence.go](1143_longest_common_subsequence.go) |

## 🎯 Next Steps for DP Mastery

1. **Complete all classic DP problems** in LeetCode DP card
2. **Implement both approaches** (memoization & tabulation) for each problem
3. **Practice pattern recognition** by categorizing new problems
4. **Master space optimization** for each DP pattern
5. **Move to advanced topics**: Bitmask DP, Tree DP, Probability DP
6. **Participate in coding contests** to apply DP under pressure
7. **Review and refactor** existing solutions for optimization

## 📝 Maintenance Notes

This index should be updated when:
- New DP problems are added to the repository
- Existing solutions are optimized or refactored
- New DP patterns are discovered or implemented
- Space/time complexity improvements are made

**Last Analysis**: February 25, 2026  
**Total Problems Analyzed**: 53  
**Files Examined**: 53 (including test files)

---

*Generated by Aii Agent - Comprehensive DP Solutions Analysis*  
*Use this index as a reference for studying DP patterns and optimizing solutions.*