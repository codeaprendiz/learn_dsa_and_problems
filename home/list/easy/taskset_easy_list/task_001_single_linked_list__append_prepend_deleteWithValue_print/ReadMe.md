# Single Linked List

## Setup

```bash
$ go test


-------------- Test Deletion From a LinkedList ----------------------------

List status : 1 2 3 4 5           Delete head (1) - Delete: 1, Expected: [2 3 4 5], Result: [2 3 4 5]    --------- Pass
List status : 2 3 4 5           Delete middle (3) - Delete: 3, Expected: [2 4 5], Result: [2 4 5]    --------- Pass
List status : 2 4 5           Delete tail (5) - Delete: 5, Expected: [2 4], Result: [2 4]    --------- Pass
List status : 2 4           Delete non-existent (6) - Delete: 6, Expected: [2 4], Result: [2 4]    --------- Pass



-------------- Test Appending To a LinkedList ----------------------------

List status : 1       Append 1 - Inputs: [1], Expected: [1], Result: [1]    --------- Pass
List status : 1 2       Append 1, 2 - Inputs: [1 2], Expected: [1 2], Result: [1 2]    --------- Pass
List status : 1 2 3       Append 1, 2, 3 - Inputs: [1 2 3], Expected: [1 2 3], Result: [1 2 3]    --------- Pass
Overall Result


-------------- Test Prepending To a LinkedList ----------------------------

List status : 1       Prepend 1 - Inputs: [1], Expected: [1], Result: [1]    --------- Pass
List status : 1 2       Prepend 2, 1 - Inputs: [2 1], Expected: [1 2], Result: [1 2]    --------- Pass
List status : 1 2 3       Prepend 3, 2, 1 - Inputs: [3 2 1], Expected: [1 2 3], Result: [1 2 3]    --------- Pass


Overall Result
PASS
ok      single_linked_list      0.438s
```

## Node

```go

type Node struct {
    val  int   // This is the value that Node contains
    next *Node // This is the pointer to the next node
}
```

```bash
Node
+-------+
| val: X |   // 'X' represents the value stored in the node
| next  |-----> [Address of Next Node / nil]
+-------+
```

## LinkedList

```go
type LinkedList struct {
    head *Node // LinkedList is a struct that just contains a pointer to the first Node
}
```

```bash
# Empty LinkedList
LinkedList
+-------+
| head  |-----> [Address of First Node / nil]
+-------+

# LinkedList with two nodes
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
                  ^              ^
                  |              |
                Node 1         Node 2
```

## Appending to a list (Adding node to the end of list)

```go
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
```

Let's visually represent the steps of appending a node to a linked list using your `appendToList` function. We'll use references to `cur` and `newNode` in the graphical representation.

**Initial LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
                  ^              ^
                  |              |
                 cur           cur.next
```

- `cur` initially points to `Node 1`.

**Step 1: Create a New Node**

```bash
New Node
+-------+
| val:X |   // 'X' is the new value
| next  |-----> nil
+-------+
```

- A new node with `val: X` is created.

**Step 2: Traverse to the End of the List**

- Since `ll.head` is not `nil`, traverse the list using `cur` to find the last node.

**Step 3: `cur` Reaches Last Node**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
                                 ^
                                 |
                                cur
```

- `cur` is now at the last node (`Node 2`).

**Step 4: Append New Node at the End**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:X |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+
                                 ^              ^
                                 |              |
                                cur          newNode
```

- `cur.next` is set to `newNode`, adding it to the end of the list.

**Final LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:X |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+
```

- The new node is successfully appended to the end of the list.

## Deleting a Node from a list

```go
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
            cur = cur.next       // Advance cur to the next node
        } else {
            prev = cur
            cur = cur.next
        }
    }
}
```

Let's break down the node deletion process into steps with a purely graphical representation, including references to `cur`, `cur.next`, `prev`, and `prev.next`.

**Initial LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+
                  ^              ^              ^
                  |              |              |
                prev            cur           cur.next
```

**Step 1: Check `cur.val` Against `data`**

- `prev` at `Node 1`, `cur` at `Node 2`.

**Step 2: `cur.val` Matches `data`**

- `cur.val` is `2`, matches `data`.

**Step 3: Delete `Node 2` by Updating Pointers**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:3 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
                  ^              ^
                  |              |
                prev           cur.next
```

- Set `prev.next` to `cur.next`.

**Step 4: Advance `cur`**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:3 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
                  ^              ^
                  |              |
                prev            cur
```

- Move `cur` to `cur.next` (now at `Node 3`).

**Step 5: Continue Loop Until `cur` is `nil`**

- No more matching nodes; loop concludes.

**Final LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:3 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
```

- `Node 2` successfully deleted.

## Insert at the front of a list, Prepend to a list

```go
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
```

Let's create a graphical representation for the `prependToList` function, similar to how we illustrated the deletion process, but this time focusing on the insertion at the beginning of the list.

**Initial LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+
|       |      | val:1 |      | val:2 |
+-------+      | next  |----->| next  |-----> nil
               +-------+      +-------+
```

**Step 1: Create a New Node (e.g., with `val: X`)**

```bash
New Node
+--------+
| val:X  | 
| next   |-----> nil
+--------+
```

**Step 2: Point New Node's Next to Current Head**

```bash
New Node
+--------+
| val:X  | 
| next   |----->+-------+      +-------+
+--------+      | val:1 |      | val:2 |
                | next  |----->| next  |-----> nil
                +-------+      +-------+
```

**Step 3: Update Head to Point to New Node**

```bash
LinkedList
+-------+
| head  |----->+--------+     +-------+      +-------+
|       |      | val:X  |     | val:1 |      | val:2 |
+-------+      | next   |---->| next  |----->| next  |-----> nil
               +--------+     +-------+      +-------+
```

**Final LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+--------+     +-------+      +-------+
|       |      | val:X  |     | val:1 |      | val:2 |
+-------+      | next   |---->| next  |----->| next  |-----> nil
               +--------+     +-------+      +-------+
```

- The new node (`val: X`) is now the first node in the list, with its `next` pointing to the former first node (`val: 1`). The `head` of the list now points to the new node.
