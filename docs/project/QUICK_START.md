# Quick Start Guide - LeetCode Solutions in Go

## 🚀 Getting Started in 2 Minutes

### 1. Run Tests
```bash
cd leetcode
go test -v
```
**Expected Output**: `ok leetcode 0.472s` with 100% pass rate

### 2. Run Demo Program
```bash
cd leetcode
go run cmd/demo/main.go
```
**Expected Output**: Examples for all 10 problems with sample inputs/outputs

### 3. View Documentation
- **Main Guide**: Open `README.md`
- **Problem Index**: Open `PROBLEM_INDEX.md`
- **Quick Reference**: See `PROBLEM_INDEX.md` Quick Reference Table

---

## 📖 Problem Overview

| # | Problem | Difficulty | Tags |
|---|---------|------------|------|
| 0001 | Two Sum | Easy | Array, Hash Table |
| 0002 | Add Two Numbers | Medium | Linked List, Math |
| 0003 | Longest Substring | Medium | String, Sliding Window |
| 0004 | Median of Two Sorted Arrays | Hard | Array, Binary Search |
| 0005 | Longest Palindromic Substring | Medium | String, DP |
| 0006 | Zigzag Conversion | Medium | String |
| 0007 | Reverse Integer | Medium | Math |
| 0008 | String to Integer | Medium | String |
| 0009 | Palindrome Number | Easy | Math |
| 0010 | Regular Expression Matching | Hard | String, DP |

---

## 🎯 Learning Path

### Start Here (Easiest)
1. **0001 - Two Sum** (Easy) - 5 minutes
   - Concept: Hash map for O(1) lookups
   - File: `0001_two_sum.go`

2. **0009 - Palindrome Number** (Easy) - 3 minutes
   - Concept: Mathematical approach
   - File: `0009_palindrome_number.go`

### Beginner to Intermediate (Medium)
3. **0003 - Longest Substring** (Medium) - 10 minutes
   - Concept: Sliding window + hash map
   - File: `0003_longest_substring.go`

4. **0006 - Zigzag Conversion** (Medium) - 8 minutes
   - Concept: Pattern recognition
   - File: `0006_zigzag_conversion.go`

5. **0007 - Reverse Integer** (Medium) - 5 minutes
   - Concept: Overflow handling
   - File: `0007_reverse_integer.go`

6. **0008 - String to Integer** (Medium) - 8 minutes
   - Concept: String parsing
   - File: `0008_string_to_integer.go`

### Intermediate to Advanced (Medium+)
7. **0002 - Add Two Numbers** (Medium) - 10 minutes
   - Concept: Linked list + carry handling
   - File: `0002_add_two_numbers.go`

8. **0005 - Longest Palindromic Substring** (Medium) - 15 minutes
   - Concept: Two approaches (expand vs DP)
   - File: `0005_longest_palindromic_substring.go`

### Advanced (Hard)
9. **0004 - Median of Two Sorted Arrays** (Hard) - 20 minutes
   - Concept: Binary search mastery
   - File: `0004_median_sorted_arrays.go`

10. **0010 - Regular Expression Matching** (Hard) - 25 minutes
    - Concept: Complex DP or recursion
    - File: `0010_regular_expression_matching.go`

---

## 💡 Key Concepts by Problem

### Hash Table Patterns
```go
// 0001 - Two Sum: Store seen values in map
numMap := make(map[int]int)
if idx, found := numMap[complement]; found {
    return []int{idx, i}
}
```

### Sliding Window Pattern
```go
// 0003 - Longest Substring: Two pointers + hash map
for i, ch := range s {
    if lastIndex, found := charIndex[ch]; found {
        start = max(start, lastIndex+1)
    }
    charIndex[ch] = i
}
```

### Linked List Pattern
```go
// 0002 - Add Two Numbers: Traverse with carry
for l1 != nil || l2 != nil || carry > 0 {
    sum := carry
    if l1 != nil {
        sum += l1.Val
        l1 = l1.Next
    }
    // ... build result
}
```

### Dynamic Programming Pattern
```go
// 0010 - Regex Matching: DP table
dp := make([][]bool, m+1)
for i := 1; i <= m; i++ {
    for j := 1; j <= n; j++ {
        if p[j-1] == '*' {
            dp[i][j] = dp[i][j-2] || (...)
        }
    }
}
```

### Binary Search Pattern
```go
// 0004 - Median: Binary search on partition
left, right := 0, len(nums1)
for left <= right {
    partition1 := (left + right) / 2
    // ... search logic
}
```

---

## 🔧 Common Commands

### Run All Tests
```bash
go test -v
```

### Run Tests for One Problem
```bash
go test -v -run 0001
go test -v -run 0002
```

### Run Benchmarks
```bash
go test -bench=. -benchmem
```

### Run Specific Benchmark
```bash
go test -bench=BenchmarkTwoSum -benchmem
```

### Check Test Coverage
```bash
go test -cover
```

### Run Demo
```bash
go run cmd/demo/main.go
```

---

## 📊 File Structure

```
leetcode/
├── 0001_two_sum.go              (Implementation)
├── 0001_two_sum_test.go         (Tests)
├── 0002_add_two_numbers.go
├── 0002_add_two_numbers_test.go
├── ...
├── 0010_regular_expression_matching.go
├── 0010_regular_expression_matching_test.go
├── list_node.go                 (Shared utility)
├── cmd/
│   └── demo/
│       └── main.go              (Demo program)
├── README.md                    (Full guide)
├── PROBLEM_INDEX.md             (Quick reference)
├── IMPLEMENTATION_SUMMARY.md    (Completion summary)
├── QUICK_START.md              (This file)
├── go.mod
└── go.sum
```

---

## 🎓 Understanding the Code

### Example: 0001 - Two Sum

```go
// TwoSum solves LeetCode problem 0001: Two Sum
// Difficulty: Easy
// Tags: Array, Hash Table
// Time complexity: O(n), Space complexity: O(n)

func TwoSum(nums []int, target int) []int {
    // Create a map to store number -> index
    numMap := make(map[int]int)
    
    for i, num := range nums {
        // Calculate the complement needed to reach target
        complement := target - num
        
        // Check if complement exists in map
        if idx, found := numMap[complement]; found {
            return []int{idx, i}
        }
        
        // Store current number and its index
        numMap[num] = i
    }
    
    return []int{}
}
```

**Key Points**:
1. Problem number in file: `0001`
2. Difficulty level: `Easy`
3. Algorithm tags: `Array, Hash Table`
4. Complexity documentation: Time `O(n)`, Space `O(n)`
5. Clear algorithm explanation in comments
6. Efficient single-pass solution

---

## 🧪 Understanding the Tests

```go
func TestTwoSum(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   []int
    }{
        {
            name:   "Example 1",
            nums:   []int{2, 7, 11, 15},
            target: 9,
            want:   []int{0, 1},
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := TwoSum(tt.nums, tt.target)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

**Key Features**:
- Subtests for organized output
- LeetCode official examples included
- Edge cases covered
- Clear error messages

---

## ⚡ Quick Tips

### Tip 1: Run Tests While Studying
```bash
# Watch tests as you read the code
go test -v -run 0001
```

### Tip 2: Compare Implementations
```bash
# Problem 0010 has two approaches
# Look at both IsMatch (DP) and IsMatchRecursive (memoization)
```

### Tip 3: Study Benchmarks
```bash
# See which implementation is faster
go test -bench=. -benchmem
```

### Tip 4: Use Demo as Learning Tool
```bash
# See actual inputs/outputs
go run cmd/demo/main.go
```

---

## 🎯 What You'll Learn

### By Difficulty
- **Easy** (2 problems): Basic hash table and math
- **Medium** (6 problems): Sliding window, DP intro, string manipulation
- **Hard** (2 problems): Advanced DP, binary search

### By Algorithm
- Hash tables for fast lookup
- Sliding window for subarray/substring problems
- Dynamic programming for complex pattern matching
- Binary search for optimized searching
- Linked list manipulation and traversal
- Mathematical optimizations for numeric problems

### By Problem Type
- Array problems
- String problems
- Linked list problems
- Math problems
- Pattern matching problems

---

## 📝 Making Notes

Each file is well-commented. Consider:

1. **Read the problem description** at the top of each file
2. **Understand the complexity** (Time/Space)
3. **Follow the algorithm** through the code with inline comments
4. **Study the tests** to understand edge cases
5. **Run the code** and experiment with different inputs

---

## 🐛 Debugging Tips

### Run a Single Test
```bash
go test -v -run "TestTwoSum/Example_1"
```

### Run Specific Problem Tests
```bash
go test -v -run "Test.*Sum"
```

### Verbose Output
```bash
go test -v
```

### With Panic Recovery
```bash
go test -v -timeout 30s
```

---

## ✅ Verification Checklist

Before moving to the next problem:
- [ ] Code runs without errors
- [ ] All tests pass for that problem
- [ ] Understand the algorithm
- [ ] Know the time/space complexity
- [ ] Can explain the approach

---

## 🚀 Ready to Code?

1. Pick a problem from the learning path above
2. Open the corresponding file (e.g., `0001_two_sum.go`)
3. Read the problem description and solution
4. Run the tests: `go test -v -run 0001`
5. Read the test cases in the test file
6. Try to understand the algorithm
7. Experiment with the demo: `go run cmd/demo/main.go`

**Good luck with your LeetCode journey! 🎓**

---

## 📚 Additional Resources

- **README.md**: Comprehensive guide
- **PROBLEM_INDEX.md**: Quick reference by topic
- **IMPLEMENTATION_SUMMARY.md**: Detailed completion summary
- **Each implementation file**: Full problem description and solution

---

**Last Updated**: December 20, 2025  
**Status**: Ready to use ✅
