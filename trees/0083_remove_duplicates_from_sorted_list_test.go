package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		// Input: [1,1,2]
		// Output: [1,2]
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}
		
		result := DeleteDuplicates(head)
		
		expectedVals := []int{1, 2}
		actualVals := []int{}
		current := result
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}
		
		assert.Equal(t, expectedVals, actualVals)
	})

	t.Run("Example 2", func(t *testing.T) {
		// Input: [1,1,2,3,3]
		// Output: [1,2,3]
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 2}
		head.Next.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next.Next = &ListNode{Val: 3}
		
		result := DeleteDuplicates(head)
		
		expectedVals := []int{1, 2, 3}
		actualVals := []int{}
		current := result
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}
		
		assert.Equal(t, expectedVals, actualVals)
	})

	t.Run("Empty list", func(t *testing.T) {
		result := DeleteDuplicates(nil)
		assert.Nil(t, result)
	})

	t.Run("Single node", func(t *testing.T) {
		head := &ListNode{Val: 1}
		result := DeleteDuplicates(head)
		
		expectedVals := []int{1}
		actualVals := []int{}
		current := result
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}
		
		assert.Equal(t, expectedVals, actualVals)
	})

	t.Run("All duplicates", func(t *testing.T) {
		// Input: [1,1,1,1]
		// Output: [1]
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 1}
		head.Next.Next = &ListNode{Val: 1}
		head.Next.Next.Next = &ListNode{Val: 1}
		
		result := DeleteDuplicates(head)
		
		expectedVals := []int{1}
		actualVals := []int{}
		current := result
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}
		
		assert.Equal(t, expectedVals, actualVals)
	})
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	// Create test data
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeleteDuplicates(head)
	}
}