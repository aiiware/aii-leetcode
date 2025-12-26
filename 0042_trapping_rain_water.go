// 0042 - Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/
// Hard - Array, Two Pointers, Dynamic Programming, Stack

package leetcode

// Trap calculates how much water can be trapped after raining.
// Time Complexity: O(n)
// Space Complexity: O(1) (two pointers) or O(n) (DP/stack)
func Trap(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    
    // Two pointers approach
    left, right := 0, n-1
    leftMax, rightMax := 0, 0
    water := 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                water += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                water += rightMax - height[right]
            }
            right--
        }
    }
    
    return water
}

// TrapDP uses dynamic programming approach
func TrapDP(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    
    // Left max array
    leftMax := make([]int, n)
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])
    }
    
    // Right max array
    rightMax := make([]int, n)
    rightMax[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i])
    }
    
    // Calculate water
    water := 0
    for i := 0; i < n; i++ {
        water += min(leftMax[i], rightMax[i]) - height[i]
    }
    
    return water
}

// TrapStack uses stack approach
func TrapStack(height []int) int {
    n := len(height)
    if n < 3 {
        return 0
    }
    
    stack := []int{} // stores indices
    water := 0
    
    for i := 0; i < n; i++ {
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            // Pop the top
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            if len(stack) == 0 {
                break
            }
            
            // Calculate distance
            distance := i - stack[len(stack)-1] - 1
            // Calculate bounded height
            boundedHeight := min(height[i], height[stack[len(stack)-1]]) - height[top]
            water += distance * boundedHeight
        }
        stack = append(stack, i)
    }
    
    return water
}

// Helper functions
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}