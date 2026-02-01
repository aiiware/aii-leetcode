package design

import "container/list"

// AllOne is a data structure that supports:
// - Inc(key): Increments the count of the key by 1.
// - Dec(key): Decrements the count of the key by 1.
// - GetMaxKey(): Returns one of the keys with maximal count.
// - GetMinKey(): Returns one of the keys with minimal count.
type AllOne struct {
	// Map from key to its node in the frequency list
	keyMap map[string]*list.Element
	// List of frequency nodes, each containing a set of keys with that frequency
	freqList *list.List
}

// freqNode represents a frequency bucket containing all keys with the same count
type freqNode struct {
	freq int
	keys map[string]bool
}

// AllOneConstructor initializes the AllOne data structure.
func AllOneConstructor() AllOne {
	return AllOne{
		keyMap:   make(map[string]*list.Element),
		freqList: list.New(),
	}
}

// Inc increments the count of the key by 1.
func (this *AllOne) Inc(key string) {
	if elem, exists := this.keyMap[key]; exists {
		// Key exists, move it to next frequency bucket
		node := elem.Value.(*freqNode)
		oldFreq := node.freq
		
		// Remove key from current node
		delete(node.keys, key)
		
		// Check if we need to create a new frequency node
		nextElem := elem.Next()
		if nextElem == nil || nextElem.Value.(*freqNode).freq != oldFreq+1 {
			// Create new frequency node
			newNode := &freqNode{
				freq: oldFreq + 1,
				keys: make(map[string]bool),
			}
			newNode.keys[key] = true
			nextElem = this.freqList.InsertAfter(newNode, elem)
		} else {
			// Add to existing frequency node
			nextNode := nextElem.Value.(*freqNode)
			nextNode.keys[key] = true
		}
		
		// Update keyMap
		this.keyMap[key] = nextElem
		
		// Remove current node if empty
		if len(node.keys) == 0 {
			this.freqList.Remove(elem)
		}
	} else {
		// Key doesn't exist, add to frequency 1 bucket
		front := this.freqList.Front()
		if front == nil || front.Value.(*freqNode).freq != 1 {
			// Create frequency 1 node
			newNode := &freqNode{
				freq: 1,
				keys: make(map[string]bool),
			}
			newNode.keys[key] = true
			front = this.freqList.PushFront(newNode)
		} else {
			// Add to existing frequency 1 node
			frontNode := front.Value.(*freqNode)
			frontNode.keys[key] = true
		}
		this.keyMap[key] = front
	}
}

// Dec decrements the count of the key by 1.
func (this *AllOne) Dec(key string) {
	elem, exists := this.keyMap[key]
	if !exists {
		return
	}
	
	node := elem.Value.(*freqNode)
	oldFreq := node.freq
	
	// Remove key from current node
	delete(node.keys, key)
	
	if oldFreq == 1 {
		// Frequency becomes 0, remove key completely
		delete(this.keyMap, key)
	} else {
		// Move to lower frequency bucket
		prevElem := elem.Prev()
		if prevElem == nil || prevElem.Value.(*freqNode).freq != oldFreq-1 {
			// Create new frequency node
			newNode := &freqNode{
				freq: oldFreq - 1,
				keys: make(map[string]bool),
			}
			newNode.keys[key] = true
			prevElem = this.freqList.InsertBefore(newNode, elem)
		} else {
			// Add to existing frequency node
			prevNode := prevElem.Value.(*freqNode)
			prevNode.keys[key] = true
		}
		
		// Update keyMap
		this.keyMap[key] = prevElem
	}
	
	// Remove current node if empty
	if len(node.keys) == 0 {
		this.freqList.Remove(elem)
	}
}

// GetMaxKey returns one of the keys with maximal count.
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

// GetMinKey returns one of the keys with minimal count.
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

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := AllOneConstructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */