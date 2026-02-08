package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			expected: 7,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			expected: 9,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name: "Single node",
			root: &TreeNode{
				Val: 5,
			},
			expected: 5,
		},
		{
			name: "Two nodes",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RobTree(tt.root)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkRobTree(b *testing.B) {
	// Create a sample tree for benchmarking
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RobTree(root)
	}
}