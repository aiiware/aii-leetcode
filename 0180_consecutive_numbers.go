package leetcode

/*
180. Consecutive Numbers

Table: Logs

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| num         | int     |
+-------------+---------+
id is the primary key for this table.
id is an autoincrement primary key.

Find all numbers that appear at least 3 times consecutively.

Return the result table in any order.

Example 1:
Input:
Logs table:
+----+-----+
| id | num |
+----+-----+
| 1  | 1   |
| 2  | 1   |
| 3  | 1   |
| 4  | 2   |
| 5  | 1   |
+----+-----+

Output:
+-----+
| ConsecutiveNums |
+-----+
| 1   |
+-----+

Explanation: 1 is the only number that appears at least 3 times consecutively.

Difficulty: Medium
Tags: Database
Companies: Apple, Yahoo
*/

// ConsecutiveLog represents a row in the Logs table
type ConsecutiveLog struct {
	ID  int
	Num int
}

// FindConsecutiveNumbers finds all numbers that appear at least 3 times consecutively.
// This is the Go equivalent of the SQL solution for LeetCode 0180.
//
// Algorithm:
// 1. Iterate through the logs and check if three consecutive IDs have the same number
// 2. Use a set to track numbers already found to avoid duplicates
// 3. Return the unique consecutive numbers
//
// Time complexity: O(n) where n is the number of logs
// Space complexity: O(k) where k is the number of unique consecutive numbers
func FindConsecutiveNumbers(logs []ConsecutiveLog) []int {
	if len(logs) < 3 {
		return []int{}
	}

	// Use a set to store unique consecutive numbers
	consecutiveSet := make(map[int]bool)

	// Check each potential triplet of consecutive IDs
	for i := 0; i < len(logs)-2; i++ {
		// Check if three consecutive logs have:
		// 1. Consecutive IDs (id[i] + 1 = id[i+1] and id[i+1] + 1 = id[i+2])
		// 2. Same number value
		if logs[i].ID+1 == logs[i+1].ID &&
			logs[i+1].ID+1 == logs[i+2].ID &&
			logs[i].Num == logs[i+1].Num &&
			logs[i+1].Num == logs[i+2].Num {
			consecutiveSet[logs[i].Num] = true
		}
	}

	// Convert set to sorted slice
	result := make([]int, 0, len(consecutiveSet))
	for num := range consecutiveSet {
		result = append(result, num)
	}

	return result
}

// FindConsecutiveNumbersOptimized is an optimized version that doesn't assume
// IDs are in the logs array in order. It works by sorting first.
//
// Time complexity: O(n log n) due to sorting
// Space complexity: O(k) where k is the number of unique consecutive numbers
func FindConsecutiveNumbersOptimized(logs []ConsecutiveLog) []int {
	if len(logs) < 3 {
		return []int{}
	}

	// Create a map of ID to Num for O(1) lookup
	idToNum := make(map[int]int)
	idSet := make(map[int]bool)

	for _, log := range logs {
		idToNum[log.ID] = log.Num
		idSet[log.ID] = true
	}

	// Track found consecutive numbers
	consecutiveSet := make(map[int]bool)

	// For each ID, check if the next two consecutive IDs have the same number
	for id := range idSet {
		// Check if all three consecutive IDs exist
		if num, ok1 := idToNum[id]; ok1 {
			if num2, ok2 := idToNum[id+1]; ok2 && num == num2 {
				if num3, ok3 := idToNum[id+2]; ok3 && num2 == num3 {
					consecutiveSet[num] = true
				}
			}
		}
	}

	// Convert set to slice
	result := make([]int, 0, len(consecutiveSet))
	for num := range consecutiveSet {
		result = append(result, num)
	}

	return result
}

// FindConsecutiveNumbersStreamingApproach finds consecutive numbers using a
// streaming approach that only keeps track of the last 2 elements.
// More memory efficient for very large datasets.
//
// Time complexity: O(n)
// Space complexity: O(1) excluding output
func FindConsecutiveNumbersStreamingApproach(logs []ConsecutiveLog) []int {
	if len(logs) < 3 {
		return []int{}
	}

	consecutiveSet := make(map[int]bool)
	
	// Keep track of the last two logs
	prev2 := logs[0]
	prev1 := logs[1]

	for i := 2; i < len(logs); i++ {
		current := logs[i]

		// Check if current forms a consecutive triplet with prev1 and prev2
		if prev2.ID+1 == prev1.ID &&
			prev1.ID+1 == current.ID &&
			prev2.Num == prev1.Num &&
			prev1.Num == current.Num {
			consecutiveSet[prev2.Num] = true
		}

		// Shift the window
		prev2 = prev1
		prev1 = current
	}

	// Convert set to slice
	result := make([]int, 0, len(consecutiveSet))
	for num := range consecutiveSet {
		result = append(result, num)
	}

	return result
}
