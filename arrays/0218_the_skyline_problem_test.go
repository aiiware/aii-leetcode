package arrays

import (
	"testing"
)

func TestGetSkyline(t *testing.T) {
	tests := []struct {
		name     string
		buildings [][]int
		want     [][]int
	}{
		{
			name: "Example 1",
			buildings: [][]int{
				{2, 9, 10},
				{3, 7, 15},
				{5, 12, 12},
				{15, 20, 10},
				{19, 24, 8},
			},
			want: [][]int{
				{2, 10},
				{3, 15},
				{7, 12},
				{12, 0},
				{15, 10},
				{20, 8},
				{24, 0},
			},
		},
		{
			name: "Example 2",
			buildings: [][]int{
				{0, 2, 3},
				{2, 5, 3},
			},
			want: [][]int{
				{0, 3},
				{5, 0},
			},
		},
		{
			name:     "Empty input",
			buildings: [][]int{},
			want:     [][]int{},
		},
		{
			name: "Single building",
			buildings: [][]int{
				{0, 5, 10},
			},
			want: [][]int{
				{0, 10},
				{5, 0},
			},
		},
		{
			name: "Overlapping buildings with same height",
			buildings: [][]int{
				{1, 5, 5},
				{2, 6, 3},
				{3, 7, 5},
			},
			// Building 1 [1,5] h=5, Building 2 [2,6] h=3, Building 3 [3,7] h=5
			// At pos 1: start 5, max=5 -> add [1,5]
			// At pos 2: start 3, max=5 -> no change
			// At pos 3: start 5, max=5 -> no change
			// At pos 5: end 5, max=5 (building 3 still active) -> no change
			// At pos 6: end 3, max=5 -> no change
			// At pos 7: end 5, max=0 -> add [7,0]
			want: [][]int{
				{1, 5},
				{7, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetSkyline(tt.buildings)
			if !equalSkylines(result, tt.want) {
				t.Errorf("GetSkyline() = %v, want %v", result, tt.want)
			}
		})
	}
}

func equalSkylines(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}
	return true
}
