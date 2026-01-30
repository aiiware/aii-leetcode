package leetcode

/*
181. Employees Earning More Than Their Managers

Table: Employee

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
| salary      | int     |
| managerId   | int     |
+-------------+---------+
id is the primary key for this table.
Each row of this table indicates the ID of an employee, their name, salary, and the ID of their manager.

Write an SQL query to find the employees who earn more than their managers.

Return the result table in any order.

Example 1:
Input: 
Employee table:
+----+-------+--------+-----------+
| id | name  | salary | managerId |
+----+-------+--------+-----------+
| 1  | Joe   | 70000  | 3         |
| 2  | Henry | 80000  | 4         |
| 3  | Sam   | 60000  | Null      |
| 4  | Max   | 90000  | Null      |
+----+-------+--------+-----------+
Output: 
+----------+
| Employee |
+----------+
| Joe      |
+----------+
Explanation: Joe is the only employee who earns more than his manager.

Difficulty: Easy
Tags: Database
Companies: Amazon, Google, Microsoft
*/

// Employee181 represents a row in the Employee table for problem 181
type Employee181 struct {
	ID        int
	Name      string
	Salary    int
	ManagerID *int
}

// EmployeesEarningMoreThanTheirManagers finds employees who earn more than their managers.
// This is the Go equivalent of the SQL solution for LeetCode 0181.
//
// Algorithm:
// 1. Create a map of manager IDs to their salaries
// 2. Iterate through employees and check if they have a manager
// 3. If they have a manager and their salary > manager's salary, add to result
//
// Time complexity: O(n) where n is the number of employees
// Space complexity: O(m) where m is the number of managers
func EmployeesEarningMoreThanTheirManagers(employees []Employee181) []string {
	// Create a map of manager IDs to their salaries
	managerSalaries := make(map[int]int)
	for _, emp := range employees {
		if emp.ManagerID == nil {
			// This employee is a manager (no manager themselves)
			managerSalaries[emp.ID] = emp.Salary
		}
	}

	// Find employees who earn more than their managers
	result := []string{}
	for _, emp := range employees {
		if emp.ManagerID != nil {
			if managerSalary, exists := managerSalaries[*emp.ManagerID]; exists {
				if emp.Salary > managerSalary {
					result = append(result, emp.Name)
				}
			}
		}
	}

	return result
}