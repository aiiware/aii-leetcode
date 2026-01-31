package graphs

/*
210. Course Schedule II

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you 
must take course bi first if you want to take course ai.

Return the ordering of courses you should take to finish all courses. If there are many 
valid answers, return any of them. If it is impossible to finish all courses, return an 
empty array.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have 
finished course 0. So the correct course order is [0,1].

Example 2:
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3] or [0,1,2,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have 
finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

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
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
    // Build adjacency list and indegree array
    adj := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
        indegree[course]++
    }
    
    // Initialize queue with courses having no prerequisites
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // Perform topological sort
    result := make([]int, 0, numCourses)
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        result = append(result, current)
        
        // Reduce indegree of neighbors
        for _, neighbor := range adj[current] {
            indegree[neighbor]--
            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    
    // If we couldn't process all courses, there's a cycle
    if len(result) != numCourses {
        return []int{}
    }
    
    return result
}

// DFS version
func findOrderDFS(numCourses int, prerequisites [][]int) []int {
    // Build adjacency list
    adj := make([][]int, numCourses)
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
    }
    
    // 0 = unvisited, 1 = visiting, 2 = visited
    visited := make([]int, numCourses)
    result := make([]int, 0, numCourses)
    hasCycle := false
    
    var dfs func(course int)
    dfs = func(course int) {
        if visited[course] == 1 {
            hasCycle = true
            return
        }
        if visited[course] == 2 {
            return
        }
        
        visited[course] = 1 // Mark as visiting
        
        for _, neighbor := range adj[course] {
            dfs(neighbor)
            if hasCycle {
                return
            }
        }
        
        visited[course] = 2 // Mark as visited
        result = append(result, course)
    }
    
    // Perform DFS from each unvisited course
    for i := 0; i < numCourses; i++ {
        if visited[i] == 0 {
            dfs(i)
            if hasCycle {
                return []int{}
            }
        }
    }
    
    // Reverse the result since we added courses after processing dependencies
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return result
}