package arrays

// TenthLine solves LeetCode problem 0195: Tenth Line
// Difficulty: Easy
// Tags: String, Text Processing
//
// Given a text file file.txt, print just the 10th line of the file.
//
// Time complexity: O(n), Space complexity: O(1)
func TenthLine(lines []string) string {
	if len(lines) >= 10 {
		return lines[9] // 0-indexed, so 10th line is at index 9
	}
	return ""
}
