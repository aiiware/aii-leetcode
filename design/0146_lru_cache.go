package design

// LRUCache implements a Least Recently Used (LRU) cache
// Difficulty: Medium
// Tags: Hash Table, Linked List, Design, Doubly-Linked List
//
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
//
// Implement the LRUCache class:
// - LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
// - int get(int key) Return the value of the key if the key exists, otherwise return -1.
// - void put(int key, int value) Update the value of the key if the key exists.
//   Otherwise, add the key-value pair to the cache. If the number of keys exceeds the
//   capacity from this operation, evict the least recently used key.
//
// The functions get and put must each run in O(1) average time complexity.
//
// Example 1:
// Input:
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output: [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// Explanation:
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // cache is {1=1}
// lRUCache.put(2, 2); // cache is {1=1, 2=2}
// lRUCache.get(1);    // return 1
// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
// lRUCache.get(2);    // returns -1 (not found)
// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
// lRUCache.get(1);    // return -1 (not found)
// lRUCache.get(3);    // return 3
// lRUCache.get(4);    // return 4
//
// Constraints:
// - 1 <= capacity <= 3000
// - 0 <= key <= 10^4
// - 0 <= value <= 10^5
// - At most 2 * 10^5 calls will be made to get and put.
//
// Time complexity: O(1) for both get and put operations
// Space complexity: O(capacity)

// LRUNode represents a node in the doubly linked list
type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

// LRUCache implements the LRU cache using hash map and doubly linked list
type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*LRUNode
	head     *LRUNode // dummy head (most recently used)
	tail     *LRUNode // dummy tail (least recently used)
}

// ConstructorLRUCache creates a new LRUCache with the given capacity
func ConstructorLRUCache(capacity int) LRUCache {
	// Create dummy head and tail nodes
	head := &LRUNode{}
	tail := &LRUNode{}
	
	// Connect head and tail
	head.next = tail
	tail.prev = head
	
	return LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*LRUNode),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value of the key if it exists, otherwise returns -1
func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		// Move the accessed node to the front (most recently used)
		lru.moveToFront(node)
		return node.value
	}
	return -1
}

// Put adds or updates a key-value pair in the cache
func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		// Update existing node
		node.value = value
		lru.moveToFront(node)
	} else {
		// Create new node
		newNode := &LRUNode{
			key:   key,
			value: value,
		}
		
		// Add to cache and list
		lru.cache[key] = newNode
		lru.addToFront(newNode)
		lru.size++
		
		// If capacity exceeded, remove least recently used
		if lru.size > lru.capacity {
			lru.removeLRU()
		}
	}
}

// moveToFront moves a node to the front of the list (most recently used)
func (lru *LRUCache) moveToFront(node *LRUNode) {
	// Remove node from its current position
	lru.removeNode(node)
	// Add it to the front
	lru.addToFront(node)
}

// addToFront adds a node to the front of the list
func (lru *LRUCache) addToFront(node *LRUNode) {
	// Insert node between head and head.next
	node.prev = lru.head
	node.next = lru.head.next
	
	// Update surrounding nodes
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode removes a node from the list
func (lru *LRUCache) removeNode(node *LRUNode) {
	// Update surrounding nodes
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeLRU removes the least recently used node (from tail)
func (lru *LRUCache) removeLRU() {
	// The node before tail is the LRU node
	lruNode := lru.tail.prev
	
	// Remove from list
	lru.removeNode(lruNode)
	
	// Remove from cache
	delete(lru.cache, lruNode.key)
	lru.size--
}

// LRUCacheSimple is a simpler implementation using container/list
// This version uses Go's built-in list package for cleaner code
type LRUCacheSimple struct {
	capacity int
	cache    map[int]*cacheEntry
	list     *doublyLinkedList
}

type cacheEntry struct {
	key   int
	value int
	node  *listNode
}

// NewLRUCacheSimple creates a new LRUCacheSimple
func NewLRUCacheSimple(capacity int) *LRUCacheSimple {
	return &LRUCacheSimple{
		capacity: capacity,
		cache:    make(map[int]*cacheEntry),
		list:     newDoublyLinkedList(),
	}
}

// Get returns the value for the key
func (lru *LRUCacheSimple) Get(key int) int {
	if entry, exists := lru.cache[key]; exists {
		// Move to front (most recently used)
		lru.list.moveToFront(entry.node)
		return entry.value
	}
	return -1
}

// Put adds or updates a key-value pair
func (lru *LRUCacheSimple) Put(key int, value int) {
	if entry, exists := lru.cache[key]; exists {
		// Update value and move to front
		entry.value = value
		lru.list.moveToFront(entry.node)
	} else {
		// Create new entry
		node := lru.list.pushFront(key)
		entry := &cacheEntry{
			key:   key,
			value: value,
			node:  node,
		}
		lru.cache[key] = entry
		
		// Remove LRU if capacity exceeded
		if len(lru.cache) > lru.capacity {
			// Remove from tail (least recently used)
			lruNode := lru.list.tail.prev
			delete(lru.cache, lruNode.key)
			lru.list.remove(lruNode)
		}
	}
}

// Custom doubly linked list implementation for LRUCacheSimple
type listNode struct {
	key  int
	prev *listNode
	next *listNode
}

type doublyLinkedList struct {
	head *listNode // dummy head
	tail *listNode // dummy tail
}

func newDoublyLinkedList() *doublyLinkedList {
	head := &listNode{}
	tail := &listNode{}
	head.next = tail
	tail.prev = head
	return &doublyLinkedList{
		head: head,
		tail: tail,
	}
}

func (dll *doublyLinkedList) pushFront(key int) *listNode {
	node := &listNode{key: key}
	
	// Insert after head
	node.prev = dll.head
	node.next = dll.head.next
	
	dll.head.next.prev = node
	dll.head.next = node
	
	return node
}

func (dll *doublyLinkedList) moveToFront(node *listNode) {
	// Remove from current position
	dll.remove(node)
	
	// Insert after head
	node.prev = dll.head
	node.next = dll.head.next
	
	dll.head.next.prev = node
	dll.head.next = node
}

func (dll *doublyLinkedList) remove(node *listNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}