package dp

import "fmt"

// Problem 118: Pascal's Triangle
// Given an integer numRows, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.

// generate returns the first numRows of Pascal's triangle.
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	// Initialize the triangle with the first row
	triangle := make([][]int, numRows)
	triangle[0] = []int{1}

	// Generate each subsequent row
	for i := 1; i < numRows; i++ {
		prevRow := triangle[i-1]
		currRow := make([]int, i+1)

		// First and last elements are always 1
		currRow[0] = 1
		currRow[i] = 1

		// Calculate middle elements
		for j := 1; j < i; j++ {
			currRow[j] = prevRow[j-1] + prevRow[j]
		}

		triangle[i] = currRow
	}

	return triangle
}

// generateOptimized is a slightly optimized version that pre-allocates all rows.
func generateOptimized(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}

	return triangle
}

// Helper function to print Pascal's triangle (for debugging/demo)
func printTriangle(triangle [][]int) string {
	if len(triangle) == 0 {
		return "[]"
	}

	result := ""
	for i, row := range triangle {
		// Add indentation for pyramid shape
		indent := len(triangle) - i - 1
		result += "  " // Base indentation

		for k := 0; k < indent; k++ {
			result += "  "
		}

		// Add row elements
		for j, num := range row {
			result += fmt.Sprintf("%3d", num)
			if j < len(row)-1 {
				result += "   "
			}
		}
		if i < len(triangle)-1 {
			result += "\n"
		}
	}
	return result
}

// Helper function to compare two triangles
func trianglesEqual(t1, t2 [][]int) bool {
	if len(t1) != len(t2) {
		return false
	}

	for i := range t1 {
		if len(t1[i]) != len(t2[i]) {
			return false
		}
		for j := range t1[i] {
			if t1[i][j] != t2[i][j] {
				return false
			}
		}
	}
	return true
}