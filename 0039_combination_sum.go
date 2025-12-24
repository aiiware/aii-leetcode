// 0039 - Combination Sum
// https://leetcode.com/problems/combination-sum/
// Medium - Array, Backtracking

package leetcode

// CombinationSum finds all unique combinations where candidate numbers sum to target.
// Time Complexity: O(N^(T/M + 1)) where N is candidates length, T is target, M is min candidate
// Space Complexity: O(T/M) for recursion depth
func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var current []int
	
	backtrackCombinationSum(candidates, target, 0, current, &result)
	return result
}

// backtrackCombinationSum is the recursive backtracking function
func backtrackCombinationSum(candidates []int, target, start int, current []int, result *[][]int) {
	if target < 0 {
		return
	}
	
	if target == 0 {
		// Make a copy of current combination
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}
	
	for i := start; i < len(candidates); i++ {
		// Include candidates[i] in the combination
		current = append(current, candidates[i])
		backtrackCombinationSum(candidates, target-candidates[i], i, current, result)
		// Backtrack
		current = current[:len(current)-1]
	}
}

// CombinationSumDP uses dynamic programming (unbounded knapsack)
func CombinationSumDP(candidates []int, target int) [][]int {
	// dp[i] stores all combinations that sum to i
	dp := make([][][]int, target+1)
	dp[0] = [][]int{{}} // Base case: one way to make sum 0 (empty combination)
	
	// Sort candidates to ensure consistent ordering
	sortInts(candidates)
	
	for _, num := range candidates {
		for t := num; t <= target; t++ {
			for _, combination := range dp[t-num] {
				// Create new combination by adding num to existing combination
				newCombination := make([]int, len(combination)+1)
				copy(newCombination, combination)
				newCombination[len(combination)] = num
				dp[t] = append(dp[t], newCombination)
			}
		}
	}
	
	// Return empty slice instead of nil
	if dp[target] == nil {
		return [][]int{}
	}
	return dp[target]
}

// CombinationSumIterative uses iterative approach with stack
func CombinationSumIterative(candidates []int, target int) [][]int {
	// Special case: target is 0
	if target == 0 {
		return [][]int{{}}
	}
	
	result := [][]int{}
	
	// Stack stores (remaining target, start index, current combination)
	stack := [][]interface{}{
		{target, 0, []int{}},
	}
	
	for len(stack) > 0 {
		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		remaining := item[0].(int)
		start := item[1].(int)
		current := item[2].([]int)
		
		if remaining == 0 {
			// Found a valid combination
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			continue
		}
		
		if remaining < 0 {
			continue
		}
		
		// Try each candidate starting from 'start'
		for i := start; i < len(candidates); i++ {
			newCurrent := make([]int, len(current)+1)
			copy(newCurrent, current)
			newCurrent[len(current)] = candidates[i]
			
			stack = append(stack, []interface{}{
				remaining - candidates[i],
				i, // Note: i not i+1 because we can reuse same element
				newCurrent,
			})
		}
	}
	
	return result
}

// sortInts sorts integers in ascending order
func sortInts(nums []int) {
	// Simple bubble sort for small arrays
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}