package sql

/*
182. Duplicate Emails

Table: Person

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| email       | varchar |
+-------------+---------+
id is the primary key for this table.
Each row of this table contains an email. The emails will not contain uppercase letters.

Write an SQL query to report all the duplicate emails. Note that it's guaranteed that the email field is not NULL.

Return the result table in any order.

Example 1:
Input: 
Person table:
+----+---------+
| id | email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
Output: 
+---------+
| Email   |
+---------+
| a@b.com |
+---------+
Explanation: a@b.com is repeated two times.

Difficulty: Easy
Tags: Database
Companies: Amazon, Google, Microsoft
*/

// Person182 represents a row in the Person table for problem 182
type Person182 struct {
	ID    int
	Email string
}

// DuplicateEmails finds all duplicate emails in the Person table.
// This is the Go equivalent of the SQL solution for LeetCode 0182.
//
// Algorithm:
// 1. Count frequency of each email using a map
// 2. Collect emails with count > 1
//
// Time complexity: O(n) where n is the number of persons
// Space complexity: O(k) where k is the number of unique emails
func DuplicateEmails(persons []Person182) []string {
	// Count frequency of each email
	emailCount := make(map[string]int)
	for _, person := range persons {
		emailCount[person.Email]++
	}

	// Collect emails with count > 1
	result := []string{}
	for email, count := range emailCount {
		if count > 1 {
			result = append(result, email)
		}
	}

	return result
}