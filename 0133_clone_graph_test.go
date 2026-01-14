package leetcode

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {
	// Helper function to create a graph from adjacency list
	createGraph := func(adjList [][]int) *GraphNode {
		if len(adjList) == 0 {
			return nil
		}
		
		// Create all nodes
		nodes := make([]*GraphNode, len(adjList))
		for i := range nodes {
			nodes[i] = NewGraphNode(i + 1) // Node values are 1-indexed
		}
		
		// Connect neighbors
		for i, neighbors := range adjList {
			for _, neighborIdx := range neighbors {
				// neighborIdx is 1-indexed in the input, convert to 0-indexed
				nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighborIdx-1])
			}
		}
		
		return nodes[0]
	}
	
	// Helper function to convert graph to adjacency list
	graphToAdjList := func(node *GraphNode) [][]int {
		if node == nil {
			return [][]int{}
		}
		
		// Use BFS to traverse the graph and collect all nodes
		visited := make(map[*GraphNode]bool)
		queue := []*GraphNode{node}
		visited[node] = true
		allNodes := []*GraphNode{}
		
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			allNodes = append(allNodes, current)
			
			for _, neighbor := range current.Neighbors {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		
		// Sort nodes by value (1, 2, 3, ...)
		// Since values are unique and from 1 to n, we can create a slice
		// indexed by value
		maxVal := 0
		for _, n := range allNodes {
			if n.Val > maxVal {
				maxVal = n.Val
			}
		}
		
		// Create a map from value to node for quick lookup
		valueToNode := make(map[int]*GraphNode, maxVal)
		for _, n := range allNodes {
			valueToNode[n.Val] = n
		}
		
		// Build adjacency list in order of node values (1, 2, 3, ...)
		adjList := make([][]int, maxVal)
		for i := 1; i <= maxVal; i++ {
			node := valueToNode[i]
			neighbors := []int{}
			for _, neighbor := range node.Neighbors {
				neighbors = append(neighbors, neighbor.Val)
			}
			adjList[i-1] = neighbors
		}
		
		return adjList
	}
	
	tests := []struct {
		name     string
		adjList  [][]int
		expected [][]int
	}{
		{
			name: "Example 1",
			adjList: [][]int{
				{2, 4},
				{1, 3},
				{2, 4},
				{1, 3},
			},
			expected: [][]int{
				{2, 4},
				{1, 3},
				{2, 4},
				{1, 3},
			},
		},
		{
			name: "Example 2",
			adjList: [][]int{
				{},
			},
			expected: [][]int{
				{},
			},
		},
		{
			name:     "Example 3",
			adjList:  [][]int{},
			expected: [][]int{},
		},
		{
			name: "Single node",
			adjList: [][]int{
				{},
			},
			expected: [][]int{
				{},
			},
		},
		{
			name: "Two connected nodes",
			adjList: [][]int{
				{2},
				{1},
			},
			expected: [][]int{
				{2},
				{1},
			},
		},
		{
			name: "Three nodes in line",
			adjList: [][]int{
				{2},
				{1, 3},
				{2},
			},
			expected: [][]int{
				{2},
				{1, 3},
				{2},
			},
		},
		{
			name: "Complete graph of 4 nodes",
			adjList: [][]int{
				{2, 3, 4},
				{1, 3, 4},
				{1, 2, 4},
				{1, 2, 3},
			},
			expected: [][]int{
				{2, 3, 4},
				{1, 3, 4},
				{1, 2, 4},
				{1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := createGraph(tt.adjList)
			cloned := CloneGraph(graph)
			result := graphToAdjList(cloned)
			
			// Compare adjacency lists
			if len(result) != len(tt.expected) {
				t.Errorf("CloneGraph() result length = %d, expected %d", len(result), len(tt.expected))
				return
			}
			
			for i := range result {
				if len(result[i]) != len(tt.expected[i]) {
					t.Errorf("CloneGraph() neighbors count for node %d = %d, expected %d", 
						i+1, len(result[i]), len(tt.expected[i]))
					return
				}
				
				// Sort neighbors for comparison (graph is undirected, order doesn't matter)
				sortInts := func(arr []int) {
					for i := 0; i < len(arr); i++ {
						for j := i + 1; j < len(arr); j++ {
							if arr[i] > arr[j] {
								arr[i], arr[j] = arr[j], arr[i]
							}
						}
					}
				}
				
				sortInts(result[i])
				sortInts(tt.expected[i])
				
				for j := range result[i] {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("CloneGraph() neighbor %d for node %d = %d, expected %d", 
							j+1, i+1, result[i][j], tt.expected[i][j])
					}
				}
			}
		})
	}
}

func BenchmarkCloneGraph(b *testing.B) {
	// Create a complex graph for benchmarking
	createComplexGraph := func(nodes int) *GraphNode {
		if nodes == 0 {
			return nil
		}
		
		// Create all nodes
		graphNodes := make([]*GraphNode, nodes)
		for i := range graphNodes {
			graphNodes[i] = NewGraphNode(i + 1)
		}
		
		// Connect each node to its next 3 neighbors (circular)
		for i := 0; i < nodes; i++ {
			for j := 1; j <= 3; j++ {
				neighborIdx := (i + j) % nodes
				graphNodes[i].Neighbors = append(graphNodes[i].Neighbors, graphNodes[neighborIdx])
			}
		}
		
		return graphNodes[0]
	}
	
	testCases := []struct {
		name  string
		nodes int
	}{
		{"Small graph", 10},
		{"Medium graph", 50},
		{"Large graph", 100},
		{"Very large graph", 200},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			graph := createComplexGraph(tc.nodes)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				CloneGraph(graph)
			}
		})
	}
}