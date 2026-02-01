package graphs

// Solve0684 solves LeetCode problem 0684: Redundant Connection
// Difficulty: Medium
// Tags: Union Find, Graph, Tree
//
// In this problem, a tree is an undirected graph that is connected and has no cycles.
//
// You are given a graph that started as a tree with n nodes labeled from 1 to n,
// with one additional edge added. The added edge has two different vertices chosen
// from 1 to n, and was not an edge that already existed.
//
// The graph is represented as an array edges of length n where edges[i] = [ai, bi]
// indicates that there is an edge between nodes ai and bi in the graph.
//
// Return an edge that can be removed so that the resulting graph is a tree of n nodes.
// If there are multiple answers, return the answer that occurs last in the input.
//
// Example 1:
// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]
//
// Example 2:
// Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
// Output: [1,4]
//
// Constraints:
// n == edges.length
// 3 <= n <= 1000
// edges[i].length == 2
// 1 <= ai < bi <= n
// ai != bi
// There are no repeated edges.
// The given graph is connected.
//
// Time complexity: O(n * α(n)) where α is inverse Ackermann function, Space complexity: O(n)
func Solve0684(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1) // nodes are 1-indexed
	rank := make([]int, n+1)

	// Initialize Union-Find
	for i := 1; i <= n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	// Find with path compression
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Union by rank
	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)

		if rootX == rootY {
			return false // Already connected, this edge creates a cycle
		}

		// Union by rank
		if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}

		return true
	}

	// Process edges
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if !union(a, b) {
			// This edge creates a cycle, return it
			return []int{a, b}
		}
	}

	// Should never reach here given problem constraints
	return []int{}
}