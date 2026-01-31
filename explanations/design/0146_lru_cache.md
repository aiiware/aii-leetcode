# 146. LRU Cache - Solution Explanation

## Problem Statement
Design a data structure that follows the constraints of a **Least Recently Used (LRU) cache**.

Implement the `LRUCache` class:
- `LRUCache(int capacity)` Initialize the LRU cache with positive size `capacity`
- `int get(int key)` Return the value of the `key` if it exists, otherwise return `-1`
- `void put(int key, int value)` Update the value of the `key` if it exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity`, evict the least recently used key.

All operations must run in **O(1)** average time complexity.

## Difficulty: Medium

## Key Insights
1. **O(1) Requirements**: Need constant time for get and put operations
2. **Tracking Recency**: Need to track most recently and least recently used items
3. **Data Structure Combination**: Hash map provides O(1) lookup, doubly linked list provides O(1) removal/insertion
4. **Capacity Management**: When capacity is exceeded, remove the tail of the list (LRU item)

## Solution Approaches

### Approach 1: Hash Map + Doubly Linked List (Standard)
**Time Complexity**: O(1) for both get and put
**Space Complexity**: O(capacity)

```go
type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node  // Most recently used
    tail     *Node  // Least recently used
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*Node),
        head:     &Node{key: -1, value: -1}, // Dummy head
        tail:     &Node{key: -1, value: -1}, // Dummy tail
    }
}

// Initialize doubly linked list with dummy nodes
func (this *LRUCache) initList() {
    this.head.next = this.tail
    this.tail.prev = this.head
}

// Remove node from its current position
func (this *LRUCache) remove(node *Node) {
    prev := node.prev
    next := node.next
    
    prev.next = next
    next.prev = prev
}

// Add node right after head (most recently used)
func (this *LRUCache) addToFront(node *Node) {
    node.prev = this.head
    node.next = this.head.next
    
    this.head.next.prev = node
    this.head.next = node
}

// Move node to front (mark as recently used)
func (this *LRUCache) moveToFront(node *Node) {
    this.remove(node)
    this.addToFront(node)
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
        // Move to front since it was recently accessed
        this.moveToFront(node)
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.cache[key]; exists {
        // Update value and move to front
        node.value = value
        this.moveToFront(node)
        return
    }
    
    // Create new node
    newNode := &Node{key: key, value: value}
    
    // Add to cache and list
    this.cache[key] = newNode
    this.addToFront(newNode)
    
    // Check capacity
    if len(this.cache) > this.capacity {
        // Remove LRU item (tail.prev since tail is dummy)
        lru := this.tail.prev
        this.remove(lru)
        delete(this.cache, lru.key)
    }
}
```

### Approach 2: Using container/list Package (Simpler)
**Time Complexity**: O(1) for both get and put
**Space Complexity**: O(capacity)

```go
import "container/list"

type LRUCacheSimple struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}

type entry struct {
    key   int
    value int
}

func ConstructorSimple(capacity int) LRUCacheSimple {
    return LRUCacheSimple{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        list:     list.New(),
    }
}

func (this *LRUCacheSimple) Get(key int) int {
    if elem, exists := this.cache[key]; exists {
        // Move to front (most recently used)
        this.list.MoveToFront(elem)
        return elem.Value.(*entry).value
    }
    return -1
}

func (this *LRUCacheSimple) Put(key int, value int) {
    if elem, exists := this.cache[key]; exists {
        // Update value and move to front
        elem.Value.(*entry).value = value
        this.list.MoveToFront(elem)
        return
    }
    
    // Create new entry
    newEntry := &entry{key: key, value: value}
    elem := this.list.PushFront(newEntry)
    this.cache[key] = elem
    
    // Check capacity
    if this.list.Len() > this.capacity {
        // Remove least recently used (back of list)
        lru := this.list.Back()
        this.list.Remove(lru)
        delete(this.cache, lru.Value.(*entry).key)
    }
}
```

## Step-by-Step Walkthrough

### Example: capacity = 2

**Initialization**:
```
cache = {}
list: head <-> tail
```

**Operation 1**: put(1, 10)
```
cache = {1: node1}
list: head <-> [1:10] <-> tail
```

**Operation 2**: put(2, 20)
```
cache = {1: node1, 2: node2}
list: head <-> [2:20] <-> [1:10] <-> tail
```

**Operation 3**: get(1) → returns 10
- Move node1 to front
```
cache = {1: node1, 2: node2}
list: head <-> [1:10] <-> [2:20] <-> tail
```

**Operation 4**: put(3, 30) → exceeds capacity
- Remove LRU (node2 with key=2)
- Add node3 to front
```
cache = {1: node1, 3: node3}
list: head <-> [3:30] <-> [1:10] <-> tail
```

**Operation 5**: get(2) → returns -1 (evicted)
**Operation 6**: get(1) → returns 10, moves to front
```
list: head <-> [1:10] <-> [3:30] <-> tail
```

## Complexity Analysis

### Time Complexity
- **Get Operation**: O(1) - Hash map lookup + constant time list operations
- **Put Operation**: O(1) - Hash map operations + constant time list operations
- **All operations** achieve O(1) average time complexity as required

### Space Complexity
- **Overall**: O(capacity) - Store up to `capacity` nodes in both hash map and linked list
- **Per Node**: O(1) - Each node stores key, value, and two pointers
- **Auxiliary**: O(1) - Constant extra space for pointers and dummy nodes

## Common Pitfalls
1. **Not using dummy nodes**: Edge cases when adding/removing first/last nodes
2. **Forgetting to update both data structures**: Must update hash map AND linked list
3. **Incorrect pointer updates**: Can cause memory leaks or incorrect list structure
4. **Not checking capacity after put**: Should evict before or after adding new item
5. **Using singly linked list**: Cannot remove arbitrary nodes in O(1) without previous pointer

## Optimization Tips
1. **Use dummy nodes**: Simplify edge case handling for head/tail operations
2. **Combine operations**: moveToFront = remove + addToFront
3. **Reuse nodes**: If updating existing key, reuse the node instead of creating new
4. **Batch operations**: If multiple puts, check capacity once at the end
5. **Memory pooling**: Pre-allocate nodes for better performance

## Edge Cases
1. **Capacity = 0**: Should reject all puts (or immediately evict)
2. **Capacity = 1**: Simple queue-like behavior
3. **Negative keys/values**: Problem states positive capacity, but keys/values can be any int
4. **Concurrent access**: Not thread-safe (would need synchronization)
5. **Large capacity**: Memory usage considerations

## Related Problems
- **460. LFU Cache** - Least Frequently Used cache (more complex)
- **588. Design In-Memory File System** - Similar design patterns
- **355. Design Twitter** - Uses similar data structures
- **716. Max Stack** - Another data structure design problem
- **380. Insert Delete GetRandom O(1)** - Similar O(1) requirements

## Practice Exercises
1. **Variation 1**: Implement MRU (Most Recently Used) cache instead
2. **Variation 2**: Add TTL (Time To Live) expiration to cache entries
3. **Variation 3**: Implement LFU (Least Frequently Used) cache
4. **Variation 4**: Make the cache thread-safe with mutexes
5. **Variation 5**: Add persistence (save/load from disk)
6. **Challenge**: Implement with only arrays (no pointers) for fixed-size cache

## Real-World Applications
1. **CPU Caches**: Hardware caches use LRU-like policies
2. **Database Query Caching**: Cache frequent query results
3. **Web Browser Cache**: Cache web pages and resources
4. **CDN Caching**: Content Delivery Networks cache popular content
5. **Memory Management**: Operating systems use similar algorithms for page replacement

## Design Patterns
This problem demonstrates several important design patterns:
1. **Decorator Pattern**: Cache wraps underlying data source
2. **Strategy Pattern**: Different eviction policies (LRU, LFU, FIFO)
3. **Factory Pattern**: Cache creation with different configurations
4. **Observer Pattern**: Notify when items are evicted (if needed)

## Additional Notes
- LRU is one of the most common cache eviction policies
- The combination of hash map and doubly linked list is classic for LRU implementation
- Many real-world systems use variations (LRU-K, ARC, etc.)
- Understanding LRU is fundamental for system design interviews
- This implementation can be extended to support various features like metrics, logging, or different eviction policies