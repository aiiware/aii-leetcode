package dp

import "leetcode/utils"

// 0161. One Edit Distance
// https://leetcode.com/problems/one-edit-distance

// isOneEditDistance is the main solution function
func isOneEditDistance(s string, t string) bool {
	// Solution 1: Two pointers approach
	return isOneEditDistanceTwoPointers(s, t)
}

// ===== Solution 1: Two pointers approach =====
// Time complexity: O(n) where n is the length of the longer string
// Space complexity: O(1)

func isOneEditDistanceTwoPointers(s string, t string) bool {
	m, n := len(s), len(t)

	// If lengths differ by more than 1, cannot be one edit distance
	if utils.Abs(m-n) > 1 {
		return false
	}

	// Find first difference
	i, j := 0, 0
	for i < m && j < n && s[i] == t[j] {
		i++
		j++
	}

	// If we reached end of both strings, they are identical (0 edit distance)
	if i == m && j == n {
		return false
	}

	// Handle different cases based on string lengths
	if m == n {
		// Same length: check if remaining parts are equal (replace operation)
		return s[i+1:] == t[j+1:]
	} else if m < n {
		// s is shorter: check if we can insert one character in s (or delete from t)
		return s[i:] == t[j+1:]
	} else {
		// s is longer: check if we can delete one character from s (or insert into t)
		return s[i+1:] == t[j:]
	}
}

// ===== Solution 2: Dynamic programming (edit distance = 1) =====
// Time complexity: O(m * n)
// Space complexity: O(min(m, n))
// More general but less efficient for this specific problem

func isOneEditDistanceDP(s string, t string) bool {
	m, n := len(s), len(t)

	// If lengths differ by more than 1, cannot be one edit distance
	if utils.Abs(m-n) > 1 {
		return false
	}

	// Ensure s is the shorter string for space optimization
	if m > n {
		s, t = t, s
		m, n = n, m
	}

	// DP array for current row
	dp := make([]int, n+1)

	// Initialize first row (empty string to t[0:j])
	for j := 0; j <= n; j++ {
		dp[j] = j
	}

	// Fill DP table
	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = i

		for j := 1; j <= n; j++ {
			temp := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = prev
			} else {
				// Minimum of insert, delete, replace
				dp[j] = utils.Min(prev, utils.Min(dp[j], dp[j-1])) + 1
			}
			prev = temp
		}
	}

	return dp[n] == 1
}

// ===== Solution 3: Early exit approach =====
// Time complexity: O(n) in average case
// Space complexity: O(1)

func isOneEditDistanceEarlyExit(s string, t string) bool {
	m, n := len(s), len(t)

	// Quick checks
	if utils.Abs(m-n) > 1 {
		return false
	}

	// If strings are identical, not one edit distance
	if s == t {
		return false
	}

	// If one string is empty, the other must have exactly one character
	if m == 0 {
		return n == 1
	}
	if n == 0 {
		return m == 1
	}

	// Compare character by character
	edits := 0
	i, j := 0, 0

	for i < m && j < n {
		if s[i] != t[j] {
			edits++
			if edits > 1 {
				return false
			}

			// Handle length difference
			if m > n {
				i++ // delete from s (or insert into t)
			} else if m < n {
				j++ // insert into s (or delete from t)
			} else {
				// same length, move both pointers (replace)
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}

	// Check remaining characters
	if i < m {
		edits += m - i
	}
	if j < n {
		edits += n - j
	}

	return edits == 1
}

// ===== Solution 4: Simple length-based approach =====
// Time complexity: O(n)
// Space complexity: O(1)

func isOneEditDistanceSimple(s string, t string) bool {
	m, n := len(s), len(t)

	// Ensure s is not longer than t
	if m > n {
		return isOneEditDistanceSimple(t, s)
	}

	// Length check
	if n-m > 1 {
		return false
	}

	for i := 0; i < m; i++ {
		if s[i] != t[i] {
			if m == n {
				// Replace: check if remaining parts are equal
				return s[i+1:] == t[i+1:]
			} else {
				// Insert/Delete: check if s[i:] == t[i+1:]
				return s[i:] == t[i+1:]
			}
		}
	}

	// All characters matched in the shorter string
	// If lengths are equal, strings are identical (0 edits)
	// If t is longer by 1, we need to insert that character
	return m != n
}