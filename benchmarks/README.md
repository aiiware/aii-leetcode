# LeetCode Algorithm Benchmarks

This directory contains performance benchmarks for key LeetCode algorithms implemented in Go. The benchmarks help compare different approaches and understand performance characteristics.

## Available Benchmarks

### 1. Two Sum (`two_sum_benchmark.go`)
- **Brute Force** vs **HashMap** approaches
- Small (22 elements) and large (10,000 elements) datasets
- Shows O(n²) vs O(n) time complexity

### 2. Climbing Stairs (`climbing_stairs_benchmark.go`)
- **Dynamic Programming** solution
- Small (n=5) to very large (n=10,000) inputs
- Demonstrates efficient O(n) time complexity

### 3. LRU Cache (`lru_cache_benchmark.go`)
- Cache operations (Put/Get) with different sizes
- Cache hit vs cache miss scenarios
- Small (10), medium (100), and large (1000) cache sizes

### 4. Graph Cloning (`graph_clone_benchmark.go`)
- **CloneGraph** algorithm using BFS
- Small (4 nodes), medium (100 nodes), large (1000 nodes), and very large (10000 nodes) graphs
- Tests deep copying of graph structures

### 5. Binary Tree Balance Checking (`binary_tree_benchmark.go`)
- **IsBalanced** and **IsBalancedBottomUp** algorithms
- Balanced vs unbalanced trees
- Tests tree traversal and balance checking

## Running Benchmarks

### Run All Benchmarks
```bash
go test -bench=. ./benchmarks/... -benchtime=3s
```

### Run Specific Benchmarks
```bash
# Two Sum benchmarks
go test -bench=TwoSum ./benchmarks/...

# Climbing Stairs benchmarks  
go test -bench=ClimbStairs ./benchmarks/...

# LRU Cache benchmarks
go test -bench=LRUCache ./benchmarks/...

# Graph Cloning benchmarks
go test -bench=CloneGraph ./benchmarks/...

# Binary Tree benchmarks
go test -bench=IsBalanced ./benchmarks/...
```

### Run with Memory Profiling
```bash
go test -bench=. ./benchmarks/... -benchmem -benchtime=3s
```

### Generate Benchmark Report
```bash
go test -bench=. ./benchmarks/... -benchtime=3s > benchmark_results.txt
```

## Example Output

```
BenchmarkTwoSumHashMap-8             5000000               342 ns/op
BenchmarkTwoSumBruteForce-8           200000              6543 ns/op
BenchmarkClimbStairsDP-8            10000000               112 ns/op
BenchmarkCloneGraphSmall-8           1000000              1234 ns/op
BenchmarkIsBalancedBalanced-8        5000000               456 ns/op
```

## Key Insights

1. **Algorithm Choice Matters**: HashMap vs Brute Force shows significant speedup
2. **Dynamic Programming**: Efficient solutions for problems like climbing stairs
3. **Data Structure Impact**: LRU cache shows O(1) vs O(n) operations
4. **Graph Algorithms**: CloneGraph demonstrates BFS traversal efficiency
5. **Tree Algorithms**: Balance checking shows different approaches

## Adding New Benchmarks

1. Create a new file `benchmarks/algorithm_name_benchmark.go`
2. Follow the naming convention: `BenchmarkAlgorithmNameScenario`
3. Include both best-case and worst-case scenarios
4. Test with different input sizes
5. Compare alternative implementations if available

## Performance Tips

- Use `b.ResetTimer()` to exclude setup time
- Test with realistic input sizes
- Include edge cases and pathological inputs
- Run benchmarks multiple times for consistency
- Use `-benchtime` to control duration for slow benchmarks