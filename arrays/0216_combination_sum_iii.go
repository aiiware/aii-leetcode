package arrays

// CombinationSum3 solves LeetCode problem 0216: Combination Sum III
// Difficulty: Medium
// Tags: Array, Backtracking
//
// Find all valid combinations of k numbers that sum up to n such that:
// - Only numbers 1 through 9 are used
// - Each number is used at most once
//
// Time complexity: O(C(9,k)), Space complexity: O(k)
func CombinationSum3(k int, n int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start, remaining int)
	backtrack = func(start, remaining int) {
		if len(current) == k {
			if remaining == 0 {
				// Make a copy of current combination
				combination := make([]int, k)
				copy(combination, current)
				result = append(result, combination)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if i > remaining {
				break
			}
			current = append(current, i)
			backtrack(i+1, remaining-i)
			current = current[:len(current)-1]
		}
	}

	backtrack(1, n)

	return result
}
