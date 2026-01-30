package dp

// Problem 0097: Interleaving String
//
// Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
//
// An interleaving of two strings s and t is a configuration where s and t are divided into n and m 
// substrings respectively, such that:
// - s = s1 + s2 + ... + sn
// - t = t1 + t2 + ... + tm
// - |n - m| <= 1
// - The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
//
// Note: a + b is the concatenation of strings a and b.
//
// Example 1:
// Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// Output: true
// Explanation: One way to obtain s3 is:
// Split s1 = "aa" + "bc" + "c", and s2 = "dbbc" + "a".
// Interleaving: "aa" + "dbbc" + "bc" + "a" + "c" = "aadbbcbcac".
//
// Example 2:
// Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// Output: false
//
// Example 3:
// Input: s1 = "", s2 = "", s3 = ""
// Output: true
//
// Constraints:
// - 0 <= s1.length, s2.length <= 100
// - 0 <= s3.length <= 200
// - s1, s2, and s3 consist of lowercase English letters.

// isInterleave is the main solution function using dynamic programming.
// Time complexity: O(m*n), Space complexity: O(m*n)
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	
	// Quick check: lengths must sum up
	if m+n != len(s3) {
		return false
	}
	
	// dp[i][j] = whether s3[0:i+j] is interleaving of s1[0:i] and s2[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	
	// Base case: empty strings
	dp[0][0] = true
	
	// First column: only using s1
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	
	// First row: only using s2
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	
	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Check if current character of s3 matches s1 or s2
			matchS1 := dp[i-1][j] && s1[i-1] == s3[i+j-1]
			matchS2 := dp[i][j-1] && s2[j-1] == s3[i+j-1]
			dp[i][j] = matchS1 || matchS2
		}
	}
	
	return dp[m][n]
}

// isInterleaveOptimized uses optimized space O(n).
func isInterleaveOptimized(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	
	if m+n != len(s3) {
		return false
	}
	
	// Use 1D DP array
	dp := make([]bool, n+1)
	dp[0] = true
	
	// Initialize first row
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}
	
	// Fill DP table row by row
	for i := 1; i <= m; i++ {
		// Update first column for current row
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		
		for j := 1; j <= n; j++ {
			matchS1 := dp[j] && s1[i-1] == s3[i+j-1]
			matchS2 := dp[j-1] && s2[j-1] == s3[i+j-1]
			dp[j] = matchS1 || matchS2
		}
	}
	
	return dp[n]
}

// isInterleaveDFS uses DFS with memoization.
func isInterleaveDFS(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	
	if m+n != len(s3) {
		return false
	}
	
	memo := make(map[[2]int]bool)
	return dfsInterleave(s1, s2, s3, 0, 0, 0, memo)
}

func dfsInterleave(s1, s2, s3 string, i, j, k int, memo map[[2]int]bool) bool {
	// If we've reached the end of s3
	if k == len(s3) {
		return i == len(s1) && j == len(s2)
	}
	
	// Check memo
	key := [2]int{i, j}
	if val, exists := memo[key]; exists {
		return val
	}
	
	result := false
	
	// Try taking from s1 if possible
	if i < len(s1) && s1[i] == s3[k] {
		result = result || dfsInterleave(s1, s2, s3, i+1, j, k+1, memo)
	}
	
	// Try taking from s2 if possible
	if j < len(s2) && s2[j] == s3[k] {
		result = result || dfsInterleave(s1, s2, s3, i, j+1, k+1, memo)
	}
	
	memo[key] = result
	return result
}

// isInterleaveBFS uses BFS approach.
func isInterleaveBFS(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	
	if m+n != len(s3) {
		return false
	}
	
	// BFS queue stores (i, j) pairs
	queue := [][2]int{{0, 0}}
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		i, j := current[0], current[1]
		k := i + j
		
		// If we've reached the end
		if k == len(s3) {
			return true
		}
		
		// Try moving right (take from s1)
		if i < m && s1[i] == s3[k] {
			next := [2]int{i + 1, j}
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
		
		// Try moving down (take from s2)
		if j < n && s2[j] == s3[k] {
			next := [2]int{i, j + 1}
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	
	return false
}

// isInterleaveRecursive uses pure recursion (exponential time).
func isInterleaveRecursive(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	
	if len(s3) == 0 {
		return true
	}
	
	result := false
	
	// Try taking first character from s1
	if len(s1) > 0 && s1[0] == s3[0] {
		result = result || isInterleaveRecursive(s1[1:], s2, s3[1:])
	}
	
	// Try taking first character from s2
	if len(s2) > 0 && s2[0] == s3[0] {
		result = result || isInterleaveRecursive(s1, s2[1:], s3[1:])
	}
	
	return result
}

// isInterleaveDP2 uses another DP formulation.
func isInterleaveDP2(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	
	if m+n != len(s3) {
		return false
	}
	
	// dp[i][j] = whether s3[i+j-1] can be formed by interleaving s1[0:i] and s2[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	
	// Fill DP table
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) ||
					(dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	
	return dp[m][n]
}

// IsInterleave is the public interface function.
// It uses the optimized DP solution by default.
func IsInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveOptimized(s1, s2, s3)
}