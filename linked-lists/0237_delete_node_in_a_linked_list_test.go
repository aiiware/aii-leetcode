package linkedlists

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		nodeVal  int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{4, 5, 1, 9},
			nodeVal:  5,
			expected: []int{4, 1, 9},
		},
		{
			name:     "Example 2",
			input:    []int{4, 5, 1, 9},
			nodeVal:  1,
			expected: []int{4, 5, 9},
		},
		{
			name:     "Delete second node",
			input:    []int{1, 2, 3, 4},
			nodeVal:  2,
			expected: []int{1, 3, 4},
		},
		{
			name:     "Delete first node (not tail)",
			input:    []int{1, 2, 3},
			nodeVal:  1,
			expected: []int{2, 3},
		},
		{
			name:     "Delete middle node",
			input:    []int{1, 2, 3, 4, 5},
			nodeVal:  3,
			expected: []int{1, 2, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create linked list
			head := createLinkedList(tt.input)
			
			// Find the node to delete
			var nodeToDelete *ListNode
			current := head
			for current != nil {
				if current.Val == tt.nodeVal {
					nodeToDelete = current
					break
				}
				current = current.Next
			}
			
			if nodeToDelete == nil {
				t.Fatalf("Node with value %d not found in input list", tt.nodeVal)
			}
			
			// Delete the node
			DeleteNode(nodeToDelete)
			
			// Verify the resulting list
			result := linkedListToSlice(head)
			if !slicesEqual237(result, tt.expected) {
				t.Errorf("DeleteNode() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	
	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func slicesEqual237(a, b []int) bool {
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