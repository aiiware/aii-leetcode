package trees

import (
	"reflect"
	"testing"
)

func TestUpsideDownBinaryTreeIterative(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Example 1: [1,2,3,4,5]",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1},
		},
		{
			name:     "Example 2: []",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "Example 3: [1]",
			input:    []interface{}{1},
			expected: []interface{}{1},
		},
		{
			name:     "Left-skewed tree: [1,2,nil,3,nil,4]",
			input:    []interface{}{1, 2, nil, 3, nil, 4},
			expected: []interface{}{4, nil, 3, nil, 2, nil, 1},
		},
		{
			name:     "Right child only: [1,nil,2,nil,3]",
			input:    []interface{}{1, nil, 2, nil, 3},
			expected: []interface{}{1, nil, 2, nil, 3},
		},
		{
			name:     "Complete tree: [1,2,3,4,5,6,7]",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1, 6, 7},
		},
		{
			name:     "Tree with only left children: [1,2,nil,3,nil,4,nil]",
			input:    []interface{}{1, 2, nil, 3, nil, 4, nil},
			expected: []interface{}{4, nil, 3, nil, 2, nil, 1},
		},
		{
			name:     "Tree with mixed structure",
			input:    []interface{}{1, 2, 3, nil, 4, 5, 6},
			expected: []interface{}{2, 3, 1, 5, 6},
		},
		{
			name:     "Larger tree",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []interface{}{8, 9, 4, nil, nil, 5, 2, 10, nil, 3, 1, nil, nil, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice(tt.input)
			result := upsideDownBinaryTreeIterative(root)
			resultSlice := treeToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("upsideDownBinaryTreeIterative(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestUpsideDownBinaryTreeRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Example 1: [1,2,3,4,5]",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1},
		},
		{
			name:     "Empty tree",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: []interface{}{1},
		},
		{
			name:     "Left-skewed tree",
			input:    []interface{}{1, 2, nil, 3, nil, 4},
			expected: []interface{}{4, nil, 3, nil, 2, nil, 1},
		},
		{
			name:     "Complete tree",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice(tt.input)
			result := upsideDownBinaryTreeRecursive(root)
			resultSlice := treeToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("upsideDownBinaryTreeRecursive(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestUpsideDownBinaryTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Main function test 1",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1},
		},
		{
			name:     "Main function test 2",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice(tt.input)
			result := upsideDownBinaryTree(root)
			resultSlice := treeToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("upsideDownBinaryTree(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestUpsideDownBinaryTreeDFS(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "DFS test 1",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1},
		},
		{
			name:     "DFS test 2 - empty",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "DFS test 3 - single node",
			input:    []interface{}{1},
			expected: []interface{}{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice(tt.input)
			result := upsideDownBinaryTreeDFS(root)
			resultSlice := treeToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("upsideDownBinaryTreeDFS(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestUpsideDownBinaryTreeStack(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Stack test 1",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1},
		},
		{
			name:     "Stack test 2 - left-skewed",
			input:    []interface{}{1, 2, nil, 3, nil, 4},
			expected: []interface{}{4, nil, 3, nil, 2, nil, 1},
		},
		{
			name:     "Stack test 3 - complete tree",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: []interface{}{4, 5, 2, nil, nil, 3, 1, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTreeFromSlice(tt.input)
			result := upsideDownBinaryTreeStack(root)
			resultSlice := treeToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("upsideDownBinaryTreeStack(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestTreeConversionHelpers(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Round trip test 1",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "Round trip test 2 - with nils",
			input:    []interface{}{1, nil, 2, nil, 3},
			expected: []interface{}{1, nil, 2, nil, 3},
		},
		{
			name:     "Round trip test 3 - complete tree",
			input:    []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Round trip test 4 - empty",
			input:    []interface{}{},
			expected: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test createTreeFromSlice and treeToSlice round trip
			root := createTreeFromSlice(tt.input)
			result := treeToSlice(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Round trip failed: createTreeFromSlice then treeToSlice(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkUpsideDownBinaryTreeIterative(b *testing.B) {
	// Create a balanced tree with 1023 nodes (height 10)
	values := make([]interface{}, 1023)
	for i := range values {
		values[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clone the tree for each iteration
		clonedRoot := createTreeFromSlice(values)
		upsideDownBinaryTreeIterative(clonedRoot)
	}
}

func BenchmarkUpsideDownBinaryTreeRecursive(b *testing.B) {
	// Create a balanced tree with 1023 nodes (height 10)
	values := make([]interface{}, 1023)
	for i := range values {
		values[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clone the tree for each iteration
		clonedRoot := createTreeFromSlice(values)
		upsideDownBinaryTreeRecursive(clonedRoot)
	}
}

func BenchmarkUpsideDownBinaryTreeDFS(b *testing.B) {
	// Create a balanced tree with 1023 nodes (height 10)
	values := make([]interface{}, 1023)
	for i := range values {
		values[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clone the tree for each iteration
		clonedRoot := createTreeFromSlice(values)
		upsideDownBinaryTreeDFS(clonedRoot)
	}
}

func BenchmarkUpsideDownBinaryTreeStack(b *testing.B) {
	// Create a balanced tree with 1023 nodes (height 10)
	values := make([]interface{}, 1023)
	for i := range values {
		values[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clone the tree for each iteration
		clonedRoot := createTreeFromSlice(values)
		upsideDownBinaryTreeStack(clonedRoot)
	}
}

func BenchmarkDifferentTreeSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run("Iterative", func(b *testing.B) {
			values := make([]interface{}, size)
			for i := range values {
				values[i] = i + 1
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				clonedRoot := createTreeFromSlice(values)
				upsideDownBinaryTreeIterative(clonedRoot)
			}
		})
		
		b.Run("Recursive", func(b *testing.B) {
			values := make([]interface{}, size)
			for i := range values {
				values[i] = i + 1
			}
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				clonedRoot := createTreeFromSlice(values)
				upsideDownBinaryTreeRecursive(clonedRoot)
			}
		})
	}
}