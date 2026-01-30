package sql

/*
0176. Second Highest Salary

Table: Employee
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
id is the primary key column for this table.
Each row of this table contains information about the salary of an employee.

Write an SQL query to report the second highest distinct salary from the Employee table.
If there is no second highest salary, the query should report null.

Example 1:
Input: 
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
Output: 
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+

Example 2:
Input: 
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
Output: 
+---------------------+
| SecondHighestSalary |
+---------------------+
| null                |
+---------------------+

Constraints:
- 1 <= Employee.length <= 10^4
- -10^4 <= salary <= 10^4

Difficulty: Medium
Tags: Database
Companies: Amazon, Google, Microsoft, Apple, Facebook
*/

// Employee represents an employee in the Employee table
type Employee struct {
	ID     int
	Salary int
}

// SecondHighestSalaryResult represents the result of the query
type SecondHighestSalaryResult struct {
	SecondHighestSalary *int // Use pointer to allow null
}

// SecondHighestSalary finds the second highest distinct salary
// Returns nil if there is no second highest salary
func SecondHighestSalary(employees []Employee) *SecondHighestSalaryResult {
	if len(employees) == 0 {
		return &SecondHighestSalaryResult{SecondHighestSalary: nil}
	}

	// Find unique salaries using a map
	salarySet := make(map[int]bool)
	for _, emp := range employees {
		salarySet[emp.Salary] = true
	}

	// Convert to slice and sort in descending order
	uniqueSalaries := make([]int, 0, len(salarySet))
	for salary := range salarySet {
		uniqueSalaries = append(uniqueSalaries, salary)
	}

	// Sort in descending order
	for i := 0; i < len(uniqueSalaries); i++ {
		for j := i + 1; j < len(uniqueSalaries); j++ {
			if uniqueSalaries[i] < uniqueSalaries[j] {
				uniqueSalaries[i], uniqueSalaries[j] = uniqueSalaries[j], uniqueSalaries[i]
			}
		}
	}

	// If we have at least 2 unique salaries, return the second one
	if len(uniqueSalaries) >= 2 {
		return &SecondHighestSalaryResult{SecondHighestSalary: &uniqueSalaries[1]}
	}

	// Otherwise return nil
	return &SecondHighestSalaryResult{SecondHighestSalary: nil}
}

// SecondHighestSalaryOptimized provides an optimized version using a single pass
// This approach finds the max and second max in one pass
func SecondHighestSalaryOptimized(employees []Employee) *SecondHighestSalaryResult {
	if len(employees) == 0 {
		return &SecondHighestSalaryResult{SecondHighestSalary: nil}
	}

	// Initialize with smallest possible values
	maxSalary := employees[0].Salary
	secondMax := employees[0].Salary
	foundSecond := false

	// Track unique salaries to handle duplicates
	seenSalaries := make(map[int]bool)

	for _, emp := range employees {
		salary := emp.Salary

		// Skip if we've seen this salary before
		if seenSalaries[salary] {
			continue
		}
		seenSalaries[salary] = true

		if salary > maxSalary {
			// Current max becomes second max, new salary becomes max
			secondMax = maxSalary
			maxSalary = salary
			foundSecond = true
		} else if salary > secondMax && salary < maxSalary {
			// Salary is between current second max and max
			secondMax = salary
			foundSecond = true
		} else if !foundSecond && salary < maxSalary {
			// We have at least 2 unique salaries now
			secondMax = salary
			foundSecond = true
		}
	}

	// If we found a second max salary, return it
	if foundSecond && secondMax != maxSalary {
		return &SecondHighestSalaryResult{SecondHighestSalary: &secondMax}
	}

	// Otherwise return nil
	return &SecondHighestSalaryResult{SecondHighestSalary: nil}
}

// SecondHighestSalarySQLStyle simulates the SQL approach using LIMIT and OFFSET
// This is the most direct translation of the SQL solution
func SecondHighestSalarySQLStyle(employees []Employee) *SecondHighestSalaryResult {
	if len(employees) == 0 {
		return &SecondHighestSalaryResult{SecondHighestSalary: nil}
	}

	// Get unique salaries
	salarySet := make(map[int]bool)
	for _, emp := range employees {
		salarySet[emp.Salary] = true
	}

	// Convert to slice
	uniqueSalaries := make([]int, 0, len(salarySet))
	for salary := range salarySet {
		uniqueSalaries = append(uniqueSalaries, salary)
	}

	// Sort in descending order (simulating ORDER BY salary DESC)
	for i := 0; i < len(uniqueSalaries); i++ {
		for j := i + 1; j < len(uniqueSalaries); j++ {
			if uniqueSalaries[i] < uniqueSalaries[j] {
				uniqueSalaries[i], uniqueSalaries[j] = uniqueSalaries[j], uniqueSalaries[i]
			}
		}
	}

	// Simulate LIMIT 1 OFFSET 1
	if len(uniqueSalaries) >= 2 {
		return &SecondHighestSalaryResult{SecondHighestSalary: &uniqueSalaries[1]}
	}

	return &SecondHighestSalaryResult{SecondHighestSalary: nil}
}