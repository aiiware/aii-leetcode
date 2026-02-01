# 432. All O(1) Data Structure

## Problem Statement
Design a data structure that supports the following operations, all in **O(1)** average time complexity:
- `inc(key)`: Increments the count of `key` by 1
- `dec(key)`: Decrements the count of `key` by 1
- `getMaxKey()`: Returns one of the keys with maximal count
- `getMinKey()`: Returns one of the keys with minimal count

If the count of a key is decremented to 0, it should be removed from the data structure.

## Example
```
Input: ["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"]
       [[], ["hello"], ["hello"], [], [], ["world"], [], []]
Output: [null, null, null, "hello", "hello", null, "hello", "world"]

Explanation:
- After inc("hello"): {"hello": 1}
- After inc("hello"): {"hello": 2}
- getMaxKey() returns "hello"
- getMinKey() returns "hello"
- After inc("world"): {"hello": 2, "world": 1}
- getMaxKey() returns "hello"
- getMinKey() returns "world"
```

## Solution Approach: Doubly Linked List + Hash Map

### Intuition
We need O(1) operations for:
- Increment/decrement frequency of a key
- Get keys with max/min frequency

This suggests using:
- **Hash Map**: O(1) access to keys
- **Doubly Linked List**: O(1) insertion/removal, maintains frequency order

### Data Structure Design
```
AllOne {
    keyMap: map[string]*list.Element    // key → node in frequency list
    freqList: *list.List                // list of frequency nodes
}

freqNode {
    freq: int                           // frequency value
    keys: map[string]bool               // set of keys with this frequency
}
```

### Algorithm
1. **inc(key)**:
   - If key exists: move it to next frequency bucket (freq+1)
   - If key doesn't exist: add to frequency 1 bucket
   - Remove empty frequency nodes

2. **dec(key)**:
   - If key doesn't exist: return
   - If freq becomes 0: remove key completely
   - Otherwise: move to previous frequency bucket (freq-1)
   - Remove empty frequency nodes

3. **getMaxKey()**:
   - Return any key from the last node in freqList (highest frequency)

4. **getMinKey()**:
   - Return any key from the first node in freqList (lowest frequency)

### Complexity Analysis
- **Time Complexity**: O(1) average for all operations
- **Space Complexity**: O(n) where n is number of unique keys

### Code Implementation
```go
type AllOne struct {
    keyMap   map[string]*list.Element
    freqList *list.List
}

type freqNode struct {
    freq int
    keys map[string]bool
}

func AllOneConstructor() AllOne {
    return AllOne{
        keyMap:   make(map[string]*list.Element),
        freqList: list.New(),
    }
}

func (this *AllOne) Inc(key string) {
    // Implementation handles:
    // 1. Existing key: move to next frequency
    // 2. New key: add to frequency 1
    // 3. Clean up empty nodes
}

func (this *AllOne) Dec(key string) {
    // Implementation handles:
    // 1. Non-existent key: return
    // 2. Frequency becomes 0: remove key
    // 3. Frequency decreases: move to previous bucket
    // 4. Clean up empty nodes
}

func (this *AllOne) GetMaxKey() string {
    if this.freqList.Back() == nil {
        return ""
    }
    backNode := this.freqList.Back().Value.(*freqNode)
    for key := range backNode.keys {
        return key
    }
    return ""
}

func (this *AllOne) GetMinKey() string {
    if this.freqList.Front() == nil {
        return ""
    }
    frontNode := this.freqList.Front().Value.(*freqNode)
    for key := range frontNode.keys {
        return key
    }
    return ""
}
```

### Key Insights
1. **Frequency Buckets**: Group keys by frequency for O(1) max/min access
2. **Lazy Deletion**: Remove frequency nodes only when empty
3. **Set Operations**: Use map[string]bool for O(1) add/remove from frequency sets
4. **Order Maintenance**: List keeps frequencies in sorted order

### Edge Cases
1. **Empty Structure**: Return empty string for getMaxKey/getMinKey
2. **Single Key**: Same key for max and min
3. **Multiple Keys Same Frequency**: Any key can be returned
4. **Frequency 0**: Key should be completely removed
5. **Non-existent dec()**: Should do nothing

### Optimization Details
1. **Node Reuse**: Reuse existing frequency nodes when possible
2. **Direct Access**: Hash map provides O(1) access to key's current node
3. **Bidirectional Links**: Doubly linked list allows O(1) insert/remove

### Real-World Applications
- Real-time leaderboard systems
- Cache eviction policies (LFU cache)
- Frequency analysis in data streams
- Trending topic detection

### Related Problems
- 460. LFU Cache (similar frequency tracking)
- 146. LRU Cache (different eviction policy)
- 895. Maximum Frequency Stack