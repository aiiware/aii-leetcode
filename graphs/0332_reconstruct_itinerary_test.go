package graphs

import (
	"reflect"
	"testing"
)

func TestSolve0332(t *testing.T) {
	tests := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		{
			name: "Example 1",
			tickets: [][]string{
				{"MUC", "LHR"},
				{"JFK", "MUC"},
				{"SFO", "SJC"},
				{"LHR", "SFO"},
			},
			expected: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name: "Example 2",
			tickets: [][]string{
				{"JFK", "SFO"},
				{"JFK", "ATL"},
				{"SFO", "ATL"},
				{"ATL", "JFK"},
				{"ATL", "SFO"},
			},
			expected: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name: "Single ticket",
			tickets: [][]string{
				{"JFK", "LAX"},
			},
			expected: []string{"JFK", "LAX"},
		},
		{
			name: "Cycle",
			tickets: [][]string{
				{"JFK", "ATL"},
				{"ATL", "JFK"},
			},
			expected: []string{"JFK", "ATL", "JFK"},
		},
		{
			name: "Multiple options chooses lexical order",
			tickets: [][]string{
				{"JFK", "AAA"},
				{"JFK", "BBB"},
				{"AAA", "JFK"},
				{"BBB", "JFK"},
			},
			expected: []string{"JFK", "AAA", "JFK", "BBB", "JFK"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solve0332(tt.tickets)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Solve0332(%v) = %v, expected %v", tt.tickets, result, tt.expected)
			}
		})
	}
}