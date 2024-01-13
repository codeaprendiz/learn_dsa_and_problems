# [Binary Search Trees](https://www.geeksforgeeks.org/binary-search-tree-data-structure)

[www.geeksforgeeks.org » Binary Search Tree](https://www.geeksforgeeks.org/binary-search-tree-data-structure/)

## [Binary Search Tree (BST) Traversals – Inorder, Preorder, Post Order](https://www.geeksforgeeks.org/binary-search-tree-traversal-inorder-preorder-post-order)

```bash
        25
       /  \
      14   30
     / \   / \
    8  18 26  40
```

Inorder :  left subtree -> root -> right subtree

```bash
8 14 18 25 26 30 40
```

Preorder : root -> left subtree -> right subtree

```bash
25 14 8 18 30 26 40
```

Postorder : left subtree -> right subtree -> root

```bash
8 18 14 26 40 30 25
```

## Structure

```go
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
```

## Operations

### Insertion

```go
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
```

### Inorder Traversal

```go
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
```

1. **Scope of `result`**: The `result` slice is declared in the `InOrderTraversal` function, shared across all calls of the inner `inorder` function.

2. **Closure**: The `inorder` function is a closure, allowing it to access and modify the `result` slice from its enclosing `InOrderTraversal` scope.

3. **Appending to `result`**: Each call to `inorder` appends to the `result` slice, modifying the same slice since slices are reference types in Go.

4. **Concurrency and Race Conditions**: The `InOrderTraversal` function operates sequentially without concurrency, eliminating the risk of race conditions.

### PreOrder Traversal

```go
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
```

### PostOrder Traversal

```go
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
```

## Run

```bash
$ go mod init bst       
go: creating new go.mod: module bst
go: to add module requirements and sums:
        go mod tidy

$ go test        


-------------- Test Binary Search Tree Insertion ----------------------------
Before Insertion: []
Insert into an empty tree - Insert: 10, Expected: [10], Result: [10]    --------- Pass
Before Insertion: [20 30 40]
Insert smaller value - Insert: 15, Expected: [15 20 30 40], Result: [15 20 30 40]    --------- Pass
Before Insertion: [20 30 40]
Insert larger value - Insert: 35, Expected: [20 30 35 40], Result: [20 30 35 40]    --------- Pass
Before Insertion: [20 30 40]
Insert equal value - Insert: 30, Expected: [20 30 40], Result: [20 30 40]    --------- Pass


-------------- Test Binary Search Tree PreOrderTraversal ----------------------------
Single element - Inputs: [10], Expected: [10], Result: [10]    --------- Pass
Multiple elements - Inputs: [25 14 30 8 18 26 40], Expected: [25 14 8 18 30 26 40], Result: [25 14 8 18 30 26 40]    --------- Pass
Empty tree - Inputs: [], Expected: [], Result: []    --------- Pass


-------------- Test Binary Search Tree PostOrderTraversal ----------------------------
Single element - Inputs: [10], Expected: [10], Result: [10]    --------- Pass
Multiple elements - Inputs: [25 14 30 8 18 26 40], Expected: [8 18 14 26 40 30 25], Result: [8 18 14 26 40 30 25]    --------- Pass
Empty tree - Inputs: [], Expected: [], Result: []    --------- Pass


-------------- Test Binary Search Tree InorderTraversal ----------------------------
Single element - Inputs: [10], Expected: [10], Result: [10]    --------- Pass
Multiple elements - Inputs: [25 14 30 8 18 26 40], Expected: [8 14 18 25 26 30 40], Result: [8 14 18 25 26 30 40]    --------- Pass
Empty tree - Inputs: [], Expected: [], Result: []    --------- Pass

Overall Result
PASS
ok      bst     0.706s

$ go run bst.go
In order traversal : 
[8 14 18 25 26 30 40]
Pre-order traversal :
[25 14 8 18 30 26 40]
Post-order traversal : 
[8 18 14 26 40 30 25]
```
