package dp

import (
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name: "example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 4,
		},
		{
			name: "example 2",
			matrix: [][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			want: 1,
		},
		{
			name:   "empty matrix",
			matrix: [][]byte{},
			want:   0,
		},
		{
			name: "single cell with 1",
			matrix: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "single cell with 0",
			matrix: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: "all ones 2x2",
			matrix: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 4,
		},
		{
			name: "all ones 3x3",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 9,
		},
		{
			name: "diagonal ones",
			matrix: [][]byte{
				{'1', '0', '0'},
				{'0', '1', '0'},
				{'0', '0', '1'},
			},
			want: 1,
		},
		{
			name: "larger example",
			matrix: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaximalSquare(tt.matrix)
			if got != tt.want {
				t.Errorf("MaximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}