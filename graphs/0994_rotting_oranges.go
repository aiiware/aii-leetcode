package graphs

// Solve0994 solves LeetCode problem 0994: Rotting Oranges
// Difficulty: Medium
// Tags: Breadth-First Search, Matrix, Array
//
// You are given an m x n grid where each cell can have one of three values:
// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
//
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange
// becomes rotten.
//
// Return the minimum number of minutes that must elapse until no cell has a fresh orange.
// If this is impossible, return -1.
//
// Example 1:
// Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
// Output: 4
//
// Example 2:
// Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
// Output: -1
// Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten.
//
// Example 3:
// Input: grid = [[0,2]]
// Output: 0
// Explanation: Since there are no fresh oranges at minute 0, the answer is just 0.
//
// Constraints:
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] is 0, 1, or 2.
//
// Time complexity: O(m * n), Space complexity: O(m * n)
func Solve0994(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	queue := make([][2]int, 0)
	freshCount := 0

	// Initialize: find all rotten oranges and count fresh ones
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	// If no fresh oranges, return 0 immediately
	if freshCount == 0 {
		return 0
	}

	// Directions: up, down, left, right
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	// BFS
	for len(queue) > 0 {
		levelSize := len(queue)
		rottedThisMinute := false

		for i := 0; i < levelSize; i++ {
			cell := queue[0]
			queue = queue[1:]
			row, col := cell[0], cell[1]

			// Check all 4 directions
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]

				// Check bounds and if it's a fresh orange
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == 1 {
					// Rot the fresh orange
					grid[newRow][newCol] = 2
					freshCount--
					queue = append(queue, [2]int{newRow, newCol})
					rottedThisMinute = true
				}
			}
		}

		// If we rotted any oranges this minute, increment minutes
		if rottedThisMinute {
			minutes++
		}
	}

	// If there are still fresh oranges, return -1
	if freshCount > 0 {
		return -1
	}

	return minutes
}