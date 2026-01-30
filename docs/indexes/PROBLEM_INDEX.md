# LeetCode Solutions - Problem Index

## Quick Reference Table

| Problem | Title | Difficulty | Tags | Time | Space | File |
|---------|-------|------------|------|------|-------|------|
| 0001 | Two Sum | Easy | Array, Hash Table | O(n) | O(n) | `0001_two_sum.go` |
| 0002 | Add Two Numbers | Medium | Linked List, Math, Recursion | O(max(m,n)) | O(max(m,n)) | `0002_add_two_numbers.go` |
| 0003 | Longest Substring Without Repeating Characters | Medium | Hash Table, String, Sliding Window | O(n) | O(min(m,n)) | `0003_longest_substring.go` |
| 0004 | Median of Two Sorted Arrays | Hard | Array, Binary Search, Divide and Conquer | O(log(min(m,n))) | O(1) | `0004_median_sorted_arrays.go` |
| 0005 | Longest Palindromic Substring | Medium | String, Dynamic Programming | O(n²) | O(n²) / O(1) | `0005_longest_palindromic_substring.go` |
| 0006 | Zigzag Conversion | Medium | String | O(n) | O(n) | `0006_zigzag_conversion.go` |
| 0007 | Reverse Integer | Medium | Math | O(log x) | O(1) | `0007_reverse_integer.go` |
| 0008 | String to Integer (atoi) | Medium | String | O(n) | O(1) | `0008_string_to_integer.go` |
| 0009 | Palindrome Number | Easy | Math | O(log x) | O(1) | `0009_palindrome_number.go` |
| 0010 | Regular Expression Matching | Hard | String, Dynamic Programming, Recursion | O(m×n) | O(m×n) | `0010_regular_expression_matching.go` |

## Organized by Difficulty

### Easy (2 problems)
- **0001** - Two Sum
  - Tags: Array, Hash Table
  - Approach: Hash map for O(1) lookups
  - Key insight: Store seen numbers in map while iterating

- **0009** - Palindrome Number
  - Tags: Math
  - Approach: Reverse half of the number mathematically
  - Key insight: No string conversion needed, O(1) space

### Medium (6 problems)
- **0002** - Add Two Numbers
  - Tags: Linked List, Math, Recursion
  - Approach: Traverse and handle carry
  - Key insight: Process least significant digits first

- **0003** - Longest Substring Without Repeating Characters
  - Tags: Hash Table, String, Sliding Window
  - Approach: Two pointers with hash map
  - Key insight: Unicode-aware character handling

- **0005** - Longest Palindromic Substring
  - Tags: String, Dynamic Programming
  - Approach: Expand around center or DP
  - Key insight: Two approaches with different tradeoffs

- **0006** - Zigzag Conversion
  - Tags: String
  - Approach: Row-based grouping with mathematical pattern
  - Key insight: Period = 2*(numRows-1)

- **0007** - Reverse Integer
  - Tags: Math
  - Approach: Extract digits and handle overflow
  - Key insight: Use int64 to detect overflow before casting to int32

- **0008** - String to Integer (atoi)
  - Tags: String
  - Approach: Character-by-character parsing
  - Key insight: Handle sign, whitespace, and overflow clamping

### Hard (2 problems)
- **0004** - Median of Two Sorted Arrays
  - Tags: Array, Binary Search, Divide and Conquer
  - Approach: Binary search on partition
  - Key insight: O(log(min(m,n))) achievable through partition search

- **0010** - Regular Expression Matching
  - Tags: String, Dynamic Programming, Recursion
  - Approach: DP table or recursive with memoization
  - Key insight: '*' matches zero or more of preceding character

## Organized by Tags

### Array (2 problems)
- 0001 - Two Sum
- 0004 - Median of Two Sorted Arrays

### Binary Search (1 problem)
- 0004 - Median of Two Sorted Arrays

### Divide and Conquer (1 problem)
- 0004 - Median of Two Sorted Arrays

### Dynamic Programming (3 problems)
- 0005 - Longest Palindromic Substring
- 0010 - Regular Expression Matching

### Hash Table (2 problems)
- 0001 - Two Sum
- 0003 - Longest Substring Without Repeating Characters

### Linked List (1 problem)
- 0002 - Add Two Numbers

### Math (4 problems)
- 0002 - Add Two Numbers
- 0007 - Reverse Integer
- 0008 - String to Integer (atoi)
- 0009 - Palindrome Number

### Recursion (2 problems)
- 0002 - Add Two Numbers
- 0010 - Regular Expression Matching

### Sliding Window (1 problem)
- 0003 - Longest Substring Without Repeating Characters

### String (5 problems)
- 0003 - Longest Substring Without Repeating Characters
- 0005 - Longest Palindromic Substring
- 0006 - Zigzag Conversion
- 0008 - String to Integer (atoi)
- 0010 - Regular Expression Matching

## Algorithm Patterns Used

### 1. Hash Map / Hash Table
**Problems**: 0001, 0003
- Used for O(1) lookups and storage
- Trade space for time optimization

**Example**: Two Sum uses hash map to store number → index mapping

### 2. Sliding Window
**Problems**: 0003
- Two pointer technique for substring/subarray problems
- Expand and contract window based on conditions

**Example**: Longest Substring uses two pointers with hash map

### 3. Dynamic Programming
**Problems**: 0005, 0010
- Build up solutions from subproblems
- Store intermediate results to avoid recomputation

**Example**: Regular Expression uses DP table for pattern matching

### 4. Recursion with Memoization
**Problems**: 0002, 0010
- Recursive approach with caching to optimize
- Alternative to bottom-up DP

**Example**: Regex Matching has both DP and recursive solutions

### 5. Binary Search
**Problems**: 0004
- Divide and conquer approach
- Find partition point in sorted arrays

**Example**: Median of Two Sorted Arrays uses binary search

### 6. Mathematical Approach
**Problems**: 0006, 0007, 0009
- Pattern recognition and mathematical formulas
- Optimize space complexity

**Example**: Zigzag uses period = 2*(n-1) formula

### 7. Two Pointers
**Problems**: 0003, 0009 (implicit)
- Navigate through data structure from different ends
- Useful for palindrome checking and substring problems

## Learning Path Recommendation

### Beginner (Start Here)
1. **0001** - Two Sum
   - Basic hash map usage
   - One-pass algorithm
   - Easy difficulty

2. **0009** - Palindrome Number
   - Simple mathematical approach
   - No data structure needed
   - Easy difficulty

### Intermediate
3. **0002** - Add Two Numbers
   - Linked list manipulation
   - Carry handling concept
   - Medium difficulty

4. **0003** - Longest Substring Without Repeating Characters
   - Sliding window introduction
   - Hash map with two pointers
   - Medium difficulty

5. **0006** - Zigzag Conversion
   - String manipulation
   - Pattern recognition
   - Medium difficulty

6. **0007** - Reverse Integer
   - Mathematical approach
   - Overflow handling
   - Medium difficulty

7. **0008** - String to Integer (atoi)
   - String parsing
   - State machine concepts
   - Medium difficulty

8. **0005** - Longest Palindromic Substring
   - Two approaches (expand vs DP)
   - Space-time tradeoff
   - Medium difficulty

### Advanced
9. **0004** - Median of Two Sorted Arrays
   - Binary search mastery
   - Partition approach
   - Hard difficulty

10. **0010** - Regular Expression Matching
    - Complex DP problem
    - Multiple solution approaches
    - Hard difficulty

## Implementation Statistics

- **Total Problems**: 10
- **Total Functions**: 15+ (including alternative implementations)
- **Test Cases**: 100+
- **Difficulty Distribution**: 2 Easy, 6 Medium, 2 Hard
- **Tag Distribution**: 
  - String: 5 problems
  - Math: 4 problems
  - Dynamic Programming: 3 problems
  - Others: 2 each

## Performance Summary

| Problem | Time Complexity | Space Complexity | Best Approach |
|---------|-----------------|------------------|----------------|
| 0001 | O(n) | O(n) | Hash Table |
| 0002 | O(max(m,n)) | O(max(m,n)) | Linked List Traversal |
| 0003 | O(n) | O(min(m,n)) | Sliding Window |
| 0004 | O(log min(m,n)) | O(1) | Binary Search |
| 0005 | O(n²) | O(1) / O(n²) | Expand Around Center / DP |
| 0006 | O(n) | O(n) | Row-based Grouping |
| 0007 | O(log x) | O(1) | Digit Extraction |
| 0008 | O(n) | O(1) | Character Parsing |
| 0009 | O(log x) | O(1) | Math Approach |
| 0010 | O(m×n) | O(m×n) | Dynamic Programming |

## Key Takeaways

1. **Hash Maps** are powerful for O(1) lookups
2. **Sliding Window** is efficient for substring problems
3. **Dynamic Programming** solves complex pattern matching
4. **Binary Search** achieves logarithmic complexity on sorted data
5. **Mathematical Approaches** can reduce space complexity
6. **Multiple Solutions** exist for most problems with different tradeoffs
7. **Testing** edge cases is crucial for correctness
8. **Overflow Handling** is important for numeric problems
