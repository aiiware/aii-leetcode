# 10. Regular Expression Matching - Solution Explanation

## Problem Statement
Given an input string `s` and a pattern `p`, implement regular expression matching with support for '.' and '*' where:
- '.' Matches any single character
- '*' Matches zero or more of the preceding element

The matching should cover the **entire** input string (not partial).

## Difficulty: Hard

## Key Insights
1. **Dynamic Programming Approach**: This is a classic DP problem similar to wildcard matching but with different rules for '*'
2. **State Definition**: `dp[i][j]` = whether `s[0..i-1]` matches `p[0..j-1]`
3. **Special Cases**:
   - '*' can match zero or more of the preceding character
   - '.' matches any single character
   - Need to handle cases where '*' means zero occurrences

## Solution Approaches

### Approach 1: Dynamic Programming (Bottom-Up)
**Time Complexity**: O(m × n) where m = len(s), n = len(p)
**Space Complexity**: O(m × n) can be optimized to O(n)

```go
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    
    // dp[i][j] = does s[0..i-1] match p[0..j-1]
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    
    // Empty string matches empty pattern
    dp[0][0] = true
    
    // Handle patterns like a*, a*b*, a*b*c* that can match empty string
    for j := 1; j <= n; j++ {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-2]  // '*' can match zero of preceding char
        }
    }
    
    // Fill DP table
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '.' || p[j-1] == s[i-1] {
                // Current characters match
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                // '*' case
                // Option 1: Match zero of preceding character
                dp[i][j] = dp[i][j-2]
                
                // Option 2: Match one or more of preceding character
                // Check if preceding character matches
                if p[j-2] == '.' || p[j-2] == s[i-1] {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            }
            // Otherwise dp[i][j] remains false
        }
    }
    
    return dp[m][n]
}
```

### Approach 2: Recursion with Memoization (Top-Down)
**Time Complexity**: O(m × n)
**Space Complexity**: O(m × n)

```go
func isMatchMemo(s string, p string) bool {
    memo := make(map[[2]int]bool)
    
    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        // Check memo
        if val, ok := memo[[2]int{i, j}]; ok {
            return val
        }
        
        // Base cases
        if j == len(p) {
            return i == len(s)
        }
        
        // Check if current characters match
        firstMatch := i < len(s) && (p[j] == '.' || p[j] == s[i])
        
        var result bool
        if j+1 < len(p) && p[j+1] == '*' {
            // Case 1: '*' matches zero of preceding character
            // Case 2: '*' matches one or more of preceding character
            result = dfs(i, j+2) || (firstMatch && dfs(i+1, j))
        } else {
            // Simple match
            result = firstMatch && dfs(i+1, j+1)
        }
        
        memo[[2]int{i, j}] = result
        return result
    }
    
    return dfs(0, 0)
}
```

## Step-by-Step Walkthrough

### Example:
```
s = "aab"
p = "c*a*b"
```

**DP Table Initialization**:
```
   "" c * a * b
"" T  F T F T F
a  F
a  F
b  F
```

**Step-by-step filling**:

1. `dp[0][0] = true` (empty matches empty)
2. Handle `*` in pattern for empty string:
   - `dp[0][2] = dp[0][0] = true` (c* matches zero c's)
   - `dp[0][4] = dp[0][2] = true` (a* matches zero a's after c*)

3. Fill row i=1 (s="a"):
   - j=1: p="c" doesn't match s="a" → false
   - j=2: p="c*" → check zero match (dp[1][0]=false) OR if 'c' matches 'a' (false) → false
   - j=3: p="a" matches 'a' → dp[0][2]=true → true
   - j=4: p="a*" → zero match (dp[1][2]=false) OR 'a' matches 'a' and dp[0][4]=true → true
   - j=5: p="b" doesn't match 'a' → false

4. Continue for all cells...

**Final result**: `dp[3][5] = true`

## Common Patterns and Tricks

### Pattern Analysis:
1. `.*` - Matches any sequence of characters
2. `a*` - Matches zero or more 'a's
3. `a.b` - Matches 'a', any character, 'b'
4. `a*.*b*` - Complex pattern requiring careful DP

### Optimization Tips:
1. **Early termination**: If pattern has too many characters that must match (excluding '*'), and string is too short
2. **Space optimization**: Use only two rows for DP instead of full matrix
3. **Pattern preprocessing**: Combine consecutive '*' (a** = a*)

## Edge Cases
1. Empty string with pattern like `a*`, `.*`, `a*b*`
2. String with repeating characters and pattern like `a*a*a*`
3. Pattern starting with '*' (invalid in problem but good to handle)
4. Multiple '.' and '*' combinations

## Related Problems
- 44. Wildcard Matching (different '*' semantics)
- 72. Edit Distance (similar DP structure)
- 97. Interleaving String

## Practice Exercises
1. Modify to support '+' (one or more) operator
2. Implement NFA (Non-deterministic Finite Automaton) approach
3. Add support for character classes like [a-z]
4. Benchmark DP vs recursion with memoization