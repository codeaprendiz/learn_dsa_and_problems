# Doubly Linked List

<br>

## Run

```bash
$ go mod init doubly_linked_list
go: creating new go.mod: module doubly_linked_list
go: to add module requirements and sums:
        go mod tidy

$ go test


-------------- Test Deletion From a Doubly LinkedList ----------------------------

List Details :  1  2  3  4 
Delete from middle of list - Delete: 3, Expected: [1 2 4], Result: [1 2 4]    --------- Pass

List Details :  1  2  3  4 
Delete head of list - Delete: 1, Expected: [2 3 4], Result: [2 3 4]    --------- Pass

List Details :  1  2  3  4 
Delete tail of list - Delete: 4, Expected: [1 2 3], Result: [1 2 3]    --------- Pass

List Details :  1  2  3  4 
Delete non-existent value - Delete: 5, Expected: [1 2 3 4], Result: [1 2 3 4]    --------- Pass

Overall Result


-------------- Test Appending To a Doubly LinkedList ----------------------------
Before Appending...
List Details : 
Append single element - Inputs: [1], Expected: [1], Result: [1]    --------- Pass
Before Appending...
List Details : 
Append multiple elements - Inputs: [1 2 3], Expected: [1 2 3], Result: [1 2 3]    --------- Pass

Overall Result


-------------- Test Prepending To a Doubly LinkedList ----------------------------

Before Prepending...

List Details : 
Prepend single element - Inputs: [1], Expected: [1], Result: [1]    --------- Pass

Before Prepending...

List Details : 
Prepend multiple elements - Inputs: [3 2 1], Expected: [1 2 3], Result: [1 2 3]    --------- Pass

Overall Result
PASS
ok      doubly_linked_list      0.613s




$ go run doubly_linked_list.go
 1  2  3
```

<br>

## Insert At The End of the List

```go
func (ll *LinkedList) appendToList(data int) {
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
```

Let's visually represent the steps of appending a node to a doubly linked list using the `appendToList` function. We'll include references to `cur`, `newNode`, and the `prev` and `next` pointers to clearly illustrate the structure of a doubly linked list.

**Initial Doubly LinkedList:**

```bash
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               | prev  |<-----| prev  |
               +-------+      +-------+
                  ^              ^
                  |              |
                Node 1         Node 2
```

- `Node 1` and `Node 2` are linked in both directions (`next` and `prev`).

**Step 1: Create a New Node (`val: X`)**

```bash
New Node
+-------+
| val:X |   // 'X' is the new value
| prev  |-----> nil
| next  |-----> nil
+-------+
```

- A new node with `val: X` is created with `nil` `prev` and `next`.

**Step 2: Traverse to the End of the List**

- Traverse the list using `cur` to find the last node (`Node 2` in this case).

**Step 3: `cur` Reaches Last Node**

```bash
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               | prev  |<-----| prev  |
               +-------+      +-------+
                                 ^
                                 |
                                cur
```

- `cur` points to the last node (`Node 2`).

**Step 4: Append New Node at the End**

```bash
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:X |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
                                 ^              ^
                                 |              |
                                cur          newNode
```

- `cur.next` is set to `newNode`, adding it to the end of the list.
- `newNode.prev` is set to `cur`, linking it back to the last node.

**Final Doubly LinkedList:**

```bash
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:X |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
```

- The new node (`val: X`) is successfully appended to the end of the list, with proper `prev` and `next` pointers maintaining the doubly linked structure.

<br>

## DeleteNode

```go
// To delete from the linkedlist
func (ll *LinkedList) deleteNodeFromListWithData(data int) {
    if ll.head == nil { // If the first Node is null
        return
    }

    if ll.head.val == data {
        ll.head = ll.head.next // head now points to the second Node, we have to delete the first Node
        if ll.head != nil {    // what is the second node is Nil
            ll.head.prev = nil // second node's prev pointer is no longer pointing to firstNode as it's deleted
        }
    }

    // Find the node to delete
    cur := ll.head
    for cur != nil && cur.val != data {
        cur = cur.next
    }

    // If node was not found
    if cur == nil {
        return
    }

    // Update the pointer to exclude the current node
    if cur.next != nil {
        cur.next.prev = cur.prev
    }
    if cur.prev != nil {
        cur.prev.next = cur.next
    }
}
```

You're correct. Let's revise the visualization to correctly depict the deletion of the node with `val:2` from a doubly linked list. We'll ensure that `cur` correctly points to the node being deleted and show the subsequent adjustments to the `prev` and `next` pointers of adjacent nodes.

**Initial Doubly LinkedList:**

```bash
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
```

- The list contains three nodes, and we aim to delete the node with `val:2`.

**Step 1: Traverse to Find the Node (`val:2`)**

- Traverse the list using `cur` to find the node with `val:2`.

```bash
Traversal
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
                                 ^
                                 |
                                cur
```

- `cur` points to the node with `val:2`.

**Step 2: Delete the Node (`val:2`)**

- Update `cur.next.prev` to `cur.prev` (link `val:3`'s `prev` to `val:1`).
- Update `cur.prev.next` to `cur.next` (link `val:1`'s `next` to `val:3`).

```bash
After Deletion
+-------+      +-------+      +-------+
| head  |----->| val:1 |----->| val:3 |-----> nil
|       |      | next  |      | next  |
+-------+      | prev  |<-----| prev  |
               +-------+      +-------+
```

- The node with `val:2` has been successfully deleted.
- The `prev` and `next` pointers of `val:1` and `val:3` are updated to exclude `val:2`.

**Final Doubly LinkedList (After Deletion of `val:2`):**

```bash
Doubly LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:3 |
+-------+      | next  |----->| next  |-----> nil
               | prev  |<-----| prev  |
               +-------+      +-------+
```

- The list now correctly reflects the removal of `val:2`, with adjacent nodes `val:1` and `val:3` linked directly to each other.

<br>

## Prepend

```go
// Prepend function, to add at the begining of the list
func (ll *LinkedList) prependNodeToListWithData(data int) {
    newNode := &Node{val: data, prev: nil, next: nil}

    if ll.head == nil {
        ll.head = newNode
        return
    }

    newNode.next = ll.head
    ll.head.prev = newNode
    ll.head = newNode
}
```

Let's visually represent the process of prepending a node to a doubly linked list using the `prependToList` function. This will illustrate how the function adds a new node at the beginning of the list and updates the `prev` and `next` pointers.

**Initial Doubly LinkedList:**

```bash
Doubly LinkedList (Before Prepension)
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               | prev  |<-----| prev  |
               +-------+      +-------+
```

- The list currently has two nodes: `val:1` and `val:2`.

**Step 1: Create a New Node (`val:X`)**

```bash
New Node
+-------+
| val:X |   // 'X' is the new value
| prev  |-----> nil
| next  |-----> nil
+-------+
```

- A new node with `val: X` is created.

**Step 2: Prepend the New Node to the List**

- The new node's `next` is set to the current head (`val:1`).
- The current head's `prev` is set to the new node.

```bash
Doubly LinkedList (After Prepension)
+-------+      +-------+      +-------+
| head  |----->| val:X |----->| val:1 |----->+-------+
|       |      | next  |      | next  |      | val:2 |
+-------+      | prev  |<-----| prev  |      | next  |-----> nil
               +-------+      +-------+      | prev  |
                                              +-------+
```

- `head` now points to the new node (`val:X`).

**Final Doubly LinkedList (After Prepension of `val:X`):**

```bash
Doubly LinkedList
+-------+      +-------+      +-------+
| head  |----->| val:X |----->| val:1 |----->| val:2 |-----> nil
|       |      | next  |      | next  |      | next  |
+-------+      | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
```

- The new node (`val: X`) is successfully prepended to the beginning of the list.
- The `prev` and `next` pointers are updated to maintain the integrity of the doubly linked structure.
