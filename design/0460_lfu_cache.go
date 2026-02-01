package design

import "container/list"

// LFUCache implements a Least Frequently Used (LFU) cache
type LFUCache struct {
	capacity int
	minFreq  int
	// key -> (value, frequency, pointer to node in frequency list)
	keyMap map[int]*lfuCacheEntry
	// frequency -> list of keys with that frequency (most recent at front)
	freqMap map[int]*list.List
}

type lfuCacheEntry struct {
	value    int
	freq     int
	listElem *list.Element // pointer to node in freqMap[freq]
}

// ConstructorLFU creates an LFU cache with the given capacity
func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		minFreq:  0,
		keyMap:   make(map[int]*lfuCacheEntry),
		freqMap:  make(map[int]*list.List),
	}
}

// Get returns the value of the key if it exists, otherwise returns -1
func (this *LFUCache) Get(key int) int {
	if entry, exists := this.keyMap[key]; exists {
		// Update frequency
		this.incrementFreq(entry, key)
		return entry.value
	}
	return -1
}

// Put updates the value of the key if present, or inserts the key if not present
func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	
	if entry, exists := this.keyMap[key]; exists {
		// Update value and frequency
		entry.value = value
		this.incrementFreq(entry, key)
		return
	}
	
	// Need to evict if at capacity
	if len(this.keyMap) >= this.capacity {
		this.evict()
	}
	
	// Create new entry with frequency 1
	entry := &lfuCacheEntry{
		value: value,
		freq:  1,
	}
	
	// Add to freqMap[1]
	if this.freqMap[1] == nil {
		this.freqMap[1] = list.New()
	}
	entry.listElem = this.freqMap[1].PushFront(key)
	this.keyMap[key] = entry
	
	// Update minFreq
	this.minFreq = 1
}

// incrementFreq increases the frequency of a key and moves it to the appropriate list
func (this *LFUCache) incrementFreq(entry *lfuCacheEntry, key int) {
	oldFreq := entry.freq
	newFreq := oldFreq + 1
	
	// Remove from old frequency list
	oldList := this.freqMap[oldFreq]
	if oldList != nil {
		oldList.Remove(entry.listElem)
		if oldList.Len() == 0 {
			delete(this.freqMap, oldFreq)
			// If this was the minFreq list that became empty, update minFreq
			if oldFreq == this.minFreq {
				this.minFreq = newFreq
			}
		}
	}
	
	// Add to new frequency list
	if this.freqMap[newFreq] == nil {
		this.freqMap[newFreq] = list.New()
	}
	entry.listElem = this.freqMap[newFreq].PushFront(key)
	entry.freq = newFreq
}

// evict removes the least frequently used key (and least recently used if tie)
func (this *LFUCache) evict() {
	// Get the list for minFreq
	minList := this.freqMap[this.minFreq]
	if minList == nil {
		return
	}
	
	// Remove the least recently used (back of the list)
	backElem := minList.Back()
	if backElem == nil {
		return
	}
	
	key := backElem.Value.(int)
	minList.Remove(backElem)
	
	// Clean up empty list
	if minList.Len() == 0 {
		delete(this.freqMap, this.minFreq)
	}
	
	// Remove from keyMap
	delete(this.keyMap, key)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */