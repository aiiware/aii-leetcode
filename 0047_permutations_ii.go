// 0047 - Permutations II
// https://leetcode.com/problems/permutations-ii/
// Medium - Array, Backtracking

package leetcode

import "sort"

// PermuteUnique generates all unique permutations of integers that may contain duplicates.
// Time Complexity: O(n * n!) where n is length of nums
// Space Complexity: O(n!) for storing results
func PermuteUnique(nums []int) [][]int {
    // Sort to handle duplicates
    sort.Ints(nums)
    
    result := [][]int{}
    used := make([]bool, len(nums))
    current := []int{}
    
    backtrackPermuteUnique(nums, used, current, &result)
    return result
}

func backtrackPermuteUnique(nums []int, used []bool, current []int, result *[][]int) {
    if len(current) == len(nums) {
        // Make a copy of current permutation
        perm := make([]int, len(current))
        copy(perm, current)
        *result = append(*result, perm)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        // Skip used elements
        if used[i] {
            continue
        }
        
        // Skip duplicates: if current element is same as previous and previous is not used
        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        
        // Choose nums[i]
        used[i] = true
        current = append(current, nums[i])
        
        // Recurse
        backtrackPermuteUnique(nums, used, current, result)
        
        // Backtrack
        used[i] = false
        current = current[:len(current)-1]
    }
}

// PermuteUniqueSwap uses swap-based approach
func PermuteUniqueSwap(nums []int) [][]int {
    result := [][]int{}
    permuteUniqueSwapHelper(nums, 0, &result)
    return result
}

func permuteUniqueSwapHelper(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        perm := make([]int, len(nums))
        copy(perm, nums)
        *result = append(*result, perm)
        return
    }
    
    seen := make(map[int]bool)
    for i := start; i < len(nums); i++ {
        // Skip duplicates
        if seen[nums[i]] {
            continue
        }
        seen[nums[i]] = true
        
        // Swap
        nums[start], nums[i] = nums[i], nums[start]
        // Recurse
        permuteUniqueSwapHelper(nums, start+1, result)
        // Backtrack
        nums[start], nums[i] = nums[i], nums[start]
    }
}

// PermuteUniqueIterative uses iterative approach
func PermuteUniqueIterative(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }
    
    // Sort to handle duplicates
    sort.Ints(nums)
    
    result := [][]int{{nums[0]}}
    
    for i := 1; i < len(nums); i++ {
        newResult := [][]int{}
        num := nums[i]
        
        for _, perm := range result {
            // Insert num at every possible position
            for j := 0; j <= len(perm); j++ {
                // Skip duplicate insertion if num equals previous element
                if j > 0 && perm[j-1] == num {
                    continue
                }
                
                newPerm := make([]int, len(perm)+1)
                copy(newPerm[:j], perm[:j])
                newPerm[j] = num
                copy(newPerm[j+1:], perm[j:])
                newResult = append(newResult, newPerm)
            }
        }
        
        result = newResult
    }
    
    return result
}