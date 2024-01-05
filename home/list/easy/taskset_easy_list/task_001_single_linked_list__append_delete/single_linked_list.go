package main

// package single_linked_list

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

func (ll *LinkedList) appendToList(data int) {
	newNode := &Node{val: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
	}

	cur := ll.head

	for cur.next != nil { // if cur OR ll.head is not nil, then cur.next OR ll.head.next must point to a valid address
		cur = cur.next
	}

	cur.next = newNode
}

func (ll *LinkedList) printList() {
	cur := ll.head

	for cur != nil {
		fmt.Println(" ", cur.val)
		cur = cur.next
	}
}

func main() {
	linkedList := LinkedList{head: nil}
	linkedList.printList()
	linkedList.appendToList(2)
	linkedList.printList()

}
