package design

/*
284. Peeking Iterator

Design an iterator that supports the peek operation on an existing iterator in addition 
to the hasNext and the next operations.

Implement the PeekingIterator class:
- PeekingIterator(Iterator<int> nums) Initializes the object with the given integer iterator iterator.
- int next() Returns the next element in the array and moves the pointer to the next element.
- boolean hasNext() Returns true if there are still elements in the array.
- int peek() Returns the next element in the array without moving the pointer.

Note: Each language may have a different implementation of the constructor and Iterator, 
but they all support the int next() and boolean hasNext() functions.

Example 1:
Input
["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 2, 2, 3, false]

Explanation
PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
peekingIterator.next();    // return 1, the pointer moves to the next element [1,2,3].
peekingIterator.peek();    // return 2, the pointer does not move [1,2,3].
peekingIterator.next();    // return 2, the pointer moves to the next element [1,2,3]
peekingIterator.next();    // return 3, the pointer moves to the next element [1,2,3]
peekingIterator.hasNext(); // return False

Constraints:
- 1 <= nums.length <= 1000
- 1 <= nums[i] <= 1000
- All the calls to next and peek are valid.
- At most 1000 calls will be made to next, hasNext, and peek.
*/

/*
Difficulty: Medium
Tags: Array, Design, Iterator
Companies: Google, Apple, Facebook, Amazon, Microsoft
*/

// Below is the interface for Iterator, which is already defined.
type Iterator struct {
    nums []int
    pos  int
}

func (this *Iterator) hasNext() bool {
    return this.pos < len(this.nums)
}

func (this *Iterator) next() int {
    if this.hasNext() {
        val := this.nums[this.pos]
        this.pos++
        return val
    }
    return -1
}

// PeekingIterator wraps an Iterator to support peek operation
type PeekingIterator struct {
    iter *Iterator
    nextVal *int  // cached next value for peek
    hasNextVal bool
}

func ConstructorPeekingIterator(iter *Iterator) *PeekingIterator {
    pi := &PeekingIterator{
        iter: iter,
        nextVal: nil,
        hasNextVal: false,
    }
    // Initialize by caching the first value
    if iter.hasNext() {
        val := iter.next()
        pi.nextVal = &val
        pi.hasNextVal = true
    }
    return pi
}

func (this *PeekingIterator) hasNext() bool {
    return this.hasNextVal
}

func (this *PeekingIterator) next() int {
    if !this.hasNextVal {
        return -1
    }
    
    val := *this.nextVal
    // Update cached value
    if this.iter.hasNext() {
        nextVal := this.iter.next()
        this.nextVal = &nextVal
        this.hasNextVal = true
    } else {
        this.nextVal = nil
        this.hasNextVal = false
    }
    
    return val
}

func (this *PeekingIterator) peek() int {
    if this.hasNextVal {
        return *this.nextVal
    }
    return -1
}

// Alternative implementation without modifying the original Iterator interface
type PeekingIteratorSimple struct {
    nums []int
    pos  int
    peeked bool
    peekVal int
}

func ConstructorPeekingIteratorSimple(nums []int) *PeekingIteratorSimple {
    return &PeekingIteratorSimple{
        nums: nums,
        pos: 0,
        peeked: false,
        peekVal: 0,
    }
}

func (this *PeekingIteratorSimple) hasNext() bool {
    return this.pos < len(this.nums) || this.peeked
}

func (this *PeekingIteratorSimple) next() int {
    if this.peeked {
        this.peeked = false
        return this.peekVal
    }
    
    if this.pos < len(this.nums) {
        val := this.nums[this.pos]
        this.pos++
        return val
    }
    
    return -1
}

func (this *PeekingIteratorSimple) peek() int {
    if this.peeked {
        return this.peekVal
    }
    
    if this.pos < len(this.nums) {
        this.peeked = true
        this.peekVal = this.nums[this.pos]
        return this.peekVal
    }
    
    return -1
}