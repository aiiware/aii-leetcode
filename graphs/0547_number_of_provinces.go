package graphs

/*
547. Number of Provinces

There are n cities. Some of them are connected, while some are not. 
If city a is connected directly with city b, and city b is connected directly with city c, 
then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city 
are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.

Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

Constraints:
- 1 <= n <= 200
- n == isConnected.length
- n == isConnected[i].length
- isConnected[i][j] is 1 or 0.
- isConnected[i][i] == 1
- isConnected[i][j] == isConnected[j][i]
*/

/*
Difficulty: Medium
Tags: Depth-First Search, Breadth-First Search, Union Find, Graph
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

// FindCircleNum is the main function that uses DFS (can be changed to BFS or Union-Find)
func FindCircleNum(isConnected [][]int) int {
	return findCircleNumDFS(isConnected)
}

// FindCircleNumBFS uses Breadth-First Search to find connected components
func FindCircleNumBFS(isConnected [][]int) int {
	return findCircleNumBFS(isConnected)
}

// FindCircleNumUnionFind uses Union-Find (Disjoint Set Union) to find connected components
func FindCircleNumUnionFind(isConnected [][]int) int {
	return findCircleNumUnionFind(isConnected)
}

// findCircleNumDFS uses Depth-First Search to find connected components
func findCircleNumDFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	count := 0

	// DFS function to mark all connected cities
	var dfs func(city int)
	dfs = func(city int) {
		visited[city] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			// If cities are connected and neighbor not visited
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	// Iterate through all cities
	for city := 0; city < n; city++ {
		if !visited[city] {
			count++
			dfs(city)
		}
	}

	return count
}

// findCircleNumBFS uses Breadth-First Search to find connected components
func findCircleNumBFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	visited := make([]bool, n)
	count := 0

	for city := 0; city < n; city++ {
		if !visited[city] {
			count++

			// BFS queue
			queue := []int{city}
			visited[city] = true

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				// Check all neighbors
				for neighbor := 0; neighbor < n; neighbor++ {
					if isConnected[current][neighbor] == 1 && !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
		}
	}

	return count
}

// findCircleNumUnionFind uses Union-Find (Disjoint Set Union) to find connected components
func findCircleNumUnionFind(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}

	// Initialize parent array for Union-Find
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// Find operation with path compression
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // Path compression
		}
		return parent[x]
	}

	// Union operation
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootX] = rootY
		}
	}

	// Perform union for all connected cities
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	// Count number of unique roots (provinces)
	provinces := make(map[int]bool)
	for i := 0; i < n; i++ {
		provinces[find(i)] = true
	}

	return len(provinces)
}
