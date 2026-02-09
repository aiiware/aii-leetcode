package arrays

// DeleteDuplicateEmails solves LeetCode problem 0196: Delete Duplicate Emails
// Difficulty: Easy
// Tags: Array, Database
//
// Write a SQL query to delete all duplicate email entries in a table named Person,
// keeping only unique emails based on its smallest Id.
//
// Time complexity: O(n), Space complexity: O(1)
func DeleteDuplicateEmails(emails []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, email := range emails {
		if !seen[email] {
			seen[email] = true
			result = append(result, email)
		}
	}

	return result
}
