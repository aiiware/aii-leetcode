# 🎯 Dynamic Programming Practice Exercises

**Last Updated**: January 28, 2026  
**Total Exercises**: 50+ problems organized by pattern

## 📋 How to Use This Guide

1. **Start with your level** - Beginner, Intermediate, or Advanced
2. **Focus on one pattern at a time** - Master it before moving on
3. **Time yourself** - 30 minutes per problem initially
4. **Implement both approaches** - Memoization AND tabulation
5. **Optimize space** - Try to reduce space complexity

## 🏆 Beginner Level (Pattern Recognition)

### Exercise Set 1: Fibonacci-like DP
**Goal**: Recognize and implement simple 1D DP

1. **0070. Climbing Stairs** (Easy)
   - Pattern: Fibonacci sequence
   - Recurrence: `dp[i] = dp[i-1] + dp[i-2]`
   - Space: Optimize to O(1)

2. **0198. House Robber** (Medium)
   - Pattern: 1D DP with constraints
   - Recurrence: `dp[i] = max(dp[i-1], dp[i-2] + nums[i])`
   - Constraint: Can't rob adjacent houses

3. **0213. House Robber II** (Medium)
   - Pattern: Circular 1D DP
   - Hint: Solve for [0:n-1] and [1:n], take max

4. **0746. Min Cost Climbing Stairs** (Easy)
   - Pattern: Fibonacci with costs
   - Recurrence: `dp[i] = cost[i] + min(dp[i-1], dp[i-2])`

### Exercise Set 2: Grid Path DP
**Goal**: Master 2D grid traversal

5. **0062. Unique Paths** (Medium) - *Already implemented*
   - Pattern: Grid path counting
   - Recurrence: `dp[i][j] = dp[i-1][j] + dp[i][j-1]`
   - Space: Optimize to O(n)

6. **0063. Unique Paths II** (Medium) - *Already implemented*
   - Pattern: Grid with obstacles
   - Recurrence: `dp[i][j] = obstacle ? 0 : dp[i-1][j] + dp[i][j-1]`

7. **0064. Minimum Path Sum** (Medium) - *Already implemented*
   - Pattern: Grid with weights
   - Recurrence: `dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])`

8. **0931. Minimum Falling Path Sum** (Medium)
   - Pattern: Grid with 3 possible moves
   - Recurrence: `dp[i][j] = matrix[i][j] + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])`

## 📈 Intermediate Level (Pattern Application)

### Exercise Set 3: String Comparison DP
**Goal**: Master string transformation problems

9. **0072. Edit Distance** (Hard) - *Already implemented*
   - Pattern: Classic string edit
   - Operations: Insert, delete, replace
   - Space: Optimize to O(n)

10. **0115. Distinct Subsequences** (Hard) - *Already implemented*
    - Pattern: Subsequence counting
    - Recurrence: `dp[i][j] = dp[i-1][j] + (s[i-1]==t[j-1] ? dp[i-1][j-1] : 0)`

11. **0583. Delete Operation for Two Strings** (Medium)
    - Pattern: Edit distance variant
    - Only deletions allowed
    - Hint: Find LCS, then delete rest

12. **0712. Minimum ASCII Delete Sum for Two Strings** (Medium)
    - Pattern: Edit distance with ASCII costs
    - Recurrence: Similar to edit distance but with ASCII values

### Exercise Set 4: Knapsack DP
**Goal**: Master selection problems

13. **0039. Combination Sum** (Medium) - *Already implemented*
    - Pattern: Unbounded knapsack
    - Items can be reused
    - Find all combinations

14. **0040. Combination Sum II** (Medium) - *Already implemented*
    - Pattern: 0/1 knapsack with duplicates
    - Each item used once
    - Handle duplicates carefully

15. **0416. Partition Equal Subset Sum** (Medium)
    - Pattern: 0/1 knapsack
    - Target: sum/2
    - Check if possible

16. **0494. Target Sum** (Medium)
    - Pattern: 0/1 knapsack with +/- signs
    - Convert to subset sum problem

### Exercise Set 5: Palindrome DP
**Goal**: Master palindrome-related problems

17. **0005. Longest Palindromic Substring** (Medium) - *Already implemented*
    - Pattern: Palindrome expansion
    - Time: O(n²) DP, O(n) Manacher's
    - Space: O(n²) or O(1)

18. **0647. Palindromic Substrings** (Medium)
    - Pattern: Count all palindromic substrings
    - Similar to #5 but count instead of find longest

19. **0516. Longest Palindromic Subsequence** (Medium)
    - Pattern: Subsequence (not substring)
    - Recurrence: `dp[i][j] = s[i]==s[j] ? 2+dp[i+1][j-1] : max(dp[i+1][j], dp[i][j-1])`

20. **0132. Palindrome Partitioning II** (Hard) - *Already implemented*
    - Pattern: Palindrome + min cuts
    - Two-step DP: palindrome check then min cuts

## 🚀 Advanced Level (Pattern Combination)

### Exercise Set 6: Pattern Matching DP
**Goal**: Master regex and wildcard matching

21. **0010. Regular Expression Matching** (Hard) - *Already implemented*
    - Pattern: Regex with '.' and '*'
    - Handle zero-or-more of preceding element

22. **0044. Wildcard Matching** (Hard) - *Already implemented*
    - Pattern: Wildcard with '?' and '*'
    - '*' matches any sequence

23. **0045. Jump Game II** (Medium) - *Already implemented*
    - Pattern: Greedy/DP hybrid
    - Find minimum jumps to reach end

24. **0055. Jump Game** (Medium)
    - Pattern: Greedy/DP
    - Check if can reach end
    - Simpler than #45

### Exercise Set 7: Parentheses DP
**Goal**: Master parentheses problems

25. **0022. Generate Parentheses** (Medium) - *Already implemented*
    - Pattern: Catalan numbers
    - Generate all valid combinations

26. **0032. Longest Valid Parentheses** (Hard) - *Already implemented*
    - Pattern: Parentheses validation
    - Find longest valid substring

27. **0301. Remove Invalid Parentheses** (Hard)
    - Pattern: BFS/DFS with pruning
    - Remove min parentheses to make valid

28. **0678. Valid Parenthesis String** (Medium)
    - Pattern: Parentheses with wildcard '*'
    - '*' can be '(', ')', or empty

### Exercise Set 8: Advanced String DP
**Goal**: Master complex string problems

29. **0097. Interleaving String** (Medium) - *Already implemented*
    - Pattern: String interleaving
    - Check if s3 is interleaving of s1 and s2

30. **0087. Scramble String** (Hard) - *Already implemented*
    - Pattern: 3D DP
    - Check if strings are scramble

31. **0139. Word Break** (Medium) - *Already implemented*
    - Pattern: String segmentation
    - Check if can be segmented into dictionary words

32. **0140. Word Break II** (Hard)
    - Pattern: Word break with output
    - Return all possible segmentations

## 🏅 Expert Level (Hard Problems)

### Exercise Set 9: Hard DP Problems
**Goal**: Solve the toughest DP problems

33. **0174. Dungeon Game** (Hard) - *Already implemented*
    - Pattern: Reverse DP
    - Calculate minimum initial health

34. **0312. Burst Balloons** (Hard)
    - Pattern: Interval DP
    - Recurrence: `dp[i][j] = max(dp[i][k-1] + nums[i-1]*nums[k]*nums[j+1] + dp[k+1][j])`

35. **0329. Longest Increasing Path in a Matrix** (Hard)
    - Pattern: DFS + memoization
    - Find longest increasing path in grid

36. **0354. Russian Doll Envelopes** (Hard)
    - Pattern: 2D LIS
    - Sort + LIS on second dimension

### Exercise Set 10: Bitmask DP
**Goal**: Learn state compression

37. **0464. Can I Win** (Medium)
    - Pattern: Game theory + bitmask
    - Players take turns picking numbers

38. **0526. Beautiful Arrangement** (Medium)
    - Pattern: Bitmask DP
    - Count arrangements where number divides index or vice versa

39. **0691. Stickers to Spell Word** (Hard)
    - Pattern: Bitmask + BFS
    - Minimum stickers to spell target word

40. **0847. Shortest Path Visiting All Nodes** (Hard)
    - Pattern: Bitmask BFS
    - Shortest path visiting all nodes in graph

## 📊 Progress Tracking

### Weekly Study Plan

**Week 1-2: Foundation**
- [ ] Complete all Fibonacci-like problems (4)
- [ ] Complete all Grid Path problems (4)
- [ ] Master space optimization techniques

**Week 3-4: Core Patterns**
- [ ] Complete String Comparison problems (4)
- [ ] Complete Knapsack problems (4)
- [ ] Complete Palindrome problems (4)

**Week 5-6: Advanced Patterns**
- [ ] Complete Pattern Matching problems (4)
- [ ] Complete Parentheses problems (4)
- [ ] Complete Advanced String problems (4)

**Week 7-8: Expert Level**
- [ ] Complete Hard DP problems (4)
- [ ] Complete Bitmask DP problems (4)
- [ ] Review and optimize all solutions

### Performance Metrics

Track your performance for each problem:
1. **Time to solve** (first attempt)
2. **Time complexity** achieved
3. **Space complexity** achieved
4. **Number of attempts** needed
5. **Optimization level** (1-5 scale)

### Pattern Mastery Checklist

For each pattern, check when you can:
- [ ] **Recognize** the pattern in new problems
- [ ] **Implement** both memoization and tabulation
- [ ] **Optimize** space complexity
- [ ] **Explain** the recurrence relation
- [ ] **Solve** 3+ problems of this pattern

## 🎯 Challenge Problems

### Mixed Pattern Challenges
These combine multiple patterns:

41. **0085. Maximal Rectangle** (Hard)
    - Pattern: Histogram DP + monotonic stack
    - Find largest rectangle in binary matrix

42. **0221. Maximal Square** (Medium)
    - Pattern: 2D DP for squares
    - Recurrence: `dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1`

43. **0304. Range Sum Query 2D - Immutable** (Medium)
    - Pattern: Prefix sum DP
    - Precompute sums for O(1) queries

44. **0363. Max Sum of Rectangle No Larger Than K** (Hard)
    - Pattern: 2D prefix sum + binary search

### Optimization Challenges
Focus on optimizing existing solutions:

45. **Take any 2D DP problem** and optimize to 1D
46. **Take any O(n²) DP** and try to optimize to O(n log n)
47. **Implement iterative solution** for recursive DP
48. **Reduce space** from O(n) to O(1) where possible

## 🔍 Pattern Recognition Drills

### Quick Identification
Given problem statement, identify pattern in 30 seconds:

1. "Count ways to decode digits to letters" → **Fibonacci-like**
2. "Minimum path sum in triangle" → **Grid DP**
3. "Transform word1 to word2 with min operations" → **String Comparison**
4. "Check if string matches pattern with '*'" → **Pattern Matching**
5. "Find combinations that sum to target" → **Knapsack**
6. "Longest palindrome substring" → **Palindrome**
7. "Generate valid parentheses" → **Parentheses**
8. "Check if s3 is interleaving of s1 and s2" → **Advanced String**

### Recurrence Writing
Write recurrence relations for:

1. **Coin Change (min coins)**: `dp[i] = min(dp[i], dp[i-coin] + 1)`
2. **Longest Common Subsequence**: `dp[i][j] = s1[i]==s2[j] ? 1+dp[i-1][j-1] : max(dp[i-1][j], dp[i][j-1])`
3. **Maximum Subarray**: `dp[i] = max(nums[i], dp[i-1] + nums[i])`
4. **Best Time to Buy/Sell Stock**: `dp[i] = max(dp[i-1], prices[i] - minPrice)`

## 📚 Resources

### In This Repository
- `DP_SOLUTIONS_INDEX.md` - Complete index of DP solutions
- `DP_PATTERN_ANALYSIS.md` - Detailed pattern analysis
- `DP_QUICK_REFERENCE.md` - Quick reference guide
- Individual problem files (`####_problem_name.go`)

### External Resources
- **LeetCode DP Card** - Curated list of DP problems
- **"Introduction to Algorithms"** - CLRS DP chapter
- **NeetCode 150** - DP section
- **AlgoExpert** - DP category

### Practice Platforms
- **LeetCode** - Filter by "Dynamic Programming" tag
- **HackerRank** - DP challenges
- **Codeforces** - DP problems in contests
- **AtCoder** - DP educational contests

## 🎯 Final Challenge

### The Ultimate DP Test
Solve these in one sitting (4 hours):

1. **0072. Edit Distance** (Hard)
2. **0322. Coin Change** (Medium)
3. **0300. Longest Increasing Subsequence** (Medium)
4. **0416. Partition Equal Subset Sum** (Medium)
5. **0312. Burst Balloons** (Hard)

**Success Criteria**:
- All solutions accepted on LeetCode
- Optimal time/space complexity
- Clean, well-commented code
- Under 30 minutes per problem

### Mastery Certification
You've mastered DP when you can:
1. **Recognize** DP patterns within 1 minute of reading problem
2. **Implement** correct solution in under 20 minutes
3. **Optimize** space complexity appropriately
4. **Explain** your solution clearly to others
5. **Teach** DP concepts to beginners

## 📈 Tracking Your Progress

### Create a Progress Log
```markdown
# DP Progress Log - [Your Name]

## Week 1: Fibonacci & Grid DP
- [x] 0070. Climbing Stairs - 15min, O(n) time, O(1) space
- [x] 0062. Unique Paths - 20min, O(mn) time, O(n) space
- [ ] 0064. Minimum Path Sum - TODO

## Week 2: String & Knapsack DP
- [ ] 0072. Edit Distance
- [ ] 0039. Combination Sum
- [ ] 0416. Partition Equal Subset Sum

## Pattern Mastery:
- [x] Fibonacci-like (3/3 problems)
- [x] Grid DP (3/4 problems)
- [ ] String Comparison (1/4 problems)
- [ ] Knapsack (0/4 problems)
```

### Review Schedule
- **Daily**: Solve 1-2 DP problems
- **Weekly**: Review one pattern category
- **Monthly**: Take mock interview with DP focus
- **Quarterly**: Re-solve all problems for retention

---
*Use this guide systematically to master Dynamic Programming.*  
*Start with patterns you find challenging, track your progress, and celebrate milestones!*

*Generated by Aii Agent - DP Practice Guide*  
*Last updated: January 28, 2026*