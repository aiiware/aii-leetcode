package sql


/*
184. Department Highest Salary

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

Write an SQL query to find employees who have the highest salary in each of the departments.

Return the result table in any order.

Example 1:
Input: 
Employee table:
+----+-------+--------+--------------+
| id | name  | salary | departmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Jim   | 90000  | 1            |
| 3  | Henry | 80000  | 2            |
| 4  | Sam   | 60000  | 2            |
| 5  | Max   | 90000  | 1            |
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
| IT         | Jim      | 90000  |
| IT         | Max      | 90000  |
| Sales      | Henry    | 80000  |
+------------+----------+--------+
Explanation: Max and Jim both have the highest salary in the IT department and Henry has the highest salary in the Sales department.

Difficulty: Medium
Tags: Database
Companies: Amazon, Google, Microsoft
*/

// Employee184 represents a row in the Employee table for problem 184
type Employee184 struct {
	ID           int
	Name         string
	Salary       int
	DepartmentID int
}

// Department184 represents a row in the Department table for problem 184
type Department184 struct {
	ID   int
	Name string
}

// DepartmentSalary184 represents the result structure for problem 184
type DepartmentSalary184 struct {
	Department string
	Employee   string
	Salary     int
}

// DepartmentHighestSalary finds employees with the highest salary in each department.
// This is the Go equivalent of the SQL solution for LeetCode 0184.
//
// Algorithm:
// 1. Create a map of department ID to department name
// 2. Find the maximum salary for each department
// 3. Find all employees who have that maximum salary in their department
//
// Time complexity: O(n + m) where n is number of employees, m is number of departments
// Space complexity: O(m) for storing department info and max salaries
func DepartmentHighestSalary(employees []Employee184, departments []Department184) []DepartmentSalary184 {
	// Create a map of department ID to department name
	deptName := make(map[int]string)
	for _, dept := range departments {
		deptName[dept.ID] = dept.Name
	}

	// Find the maximum salary for each department
	maxSalaryByDept := make(map[int]int)
	for _, emp := range employees {
		if currentMax, exists := maxSalaryByDept[emp.DepartmentID]; !exists || emp.Salary > currentMax {
			maxSalaryByDept[emp.DepartmentID] = emp.Salary
		}
	}

	// Find employees who have the maximum salary in their department
	result := []DepartmentSalary184{}
	for _, emp := range employees {
		if maxSalary, exists := maxSalaryByDept[emp.DepartmentID]; exists && emp.Salary == maxSalary {
			result = append(result, DepartmentSalary184{
				Department: deptName[emp.DepartmentID],
				Employee:   emp.Name,
				Salary:     emp.Salary,
			})
		}
	}

	return result
}