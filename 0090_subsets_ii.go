package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

// Problem 0090: Subsets II
//
// Given an integer array nums that may contain duplicates, return all possible 
// subsets (the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Example 1:
// Input: nums = [1,2,2]
// Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
//
// Example 2:
// Input: nums = [0]
// Output: [[],[0]]
//
// Constraints:
// - 1 <= nums.length <= 10
// - -10 <= nums[i] <= 10

// subsetsWithDup is the main solution function using backtracking with sorting.
// Time complexity: O(n * 2^n), Space complexity: O(n * 2^n)
func subsetsWithDup(nums []int) [][]int {
	// Sort to handle duplicates
	sort.Ints(nums)
	
	result := [][]int{}
	current := []int{}
	
	var backtrack func(int)
	backtrack = func(start int) {
		// Add current subset to result
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)
		
		// Explore further
		for i := start; i < len(nums); i++ {
			// Skip duplicates
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			
			// Include nums[i]
			current = append(current, nums[i])
			backtrack(i + 1)
			
			// Backtrack
			current = current[:len(current)-1]
		}
	}
	
	backtrack(0)
	return result
}

// subsetsWithDupIterative uses an iterative approach.
func subsetsWithDupIterative(nums []int) [][]int {
	// Sort to handle duplicates
	sort.Ints(nums)
	
	result := [][]int{{}} // Start with empty subset
	
	size := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		// If current element is same as previous, only add to subsets created in previous iteration
		if i > 0 && nums[i] == nums[i-1] {
			start = size
		} else {
			start = 0
		}
		
		size = len(result)
		for j := start; j < size; j++ {
			// Create new subset by adding nums[i] to existing subset
			newSubset := make([]int, len(result[j]))
			copy(newSubset, result[j])
			newSubset = append(newSubset, nums[i])
			result = append(result, newSubset)
		}
	}
	
	return result
}

// subsetsWithDupBitmask uses bitmask approach (less efficient with duplicates).
func subsetsWithDupBitmask(nums []int) [][]int {
	// Sort to make duplicates adjacent
	sort.Ints(nums)
	
	n := len(nums)
	total := 1 << n
	result := [][]int{}
	seen := make(map[string]bool)
	
	for mask := 0; mask < total; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		
		// Create a string key to check for duplicates
		key := fmtSubset(subset)
		if !seen[key] {
			seen[key] = true
			result = append(result, subset)
		}
	}
	
	return result
}

// fmtSubset creates a string representation of a subset for duplicate checking
func fmtSubset(subset []int) string {
	// Simple string representation
	var sb strings.Builder
	for _, num := range subset {
		sb.WriteString(fmt.Sprintf("%d,", num))
	}
	return sb.String()
}

// subsetsWithDupDFS uses DFS approach.
func subsetsWithDupDFS(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	
	var dfs func(int, []int)
	dfs = func(index int, current []int) {
		if index == len(nums) {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		
		// Include current element
		current = append(current, nums[index])
		dfs(index+1, current)
		current = current[:len(current)-1]
		
		// Skip duplicates
		next := index + 1
		for next < len(nums) && nums[next] == nums[index] {
			next++
		}
		dfs(next, current)
	}
	
	dfs(0, []int{})
	return result
}

// subsetsWithDupBFS uses BFS approach.
func subsetsWithDupBFS(nums []int) [][]int {
	sort.Ints(nums)
	
	result := [][]int{{}} // Start with empty subset
	
	size := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		// If current element is same as previous, only add to subsets created in previous iteration
		if i > 0 && nums[i] == nums[i-1] {
			start = size
		} else {
			start = 0
		}
		
		size = len(result)
		for j := start; j < size; j++ {
			// Create new subset by adding nums[i] to existing subset
			newSubset := make([]int, len(result[j]))
			copy(newSubset, result[j])
			newSubset = append(newSubset, nums[i])
			result = append(result, newSubset)
		}
	}
	
	return result
}

// subsetsWithDupOptimized is an optimized version.
func subsetsWithDupOptimized(nums []int) [][]int {
	sort.Ints(nums)
	
	result := [][]int{}
	current := []int{}
	
	var backtrack func(int)
	backtrack = func(start int) {
		result = append(result, append([]int{}, current...))
		
		for i := start; i < len(nums); i++ {
			// Skip duplicates
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}
	
	backtrack(0)
	return result
}

// SubsetsWithDup is the public interface function.
// It uses the optimized backtracking solution by default.
func SubsetsWithDup(nums []int) [][]int {
	return subsetsWithDupOptimized(nums)
}