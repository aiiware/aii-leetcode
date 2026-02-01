package graphs

import (
	"container/heap"
	"math"
)

// Solve0743 solves LeetCode problem 0743: Network Delay Time
// Difficulty: Medium
// Tags: Breadth-First Search, Graph, Heap (Priority Queue), Shortest Path
//
// You are given a network of n nodes, labeled from 1 to n. You are also given times,
// a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the
// source node, vi is the target node, and wi is the time it takes for a signal to
// travel from source to target.
//
// We will send a signal from a given node k. Return the minimum time it takes for
// all the n nodes to receive the signal. If it is impossible for all the n nodes
// to receive the signal, return -1.
//
// Example 1:
// Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
// Output: 2
//
// Example 2:
// Input: times = [[1,2,1]], n = 2, k = 1
// Output: 1
//
// Example 3:
// Input: times = [[1,2,1]], n = 2, k = 2
// Output: -1
//
// Constraints:
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// All the pairs (ui, vi) are unique. (i.e., no multiple edges.)
//
// Time complexity: O(E log V) using Dijkstra's algorithm, Space complexity: O(V + E)
func Solve0743(times [][]int, n int, k int) int {
	// Build adjacency list
	adj := make([][][2]int, n+1) // 1-indexed
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		adj[u] = append(adj[u], [2]int{v, w})
	}

	// Dijkstra's algorithm
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	// Min-heap priority queue
	pq := &minHeap{}
	heap.Init(pq)
	heap.Push(pq, [2]int{k, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).([2]int)
		node, time := current[0], current[1]

		// If we found a better path already, skip
		if time > dist[node] {
			continue
		}

		// Explore neighbors
		for _, neighbor := range adj[node] {
			nextNode, weight := neighbor[0], neighbor[1]
			newDist := time + weight

			if newDist < dist[nextNode] {
				dist[nextNode] = newDist
				heap.Push(pq, [2]int{nextNode, newDist})
			}
		}
	}

	// Find maximum distance (time for all nodes to receive signal)
	maxTime := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1 // Some node is unreachable
		}
		if dist[i] > maxTime {
			maxTime = dist[i]
		}
	}

	return maxTime
}

// Min-heap implementation for Dijkstra's algorithm
type minHeap [][2]int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}