package leetcode

// Problem 0095: Unique Binary Search Trees II
//
// Given an integer n, return all the structurally unique BST's (binary search trees), 
// which has exactly n nodes of unique values from 1 to n. Return the answer in any order.
//
// Example 1:
// Input: n = 3
// Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
// Example 2:
// Input: n = 1
// Output: [[1]]
//
// Constraints:
// - 1 <= n <= 8

// generateTrees is the main solution function using recursion with memoization.
// Time complexity: O(4^n / n^(3/2)) (Catalan number), Space complexity: O(4^n / n^(3/2))
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	memo := make(map[[2]int][]*TreeNode)
	return generateTreesHelper(1, n, memo)
}

func generateTreesHelper(start, end int, memo map[[2]int][]*TreeNode) []*TreeNode {
	// Check memo
	key := [2]int{start, end}
	if val, exists := memo[key]; exists {
		return val
	}

	var result []*TreeNode

	// Base case: empty tree
	if start > end {
		result = append(result, nil)
		memo[key] = result
		return result
	}

	// Try each number as root
	for rootVal := start; rootVal <= end; rootVal++ {
		// Generate all left subtrees
		leftSubtrees := generateTreesHelper(start, rootVal-1, memo)
		
		// Generate all right subtrees
		rightSubtrees := generateTreesHelper(rootVal+1, end, memo)
		
		// Combine left and right subtrees with current root
		for _, left := range leftSubtrees {
			for _, right := range rightSubtrees {
				root := &TreeNode{
					Val:   rootVal,
					Left:  left,
					Right: right,
				}
				result = append(result, root)
			}
		}
	}

	memo[key] = result
	return result
}

// generateTreesDP uses dynamic programming (bottom-up).
func generateTreesDP(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	// dp[i] stores all BSTs for i nodes (values 1..i)
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil} // Empty tree

	// Build up from 1 to n
	for length := 1; length <= n; length++ {
		var trees []*TreeNode
		
		// For trees of size 'length', root can be any value from 1 to length
		for rootVal := 1; rootVal <= length; rootVal++ {
			// Left subtree: values 1..rootVal-1 (size = rootVal-1)
			leftTrees := dp[rootVal-1]
			
			// Right subtree: values rootVal+1..length
			// Need to adjust values by adding rootVal
			rightTrees := dp[length-rootVal]
			
			// Combine
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					// For right subtree, we need to add rootVal to all node values
					adjustedRight := cloneAndAdd(right, rootVal)
					
					root := &TreeNode{
						Val:   rootVal,
						Left:  left,
						Right: adjustedRight,
					}
					trees = append(trees, root)
				}
			}
		}
		dp[length] = trees
	}

	return dp[n]
}

// cloneAndAdd clones a tree and adds 'offset' to all node values
func cloneAndAdd(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val + offset,
		Left:  cloneAndAdd(root.Left, offset),
		Right: cloneAndAdd(root.Right, offset),
	}
}

// generateTreesIterative uses iterative approach.
func generateTreesIterative(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	// Start with list containing only nil (empty tree)
	trees := [][]*TreeNode{{nil}}

	// Build trees of increasing size
	for length := 1; length <= n; length++ {
		var currentTrees []*TreeNode
		
		// For each possible root value
		for rootVal := 1; rootVal <= length; rootVal++ {
			// Get left trees of size rootVal-1
			leftTrees := trees[rootVal-1]
			
			// Get right trees of size length-rootVal
			rightTrees := trees[length-rootVal]
			
			// Combine
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					// Adjust right tree values
					adjustedRight := cloneAndAdd(right, rootVal)
					
					currentTrees = append(currentTrees, &TreeNode{
						Val:   rootVal,
						Left:  left,
						Right: adjustedRight,
					})
				}
			}
		}
		trees = append(trees, currentTrees)
	}

	return trees[n]
}

// generateTreesBacktracking uses backtracking approach.
func generateTreesBacktracking(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	// Generate all permutations of 1..n and build BSTs
	// This is inefficient but included for completeness
	var result []*TreeNode
	used := make([]bool, n+1)
	current := make([]int, n)
	
	var backtrack func(int)
	backtrack = func(pos int) {
		if pos == n {
			// Build BST from permutation
			root := buildBSTFromPermutation(current)
			// Check if this BST is unique
			if isUniqueBST(root, result) {
				result = append(result, root)
			}
			return
		}
		
		for i := 1; i <= n; i++ {
			if !used[i] {
				used[i] = true
				current[pos] = i
				backtrack(pos + 1)
				used[i] = false
			}
		}
	}
	
	backtrack(0)
	return result
}

// buildBSTFromPermutation builds a BST by inserting values in given order
func buildBSTFromPermutation(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	
	root := &TreeNode{Val: values[0]}
	for i := 1; i < len(values); i++ {
		insertIntoBST(root, values[i])
	}
	return root
}

// insertIntoBST inserts a value into BST
func insertIntoBST(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertIntoBST(root.Right, val)
		}
	}
}

// isUniqueBST checks if a BST is unique compared to existing ones
func isUniqueBST(root *TreeNode, existing []*TreeNode) bool {
	for _, tree := range existing {
		if treesEqual(root, tree) {
			return false
		}
	}
	return true
}

// treesEqual checks if two BSTs are equal (structurally and values)
func treesEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return treesEqual(a.Left, b.Left) && treesEqual(a.Right, b.Right)
}

// generateTreesCatalan uses Catalan number property.
func generateTreesCatalan(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	
	// The number of unique BSTs is the nth Catalan number
	// C(n) = (2n)! / ((n+1)! * n!)
	// We'll generate using the recursive definition
	
	var generate func(int, int) []*TreeNode
	generate = func(start, end int) []*TreeNode {
		var result []*TreeNode
		
		if start > end {
			result = append(result, nil)
			return result
		}
		
		for i := start; i <= end; i++ {
			leftTrees := generate(start, i-1)
			rightTrees := generate(i+1, end)
			
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					result = append(result, &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					})
				}
			}
		}
		
		return result
	}
	
	return generate(1, n)
}

// generateTreesOptimized is an optimized version with memoization.
func generateTreesOptimized(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	
	// Pre-compute Catalan numbers for bounds checking
	// C(0)=1, C(1)=1, C(2)=2, C(3)=5, C(4)=14, C(5)=42, C(6)=132, C(7)=429, C(8)=1430
	catalan := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430}
	if n < len(catalan) {
		// We know exactly how many trees to expect
		result := make([]*TreeNode, 0, catalan[n])
		memo := make(map[[2]int][]*TreeNode)
		result = generateTreesHelper(1, n, memo)
		return result
	}
	
	// For n > 8 (beyond constraints), use standard approach
	return generateTrees(n)
}

// GenerateTrees is the public interface function.
// It uses the optimized memoization solution by default.
func GenerateTrees(n int) []*TreeNode {
	return generateTreesOptimized(n)
}