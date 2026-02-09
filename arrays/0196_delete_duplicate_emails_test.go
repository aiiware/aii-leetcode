package arrays

import "testing"

func TestDeleteDuplicateEmails(t *testing.T) {
	tests := []struct {
		name     string
		emails   []string
		expected []string
	}{
		{
			name:     "no duplicates",
			emails:   []string{"a@example.com", "b@example.com", "c@example.com"},
			expected: []string{"a@example.com", "b@example.com", "c@example.com"},
		},
		{
			name:     "with duplicates",
			emails:   []string{"a@example.com", "b@example.com", "a@example.com", "c@example.com", "b@example.com"},
			expected: []string{"a@example.com", "b@example.com", "c@example.com"},
		},
		{
			name:     "all same",
			emails:   []string{"test@example.com", "test@example.com", "test@example.com"},
			expected: []string{"test@example.com"},
		},
		{
			name:     "empty",
			emails:   []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeleteDuplicateEmails(tt.emails)
			if len(result) != len(tt.expected) {
				t.Errorf("DeleteDuplicateEmails() = %v, want %v", result, tt.expected)
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("DeleteDuplicateEmails()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func BenchmarkDeleteDuplicateEmails(b *testing.B) {
	emails := []string{
		"a@example.com",
		"b@example.com",
		"a@example.com",
		"c@example.com",
		"b@example.com",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeleteDuplicateEmails(emails)
	}
}
