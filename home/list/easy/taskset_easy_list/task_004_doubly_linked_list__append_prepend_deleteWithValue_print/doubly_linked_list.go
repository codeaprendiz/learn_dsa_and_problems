package main

import (
	"fmt"
)

type Node struct {
	val  int   // This is the value that Node contains
	prev *Node // This is the pointer to the prev Node
	next *Node // This is the pointer to the next Node
}

type LinkedList struct {
	head *Node // This is the pointer to the first Node
	tail *Node // This is the pointer to the last Node
}

func (ll *LinkedList) appendNodeToListWithData(data int) {
	newNode := &Node{val: data, prev: nil, next: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	// Traverse to the end of the list
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}

	// Link the new node
	cur.next = newNode
	newNode.prev = cur

	// Update the tail of the list
	ll.tail = newNode
}

// To delete from the linkedlist
func (ll *LinkedList) deleteNodeFromListWithData(data int) {

}

// To display list in one line
func (ll *LinkedList) displayList() {
	cur := ll.head
	for cur != nil {
		fmt.Printf(" %v ", cur.val)
		cur = cur.next
	}
}

func main() {
	ll := LinkedList{head: nil, tail: nil}
	ll.appendNodeToListWithData(1)
	ll.appendNodeToListWithData(2)
	ll.appendNodeToListWithData(3)
	ll.displayList()
}
