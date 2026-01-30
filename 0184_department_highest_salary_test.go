package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepartmentHighestSalary(t *testing.T) {
	tests := []struct {
		name       string
		employees  []Employee184
		departments []Department184
		expected   []DepartmentSalary184
	}{
		{
			name: "Example 1",
			employees: []Employee184{
				{ID: 1, Name: "Joe", Salary: 70000, DepartmentID: 1},
				{ID: 2, Name: "Jim", Salary: 90000, DepartmentID: 1},
				{ID: 3, Name: "Henry", Salary: 80000, DepartmentID: 2},
				{ID: 4, Name: "Sam", Salary: 60000, DepartmentID: 2},
				{ID: 5, Name: "Max", Salary: 90000, DepartmentID: 1},
			},
			departments: []Department184{
				{ID: 1, Name: "IT"},
				{ID: 2, Name: "Sales"},
			},
			expected: []DepartmentSalary184{
				{Department: "IT", Employee: "Jim", Salary: 90000},
				{Department: "IT", Employee: "Max", Salary: 90000},
				{Department: "Sales", Employee: "Henry", Salary: 80000},
			},
		},
		{
			name: "Single department, single employee",
			employees: []Employee184{
				{ID: 1, Name: "Alice", Salary: 50000, DepartmentID: 1},
			},
			departments: []Department184{
				{ID: 1, Name: "HR"},
			},
			expected: []DepartmentSalary184{
				{Department: "HR", Employee: "Alice", Salary: 50000},
			},
		},
		{
			name: "Multiple employees with same salary",
			employees: []Employee184{
				{ID: 1, Name: "Bob", Salary: 60000, DepartmentID: 1},
				{ID: 2, Name: "Charlie", Salary: 60000, DepartmentID: 1},
				{ID: 3, Name: "David", Salary: 55000, DepartmentID: 1},
			},
			departments: []Department184{
				{ID: 1, Name: "Engineering"},
			},
			expected: []DepartmentSalary184{
				{Department: "Engineering", Employee: "Bob", Salary: 60000},
				{Department: "Engineering", Employee: "Charlie", Salary: 60000},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DepartmentHighestSalary(tt.employees, tt.departments)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}