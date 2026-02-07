package dp

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n = 0",
			n:    0,
			want: []int{0},
		},
		{
			name: "n = 1",
			n:    1,
			want: []int{0, 1},
		},
		{
			name: "n = 2",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "n = 5",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "n = 8",
			n:    8,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
		{
			name: "n = 15",
			n:    15,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountBits(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits(%d) = %v, want %v", tt.n, got, tt.want)
			}
			
			// Also test the alternative implementation
			got2 := CountBits2(tt.n)
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("CountBits2(%d) = %v, want %v", tt.n, got2, tt.want)
			}
		})
	}
}