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

// Returns head of the listNode
func (l *LinkedList) Insert(v int, pos int) error {
	if pos < 0 {
		return errors.New("invalid position")
	}
	cNode := l.head
	// Loop for pos number of times..
	for range make([]struct{}, pos-1) {
		if cNode == nil {
			return fmt.Errorf("insertion index out of bounds position: %d with list of size: %d", pos, l.size)
		}
		cNode = cNode.next
	}

	newNode := node{val: v, next: cNode.next}
	cNode.next = &newNode

	l.size += 1
	return nil
}

func (l *LinkedList) GetListAsString() string {
	result := ""
	cNode := l.head
	for cNode != nil {
		result += fmt.Sprintf(" %v ->", cNode.val)
		cNode = cNode.next
	}
	result = strings.TrimSuffix(result, "->")
	return result
}
