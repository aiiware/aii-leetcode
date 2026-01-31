package data_structures

/*
705. Design HashSet

Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:
- void add(key) Inserts the value key into the HashSet.
- bool contains(key) Returns whether the value key exists in the HashSet or not.
- void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

Example 1:
Input
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
Output
[null, null, null, true, false, null, true, null, false]

Explanation
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // return True
myHashSet.contains(3); // return False, (not found)
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // return True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // return False, (already removed)

Constraints:
- 0 <= key <= 10^6
- At most 10^4 calls will be made to add, remove, and contains.
*/

/*
Difficulty: Easy
Tags: Array, Hash Table, Linked List, Design, Hash Function
Companies: Amazon, Google, Microsoft, Apple, Bloomberg
*/

// MyHashSet implements a hash set using separate chaining
type MyHashSet struct {
    buckets []*Node
    size    int
}

// Node represents a node in the linked list chain
type Node struct {
    key  int
    next *Node
}

func ConstructorHashSet() MyHashSet {
    size := 1000 // Choose a prime number for better distribution
    return MyHashSet{
        buckets: make([]*Node, size),
        size:    size,
    }
}

// hash function
func (this *MyHashSet) hash(key int) int {
    return key % this.size
}

func (this *MyHashSet) Add(key int) {
    index := this.hash(key)
    
    // Check if key already exists
    curr := this.buckets[index]
    for curr != nil {
        if curr.key == key {
            return // Key already exists
        }
        curr = curr.next
    }
    
    // Add new node at the beginning of the chain
    newNode := &Node{
        key:  key,
        next: this.buckets[index],
    }
    this.buckets[index] = newNode
}

func (this *MyHashSet) Remove(key int) {
    index := this.hash(key)
    curr := this.buckets[index]
    
    // Handle case where key is at the head
    if curr != nil && curr.key == key {
        this.buckets[index] = curr.next
        return
    }
    
    // Search for key in the chain
    prev := curr
    for curr != nil {
        if curr.key == key {
            prev.next = curr.next
            return
        }
        prev = curr
        curr = curr.next
    }
}

func (this *MyHashSet) Contains(key int) bool {
    index := this.hash(key)
    curr := this.buckets[index]
    
    for curr != nil {
        if curr.key == key {
            return true
        }
        curr = curr.next
    }
    
    return false
}

// Alternative implementation using boolean array (simpler but less efficient for large ranges)
type MyHashSetArray struct {
    data []bool
}

func ConstructorHashSetArray() MyHashSetArray {
    // Initialize with false
    data := make([]bool, 1000001)
    return MyHashSetArray{data: data}
}

func (this *MyHashSetArray) Add(key int) {
    this.data[key] = true
}

func (this *MyHashSetArray) Remove(key int) {
    this.data[key] = false
}

func (this *MyHashSetArray) Contains(key int) bool {
    return this.data[key]
}

// BitSet implementation for memory efficiency
type MyHashSetBitSet struct {
    data []uint64
}

func ConstructorHashSetBitSet() MyHashSetBitSet {
    // We need enough bits for keys 0..10^6
    // 10^6 / 64 = 15625 (rounded up)
    size := 15625
    return MyHashSetBitSet{
        data: make([]uint64, size),
    }
}

func (this *MyHashSetBitSet) Add(key int) {
    index := key / 64
    bit := uint64(1) << (key % 64)
    this.data[index] |= bit
}

func (this *MyHashSetBitSet) Remove(key int) {
    index := key / 64
    bit := uint64(1) << (key % 64)
    this.data[index] &^= bit // AND NOT operation
}

func (this *MyHashSetBitSet) Contains(key int) bool {
    index := key / 64
    bit := uint64(1) << (key % 64)
    return (this.data[index] & bit) != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */