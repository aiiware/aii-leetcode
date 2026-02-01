# 0170. Two Sum III - Data structure design - Solution Explanation

## Problem Statement
Design a data structure that accepts a stream of integers and checks if it has a pair of integers that sum up to a particular value.

Implement the TwoSum class:
- TwoSum() Initializes the TwoSum object, with an empty data structure initially.
- void add(int number) Adds number to the data structure.
- boolean find(int value) Returns true if there exists any pair of numbers whose sum is equal to value, otherwise, it returns false.

Example 1:
Input:
["TwoSum", "add", "add", "add", "find", "find"]
[[], [1], [3], [5], [4], [7]]
Output:
[null, null, null, null, true, false]

Explanation:
TwoSum twoSum = new TwoSum();
twoSum.add(1);   // [] --> [1]
twoSum.add(3);   // [1] --> [1,3]
twoSum.add(5);   // [1,3] --> [1,3,5]
twoSum.find(4);  // 1 + 3 = 4, return true
twoSum.find(7);  // No two integers sum up to 7, return false

Constraints:
- -10^5 <= number <= 10^5
- -2^31 <= value <= 2^31 - 1
- At most 10^4 calls will be made to add and find.

## Difficulty: Easy

## Key Insights
1. **Frequency Map Approach**: Use a hash map to store the frequency of each number to handle duplicates efficiently.
2. **Complement Search**: For each number in the map, check if its complement (value - num) exists in the map.
3. **Duplicate Handling**: Special case when complement equals the current number - need at least two occurrences.

## Solution Approaches

### Approach 1: Hash Map with Frequency Count
**Time Complexity**: 
- Add: O(1) average case
- Find: O(n) worst case, where n is number of unique numbers
**Space Complexity**: O(n) for storing the frequency map

```go
// TwoSumIII implements a data structure for Two Sum III problem
type TwoSumIII struct {
    freq map[int]int // frequency map of numbers
}

// NewTwoSumIII initializes the TwoSumIII object
func NewTwoSumIII() TwoSumIII {
    return TwoSumIII{
        freq: make(map[int]int),
    }
}

// Add adds number to the data structure
func (this *TwoSumIII) Add(number int) {
    this.freq[number]++
}

// Find returns true if there exists any pair of numbers whose sum is equal to value
func (this *TwoSumIII) Find(value int) bool {
    for num, count := range this.freq {
        complement := value - num
        
        if complement == num {
            // Need at least two occurrences of the same number
            if count >= 2 {
                return true
            }
        } else {
            if _, exists := this.freq[complement]; exists {
                return true
            }
        }
    }
    return false
}
```

### Approach 2: Two Hash Maps (Optimized for Frequent Finds)
**Time Complexity**: 
- Add: O(1) average case  
- Find: O(1) average case with precomputed sums
**Space Complexity**: O(n²) worst case for storing all possible sums

```go
// Alternative: Store all possible sums when adding numbers
// This approach is better when find operations are much more frequent than add operations
```

## Step-by-Step Walkthrough

### Example:
```
Operations: add(1), add(3), add(5), find(4), find(7)
```

**Step 1**: Initialize empty frequency map
- `freq = {}`

**Step 2**: Add number 1
- `freq[1] = 1`
- Map becomes: `{1: 1}`

**Step 3**: Add number 3
- `freq[3] = 1`
- Map becomes: `{1: 1, 3: 1}`

**Step 4**: Add number 5
- `freq[5] = 1`
- Map becomes: `{1: 1, 3: 1, 5: 1}`

**Step 5**: Find value 4
- Iterate through map:
  - For num=1, complement=3 → complement exists in map → return true
- Result: true (1 + 3 = 4)

**Step 6**: Find value 7
- Iterate through map:
  - For num=1, complement=6 → not in map
  - For num=3, complement=4 → not in map  
  - For num=5, complement=2 → not in map
- Result: false (no pair sums to 7)

**Final Result**: [true, false]

## Complexity Analysis
### Time Complexity
- **Add Operation**: O(1) average case - hash map insertion
- **Find Operation**: O(n) worst case - iterate through all unique numbers
- **Amortized**: Depends on usage pattern - if finds are rare, this is efficient

### Space Complexity
- **Auxiliary Space**: O(n) - storing frequency map
- **Total Space**: O(n) - linear in number of unique numbers added

## Common Pitfalls
1. **Forgetting duplicate handling**: When complement equals current number, need to check if we have at least two occurrences.
2. **Integer overflow**: With constraints up to 2^31, complement calculation could overflow in some languages (not an issue in Go with int type).
3. **Inefficient find**: If find operations are very frequent, consider Approach 2 with precomputed sums.

## Optimization Tips
1. **Trade-off consideration**: Choose between Approach 1 (fast add, slow find) and Approach 2 (slow add, fast find) based on expected usage pattern.
2. **Early exit**: In find operation, we can return as soon as we find a valid pair.
3. **Memory optimization**: Approach 1 uses less memory, which is important for large streams.

## Edge Cases
1. **Empty data structure**: find should return false.
2. **Single element**: find should return false (need at least two numbers).
3. **Large numbers**: Handle numbers at constraint boundaries (-10^5 to 10^5).
4. **Many duplicates**: Frequency map handles this correctly.
5. **Zero values**: Works normally, complement calculation handles zero correctly.

## Related Problems
- [0001. Two Sum](https://leetcode.com/problems/two-sum/) - Classic Two Sum problem
- [0167. Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) - Two Sum with sorted input
- [0653. Two Sum IV - Input is a BST](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/) - Two Sum in Binary Search Tree

## Practice Exercises
1. **Exercise 1**: Modify to return all pairs that sum to the target value.
2. **Exercise 2**: Implement Approach 2 that precomputes all possible sums on add operation.
3. **Exercise 3**: Analyze the trade-off between Approach 1 and Approach 2 for different add/find ratios.

## Additional Notes
- This problem is a good example of the classic space-time tradeoff in algorithm design.
- The frequency map approach is generally preferred in interviews as it's simpler and uses reasonable space.
- In real-world scenarios, the choice between approaches depends on the expected ratio of add vs find operations.
- The problem constraints (10^4 operations) make both approaches feasible, but Approach 1 is more memory-efficient.