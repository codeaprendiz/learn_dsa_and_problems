# Single Linked List

## Setup

```bash
$ go mod init single_linked_list             
go: creating new go.mod: module single_linked_list
go: to add module requirements and sums:
        go mod tidy

$ go test

List status : 1 2 3 4 5           Delete head (1) - Delete: 1, Expected: [2 3 4 5], Result: [2 3 4 5]    --------- Pass
List status : 2 3 4 5           Delete middle (3) - Delete: 3, Expected: [2 4 5], Result: [2 4 5]    --------- Pass
List status : 2 4 5           Delete tail (5) - Delete: 5, Expected: [2 4], Result: [2 4]    --------- Pass
List status : 2 4           Delete non-existent (6) - Delete: 6, Expected: [2 4], Result: [2 4]    --------- Pass


List status : 1       Append 1 - Inputs: [1], Expected: [1], Result: [1]    --------- Pass
List status : 1 2       Append 1, 2 - Inputs: [1 2], Expected: [1 2], Result: [1 2]    --------- Pass
List status : 1 2 3       Append 1, 2, 3 - Inputs: [1 2 3], Expected: [1 2 3], Result: [1 2 3]    --------- Pass
Overall Result
PASS
ok      single_linked_list      1.124s
```
