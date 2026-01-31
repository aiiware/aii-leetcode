package data_structures

/*
707. Design Linked List

Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the 
current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the 
previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement the MyLinkedList class:
- MyLinkedList() Initializes the MyLinkedList object.
- int get(int index) Get the value of the index-th node in the linked list. If the index is invalid, return -1.
- void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
- void addAtTail(int val) Append a node of value val as the last element of the linked list.
- void addAtIndex(int index, int val) Add a node of value val before the index-th node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
- void deleteAtIndex(int index) Delete the index-th node in the linked list, if the index is valid.

Example 1:
Input
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
Output
[null, null, null, null, 2, null, 3]

Explanation
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
myLinkedList.get(1);              // return 2
myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
myLinkedList.get(1);              // return 3

Constraints:
- 0 <= index, val <= 1000
- At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex.
*/

/*
Difficulty: Medium
Tags: Linked List, Design
Companies: Amazon, Microsoft, Google, Apple, Bloomberg
*/

// ListNode represents a node in a singly linked list
type ListNode struct {
    val  int
    next *ListNode
}

// MyLinkedListSingly implements a singly linked list
type MyLinkedListSingly struct {
    head *ListNode
    size int
}

func ConstructorLinkedListSingly() MyLinkedListSingly {
    return MyLinkedListSingly{
        head: nil,
        size: 0,
    }
}

func (this *MyLinkedListSingly) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    
    curr := this.head
    for i := 0; i < index; i++ {
        curr = curr.next
    }
    return curr.val
}

func (this *MyLinkedListSingly) AddAtHead(val int) {
    newNode := &ListNode{
        val:  val,
        next: this.head,
    }
    this.head = newNode
    this.size++
}

func (this *MyLinkedListSingly) AddAtTail(val int) {
    newNode := &ListNode{val: val}
    
    if this.head == nil {
        this.head = newNode
    } else {
        curr := this.head
        for curr.next != nil {
            curr = curr.next
        }
        curr.next = newNode
    }
    this.size++
}

func (this *MyLinkedListSingly) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }
    
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    
    if index == this.size {
        this.AddAtTail(val)
        return
    }
    
    // Find the node before the insertion point
    prev := this.head
    for i := 0; i < index-1; i++ {
        prev = prev.next
    }
    
    newNode := &ListNode{
        val:  val,
        next: prev.next,
    }
    prev.next = newNode
    this.size++
}

func (this *MyLinkedListSingly) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    
    if index == 0 {
        this.head = this.head.next
    } else {
        prev := this.head
        for i := 0; i < index-1; i++ {
            prev = prev.next
        }
        prev.next = prev.next.next
    }
    this.size--
}

// Doubly linked list implementation
type DoublyListNode struct {
    val  int
    prev *DoublyListNode
    next *DoublyListNode
}

// MyLinkedListDoubly implements a doubly linked list with sentinel nodes
type MyLinkedListDoubly struct {
    head *DoublyListNode // sentinel head
    tail *DoublyListNode // sentinel tail
    size int
}

func ConstructorLinkedListDoubly() MyLinkedListDoubly {
    head := &DoublyListNode{val: -1}
    tail := &DoublyListNode{val: -1}
    head.next = tail
    tail.prev = head
    
    return MyLinkedListDoubly{
        head: head,
        tail: tail,
        size: 0,
    }
}

func (this *MyLinkedListDoubly) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    
    // Choose the closer end to start from
    var curr *DoublyListNode
    if index < this.size/2 {
        // Start from head
        curr = this.head.next
        for i := 0; i < index; i++ {
            curr = curr.next
        }
    } else {
        // Start from tail
        curr = this.tail.prev
        for i := 0; i < this.size-index-1; i++ {
            curr = curr.prev
        }
    }
    
    return curr.val
}

func (this *MyLinkedListDoubly) AddAtHead(val int) {
    newNode := &DoublyListNode{
        val:  val,
        prev: this.head,
        next: this.head.next,
    }
    
    this.head.next.prev = newNode
    this.head.next = newNode
    this.size++
}

func (this *MyLinkedListDoubly) AddAtTail(val int) {
    newNode := &DoublyListNode{
        val:  val,
        prev: this.tail.prev,
        next: this.tail,
    }
    
    this.tail.prev.next = newNode
    this.tail.prev = newNode
    this.size++
}

func (this *MyLinkedListDoubly) AddAtIndex(index int, val int) {
    if index < 0 || index > this.size {
        return
    }
    
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    
    if index == this.size {
        this.AddAtTail(val)
        return
    }
    
    // Find the node at the index
    var curr *DoublyListNode
    if index < this.size/2 {
        curr = this.head.next
        for i := 0; i < index; i++ {
            curr = curr.next
        }
    } else {
        curr = this.tail.prev
        for i := 0; i < this.size-index-1; i++ {
            curr = curr.prev
        }
    }
    
    newNode := &DoublyListNode{
        val:  val,
        prev: curr.prev,
        next: curr,
    }
    
    curr.prev.next = newNode
    curr.prev = newNode
    this.size++
}

func (this *MyLinkedListDoubly) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    
    // Find the node to delete
    var curr *DoublyListNode
    if index < this.size/2 {
        curr = this.head.next
        for i := 0; i < index; i++ {
            curr = curr.next
        }
    } else {
        curr = this.tail.prev
        for i := 0; i < this.size-index-1; i++ {
            curr = curr.prev
        }
    }
    
    curr.prev.next = curr.next
    curr.next.prev = curr.prev
    this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */