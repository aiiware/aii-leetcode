package graphs

// 0547. Number of Provinces
// https://leetcode.com/problems/number-of-provinces/

// FindCircleNum finds the number of connected components in an undirected graph
// represented by an adjacency matrix (isConnected[i][j] = 1 if cities i and j are connected)
func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	provinces := 0

	// DFS to mark all cities in the same province
	var dfs func(city int)
	dfs = func(city int) {
		visited[city] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	// Count provinces
	for city := 0; city < n; city++ {
		if !visited[city] {
			provinces++
			dfs(city)
		}
	}

	return provinces
}

// FindCircleNumBFS finds provinces using BFS
func FindCircleNumBFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	provinces := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			provinces++
			
			// BFS for this province
			queue := []int{i}
			visited[i] = true
			
			for len(queue) > 0 {
				city := queue[0]
				queue = queue[1:]
				
				for neighbor := 0; neighbor < n; neighbor++ {
					if isConnected[city][neighbor] == 1 && !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
		}
	}

	return provinces
}

// FindCircleNumUnionFind finds provinces using Union-Find (Disjoint Set Union)
func FindCircleNumUnionFind(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	// Initialize Union-Find
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	// Find with path compression
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // Path compression
			x = parent[x]
		}
		return x
	}

	// Union by rank
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return
		}
		
		if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
	}

	// Union connected cities
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	// Count unique roots
	roots := make(map[int]bool)
	for i := 0; i < n; i++ {
		roots[find(i)] = true
	}

	return len(roots)
}