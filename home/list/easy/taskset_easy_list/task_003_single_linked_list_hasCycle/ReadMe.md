# Has Cycle

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

## has cycle

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
