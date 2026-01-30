package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateEmails(t *testing.T) {
	tests := []struct {
		name     string
		persons  []Person182
		expected []string
	}{
		{
			name: "Example 1",
			persons: []Person182{
				{ID: 1, Email: "a@b.com"},
				{ID: 2, Email: "c@d.com"},
				{ID: 3, Email: "a@b.com"},
			},
			expected: []string{"a@b.com"},
		},
		{
			name: "Multiple duplicates",
			persons: []Person182{
				{ID: 1, Email: "test1@example.com"},
				{ID: 2, Email: "test2@example.com"},
				{ID: 3, Email: "test1@example.com"},
				{ID: 4, Email: "test2@example.com"},
				{ID: 5, Email: "test3@example.com"},
			},
			expected: []string{"test1@example.com", "test2@example.com"},
		},
		{
			name: "No duplicates",
			persons: []Person182{
				{ID: 1, Email: "unique1@example.com"},
				{ID: 2, Email: "unique2@example.com"},
				{ID: 3, Email: "unique3@example.com"},
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DuplicateEmails(tt.persons)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}