package graphs

import (
	"testing"
)

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{
			name: "bipartite graph",
			graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			want: true,
		},
		{
			name: "not bipartite (triangle)",
			graph: [][]int{
				{1, 2},
				{0, 2},
				{0, 1},
			},
			want: false,
		},
		{
			name: "single node",
			graph: [][]int{
				{},
			},
			want: true,
		},
		{
			name: "two disconnected components",
			graph: [][]int{
				{1},
				{0},
				{3},
				{2},
			},
			want: true,
		},
		{
			name: "odd cycle",
			graph: [][]int{
				{1, 4},
				{0, 2},
				{1, 3},
				{2, 4},
				{3, 0},
			},
			want: false,
		},
		{
			name: "empty graph",
			graph: [][]int{},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBipartite(tt.graph); got != tt.want {
				t.Errorf("IsBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBipartiteDFS(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{
			name: "bipartite graph",
			graph: [][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			want: true,
		},
		{
			name: "not bipartite (triangle)",
			graph: [][]int{
				{1, 2},
				{0, 2},
				{0, 1},
			},
			want: false,
		},
		{
			name: "single node",
			graph: [][]int{
				{},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBipartiteDFS(tt.graph); got != tt.want {
				t.Errorf("IsBipartiteDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsBipartite(b *testing.B) {
	graph := [][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}
	for i := 0; i < b.N; i++ {
		IsBipartite(graph)
	}
}

func BenchmarkIsBipartiteDFS(b *testing.B) {
	graph := [][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}
	for i := 0; i < b.N; i++ {
		IsBipartiteDFS(graph)
	}
}