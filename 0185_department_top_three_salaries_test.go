package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepartmentTopThreeSalaries(t *testing.T) {
	tests := []struct {
		name       string
		employees  []Employee185
		departments []Department185
		expected   []DepartmentSalary185
	}{
		{
			name: "Example 1",
			employees: []Employee185{
				{ID: 1, Name: "Joe", Salary: 85000, DepartmentID: 1},
				{ID: 2, Name: "Henry", Salary: 80000, DepartmentID: 2},
				{ID: 3, Name: "Sam", Salary: 60000, DepartmentID: 2},
				{ID: 4, Name: "Max", Salary: 90000, DepartmentID: 1},
				{ID: 5, Name: "Janet", Salary: 69000, DepartmentID: 1},
				{ID: 6, Name: "Randy", Salary: 85000, DepartmentID: 1},
				{ID: 7, Name: "Will", Salary: 70000, DepartmentID: 1},
			},
			departments: []Department185{
				{ID: 1, Name: "IT"},
				{ID: 2, Name: "Sales"},
			},
			expected: []DepartmentSalary185{
				{Department: "IT", Employee: "Max", Salary: 90000},
				{Department: "IT", Employee: "Joe", Salary: 85000},
				{Department: "IT", Employee: "Randy", Salary: 85000},
				{Department: "IT", Employee: "Will", Salary: 70000},
				{Department: "Sales", Employee: "Henry", Salary: 80000},
				{Department: "Sales", Employee: "Sam", Salary: 60000},
			},
		},
		{
			name: "Less than 3 employees in department",
			employees: []Employee185{
				{ID: 1, Name: "Alice", Salary: 50000, DepartmentID: 1},
				{ID: 2, Name: "Bob", Salary: 45000, DepartmentID: 1},
			},
			departments: []Department185{
				{ID: 1, Name: "HR"},
			},
			expected: []DepartmentSalary185{
				{Department: "HR", Employee: "Alice", Salary: 50000},
				{Department: "HR", Employee: "Bob", Salary: 45000},
			},
		},
		{
			name: "Ties in salaries - all employees with top 2 unique salaries should be included",
			employees: []Employee185{
				{ID: 1, Name: "E1", Salary: 100000, DepartmentID: 1},
				{ID: 2, Name: "E2", Salary: 100000, DepartmentID: 1},
				{ID: 3, Name: "E3", Salary: 100000, DepartmentID: 1},
				{ID: 4, Name: "E4", Salary: 90000, DepartmentID: 1},
			},
			departments: []Department185{
				{ID: 1, Name: "Engineering"},
			},
			expected: []DepartmentSalary185{
				{Department: "Engineering", Employee: "E1", Salary: 100000},
				{Department: "Engineering", Employee: "E2", Salary: 100000},
				{Department: "Engineering", Employee: "E3", Salary: 100000},
				{Department: "Engineering", Employee: "E4", Salary: 90000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DepartmentTopThreeSalaries(tt.employees, tt.departments)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}