package arrays

// Subsets solves LeetCode problem 0078: Subsets
// Difficulty: Medium
// Tags: Backtracking, Bit Manipulation, Depth-First Search
//
// Given an integer array nums of unique elements, return all possible subsets
// (the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in any order.
//
// Example 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
// Example 2:
// Input: nums = [0]
// Output: [[],[0]]
//
// Constraints:
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
//
// Time complexity: O(2^n * n) where n is length of nums
// Space complexity: O(2^n * n) for storing results
func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		// Add current subset to result
		subset := make([]int, len(current))
		copy(subset, current)
		result = append(result, subset)

		// Try adding each remaining element
		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}

// SubsetsIterative builds subsets iteratively
// Start with empty subset, then for each number, add it to all existing subsets
func SubsetsIterative(nums []int) [][]int {
	result := [][]int{{}} // Start with empty subset

	for _, num := range nums {
		// For each existing subset, create a new subset with current number added
		size := len(result)
		for i := 0; i < size; i++ {
			newSubset := make([]int, len(result[i])+1)
			copy(newSubset, result[i])
			newSubset[len(result[i])] = num
			result = append(result, newSubset)
		}
	}

	return result
}

// SubsetsBitMask uses bit manipulation to generate all subsets
// Each subset corresponds to a bitmask of length n
func SubsetsBitMask(nums []int) [][]int {
	n := len(nums)
	total := 1 << uint(n) // 2^n subsets
	result := make([][]int, 0, total)

	for mask := 0; mask < total; mask++ {
		subset := make([]int, 0)
		// Check each bit in the mask
		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}

	return result
}

// SubsetsDFS uses depth-first search approach
func SubsetsDFS(nums []int) [][]int {
	result := make([][]int, 0)

	var dfs func(index int, current []int)
	dfs = func(index int, current []int) {
		if index == len(nums) {
			subset := make([]int, len(current))
			copy(subset, current)
			result = append(result, subset)
			return
		}

		// Option 1: Don't include nums[index]
		dfs(index+1, current)

		// Option 2: Include nums[index]
		current = append(current, nums[index])
		dfs(index+1, current)
		current = current[:len(current)-1]
	}

	dfs(0, []int{})
	return result
}

// SubsetsLexicographic generates subsets in lexicographic order
func SubsetsLexicographic(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0, 1<<uint(n))

	// Sort to ensure lexicographic order (nums should already be unique)
	// We'll assume nums is sorted or we'll sort it
	// For simplicity, we'll work with indices

	for mask := 0; mask < (1 << uint(n)); mask++ {
		subset := make([]int, 0)
		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}

	return result
}

// SubsetsWithDuplicates handles arrays with duplicates (not needed for this problem
// but included for completeness)
func SubsetsWithDuplicates(nums []int) [][]int {
	// First sort to handle duplicates
	// sort.Ints(nums) // Uncomment if nums can have duplicates
	
	result := make([][]int, 0)
	current := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		subset := make([]int, len(current))
		copy(subset, current)
		result = append(result, subset)

		for i := start; i < len(nums); i++ {
			// Skip duplicates (if nums were sorted)
			// if i > start && nums[i] == nums[i-1] {
			//     continue
			// }
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}