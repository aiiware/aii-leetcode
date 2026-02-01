package graphs

// Solve0310 solves LeetCode problem 0310: Minimum Height Trees
// Difficulty: Medium
// Tags: Breadth-First Search, Graph, Topological Sort
//
// A tree is an undirected graph in which any two vertices are connected by exactly one path.
// In other words, any connected graph without simple cycles is a tree.
//
// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where
// edges[i] = [ai, bi] indicates that there is an undirected edge between nodes ai and bi
// in the tree, you can choose any node as the root. When you select a node x as the root,
// the result tree has height h. Among all possible rooted trees, those with minimum height
// (i.e., min(h)) are called minimum height trees (MHTs).
//
// Return a list of all MHTs' root labels. You can return the answer in any order.
//
// The height of a rooted tree is the number of edges on the longest downward path between
// the root and a leaf.
//
// Example 1:
// Input: n = 4, edges = [[1,0],[1,2],[1,3]]
// Output: [1]
// Explanation: As shown, the height of the tree when root is 1 is 1, while the heights
// of the trees when roots are 0, 2, or 3 are 2.
//
// Example 2:
// Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
// Output: [3,4]
//
// Constraints:
// 1 <= n <= 2 * 10^4
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// All the pairs (ai, bi) are distinct.
// The given input is guaranteed to be a tree and there will be no repeated edges.
//
// Time complexity: O(n), Space complexity: O(n)
func Solve0310(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// Build adjacency list
	adj := make([][]int, n)
	degree := make([]int, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		degree[a]++
		degree[b]++
	}

	// Initialize leaves queue (nodes with degree 1)
	leaves := make([]int, 0)
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			leaves = append(leaves, i)
		}
	}

	// Trim leaves layer by layer until we have 1 or 2 nodes left
	remainingNodes := n
	for remainingNodes > 2 {
		leavesCount := len(leaves)
		remainingNodes -= leavesCount

		// Process current layer of leaves
		newLeaves := make([]int, 0)
		for i := 0; i < leavesCount; i++ {
			leaf := leaves[i]

			// For each neighbor of the leaf
			for _, neighbor := range adj[leaf] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}