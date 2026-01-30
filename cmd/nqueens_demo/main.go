package main

import (
	"fmt"
	"time"

	"leetcode/arrays"
)

func main() {
	fmt.Println("=== N-Queens Solutions Demo ===")
	
	// Test different board sizes
	for n := 1; n <= 6; n++ {
		fmt.Printf("\n=== n = %d ===\n", n)
		
		// Get all solutions
		solutions := arrays.SolveNQueens(n)
		fmt.Printf("Number of solutions: %d\n", len(solutions))
		
		// Show first solution if it exists
		if len(solutions) > 0 {
			fmt.Println("First solution:")
			for _, row := range solutions[0] {
				fmt.Println(row)
			}
		}
		
		// Compare with optimized version
		optimizedSolutions := arrays.SolveNQueensOptimized(n)
		fmt.Printf("Optimized version found: %d solutions\n", len(optimizedSolutions))
		
		// Show count
		count := arrays.SolveNQueensCount(n)
		fmt.Printf("Count-only version: %d solutions\n", count)
	}
	
	// Show performance comparison for larger n
	fmt.Println("\n=== Performance Comparison ===")
	fmt.Println("n | Basic (ns) | Optimized (ns) | Count (ns)")
	fmt.Println("--|------------|----------------|-----------")
	
	// Run quick benchmarks
	for n := 4; n <= 8; n++ {
		// Quick benchmark - run each 1000 times
		iterations := 1000
		
		// Basic
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_ = arrays.SolveNQueens(n)
		}
		basicTime := time.Since(start).Nanoseconds() / int64(iterations)
		
		// Optimized
		start = time.Now()
		for i := 0; i < iterations; i++ {
			_ = arrays.SolveNQueensOptimized(n)
		}
		optimizedTime := time.Since(start).Nanoseconds() / int64(iterations)
		
		// Count
		start = time.Now()
		for i := 0; i < iterations; i++ {
			_ = arrays.SolveNQueensCount(n)
		}
		countTime := time.Since(start).Nanoseconds() / int64(iterations)
		
		fmt.Printf("%d | %9d | %14d | %9d\n", n, basicTime, optimizedTime, countTime)
	}
	
	// Show all solutions for n=4
	fmt.Println("\n=== All Solutions for n=4 ===")
	solutions := arrays.SolveNQueens(4)
	for i, solution := range solutions {
		fmt.Printf("\nSolution %d:\n", i+1)
		for _, row := range solution {
			fmt.Println(row)
		}
	}
}