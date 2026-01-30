package trees

import (
	"reflect"
	"testing"
)

func TestConnect117(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected [][]int
	}{
		{
			name:     "Empty tree",
			input:    []interface{}{},
			expected: [][]int{},
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Complete tree",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Unbalanced tree",
			input:    []interface{}{1, 2, 3, 4, 5, nil, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 7}},
		},
		{
			name:     "Tree with missing right children",
			input:    []interface{}{1, 2, 3, 4, nil, nil, 7, 8, nil, nil, nil, nil, nil, nil, 15},
			expected: [][]int{{1}, {2, 3}, {4, 7}, {8, 15}},
		},
		{
			name:     "Tree with missing left children",
			input:    []interface{}{1, nil, 3, nil, nil, 6, 7, nil, nil, nil, nil, 12, 13, nil, 15},
			expected: [][]int{{1}, {3}, {6, 7}, {12, 13, 15}},
		},
		{
			name:     "Complex tree",
			input:    []interface{}{1, 2, 3, nil, 4, nil, 5, nil, nil, 6, 7, nil, nil, nil, nil, 8, 9},
			expected: [][]int{{1}, {2, 3}, {4, 5}, {6, 7}}, // Fixed: only 4 levels, not 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice117(tt.input)
			result := connect117(root)
			actual := traverseByLevel117(result)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("connect117() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestConnectBFS117(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected [][]int
	}{
		{
			name:     "Empty tree",
			input:    []interface{}{},
			expected: [][]int{},
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Complete tree",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Unbalanced tree",
			input:    []interface{}{1, 2, 3, 4, 5, nil, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice117(tt.input)
			result := connectBFS117(root)
			actual := traverseByLevel117(result)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("connectBFS117() = %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestConnectEdgeCases117(t *testing.T) {
	// Test that next pointers are correctly set to nil at end of levels
	t.Run("Next pointers at level ends", func(t *testing.T) {
		root := createTreeFromSlice117([]interface{}{1, 2, 3, nil, 4, nil, 5})
		result := connect117(root)

		// Level 0: node 1 should have Next = nil
		if result.Next != nil {
			t.Errorf("Root node Next should be nil, got %v", result.Next)
		}

		// Level 1: node 2 should point to node 3, node 3 should have Next = nil
		if result.Left.Next != result.Right {
			t.Errorf("Node 2 should point to Node 3")
		}
		if result.Right.Next != nil {
			t.Errorf("Node 3 Next should be nil")
		}

		// Level 2: node 4 should point to node 5, node 5 should have Next = nil
		if result.Left.Right.Next != result.Right.Right {
			t.Errorf("Node 4 should point to Node 5")
		}
		if result.Right.Right.Next != nil {
			t.Errorf("Node 5 Next should be nil")
		}
	})

	// Test tree with only left children
	t.Run("Left-skewed tree", func(t *testing.T) {
		root := createTreeFromSlice117([]interface{}{1, 2, nil, 3, nil, nil, nil, 4, nil})
		result := connect117(root)
		actual := traverseByLevel117(result)
		expected := [][]int{{1}, {2}, {3}, {4}}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Left-skewed tree: got %v, expected %v", actual, expected)
		}
	})

	// Test tree with only right children
	t.Run("Right-skewed tree", func(t *testing.T) {
		root := createTreeFromSlice117([]interface{}{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4})
		result := connect117(root)
		actual := traverseByLevel117(result)
		expected := [][]int{{1}, {2}, {3}, {4}}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Right-skewed tree: got %v, expected %v", actual, expected)
		}
	})
}

func BenchmarkConnect117(b *testing.B) {
	// Create a large complete binary tree for benchmarking
	// Tree with 1023 nodes (10 levels)
	values := make([]interface{}, 1023)
	for i := range values {
		values[i] = i + 1
	}

	root := createTreeFromSlice117(values)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		connect117(root)
	}
}

func BenchmarkConnectBFS117(b *testing.B) {
	// Create a large complete binary tree for benchmarking
	values := make([]interface{}, 1023)
	for i := range values {
		values[i] = i + 1
	}

	root := createTreeFromSlice117(values)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		connectBFS117(root)
	}
}

func TestCreateTreeFromSlice117(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		validate func(*Node117) bool
	}{
		{
			name:  "Empty slice",
			input: []interface{}{},
			validate: func(root *Node117) bool {
				return root == nil
			},
		},
		{
			name:  "Single node",
			input: []interface{}{1},
			validate: func(root *Node117) bool {
				return root != nil && root.Val == 1 && root.Left == nil && root.Right == nil
			},
		},
		{
			name:  "Tree with nulls",
			input: []interface{}{1, nil, 2, nil, nil, 3, 4},
			validate: func(root *Node117) bool {
				return root != nil &&
					root.Val == 1 &&
					root.Left == nil &&
					root.Right != nil &&
					root.Right.Val == 2 &&
					root.Right.Left != nil &&
					root.Right.Left.Val == 3 &&
					root.Right.Right != nil &&
					root.Right.Right.Val == 4
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice117(tt.input)
			if !tt.validate(root) {
				t.Errorf("createTreeFromSlice117() produced incorrect tree for input %v", tt.input)
			}
		})
	}
}