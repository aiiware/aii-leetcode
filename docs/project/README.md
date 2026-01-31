# LeetCode Solutions in Go

A comprehensive collection of LeetCode problem solutions implemented in Go with detailed documentation, tests, and benchmarks.

## Overview

This package contains solutions for **210+ LeetCode problems** organized by algorithmic category with:
- ✅ Clean, idiomatic Go implementations
- ✅ Comprehensive test coverage with edge cases
- ✅ Performance benchmarks
- ✅ Time and space complexity analysis
- ✅ Difficulty levels and topic tags
- ✅ Multiple approaches for some problems
- ✅ Interactive tutorials and learning resources

## 📊 Project Statistics

| Metric | Count |
|--------|-------|
| **Total Problems** | 210+ |
| **Implementation Files** | 210+ |
| **Test Files** | 210+ |
| **Total Go Files** | 420+ |
| **Documentation Files** | 25+ |
| **Lines of Code** | ~85,000 |
| **Test Cases** | 2,500+ |
| **Test Pass Rate** | 100% ✅ |

## 🗂️ Project Structure

```
leetcode/
├── arrays/                    # Array problems (58 problems)
├── binary-tree/              # Binary tree problems (20 problems)
├── data_structures/          # Data structure problems (3 problems)
├── design/                   # Design problems (10 problems)
├── dp/                       # Dynamic programming (28 problems)
├── graphs/                   # Graph problems (10 problems)
├── linked-lists/             # Linked list problems (14 problems)
├── math/                     # Math problems (17 problems)
├── sorting/                  # Sorting problems (10 problems)
├── sql/                      # SQL problems (9 problems)
├── strings/                  # String problems (30 problems)
├── cmd/                      # Command-line tools and demos
│   ├── analyze/              # Analysis tools
│   ├── demo/                 # Demo program with all examples
│   └── ... (other demos)
├── data_structures/          # Data structure implementations
├── debug/                    # Debug utilities and test cases
├── explanations/             # Detailed solution explanations
├── indexes/                  # Index files for navigation
│   ├── by_category.md       # Problems by category
│   ├── by_difficulty.md     # Problems by difficulty
│   └── by_number.md         # Problems by number
├── scripts/                  # Utility scripts
├── testutils/                # Testing utilities
├── utils/                    # Utility functions
└── docs/                     # Documentation
```

## 📈 Problem Distribution by Category

| Category | Count | Percentage |
|----------|-------|------------|
| Arrays | 58 | 27.6% |
| Strings | 30 | 14.3% |
| Dynamic Programming | 28 | 13.3% |
| Math | 17 | 8.1% |
| Binary Tree | 20 | 9.5% |
| Linked Lists | 14 | 6.7% |
| Sorting | 10 | 4.8% |
| Design | 10 | 4.8% |
| Graphs | 10 | 4.8% |
| SQL | 9 | 4.3% |
| Data Structures | 3 | 1.4% |
| **Total** | **210** | **100%** |

## 📈 Problem Distribution by Difficulty

| Difficulty | Count | Percentage |
|------------|-------|------------|
| Easy | 49 | 23.3% |
| Medium | 118 | 56.2% |
| Hard | 42 | 20.0% |
| **Total** | **210** | **100%** |

## 🔗 Quick Navigation

- **[By Category](indexes/by_category.md)** - Browse problems by algorithmic category
- **[By Difficulty](indexes/by_difficulty.md)** - Browse problems by difficulty level
- **[By Number](indexes/by_number.md)** - Browse problems in numerical order
- **[Solution Explanations](explanations/)** - Detailed explanations for complex problems

## Usage

### Running Tests

```bash
cd leetcode
go test ./...              # Run all tests
go test -v ./...           # Run all tests with verbose output
go test ./arrays/...       # Run tests for specific category
go test -cover ./...       # Run tests with coverage report
```

### Running Benchmarks

```bash
cd leetcode
go test -bench=. -benchmem ./...          # Run all benchmarks
go test -bench=BenchmarkTwoSum -benchmem ./arrays  # Run specific benchmark
go test -bench=. -benchmem -count=3 ./... # Run benchmarks 3 times
```

### Running Demo

```bash
cd leetcode
go run cmd/demo/main.go
```

### Using in Your Code

```go
package main

import (
    "fmt"
    "github.com/yourusername/leetcode/arrays"
    "github.com/yourusername/leetcode/strings"
)

func main() {
    // Problem 0001: Two Sum
    nums := []int{2, 7, 11, 15}
    target := 9
    result := arrays.TwoSum(nums, target)
    fmt.Println(result) // Output: [0 1]
    
    // Problem 0020: Valid Parentheses
    s := "()[]{}"
    valid := strings.IsValid(s)
    fmt.Println(valid) // Output: true
}
```

## Key Features

### Comprehensive Testing
- **LeetCode Examples**: All official examples included in tests
- **Edge Cases**: Empty inputs, single elements, boundary values, overflows
- **Performance**: Benchmarks with different input sizes
- **100% Pass Rate**: All tests passing

### Code Quality
- **Documentation**: Clear docstrings with complexity analysis
- **Best Practices**: Idiomatic Go with proper error handling
- **Multiple Approaches**: Different solutions for some problems
- **Performance Optimization**: Optimized implementations where possible

### Interactive Learning Resources
- **Algorithm Tutorials**: Step-by-step guides by category
- **Problem-Solving Guides**: How to approach specific problem types
- **Interactive Exercises**: Hands-on coding challenges
- **Visualizations**: Algorithm animations and explanations

### Difficulty Levels and Tags
- Each problem clearly labeled with:
  - LeetCode difficulty (Easy, Medium, Hard)
  - Relevant algorithm tags
  - Time and space complexity
  - Problem description

## Implementation Highlights

### Problem-Specific Optimizations

1. **0001 - Two Sum**: Single-pass hash map for optimal O(n) complexity
2. **0004 - Median**: Binary search on smaller array for O(log min(m,n))
3. **0003 - Longest Substring**: Unicode-aware character handling
4. **0006 - Zigzag**: Mathematical pattern recognition for O(n) time
5. **0010 - Regular Expression**: Both DP and recursive solutions
6. **0023 - Merge k Lists**: Both heap-based and divide-and-conquer approaches
7. **0025 - Reverse k-Group**: Both recursive and iterative implementations
8. **0030 - Substring Concatenation**: Sliding window with hash map optimization
9. **0031 - Next Permutation**: In-place algorithm with O(n) time complexity
10. **0042 - Trapping Rain Water**: Multiple approaches (two pointers, DP, stack)

### Reusable Components

- **ListNode**: Shared data structure for linked list problems
  - Helper functions: `NewListFromSlice()`, `ToSlice()`, `Equal()`
  - Used by problems: 0002, 0019, 0021, 0023, 0024, 0025, 0061, 0082, 0083, 0086, 0092, 0109, 0141-0148

- **TreeNode**: Shared data structure for binary tree problems
  - Helper functions: `NewTreeFromSlice()`, `TreeToSlice()`, `IntPtr()`
  - Used by problems: 0094, 0095, 0096, 0098, 0099, 0100, 0101, 0102, 0103, 0104, 0105, 0106, 0107, 0108, 0109, 0110, 0111, 0112, 0113, 0114, 0116, 0117, 0124, 0129, 0144, 0145

## Recent Improvements

✅ **Completed**: Added difficulty tags to all 210 solution files  
✅ **Completed**: Strengthened Graphs category (now 10 problems)  
✅ **Completed**: Strengthened Design category (now 10 problems)  
✅ **Completed**: Created Data Structures category (3 problems)  
✅ **Completed**: Strengthened Sorting category (now 10 problems)  
✅ **In Progress**: Creating solution explanations for complex problems  

## Project Planning

For detailed project planning and backlog, see **[BACKLOG.md](BACKLOG.md)** which includes:
- ✅ **Completed tasks**: All critical and high-priority items completed
- ⏳ **Current work**: Adding solution explanations for complex problems
- 📋 **Future plans**: Adding more problems, creating learning paths
- 📊 **Progress tracking**: Current status and success metrics

## Contributing

To add more problems:

1. Create implementation file in appropriate category: `NNNN_problem_name.go`
2. Add test file: `NNNN_problem_name_test.go`
3. Update demo program: `cmd/demo/main.go`
4. Update index files in `indexes/` directory
5. Update this README with problem details

## Testing Status

✅ All tests passing (100% pass rate)
✅ All benchmarks running
✅ Demo program working
✅ Code compiles without warnings

## License

This project is for educational purposes.

## References

- [LeetCode](https://leetcode.com/)
- [Go Documentation](https://golang.org/doc/)