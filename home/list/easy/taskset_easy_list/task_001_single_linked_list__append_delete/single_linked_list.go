package single_linked_list

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
func (ll *LinkedList) appendToList(data int) {
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

// Delete a Node
func (ll *LinkedList) deleteFromList(data int) {
	if ll.head == nil {
		return
	}

	if ll.head.val == data {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.val == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// Display a linked list
func (ll *LinkedList) displayList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
}

// func main() {
// 	linkedList := LinkedList{}
// 	fmt.Println("Linked List Before Appending")
// 	linkedList.displayList()
// 	fmt.Println("Linked List After Appending")
// 	linkedList.appendToList(1)
// 	linkedList.appendToList(2)
// 	linkedList.displayList()
// }
