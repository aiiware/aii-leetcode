package graphs

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name         string
		isConnected  [][]int
		wantProvinces int
	}{
		{
			name: "three cities, two provinces",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			wantProvinces: 2,
		},
		{
			name: "all cities connected",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			wantProvinces: 1,
		},
		{
			name: "no connections",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			wantProvinces: 3,
		},
		{
			name: "single city",
			isConnected: [][]int{
				{1},
			},
			wantProvinces: 1,
		},
		{
			name: "chain of cities",
			isConnected: [][]int{
				{1, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 1, 1},
				{0, 0, 1, 1},
			},
			wantProvinces: 1,
		},
		{
			name: "two separate chains",
			isConnected: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 1, 0},
				{0, 0, 1, 1, 1},
				{0, 0, 0, 1, 1},
			},
			wantProvinces: 2,
		},
		{
			name:         "empty matrix",
			isConnected:  [][]int{},
			wantProvinces: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCircleNum(tt.isConnected); got != tt.wantProvinces {
				t.Errorf("FindCircleNum() = %v, want %v", got, tt.wantProvinces)
			}
		})
	}
}

func TestFindCircleNumBFS(t *testing.T) {
	tests := []struct {
		name         string
		isConnected  [][]int
		wantProvinces int
	}{
		{
			name: "three cities, two provinces",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			wantProvinces: 2,
		},
		{
			name: "all cities connected",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			wantProvinces: 1,
		},
		{
			name: "no connections",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			wantProvinces: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCircleNumBFS(tt.isConnected); got != tt.wantProvinces {
				t.Errorf("FindCircleNumBFS() = %v, want %v", got, tt.wantProvinces)
			}
		})
	}
}

func TestFindCircleNumUnionFind(t *testing.T) {
	tests := []struct {
		name         string
		isConnected  [][]int
		wantProvinces int
	}{
		{
			name: "three cities, two provinces",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			wantProvinces: 2,
		},
		{
			name: "all cities connected",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			wantProvinces: 1,
		},
		{
			name: "no connections",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			wantProvinces: 3,
		},
		{
			name: "complex connections",
			isConnected: [][]int{
				{1, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0},
				{0, 0, 0, 1, 1, 0},
				{0, 0, 0, 0, 0, 1},
			},
			wantProvinces: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCircleNumUnionFind(tt.isConnected); got != tt.wantProvinces {
				t.Errorf("FindCircleNumUnionFind() = %v, want %v", got, tt.wantProvinces)
			}
		})
	}
}

func BenchmarkFindCircleNum(b *testing.B) {
	isConnected := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1},
	}
	for i := 0; i < b.N; i++ {
		FindCircleNum(isConnected)
	}
}

func BenchmarkFindCircleNumBFS(b *testing.B) {
	isConnected := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1},
	}
	for i := 0; i < b.N; i++ {
		FindCircleNumBFS(isConnected)
	}
}

func BenchmarkFindCircleNumUnionFind(b *testing.B) {
	isConnected := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1},
	}
	for i := 0; i < b.N; i++ {
		FindCircleNumUnionFind(isConnected)
	}
}