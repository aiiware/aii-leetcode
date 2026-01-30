package trees

import (
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		levels   int
		expected [][]int
	}{
		{
			name:   "Empty tree",
			levels: 0,
			expected: [][]int{},
		},
		{
			name:   "Single node",
			levels: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:   "Perfect tree with 2 levels",
			levels: 2,
			expected: [][]int{
				{1},
				{2, 3},
			},
		},
		{
			name:   "Perfect tree with 3 levels",
			levels: 3,
			expected: [][]int{
				{1},
				{2, 3},
				{4, 5, 6, 7},
			},
		},
		{
			name:   "Perfect tree with 4 levels",
			levels: 4,
			expected: [][]int{
				{1},
				{2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11, 12, 13, 14, 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create perfect tree
			root := NewPerfectTree(tt.levels)
			
			// Connect next pointers
			connected := connect(root)
			
			// Verify structure
			if tt.levels == 0 {
				if connected != nil {
					t.Errorf("connect() = %v, want nil", connected)
				}
				return
			}
			
			// Get level representation using next pointers
			result := connected.ToSliceWithNext()
			
			// Compare with expected
			if len(result) != len(tt.expected) {
				t.Errorf("ToSliceWithNext() levels = %v, want %v", len(result), len(tt.expected))
				return
			}
			
			for i := 0; i < len(result); i++ {
				if len(result[i]) != len(tt.expected[i]) {
					t.Errorf("ToSliceWithNext()[%d] length = %v, want %v", i, len(result[i]), len(tt.expected[i]))
					continue
				}
				
				for j := 0; j < len(result[i]); j++ {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("ToSliceWithNext()[%d][%d] = %v, want %v", i, j, result[i][j], tt.expected[i][j])
					}
				}
			}
			
			// Verify next pointers are correctly set
			verifyNextPointers(t, connected)
		})
	}
}

func TestConnectBFS(t *testing.T) {
	tests := []struct {
		name     string
		levels   int
		expected [][]int
	}{
		{
			name:   "Empty tree",
			levels: 0,
			expected: [][]int{},
		},
		{
			name:   "Perfect tree with 3 levels",
			levels: 3,
			expected: [][]int{
				{1},
				{2, 3},
				{4, 5, 6, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create perfect tree
			root := NewPerfectTree(tt.levels)
			
			// Connect next pointers using BFS
			connected := connectBFS(root)
			
			// Verify structure
			if tt.levels == 0 {
				if connected != nil {
					t.Errorf("connectBFS() = %v, want nil", connected)
				}
				return
			}
			
			// Get level representation using next pointers
			result := connected.ToSliceWithNext()
			
			// Compare with expected
			if len(result) != len(tt.expected) {
				t.Errorf("ToSliceWithNext() levels = %v, want %v", len(result), len(tt.expected))
				return
			}
			
			for i := 0; i < len(result); i++ {
				if len(result[i]) != len(tt.expected[i]) {
					t.Errorf("ToSliceWithNext()[%d] length = %v, want %v", i, len(result[i]), len(tt.expected[i]))
					continue
				}
				
				for j := 0; j < len(result[i]); j++ {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("ToSliceWithNext()[%d][%d] = %v, want %v", i, j, result[i][j], tt.expected[i][j])
					}
				}
			}
			
			// Verify next pointers are correctly set
			verifyNextPointers(t, connected)
		})
	}
}

func TestGetNextValues(t *testing.T) {
	tests := []struct {
		name     string
		levels   int
		expected []int
	}{
		{
			name:     "Single node",
			levels:   1,
			expected: []int{-1}, // Only node has no next
		},
		{
			name:     "Perfect tree with 2 levels",
			levels:   2,
			expected: []int{-1, 3, -1}, // Node 1 -> nil, Node 2 -> 3, Node 3 -> nil
		},
		{
			name:     "Perfect tree with 3 levels",
			levels:   3,
			expected: []int{-1, 3, -1, 5, 6, 7, -1}, // 7 nodes in BFS order
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create perfect tree
			root := NewPerfectTree(tt.levels)
			
			// Connect next pointers
			connected := connect(root)
			
			// Get next values
			result := connected.GetNextValues()
			
			// Compare with expected
			if len(result) != len(tt.expected) {
				t.Errorf("GetNextValues() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			
			for i := 0; i < len(result); i++ {
				if result[i] != tt.expected[i] {
					t.Errorf("GetNextValues()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestClone(t *testing.T) {
	// Create a perfect tree with 3 levels
	original := NewPerfectTree(3)
	
	// Clone the tree
	cloned := original.Clone()
	
	// Verify structure is the same
	verifyTreeStructure(t, original, cloned)
	
	// Connect next pointers in original
	connectedOriginal := connect(original)
	
	// Clone should not have next pointers set
	if cloned.Next != nil {
		t.Errorf("Clone should not copy next pointers")
	}
	
	// Connect next pointers in clone
	connectedClone := connect(cloned)
	
	// Verify both have same structure after connection
	verifyTreeStructure(t, connectedOriginal, connectedClone)
}

// Helper function to verify next pointers are correctly set
func verifyNextPointers(t *testing.T, root *Node) {
	t.Helper()
	
	if root == nil {
		return
	}
	
	// Traverse using leftmost pointer
	current := root
	for current != nil {
		node := current
		
		// Check all nodes in current level
		for node != nil {
			// If node has left child, verify it points to right child
			if node.Left != nil {
				if node.Left.Next != node.Right {
					t.Errorf("Node %d.Left.Next = %v, want %v", 
						node.Val, 
						safeVal(node.Left.Next), 
						safeVal(node.Right))
				}
			}
			
			// If node has right child and has next, verify right child points to next's left
			if node.Right != nil && node.Next != nil {
				if node.Right.Next != node.Next.Left {
					t.Errorf("Node %d.Right.Next = %v, want %v", 
						node.Val, 
						safeVal(node.Right.Next), 
						safeVal(node.Next.Left))
				}
			}
			
			// Last node in level should have nil next
			if node.Next == nil {
				// Verify this is indeed the last node by checking if we're at the end of traversal
				// (This is just for debugging, not a strict requirement)
			}
			
			node = node.Next
		}
		
		current = current.Left
	}
}

// Helper function to verify tree structure (without next pointers)
func verifyTreeStructure(t *testing.T, a, b *Node) {
	t.Helper()
	
	if a == nil && b == nil {
		return
	}
	
	if a == nil || b == nil {
		t.Errorf("Tree structure mismatch: one is nil, other is not")
		return
	}
	
	if a.Val != b.Val {
		t.Errorf("Node values differ: %v vs %v", a.Val, b.Val)
	}
	
	verifyTreeStructure(t, a.Left, b.Left)
	verifyTreeStructure(t, a.Right, b.Right)
}

// Helper function to safely get node value (returns -1 if nil)
func safeVal(node *Node) int {
	if node == nil {
		return -1
	}
	return node.Val
}

func BenchmarkConnect(b *testing.B) {
	// Create a large perfect tree for benchmarking
	root := NewPerfectTree(10) // 1023 nodes
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clone the tree for each iteration
		cloned := root.Clone()
		connect(cloned)
	}
}

func BenchmarkConnectBFS(b *testing.B) {
	// Create a large perfect tree for benchmarking
	root := NewPerfectTree(10) // 1023 nodes
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clone the tree for each iteration
		cloned := root.Clone()
		connectBFS(cloned)
	}
}