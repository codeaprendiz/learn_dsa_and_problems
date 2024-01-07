# Doubly Linked List

## Run

```bash
$ go mod init doubly_linked_list
go: creating new go.mod: module doubly_linked_list
go: to add module requirements and sums:
        go mod tidy


```

## Insert At The End of the List

```go
func (ll *LinkedList) appendToList(data int) {
    newNode := &Node{val: data, prev: nil, next: nil}

    if ll.head == nil {
        ll.head = newNode
        ll.tail = newNode
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

    // Update the tail of the list
    ll.tail = newNode
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
                                cur          newNode / tail
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
                                                 ^
                                                 |
                                                tail
```

- The new node (`val: X`) is successfully appended to the end of the list, with proper `prev` and `next` pointers maintaining the doubly linked structure.
