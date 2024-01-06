package insert_after_address

// package main

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

// insertAfter inserts a new node with the given data after the specified node
func (ll *LinkedList) insertAfter(prev *Node, data int) {
	if prev == nil {
		return
	}

	cur := &Node{val: data, next: nil}
	cur.next = prev.next
	prev.next = cur
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
	fmt.Println()
}

func main() {
	ll := LinkedList{head: nil}

	// Append additional nodes to the list
	nodesToAdd := []int{1, 2, 3, 4}
	for _, val := range nodesToAdd {
		ll.appendToList(val)
	}

	fmt.Println("List after initial append:")
	ll.displayList()

	// Insert a node after the second node (which currently holds value 2)
	secondNode := ll.head.next
	ll.insertAfter(secondNode, 5)

	fmt.Println("List after inserting 5 after the second node:")
	ll.displayList()

	// Print detailed list
	fmt.Println("\n\nDetailed list:")
	ll.printList()
}
