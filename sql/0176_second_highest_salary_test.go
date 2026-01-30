package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondHighestSalary(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee
		expected  *int
	}{
		{
			name: "Example 1 from LeetCode",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			expected: intPtr176(200),
		},
		{
			name: "Example 2 from LeetCode - single employee",
			employees: []Employee{
				{ID: 1, Salary: 100},
			},
			expected: nil,
		},
		{
			name: "All same salaries - no second highest",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 100},
				{ID: 3, Salary: 100},
			},
			expected: nil,
		},
		{
			name: "Multiple employees with duplicate salaries",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
				{ID: 5, Salary: 300},
			},
			expected: intPtr176(200), // Second highest distinct salary
		},
		{
			name: "Negative salaries",
			employees: []Employee{
				{ID: 1, Salary: -100},
				{ID: 2, Salary: -200},
				{ID: 3, Salary: -50},
			},
			expected: intPtr176(-100), // -100 is second highest (after -50)
		},
		{
			name: "Mixed positive and negative salaries",
			employees: []Employee{
				{ID: 1, Salary: -100},
				{ID: 2, Salary: 0},
				{ID: 3, Salary: 100},
				{ID: 4, Salary: 200},
			},
			expected: intPtr176(100),
		},
		{
			name: "Empty employee list",
			employees: []Employee{},
			expected: nil,
		},
		{
			name: "Only two distinct salaries",
			employees: []Employee{
				{ID: 1, Salary: 500},
				{ID: 2, Salary: 500},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 300},
			},
			expected: intPtr176(300),
		},
		{
			name: "Large salary range",
			employees: []Employee{
				{ID: 1, Salary: 10000},
				{ID: 2, Salary: 5000},
				{ID: 3, Salary: 7500},
				{ID: 4, Salary: 2500},
			},
			expected: intPtr176(7500),
		},
		{
			name: "Second highest is first element in unsorted list",
			employees: []Employee{
				{ID: 1, Salary: 500},  // This is second highest
				{ID: 2, Salary: 1000}, // This is highest
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 200},
			},
			expected: intPtr176(500),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SecondHighestSalary(tt.employees)
			assertSecondHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestSecondHighestSalaryOptimized(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee
		expected  *int
	}{
		{
			name: "Basic example",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			expected: intPtr176(200),
		},
		{
			name: "Single employee",
			employees: []Employee{
				{ID: 1, Salary: 100},
			},
			expected: nil,
		},
		{
			name: "All duplicates",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 100},
				{ID: 3, Salary: 100},
			},
			expected: nil,
		},
		{
			name: "With duplicates",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
			},
			expected: intPtr176(200),
		},
		{
			name: "Negative values",
			employees: []Employee{
				{ID: 1, Salary: -300},
				{ID: 2, Salary: -200},
				{ID: 3, Salary: -100},
			},
			expected: intPtr176(-200),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SecondHighestSalaryOptimized(tt.employees)
			assertSecondHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestSecondHighestSalarySQLStyle(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee
		expected  *int
	}{
		{
			name: "Basic example",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			expected: intPtr176(200),
		},
		{
			name: "Single employee",
			employees: []Employee{
				{ID: 1, Salary: 100},
			},
			expected: nil,
		},
		{
			name: "With duplicates",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
			},
			expected: intPtr176(200),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SecondHighestSalarySQLStyle(tt.employees)
			assertSecondHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestSecondHighestSalary_Consistency(t *testing.T) {
	// Test that all implementations produce the same results
	testCases := []struct {
		name      string
		employees []Employee
	}{
		{
			name: "Normal case",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 400},
			},
		},
		{
			name: "With duplicates",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
				{ID: 5, Salary: 300},
			},
		},
		{
			name: "Single employee",
			employees: []Employee{
				{ID: 1, Salary: 100},
			},
		},
		{
			name: "All same salaries",
			employees: []Employee{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 100},
				{ID: 3, Salary: 100},
			},
		},
		{
			name: "Empty list",
			employees: []Employee{},
		},
		{
			name: "Negative salaries",
			employees: []Employee{
				{ID: 1, Salary: -100},
				{ID: 2, Salary: -200},
				{ID: 3, Salary: -50},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result1 := SecondHighestSalary(tc.employees)
			result2 := SecondHighestSalaryOptimized(tc.employees)
			result3 := SecondHighestSalarySQLStyle(tc.employees)

			// All should have same nil status
			if result1.SecondHighestSalary == nil {
				assert.Nil(t, result2.SecondHighestSalary, "Optimized should also return nil")
				assert.Nil(t, result3.SecondHighestSalary, "SQLStyle should also return nil")
			} else {
				assert.NotNil(t, result2.SecondHighestSalary, "Optimized should not return nil")
				assert.NotNil(t, result3.SecondHighestSalary, "SQLStyle should not return nil")
				if result2.SecondHighestSalary != nil && result3.SecondHighestSalary != nil {
					assert.Equal(t, *result1.SecondHighestSalary, *result2.SecondHighestSalary, "Basic and Optimized should match")
					assert.Equal(t, *result1.SecondHighestSalary, *result3.SecondHighestSalary, "Basic and SQLStyle should match")
				}
			}
		})
	}
}

func BenchmarkSecondHighestSalary(b *testing.B) {
	// Create test data with 10000 employees (max constraint)
	employees := make([]Employee, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee{
			ID:     i + 1,
			Salary: (i * 7) % 5000, // Generate varied salaries with some duplicates
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SecondHighestSalary(employees)
	}
}

func BenchmarkSecondHighestSalaryOptimized(b *testing.B) {
	// Create test data with 10000 employees
	employees := make([]Employee, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee{
			ID:     i + 1,
			Salary: (i * 7) % 5000,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SecondHighestSalaryOptimized(employees)
	}
}

func BenchmarkSecondHighestSalarySQLStyle(b *testing.B) {
	// Create test data with 10000 employees
	employees := make([]Employee, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee{
			ID:     i + 1,
			Salary: (i * 7) % 5000,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SecondHighestSalarySQLStyle(employees)
	}
}

// Helper function to assert SecondHighestSalaryResult
func assertSecondHighestSalaryResult(t *testing.T, result *SecondHighestSalaryResult, expected *int) {
	if expected == nil {
		assert.Nil(t, result.SecondHighestSalary, "Expected nil SecondHighestSalary")
	} else {
		assert.NotNil(t, result.SecondHighestSalary, "Expected non-nil SecondHighestSalary")
		if result.SecondHighestSalary != nil {
			assert.Equal(t, *expected, *result.SecondHighestSalary, "SecondHighestSalary value mismatch")
		}
	}
}

// Helper function to create int pointers
func intPtr176(i int) *int {
	return &i
}