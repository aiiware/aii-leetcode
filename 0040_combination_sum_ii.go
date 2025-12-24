// 0040 - Combination Sum II
// https://leetcode.com/problems/combination-sum-ii/
// Medium - Array, Backtracking

package leetcode

import "sort"

// CombinationSum2 finds all unique combinations where candidate numbers sum to target.
// Each number in candidates may only be used once in the combination.
// Time Complexity: O(2^n) where n is length of candidates
// Space Complexity: O(n) for recursion depth
func CombinationSum2(candidates []int, target int) [][]int {
	// Sort candidates to handle duplicates
	sort.Ints(candidates)
	
	result := [][]int{}
	var current []int
	
	backtrackCombinationSum2(candidates, target, 0, current, &result)
	return result
}

// backtrackCombinationSum2 is the recursive backtracking function
func backtrackCombinationSum2(candidates []int, target, start int, current []int, result *[][]int) {
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
		// Skip duplicates to avoid duplicate combinations
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		
		// Prune: if current candidate is greater than remaining target, skip
		if candidates[i] > target {
			break
		}
		
		// Include candidates[i] in the combination
		current = append(current, candidates[i])
		backtrackCombinationSum2(candidates, target-candidates[i], i+1, current, result)
		// Backtrack
		current = current[:len(current)-1]
	}
}

// CombinationSum2DP uses dynamic programming (0/1 knapsack)
func CombinationSum2DP(candidates []int, target int) [][]int {
	// Sort candidates
	sort.Ints(candidates)
	
	// dp[i] stores all combinations that sum to i
	dp := make([][][]int, target+1)
	dp[0] = [][]int{{}} // Base case
	
	// Track last used index for each combination to avoid using same element multiple times
	lastIndex := make([][]int, target+1)
	lastIndex[0] = []int{-1} // Base case
	
	for i, num := range candidates {
		// Process from high to low to avoid using same element multiple times
		for t := target; t >= num; t-- {
			if len(dp[t-num]) > 0 {
				for j, combination := range dp[t-num] {
					// Check if we can use this candidate (not used before in this path)
					if lastIndex[t-num][j] < i {
						// Create new combination
						newCombination := make([]int, len(combination)+1)
						copy(newCombination, combination)
						newCombination[len(combination)] = num
						
						dp[t] = append(dp[t], newCombination)
						lastIndex[t] = append(lastIndex[t], i)
					}
				}
			}
		}
	}
	
	// Remove duplicates from result
	uniqueMap := make(map[string]bool)
	var uniqueResult [][]int
	
	for _, combination := range dp[target] {
		// Sort combination to create consistent key
		sortedComb := make([]int, len(combination))
		copy(sortedComb, combination)
		sort.Ints(sortedComb)
		
		// Create key
		key := ""
		for _, num := range sortedComb {
			key += string(rune(num)) + ","
		}
		
		if !uniqueMap[key] {
			uniqueMap[key] = true
			uniqueResult = append(uniqueResult, combination)
		}
	}
	
	// Return empty slice instead of nil
	if uniqueResult == nil {
		return [][]int{}
	}
	return uniqueResult
}

// CombinationSum2Iterative uses iterative approach with stack
func CombinationSum2Iterative(candidates []int, target int) [][]int {
	// Sort candidates
	sort.Ints(candidates)
	
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
			// Skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			
			// Prune
			if candidates[i] > remaining {
				break
			}
			
			newCurrent := make([]int, len(current)+1)
			copy(newCurrent, current)
			newCurrent[len(current)] = candidates[i]
			
			stack = append(stack, []interface{}{
				remaining - candidates[i],
				i + 1, // i+1 because we cannot reuse same element
				newCurrent,
			})
		}
	}
	
	return result
}