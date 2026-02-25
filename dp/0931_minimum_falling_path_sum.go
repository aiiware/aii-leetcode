package dp

// 931. Minimum Falling Path Sum
//
// Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
//
// A falling path starts at any element in the first row and chooses the element in the next row that is
// either directly below or diagonally left/right. Specifically, the next element from position (row, col)
// will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
//
// Example 1:
// Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
// Output: 13
// Explanation: There are two falling paths with a minimum sum:
// 1) 2 → 1 → 7 = 10
// 2) 2 → 5 → 8 = 15
// 3) 2 → 4 → 9 = 15
// 4) 1 → 5 → 7 = 13
// 5) 1 → 4 → 8 = 13
// 6) 1 → 4 → 9 = 14
// 7) 3 → 5 → 7 = 15
// 8) 3 → 4 → 8 = 15
// 9) 3 → 4 → 9 = 16
// Minimum is 13 (path: 1 → 5 → 7 or 1 → 4 → 8)
//
// Example 2:
// Input: matrix = [[-19,57],[-40,-5]]
// Output: -59
// Explanation: The falling path with a minimum sum is -19 → -40 = -59.
//
// Example 3:
// Input: matrix = [[-48]]
// Output: -48
//
// Constraints:
// - n == matrix.length == matrix[i].length
// - 1 <= n <= 100
// - -100 <= matrix[i][j] <= 100

// MinFallingPathSum returns the minimum sum of any falling path through matrix.
// This is the main public function that uses the optimized solution by default.
func MinFallingPathSum(matrix [][]int) int {
	return minFallingPathSumOptimized(matrix)
}

// minFallingPathSumDP solves using dynamic programming with O(n²) time and O(n²) space.
// It creates a DP table where dp[i][j] represents the minimum falling path sum ending at (i, j).
func minFallingPathSumDP(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	
	// Create DP table
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// Initialize first row (base case)
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}
	
	// Fill DP table row by row
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			// Current cell value
			current := matrix[i][j]
			
			// Minimum from possible previous positions
			minPrev := dp[i-1][j] // from directly above
			
			// Check top-left if exists
			if j > 0 && dp[i-1][j-1] < minPrev {
				minPrev = dp[i-1][j-1]
			}
			
			// Check top-right if exists
			if j < n-1 && dp[i-1][j+1] < minPrev {
				minPrev = dp[i-1][j+1]
			}
			
			dp[i][j] = current + minPrev
		}
	}
	
	// Find minimum in last row
	result := dp[n-1][0]
	for j := 1; j < n; j++ {
		if dp[n-1][j] < result {
			result = dp[n-1][j]
		}
	}
	
	return result
}

// minFallingPathSumOptimized solves using dynamic programming with O(n²) time and O(n) space.
// It uses only two rows (previous and current) instead of the full DP table.
func minFallingPathSumOptimized(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	
	// Initialize previous row with first row values
	prev := make([]int, n)
	for j := 0; j < n; j++ {
		prev[j] = matrix[0][j]
	}
	
	// Process remaining rows
	for i := 1; i < n; i++ {
		curr := make([]int, n)
		
		for j := 0; j < n; j++ {
			// Current cell value
			current := matrix[i][j]
			
			// Minimum from possible previous positions
			minPrev := prev[j] // from directly above
			
			// Check top-left if exists
			if j > 0 && prev[j-1] < minPrev {
				minPrev = prev[j-1]
			}
			
			// Check top-right if exists
			if j < n-1 && prev[j+1] < minPrev {
				minPrev = prev[j+1]
			}
			
			curr[j] = current + minPrev
		}
		
		// Update previous row for next iteration
		prev = curr
	}
	
	// Find minimum in last row (now stored in prev)
	result := prev[0]
	for j := 1; j < n; j++ {
		if prev[j] < result {
			result = prev[j]
		}
	}
	
	return result
}

// minFallingPathSumInPlace solves by modifying the input matrix in-place.
// This has O(n²) time and O(1) extra space (modifies input).
func minFallingPathSumInPlace(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	
	// Process rows starting from the second row
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			// Minimum from possible previous positions
			minPrev := matrix[i-1][j] // from directly above
			
			// Check top-left if exists
			if j > 0 && matrix[i-1][j-1] < minPrev {
				minPrev = matrix[i-1][j-1]
			}
			
			// Check top-right if exists
			if j < n-1 && matrix[i-1][j+1] < minPrev {
				minPrev = matrix[i-1][j+1]
			}
			
			// Update current cell with minimum path sum ending here
			matrix[i][j] += minPrev
		}
	}
	
	// Find minimum in last row
	result := matrix[n-1][0]
	for j := 1; j < n; j++ {
		if matrix[n-1][j] < result {
			result = matrix[n-1][j]
		}
	}
	
	return result
}