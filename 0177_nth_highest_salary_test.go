package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthHighestSalary(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee177
		n         int
		expected  *int
	}{
		{
			name: "Example 1 from LeetCode - 2nd highest",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200),
		},
		{
			name: "Example 2 from LeetCode - 2nd highest with single employee",
			employees: []Employee177{
				{ID: 1, Salary: 100},
			},
			n:        2,
			expected: nil,
		},
		{
			name: "1st highest salary",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        1,
			expected: intPtr177(300),
		},
		{
			name: "3rd highest salary",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 400},
			},
			n:        3,
			expected: intPtr177(200),
		},
		{
			name: "With duplicate salaries",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
				{ID: 5, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200), // Distinct salaries: [300, 200, 100]
		},
		{
			name: "n equals number of distinct salaries",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        3,
			expected: intPtr177(100),
		},
		{
			name: "n greater than number of distinct salaries",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        4,
			expected: nil,
		},
		{
			name: "All same salaries - asking for 2nd highest",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 100},
				{ID: 3, Salary: 100},
			},
			n:        2,
			expected: nil, // Only 1 distinct salary
		},
		{
			name: "Negative salaries",
			employees: []Employee177{
				{ID: 1, Salary: -100},
				{ID: 2, Salary: -200},
				{ID: 3, Salary: -50},
				{ID: 4, Salary: 0},
			},
			n:        2,
			expected: intPtr177(-50), // Distinct: [0, -50, -100, -200], 2nd highest is -50
		},
		{
			name:      "Empty employee list",
			employees: []Employee177{},
			n:         1,
			expected:  nil,
		},
		{
			name: "n = 0 (invalid)",
			employees: []Employee177{
				{ID: 1, Salary: 100},
			},
			n:        0,
			expected: nil,
		},
		{
			name: "Large n value",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 400},
				{ID: 5, Salary: 500},
			},
			n:        5,
			expected: intPtr177(100),
		},
		{
			name: "Mixed scenario with many duplicates",
			employees: []Employee177{
				{ID: 1, Salary: 500},
				{ID: 2, Salary: 500},
				{ID: 3, Salary: 400},
				{ID: 4, Salary: 400},
				{ID: 5, Salary: 300},
				{ID: 6, Salary: 300},
				{ID: 7, Salary: 200},
				{ID: 8, Salary: 100},
			},
			n:        3, // Distinct: [500, 400, 300, 200, 100]
			expected: intPtr177(300),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NthHighestSalary(tt.employees, tt.n)
			assertNthHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestNthHighestSalaryOptimized(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee177
		n         int
		expected  *int
	}{
		{
			name: "Basic example - 2nd highest",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200),
		},
		{
			name: "Quickselect with duplicates",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200),
		},
		{
			name: "n too large",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
			},
			n:        3,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NthHighestSalaryOptimized(tt.employees, tt.n)
			assertNthHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestNthHighestSalarySQLStyle(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee177
		n         int
		expected  *int
	}{
		{
			name: "Basic example",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200),
		},
		{
			name: "With offset calculation",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 400},
			},
			n:        3,
			expected: intPtr177(200),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NthHighestSalarySQLStyle(tt.employees, tt.n)
			assertNthHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestNthHighestSalaryWithHeap(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee177
		n         int
		expected  *int
	}{
		{
			name: "Basic heap example",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200),
		},
		{
			name: "Heap with many elements",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 400},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 500},
				{ID: 5, Salary: 200},
			},
			n:        3, // Distinct: [500, 400, 300, 200, 100]
			expected: intPtr177(300),
		},
		{
			name: "Heap with duplicates",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
				{ID: 5, Salary: 300},
			},
			n:        2,
			expected: intPtr177(200),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NthHighestSalaryWithHeap(tt.employees, tt.n)
			assertNthHighestSalaryResult(t, result, tt.expected)
		})
	}
}

func TestNthHighestSalary_Consistency(t *testing.T) {
	// Test that all implementations produce the same results
	testCases := []struct {
		name      string
		employees []Employee177
		n         int
	}{
		{
			name: "Normal case n=2",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
				{ID: 4, Salary: 400},
			},
			n: 2,
		},
		{
			name: "With duplicates n=3",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 200},
				{ID: 4, Salary: 300},
				{ID: 5, Salary: 300},
			},
			n: 3,
		},
		{
			name: "n=1 (highest)",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 500},
				{ID: 3, Salary: 300},
			},
			n: 1,
		},
		{
			name: "n equals distinct count",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
				{ID: 3, Salary: 300},
			},
			n: 3,
		},
		{
			name: "n too large",
			employees: []Employee177{
				{ID: 1, Salary: 100},
				{ID: 2, Salary: 200},
			},
			n: 5,
		},
		{
			name:      "Empty list",
			employees: []Employee177{},
			n:         1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result1 := NthHighestSalary(tc.employees, tc.n)
			result2 := NthHighestSalaryOptimized(tc.employees, tc.n)
			result3 := NthHighestSalarySQLStyle(tc.employees, tc.n)
			result4 := NthHighestSalaryWithHeap(tc.employees, tc.n)

			// All should have same nil status
			results := []*NthHighestSalaryResult{result1, result2, result3, result4}
			for i := 1; i < len(results); i++ {
				if result1.NthHighestSalary == nil {
					assert.Nil(t, results[i].NthHighestSalary, "Implementation %d should also return nil", i)
				} else {
					assert.NotNil(t, results[i].NthHighestSalary, "Implementation %d should not return nil", i)
					if results[i].NthHighestSalary != nil {
						assert.Equal(t, *result1.NthHighestSalary, *results[i].NthHighestSalary, 
							"Implementation %d should match basic implementation", i)
					}
				}
			}
		})
	}
}

func BenchmarkNthHighestSalary(b *testing.B) {
	// Create test data with 10000 employees
	employees := make([]Employee177, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee177{
			ID:     i + 1,
			Salary: (i * 13) % 5000, // Generate varied salaries
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NthHighestSalary(employees, 100) // Find 100th highest
	}
}

func BenchmarkNthHighestSalaryOptimized(b *testing.B) {
	employees := make([]Employee177, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee177{
			ID:     i + 1,
			Salary: (i * 13) % 5000,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NthHighestSalaryOptimized(employees, 100)
	}
}

func BenchmarkNthHighestSalarySQLStyle(b *testing.B) {
	employees := make([]Employee177, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee177{
			ID:     i + 1,
			Salary: (i * 13) % 5000,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NthHighestSalarySQLStyle(employees, 100)
	}
}

func BenchmarkNthHighestSalaryWithHeap(b *testing.B) {
	employees := make([]Employee177, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee177{
			ID:     i + 1,
			Salary: (i * 13) % 5000,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NthHighestSalaryWithHeap(employees, 100)
	}
}

func BenchmarkNthHighestSalary_VaryingN(b *testing.B) {
	employees := make([]Employee177, 10000)
	for i := 0; i < 10000; i++ {
		employees[i] = Employee177{
			ID:     i + 1,
			Salary: (i * 13) % 5000,
		}
	}

	// Benchmark different n values
	nValues := []int{1, 10, 100, 1000, 5000}
	for _, n := range nValues {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NthHighestSalaryOptimized(employees, n)
			}
		})
	}
}

// Helper function to assert NthHighestSalaryResult
func assertNthHighestSalaryResult(t *testing.T, result *NthHighestSalaryResult, expected *int) {
	if expected == nil {
		assert.Nil(t, result.NthHighestSalary, "Expected nil NthHighestSalary")
	} else {
		assert.NotNil(t, result.NthHighestSalary, "Expected non-nil NthHighestSalary")
		if result.NthHighestSalary != nil {
			assert.Equal(t, *expected, *result.NthHighestSalary, "NthHighestSalary value mismatch")
		}
	}
}

// Helper function to create int pointers
func intPtr177(i int) *int {
	return &i
}