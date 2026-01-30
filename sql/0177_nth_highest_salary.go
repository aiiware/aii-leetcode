package sql

/*
0177. Nth Highest Salary

Table: Employee
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
id is the primary key column for this table.
Each row of this table contains information about the salary of an employee.

Write an SQL query to report the nth highest distinct salary from the Employee table.
If there is no nth highest salary, the query should report null.

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
n = 2
Output: 
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+

Example 2:
Input: 
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
n = 2
Output: 
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| null                   |
+------------------------+

Constraints:
- 1 <= Employee.length <= 10^4
- -10^4 <= salary <= 10^4
- 1 <= n <= Employee.length

Difficulty: Medium
Tags: Database
Companies: Amazon, Google, Microsoft, Facebook, Apple
*/

// Employee177 represents an employee in the Employee table for problem 177
// Note: Using Employee177 to avoid conflict with Employee in 0176_second_highest_salary.go
type Employee177 struct {
	ID     int
	Salary int
}

// NthHighestSalaryResult represents the result of the query
type NthHighestSalaryResult struct {
	NthHighestSalary *int // Use pointer to allow null
}

// NthHighestSalary finds the nth highest distinct salary
// Returns nil if there is no nth highest salary
// n is 1-indexed (1 = highest salary, 2 = second highest, etc.)
func NthHighestSalary(employees []Employee177, n int) *NthHighestSalaryResult {
	if len(employees) == 0 || n <= 0 {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
	}

	// Find unique salaries using a map
	salarySet := make(map[int]bool)
	for _, emp := range employees {
		salarySet[emp.Salary] = true
	}

	// Convert to slice
	uniqueSalaries := make([]int, 0, len(salarySet))
	for salary := range salarySet {
		uniqueSalaries = append(uniqueSalaries, salary)
	}

	// If n is greater than number of unique salaries, return nil
	if n > len(uniqueSalaries) {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
	}

	// Sort in descending order
	for i := 0; i < len(uniqueSalaries); i++ {
		for j := i + 1; j < len(uniqueSalaries); j++ {
			if uniqueSalaries[i] < uniqueSalaries[j] {
				uniqueSalaries[i], uniqueSalaries[j] = uniqueSalaries[j], uniqueSalaries[i]
			}
		}
	}

	// Return the nth highest salary (n-1 because n is 1-indexed)
	return &NthHighestSalaryResult{NthHighestSalary: &uniqueSalaries[n-1]}
}

// NthHighestSalaryOptimized provides an optimized version using quickselect
// This approach finds the nth highest without fully sorting the array
func NthHighestSalaryOptimized(employees []Employee177, n int) *NthHighestSalaryResult {
	if len(employees) == 0 || n <= 0 {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
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

	// If n is greater than number of unique salaries, return nil
	if n > len(uniqueSalaries) {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
	}

	// Use quickselect to find the nth highest element
	// We want the (n-1)th element in descending order
	result := quickSelect(uniqueSalaries, 0, len(uniqueSalaries)-1, n-1)
	return &NthHighestSalaryResult{NthHighestSalary: &result}
}

// quickSelect finds the kth largest element (0-indexed) using quickselect algorithm
func quickSelect(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	pivotIndex := partition(arr, left, right)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect(arr, left, pivotIndex-1, k)
	} else {
		return quickSelect(arr, pivotIndex+1, right, k)
	}
}

// partition rearranges elements so that elements greater than pivot come before it
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		// Sort in descending order (larger elements first)
		if arr[j] > pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

// NthHighestSalarySQLStyle simulates the SQL approach using LIMIT and OFFSET
// This is the most direct translation of the SQL solution
func NthHighestSalarySQLStyle(employees []Employee177, n int) *NthHighestSalaryResult {
	if len(employees) == 0 || n <= 0 {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
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

	// If n is greater than number of unique salaries, return nil
	if n > len(uniqueSalaries) {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
	}

	// Sort in descending order (simulating ORDER BY salary DESC)
	for i := 0; i < len(uniqueSalaries); i++ {
		for j := i + 1; j < len(uniqueSalaries); j++ {
			if uniqueSalaries[i] < uniqueSalaries[j] {
				uniqueSalaries[i], uniqueSalaries[j] = uniqueSalaries[j], uniqueSalaries[i]
			}
		}
	}

	// Simulate LIMIT 1 OFFSET (n-1)
	return &NthHighestSalaryResult{NthHighestSalary: &uniqueSalaries[n-1]}
}

// NthHighestSalaryWithHeap uses a min-heap to find the nth highest salary
// More efficient for very large n values close to the total count
func NthHighestSalaryWithHeap(employees []Employee177, n int) *NthHighestSalaryResult {
	if len(employees) == 0 || n <= 0 {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
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

	// If n is greater than number of unique salaries, return nil
	if n > len(uniqueSalaries) {
		return &NthHighestSalaryResult{NthHighestSalary: nil}
	}

	// Build a min-heap of size n
	heap := make([]int, n)
	copy(heap, uniqueSalaries[:n])
	buildMinHeap(heap)

	// Process remaining salaries
	for i := n; i < len(uniqueSalaries); i++ {
		if uniqueSalaries[i] > heap[0] {
			heap[0] = uniqueSalaries[i]
			minHeapify(heap, 0, n)
		}
	}

	// The root of the min-heap is the nth largest element
	return &NthHighestSalaryResult{NthHighestSalary: &heap[0]}
}

// buildMinHeap builds a min-heap from an array
func buildMinHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(arr, i, n)
	}
}

// minHeapify maintains the min-heap property
func minHeapify(arr []int, i, n int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, smallest, n)
	}
}