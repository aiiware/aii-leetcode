package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		// Create test lists: [1->4->5], [1->3->4], [2->6]
		list1 := &ListNode{Val: 1}
		list1.Next = &ListNode{Val: 4}
		list1.Next.Next = &ListNode{Val: 5}

		list2 := &ListNode{Val: 1}
		list2.Next = &ListNode{Val: 3}
		list2.Next.Next = &ListNode{Val: 4}

		list3 := &ListNode{Val: 2}
		list3.Next = &ListNode{Val: 6}

		lists := []*ListNode{list1, list2, list3}
		result := MergeKLists(lists)
		
		// Expected: 1->1->2->3->4->4->5->6
		expectedVals := []int{1, 1, 2, 3, 4, 4, 5, 6}
		actualVals := []int{}
		current := result
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}
		
		assert.Equal(t, expectedVals, actualVals)
	})

	t.Run("Empty input", func(t *testing.T) {
		result := MergeKLists([]*ListNode{})
		assert.Nil(t, result)
	})

	t.Run("Single list", func(t *testing.T) {
		list := &ListNode{Val: 1}
		list.Next = &ListNode{Val: 2}
		list.Next.Next = &ListNode{Val: 3}
		
		result := MergeKLists([]*ListNode{list})
		expectedVals := []int{1, 2, 3}
		actualVals := []int{}
		current := result
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}
		
		assert.Equal(t, expectedVals, actualVals)
	})

	t.Run("One empty list", func(t *testing.T) {
		list := &ListNode{}
		result := MergeKLists([]*ListNode{list})
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.Val)
		assert.Nil(t, result.Next)
	})
}

func BenchmarkMergeKLists(b *testing.B) {
	// Create test data
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	list3 := &ListNode{Val: 2}
	list3.Next = &ListNode{Val: 6}
	
	lists := []*ListNode{list1, list2, list3}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeKLists(lists)
	}
}