// package main
package has_cycle

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
		return
	}

	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = newNode
}

func (ll *LinkedList) hasCycle() bool {
	slow := ll.head
	fast := ll.head

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
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
	ll := LinkedList{}

	// Append nodes to create a linked list
	for _, val := range []int{1, 2, 3, 4, 5} {
		ll.appendToList(val)
	}

	// Test for cycle (should be false)
	fmt.Println("Does the list have a cycle? ", ll.hasCycle())

	// Creating a cycle for testing: Pointing the next of the last node to the second node
	lastNode := ll.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = ll.head.next // Creating a cycle

	// Test for cycle (should be true)
	fmt.Println("Does the list have a cycle? ", ll.hasCycle())
}
