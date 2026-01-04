package leetcode

// GetPermutation solves LeetCode problem 0060: Permutation Sequence
// Difficulty: Hard
// Tags: Math, Backtracking
//
// The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
// "123", "132", "213", "231", "312", "321"
//
// Given n and k, return the kth permutation sequence.
//
// Example 1:
// Input: n = 3, k = 3
// Output: "213"
//
// Example 2:
// Input: n = 4, k = 9
// Output: "2314"
//
// Example 3:
// Input: n = 3, k = 1
// Output: "123"
//
// Constraints:
// 1 <= n <= 9
// 1 <= k <= n!
//
// Time complexity: O(n^2), Space complexity: O(n)
func GetPermutation(n int, k int) string {
	if n <= 0 || k <= 0 {
		return ""
	}

	// Calculate factorials up to n
	factorials := make([]int, n+1)
	factorials[0] = 1
	for i := 1; i <= n; i++ {
		factorials[i] = factorials[i-1] * i
	}

	// Validate k
	if k > factorials[n] {
		return ""
	}

	// Create list of available numbers
	numbers := make([]byte, n)
	for i := 0; i < n; i++ {
		numbers[i] = byte('1' + i)
	}

	// Convert k to 0-based index
	k--

	result := make([]byte, 0, n)

	for i := n; i > 0; i-- {
		// Calculate index of current digit
		index := k / factorials[i-1]
		k %= factorials[i-1]

		// Add selected digit to result
		result = append(result, numbers[index])

		// Remove selected digit from available numbers
		numbers = append(numbers[:index], numbers[index+1:]...)
	}

	return string(result)
}