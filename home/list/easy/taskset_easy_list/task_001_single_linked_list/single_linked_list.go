package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Add a new node at the end of LinkedList
func (ll *LinkedList) append(data int) {
	newNode := &Node{val: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Display a linked list
func (ll *LinkedList) print() {
	current := ll.head
	for current != nil {
		fmt.Println("Data : ", current.val)
		current = current.next
	}
}

func main() {
	linkedList := LinkedList{}
	fmt.Println("Linked List Before Appending")
	linkedList.print()
	fmt.Println("Linked List After Appending")
	linkedList.append(1)
	linkedList.append(2)
	linkedList.print()
}
