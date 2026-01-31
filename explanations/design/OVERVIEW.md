# Design Patterns Category Overview

## Introduction to Design Problems

Design problems in LeetCode focus on implementing data structures or systems that meet specific requirements, often with constraints on time and space complexity. These problems test your ability to design efficient, scalable systems and choose appropriate data structures.

### Key Characteristics of Design Problems
1. **API Design**: Define clear interfaces and method signatures
2. **Efficiency Requirements**: Often require O(1) or O(log n) operations
3. **Trade-off Analysis**: Balance between time complexity, space complexity, and implementation complexity
4. **Edge Case Handling**: Robust handling of boundary conditions
5. **Concurrency Considerations**: Some problems may require thread-safe designs

## Design Problem Patterns

### 1. Cache Design
**Characteristics**: Implement caching mechanisms with eviction policies
**Key Concepts**: LRU, LFU, TTL, cache invalidation
**Examples**:
- **0146 - LRU Cache**: Least Recently Used cache with O(1) operations
- **0460 - LFU Cache**: Least Frequently Used cache (more complex)
- **0588 - Design In-Memory File System**: Hierarchical cache-like structure

### 2. Stack/Queue Design
**Characteristics**: Implement stacks or queues with additional functionality
**Key Concepts**: Two-stack approach, monotonic structures, amortized analysis
**Examples**:
- **0155 - Min Stack**: Stack that supports getMin() in O(1)
- **0232 - Implement Queue using Stacks**: Queue with stack operations
- **0716 - Max Stack**: Stack that supports getMax() and popMax()

### 3. Data Stream Design
**Characteristics**: Process data arriving in streams with real-time queries
**Key Concepts**: Sliding windows, heaps, reservoirs, online algorithms
**Examples**:
- **0295 - Find Median from Data Stream**: Maintain median of streaming data
- **0346 - Moving Average from Data Stream**: Sliding window average
- **0352 - Data Stream as Disjoint Intervals**: Merge intervals from stream

### 4. Randomized Data Structures
**Characteristics**: Support random access or random operations efficiently
**Key Concepts**: Hash maps, arrays, reservoir sampling
**Examples**:
- **0380 - Insert Delete GetRandom O(1)**: Randomized set with O(1) operations
- **0381 - Insert Delete GetRandom O(1) - Duplicates allowed**: Extended version
- **0398 - Random Pick Index**: Reservoir sampling application

### 5. System Design
**Characteristics**: Design larger systems with multiple components
**Key Concepts**: Scalability, consistency, availability, partitioning
**Examples**:
- **0355 - Design Twitter**: Social media feed system
- **0359 - Logger Rate Limiter**: Rate limiting system
- **0642 - Design Search Autocomplete System**: Trie-based autocomplete

## Design Solution Strategies

### 1. Combine Multiple Data Structures
Often the key to O(1) operations is combining different data structures:

```go
// Example: LRU Cache combines hash map and doubly linked list
type LRUCache struct {
    capacity int
    cache    map[int]*Node      // O(1) lookup
    head     *Node              // Most recently used
    tail     *Node              // Least recently used
}
```

### 2. Trade Space for Time
Use additional data structures to cache results or maintain state:

```go
// Example: Min Stack stores pairs (value, currentMin)
type MinStack struct {
    stack []item
}

type item struct {
    val int
    min int  // Minimum value at or below this point
}
```

### 3. Amortized Analysis
Some operations may be expensive occasionally but cheap on average:

```go
// Example: Queue using two stacks
// Enqueue: O(1) amortized, Dequeue: O(1) amortized
type MyQueue struct {
    inStack  []int  // For enqueue operations
    outStack []int  // For dequeue operations
}
```

### 4. Lazy Evaluation
Compute results only when needed:

```go
// Example: Compute median only when getMedian() is called
type MedianFinder struct {
    small *MaxHeap  // Lower half (max heap)
    large *MinHeap  // Upper half (min heap)
    // Median computed on demand from heap tops
}
```

## Complexity Analysis

### Time Complexity Goals
- **Basic Operations**: O(1) for get, put, add, remove
- **Complex Operations**: O(log n) for ordered operations
- **Batch Operations**: O(n) amortized or O(n log n)

### Space Complexity Considerations
- **Auxiliary Structures**: Additional O(n) space often acceptable
- **Trade-offs**: More space for better time complexity
- **Memory Efficiency**: Consider object overhead in language-specific implementations

## Common Design Techniques

### 1. Two-Pointer/Two-Data-Structure Approach
Use two complementary data structures to maintain different views of the data.

### 2. Sentinel/Dummy Nodes
Simplify edge case handling in linked structures.

### 3. Circular Buffers
Efficient implementation of fixed-size queues or caches.

### 4. Index Mapping
Maintain bidirectional mapping between keys and indices.

### 5. Versioning/Timestamps
Track when data was last accessed or modified.

## Design vs Implementation

### Design Phase
1. **Requirements Analysis**: Understand constraints and use cases
2. **Interface Design**: Define clear API
3. **Data Structure Selection**: Choose appropriate structures
4. **Algorithm Design**: Plan operations and their complexity
5. **Trade-off Analysis**: Balance different requirements

### Implementation Phase
1. **Edge Case Handling**: Implement robust error handling
2. **Testing**: Verify correctness and performance
3. **Optimization**: Refine based on profiling
4. **Documentation**: Explain design decisions

## Learning Path

### Beginner Level
1. **0155 - Min Stack**: Simple two-stack design
2. **0232 - Implement Queue using Stacks**: Classic two-stack approach
3. **0225 - Implement Stack using Queues**: Reverse thinking

### Intermediate Level
1. **0146 - LRU Cache**: Hash map + doubly linked list
2. **0295 - Find Median from Data Stream**: Two-heap approach
3. **0380 - Insert Delete GetRandom O(1)**: Array + hash map
4. **0353 - Design Snake Game**: Game state management

### Advanced Level
1. **0460 - LFU Cache**: Complex frequency-based design
2. **0355 - Design Twitter**: System design with multiple components
3. **0642 - Design Search Autocomplete System**: Trie-based design
4. **0715 - Range Module**: Interval tree design

## Practice Problems by Difficulty

### Easy
- 0155 - Min Stack
- 0225 - Implement Stack using Queues
- 0232 - Implement Queue using Stacks
- 0346 - Moving Average from Data Stream

### Medium
- 0146 - LRU Cache
- 0295 - Find Median from Data Stream
- 0380 - Insert Delete GetRandom O(1)
- 0353 - Design Snake Game
- 0359 - Logger Rate Limiter
- 0622 - Design Circular Queue

### Hard
- 0297 - Serialize and Deserialize Binary Tree
- 0460 - LFU Cache
- 0355 - Design Twitter
- 0642 - Design Search Autocomplete System
- 0715 - Range Module
- 0716 - Max Stack

## Optimization Tips

### Memory Optimization
- Use arrays instead of linked lists when possible
- Reuse objects instead of creating new ones
- Consider memory alignment and padding
- Use primitive types instead of objects

### Time Optimization
- Precompute values when possible
- Use lazy evaluation for expensive operations
- Batch operations to reduce overhead
- Choose data structures with good cache locality

### Code Organization
- Separate interface from implementation
- Use helper methods for complex operations
- Document design decisions and trade-offs
- Write comprehensive tests for edge cases

## Common Mistakes

1. **Over-engineering**: Adding unnecessary complexity
2. **Under-engineering**: Missing important requirements
3. **Wrong Data Structure Choice**: Using inefficient structures
4. **Not Handling Concurrency**: Race conditions in multi-threaded scenarios
5. **Memory Leaks**: Not cleaning up unused resources
6. **API Design Flaws**: Unclear or inconsistent interfaces

## Real-World Applications

### 1. Database Systems
- Query optimization caches
- Index structures (B-trees, hash indexes)
- Transaction management

### 2. Operating Systems
- Page replacement algorithms (LRU, LFU)
- File system caches
- Process scheduling queues

### 3. Web Applications
- Session management
- Rate limiting
- Caching layers (Redis, Memcached)

### 4. Distributed Systems
- Consistent hashing
- Load balancing
- Distributed caches

### 5. Game Development
- Game state management
- AI decision trees
- Physics engines

## Design Patterns in Software Engineering

### 1. Creational Patterns
- **Singleton**: Ensure single instance (e.g., cache manager)
- **Factory**: Create objects without specifying exact class
- **Builder**: Construct complex objects step by step

### 2. Structural Patterns
- **Adapter**: Convert interface of a class
- **Decorator**: Add responsibilities dynamically
- **Composite**: Treat individual and composite objects uniformly

### 3. Behavioral Patterns
- **Iterator**: Access elements sequentially
- **Observer**: Notify dependents of state changes
- **Strategy**: Define family of algorithms

## Additional Resources

### Books
- "Designing Data-Intensive Applications" - Martin Kleppmann
- "System Design Interview" - Alex Xu
- "Clean Architecture" - Robert C. Martin

### Online Courses
- Grokking the System Design Interview (Educative)
- System Design Primer (GitHub)
- High Scalability (Blog)

### Practice Platforms
- LeetCode: Design problems by frequency
- Pramp: Mock system design interviews
- Interviewing.io: Practice with engineers

### Tools and Frameworks
- Redis: In-memory data structure store
- Apache Kafka: Distributed event streaming
- ZooKeeper: Distributed coordination service

---

*Last Updated: 2026-01-31*  
*Next: Create Arrays category overview*