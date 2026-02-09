package arrays

// Combine solves LeetCode problem 0077: Combinations
// Difficulty: Medium
// Tags: Backtracking, Depth-First Search, Recursion
//
// Given two integers n and k, return all possible combinations of k numbers
// chosen from the range [1, n].
//
// You may return the answer in any order.
//
// Example 1:
// Input: n = 4, k = 2
// Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Explanation: There are 4 choose 2 = 6 total combinations.
//
// Example 2:
// Input: n = 1, k = 1
// Output: [[1]]
// Explanation: There is only 1 choose 1 = 1 combination.
//
// Constraints:
// 1 <= n <= 20
// 1 <= k <= n
//
// Time complexity: O(C(n,k) * k) where C(n,k) is binomial coefficient
// Space complexity: O(C(n,k) * k) for storing results
func Combine(n int, k int) [][]int {
	if n < 1 || k < 1 || k > n {
		return [][]int{}
	}

	result := make([][]int, 0)
	current := make([]int, 0, k)

	var backtrack func(start int)
	backtrack = func(start int) {
		// If we have k numbers, add to result
		if len(current) == k {
			// Make a copy of current combination
			combination := make([]int, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// Try all numbers from start to n
		for i := start; i <= n; i++ {
			// Add current number to combination
			current = append(current, i)
			// Recursively build the rest of the combination
			backtrack(i + 1)
			// Backtrack: remove last number
			current = current[:len(current)-1]
		}
	}

	backtrack(1)
	return result
}

// CombineOptimized is an optimized version with pruning
// We can prune the search space by limiting the maximum value we can choose
// based on how many numbers we still need to pick
func CombineOptimized(n int, k int) [][]int {
	if n < 1 || k < 1 || k > n {
		return [][]int{}
	}

	result := make([][]int, 0)
	current := make([]int, 0, k)

	var backtrack func(start int)
	backtrack = func(start int) {
		// If we have k numbers, add to result
		if len(current) == k {
			combination := make([]int, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// Pruning: we need (k - len(current)) more numbers
		// The maximum value we can choose is n - (k - len(current)) + 1
		// This ensures we have enough numbers left to complete the combination
		maxVal := n - (k - len(current)) + 1

		for i := start; i <= maxVal; i++ {
			current = append(current, i)
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(1)
	return result
}

// CombineIterative uses an iterative approach with bit manipulation
// This approach generates all subsets of size k from n elements
func CombineIterative(n int, k int) [][]int {
	if n < 1 || k < 1 || k > n {
		return [][]int{}
	}

	result := make([][]int, 0)

	// Generate all bitmasks of length n with exactly k bits set
	// We iterate through all bitmasks from 0 to 2^n - 1
	for mask := 0; mask < (1 << uint(n)); mask++ {
		// Count number of set bits
		count := 0
		temp := mask
		for temp > 0 {
			count += temp & 1
			temp >>= 1
		}

		// If mask has exactly k bits set, it's a valid combination
		if count == k {
			combination := make([]int, 0, k)
			// Extract numbers from the bitmask
			for i := 0; i < n; i++ {
				if mask&(1<<uint(i)) != 0 {
					combination = append(combination, i+1) // +1 because numbers start from 1
				}
			}
			result = append(result, combination)
		}
	}

	return result
}

// CombineDP uses dynamic programming to build combinations
// This approach builds combinations from smaller to larger
func CombineDP(n int, k int) [][]int {
	if n < 1 || k < 1 || k > n {
		return [][]int{}
	}

	// dp[i][j] stores all combinations of size j from numbers 1..i
	dp := make([][][][]int, n+1)
	for i := range dp {
		dp[i] = make([][][]int, k+1)
	}

	// Base case: combinations of size 0
	for i := 0; i <= n; i++ {
		dp[i][0] = [][]int{{}}
	}

	// Build dp table
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			// Combinations that don't include i
			combinationsWithoutI := dp[i-1][j]

			// Combinations that include i
			combinationsWithI := make([][]int, 0)
			for _, comb := range dp[i-1][j-1] {
				// Create new combination by adding i
				newComb := make([]int, len(comb)+1)
				copy(newComb, comb)
				newComb[len(comb)] = i
				combinationsWithI = append(combinationsWithI, newComb)
			}

			// Combine both
			dp[i][j] = append(combinationsWithoutI, combinationsWithI...)
		}
	}

	return dp[n][k]
}

// CombineMath uses mathematical formula to generate combinations
// This is an alternative approach using combinatorial number system
func CombineMath(n int, k int) [][]int {
	if n < 1 || k < 1 || k > n {
		return [][]int{}
	}

	result := make([][]int, 0)

	// Initialize first combination: [1, 2, ..., k]
	current := make([]int, k)
	for i := 0; i < k; i++ {
		current[i] = i + 1
	}

	for {
		// Add current combination to result
		combination := make([]int, k)
		copy(combination, current)
		result = append(result, combination)

		// Find the rightmost element that can be incremented
		i := k - 1
		for i >= 0 && current[i] == n-k+i+1 {
			i--
		}

		// If no element can be incremented, we're done
		if i < 0 {
			break
		}

		// Increment this element
		current[i]++

		// Reset all elements to the right
		for j := i + 1; j < k; j++ {
			current[j] = current[j-1] + 1
		}
	}

	return result
}
