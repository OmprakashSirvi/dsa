package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
	size int
	tail *node
}

func (l *LinkedList) GetHeadNode() *node {
	return l.head
}

func (l *LinkedList) Size() int {
	return l.size
}

// Returns the first value in list, returns error, if the list is empty
func (l *LinkedList) GetHeadVal() (int, error) {
	if l.head == nil {
		return 0, errors.New("list is empty")
	}
	return l.head.val, nil
}

// Create a new empty linkedList
func New() *LinkedList {
	head := &LinkedList{nil, 0, nil}
	return head
}

// Create new linked list with an entry: v
func NewWithVal(v int) *LinkedList {
	newNode := &LinkedList{&node{val: v, next: nil}, 1, nil}
	newNode.tail = newNode.head
	return newNode
}

// Appends l with value v at the end of l
func (l *LinkedList) Append(v int) {
	newNode := node{val: v, next: nil}
	// For empty linked lists
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		l.size = 1
		return
	}
	l.tail.next = &newNode
	l.tail = &newNode
	l.size += 1
}

func (l *LinkedList) InsertAtHead(v int) {
	newNode := &node{val: v, next: l.head}
	if l.size == 0 {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
}

func (l *LinkedList) Insert(v int, pos int) error {
	if pos < 0 {
		return errors.New("invalid position")
	}
	if pos > l.size {
		return fmt.Errorf("insertion index out of bounds: pos=%d, size=%d", pos, l.size)
	}

	// Determining whether this is an append operation
	if pos == l.size {
		l.Append(v)
		return nil
	}

	// Determine whether to insert at head
	if pos == 0 {
		l.InsertAtHead(v)
		return nil
	}

	newNode := &node{val: v}

	// Traverse to the node before position
	current := l.head
	for i := 0; i < pos-1; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
	l.size++
	return nil
}

// get string representation of the linked list
func (l *LinkedList) GetListAsString() string {
	return GetListAsStringFromNode(l.head)
}

// get strings representation of linked list from the given node
func GetListAsStringFromNode(head *node) string {
	resultBuilder := strings.Builder{}
	currNode := head
	for currNode != nil {
		resultBuilder.WriteString(fmt.Sprintf(" %v ->", currNode.val))
		currNode = currNode.next
	}

	result := strings.TrimSuffix(resultBuilder.String(), "->")
	return result
}

// Remove last element in the list
func (l *LinkedList) RemoveLast() {
	node := l.head
	// Do not do anything, cannot remove from an empty list
	if l.size == 0 {
		return
	}

	// Check if we just want to remove head
	if l.size == 1 {
		l.RemoveHead()
		return
	}

	for i := 0; i < l.size-2; i++ {
		node = node.next
	}

	// We will have the penultimate node now.
	node.next = nil
	l.tail = node
	l.size--
}

// Remove first element in the list
func (l *LinkedList) RemoveHead() {
	// If there are not elements in list, then just return
	if l.size == 0 {
		return
	}

	node := l.head.next
	l.head = node
	l.size--
	if l.size == 1 {
		// Here node will be nil
		l.tail = nil
	}
}

func (l *LinkedList) RemoveElement(pos int) error {
	if pos < 0 {
		return errors.New("invalid position")
	}
	if pos >= l.size {
		return fmt.Errorf("deletion at index out of bounds: pos=%d, size=%d", pos, l.size)
	}

	if pos == 0 {
		l.RemoveHead()
		return nil
	}

	if pos == l.size-1 {
		l.RemoveLast()
		return nil
	}

	node := l.head
	// Keep traversing till we reach the prev node to the deletion point
	for i := 0; i < pos-1; i++ {
		node = node.next
	}
	// We are sure that this will not panic
	node.next = node.next.next
	l.size--
	return nil
}

func (l *LinkedList) IsPresent(val int) bool {
	curr := l.head
	for curr != nil {
		if curr.val == val {
			return true
		}
		curr = curr.next
	}

	return false
}

func (l *LinkedList) GetMiddle() int {
	curr := l.head
	loop := l.size / 2
	for loop > 0 {
		curr = curr.next
		loop--
	}

	return curr.val
}

// Returns the head of the reverse linked list
func (l *LinkedList) Reverse() *node {
	var prev *node
	curr := l.head
	if curr == nil {
		// This current list is empty so returning nil
		return nil
	}

	// Loop till current node is empty
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.tail = l.head
	l.head = prev
	return l.head
}

// Returns the head of the reverse linked list
func ReverseRec(head *node) *node {
	// Base case
	if head == nil || head.next == nil {
		return head
	}

	newHead := ReverseRec(head.next)

	front := head.next
	front.next = head
	head.next = nil

	return newHead
}

func HasLoop(head *node) bool {
	fast := head
	slow := head

	// There should be at-least two values in the list
	if fast == nil || fast.next == nil {
		return false
	}

	for {
		fast = fast.next.next
		slow = slow.next

		// make sure we are not in the endgame..
		if fast == nil || fast.next == nil {
			return false
		}

		// Is the value same?
		// Equality check for two nodes
		if fast.val == slow.val && fast.next == slow.next {
			return true
		}
	}
}

func (l *LinkedList) HasLoop() bool {
	return HasLoop(l.head)
}

func (l *LinkedList) StartOfLoop() *node {
	fast := l.head
	slow := l.head

	if fast == nil || fast.next == nil {
		return nil
	}

	for {
		fast = fast.next.next
		slow = slow.next

		// This means we are not in loop
		if fast == nil || fast.next == nil {
			return nil
		}

		// Check if the slow and fast pointers are pointing to the same location or not..
		if fast.val == slow.val && fast.next == slow.next {
			// from here on we move these pointers one by one..
			break
		}
	}

	fast = l.head

	for {
		if fast.val == slow.val && fast.next == slow.next {
			return fast
		}

		fast = fast.next
		slow = slow.next

	}
}

func (l *LinkedList) Sort() *node {
	sortedHead := l.head.sort()
	l.head = sortedHead
	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	l.tail = curr
	return sortedHead
}

func (head *node) sort() *node {
	// base case
	if head == nil || head.next == nil {
		return head
	}

	slow := head
	fast := head

	for !(fast.next == nil || fast.next.next == nil) {
		slow = slow.next
		fast = fast.next.next
	}

	// break the list into two halves
	fast = slow.next
	slow.next = nil
	slow = head

	// time.Sleep(3 * time.Second)
	sortedSlow := slow.sort()
	sortedFast := fast.sort()

	mergedHead := merge(sortedSlow, sortedFast)

	return mergedHead
}

// we know that the two list's head is already sorted
// return the head of the merged list
func merge(head1 *node, head2 *node) *node {
	h1n := head1
	h2n := head2
	var temp *node

	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	// assign temp node to the head of the node which is smaller.
	if h1n.val <= h2n.val {
		temp = h1n
		h1n = h1n.next
	} else {
		temp = h2n
		h2n = h2n.next
	}
	retHead := temp

	for {
		// break conditions
		if h1n == nil {
			temp.next = h2n
			break
		}
		if h2n == nil {
			temp.next = h1n
			break
		}

		// merge conditions
		if h1n.val <= h2n.val {
			temp.next = h1n
			h1n = h1n.next
			temp = temp.next
		} else {
			temp.next = h2n
			h2n = h2n.next
			temp = temp.next
		}
	}

	return retHead
}
