// 0046 - Permutations
// https://leetcode.com/problems/permutations/
// Medium - Array, Backtracking

package leetcode

// Permute generates all permutations of distinct integers.
// Time Complexity: O(n * n!) where n is length of nums
// Space Complexity: O(n!) for storing results
func Permute(nums []int) [][]int {
    result := [][]int{}
    backtrackPermute(nums, 0, &result)
    return result
}

// backtrackPermute is the recursive backtracking function
func backtrackPermute(nums []int, start int, result *[][]int) {
    if start == len(nums) {
        // Make a copy of current permutation
        perm := make([]int, len(nums))
        copy(perm, nums)
        *result = append(*result, perm)
        return
    }
    
    for i := start; i < len(nums); i++ {
        // Swap nums[start] with nums[i]
        nums[start], nums[i] = nums[i], nums[start]
        // Recurse
        backtrackPermute(nums, start+1, result)
        // Backtrack (swap back)
        nums[start], nums[i] = nums[i], nums[start]
    }
}

// PermuteHeap uses Heap's algorithm (correct implementation)
func PermuteHeap(nums []int) [][]int {
    result := [][]int{}
    n := len(nums)
    
    // Helper function for Heap's algorithm
    var generate func(int)
    generate = func(k int) {
        if k == 1 {
            perm := make([]int, n)
            copy(perm, nums)
            result = append(result, perm)
            return
        }
        
        // Generate permutations with k-th unaltered
        generate(k - 1)
        
        // Generate permutations for k-th swapped with each k-1 initial
        for i := 0; i < k-1; i++ {
            if k%2 == 0 {
                // Even k: swap i and k-1
                nums[i], nums[k-1] = nums[k-1], nums[i]
            } else {
                // Odd k: swap 0 and k-1
                nums[0], nums[k-1] = nums[k-1], nums[0]
            }
            generate(k - 1)
        }
    }
    
    generate(n)
    return result
}

// PermuteIterative uses iterative approach
func PermuteIterative(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{{}} // Fixed: empty array has one permutation: empty permutation
    }
    
    result := [][]int{{nums[0]}}
    
    for i := 1; i < len(nums); i++ {
        newResult := [][]int{}
        num := nums[i]
        
        for _, perm := range result {
            // Insert num at every possible position in each permutation
            for j := 0; j <= len(perm); j++ {
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