package benchmarks

import (
	"leetcode/graphs"
	"testing"
)

// BenchmarkCloneGraph benchmarks graph cloning
func BenchmarkCloneGraph(b *testing.B) {
	// Create a simple graph for testing
	// Graph structure: 1-2-3, 1-4, 2-5
	node1 := &graphs.GraphNode{Val: 1}
	node2 := &graphs.GraphNode{Val: 2}
	node3 := &graphs.GraphNode{Val: 3}
	node4 := &graphs.GraphNode{Val: 4}
	node5 := &graphs.GraphNode{Val: 5}
	
	// Set up neighbors
	node1.Neighbors = []*graphs.GraphNode{node2, node4}
	node2.Neighbors = []*graphs.GraphNode{node1, node3, node5}
	node3.Neighbors = []*graphs.GraphNode{node2}
	node4.Neighbors = []*graphs.GraphNode{node1}
	node5.Neighbors = []*graphs.GraphNode{node2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graphs.CloneGraph(node1)
	}
}

// BenchmarkCloneGraphSmall benchmarks with small graph
func BenchmarkCloneGraphSmall(b *testing.B) {
	// Single node graph
	node1 := &graphs.GraphNode{Val: 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graphs.CloneGraph(node1)
	}
}

// BenchmarkCloneGraphLarge benchmarks with larger graph
func BenchmarkCloneGraphLarge(b *testing.B) {
	// Create a graph with 100 nodes in a chain
	nodes := make([]*graphs.GraphNode, 100)
	for i := range nodes {
		nodes[i] = &graphs.GraphNode{Val: i + 1}
	}
	
	// Connect them in a chain
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Neighbors = []*graphs.GraphNode{nodes[i+1]}
		nodes[i+1].Neighbors = []*graphs.GraphNode{nodes[i]}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graphs.CloneGraph(nodes[0])
	}
}

// BenchmarkCloneGraphComplete benchmarks with complete graph
func BenchmarkCloneGraphComplete(b *testing.B) {
	// Create a complete graph with 10 nodes (each node connected to all others)
	size := 10
	nodes := make([]*graphs.GraphNode, size)
	for i := range nodes {
		nodes[i] = &graphs.GraphNode{Val: i + 1}
	}
	
	// Connect each node to all other nodes
	for i := 0; i < size; i++ {
		neighbors := make([]*graphs.GraphNode, 0, size-1)
		for j := 0; j < size; j++ {
			if i != j {
				neighbors = append(neighbors, nodes[j])
			}
		}
		nodes[i].Neighbors = neighbors
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graphs.CloneGraph(nodes[0])
	}
}