package leetcode

import (
	"slices"
)

/*
185. Department Top Three Salaries

Table: Employee

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| id           | int     |
| name         | varchar |
| salary       | int     |
| departmentId | int     |
+--------------+---------+
id is the primary key for this table.
departmentId is a foreign key of the ID from the Department table.
Each row of this table indicates the ID, name, and salary of an employee. It also contains the ID of their department.

Table: Department

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
id is the primary key for this table.
Each row of this table indicates the ID and the name of a department.

Write an SQL query to find the employees who are high earners in each of the departments.
A high earner in a department is an employee who has a salary in the top three unique salaries for that department.

Return the result table in any order.

Example 1:
Input: 
Employee table:
+----+-------+--------+--------------+
| id | name  | salary | departmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 85000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
| 7  | Will  | 70000  | 1            |
+----+-------+--------+--------------+
Department table:
+----+-------+
| id | name  |
+----+-------+
| 1  | IT    |
| 2  | Sales |
+----+-------+
Output: 
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Joe      | 85000  |
| IT         | Randy    | 85000  |
| IT         | Will     | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
+------------+----------+--------+
Explanation: 
In the IT department:
- Max earns the highest unique salary
- Both Randy and Joe earn the second-highest unique salary
- Will earns the third-highest unique salary

In the Sales department:
- Henry earns the highest salary
- Sam earns the second-highest salary
There is no third-highest salary as there are only two employees.

Difficulty: Hard
Tags: Database
Companies: Amazon, Google, Microsoft
*/

// Employee185 represents a row in the Employee table for problem 185
type Employee185 struct {
	ID           int
	Name         string
	Salary       int
	DepartmentID int
}

// Department185 represents a row in the Department table for problem 185
type Department185 struct {
	ID   int
	Name string
}

// DepartmentSalary185 represents the result structure for problem 185
type DepartmentSalary185 struct {
	Department string
	Employee   string
	Salary     int
}

// DepartmentTopThreeSalaries finds employees with top three unique salaries in each department.
// This is the Go equivalent of the SQL solution for LeetCode 0185.
//
// Algorithm:
// 1. Group employees by department
// 2. For each department, get unique salaries and sort in descending order
// 3. Take top 3 unique salaries
// 4. Find all employees with those salaries
//
// Time complexity: O(n log n) for sorting salaries in each department
// Space complexity: O(n) for storing grouped employees
func DepartmentTopThreeSalaries(employees []Employee185, departments []Department185) []DepartmentSalary185 {
	// Create a map of department ID to department name
	deptName := make(map[int]string)
	for _, dept := range departments {
		deptName[dept.ID] = dept.Name
	}

	// Group employees by department
	employeesByDept := make(map[int][]Employee185)
	for _, emp := range employees {
		employeesByDept[emp.DepartmentID] = append(employeesByDept[emp.DepartmentID], emp)
	}

	result := []DepartmentSalary185{}

	// Process each department
	for deptID, deptEmployees := range employeesByDept {
		// Get unique salaries for this department
		salarySet := make(map[int]bool)
		for _, emp := range deptEmployees {
			salarySet[emp.Salary] = true
		}

		// Convert to slice and sort in descending order
		uniqueSalaries := make([]int, 0, len(salarySet))
		for salary := range salarySet {
			uniqueSalaries = append(uniqueSalaries, salary)
		}
		slices.SortFunc(uniqueSalaries, func(a, b int) int {
			return b - a // descending order
		})

		// Take top 3 unique salaries
		topSalaries := []int{}
		for i := 0; i < len(uniqueSalaries) && i < 3; i++ {
			topSalaries = append(topSalaries, uniqueSalaries[i])
		}

		// Create a set for O(1) lookup
		topSalarySet := make(map[int]bool)
		for _, salary := range topSalaries {
			topSalarySet[salary] = true
		}

		// Find employees with top salaries
		for _, emp := range deptEmployees {
			if topSalarySet[emp.Salary] {
				result = append(result, DepartmentSalary185{
					Department: deptName[deptID],
					Employee:   emp.Name,
					Salary:     emp.Salary,
				})
			}
		}
	}

	return result
}