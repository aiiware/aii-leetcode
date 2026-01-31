package graphs

// 0785. Is Graph Bipartite?
// https://leetcode.com/problems/is-graph-bipartite/

// IsBipartite checks if a graph is bipartite using BFS coloring approach
// A graph is bipartite if we can color its vertices with two colors such that
// no two adjacent vertices share the same color
func IsBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n) // 0 = uncolored, 1 = color A, -1 = color B

	for i := 0; i < n; i++ {
		// If node is already colored, skip
		if colors[i] != 0 {
			continue
		}

		// Start BFS from this node
		queue := []int{i}
		colors[i] = 1 // Color with first color

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				// If neighbor has same color as current node, graph is not bipartite
				if colors[neighbor] == colors[node] {
					return false
				}

				// If neighbor is uncolored, color it with opposite color and add to queue
				if colors[neighbor] == 0 {
					colors[neighbor] = -colors[node]
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return true
}

// IsBipartiteDFS checks if a graph is bipartite using DFS coloring approach
func IsBipartiteDFS(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)

	var dfs func(node, color int) bool
	dfs = func(node, color int) bool {
		// Color the current node
		colors[node] = color

		// Check all neighbors
		for _, neighbor := range graph[node] {
			// If neighbor has same color, graph is not bipartite
			if colors[neighbor] == color {
				return false
			}
			// If neighbor is uncolored, recursively color it with opposite color
			if colors[neighbor] == 0 && !dfs(neighbor, -color) {
				return false
			}
		}
		return true
	}

	// Check all connected components
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}