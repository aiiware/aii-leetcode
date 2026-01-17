package leetcode

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		
		if min := minStack.GetMin(); min != -3 {
			t.Errorf("GetMin() = %d, expected -3", min)
		}
		
		minStack.Pop()
		
		if top := minStack.Top(); top != 0 {
			t.Errorf("Top() = %d, expected 0", top)
		}
		
		if min := minStack.GetMin(); min != -2 {
			t.Errorf("GetMin() = %d, expected -2", min)
		}
	})

	t.Run("Duplicate minimum values", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(2)
		minStack.Push(0)
		minStack.Push(3)
		minStack.Push(0)
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() = %d, expected 0", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() after first pop = %d, expected 0", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() after second pop = %d, expected 0", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() after third pop = %d, expected 2", min)
		}
	})

	t.Run("Ascending values", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(1)
		minStack.Push(2)
		minStack.Push(3)
		minStack.Push(4)
		minStack.Push(5)
		
		if min := minStack.GetMin(); min != 1 {
			t.Errorf("GetMin() = %d, expected 1", min)
		}
		
		minStack.Pop()
		minStack.Pop()
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 1 {
			t.Errorf("GetMin() after pops = %d, expected 1", min)
		}
	})

	t.Run("Descending values", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(5)
		minStack.Push(4)
		minStack.Push(3)
		minStack.Push(2)
		minStack.Push(1)
		
		if min := minStack.GetMin(); min != 1 {
			t.Errorf("GetMin() = %d, expected 1", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() after pop = %d, expected 2", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 3 {
			t.Errorf("GetMin() after second pop = %d, expected 3", min)
		}
	})

	t.Run("Negative numbers", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(-5)
		minStack.Push(-3)
		minStack.Push(-7)
		minStack.Push(-2)
		minStack.Push(-1)
		
		if min := minStack.GetMin(); min != -7 {
			t.Errorf("GetMin() = %d, expected -7", min)
		}
		
		minStack.Pop()
		minStack.Pop()
		
		if min := minStack.GetMin(); min != -7 {
			t.Errorf("GetMin() after pops = %d, expected -7", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != -5 {
			t.Errorf("GetMin() after third pop = %d, expected -5", min)
		}
	})

	t.Run("Mixed positive and negative", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(3)
		minStack.Push(-2)
		minStack.Push(5)
		minStack.Push(-4)
		minStack.Push(1)
		
		if min := minStack.GetMin(); min != -4 {
			t.Errorf("GetMin() = %d, expected -4", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != -4 {
			t.Errorf("GetMin() after first pop = %d, expected -4", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != -2 {
			t.Errorf("GetMin() after second pop = %d, expected -2", min)
		}
	})

	t.Run("Single element", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(42)
		
		if top := minStack.Top(); top != 42 {
			t.Errorf("Top() = %d, expected 42", top)
		}
		
		if min := minStack.GetMin(); min != 42 {
			t.Errorf("GetMin() = %d, expected 42", min)
		}
	})

	t.Run("Empty stack operations (should not panic)", func(t *testing.T) {
		minStack := MinStackConstructor()
		
		// These should not panic
		minStack.Pop()
		_ = minStack.Top()
		_ = minStack.GetMin()
		
		// Now add elements and verify
		minStack.Push(10)
		
		if top := minStack.Top(); top != 10 {
			t.Errorf("Top() = %d, expected 10", top)
		}
		
		if min := minStack.GetMin(); min != 10 {
			t.Errorf("GetMin() = %d, expected 10", min)
		}
	})
}

func TestMinStackSingleStack(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		minStack := MinStackSingleStackConstructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		
		if min := minStack.GetMin(); min != -3 {
			t.Errorf("GetMin() = %d, expected -3", min)
		}
		
		minStack.Pop()
		
		if top := minStack.Top(); top != 0 {
			t.Errorf("Top() = %d, expected 0", top)
		}
		
		if min := minStack.GetMin(); min != -2 {
			t.Errorf("GetMin() = %d, expected -2", min)
		}
	})

	t.Run("Duplicate minimum values", func(t *testing.T) {
		minStack := MinStackSingleStackConstructor()
		minStack.Push(2)
		minStack.Push(0)
		minStack.Push(3)
		minStack.Push(0)
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() = %d, expected 0", min)
		}
		
		minStack.Pop()
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() after first pop = %d, expected 0", min)
		}
	})
}

func TestMinStackLinkedList(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		minStack := MinStackLinkedListConstructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		
		if min := minStack.GetMin(); min != -3 {
			t.Errorf("GetMin() = %d, expected -3", min)
		}
		
		minStack.Pop()
		
		if top := minStack.Top(); top != 0 {
			t.Errorf("Top() = %d, expected 0", top)
		}
		
		if min := minStack.GetMin(); min != -2 {
			t.Errorf("GetMin() = %d, expected -2", min)
		}
	})

	t.Run("Multiple operations", func(t *testing.T) {
		minStack := MinStackLinkedListConstructor()
		minStack.Push(5)
		minStack.Push(3)
		minStack.Push(7)
		minStack.Push(2)
		minStack.Push(8)
		
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() = %d, expected 2", min)
		}
		
		minStack.Pop() // pop 8
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() after first pop = %d, expected 2", min)
		}
		
		minStack.Pop() // pop 2
		if min := minStack.GetMin(); min != 3 {
			t.Errorf("GetMin() after second pop = %d, expected 3", min)
		}
		
		minStack.Pop() // pop 7
		if min := minStack.GetMin(); min != 3 {
			t.Errorf("GetMin() after third pop = %d, expected 3", min)
		}
	})
}

func TestMinStackOptimized(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		minStack := MinStackOptimizedConstructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		
		if min := minStack.GetMin(); min != -3 {
			t.Errorf("GetMin() = %d, expected -3", min)
		}
		
		minStack.Pop()
		
		if top := minStack.Top(); top != 0 {
			t.Errorf("Top() = %d, expected 0", top)
		}
		
		if min := minStack.GetMin(); min != -2 {
			t.Errorf("GetMin() = %d, expected -2", min)
		}
	})

	t.Run("Duplicate minimum values", func(t *testing.T) {
		minStack := MinStackOptimizedConstructor()
		minStack.Push(2)
		minStack.Push(0)
		minStack.Push(3)
		minStack.Push(0)
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() = %d, expected 0", min)
		}
		
		minStack.Pop() // pop 0
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() after first pop = %d, expected 0", min)
		}
		
		minStack.Pop() // pop 3
		
		if min := minStack.GetMin(); min != 0 {
			t.Errorf("GetMin() after second pop = %d, expected 0", min)
		}
		
		minStack.Pop() // pop 0
		
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() after third pop = %d, expected 2", min)
		}
	})

	t.Run("Space optimization test", func(t *testing.T) {
		minStack := MinStackOptimizedConstructor()
		// Push values where min only changes occasionally
		minStack.Push(5)
		minStack.Push(6)
		minStack.Push(7)
		minStack.Push(8)
		minStack.Push(3) // New min
		minStack.Push(4)
		minStack.Push(9)
		minStack.Push(2) // New min
		minStack.Push(10)
		
		// minStack should only have [5, 3, 2] (not 8 elements)
		// We can't directly check internal state, but we can verify operations work
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() = %d, expected 2", min)
		}
		
		minStack.Pop() // pop 10
		if min := minStack.GetMin(); min != 2 {
			t.Errorf("GetMin() after pop 10 = %d, expected 2", min)
		}
		
		minStack.Pop() // pop 2
		if min := minStack.GetMin(); min != 3 {
			t.Errorf("GetMin() after pop 2 = %d, expected 3", min)
		}
		
		minStack.Pop() // pop 9
		minStack.Pop() // pop 4
		if min := minStack.GetMin(); min != 3 {
			t.Errorf("GetMin() after pops = %d, expected 3", min)
		}
		
		minStack.Pop() // pop 3
		if min := minStack.GetMin(); min != 5 {
			t.Errorf("GetMin() after pop 3 = %d, expected 5", min)
		}
	})
}

func BenchmarkMinStackPushPop(b *testing.B) {
	b.Run("StandardTwoStacks", func(b *testing.B) {
		minStack := MinStackConstructor()
		for i := 0; i < b.N; i++ {
			minStack.Push(i)
			minStack.Push(i - 1)
			minStack.GetMin()
			minStack.Pop()
			minStack.Top()
		}
	})

	b.Run("SingleStack", func(b *testing.B) {
		minStack := MinStackSingleStackConstructor()
		for i := 0; i < b.N; i++ {
			minStack.Push(i)
			minStack.Push(i - 1)
			minStack.GetMin()
			minStack.Pop()
			minStack.Top()
		}
	})

	b.Run("LinkedList", func(b *testing.B) {
		minStack := MinStackLinkedListConstructor()
		for i := 0; i < b.N; i++ {
			minStack.Push(i)
			minStack.Push(i - 1)
			minStack.GetMin()
			minStack.Pop()
			minStack.Top()
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		minStack := MinStackOptimizedConstructor()
		for i := 0; i < b.N; i++ {
			minStack.Push(i)
			minStack.Push(i - 1)
			minStack.GetMin()
			minStack.Pop()
			minStack.Top()
		}
	})
}

func BenchmarkMinStackGetMin(b *testing.B) {
	// Test GetMin performance with large stack
	minStack := MinStackConstructor()
	for i := 0; i < 10000; i++ {
		minStack.Push(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minStack.GetMin()
	}
}

func BenchmarkMinStackMixedOperations(b *testing.B) {
	minStack := MinStackConstructor()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Mix of operations
		minStack.Push(i % 100)
		if i%3 == 0 {
			minStack.GetMin()
		}
		if i%4 == 0 {
			minStack.Top()
		}
		if i%5 == 0 && i > 0 {
			minStack.Pop()
		}
	}
}