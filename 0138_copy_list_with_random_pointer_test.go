package leetcode

import (
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	// Helper function to create a linked list with random pointers from slices
	createList := func(values []int, randomIndices []int) *RandomListNode {
		if len(values) == 0 {
			return nil
		}
		
		// Create all nodes
		nodes := make([]*RandomListNode, len(values))
		for i, val := range values {
			nodes[i] = NewRandomListNode(val)
		}
		
		// Set next pointers
		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].Next = nodes[i+1]
		}
		
		// Set random pointers
		for i, randomIdx := range randomIndices {
			if randomIdx >= 0 && randomIdx < len(nodes) {
				nodes[i].Random = nodes[randomIdx]
			}
		}
		
		return nodes[0]
	}
	
	// Helper function to convert list to slices for comparison
	listToSlices := func(head *RandomListNode) ([]int, []int) {
		if head == nil {
			return []int{}, []int{}
		}
		
		// Create a map from node to index
		nodeToIndex := make(map[*RandomListNode]int)
		current := head
		index := 0
		for current != nil {
			nodeToIndex[current] = index
			current = current.Next
			index++
		}
		
		// Extract values and random indices
		values := []int{}
		randomIndices := []int{}
		current = head
		for current != nil {
			values = append(values, current.Val)
			
			if current.Random != nil {
				randomIndices = append(randomIndices, nodeToIndex[current.Random])
			} else {
				randomIndices = append(randomIndices, -1)
			}
			
			current = current.Next
		}
		
		return values, randomIndices
	}
	
	tests := []struct {
		name           string
		values         []int
		randomIndices  []int
		expectedValues []int
		expectedRandom []int
	}{
		{
			name:           "Example 1",
			values:         []int{7, 13, 11, 10, 1},
			randomIndices:  []int{-1, 0, 4, 2, 0},
			expectedValues: []int{7, 13, 11, 10, 1},
			expectedRandom: []int{-1, 0, 4, 2, 0},
		},
		{
			name:           "Example 2",
			values:         []int{1, 2},
			randomIndices:  []int{1, 1},
			expectedValues: []int{1, 2},
			expectedRandom: []int{1, 1},
		},
		{
			name:           "Example 3",
			values:         []int{3, 3, 3},
			randomIndices:  []int{-1, 0, -1},
			expectedValues: []int{3, 3, 3},
			expectedRandom: []int{-1, 0, -1},
		},
		{
			name:           "Empty list",
			values:         []int{},
			randomIndices:  []int{},
			expectedValues: []int{},
			expectedRandom: []int{},
		},
		{
			name:           "Single node no random",
			values:         []int{5},
			randomIndices:  []int{-1},
			expectedValues: []int{5},
			expectedRandom: []int{-1},
		},
		{
			name:           "Single node self random",
			values:         []int{5},
			randomIndices:  []int{0},
			expectedValues: []int{5},
			expectedRandom: []int{0},
		},
		{
			name:           "All random pointers null",
			values:         []int{1, 2, 3, 4, 5},
			randomIndices:  []int{-1, -1, -1, -1, -1},
			expectedValues: []int{1, 2, 3, 4, 5},
			expectedRandom: []int{-1, -1, -1, -1, -1},
		},
		{
			name:           "Random pointers in reverse order",
			values:         []int{1, 2, 3, 4, 5},
			randomIndices:  []int{4, 3, 2, 1, 0},
			expectedValues: []int{1, 2, 3, 4, 5},
			expectedRandom: []int{4, 3, 2, 1, 0},
		},
		{
			name:           "Complex random pointers",
			values:         []int{1, 2, 3, 4, 5, 6},
			randomIndices:  []int{2, -1, 5, 1, 0, 3},
			expectedValues: []int{1, 2, 3, 4, 5, 6},
			expectedRandom: []int{2, -1, 5, 1, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test hash map approach
			original := createList(tt.values, tt.randomIndices)
			copied := CopyRandomList(original)
			values, randomIndices := listToSlices(copied)
			
			if !equalIntSlices(values, tt.expectedValues) {
				t.Errorf("CopyRandomList() values = %v, expected %v", values, tt.expectedValues)
			}
			if !equalIntSlices(randomIndices, tt.expectedRandom) {
				t.Errorf("CopyRandomList() random indices = %v, expected %v", randomIndices, tt.expectedRandom)
			}
			
			// Test optimized approach
			original2 := createList(tt.values, tt.randomIndices)
			copied2 := CopyRandomListOptimized(original2)
			values2, randomIndices2 := listToSlices(copied2)
			
			if !equalIntSlices(values2, tt.expectedValues) {
				t.Errorf("CopyRandomListOptimized() values = %v, expected %v", values2, tt.expectedValues)
			}
			if !equalIntSlices(randomIndices2, tt.expectedRandom) {
				t.Errorf("CopyRandomListOptimized() random indices = %v, expected %v", randomIndices2, tt.expectedRandom)
			}
		})
	}
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkCopyRandomList(b *testing.B) {
	// Create a complex linked list for benchmarking
	createComplexList := func(n int) *RandomListNode {
		if n == 0 {
			return nil
		}
		
		// Create nodes
		nodes := make([]*RandomListNode, n)
		for i := 0; i < n; i++ {
			nodes[i] = NewRandomListNode(i + 1)
		}
		
		// Set next pointers
		for i := 0; i < n-1; i++ {
			nodes[i].Next = nodes[i+1]
		}
		
		// Set random pointers: each node points to a random node
		for i := 0; i < n; i++ {
			randomIdx := (i * 7) % n // Some pseudo-random pattern
			nodes[i].Random = nodes[randomIdx]
		}
		
		return nodes[0]
	}
	
	testCases := []struct {
		name string
		n    int
	}{
		{"Small list", 10},
		{"Medium list", 100},
		{"Large list", 1000},
		{"Very large list", 10000},
	}

	for _, tc := range testCases {
		b.Run(tc.name+"_hashmap", func(b *testing.B) {
			list := createComplexList(tc.n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				CopyRandomList(list)
			}
		})
		
		b.Run(tc.name+"_optimized", func(b *testing.B) {
			list := createComplexList(tc.n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				CopyRandomListOptimized(list)
			}
		})
	}
}