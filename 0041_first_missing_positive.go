// 0041 - First Missing Positive
// https://leetcode.com/problems/first-missing-positive/
// Hard - Array, Hash Table

package leetcode

// FirstMissingPositive finds the smallest missing positive integer.
// Time Complexity: O(n)
// Space Complexity: O(1)
func FirstMissingPositive(nums []int) int {
    n := len(nums)
    
    // Step 1: Place each number in its correct position if possible
    // nums[i] should be at position nums[i]-1 (if 1 <= nums[i] <= n)
    for i := 0; i < n; i++ {
        // While nums[i] is in range [1, n] and not in correct position
        for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            // Swap nums[i] with nums[nums[i]-1]
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    
    // Step 2: Find first position where nums[i] != i+1
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }
    
    // All numbers 1..n are present, so missing is n+1
    return n + 1
}

// FirstMissingPositiveHash uses hash table approach
func FirstMissingPositiveHash(nums []int) int {
    n := len(nums)
    
    // Mark numbers that are out of range as 0
    for i := 0; i < n; i++ {
        if nums[i] <= 0 || nums[i] > n {
            nums[i] = 0
        }
    }
    
    // Use index as hash: for each number, mark its position as visited
    for i := 0; i < n; i++ {
        if nums[i] != 0 {
            // Get the original value (might have been modified)
            val := nums[i]
            if val < 0 {
                val = -val
            }
            
            // Mark position val-1 as visited by making it negative
            if nums[val-1] == 0 {
                nums[val-1] = -val
            } else if nums[val-1] > 0 {
                nums[val-1] = -nums[val-1]
            }
        }
    }
    
    // Find first positive index
    for i := 0; i < n; i++ {
        if nums[i] >= 0 {
            return i + 1
        }
    }
    
    return n + 1
}