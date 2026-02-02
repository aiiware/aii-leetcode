package design

import (
	"testing"

	"leetcode/utils"
)

func TestBSTIterator(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		commands []string
		args     [][]interface{}
		expected []interface{}
	}{
		{
			name: "Example 1",
			root: []*int{utils.IntPtr(7), utils.IntPtr(3), utils.IntPtr(15), nil, nil, utils.IntPtr(9), utils.IntPtr(20)},
			commands: []string{
				"BSTIterator", "next", "next", "hasNext", "next",
				"hasNext", "next", "hasNext", "next", "hasNext",
			},
			args: [][]interface{}{
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
			},
			expected: []interface{}{
				nil, 3, 7, true, 9, true, 15, true, 20, false,
			},
		},
		{
			name: "Single node tree",
			root: []*int{utils.IntPtr(5)},
			commands: []string{
				"BSTIterator", "hasNext", "next", "hasNext",
			},
			args: [][]interface{}{
				{},
				{},
				{},
				{},
			},
			expected: []interface{}{
				nil, true, 5, false,
			},
		},
		{
			name: "Left-skewed tree",
			root: []*int{utils.IntPtr(5), utils.IntPtr(3), nil, utils.IntPtr(1), nil},
			commands: []string{
				"BSTIterator", "next", "next", "next", "hasNext",
			},
			args: [][]interface{}{
				{},
				{},
				{},
				{},
				{},
			},
			expected: []interface{}{
				nil, 1, 3, 5, false,
			},
		},
		{
			name: "Right-skewed tree",
			root: []*int{utils.IntPtr(1), nil, utils.IntPtr(3), nil, utils.IntPtr(5)},
			commands: []string{
				"BSTIterator", "next", "next", "next", "hasNext",
			},
			args: [][]interface{}{
				{},
				{},
				{},
				{},
				{},
			},
			expected: []interface{}{
				nil, 1, 3, 5, false,
			},
		},
		{
			name: "Complex tree",
			root: []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(3), utils.IntPtr(7), utils.IntPtr(13), utils.IntPtr(18), utils.IntPtr(1), nil, utils.IntPtr(6)},
			commands: []string{
				"BSTIterator", "next", "next", "next", "next",
				"next", "next", "next", "next", "hasNext",
			},
			args: [][]interface{}{
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
				{},
			},
			expected: []interface{}{
				nil, 1, 3, 5, 6, 7, 10, 13, 15, true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := utils.NewTreeFromSlice(tt.root)
			var iterator BSTIterator

			for i, cmd := range tt.commands {
				switch cmd {
				case "BSTIterator":
					iterator = BSTIteratorConstructor(root)
				case "next":
					result := iterator.Next()
					expected := tt.expected[i].(int)
					if result != expected {
						t.Errorf("Next() = %v, want %v", result, expected)
					}
				case "hasNext":
					result := iterator.HasNext()
					expected := tt.expected[i].(bool)
					if result != expected {
						t.Errorf("HasNext() = %v, want %v", result, expected)
					}
				}
			}
		})
	}
}

func TestBSTIterator_EdgeCases(t *testing.T) {
	// Test with nil root (should handle gracefully)
	root := utils.NewTreeFromSlice([]*int{})
	iterator := BSTIteratorConstructor(root)
	
	if iterator.HasNext() {
		t.Error("HasNext() should return false for empty tree")
	}
}

func TestBSTIterator_MultipleCalls(t *testing.T) {
	root := utils.NewTreeFromSlice([]*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(7)})
	iterator := BSTIteratorConstructor(root)

	// Test sequence of calls
	expectedSequence := []int{1, 2, 3, 4, 5, 6, 7}
	for _, expected := range expectedSequence {
		if !iterator.HasNext() {
			t.Errorf("HasNext() should return true before getting %v", expected)
		}
		val := iterator.Next()
		if val != expected {
			t.Errorf("Next() = %v, want %v", val, expected)
		}
	}

	if iterator.HasNext() {
		t.Error("HasNext() should return false after all elements")
	}
}

func BenchmarkBSTIterator(b *testing.B) {
	// Create a balanced tree with 1000 nodes
	// For simplicity, we'll create a smaller tree for benchmarking
	root := utils.NewTreeFromSlice([]*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(3), utils.IntPtr(7), utils.IntPtr(13), utils.IntPtr(18), utils.IntPtr(1), utils.IntPtr(4), utils.IntPtr(6), utils.IntPtr(8), utils.IntPtr(12), utils.IntPtr(14), utils.IntPtr(16), utils.IntPtr(20)})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iterator := BSTIteratorConstructor(root)
		for iterator.HasNext() {
			iterator.Next()
		}
	}
}