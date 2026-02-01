package graphs

import (
	"sort"
)

// Solve0332 solves LeetCode problem 0332: Reconstruct Itinerary
// Difficulty: Hard
// Tags: Depth-First Search, Graph, Eulerian Path
//
// You are given a list of airline tickets where tickets[i] = [fromi, toi] represent
// the departure and the arrival airports of one flight. Reconstruct the itinerary
// in order and return it.
//
// All of the tickets belong to a man who departs from "JFK", thus the itinerary
// must begin with "JFK". If there are multiple valid itineraries, you should return
// the itinerary that has the smallest lexical order when read as a single string.
//
// For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
// You may assume all tickets form at least one valid itinerary. You must use all the tickets
// once and only once.
//
// Example 1:
// Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
// Output: ["JFK","MUC","LHR","SFO","SJC"]
//
// Example 2:
// Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
// Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
// Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"]
// but it is larger in lexical order.
//
// Constraints:
// 1 <= tickets.length <= 300
// tickets[i].length == 2
// fromi.length == 3
// toi.length == 3
// fromi and toi consist of uppercase English letters.
// fromi != toi
//
// Time complexity: O(E log E) where E is number of edges, Space complexity: O(V + E)
func Solve0332(tickets [][]string) []string {
	// Build adjacency list with sorted destinations
	graph := make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		graph[from] = append(graph[from], to)
	}

	// Sort destinations in lexical order for each airport
	for airport := range graph {
		sort.Strings(graph[airport])
	}

	// Use Hierholzer's algorithm for Eulerian path
	itinerary := make([]string, 0)
	var dfs func(airport string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			// Get the next destination (smallest lexical order)
			next := graph[airport][0]
			// Remove this edge (ticket)
			graph[airport] = graph[airport][1:]
			dfs(next)
		}
		// Post-order: add airport to itinerary
		itinerary = append(itinerary, airport)
	}

	// Start from JFK
	dfs("JFK")

	// Reverse the itinerary (post-order gives reverse path)
	for i, j := 0, len(itinerary)-1; i < j; i, j = i+1, j-1 {
		itinerary[i], itinerary[j] = itinerary[j], itinerary[i]
	}

	return itinerary
}