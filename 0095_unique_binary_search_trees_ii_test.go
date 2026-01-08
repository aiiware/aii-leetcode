package leetcode

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		validate func([]*TreeNode) bool
	}{
		{
			name: "n = 0",
			n:    0,
			validate: func(trees []*TreeNode) bool {
				return len(trees) == 0
			},
		},
		{
			name: "n = 1",
			n:    1,
			validate: func(trees []*TreeNode) bool {
				if len(trees) != 1 {
					return false
				}
				// Should be a single node with value 1
				return trees[0].Val == 1 && trees[0].Left == nil && trees[0].Right == nil
			},
		},
		{
			name: "n = 2",
			n:    2,
			validate: func(trees []*TreeNode) bool {
				// Should have 2 unique BSTs
				if len(trees) != 2 {
					return false
				}
				// Check all trees are valid BSTs with values 1..n
				for _, tree := range trees {
					if !isValidBSTHelper95(tree, 1, 2) {
						return false
					}
					if !containsValues1toN(tree, 2) {
						return false
					}
				}
				// Check all trees are unique
				return allTreesUnique(trees)
			},
		},
		{
			name: "n = 3",
			n:    3,
			validate: func(trees []*TreeNode) bool {
				// Should have 5 unique BSTs (3rd Catalan number)
				if len(trees) != 5 {
					return false
				}
				// Check all trees are valid BSTs with values 1..n
				for _, tree := range trees {
					if !isValidBSTHelper95(tree, 1, 3) {
						return false
					}
					if !containsValues1toN(tree, 3) {
						return false
					}
				}
				// Check all trees are unique
				return allTreesUnique(trees)
			},
		},
		{
			name: "n = 4",
			n:    4,
			validate: func(trees []*TreeNode) bool {
				// Should have 14 unique BSTs (4th Catalan number)
				if len(trees) != 14 {
					return false
				}
				// Check all trees are valid BSTs with values 1..n
				for _, tree := range trees {
					if !isValidBSTHelper95(tree, 1, 4) {
						return false
					}
					if !containsValues1toN(tree, 4) {
						return false
					}
				}
				// Check all trees are unique
				return allTreesUnique(trees)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateTrees(tt.n)
			if !tt.validate(result) {
				t.Errorf("GenerateTrees(%d) failed validation", tt.n)
				// Print trees for debugging
				for i, tree := range result {
					t.Logf("Tree %d: %v", i, tree.ToSlice())
				}
			}
		})
	}
}

func TestAllGenerateTreesImplementations(t *testing.T) {
	testCases := []struct {
		name string
		n    int
	}{
		{"n=0", 0},
		{"n=1", 1},
		{"n=2", 2},
		{"n=3", 3},
		{"n=4", 4},
	}

	implementations := []struct {
		name string
		fn   func(int) []*TreeNode
	}{
		{"generateTrees", generateTrees},
		{"generateTreesDP", generateTreesDP},
		{"generateTreesIterative", generateTreesIterative},
		{"generateTreesCatalan", generateTreesCatalan},
		{"generateTreesOptimized", generateTreesOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := GenerateTrees(tc.n)
			expectedCount := len(expected)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.n)
					
					// Check count matches
					if len(result) != expectedCount {
						t.Errorf("%s(%d) returned %d trees, expected %d",
							impl.name, tc.n, len(result), expectedCount)
					}

					// Check all trees are valid BSTs with values 1..n
					for _, tree := range result {
						if !isValidBSTHelper95(tree, 1, tc.n) {
							t.Errorf("%s(%d) produced invalid BST: %v",
								impl.name, tc.n, tree.ToSlice())
						}
						if !containsValues1toN(tree, tc.n) {
							t.Errorf("%s(%d) tree doesn't contain all values 1..%d: %v",
								impl.name, tc.n, tc.n, tree.ToSlice())
						}
					}

					// Check all trees are unique (for n > 0)
					if tc.n > 0 && !allTreesUnique(result) {
						t.Errorf("%s(%d) produced duplicate trees",
							impl.name, tc.n)
					}
				})
			}
		})
	}
}

func TestGenerateTreesEdgeCases(t *testing.T) {
	t.Run("n = 0", func(t *testing.T) {
		result := GenerateTrees(0)
		if len(result) != 0 {
			t.Errorf("GenerateTrees(0) should return empty slice, got %d trees", len(result))
		}
	})

	t.Run("n = 1", func(t *testing.T) {
		result := GenerateTrees(1)
		if len(result) != 1 {
			t.Errorf("GenerateTrees(1) should return 1 tree, got %d", len(result))
		}
		tree := result[0]
		if tree.Val != 1 || tree.Left != nil || tree.Right != nil {
			t.Errorf("GenerateTrees(1) should return single node with value 1, got %v", tree.ToSlice())
		}
	})

	t.Run("n = 8 (maximum per constraints)", func(t *testing.T) {
		result := GenerateTrees(8)
		// 8th Catalan number is 1430
		expectedCount := 1430
		if len(result) != expectedCount {
			t.Errorf("GenerateTrees(8) should return %d trees (8th Catalan number), got %d",
				expectedCount, len(result))
		}

		// Check a sample of trees
		for i := 0; i < min(10, len(result)); i++ {
			tree := result[i]
			if !isValidBSTHelper95(tree, 1, 8) {
				t.Errorf("Tree %d is not a valid BST: %v", i, tree.ToSlice())
			}
			if !containsValues1toN(tree, 8) {
				t.Errorf("Tree %d doesn't contain all values 1..8: %v", i, tree.ToSlice())
			}
		}

		// Check uniqueness
		if !allTreesUnique(result) {
			t.Errorf("GenerateTrees(8) produced duplicate trees")
		}
	})

	t.Run("All trees have exactly n nodes", func(t *testing.T) {
		for n := 1; n <= 5; n++ {
			t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
				result := GenerateTrees(n)
				for i, tree := range result {
					nodeCount := countNodesHelper(tree)
					if nodeCount != n {
						t.Errorf("Tree %d has %d nodes, expected %d: %v",
							i, nodeCount, n, tree.ToSlice())
					}
				}
			})
		}
	})

	t.Run("All values are 1..n", func(t *testing.T) {
		for n := 1; n <= 5; n++ {
			t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
				result := GenerateTrees(n)
				for i, tree := range result {
					values := getAllValuesHelper(tree)
					// Check we have exactly n values
					if len(values) != n {
						t.Errorf("Tree %d has %d values, expected %d: %v",
							i, len(values), n, tree.ToSlice())
					}
					// Check all values are in range 1..n
					for _, val := range values {
						if val < 1 || val > n {
							t.Errorf("Tree %d has value %d outside range 1..%d: %v",
								i, val, n, tree.ToSlice())
						}
					}
					// Check no duplicates
					valueSet := make(map[int]bool)
					for _, val := range values {
						if valueSet[val] {
							t.Errorf("Tree %d has duplicate value %d: %v",
								i, val, tree.ToSlice())
						}
						valueSet[val] = true
					}
				}
			})
		}
	})
}

func TestGenerateTreesProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(int) []*TreeNode
	}{
		{"generateTrees", generateTrees},
		{"generateTreesDP", generateTreesDP},
		{"generateTreesIterative", generateTreesIterative},
		{"generateTreesCatalan", generateTreesCatalan},
		{"generateTreesOptimized", generateTreesOptimized},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for n := 1; n <= 5; n++ {
				t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
					result := impl.fn(n)

					// Property 1: Number of trees should be nth Catalan number
					catalanNumbers := []int{1, 1, 2, 5, 14, 42}
					expectedCount := catalanNumbers[n]
					if len(result) != expectedCount {
						t.Errorf("Expected %d trees (C%d), got %d",
							expectedCount, n, len(result))
					}

					// Property 2: All trees should be valid BSTs
					for i, tree := range result {
						if !isValidBSTHelper95(tree, 1, n) {
							t.Errorf("Tree %d is not a valid BST: %v",
								i, tree.ToSlice())
						}
					}

					// Property 3: All trees should contain values 1..n exactly once
					for i, tree := range result {
						if !containsValues1toN(tree, n) {
							t.Errorf("Tree %d doesn't contain all values 1..%d: %v",
								i, n, tree.ToSlice())
						}
					}

					// Property 4: All trees should be unique
					if !allTreesUnique(result) {
						t.Errorf("Found duplicate trees")
					}

					// Property 5: All trees should have exactly n nodes
					for i, tree := range result {
						nodeCount := countNodesHelper(tree)
						if nodeCount != n {
							t.Errorf("Tree %d has %d nodes, expected %d",
								i, nodeCount, n)
						}
					}
				})
			}
		})
	}
}

func BenchmarkGenerateTrees(b *testing.B) {
	// Test cases up to n=8 (maximum per constraints)
	testCases := []struct {
		name string
		n    int
	}{
		{"n=1", 1},
		{"n=2", 2},
		{"n=3", 3},
		{"n=4", 4},
		{"n=5", 5},
		{"n=6", 6},
		{"n=7", 7},
		{"n=8", 8},
	}

	implementations := []struct {
		name string
		fn   func(int) []*TreeNode
	}{
		{"generateTrees", generateTrees},
		{"generateTreesDP", generateTreesDP},
		{"generateTreesIterative", generateTreesIterative},
		{"generateTreesCatalan", generateTreesCatalan},
		{"generateTreesOptimized", generateTreesOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.n)
					}
				})
			}
		})
	}
}

func BenchmarkGenerateTreesWorstCase(b *testing.B) {
	// n=8 is the worst case within constraints
	n := 8

	b.ResetTimer()

	b.Run("generateTrees", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateTrees(n)
		}
	})

	b.Run("generateTreesDP", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateTreesDP(n)
		}
	})

	b.Run("generateTreesOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generateTreesOptimized(n)
		}
	})
}

// Helper functions (renamed to avoid conflicts)

func isValidBSTHelper95(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val < minVal || root.Val > maxVal {
		return false
	}
	return isValidBSTHelper95(root.Left, minVal, root.Val-1) &&
		isValidBSTHelper95(root.Right, root.Val+1, maxVal)
}

func containsValues1toN(root *TreeNode, n int) bool {
	values := getAllValuesHelper(root)
	if len(values) != n {
		return false
	}
	// Check we have all values 1..n
	present := make([]bool, n+1)
	for _, val := range values {
		if val < 1 || val > n {
			return false
		}
		present[val] = true
	}
	for i := 1; i <= n; i++ {
		if !present[i] {
			return false
		}
	}
	return true
}

func allTreesUnique(trees []*TreeNode) bool {
	seen := make(map[string]bool)
	for _, tree := range trees {
		// Create a canonical representation
		repr := treeToString(tree)
		if seen[repr] {
			return false
		}
		seen[repr] = true
	}
	return true
}

func treeToString(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return fmt.Sprintf("(%d %s %s)", root.Val, treeToString(root.Left), treeToString(root.Right))
}

func countNodesHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodesHelper(root.Left) + countNodesHelper(root.Right)
}

func getAllValuesHelper(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := getAllValuesHelper(root.Left)
	right := getAllValuesHelper(root.Right)
	result := append(left, root.Val)
	result = append(result, right...)
	return result
}

// Test that backtracking implementation works (it's too slow for n>4)
func TestGenerateTreesBacktracking(t *testing.T) {
	// Only test small n because backtracking is O(n!)
	for n := 1; n <= 4; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			result := generateTreesBacktracking(n)
			expected := GenerateTrees(n)
			
			// Check counts match
			if len(result) != len(expected) {
				t.Errorf("generateTreesBacktracking(%d) returned %d trees, expected %d",
					n, len(result), len(expected))
			}
			
			// Check all trees are valid
			for _, tree := range result {
				if !isValidBSTHelper95(tree, 1, n) {
					t.Errorf("Invalid BST from backtracking: %v", tree.ToSlice())
				}
			}
		})
	}
}