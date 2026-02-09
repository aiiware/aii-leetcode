package arrays

/*
174. Dungeon Game

The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon.
The dungeon consists of m x n rooms laid out in a 2D grid. Our valiant knight was initially
positioned in the top-left room and must fight his way through dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his
health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons (represented by negative integers), so the knight loses
health upon entering these rooms; other rooms are either empty (represented as 0) or contain
magic orbs that increase the knight's health (represented by positive integers).

To reach the princess as quickly as possible, the knight decides to move only rightward or
downward in each step.

Return the knight's minimum initial health so that he can rescue the princess.

Note that any room can contain threats or power-ups, even the first room the knight enters and
the bottom-right room where the princess is imprisoned.

Example 1:
Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
Output: 7
Explanation: The initial health of the knight must be at least 7 if he follows the optimal path:
RIGHT -> RIGHT -> DOWN -> DOWN.

Example 2:
Input: dungeon = [[0]]
Output: 1

Constraints:
- m == dungeon.length
- n == dungeon[i].length
- 1 <= m, n <= 200
- -1000 <= dungeon[i][j] <= 1000

Difficulty: Hard
Tags: Array, Dynamic Programming, Matrix
Companies: Microsoft, Google, Amazon, Bloomberg
*/

// calculateMinimumHP calculates the minimum initial health required to rescue the princess
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	m, n := len(dungeon), len(dungeon[0])

	// Create DP table with extra row and column for boundary conditions
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = 1<<31 - 1 // Initialize with large value (max int)
		}
	}

	// The princess room: knight needs at least 1 health after this cell
	// If the cell has negative value, he needs more health to survive
	// If the cell has positive value, he needs at least 1 health
	dp[m][n-1], dp[m-1][n] = 1, 1

	// Fill DP table from bottom-right to top-left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// Minimum health needed from (i,j) to reach princess
			// We need to choose the path that requires minimum initial health
			minHealth := min(dp[i+1][j], dp[i][j+1])

			// If dungeon[i][j] provides health, we need less initial health
			// If dungeon[i][j] consumes health, we need more initial health
			// The formula: dp[i][j] = max(1, minHealth - dungeon[i][j])
			// We need at least 1 health at any point
			dp[i][j] = maximum(1, minHealth-dungeon[i][j])
		}
	}

	return dp[0][0]
}

// Alternative solution with space optimization (O(n) space)
func calculateMinimumHP2(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	m, n := len(dungeon), len(dungeon[0])

	// Use only 1D DP array for space optimization
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1<<31 - 1
	}
	dp[n-1] = 1 // Initialize the cell to the right of princess

	// Fill DP table from bottom-right to top-left
	for i := m - 1; i >= 0; i-- {
		// Update dp[n] for current row (boundary condition)
		dp[n] = 1<<31 - 1
		if i == m-1 {
			dp[n] = 1
		}

		for j := n - 1; j >= 0; j-- {
			// Minimum health needed from (i,j) to reach princess
			minHealth := min(dp[j], dp[j+1])
			dp[j] = maximum(1, minHealth-dungeon[i][j])
		}
	}

	return dp[0]
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
