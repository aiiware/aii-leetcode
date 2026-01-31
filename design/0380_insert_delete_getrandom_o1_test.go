package design

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	// Test 1: Basic operations
	t.Run("Basic operations", func(t *testing.T) {
		rs := Constructor()
		
		// Insert 1
		if !rs.Insert(1) {
			t.Error("Insert(1) should return true")
		}
		
		// Insert 1 again (should return false)
		if rs.Insert(1) {
			t.Error("Insert(1) again should return false")
		}
		
		// Remove 2 (doesn't exist)
		if rs.Remove(2) {
			t.Error("Remove(2) should return false")
		}
		
		// Remove 1
		if !rs.Remove(1) {
			t.Error("Remove(1) should return true")
		}
		
		// Insert 2
		if !rs.Insert(2) {
			t.Error("Insert(2) should return true")
		}
		
		// GetRandom should return 2
		val := rs.GetRandom()
		if val != 2 {
			t.Errorf("GetRandom() should return 2, got %d", val)
		}
	})
	
	// Test 2: Multiple inserts and random distribution
	t.Run("Multiple inserts and random distribution", func(t *testing.T) {
		rs := Constructor()
		
		// Insert 10 values
		for i := 0; i < 10; i++ {
			if !rs.Insert(i) {
				t.Errorf("Insert(%d) should return true", i)
			}
		}
		
		// GetRandom 100 times and count distribution
		counts := make(map[int]int)
		for i := 0; i < 1000; i++ {
			val := rs.GetRandom()
			counts[val]++
		}
		
		// All values should appear roughly equally
		for i := 0; i < 10; i++ {
			if counts[i] == 0 {
				t.Errorf("Value %d never appeared in 1000 random samples", i)
			}
			// Rough check: each should appear ~100 times (1000/10 = 100)
			// Allow some variance: 50-150
			if counts[i] < 50 || counts[i] > 150 {
				t.Errorf("Value %d appears %d times, expected ~100", i, counts[i])
			}
		}
	})
	
	// Test 3: Insert, remove, insert same value
	t.Run("Insert remove insert same value", func(t *testing.T) {
		rs := Constructor()
		
		// Insert 5
		if !rs.Insert(5) {
			t.Error("Insert(5) should return true")
		}
		
		// Remove 5
		if !rs.Remove(5) {
			t.Error("Remove(5) should return true")
		}
		
		// Insert 5 again
		if !rs.Insert(5) {
			t.Error("Insert(5) again should return true")
		}
		
		// GetRandom should return 5
		val := rs.GetRandom()
		if val != 5 {
			t.Errorf("GetRandom() should return 5, got %d", val)
		}
	})
	
	// Test 4: Remove from middle and random still works
	t.Run("Remove from middle", func(t *testing.T) {
		rs := Constructor()
		
		// Insert 1, 2, 3
		for i := 1; i <= 3; i++ {
			if !rs.Insert(i) {
				t.Errorf("Insert(%d) should return true", i)
			}
		}
		
		// Remove 2
		if !rs.Remove(2) {
			t.Error("Remove(2) should return true")
		}
		
		// GetRandom should return either 1 or 3
		for i := 0; i < 100; i++ {
			val := rs.GetRandom()
			if val != 1 && val != 3 {
				t.Errorf("GetRandom() should return 1 or 3, got %d", val)
			}
		}
		
		// Insert 4
		if !rs.Insert(4) {
			t.Error("Insert(4) should return true")
		}
		
		// Now random should return 1, 3, or 4
		counts := make(map[int]int)
		for i := 0; i < 300; i++ {
			val := rs.GetRandom()
			counts[val]++
		}
		
		// Check all three values appear
		for _, val := range []int{1, 3, 4} {
			if counts[val] == 0 {
				t.Errorf("Value %d never appeared in 300 random samples", val)
			}
		}
	})
	
	// Test 5: Empty set edge cases
	t.Run("Empty set", func(t *testing.T) {
		rs := Constructor()
		
		// Remove from empty set
		if rs.Remove(1) {
			t.Error("Remove(1) from empty set should return false")
		}
		
		// GetRandom from empty set (should return -1 per our implementation)
		val := rs.GetRandom()
		if val != -1 {
			t.Errorf("GetRandom() from empty set should return -1, got %d", val)
		}
	})
	
	// Test 6: Large number of operations
	t.Run("Large number of operations", func(t *testing.T) {
		rs := Constructor()
		
		// Insert 1000 values
		for i := 0; i < 1000; i++ {
			if !rs.Insert(i) {
				t.Errorf("Insert(%d) should return true", i)
			}
		}
		
		// Remove every other value
		for i := 0; i < 1000; i += 2 {
			if !rs.Remove(i) {
				t.Errorf("Remove(%d) should return true", i)
			}
		}
		
		// Insert some new values
		for i := 1000; i < 1100; i++ {
			if !rs.Insert(i) {
				t.Errorf("Insert(%d) should return true", i)
			}
		}
		
		// GetRandom should only return odd numbers 1-999 and 1000-1099
		counts := make(map[int]int)
		for i := 0; i < 10000; i++ {
			val := rs.GetRandom()
			counts[val]++
			
			// Check value is valid
			if val%2 == 0 && val < 1000 {
				t.Errorf("GetRandom() returned even number %d that should have been removed", val)
			}
			if val >= 1100 {
				t.Errorf("GetRandom() returned value %d >= 1100", val)
			}
		}
		
		// Check all expected values appear
		for i := 1; i < 1000; i += 2 {
			if counts[i] == 0 {
				t.Errorf("Odd value %d never appeared in 10000 random samples", i)
			}
		}
		for i := 1000; i < 1100; i++ {
			if counts[i] == 0 {
				t.Errorf("New value %d never appeared in 10000 random samples", i)
			}
		}
	})
	
	// Test 7: Optimized version
	t.Run("Optimized version", func(t *testing.T) {
		rs := ConstructorOptimized()
		
		// Insert values
		for i := 0; i < 100; i++ {
			if !rs.Insert(i) {
				t.Errorf("Insert(%d) should return true", i)
			}
		}
		
		// Remove some values
		for i := 0; i < 50; i++ {
			if !rs.Remove(i) {
				t.Errorf("Remove(%d) should return true", i)
			}
		}
		
		// GetRandom should only return values 50-99
		for i := 0; i < 100; i++ {
			val := rs.GetRandom()
			if val < 50 || val >= 100 {
				t.Errorf("GetRandom() should return value between 50-99, got %d", val)
			}
		}
	})
}

func BenchmarkRandomizedSet(b *testing.B) {
	// Benchmark insert operations
	b.Run("Insert", func(b *testing.B) {
		rs := Constructor()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			rs.Insert(i)
		}
	})
	
	// Benchmark remove operations
	b.Run("Remove", func(b *testing.B) {
		rs := Constructor()
		// Pre-insert values
		for i := 0; i < b.N; i++ {
			rs.Insert(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			rs.Remove(i)
		}
	})
	
	// Benchmark GetRandom operations
	b.Run("GetRandom", func(b *testing.B) {
		rs := Constructor()
		// Insert some values
		for i := 0; i < 1000; i++ {
			rs.Insert(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = rs.GetRandom()
		}
	})
	
	// Benchmark mixed operations
	b.Run("Mixed", func(b *testing.B) {
		rs := Constructor()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Alternate operations
			if i%3 == 0 {
				rs.Insert(i)
			} else if i%3 == 1 {
				rs.Remove(i - 1)
			} else {
				if len(rs.values) > 0 {
					_ = rs.GetRandom()
				}
			}
		}
	})
	
	// Benchmark optimized version
	b.Run("Optimized Insert", func(b *testing.B) {
		rs := ConstructorOptimized()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			rs.Insert(i)
		}
	})
}