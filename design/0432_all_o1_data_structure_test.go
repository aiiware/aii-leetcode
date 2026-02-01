package design

import (
	"testing"
)

func TestAllOne(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		obj := AllOneConstructor()
		
		// Test empty
		if max := obj.GetMaxKey(); max != "" {
			t.Errorf("GetMaxKey() on empty = %v, want \"\"", max)
		}
		if min := obj.GetMinKey(); min != "" {
			t.Errorf("GetMinKey() on empty = %v, want \"\"", min)
		}
		
		// Test Inc
		obj.Inc("hello")
		if max := obj.GetMaxKey(); max != "hello" {
			t.Errorf("GetMaxKey() after Inc(\"hello\") = %v, want \"hello\"", max)
		}
		if min := obj.GetMinKey(); min != "hello" {
			t.Errorf("GetMinKey() after Inc(\"hello\") = %v, want \"hello\"", min)
		}
		
		// Test Inc same key
		obj.Inc("hello")
		if max := obj.GetMaxKey(); max != "hello" {
			t.Errorf("GetMaxKey() after second Inc(\"hello\") = %v, want \"hello\"", max)
		}
		
		// Test Inc different key
		obj.Inc("world")
		if min := obj.GetMinKey(); min != "world" {
			t.Errorf("GetMinKey() after Inc(\"world\") = %v, want \"world\"", min)
		}
		
		// Test Dec
		obj.Dec("hello")
		// After Dec("hello"), both "hello" and "world" have frequency 1
		// GetMaxKey() could return either
		max := obj.GetMaxKey()
		if max != "hello" && max != "world" {
			t.Errorf("GetMaxKey() after Dec(\"hello\") = %v, want \"hello\" or \"world\"", max)
		}
		
		// Test Dec to zero
		obj.Dec("world")
		if min := obj.GetMinKey(); min != "hello" {
			t.Errorf("GetMinKey() after Dec(\"world\") = %v, want \"hello\"", min)
		}
		
		// Test Dec non-existent key
		obj.Dec("nonexistent") // Should not panic
	})
	
	t.Run("Multiple keys with same frequency", func(t *testing.T) {
		obj := AllOneConstructor()
		
		obj.Inc("a")
		obj.Inc("b")
		obj.Inc("c")
		
		// All keys have frequency 1, any of them could be returned
		max := obj.GetMaxKey()
		if max != "a" && max != "b" && max != "c" {
			t.Errorf("GetMaxKey() = %v, want one of [a, b, c]", max)
		}
		
		min := obj.GetMinKey()
		if min != "a" && min != "b" && min != "c" {
			t.Errorf("GetMinKey() = %v, want one of [a, b, c]", min)
		}
	})
	
	t.Run("Complex scenario", func(t *testing.T) {
		obj := AllOneConstructor()
		
		// Setup: a:3, b:2, c:1
		obj.Inc("a")
		obj.Inc("a")
		obj.Inc("a")
		obj.Inc("b")
		obj.Inc("b")
		obj.Inc("c")
		
		if max := obj.GetMaxKey(); max != "a" {
			t.Errorf("GetMaxKey() = %v, want \"a\"", max)
		}
		if min := obj.GetMinKey(); min != "c" {
			t.Errorf("GetMinKey() = %v, want \"c\"", min)
		}
		
		// Update: a:2, b:2, c:1
		obj.Dec("a")
		max := obj.GetMaxKey()
		if max != "a" && max != "b" {
			t.Errorf("GetMaxKey() after Dec(\"a\") = %v, want \"a\" or \"b\"", max)
		}
		
		// Update: a:2, b:1, c:1
		obj.Dec("b")
		if min := obj.GetMinKey(); min != "b" && min != "c" {
			t.Errorf("GetMinKey() after Dec(\"b\") = %v, want \"b\" or \"c\"", min)
		}
	})
}