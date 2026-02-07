package dp

import (
	"testing"
)

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n = 12",
			n:    12,
			want: 3, // 4 + 4 + 4
		},
		{
			name: "n = 13",
			n:    13,
			want: 2, // 4 + 9
		},
		{
			name: "n = 1",
			n:    1,
			want: 1, // 1
		},
		{
			name: "n = 4",
			n:    4,
			want: 1, // 4
		},
		{
			name: "n = 9",
			n:    9,
			want: 1, // 9
		},
		{
			name: "n = 16",
			n:    16,
			want: 1, // 16
		},
		{
			name: "n = 2",
			n:    2,
			want: 2, // 1 + 1
		},
		{
			name: "n = 3",
			n:    3,
			want: 3, // 1 + 1 + 1
		},
		{
			name: "n = 7",
			n:    7,
			want: 4, // 4 + 1 + 1 + 1
		},
		{
			name: "n = 18",
			n:    18,
			want: 2, // 9 + 9
		},
		{
			name: "n = 43",
			n:    43,
			want: 3, // 25 + 9 + 9
		},
		{
			name: "n = 0",
			n:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumSquares(tt.n)
			if got != tt.want {
				t.Errorf("NumSquares(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}