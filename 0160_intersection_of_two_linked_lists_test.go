package leetcode

import (
	"testing"
)

// Helper function to create intersecting linked lists
func createIntersectingLists(listA, listB []int, intersectVal int, skipA, skipB int) (*ListNode, *ListNode, *ListNode) {
	if intersectVal == 0 {
		// No intersection
		return NewListFromSlice(listA), NewListFromSlice(listB), nil
	}

	// Create list A up to intersection point
	var headA, tailA *ListNode
	if len(listA) > 0 {
		headA = &ListNode{Val: listA[0]}
		current := headA
		for i := 1; i < skipA; i++ {
			current.Next = &ListNode{Val: listA[i]}
			current = current.Next
		}
		tailA = current
	}

	// Create list B up to intersection point
	var headB, tailB *ListNode
	if len(listB) > 0 {
		headB = &ListNode{Val: listB[0]}
		current := headB
		for i := 1; i < skipB; i++ {
			current.Next = &ListNode{Val: listB[i]}
			current = current.Next
		}
		tailB = current
	}

	// Create intersection node
	var intersectNode *ListNode
	if intersectVal > 0 {
		intersectNode = &ListNode{Val: intersectVal}
		if tailA != nil {
			tailA.Next = intersectNode
		} else {
			headA = intersectNode
		}
		if tailB != nil {
			tailB.Next = intersectNode
		} else {
			headB = intersectNode
		}

		// Add remaining nodes after intersection
		current := intersectNode
		for i := skipA; i < len(listA); i++ {
			current.Next = &ListNode{Val: listA[i]}
			current = current.Next
		}
	}

	return headA, headB, intersectNode
}

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name        string
		listA       []int
		listB       []int
		intersectVal int
		skipA       int
		skipB       int
		expect      *ListNode
	}{
		{
			name:        "Example 1: Intersection at node with value 8",
			listA:       []int{4, 1, 8, 4, 5},
			listB:       []int{5, 6, 1, 8, 4, 5},
			intersectVal: 8,
			skipA:       2,
			skipB:       3,
			expect:      &ListNode{Val: 8},
		},
		{
			name:        "Example 2: Intersection at node with value 2",
			listA:       []int{1, 9, 1, 2, 4},
			listB:       []int{3, 2, 4},
			intersectVal: 2,
			skipA:       3,
			skipB:       1,
			expect:      &ListNode{Val: 2},
		},
		{
			name:        "Example 3: No intersection",
			listA:       []int{2, 6, 4},
			listB:       []int{1, 5},
			intersectVal: 0,
			skipA:       3,
			skipB:       2,
			expect:      nil,
		},
		{
			name:        "Empty list A",
			listA:       []int{},
			listB:       []int{1, 2, 3},
			intersectVal: 0,
			skipA:       0,
			skipB:       3,
			expect:      nil,
		},
		{
			name:        "Empty list B",
			listA:       []int{1, 2, 3},
			listB:       []int{},
			intersectVal: 0,
			skipA:       3,
			skipB:       0,
			expect:      nil,
		},
		{
			name:        "Both lists empty",
			listA:       []int{},
			listB:       []int{},
			intersectVal: 0,
			skipA:       0,
			skipB:       0,
			expect:      nil,
		},
		{
			name:        "Intersection at first node",
			listA:       []int{1, 2, 3},
			listB:       []int{1, 2, 3},
			intersectVal: 1,
			skipA:       0,
			skipB:       0,
			expect:      &ListNode{Val: 1},
		},
		{
			name:        "Intersection at last node",
			listA:       []int{1, 2, 3, 4, 5},
			listB:       []int{6, 7, 5},
			intersectVal: 5,
			skipA:       4,
			skipB:       2,
			expect:      &ListNode{Val: 5},
		},
		{
			name:        "Same list (identical)",
			listA:       []int{1, 2, 3, 4, 5},
			listB:       []int{1, 2, 3, 4, 5},
			intersectVal: 1,
			skipA:       0,
			skipB:       0,
			expect:      &ListNode{Val: 1},
		},
		{
			name:        "Long lists with intersection in middle",
			listA:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			listB:       []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10},
			intersectVal: 6,
			skipA:       5,
			skipB:       5,
			expect:      &ListNode{Val: 6},
		},
		{
			name:        "Single node lists intersecting",
			listA:       []int{1},
			listB:       []int{1},
			intersectVal: 1,
			skipA:       0,
			skipB:       0,
			expect:      &ListNode{Val: 1},
		},
		{
			name:        "Single node lists not intersecting",
			listA:       []int{1},
			listB:       []int{2},
			intersectVal: 0,
			skipA:       1,
			skipB:       1,
			expect:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expectedNode := createIntersectingLists(tt.listA, tt.listB, tt.intersectVal, tt.skipA, tt.skipB)
			
			result := getIntersectionNode(headA, headB)
			
			// Check if result matches expected
			if expectedNode == nil {
				if result != nil {
					t.Errorf("getIntersectionNode() = %v, want nil", result)
				}
			} else {
				if result == nil {
					t.Errorf("getIntersectionNode() = nil, want node with value %d", expectedNode.Val)
				} else if result.Val != expectedNode.Val {
					t.Errorf("getIntersectionNode() = node with value %d, want node with value %d", result.Val, expectedNode.Val)
				}
				
				// Verify that the intersection node is actually the same node (not just same value)
				// by checking that the rest of the lists match
				if result != nil && expectedNode != nil {
					// Check a few nodes after intersection to ensure it's the same node
					currentResult := result
					currentExpected := expectedNode
					for i := 0; i < 3 && currentResult != nil && currentExpected != nil; i++ {
						if currentResult.Val != currentExpected.Val {
							t.Errorf("Nodes after intersection don't match at position %d: got %d, want %d", 
								i, currentResult.Val, currentExpected.Val)
							break
						}
						currentResult = currentResult.Next
						currentExpected = currentExpected.Next
					}
				}
			}
		})
	}
}

func TestGetIntersectionNodeTwoPointers(t *testing.T) {
	// Test the specific implementation
	tests := []struct {
		name        string
		listA       []int
		listB       []int
		intersectVal int
		skipA       int
		skipB       int
	}{
		{
			name:        "Standard intersection",
			listA:       []int{4, 1, 8, 4, 5},
			listB:       []int{5, 6, 1, 8, 4, 5},
			intersectVal: 8,
			skipA:       2,
			skipB:       3,
		},
		{
			name:        "No intersection",
			listA:       []int{2, 6, 4},
			listB:       []int{1, 5},
			intersectVal: 0,
			skipA:       3,
			skipB:       2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expectedNode := createIntersectingLists(tt.listA, tt.listB, tt.intersectVal, tt.skipA, tt.skipB)
			
			result := getIntersectionNodeTwoPointers(headA, headB)
			
			if expectedNode == nil {
				if result != nil {
					t.Errorf("getIntersectionNodeTwoPointers() = %v, want nil", result)
				}
			} else {
				if result == nil {
					t.Errorf("getIntersectionNodeTwoPointers() = nil, want node with value %d", expectedNode.Val)
				} else if result.Val != expectedNode.Val {
					t.Errorf("getIntersectionNodeTwoPointers() = node with value %d, want node with value %d", result.Val, expectedNode.Val)
				}
			}
		})
	}
}

func TestGetIntersectionNodeLength(t *testing.T) {
	// Test the length-based implementation
	tests := []struct {
		name        string
		listA       []int
		listB       []int
		intersectVal int
		skipA       int
		skipB       int
	}{
		{
			name:        "List A longer",
			listA:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			listB:       []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10},
			intersectVal: 6,
			skipA:       5,
			skipB:       5,
		},
		{
			name:        "List B longer",
			listA:       []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10},
			listB:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			intersectVal: 6,
			skipA:       5,
			skipB:       5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expectedNode := createIntersectingLists(tt.listA, tt.listB, tt.intersectVal, tt.skipA, tt.skipB)
			
			result := getIntersectionNodeLength(headA, headB)
			
			if expectedNode == nil {
				if result != nil {
					t.Errorf("getIntersectionNodeLength() = %v, want nil", result)
				}
			} else {
				if result == nil {
					t.Errorf("getIntersectionNodeLength() = nil, want node with value %d", expectedNode.Val)
				} else if result.Val != expectedNode.Val {
					t.Errorf("getIntersectionNodeLength() = node with value %d, want node with value %d", result.Val, expectedNode.Val)
				}
			}
		})
	}
}

func TestGetIntersectionNodeHash(t *testing.T) {
	// Test the hash-based implementation
	tests := []struct {
		name        string
		listA       []int
		listB       []int
		intersectVal int
		skipA       int
		skipB       int
	}{
		{
			name:        "Hash implementation test",
			listA:       []int{4, 1, 8, 4, 5},
			listB:       []int{5, 6, 1, 8, 4, 5},
			intersectVal: 8,
			skipA:       2,
			skipB:       3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expectedNode := createIntersectingLists(tt.listA, tt.listB, tt.intersectVal, tt.skipA, tt.skipB)
			
			result := getIntersectionNodeHash(headA, headB)
			
			if expectedNode == nil {
				if result != nil {
					t.Errorf("getIntersectionNodeHash() = %v, want nil", result)
				}
			} else {
				if result == nil {
					t.Errorf("getIntersectionNodeHash() = nil, want node with value %d", expectedNode.Val)
				} else if result.Val != expectedNode.Val {
					t.Errorf("getIntersectionNodeHash() = node with value %d, want node with value %d", result.Val, expectedNode.Val)
				}
			}
		})
	}
}

func TestGetIntersectionNodeCycle(t *testing.T) {
	// Test the cycle detection implementation
	tests := []struct {
		name        string
		listA       []int
		listB       []int
		intersectVal int
		skipA       int
		skipB       int
	}{
		{
			name:        "Cycle detection test",
			listA:       []int{4, 1, 8, 4, 5},
			listB:       []int{5, 6, 1, 8, 4, 5},
			intersectVal: 8,
			skipA:       2,
			skipB:       3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB, expectedNode := createIntersectingLists(tt.listA, tt.listB, tt.intersectVal, tt.skipA, tt.skipB)
			
			result := getIntersectionNodeCycle(headA, headB)
			
			if expectedNode == nil {
				if result != nil {
					t.Errorf("getIntersectionNodeCycle() = %v, want nil", result)
				}
			} else {
				if result == nil {
					t.Errorf("getIntersectionNodeCycle() = nil, want node with value %d", expectedNode.Val)
				} else if result.Val != expectedNode.Val {
					t.Errorf("getIntersectionNodeCycle() = node with value %d, want node with value %d", result.Val, expectedNode.Val)
				}
			}
		})
	}
}

func BenchmarkGetIntersectionNode(b *testing.B) {
	// Create test lists for benchmarking
	listA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listB := []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10}
	headA, headB, _ := createIntersectingLists(listA, listB, 6, 5, 5)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNode(headA, headB)
	}
}

func BenchmarkGetIntersectionNodeTwoPointers(b *testing.B) {
	listA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listB := []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10}
	headA, headB, _ := createIntersectingLists(listA, listB, 6, 5, 5)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNodeTwoPointers(headA, headB)
	}
}

func BenchmarkGetIntersectionNodeLength(b *testing.B) {
	listA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listB := []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10}
	headA, headB, _ := createIntersectingLists(listA, listB, 6, 5, 5)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNodeLength(headA, headB)
	}
}

func BenchmarkGetIntersectionNodeHash(b *testing.B) {
	listA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listB := []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10}
	headA, headB, _ := createIntersectingLists(listA, listB, 6, 5, 5)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNodeHash(headA, headB)
	}
}

func BenchmarkGetIntersectionNodeCycle(b *testing.B) {
	listA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listB := []int{11, 12, 13, 14, 15, 6, 7, 8, 9, 10}
	headA, headB, _ := createIntersectingLists(listA, listB, 6, 5, 5)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIntersectionNodeCycle(headA, headB)
	}
}