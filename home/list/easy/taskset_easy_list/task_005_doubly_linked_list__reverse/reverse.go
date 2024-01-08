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
}

func (ll *LinkedList) appendNodeToListWithData(data int) {
	newNode := &Node{val: data, prev: nil, next: nil}

	if ll.head == nil {
		ll.head = newNode
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
}

// To reverse a linked list
func (ll *LinkedList) reverse() {
	var temp *Node
	cur := ll.head

	for cur != nil {
		// Swap the cur.prev and cur.next
		temp = cur.prev
		cur.prev = cur.next
		cur.next = temp

		// Move to the next node (which is current.prev after swapping)
		cur = cur.prev
	}

	// Before changing head, check for the cases like
	// empty list and list with only one node
	if temp != nil {
		ll.head = temp.prev
	}
}

// To display list in one line
func (ll *LinkedList) displayList() {
	cur := ll.head
	fmt.Printf("\nList Details : ")
	for cur != nil {
		fmt.Printf(" %v ", cur.val)
		cur = cur.next
	}
	fmt.Print("\n")
}

func main() {
	ll := LinkedList{}

	// Example usage
	ll.appendNodeToListWithData(1)
	ll.appendNodeToListWithData(2)
	ll.appendNodeToListWithData(3)

	fmt.Println("Original list:")
	ll.displayList()

	// Reverse the list
	ll.reverse()

	fmt.Println("Reversed list:")
	ll.displayList()
}
