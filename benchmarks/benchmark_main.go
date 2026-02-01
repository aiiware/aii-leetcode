package benchmarks

import (
	"fmt"
	"os"
	"testing"
)

// TestMain provides custom setup/teardown for benchmarks
func TestMain(m *testing.M) {
	fmt.Println("=== LeetCode Algorithm Benchmarks ===")
	fmt.Println("Running benchmarks for key algorithms...")
	fmt.Println()
	
	// Run benchmarks
	code := m.Run()
	
	fmt.Println()
	fmt.Println("=== Benchmark Summary ===")
	fmt.Println("Benchmarks completed. Use 'go test -bench=. -benchmem ./benchmarks'")
	fmt.Println("to see detailed memory allocation statistics.")
	
	os.Exit(code)
}

// Example benchmark that won't run by default (name doesn't start with Benchmark)
func ExampleRunAllBenchmarks() {
	fmt.Println("To run all benchmarks:")
	fmt.Println("  go test -bench=. ./benchmarks")
	fmt.Println()
	fmt.Println("To run specific benchmarks:")
	fmt.Println("  go test -bench=TwoSum ./benchmarks")
	fmt.Println("  go test -bench=LRUCache ./benchmarks")
	fmt.Println("  go test -bench=NumIslands ./benchmarks")
	fmt.Println()
	fmt.Println("To include memory allocation stats:")
	fmt.Println("  go test -bench=. -benchmem ./benchmarks")
}