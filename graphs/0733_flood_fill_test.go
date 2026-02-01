package graphs

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		sr     int
		sc     int
		color  int
		expect [][]int
	}{
		{
			name: "Example 1",
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:    1,
			sc:    1,
			color: 2,
			expect: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "Example 2",
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:    0,
			sc:    0,
			color: 0,
			expect: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "Single pixel",
			image: [][]int{
				{5},
			},
			sr:    0,
			sc:    0,
			color: 9,
			expect: [][]int{
				{9},
			},
		},
		{
			name: "Different starting color",
			image: [][]int{
				{1, 2, 1},
				{2, 1, 2},
				{1, 2, 1},
			},
			sr:    1,
			sc:    1,
			color: 3,
			expect: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{1, 2, 1},
			},
		},
		{
			name: "Fill entire grid",
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:    0,
			sc:    0,
			color: 1,
			expect: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test BFS implementation
			imageCopy := make([][]int, len(tt.image))
			for i := range tt.image {
				imageCopy[i] = make([]int, len(tt.image[i]))
				copy(imageCopy[i], tt.image[i])
			}

			result := floodFill(imageCopy, tt.sr, tt.sc, tt.color)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("floodFill(BFS) failed for %s\nGot: %v\nWant: %v", tt.name, result, tt.expect)
			}

			// Test DFS implementation
			imageCopy2 := make([][]int, len(tt.image))
			for i := range tt.image {
				imageCopy2[i] = make([]int, len(tt.image[i]))
				copy(imageCopy2[i], tt.image[i])
			}

			result2 := floodFillDFS(imageCopy2, tt.sr, tt.sc, tt.color)
			if !reflect.DeepEqual(result2, tt.expect) {
				t.Errorf("floodFill(DFS) failed for %s\nGot: %v\nWant: %v", tt.name, result2, tt.expect)
			}
		})
	}
}

func TestFloodFillEdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		sr     int
		sc     int
		color  int
	}{
		{
			name:  "Empty image",
			image: [][]int{},
			sr:    0,
			sc:    0,
			color: 1,
		},
		{
			name: "Empty row",
			image: [][]int{
				{},
			},
			sr:    0,
			sc:    0,
			color: 1,
		},
		{
			name: "Out of bounds start",
			image: [][]int{
				{1, 2},
				{3, 4},
			},
			sr:    5,
			sc:    5,
			color: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just ensure no panic
			_ = floodFill(tt.image, tt.sr, tt.sc, tt.color)
			_ = floodFillDFS(tt.image, tt.sr, tt.sc, tt.color)
		})
	}
}