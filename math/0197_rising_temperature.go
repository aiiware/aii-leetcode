package math

// RisingTemperature solves LeetCode problem 0197: Rising Temperature
// Difficulty: Easy
// Tags: Database, Array
//
// Given a Weather table, write a SQL query to find all dates' Ids with higher temperature compared to its previous dates (yesterday).
//
// Time complexity: O(n), Space complexity: O(1)
func RisingTemperature(temperatures []int) []int {
	var result []int
	
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] > temperatures[i-1] {
			result = append(result, i) // Return index (representing day Id)
		}
	}
	
	return result
}
