# Has Cycle

[leetcode.com » Linked List Cycle](https://leetcode.com/problems/linked-list-cycle)

<br>

## Run

```bash
$ go mod init has_cycle                                       
go: creating new go.mod: module has_cycle
go: to add module requirements and sums:
        go mod tidy

$ go test            


-------------- Test Cycle Detection in a LinkedList ____________________________----

List before making a cycle (No cycle): 1 2 3 4 
      No cycle - MakeCycle: false, Expected: false, Result: false    --------- Pass

List before making a cycle (Cycle exists): 1 2 3 4 
      Cycle exists - MakeCycle: true, Expected: true, Result: true    --------- Pass



Overall Result
PASS
ok      has_cycle       0.978s
```

<br>

## has cycle

```go
func (ll *LinkedList) hasCycle() bool {
    slow := ll.head
    fast := ll.head

    for slow != nil && fast != nil && fast.next != nil {
        slow = slow.next        // this has to come before slow == fast condition, otherwise both are at head and you return true
        fast = fast.next.next

        if slow == fast {
            return true
        }
    }

    return false
}
```

**Initial LinkedList (No Cycle):**

```bash
LinkedList
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+
```

- Initially, the list has no cycle.

**Step 1: Initialize Two Pointers (`slow` and `fast`):**

```bash
+-------+
| head  |----->+-------+      +-------+      +-------+
|       |      | val:1 |      | val:2 |      | val:3 |
+-------+      | next  |----->| next  |----->| next  |-----> nil
               +-------+      +-------+      +-------+
                  ^              ^
                  |              |
                slow           fast
```

- Both `slow` and `fast` start at the head of the list.

**Step 2: Move `slow` by One Step and `fast` by Two Steps:**

- `slow` moves to the next node, while `fast` moves two nodes ahead.

**Step 3: Continue Moving `slow` and `fast`:**

- The process continues with `slow` moving one step and `fast` two steps in each iteration.
- If there's no cycle, `fast` will eventually reach the end (`nil`).

**LinkedList With a Cycle:**

```bash
LinkedList
+-------+      +-------+      +-------+
| head  |----->| val:1 |----->| val:2 |<----------------+
|       |      | next  |      | next  |      +-------+  |
+-------+      +-------+      +-------+----->| val:3 |  |
                                             | next  |--+
                                             +-------+
```

- In this scenario, the list has a cycle (`val:3`'s `next` points back to `val:2`).

**Step 4: Detecting the Cycle:**

- If `slow` and `fast` meet at the same node, a cycle is detected.

**Conclusion:**

- If `slow == fast` at any point, `hasCycle` returns `true`, indicating the presence of a cycle.
- If `fast` reaches the end of the list (`nil`), `hasCycle` returns `false`, indicating no cycle.

<br>

## Time And Space Complexity

let's analyze the `hasCycle` function in terms of its Time Complexity, Auxiliary Space Complexity, and Space Complexity:

1. **Time Complexity:**
   - The `hasCycle` function uses a for loop that runs as long as both `slow` and `fast` pointers do not encounter the end of the list (`nil`).
   - In the worst case (no cycle), `fast` will reach the end of the list in `O(n/2)` iterations, where `n` is the number of nodes. This simplifies to `O(n)`.
   - In the presence of a cycle, the loop will terminate earlier when `slow` and `fast` meet, but this will still happen in a number of steps proportional to `n`.
   - Thus, the Time Complexity of `hasCycle` is `O(n)`.

2. **Auxiliary Space Complexity:**
   - The function only uses two pointers (`slow` and `fast`), regardless of the size of the linked list.
   - No other additional data structures or recursive calls that grow with the size of the input are used.
   - Therefore, the Auxiliary Space Complexity is `O(1)`, as the space used is constant and does not scale with the input size.

3. **Space Complexity:**
   - Space Complexity includes both the space used by the input and the auxiliary space used by the algorithm.
   - Since the linked list itself is the input and the function does not modify or copy it, and the auxiliary space is constant, the overall Space Complexity is also `O(1)`.
   - The constant space for the two pointers is the only additional space used.

In summary, the `hasCycle` function is efficient in terms of both time and space, with a linear Time Complexity (`O(n)`) and constant Space Complexity (`O(1)`), making it well-suited for large linked lists.