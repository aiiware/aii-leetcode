package graphs

/*
207. Course Schedule

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you 
must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.

Example 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should 
also have finished course 1. So it is impossible.

Constraints:
- 1 <= numCourses <= 2000
- 0 <= prerequisites.length <= 5000
- prerequisites[i].length == 2
- 0 <= ai, bi < numCourses
- All the pairs prerequisites[i] are unique.
*/

/*
Difficulty: Medium
Tags: Depth-First Search, Breadth-First Search, Graph, Topological Sort
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn, Yelp
*/

// canFinish uses DFS to detect cycles in a directed graph (Kahn's algorithm - topological sort)
func canFinish(numCourses int, prerequisites [][]int) bool {
    // Build adjacency list and indegree array
    adj := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
        indegree[course]++
    }
    
    // Find all courses with no prerequisites (indegree = 0)
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // Process courses in topological order
    count := 0
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        count++
        
        // Reduce indegree of neighbors
        for _, neighbor := range adj[current] {
            indegree[neighbor]--
            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    
    // If we processed all courses, no cycle exists
    return count == numCourses
}

// canFinishDFS uses DFS with coloring to detect cycles
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
    // Build adjacency list
    adj := make([][]int, numCourses)
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
    }
    
    // 0 = unvisited, 1 = visiting, 2 = visited
    visited := make([]int, numCourses)
    
    var hasCycle func(course int) bool
    hasCycle = func(course int) bool {
        if visited[course] == 1 {
            return true // Cycle detected
        }
        if visited[course] == 2 {
            return false // Already processed
        }
        
        visited[course] = 1 // Mark as visiting
        
        for _, neighbor := range adj[course] {
            if hasCycle(neighbor) {
                return true
            }
        }
        
        visited[course] = 2 // Mark as visited
        return false
    }
    
    // Check for cycles starting from each course
    for i := 0; i < numCourses; i++ {
        if visited[i] == 0 {
            if hasCycle(i) {
                return false
            }
        }
    }
    
    return true
}