package math

import "testing"

func TestRisingTemperature(t *testing.T) {
	tests := []struct {
		name     string
		temps    []int
		expected []int
	}{
		{
			name:     "increasing temps",
			temps:    []int{10, 20, 30, 40},
			expected: []int{1, 2, 3},
		},
		{
			name:     "decreasing temps",
			temps:    []int{40, 30, 20, 10},
			expected: []int{},
		},
		{
			name:     "mixed temps",
			temps:    []int{10, 20, 15, 25, 20},
			expected: []int{1, 3},
		},
		{
			name:     "same temps",
			temps:    []int{20, 20, 20},
			expected: []int{},
		},
		{
			name:     "empty",
			temps:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RisingTemperature(tt.temps)
			if len(result) != len(tt.expected) {
				t.Errorf("RisingTemperature() = %v, want %v", result, tt.expected)
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("RisingTemperature()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func BenchmarkRisingTemperature(b *testing.B) {
	temps := []int{10, 20, 15, 25, 20, 30, 25, 35, 40, 38}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RisingTemperature(temps)
	}
}
