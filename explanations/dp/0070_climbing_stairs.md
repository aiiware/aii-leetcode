# 70. Climbing Stairs - Solution Explanation

## Problem Statement
You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Difficulty: Easy

## Key Insights
1. **Fibonacci Sequence**: The number of ways to reach step `n` equals the number of ways to reach step `n-1` plus the number of ways to reach step `n-2`
2. **Dynamic Programming**: This is a classic DP problem where we can build up the solution from base cases
3. **Optimal Substructure**: The solution for `n` depends on solutions for `n-1` and `n-2`

## Solution Approaches

### Approach 1: Dynamic Programming (Bottom-Up)
**Time Complexity**: O(n)
**Space Complexity**: O(n) - can be optimized to O(1)

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    dp := make([]int, n+1)
    dp[1] = 1  // 1 way to reach step 1: [1]
    dp[2] = 2  // 2 ways to reach step 2: [1,1] or [2]
    
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    
    return dp[n]
}
```

### Approach 2: Optimized DP (Constant Space)
**Time Complexity**: O(n)
**Space Complexity**: O(1)

```go
func climbStairsOptimized(n int) int {
    if n <= 2 {
        return n
    }
    
    // Track only the last two values
    prev2 := 1  // ways to reach step i-2
    prev1 := 2  // ways to reach step i-1
    
    for i := 3; i <= n; i++ {
        current := prev1 + prev2
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}
```

### Approach 3: Matrix Exponentiation (Advanced)
**Time Complexity**: O(log n)
**Space Complexity**: O(1)

```go
func climbStairsMatrix(n int) int {
    if n <= 2 {
        return n
    }
    
    // Using matrix exponentiation to compute Fibonacci
    // [F(n+1) F(n)  ] = [1 1]^n
    // [F(n)   F(n-1)]   [1 0]
    
    // For simplicity, using DP for this example
    // In practice, matrix exponentiation would be implemented
    return climbStairsOptimized(n)
}
```

## Step-by-Step Walkthrough

### Example: n = 5

**Step 1**: Base cases
- Step 1: 1 way ([1])
- Step 2: 2 ways ([1,1], [2])

**Step 2**: Build up solution
- Step 3 = Step 2 + Step 1 = 2 + 1 = 3 ways
  - From Step 2: [1,1,1], [2,1]
  - From Step 1: [1,2]
  
- Step 4 = Step 3 + Step 2 = 3 + 2 = 5 ways
  - From Step 3: [1,1,1,1], [2,1,1], [1,2,1]
  - From Step 2: [1,1,2], [2,2]
  
- Step 5 = Step 4 + Step 3 = 5 + 3 = 8 ways
  - From Step 4: [1,1,1,1,1], [2,1,1,1], [1,2,1,1], [1,1,2,1], [2,2,1]
  - From Step 3: [1,1,1,2], [2,1,2], [1,2,2]

**Visualization**:
```
Step: 0   1   2   3   4   5
Ways: 1   1   2   3   5   8
      ↑   ↑   ↑   ↑   ↑   ↑
      Base Ways to reach each step
```

## Complexity Analysis

### Time Complexity
- **DP Approach**: O(n) - Single pass through steps 3 to n
- **Optimized DP**: O(n) - Same time, less space
- **Matrix Exponentiation**: O(log n) - Using fast exponentiation

### Space Complexity
- **DP Array**: O(n) - Store results for all steps
- **Optimized DP**: O(1) - Store only last two values
- **Matrix Exponentiation**: O(1) - Constant space for matrix operations

## Common Pitfalls
1. **Off-by-one errors**: Confusing step count with array indices
2. **Missing base cases**: Forgetting to handle n = 0, 1, 2
3. **Integer overflow**: For large n, results can exceed int32 limits (use int64)
4. **Recursion without memoization**: Naive recursion leads to O(2^n) time

## Optimization Tips
1. **Use constant space**: Only need last two values, not entire array
2. **Precompute for multiple queries**: If solving for many n values, cache results
3. **Consider matrix exponentiation**: For extremely large n (n > 10^9)
4. **Use modulo arithmetic**: If problem asks for result modulo some number

## Edge Cases
1. **n = 0**: 1 way (do nothing) - though problem states n ≥ 1
2. **n = 1**: 1 way ([1])
3. **n = 2**: 2 ways ([1,1], [2])
4. **Large n**: Use 64-bit integers to prevent overflow

## Related Problems
- **509. Fibonacci Number** - Same recurrence relation
- **746. Min Cost Climbing Stairs** - Variation with costs
- **1137. N-th Tribonacci Number** - Three-term recurrence
- **198. House Robber** - Similar DP structure with constraints

## Practice Exercises
1. **Variation 1**: Allow steps of 1, 2, or 3 at a time
2. **Variation 2**: Some steps are broken (cannot step on them)
3. **Variation 3**: Find all possible sequences, not just count
4. **Variation 4**: Minimum number of steps to reach top
5. **Challenge**: Implement matrix exponentiation solution

## Mathematical Insight
The number of ways to climb n stairs is the (n+1)th Fibonacci number:
- F(0) = 0, F(1) = 1, F(2) = 1, F(3) = 2, F(4) = 3, F(5) = 5, F(6) = 8
- climbStairs(n) = F(n+1)

This can be proven by induction:
- Base: climbStairs(1) = 1 = F(2), climbStairs(2) = 2 = F(3)
- Inductive: climbStairs(n) = climbStairs(n-1) + climbStairs(n-2) = F(n) + F(n-1) = F(n+1)

## Additional Notes
- This problem is often used as an introduction to dynamic programming
- The Fibonacci connection makes it a good example of mathematical patterns in DP
- Many interview questions build upon this basic pattern
- Understanding this problem helps with more complex DP problems like "Decode Ways" or "Unique Paths"