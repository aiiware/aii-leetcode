package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomersWhoNeverOrder(t *testing.T) {
	tests := []struct {
		name      string
		customers []Customer183
		orders    []Order183
		expected  []string
	}{
		{
			name: "Example 1",
			customers: []Customer183{
				{ID: 1, Name: "Joe"},
				{ID: 2, Name: "Henry"},
				{ID: 3, Name: "Sam"},
				{ID: 4, Name: "Max"},
			},
			orders: []Order183{
				{ID: 1, CustomerID: 3},
				{ID: 2, CustomerID: 1},
			},
			expected: []string{"Henry", "Max"},
		},
		{
			name: "All customers have orders",
			customers: []Customer183{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
			},
			orders: []Order183{
				{ID: 1, CustomerID: 1},
				{ID: 2, CustomerID: 2},
			},
			expected: []string{},
		},
		{
			name: "No orders at all",
			customers: []Customer183{
				{ID: 1, Name: "Customer1"},
				{ID: 2, Name: "Customer2"},
				{ID: 3, Name: "Customer3"},
			},
			orders:   []Order183{},
			expected: []string{"Customer1", "Customer2", "Customer3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CustomersWhoNeverOrder(tt.customers, tt.orders)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}