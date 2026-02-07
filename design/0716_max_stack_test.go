package design

import (
	"testing"
)

func TestMaxStack(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		stack := ConstructorMaxStack()
		
		// Test empty stack
		if got := stack.Top(); got != -1 {
			t.Errorf("Top() on empty stack = %d, want -1", got)
		}
		if got := stack.Pop(); got != -1 {
			t.Errorf("Pop() on empty stack = %d, want -1", got)
		}
		if got := stack.PeekMax(); got != -1 {
			t.Errorf("PeekMax() on empty stack = %d, want -1", got)
		}
		if got := stack.PopMax(); got != -1 {
			t.Errorf("PopMax() on empty stack = %d, want -1", got)
		}
		
		// Push elements
		stack.Push(5)
		stack.Push(1)
		stack.Push(5)
		
		// Test Top
		if got := stack.Top(); got != 5 {
			t.Errorf("Top() = %d, want 5", got)
		}
		
		// Test PeekMax
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() = %d, want 5", got)
		}
		
		// Test Pop
		if got := stack.Pop(); got != 5 {
			t.Errorf("Pop() = %d, want 5", got)
		}
		
		// Test Top after pop
		if got := stack.Top(); got != 1 {
			t.Errorf("Top() after pop = %d, want 1", got)
		}
		
		// Test PeekMax after pop
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() after pop = %d, want 5", got)
		}
	})
	
	t.Run("PopMax operations", func(t *testing.T) {
		stack := ConstructorMaxStack()
		
		stack.Push(5)
		stack.Push(1)
		stack.Push(5)
		
		// PopMax should return 5
		if got := stack.PopMax(); got != 5 {
			t.Errorf("PopMax() = %d, want 5", got)
		}
		
		// Top should now be 1
		if got := stack.Top(); got != 1 {
			t.Errorf("Top() after PopMax = %d, want 1", got)
		}
		
		// PeekMax should now be 5 (other 5 is still there)
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() after PopMax = %d, want 5", got)
		}
		
		// PopMax again should return 5
		if got := stack.PopMax(); got != 5 {
			t.Errorf("Second PopMax() = %d, want 5", got)
		}
		
		// PeekMax should now be 1
		if got := stack.PeekMax(); got != 1 {
			t.Errorf("PeekMax() after second PopMax = %d, want 1", got)
		}
	})
	
	t.Run("Complex sequence", func(t *testing.T) {
		stack := ConstructorMaxStack()
		
		// Push sequence: 5, 1, 5, 3, 7, 2
		stack.Push(5)
		stack.Push(1)
		stack.Push(5)
		stack.Push(3)
		stack.Push(7)
		stack.Push(2)
		
		// Current stack: [5, 1, 5, 3, 7, 2]
		// Top should be 2
		if got := stack.Top(); got != 2 {
			t.Errorf("Top() = %d, want 2", got)
		}
		
		// PeekMax should be 7
		if got := stack.PeekMax(); got != 7 {
			t.Errorf("PeekMax() = %d, want 7", got)
		}
		
		// PopMax should return 7
		if got := stack.PopMax(); got != 7 {
			t.Errorf("PopMax() = %d, want 7", got)
		}
		
		// Stack should now be: [5, 1, 5, 3, 2]
		// Top should be 2
		if got := stack.Top(); got != 2 {
			t.Errorf("Top() after PopMax = %d, want 2", got)
		}
		
		// PeekMax should be 5
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() after PopMax = %d, want 5", got)
		}
		
		// Pop should return 2
		if got := stack.Pop(); got != 2 {
			t.Errorf("Pop() = %d, want 2", got)
		}
		
		// Stack should now be: [5, 1, 5, 3]
		// Top should be 3
		if got := stack.Top(); got != 3 {
			t.Errorf("Top() after Pop = %d, want 3", got)
		}
		
		// PopMax should return 5
		if got := stack.PopMax(); got != 5 {
			t.Errorf("Second PopMax() = %d, want 5", got)
		}
		
		// Stack should now be: [5, 1, 3]
		// Top should be 3
		if got := stack.Top(); got != 3 {
			t.Errorf("Top() after second PopMax = %d, want 3", got)
		}
		
		// PeekMax should be 5
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() after second PopMax = %d, want 5", got)
		}
	})
	
	t.Run("Edge cases", func(t *testing.T) {
		stack := ConstructorMaxStack()
		
		// Single element
		stack.Push(10)
		
		if got := stack.Top(); got != 10 {
			t.Errorf("Top() single element = %d, want 10", got)
		}
		if got := stack.PeekMax(); got != 10 {
			t.Errorf("PeekMax() single element = %d, want 10", got)
		}
		if got := stack.PopMax(); got != 10 {
			t.Errorf("PopMax() single element = %d, want 10", got)
		}
		
		// Stack should be empty
		if got := stack.Top(); got != -1 {
			t.Errorf("Top() after removing all = %d, want -1", got)
		}
		
		// All equal elements
		stack.Push(3)
		stack.Push(3)
		stack.Push(3)
		
		if got := stack.PeekMax(); got != 3 {
			t.Errorf("PeekMax() all equal = %d, want 3", got)
		}
		
		// PopMax should remove the most recent 3
		if got := stack.PopMax(); got != 3 {
			t.Errorf("PopMax() all equal = %d, want 3", got)
		}
		
		// Top should still be 3 (two left)
		if got := stack.Top(); got != 3 {
			t.Errorf("Top() after PopMax equal = %d, want 3", got)
		}
	})
	
	t.Run("Descending values", func(t *testing.T) {
		stack := ConstructorMaxStack()
		
		stack.Push(5)
		stack.Push(4)
		stack.Push(3)
		stack.Push(2)
		stack.Push(1)
		
		// Max should always be 5
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() descending = %d, want 5", got)
		}
		
		// PopMax should return 5
		if got := stack.PopMax(); got != 5 {
			t.Errorf("PopMax() descending = %d, want 5", got)
		}
		
		// Now max should be 4
		if got := stack.PeekMax(); got != 4 {
			t.Errorf("PeekMax() after PopMax descending = %d, want 4", got)
		}
	})
	
	t.Run("Ascending values", func(t *testing.T) {
		stack := ConstructorMaxStack()
		
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		stack.Push(5)
		
		// Max should be 5
		if got := stack.PeekMax(); got != 5 {
			t.Errorf("PeekMax() ascending = %d, want 5", got)
		}
		
		// Top should be 5
		if got := stack.Top(); got != 5 {
			t.Errorf("Top() ascending = %d, want 5", got)
		}
		
		// Pop should return 5
		if got := stack.Pop(); got != 5 {
			t.Errorf("Pop() ascending = %d, want 5", got)
		}
		
		// Now max should be 4
		if got := stack.PeekMax(); got != 4 {
			t.Errorf("PeekMax() after Pop ascending = %d, want 4", got)
		}
		
		// PopMax should return 4
		if got := stack.PopMax(); got != 4 {
			t.Errorf("PopMax() ascending = %d, want 4", got)
		}
		
		// Now max should be 3
		if got := stack.PeekMax(); got != 3 {
			t.Errorf("PeekMax() after PopMax ascending = %d, want 3", got)
		}
	})
}