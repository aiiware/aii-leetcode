package data_structures

/*
706. Design HashMap

Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:
- MyHashMap() initializes the object with an empty map.
- void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
- int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
- void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.

Example 1:
Input
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Output
[null, null, null, 1, -1, null, 1, null, -1]

Explanation
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // The map is now [[1,1]]
myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (update the existing value)
myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]

Constraints:
- 0 <= key, value <= 10^6
- At most 10^4 calls will be made to put, get, and remove.
*/

/*
Difficulty: Easy
Tags: Array, Hash Table, Linked List, Design, Hash Function
Companies: Amazon, Google, Microsoft, Apple, Bloomberg
*/

// MyHashMap implements a simple hash map using separate chaining
type MyHashMap struct {
    buckets []*Entry
    size    int
}

// Entry represents a key-value pair in the hash map
type Entry struct {
    key   int
    value int
    next  *Entry
}

func ConstructorHashMap() MyHashMap {
    size := 1000 // Choose a prime number for better distribution
    return MyHashMap{
        buckets: make([]*Entry, size),
        size:    size,
    }
}

// hash function
func (this *MyHashMap) hash(key int) int {
    return key % this.size
}

func (this *MyHashMap) Put(key int, value int) {
    index := this.hash(key)
    
    // Check if key already exists
    curr := this.buckets[index]
    for curr != nil {
        if curr.key == key {
            curr.value = value // Update existing key
            return
        }
        curr = curr.next
    }
    
    // Insert new entry at the beginning of the chain
    newEntry := &Entry{
        key:   key,
        value: value,
        next:  this.buckets[index],
    }
    this.buckets[index] = newEntry
}

func (this *MyHashMap) Get(key int) int {
    index := this.hash(key)
    curr := this.buckets[index]
    
    for curr != nil {
        if curr.key == key {
            return curr.value
        }
        curr = curr.next
    }
    
    return -1
}

func (this *MyHashMap) Remove(key int) {
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

// Alternative implementation using array (simpler but less efficient for large ranges)
type MyHashMapArray struct {
    data []int
}

func ConstructorHashMapArray() MyHashMapArray {
    // Initialize with -1 since 0 <= key, value <= 10^6
    data := make([]int, 1000001)
    for i := range data {
        data[i] = -1
    }
    return MyHashMapArray{data: data}
}

func (this *MyHashMapArray) Put(key int, value int) {
    this.data[key] = value
}

func (this *MyHashMapArray) Get(key int) int {
    return this.data[key]
}

func (this *MyHashMapArray) Remove(key int) {
    this.data[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */