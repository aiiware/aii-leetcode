# LeetCode Solutions in Go

A comprehensive collection of LeetCode problem solutions implemented in Go with detailed documentation, tests, and benchmarks.

## Overview

This package contains solutions for LeetCode problems **0001-0030** with:
- ✅ Clean, idiomatic Go implementations
- ✅ Comprehensive test coverage with edge cases
- ✅ Performance benchmarks
- ✅ Time and space complexity analysis
- ✅ Difficulty levels and topic tags
- ✅ Multiple approaches for some problems

## Problems Implemented

### Problems 0001-0010

| # | Problem | Difficulty | Tags | File |
|---|---------|------------|------|------|
| 0001 | Two Sum | Easy | Array, Hash Table | `0001_two_sum.go` |
| 0002 | Add Two Numbers | Medium | Linked List, Math | `0002_add_two_numbers.go` |
| 0003 | Longest Substring Without Repeating Characters | Medium | Hash Table, String, Sliding Window | `0003_longest_substring.go` |
| 0004 | Median of Two Sorted Arrays | Hard | Array, Binary Search, Divide and Conquer | `0004_median_sorted_arrays.go` |
| 0005 | Longest Palindromic Substring | Medium | String, Dynamic Programming | `0005_longest_palindromic_substring.go` |
| 0006 | Zigzag Conversion | Medium | String | `0006_zigzag_conversion.go` |
| 0007 | Reverse Integer | Medium | Math | `0007_reverse_integer.go` |
| 0008 | String to Integer (atoi) | Medium | String | `0008_string_to_integer.go` |
| 0009 | Palindrome Number | Easy | Math | `0009_palindrome_number.go` |
| 0010 | Regular Expression Matching | Hard | String, Dynamic Programming, Recursion | `0010_regular_expression_matching.go` |

### Problems 0011-0020

| # | Problem | Difficulty | Tags | File |
|---|---------|------------|------|------|
| 0011 | Container With Most Water | Hard | Array, Two Pointers | `0011_container_with_most_water.go` |
| 0012 | Integer to Roman | Medium | Hash Table, Math, String | `0012_integer_to_roman.go` |
| 0013 | Roman to Integer | Easy | Hash Table, Math, String | `0013_roman_to_integer.go` |
| 0014 | Longest Common Prefix | Easy | String, Trie | `0014_longest_common_prefix.go` |
| 0015 | 3Sum | Medium | Array, Two Pointers, Sorting | `0015_3sum.go` |
| 0016 | 3Sum Closest | Medium | Array, Two Pointers, Sorting | `0016_3sum_closest.go` |
| 0017 | Letter Combinations of a Phone Number | Medium | String, Backtracking, Combinatorics | `0017_letter_combinations.go` |
| 0018 | 4Sum | Medium | Array, Two Pointers, Sorting | `0018_4sum.go` |
| 0019 | Remove Nth Node From End of List | Medium | Linked List, Two Pointers | `0019_remove_nth_node.go` |
| 0020 | Valid Parentheses | Easy | String, Stack | `0020_valid_parentheses.go` |

### Problems 0021-0030

| # | Problem | Difficulty | Tags | File |
|---|---------|------------|------|------|
| 0021 | Merge Two Sorted Lists | Easy | Linked List, Recursion | `0021_merge_two_sorted_lists.go` |
| 0022 | Generate Parentheses | Medium | String, Backtracking, Dynamic Programming | `0022_generate_parentheses.go` |
| 0023 | Merge k Sorted Lists | Hard | Linked List, Divide and Conquer, Heap | `0023_merge_k_sorted_lists.go` |
| 0024 | Swap Nodes in Pairs | Medium | Linked List, Recursion | `0024_swap_nodes_in_pairs.go` |
| 0025 | Reverse Nodes in k-Group | Hard | Linked List, Recursion | `0025_reverse_nodes_in_k_group.go` |
| 0026 | Remove Duplicates from Sorted Array | Easy | Array, Two Pointers | `0026_remove_duplicates_from_sorted_array.go` |
| 0027 | Remove Element | Easy | Array, Two Pointers | `0027_remove_element.go` |
| 0028 | Find the Index of the First Occurrence in a String | Easy | String, Two Pointers, String Matching | `0028_str_str.go` |
| 0029 | Divide Two Integers | Medium | Math, Bit Manipulation | `0029_divide_two_integers.go` |
| 0030 | Substring with Concatenation of All Words | Hard | Hash Table, String, Sliding Window | `0030_substring_with_concatenation_of_all_words.go` |

## File Structure

```
leetcode/
├── cmd/
│   └── demo/
│       └── main.go              # Demo program with all examples
├── 0001_two_sum.go              # Problem 0001 (Easy)
├── 0002_add_two_numbers.go      # Problem 0002 (Medium)
├── 0003_longest_substring.go    # Problem 0003 (Medium)
├── 0004_median_sorted_arrays.go # Problem 0004 (Hard)
├── 0005_longest_palindromic_substring.go # Problem 0005 (Medium)
├── 0006_zigzag_conversion.go    # Problem 0006 (Medium)
├── 0007_reverse_integer.go      # Problem 0007 (Medium)
├── 0008_string_to_integer.go    # Problem 0008 (Medium)
├── 0009_palindrome_number.go    # Problem 0009 (Easy)
├── 0010_regular_expression_matching.go # Problem 0010 (Hard)
├── 0011_container_with_most_water.go # Problem 0011 (Hard)
├── 0012_integer_to_roman.go     # Problem 0012 (Medium)
├── 0013_roman_to_integer.go     # Problem 0013 (Easy)
├── 0014_longest_common_prefix.go # Problem 0014 (Easy)
├── 0015_3sum.go                 # Problem 0015 (Medium)
├── 0016_3sum_closest.go         # Problem 0016 (Medium)
├── 0017_letter_combinations.go  # Problem 0017 (Medium)
├── 0018_4sum.go                 # Problem 0018 (Medium)
├── 0019_remove_nth_node.go      # Problem 0019 (Medium)
├── 0020_valid_parentheses.go    # Problem 0020 (Easy)
├── 0021_merge_two_sorted_lists.go # Problem 0021 (Easy)
├── 0022_generate_parentheses.go # Problem 0022 (Medium)
├── 0023_merge_k_sorted_lists.go # Problem 0023 (Hard)
├── 0024_swap_nodes_in_pairs.go  # Problem 0024 (Medium)
├── 0025_reverse_nodes_in_k_group.go # Problem 0025 (Hard)
├── 0026_remove_duplicates_from_sorted_array.go # Problem 0026 (Easy)
├── 0027_remove_element.go       # Problem 0027 (Easy)
├── 0028_str_str.go              # Problem 0028 (Easy)
├── 0029_divide_two_integers.go  # Problem 0029 (Medium)
├── 0030_substring_with_concatenation_of_all_words.go # Problem 0030 (Hard)
├── 30 test files (*_test.go)    # Comprehensive tests
├── list_node.go                 # Shared ListNode struct
├── go.mod                       # Go module
├── go.sum                       # Dependencies
└── 7 documentation files (*.md) # Complete documentation
```

## Usage

### Running Tests

```bash
cd leetcode
go test -v              # Run all tests with verbose output
go test -v -run 0001    # Run tests for specific problem
go test -cover          # Run tests with coverage report
```

### Running Benchmarks

```bash
cd leetcode
go test -bench=. -benchmem          # Run all benchmarks
go test -bench=BenchmarkTwoSum -benchmem  # Run specific benchmark
go test -bench=. -benchmem -count=3 # Run benchmarks 3 times
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
	"path/to/leetcode"
)

func main() {
	// Problem 0001: Two Sum
	nums := []int{2, 7, 11, 15}
	target := 9
	result := leetcode.TwoSum(nums, target)
	fmt.Println(result) // Output: [0 1]
	
	// Problem 0020: Valid Parentheses
	s := "()[]{}"
	valid := leetcode.IsValid(s)
	fmt.Println(valid) // Output: true
	
	// Problem 0021: Merge Two Sorted Lists
	l1 := leetcode.NewListFromSlice([]int{1, 2, 4})
	l2 := leetcode.NewListFromSlice([]int{1, 3, 4})
	result := leetcode.MergeTwoLists(l1, l2)
	fmt.Println(result.ToSlice()) // Output: [1 1 2 3 4 4]
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

### Difficulty Levels and Tags
- Each problem clearly labeled with:
  - LeetCode difficulty (Easy, Medium, Hard)
  - Relevant algorithm tags
  - Time and space complexity
  - Problem description

## Benchmark Results

Sample benchmark results (Apple M4 Max):

```
BenchmarkTwoSum-16                           3084    384961 ns/op    591503 B/op    80 allocs/op
BenchmarkAddTwoNumbers-16                     20   51726980 ns/op   1008004 B/op    20 allocs/op
BenchmarkLengthOfLongestSubstring-16      39648     30207 ns/op       2048 B/op    10 allocs/op
BenchmarkMedianSortedArrays-16           153841      7681 ns/op       1232 B/op     2 allocs/op
BenchmarkLongestPalindrome-16               598    1989234 ns/op    1032192 B/op    11 allocs/op
BenchmarkConvert-16                     1000000       1056 ns/op        256 B/op     1 allocs/op
BenchmarkReverse-16                   1000000000         1.109 ns/op      0 B/op     0 allocs/op
BenchmarkMyAtoi-16                     1000000       1024 ns/op        512 B/op     3 allocs/op
BenchmarkIsPalindrome-16               1000000000      1.037 ns/op      0 B/op     0 allocs/op
BenchmarkIsMatch-16                        100   14005923 ns/op    1048576 B/op    17 allocs/op
```

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

### Reusable Components

- **ListNode**: Shared data structure for linked list problems
  - Helper functions: `NewListFromSlice()`, `ToSlice()`, `Equal()`
  - Used by problems: 0002, 0019, 0021, 0023, 0024, 0025

## Future Improvements

- [ ] Add more LeetCode problems (0031-0050, etc.)
- [ ] Interactive problem selector
- [ ] Performance comparison tool
- [ ] Algorithm visualization
- [ ] Alternative implementations
- [ ] Extended problem set with variations

## Contributing

To add more problems:

1. Create implementation file: `NNNN_problem_name.go`
2. Add test file: `NNNN_problem_name_test.go`
3. Update demo program: `cmd/demo/main.go`
4. Update this README with problem details

## Testing Status

✅ All tests passing (100% pass rate)
✅ All benchmarks running
✅ Demo program working
✅ Code compiles without warnings

## Project Statistics

| Metric | Count |
|--------|-------|
| Problems Implemented | 30 |
| Implementation Files | 30 |
| Test Files | 30 |
| Total Go Files | 61 |
| Documentation Files | 7 |
| Total Files | 70+ |
| Lines of Code | ~6,000 |
| Test Cases | 300+ |
| Test Pass Rate | 100% ✅ |

## License

This project is for educational purposes.

## References

- [LeetCode](https://leetcode.com/)
- [Go Documentation](https://golang.org/doc/)
