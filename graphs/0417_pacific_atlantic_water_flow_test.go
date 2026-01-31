package graphs

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "example 1",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			name: "single cell",
			heights: [][]int{
				{1},
			},
			want: [][]int{{0, 0}},
		},
		{
			name: "2x2 grid",
			heights: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{{0, 1}, {1, 0}, {1, 1}},
		},
		{
			name: "all same height",
			heights: [][]int{
				{5, 5},
				{5, 5},
			},
			want: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name:    "empty grid",
			heights: [][]int{},
			want:    [][]int{},
		},
		{
			name: "waterfall pattern",
			heights: [][]int{
				{3, 3, 3},
				{3, 1, 3},
				{3, 3, 3},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PacificAtlantic(tt.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PacificAtlantic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPacificAtlanticBFS(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "example 1",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			name: "single cell",
			heights: [][]int{
				{1},
			},
			want: [][]int{{0, 0}},
		},
		{
			name: "2x2 grid",
			heights: [][]int{
				{1, 2},
				{3, 4},
			},
			want: [][]int{{0, 1}, {1, 0}, {1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PacificAtlanticBFS(tt.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PacificAtlanticBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPacificAtlantic(b *testing.B) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	for i := 0; i < b.N; i++ {
		PacificAtlantic(heights)
	}
}

func BenchmarkPacificAtlanticBFS(b *testing.B) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	for i := 0; i < b.N; i++ {
		PacificAtlanticBFS(heights)
	}
}