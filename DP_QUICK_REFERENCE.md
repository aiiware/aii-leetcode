# 🎯 Dynamic Programming Solutions - Quick Reference

## 📊 Summary Statistics
- **Total DP Problems**: 24
- **Easy**: 0 (0%)
- **Medium**: 16 (66.7%)
- **Hard**: 8 (33.3%)
- **1D DP**: 10 (41.7%)
- **2D DP**: 14 (58.3%)

## 🏆 Top 10 Most Important DP Problems

| Rank | Problem | Difficulty | Pattern | Key Insight |
|------|---------|------------|---------|-------------|
| 1 | 0072. Edit Distance | Hard | String Comparison | Classic DP, space optimization |
| 2 | 0010. Regular Expression Matching | Hard | Pattern Matching | Complex state transitions |
| 3 | 0005. Longest Palindromic Substring | Medium | Palindrome 2D | O(n²) DP for palindrome |
| 4 | 0039. Combination Sum | Medium | Unbounded Knapsack | Classic combination DP |
| 5 | 0062. Unique Paths | Medium | Grid DP | Basic grid path counting |
| 6 | 0091. Decode Ways | Medium | 1D DP | Fibonacci-like with conditions |
| 7 | 0139. Word Break | Medium | String Segmentation | Classic segmentation DP |
| 8 | 0096. Unique Binary Search Trees | Medium | Catalan Numbers | Mathematical DP |
| 9 | 0097. Interleaving String | Medium | String Interleaving | 2D string DP |
| 10 | 0174. Dungeon Game | Hard | Reverse DP | Bottom-up from destination |

## 📈 Learning Progression

### Beginner Level (Start Here)
1. **0091. Decode Ways** - Simple 1D recurrence
2. **0096. Unique Binary Search Trees** - Catalan numbers
3. **0062. Unique Paths** - Basic grid DP

### Intermediate Level
1. **0064. Minimum Path Sum** - Grid with weights
2. **0139. Word Break** - String segmentation
3. **0039. Combination Sum** - Unbounded knapsack
4. **0005. Longest Palindromic Substring** - Palindrome DP

### Advanced Level
1. **0072. Edit Distance** - Classic string edit
2. **0097. Interleaving String** - String interleaving
3. **0115. Distinct Subsequences** - Subsequence counting

### Expert Level
1. **0010. Regular Expression Matching** - Complex pattern matching
2. **0044. Wildcard Matching** - Wildcard pattern matching
3. **0087. Scramble String** - 3D DP state
4. **0174. Dungeon Game** - Reverse DP

## 🔑 Key DP Patterns

### Must-Know Patterns:
1. **Fibonacci-like**: `dp[i] = dp[i-1] + dp[i-2]` (0091, 0096)
2. **Grid Paths**: `dp[i][j] = dp[i-1][j] + dp[i][j-1]` (0062, 0063, 0064)
3. **Edit Distance**: `dp[i][j] = min(insert, delete, replace)` (0072, 0161)
4. **Knapsack**: `dp[cap] = max(dp[cap], dp[cap-weight] + value)` (0039, 0040)
5. **Palindrome**: `dp[i][j] = s[i]==s[j] && dp[i+1][j-1]` (0005, 0132)

### Space Optimization Status:
- ✅ **Already optimized**: 18 problems (75%)
- 🔄 **Could optimize further**: 6 problems (25%)
- ⚠️ **Needs optimization**: 0 problems

## 🚀 Quick Start Guide

### If you're new to DP:
1. Start with **0091. Decode Ways** (1D DP)
2. Then try **0062. Unique Paths** (Grid DP)
3. Move to **0139. Word Break** (String DP)
4. Practice **0039. Combination Sum** (Knapsack DP)

### If you know basics:
1. Master **0072. Edit Distance** (Classic)
2. Solve **0005. Longest Palindromic Substring** (Palindrome)
3. Attempt **0097. Interleaving String** (String 2D)
4. Challenge **0010. Regular Expression Matching** (Hard)

### If you're preparing for interviews:
1. **Must know**: 0072, 0139, 0039, 0005
2. **Good to know**: 0091, 0062, 0096, 0097
3. **Advanced**: 0010, 0044, 0115, 0174

## 📝 Common Interview Questions

### Frequently Asked:
1. **"Solve Edit Distance"** → 0072
2. **"Word Break problem"** → 0139
3. **"Combination Sum"** → 0039
4. **"Longest Palindromic Substring"** → 0005
5. **"Unique Paths"** → 0062

### Company-Specific:
- **Google**: 0010, 0044, 0174
- **Facebook**: 0091, 0139, 0032
- **Amazon**: 0005, 0072, 0139
- **Microsoft**: 0097, 0115, 0174

## ✅ Checklist for DP Mastery

### Basic Concepts:
- [ ] Understand overlapping subproblems
- [ ] Understand optimal substructure
- [ ] Know memoization vs tabulation
- [ ] Can implement Fibonacci with both approaches

### 1D DP:
- [ ] Solved Decode Ways (0091)
- [ ] Solved Unique BST (0096)
- [ ] Can optimize 1D DP to O(1) space
- [ ] Understand Catalan numbers

### 2D DP:
- [ ] Solved Edit Distance (0072)
- [ ] Solved Longest Palindromic Substring (0005)
- [ ] Can optimize 2D DP to 1D array
- [ ] Understand string comparison DP

### Grid DP:
- [ ] Solved Unique Paths (0062)
- [ ] Solved Minimum Path Sum (0064)
- [ ] Can handle obstacles (0063)
- [ ] Understand rolling array optimization

### Advanced Patterns:
- [ ] Solved Regular Expression Matching (0010)
- [ ] Solved Combination Sum (0039)
- [ ] Solved Word Break (0139)
- [ ] Understand reverse DP (0174)

## 🔍 Missing Problems to Add

### High Priority:
1. **0070. Climbing Stairs** - Fibonacci classic
2. **0198. House Robber** - 1D DP with constraints
3. **0322. Coin Change** - Unbounded knapsack
4. **0300. Longest Increasing Subsequence** - Classic LIS
5. **0114. Flatten Binary Tree to Linked List** - Tree DP

### Medium Priority:
1. **0416. Partition Equal Subset Sum** - 0/1 knapsack
2. **0494. Target Sum** - DP with sum constraints
3. **0646. Maximum Length of Pair Chain** - Interval DP
4. **0718. Maximum Length of Repeated Subarray** - 2D array DP

### Nice to Have:
1. **0877. Stone Game** - Game theory DP
2. **0983. Minimum Cost For Tickets** - DP with days
3. **1049. Last Stone Weight II** - Partition DP
4. **1143. Longest Common Subsequence** - Classic LCS

## 📚 Resources

### In this Repository:
- `tutorials/categories/03_dynamic_programming.md` - DP tutorial
- `DP_SOLUTIONS_INDEX.md` - Complete index (this file)
- Individual problem files (####_problem_name.go)

### External Resources:
- LeetCode Dynamic Programming Card
- "Introduction to Algorithms" DP chapter
- MIT OpenCourseWare: Introduction to Algorithms
- Dynamic Programming for Coding Interviews (book)

## 🎯 Next Actions

### Immediate (This Week):
1. Review 5 core DP problems (0072, 0139, 0039, 0005, 0062)
2. Practice space optimization for each
3. Solve 2 new DP problems from missing list

### Short-term (Next Month):
1. Complete all Medium DP problems
2. Master pattern recognition
3. Implement both memoization and tabulation

### Long-term (3 Months):
1. Solve all Hard DP problems
2. Learn advanced DP patterns (bitmask, tree, probability)
3. Participate in coding contests

---

*Use this quick reference for daily practice and interview preparation.*  
*Last updated: January 28, 2026*