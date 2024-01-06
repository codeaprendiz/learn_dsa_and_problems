// package main

package single_linked_list

import (
	"fmt"
)

type Node struct {
	val  int   // This is the value that Node contains
	next *Node // This is the pointer to the next node
}

type LinkedList struct {
	head *Node // LinkedList is a struct that just contains a pointer to the first Node
}

// Insert at the end
func (ll *LinkedList) appendToList(data int) {

	newNode := &Node{val: data, next: nil} // newNode contains the address of the first Node

	if ll.head == nil { // if Pointer to the first Node is nil i.e. if head is nil
		ll.head = newNode
		return
	}

	cur := ll.head // cur has pointer to the first Node and it is definately not nil yet

	for cur.next != nil { // since head was not nil, now traverse until you find nil
		cur = cur.next
	}

	cur.next = newNode
}

// Insert at the begining of list
func (ll *LinkedList) prependToList(data int) {
	newNode := &Node{val: data, next: nil} // Create a new Node

	if ll.head == nil { // If the list is empty, make the new node the head
		ll.head = newNode
	} else { // If the list is not empty
		newNode.next = ll.head // Point the new node's 'next' to the current head
		ll.head = newNode      // Update the head of the list to the new node
	}
}

// Delete from list the value data, it can be at multiple places in the list
func (ll *LinkedList) deleteFromList(data int) {
	cur := ll.head.next // Pointer to the second Node if exists
	prev := ll.head     // Pointer to the first Node if exists

	if prev == nil { // if head is nil
		return
	}

	if prev.val == data {
		ll.head = cur // if the value is present at first Node, then point head to second node
		return
	}

	for cur != nil {

		if cur.val == data {
			prev.next = cur.next // Delete the current node
			cur = cur.next       // Advance cur to the next node, we do not return because data can be at multiple places
		} else {
			prev = cur
			cur = cur.next
		}
	}
}

// Function to print list with each Node details in separate lines
func (ll *LinkedList) printList() {
	cur := ll.head

	for cur != nil {
		fmt.Println("Node details : ")
		fmt.Println("val  : ", cur.val)
		fmt.Println("next : ", cur.next)
		cur = cur.next
	}
}

// Function to print the list details in one line
func (ll *LinkedList) displayList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
}

func main() {
	linkedList := LinkedList{head: nil}
	// linkedList.head.printNode()
	linkedList.printList()
	linkedList.appendToList(2)
	linkedList.appendToList(3)
	linkedList.printList()

}
