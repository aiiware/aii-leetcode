package leetcode

// Problem 0096: Unique Binary Search Trees
//
// Given an integer n, return the number of structurally unique BST's 
// (binary search trees) which has exactly n nodes of unique values from 1 to n.
//
// Example 1:
// Input: n = 3
// Output: 5
//
// Example 2:
// Input: n = 1
// Output: 1
//
// Constraints:
// - 1 <= n <= 19

// numTrees is the main solution function using dynamic programming (Catalan numbers).
// Time complexity: O(n^2), Space complexity: O(n)
func numTrees(n int) int {
	if n <= 0 {
		return 0
	}

	// dp[i] = number of unique BSTs with i nodes
	dp := make([]int, n+1)
	dp[0] = 1 // Empty tree
	dp[1] = 1 // Single node tree

	// Build up to n
	for i := 2; i <= n; i++ {
		// For i nodes, root can be any node from 1 to i
		// If root is j, left subtree has j-1 nodes, right subtree has i-j nodes
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

// numTreesCatalan uses direct Catalan number formula.
// Time complexity: O(n), Space complexity: O(1)
func numTreesCatalan(n int) int {
	if n <= 0 {
		return 0
	}

	// Catalan number formula: C(n) = (2n)! / ((n+1)! * n!)
	// Using iterative calculation to avoid overflow
	result := 1
	for i := 0; i < n; i++ {
		result = result * (2*n - i) / (i + 1)
	}
	return result / (n + 1)
}

// numTreesRecursive uses recursion with memoization.
func numTreesRecursive(n int) int {
	if n <= 0 {
		return 0
	}

	memo := make(map[int]int)
	return numTreesHelper(n, memo)
}

func numTreesHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return 1
	}

	if val, exists := memo[n]; exists {
		return val
	}

	total := 0
	// For each root position
	for root := 1; root <= n; root++ {
		left := numTreesHelper(root-1, memo)
		right := numTreesHelper(n-root, memo)
		total += left * right
	}

	memo[n] = total
	return total
}

// numTreesIterative uses iterative DP with optimized space.
func numTreesIterative(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

// numTreesMath uses mathematical formula with factorial.
// Note: May overflow for large n.
func numTreesMath(n int) int {
	if n <= 0 {
		return 0
	}

	// C(n) = (2n)! / ((n+1)! * n!)
	// We compute factorials iteratively
	n64 := int64(n)
	
	// Compute (2n)!
	numerator := int64(1)
	for i := int64(1); i <= 2*n64; i++ {
		numerator *= i
	}

	// Compute (n+1)!
	denom1 := int64(1)
	for i := int64(1); i <= n64+1; i++ {
		denom1 *= i
	}

	// Compute n!
	denom2 := int64(1)
	for i := int64(1); i <= n64; i++ {
		denom2 *= i
	}

	result := numerator / (denom1 * denom2)
	return int(result)
}

// numTreesOptimized is an optimized version using Catalan numbers.
func numTreesOptimized(n int) int {
	if n <= 0 {
		return 0
	}

	// Use 64-bit to avoid overflow for n up to 19
	var result int64 = 1
	for i := int64(0); i < int64(n); i++ {
		result = result * (2*int64(n) - i) / (i + 1)
	}
	result = result / (int64(n) + 1)
	return int(result)
}

// numTreesDP2 uses another DP formulation.
func numTreesDP2(n int) int {
	if n <= 0 {
		return 0
	}

	// dp[i][j] = number of BSTs with values i..j
	// But we only need dp[1][n], and it depends on length
	// So we can use 1D dp where dp[k] = number of BSTs with k consecutive values
	dp := make([]int, n+1)
	dp[0] = 1

	for length := 1; length <= n; length++ {
		for root := 1; root <= length; root++ {
			dp[length] += dp[root-1] * dp[length-root]
		}
	}

	return dp[n]
}

// NumTrees is the public interface function.
// It uses the optimized Catalan number solution by default.
func NumTrees(n int) int {
	return numTreesOptimized(n)
}