package leetcode

/*
170. Two Sum III - Data structure design

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

Difficulty: Easy
Tags: Hash Table, Design, Data Stream
Companies: LinkedIn
*/

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

/**
 * Your TwoSumIII object will be instantiated and called as such:
 * obj := NewTwoSumIII();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */