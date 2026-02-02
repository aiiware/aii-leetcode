package arrays

// 0240. Search a 2D Matrix II
// https://leetcode.com/problems/search-a-2d-matrix-ii/
//
// Write an efficient algorithm that searches for a value target in an m x n
// integer matrix matrix. This matrix has the following properties:
// - Integers in each row are sorted in ascending from left to right.
// - Integers in each column are sorted in ascending from top to bottom.

// SearchMatrixII uses the "search from top-right corner" technique.
// Starting from top-right, we can eliminate either a row or column at each step.
// Time complexity: O(m + n), Space complexity: O(1)
func SearchMatrixII(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	
	// Start from top-right corner
	row, col := 0, n-1

	for row < m && col >= 0 {
		current := matrix[row][col]
		
		if current == target {
			return true
		} else if current > target {
			// If current is greater than target, target cannot be in this column
			// (since all elements in this column are >= current)
			col--
		} else {
			// If current is less than target, target cannot be in this row
			// (since all elements in this row are <= current)
			row++
		}
	}

	return false
}

// SearchMatrixIIBinarySearch uses binary search on each row.
// Since each row is sorted, we can binary search each row.
// Time complexity: O(m log n), Space complexity: O(1)
func SearchMatrixIIBinarySearch(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		// Skip rows where target is out of range
		if len(row) == 0 || target < row[0] || target > row[len(row)-1] {
			continue
		}
		
		// Binary search in this row
		left, right := 0, len(row)-1
		for left <= right {
			mid := left + (right-left)/2
			if row[mid] == target {
				return true
			} else if row[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

// SearchMatrixIIDivideConquer is a recursive divide and conquer approach.
// Not as efficient but demonstrates another way to think about the problem.
func SearchMatrixIIDivideConquer(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return searchMatrixIIRecursive(matrix, target, 0, 0, len(matrix)-1, len(matrix[0])-1)
}

func searchMatrixIIRecursive(matrix [][]int, target, top, left, bottom, right int) bool {
	// Base case: invalid submatrix
	if top > bottom || left > right {
		return false
	}

	// Base case: single cell
	if top == bottom && left == right {
		return matrix[top][left] == target
	}

	// Find middle row and column
	midRow := top + (bottom-top)/2
	midCol := left + (right-left)/2
	midVal := matrix[midRow][midCol]

	if midVal == target {
		return true
	} else if midVal > target {
		// Target is smaller than middle value
		// Search in top-left, top-right, and bottom-left quadrants
		return searchMatrixIIRecursive(matrix, target, top, left, midRow, midCol) ||
			searchMatrixIIRecursive(matrix, target, top, midCol+1, midRow, right) ||
			searchMatrixIIRecursive(matrix, target, midRow+1, left, bottom, midCol)
	} else {
		// Target is larger than middle value
		// Search in top-right, bottom-left, and bottom-right quadrants
		return searchMatrixIIRecursive(matrix, target, top, midCol+1, midRow, right) ||
			searchMatrixIIRecursive(matrix, target, midRow+1, left, bottom, midCol) ||
			searchMatrixIIRecursive(matrix, target, midRow+1, midCol+1, bottom, right)
	}
}