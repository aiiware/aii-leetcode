package design

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		stack := ConstructorMyStack()
		stack.Push(1)
		stack.Push(2)
		
		if top := stack.Top(); top != 2 {
			t.Errorf("Top() = %d, expected 2", top)
		}
		
		if val := stack.Pop(); val != 2 {
			t.Errorf("Pop() = %d, expected 2", val)
		}
		
		if val := stack.Pop(); val != 1 {
			t.Errorf("Pop() = %d, expected 1", val)
		}
	})

	t.Run("Example from problem", func(t *testing.T) {
		stack := ConstructorMyStack()
		
		stack.Push(1)
		stack.Push(2)
		
		if top := stack.Top(); top != 2 {
			t.Errorf("Top() = %d, expected 2", top)
		}
		
		if val := stack.Pop(); val != 2 {
			t.Errorf("Pop() = %d, expected 2", val)
		}
		
		if empty := stack.Empty(); empty != false {
			t.Errorf("Empty() = %v, expected false", empty)
		}
	})

	t.Run("Single element", func(t *testing.T) {
		stack := ConstructorMyStack()
		stack.Push(42)
		
		if top := stack.Top(); top != 42 {
			t.Errorf("Top() = %d, expected 42", top)
		}
		
		if val := stack.Pop(); val != 42 {
			t.Errorf("Pop() = %d, expected 42", val)
		}
		
		if empty := stack.Empty(); empty != true {
			t.Errorf("Empty() = %v, expected true", empty)
		}
	})

	t.Run("Empty stack", func(t *testing.T) {
		stack := ConstructorMyStack()
		
		if empty := stack.Empty(); empty != true {
			t.Errorf("Empty() = %v, expected true", empty)
		}
	})

	t.Run("Multiple elements", func(t *testing.T) {
		stack := ConstructorMyStack()
		
		for i := 1; i <= 5; i++ {
			stack.Push(i)
		}
		
		// Pop in reverse order (LIFO)
		for i := 5; i >= 1; i-- {
			if val := stack.Pop(); val != i {
				t.Errorf("Pop() = %d, expected %d", val, i)
			}
		}
	})

	t.Run("Negative numbers", func(t *testing.T) {
		stack := ConstructorMyStack()
		
		stack.Push(-1)
		stack.Push(-2)
		stack.Push(-3)
		
		if val := stack.Pop(); val != -3 {
			t.Errorf("Pop() = %d, expected -3", val)
		}
		
		if val := stack.Pop(); val != -2 {
			t.Errorf("Pop() = %d, expected -2", val)
		}
		
		if val := stack.Pop(); val != -1 {
			t.Errorf("Pop() = %d, expected -1", val)
		}
	})

	t.Run("Mixed positive and negative", func(t *testing.T) {
		stack := ConstructorMyStack()
		
		stack.Push(5)
		stack.Push(-3)
		stack.Push(10)
		stack.Push(-7)
		
		if val := stack.Pop(); val != -7 {
			t.Errorf("Pop() = %d, expected -7", val)
		}
		
		if val := stack.Pop(); val != 10 {
			t.Errorf("Pop() = %d, expected 10", val)
		}
		
		if val := stack.Pop(); val != -3 {
			t.Errorf("Pop() = %d, expected -3", val)
		}
		
		if val := stack.Pop(); val != 5 {
			t.Errorf("Pop() = %d, expected 5", val)
		}
	})
}

func TestMyStackSingleQueue(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		stack := ConstructorMyStackSingleQueue()
		stack.Push(1)
		stack.Push(2)
		
		if top := stack.Top(); top != 2 {
			t.Errorf("Top() = %d, expected 2", top)
		}
		
		if val := stack.Pop(); val != 2 {
			t.Errorf("Pop() = %d, expected 2", val)
		}
		
		if val := stack.Pop(); val != 1 {
			t.Errorf("Pop() = %d, expected 1", val)
		}
	})

	t.Run("Single element", func(t *testing.T) {
		stack := ConstructorMyStackSingleQueue()
		stack.Push(42)
		
		if top := stack.Top(); top != 42 {
			t.Errorf("Top() = %d, expected 42", top)
		}
		
		if val := stack.Pop(); val != 42 {
			t.Errorf("Pop() = %d, expected 42", val)
		}
		
		if empty := stack.Empty(); empty != true {
			t.Errorf("Empty() = %v, expected true", empty)
		}
	})

	t.Run("Multiple elements", func(t *testing.T) {
		stack := ConstructorMyStackSingleQueue()
		
		for i := 1; i <= 5; i++ {
			stack.Push(i)
		}
		
		// Pop in reverse order (LIFO)
		for i := 5; i >= 1; i-- {
			if val := stack.Pop(); val != i {
				t.Errorf("Pop() = %d, expected %d", val, i)
			}
		}
	})
}

func BenchmarkMyStackPushPop(b *testing.B) {
	stack := ConstructorMyStack()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
		stack.Pop()
	}
}

func BenchmarkMyStackSingleQueuePushPop(b *testing.B) {
	stack := ConstructorMyStackSingleQueue()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
		stack.Pop()
	}
}

func BenchmarkMyStackTop(b *testing.B) {
	stack := ConstructorMyStack()
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Top()
	}
}

func BenchmarkMyStackEmpty(b *testing.B) {
	stack := ConstructorMyStack()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Empty()
	}
}
