package graphs


/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

/*
# 0133 - Clone Graph
## Problem Description
Given a reference of a node in a connected undirected graph.
Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

## Examples
Example 1:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val=1)'s neighbors are 2nd node (val=2) and 4th node (val=4).
2nd node (val=2)'s neighbors are 1st node (val=1) and 3rd node (val=3).
3rd node (val=3)'s neighbors are 2nd node (val=2) and 4th node (val=4).
4th node (val=4)'s neighbors are 1st node (val=1) and 3rd node (val=3).

Example 2:
Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph has only one node with val=1 and it does not have any neighbors.

Example 3:
Input: adjList = []
Output: []

## Constraints
- The number of nodes in the graph is in the range [0, 100].
- 1 <= Node.val <= 100
- Node.val is unique for each node.
- There are no repeated edges and no self-loops in the graph.
- The graph is connected and undirected.

## Solution Approach
This problem can be solved using BFS or DFS with a hash map:
1. Use a map to store mapping from original nodes to cloned nodes
2. Traverse the graph using BFS/DFS
3. For each node, create a clone if not already created
4. For each neighbor, create clone and add to neighbors list
5. Return the clone of the starting node

Time Complexity: O(V + E) where V is vertices and E is edges
Space Complexity: O(V) for the hash map and recursion stack/BFS queue
*/

// GraphNode represents a node in an undirected graph
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// NewGraphNode creates a new GraphNode with the given value
func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: []*GraphNode{},
	}
}

// CloneGraph returns a deep copy of the graph starting from the given node
func CloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	
	// Map to store mapping from original nodes to cloned nodes
	visited := make(map[*GraphNode]*GraphNode)
	
	// Use BFS to traverse the graph
	queue := []*GraphNode{node}
	visited[node] = NewGraphNode(node.Val)
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		// Process all neighbors
		for _, neighbor := range current.Neighbors {
			// If neighbor hasn't been visited/cloned yet
			if _, exists := visited[neighbor]; !exists {
				// Create clone of neighbor
				visited[neighbor] = NewGraphNode(neighbor.Val)
				// Add to queue for BFS
				queue = append(queue, neighbor)
			}
			// Add cloned neighbor to cloned current node's neighbors
			visited[current].Neighbors = append(visited[current].Neighbors, visited[neighbor])
		}
	}
	
	return visited[node]
}