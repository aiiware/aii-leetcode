package leetcode

// Convert solves LeetCode problem 0006: Zigzag Conversion
// Difficulty: Medium
// Tags: String
//
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows
// and then read line by line. This function converts a string to its zigzag pattern.
//
// Example:
//   Input: s = "PAYPALISHIRING", numRows = 3
//   Output: "PAHNAPLSIIGYIR"
//
// Time complexity: O(n), Space complexity: O(n)
func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Create rows
	rows := make([][]byte, numRows)
	for i := range rows {
		rows[i] = make([]byte, 0)
	}

	// Fill rows in zigzag pattern
	row := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])

		// Change direction when reaching top or bottom row
		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			row++
		} else {
			row--
		}
	}

	// Combine rows
	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}

// ConvertOptimized optimized mathematical approach.
// Time complexity: O(n), Space complexity: O(n)
func ConvertOptimized(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]byte, len(s))
	index := 0
	cycleLen := 2*numRows - 2

	for row := 0; row < numRows; row++ {
		for i := row; i < len(s); i += cycleLen {
			result[index] = s[i]
			index++

			// For middle rows, add the "diagonal" character
			if row > 0 && row < numRows-1 {
				diagIdx := i + cycleLen - 2*row
				if diagIdx < len(s) {
					result[index] = s[diagIdx]
					index++
				}
			}
		}
	}

	return string(result)
}