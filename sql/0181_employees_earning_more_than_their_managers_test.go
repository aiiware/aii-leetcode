package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
    "leetcode/utils"
)

func TestEmployeesEarningMoreThanTheirManagers(t *testing.T) {
	tests := []struct {
		name     string
		employees []Employee181
		expected []string
	}{
		{
			name: "Example 1",
			employees: []Employee181{
				{ID: 1, Name: "Joe", Salary: 70000, ManagerID: utils.IntPtr(3)},
				{ID: 2, Name: "Henry", Salary: 80000, ManagerID: utils.IntPtr(4)},
				{ID: 3, Name: "Sam", Salary: 60000, ManagerID: nil},
				{ID: 4, Name: "Max", Salary: 90000, ManagerID: nil},
			},
			expected: []string{"Joe"},
		},
		{
			name: "No employees earn more than managers",
			employees: []Employee181{
				{ID: 1, Name: "Alice", Salary: 50000, ManagerID: utils.IntPtr(3)},
				{ID: 2, Name: "Bob", Salary: 60000, ManagerID: utils.IntPtr(4)},
				{ID: 3, Name: "Charlie", Salary: 70000, ManagerID: nil},
				{ID: 4, Name: "David", Salary: 80000, ManagerID: nil},
			},
			expected: []string{},
		},
		{
			name: "Multiple employees earn more than managers",
			employees: []Employee181{
				{ID: 1, Name: "Employee1", Salary: 100000, ManagerID: utils.IntPtr(3)},
				{ID: 2, Name: "Employee2", Salary: 120000, ManagerID: utils.IntPtr(4)},
				{ID: 3, Name: "Manager1", Salary: 90000, ManagerID: nil},
				{ID: 4, Name: "Manager2", Salary: 110000, ManagerID: nil},
			},
			expected: []string{"Employee1", "Employee2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EmployeesEarningMoreThanTheirManagers(tt.employees)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}