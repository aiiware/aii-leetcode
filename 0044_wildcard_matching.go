// 0044 - Wildcard Matching
// https://leetcode.com/problems/wildcard-matching/
// Hard - String, Dynamic Programming, Greedy, Recursion

package leetcode

// IsMatchWildcard determines if pattern matches string with wildcards.
// Pattern supports '?' (matches any single character) and '*' (matches any sequence).
// Time Complexity: O(m * n) where m and n are lengths of s and p
// Space Complexity: O(m * n) for DP, O(1) for greedy
func IsMatchWildcard(s string, p string) bool {
    m, n := len(s), len(p)
    
    // DP approach
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    
    // Empty pattern matches empty string
    dp[0][0] = true
    
    // Handle patterns starting with '*'
    for j := 1; j <= n; j++ {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-1]
        }
    }
    
    // Fill DP table
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                // '*' can match empty sequence or any sequence
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            } else if p[j-1] == '?' || s[i-1] == p[j-1] {
                // '?' matches any single character, or characters match
                dp[i][j] = dp[i-1][j-1]
            }
            // else dp[i][j] remains false
        }
    }
    
    return dp[m][n]
}

// IsMatchWildcardGreedy uses greedy approach
func IsMatchWildcardGreedy(s string, p string) bool {
    i, j := 0, 0
    starIdx, matchIdx := -1, -1
    
    for i < len(s) {
        if j < len(p) && (p[j] == '?' || p[j] == s[i]) {
            // Characters match or '?' matches any
            i++
            j++
        } else if j < len(p) && p[j] == '*' {
            // Found '*', record positions
            starIdx = j
            matchIdx = i
            j++
        } else if starIdx != -1 {
            // Use '*' to match more characters
            j = starIdx + 1
            matchIdx++
            i = matchIdx
        } else {
            // No match and no '*'
            return false
        }
    }
    
    // Skip remaining '*' in pattern
    for j < len(p) && p[j] == '*' {
        j++
    }
    
    return j == len(p)
}

// IsMatchWildcardRecursive uses recursive approach with memoization
func IsMatchWildcardRecursive(s string, p string) bool {
    memo := make(map[[2]int]bool)
    return isMatchRecursive(s, p, 0, 0, memo)
}

func isMatchRecursive(s, p string, i, j int, memo map[[2]int]bool) bool {
    key := [2]int{i, j}
    if val, ok := memo[key]; ok {
        return val
    }
    
    // Base cases
    if j == len(p) {
        result := i == len(s)
        memo[key] = result
        return result
    }
    
    if i == len(s) {
        // Check if remaining pattern is all '*'
        for k := j; k < len(p); k++ {
            if p[k] != '*' {
                memo[key] = false
                return false
            }
        }
        memo[key] = true
        return true
    }
    
    var result bool
    if p[j] == '*' {
        // '*' can match empty or one or more characters
        result = isMatchRecursive(s, p, i, j+1, memo) || // match empty
                 isMatchRecursive(s, p, i+1, j, memo)   // match one character and continue
    } else if p[j] == '?' || s[i] == p[j] {
        // Characters match
        result = isMatchRecursive(s, p, i+1, j+1, memo)
    } else {
        result = false
    }
    
    memo[key] = result
    return result
}