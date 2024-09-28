// package bst
package main

import (
	"fmt"
)

// Node represents a node in the BST
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// BST represents Binary Search Tree
type BST struct {
	Root *Node
}

// Insert adds a new key to BST
func (bst *BST) Insert(key int) {
	bst.Root = insertNode(bst.Root, key)
}

// Helper function to insert a key into BST
func insertNode(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key}
	}

	if key < root.Key {
		root.Left = insertNode(root.Left, key)
	} else if key > root.Key {
		root.Right = insertNode(root.Right, key)
	}

	return root
}

// InOrderTraversal performs in-order traversal of the BST
func (bst *BST) InOrderTraversal() []int {
	var sliceOfInorder []int
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == nil {
			return
		}
		inorder(node.Left)
		sliceOfInorder = append(sliceOfInorder, node.Key)
		inorder(node.Right)
	}
	inorder(bst.Root)

	return sliceOfInorder
}

// PreOrderTraversal performs pre-order traversal of the BST
func (bst *BST) PreOrderTraversal() []int {
	var sliceOfPreOrder []int
	var preorder func(node *Node)
	preorder = func(node *Node) {
		if node == nil {
			return
		}
		sliceOfPreOrder = append(sliceOfPreOrder, node.Key)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(bst.Root)

	return sliceOfPreOrder
}

// PostOrderTraversal performs post-order traversal of the BST
func (bst *BST) PostOrderTraversal() []int {
	var sliceOfPostOrder []int
	var postorder func(node *Node)
	postorder = func(node *Node) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		sliceOfPostOrder = append(sliceOfPostOrder, node.Key)
	}

	postorder(bst.Root)

	return sliceOfPostOrder
}

func main() {
	bst := BST{Root: nil}

	// Inserting keys into BST
	keys := []int{25, 14, 30, 8, 18, 26, 40}
	for _, value := range keys {
		bst.Insert(value)
	}

	// In-order traversal to display the keys in sorted order
	fmt.Println("In order traversal : ")
	fmt.Println(bst.InOrderTraversal())

	// Pre-order traversal to display the keys
	fmt.Println("Pre-order traversal :")
	fmt.Println(bst.PreOrderTraversal())

	// Post-order traversal to display the keys
	fmt.Println("Post-order traversal : ")
	fmt.Println(bst.PostOrderTraversal())
}
