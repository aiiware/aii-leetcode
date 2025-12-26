// 0045 - Jump Game II
// https://leetcode.com/problems/jump-game-ii/
// Medium - Array, Dynamic Programming, Greedy

package leetcode

// Jump calculates the minimum number of jumps to reach the last index.
// Time Complexity: O(n)
// Space Complexity: O(1)
func Jump(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }
    
    jumps := 0
    currentEnd := 0
    farthest := 0
    
    for i := 0; i < n-1; i++ {
        // Update the farthest we can reach
        if i+nums[i] > farthest {
            farthest = i + nums[i]
        }
        
        // If we've reached the end of current jump
        if i == currentEnd {
            jumps++
            currentEnd = farthest
            
            // If we can reach or exceed the last index
            if currentEnd >= n-1 {
                break
            }
        }
    }
    
    return jumps
}

// JumpDP uses dynamic programming approach
func JumpDP(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }
    
    // dp[i] = minimum jumps to reach index i
    dp := make([]int, n)
    for i := range dp {
        dp[i] = n // Initialize with max value
    }
    dp[0] = 0
    
    for i := 0; i < n; i++ {
        maxJump := nums[i]
        for j := 1; j <= maxJump && i+j < n; j++ {
            if dp[i]+1 < dp[i+j] {
                dp[i+j] = dp[i] + 1
            }
        }
    }
    
    return dp[n-1]
}

// JumpBFS uses BFS-like approach
func JumpBFS(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }
    
    level := 0
    currentMax := 0
    nextMax := 0
    i := 0
    
    for currentMax < n-1 {
        level++
        for ; i <= currentMax; i++ {
            if i+nums[i] > nextMax {
                nextMax = i + nums[i]
            }
            if nextMax >= n-1 {
                return level
            }
        }
        currentMax = nextMax
    }
    
    return level
}