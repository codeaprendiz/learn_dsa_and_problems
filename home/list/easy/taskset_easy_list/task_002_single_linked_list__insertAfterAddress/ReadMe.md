# InsertAfterAddress

## Setup And Run

```bash
## Running using main
$ go run insert_after_address.go
List after initial append:
1 2 3 4 
List after inserting 5 after the second node:
1 2 5 3 4 


Detailed list:
Node details : 
val  :  1
next :  &{2 0x1400008e060}
Node details : 
val  :  2
next :  &{5 0x1400008e040}
Node details : 
val  :  5
next :  &{3 0x1400008e050}
Node details : 
val  :  3
next :  &{4 <nil>}
Node details : 
val  :  4
next :  <nil>
```

```bash
$ go mod init insert_after_address
go: creating new go.mod: module insert_after_address
go: to add module requirements and sums:
        go mod tidy

$ go test


-------------- Test Inserting After a Node in a LinkedList ----------------------------

List status : 1 2 3 4 
      Insert 5 after first node (1) - Insert: 5 After: 1, Expected: [1 5 2 3 4], Result: [1 5 2 3 4]    --------- Pass
List status : 1 2 3 4 
      Insert 6 after second node (2) - Insert: 6 After: 2, Expected: [1 2 6 3 4], Result: [1 2 6 3 4]    --------- Pass
List status : 1 2 3 4 
      Insert 7 after non-existent node (8) - Insert: 7 After: 8, Expected: [1 2 3 4], Result: [1 2 3 4]    --------- Pass


Overall Result
PASS
ok      insert_after_address    1.267s
```

## InsertAfterAddress Function

Let's visually represent the steps of inserting a node after a specified node in a linked list using the `insertAfter` function.

**Initial LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+
                  ^              ^
                  |              |
                prev           prev.next
```

- `prev` initially points to `Node 1` (for example).

**Step 1: Create a New Node (`val: X`)**

```bash
New Node
+-------+
| val:X |   // 'X' is the new value
| next  |-----> nil
+-------+
```

- A new node with `val: X` is created.

**Step 2: Insert New Node After `prev`**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:X |      | val:2 |
+-------+      | next  |----->| next  |----->| next  |----->+-------+
               +-------+      +-------+                       | val:3 |
                                                               | next  |-----> nil
                                                               +-------+
                  ^              ^
                  |              |
                prev          newNode
```

- `newNode.next` is set to `prev.next`.
- `prev.next` is updated to `newNode`.

**Final LinkedList:**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+      +-------+
|       |      | val:1 |      | val:X |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+      +-------+
```

- The new node (`val: X`) is successfully inserted after `Node 1` (`prev`).
