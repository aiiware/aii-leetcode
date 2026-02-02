package advanced_algorithms

import (
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	n := 10
	uf := NewUnionFind(n)
	
	if uf.Count() != n {
		t.Errorf("Expected %d components, got %d", n, uf.Count())
	}
	
	for i := 0; i < n; i++ {
		if uf.Find(i) != i {
			t.Errorf("Expected root %d for element %d, got %d", i, i, uf.Find(i))
		}
	}
}

func TestUnionFind_Union(t *testing.T) {
	uf := NewUnionFind(5)
	
	// Test union operation
	if !uf.Union(0, 1) {
		t.Error("Expected union to return true for different sets")
	}
	
	if uf.Count() != 4 {
		t.Errorf("Expected 4 components after union, got %d", uf.Count())
	}
	
	if !uf.Connected(0, 1) {
		t.Error("Expected 0 and 1 to be connected after union")
	}
	
	// Test union on already connected elements
	if uf.Union(0, 1) {
		t.Error("Expected union to return false for already connected elements")
	}
	
	// Test transitive union
	uf.Union(1, 2)
	if !uf.Connected(0, 2) {
		t.Error("Expected transitive connection: 0 connected to 2 through 1")
	}
	
	if uf.Count() != 3 {
		t.Errorf("Expected 3 components after second union, got %d", uf.Count())
	}
}

func TestUnionFind_Connected(t *testing.T) {
	uf := NewUnionFind(5)
	
	// Initially, no elements are connected
	if uf.Connected(0, 1) {
		t.Error("Expected 0 and 1 to not be connected initially")
	}
	
	// After union, they should be connected
	uf.Union(0, 1)
	if !uf.Connected(0, 1) {
		t.Error("Expected 0 and 1 to be connected after union")
	}
	
	// Test reflexivity
	if !uf.Connected(0, 0) {
		t.Error("Expected element to be connected to itself")
	}
}

func TestUnionFind_GetComponents(t *testing.T) {
	uf := NewUnionFind(6)
	
	// Create two components: {0, 1, 2} and {3, 4, 5}
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(4, 5)
	
	components := uf.GetComponents()
	
	if len(components) != 2 {
		t.Errorf("Expected 2 components, got %d", len(components))
	}
	
	// Check that each component has 3 elements
	for _, comp := range components {
		if len(comp) != 3 {
			t.Errorf("Expected component size 3, got %d", len(comp))
		}
	}
	
	// Verify elements are in correct components
	for i := 0; i < 3; i++ {
		if !uf.Connected(0, i) {
			t.Errorf("Expected %d to be in same component as 0", i)
		}
	}
	
	for i := 3; i < 6; i++ {
		if !uf.Connected(3, i) {
			t.Errorf("Expected %d to be in same component as 3", i)
		}
	}
}

func TestUnionFind_Reset(t *testing.T) {
	uf := NewUnionFind(5)
	
	// Perform some unions
	uf.Union(0, 1)
	uf.Union(2, 3)
	
	if uf.Count() != 3 {
		t.Errorf("Expected 3 components before reset, got %d", uf.Count())
	}
	
	// Reset the structure
	uf.Reset()
	
	if uf.Count() != 5 {
		t.Errorf("Expected 5 components after reset, got %d", uf.Count())
	}
	
	// Verify all elements are disconnected
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if uf.Connected(i, j) {
				t.Errorf("Expected %d and %d to be disconnected after reset", i, j)
			}
		}
	}
}

func TestUnionFind_PathCompression(t *testing.T) {
	uf := NewUnionFind(10)
	
	// Create a chain: 0-1-2-3-4-5-6-7-8-9
	for i := 0; i < 9; i++ {
		uf.Union(i, i+1)
	}
	
	// After path compression, all should have root 0
	for i := 0; i < 10; i++ {
		root := uf.Find(i)
		if root != 0 {
			t.Errorf("Expected root 0 for element %d after path compression, got %d", i, root)
		}
	}
}

func BenchmarkUnionFind_Union(b *testing.B) {
	uf := NewUnionFind(b.N)
	b.ResetTimer()
	
	for i := 0; i < b.N-1; i++ {
		uf.Union(i, i+1)
	}
}

func BenchmarkUnionFind_Find(b *testing.B) {
	uf := NewUnionFind(b.N)
	// Create a chain to test path compression
	for i := 0; i < b.N-1; i++ {
		uf.Union(i, i+1)
	}
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		uf.Find(i)
	}
}