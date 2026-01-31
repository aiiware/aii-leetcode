package dp


/*
Difficulty: Easy
Tags: [Add relevant tags]
Companies: [Add company names]
*/

// Problem 119: Pascal's Triangle II
// Given an integer rowIndex, return the rowIndex-th (0-indexed) row of Pascal's triangle.
// Could you optimize your algorithm to use only O(rowIndex) extra space?

// getRow returns the rowIndex-th row of Pascal's triangle.
// This implementation uses O(rowIndex) extra space.
func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	// Initialize with first row [1]
	row := make([]int, rowIndex+1)
	row[0] = 1

	// Generate each row iteratively
	for i := 1; i <= rowIndex; i++ {
		// Update from right to left to avoid overwriting values we need
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}

// getRowTwoArrays uses two arrays for clarity (still O(rowIndex) space).
func getRowTwoArrays(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	prev := make([]int, rowIndex+1)
	curr := make([]int, rowIndex+1)
	prev[0] = 1

	for i := 1; i <= rowIndex; i++ {
		curr[0] = 1
		curr[i] = 1
		for j := 1; j < i; j++ {
			curr[j] = prev[j-1] + prev[j]
		}
		// Swap prev and curr for next iteration
		prev, curr = curr, prev
	}

	// After the last iteration, prev contains the result
	return prev[:rowIndex+1]
}

// getRowMath uses combinatorial formula: C(rowIndex, k) = rowIndex! / (k! * (rowIndex - k)!)
// This is O(rowIndex) time and O(1) extra space (excluding output).
func getRowMath(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	row := make([]int, rowIndex+1)
	row[0] = 1

	// Use the property: C(n, k) = C(n, k-1) * (n - k + 1) / k
	for k := 1; k <= rowIndex; k++ {
		// Calculate using integer arithmetic to avoid overflow
		// row[k] = row[k-1] * (rowIndex - k + 1) / k
		val := row[k-1]
		val *= (rowIndex - k + 1)
		val /= k
		row[k] = val
	}

	return row
}

// Helper function to compare two rows
func rowsEqual(r1, r2 []int) bool {
	if len(r1) != len(r2) {
		return false
	}
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

// Helper function to get multiple rows (for testing/debugging)
func getRowsUpTo(rowIndex int) [][]int {
	result := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = getRow(i)
	}
	return result
}