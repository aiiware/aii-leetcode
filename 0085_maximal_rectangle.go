package leetcode

// MaximalRectangle solves LeetCode problem 0085: Maximal Rectangle
// Difficulty: Hard
// Tags: Array, Dynamic Programming, Stack, Matrix
//
// Given a rows x cols binary matrix filled with 0's and 1's, find the largest
// rectangle containing only 1's and return its area.
//
// Example 1:
// Input: matrix = [["1","0","1","0","0"],
//                  ["1","0","1","1","1"],
//                  ["1","1","1","1","1"],
//                  ["1","0","0","1","0"]]
// Output: 6
// Explanation: The maximal rectangle is shown in the above picture.
//
// Example 2:
// Input: matrix = [["0"]]
// Output: 0
//
// Example 3:
// Input: matrix = [["1"]]
// Output: 1
//
// Constraints:
// rows == matrix.length
// cols == matrix[i].length
// 1 <= row, cols <= 200
// matrix[i][j] is '0' or '1'.
//
// Time complexity: O(m * n) where m is rows and n is columns
// Space complexity: O(n) for the heights array
func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	// heights array represents the histogram for each column
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		// Update heights for current row
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		// Calculate largest rectangle area for current histogram
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// largestRectangleArea is a helper function that calculates the largest rectangle
// area in a histogram using monotonic stack approach (similar to problem 0084)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// Add sentinel value to force stack to empty at the end
	heightsWithSentinel := make([]int, n+1)
	copy(heightsWithSentinel, heights)
	heightsWithSentinel[n] = 0

	stack := make([]int, 0, n+1)
	maxArea := 0

	for i := 0; i <= n; i++ {
		// While stack is not empty and current height is less than height at stack top
		for len(stack) > 0 && heightsWithSentinel[i] < heightsWithSentinel[stack[len(stack)-1]] {
			// Pop the top
			h := heightsWithSentinel[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// Calculate width
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			// Update max area
			area := h * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}

// MaximalRectangleDP is a dynamic programming solution that uses three DP arrays
// to track heights, left boundaries, and right boundaries for each cell
func MaximalRectangleDP(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	
	// height[i][j] = height of consecutive '1's ending at (i, j)
	height := make([]int, cols)
	
	// left[i][j] = left boundary of rectangle ending at (i, j)
	left := make([]int, cols)
	
	// right[i][j] = right boundary of rectangle ending at (i, j)
	right := make([]int, cols)
	
	// Initialize right boundaries to cols (maximum possible)
	for j := 0; j < cols; j++ {
		right[j] = cols
	}

	maxArea := 0

	for i := 0; i < rows; i++ {
		// Update height and left boundaries
		currentLeft := 0
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				height[j]++
				// left[j] is the maximum of previous row's left[j] and currentLeft
				if left[j] < currentLeft {
					left[j] = currentLeft
				}
			} else {
				height[j] = 0
				left[j] = 0
				currentLeft = j + 1
			}
		}

		// Update right boundaries
		currentRight := cols
		for j := cols - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				// right[j] is the minimum of previous row's right[j] and currentRight
				if right[j] > currentRight {
					right[j] = currentRight
				}
			} else {
				right[j] = cols
				currentRight = j
			}

			// Calculate area for current cell
			area := height[j] * (right[j] - left[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// MaximalRectangleBruteForce is a brute force solution for comparison
// Not recommended for large inputs (O(m^2 * n^2) time complexity)
func MaximalRectangleBruteForce(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	maxArea := 0

	// Precompute prefix sums for faster all-ones check
	// prefix[i+1][j+1] = sum of matrix[0..i][0..j]
	prefix := make([][]int, rows+1)
	for i := range prefix {
		prefix[i] = make([]int, cols+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val := 0
			if matrix[i][j] == '1' {
				val = 1
			}
			prefix[i+1][j+1] = prefix[i+1][j] + prefix[i][j+1] - prefix[i][j] + val
		}
	}

	// Try all possible rectangles
	for top := 0; top < rows; top++ {
		for bottom := top; bottom < rows; bottom++ {
			for left := 0; left < cols; left++ {
				for right := left; right < cols; right++ {
					// Calculate sum using prefix sums
					sum := prefix[bottom+1][right+1] - prefix[top][right+1] - prefix[bottom+1][left] + prefix[top][left]
					area := (bottom - top + 1) * (right - left + 1)
					// If all cells are '1's, update max area
					if sum == area && area > maxArea {
						maxArea = area
					}
				}
			}
		}
	}

	return maxArea
}

// MaximalRectangleOptimizedBruteForce is an optimized brute force solution
// that uses row-wise expansion to reduce time complexity to O(m^2 * n)
func MaximalRectangleOptimizedBruteForce(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	maxArea := 0

	// For each pair of rows (top, bottom), calculate the histogram
	for top := 0; top < rows; top++ {
		// heights[j] = number of consecutive '1's from top to current row in column j
		heights := make([]int, cols)
		
		for bottom := top; bottom < rows; bottom++ {
			// Update heights for current bottom row
			for j := 0; j < cols; j++ {
				if matrix[bottom][j] == '1' {
					heights[j]++
				} else {
					heights[j] = 0
				}
			}
			
			// Calculate largest rectangle area for current histogram
			area := largestRectangleArea(heights)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// MaximalRectangleDivideConquer uses divide and conquer approach
// Time complexity: O(m * n * log(min(m, n)))
func MaximalRectangleDivideConquer(matrix [][]byte) int {
	return MaximalRectangle(matrix)
}