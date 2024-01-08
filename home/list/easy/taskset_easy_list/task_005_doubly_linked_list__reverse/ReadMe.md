# Reverse

## Run

```bash
$ go mod init reverse
go: creating new go.mod: module reverse
go: to add module requirements and sums:
        go mod tidy

$ go test
Reverse single element list - Inputs: [1], Expected: [1], Result: [1]    --------- Pass
Reverse multiple element list - Inputs: [1 2 3], Expected: [3 2 1], Result: [3 2 1]    --------- Pass
Reverse empty list - Inputs: [], Expected: [], Result: []    --------- Pass

Overall Result
PASS
ok      reverse 0.002s

$ go run reverse.go
Original list:

List Details :  1  2  3 
Reversed list:

List Details :  3  2  1
```

## reverse doubly linked list

You are absolutely correct, and I apologize for the confusion. Let's clarify the reversal process more accurately:

**Initial Doubly LinkedList:**

```bash
Doubly LinkedList (Before Reversal)
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
```

**Step 1: Swap `prev` and `next` for Each Node**

- Start with `cur` at `head` (`val:1`) and swap `cur.prev` and `cur.next`.
- Proceed to the next node (now `cur.prev` due to the swap).

```bash
During Reversal
+-------+
| head  |          +-------+      +-------+      +-------+
|       |--------->| val:1 |      | val:2 |      | val:3 |
+-------+          | prev  |----->| prev  |----->| prev  |-----> nil
          nil<-----| next  |<-----| next  |<-----| next  |
                   +-------+      +-------+      +-------+
                    ^               ^              ^
                    |               |              |
                  cur/temp        cur/temp       cur/temp
```

- Continue swapping `cur.prev` and `cur.next` for each node.
- After processing the last node (`val:3`), `cur` will point to `nil`, and `temp` will be on the last processed node (`val:1`).

**Step 2: Update Head of the List**

- After the loop, `temp` is at the original first node (`val:1`).
- The new head of the list should be the last node processed (`val:3`).

**Final Doubly LinkedList (After Reversal):**

```bash
Doubly LinkedList (After Reversal)
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:3 |      | val:2 |      | val:1 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               | prev  |<-----| prev  |<-----| prev  |
               +-------+      +-------+      +-------+
```

- Update `ll.head` to the node that `temp` was pointing to right before `cur` became `nil` (the node `val:3`).
- The list is now correctly reversed, with `val:3` as the new head.
