# List Comprehension

- [List Comprehension](#list-comprehension)
  - [Docs](#docs)
  - [Theory](#theory)
    - [General Syntax of a List Comprehension](#general-syntax-of-a-list-comprehension)
    - [Components](#components)
    - [Your Specific Example](#your-specific-example)
    - [Equivalent Without List Comprehension](#equivalent-without-list-comprehension)
    - [How is the list comprehension different?](#how-is-the-list-comprehension-different)
  - [Run](#run)
  - [Unit Tests](#unit-tests)

<br>

## Docs

[hackerrank.com » List Comprehension](https://www.hackerrank.com/challenges/list-comprehensions/problem?isFullScreen=true)

<br>

## Theory

A list comprehension is a concise way to create lists in Python. It allows you to generate a list from an iterable (like a `range`, list, or any iterable structure) while optionally applying conditions (filters).

Here’s how a list comprehension works step by step

<br>

### General Syntax of a List Comprehension

```python
[expression for item in iterable if condition]
```

<br>

### Components

1. **expression**: This is the value that you want to include in the resulting list. In this case, it’s `[i, j, k]`.
2. **for item in iterable**: This part iterates over the iterable, like `range(x+1)` which generates numbers from `0` to `x`.
3. **if condition**: This is optional and filters the values. The result will only include items that satisfy this condition.

<br>

### Your Specific Example

```python
result = [[i, j, k] for i in range(x+1) for j in range(y+1) for k in range(z+1) if i + j + k != n]
```

Let’s break it down:

- `[[i, j, k]` is the **expression**—you want to generate a list containing the values `[i, j, k]`.
- `for i in range(x+1)` means you are iterating over the values `i` from `0` to `x`.
- `for j in range(y+1)` means for each value of `i`, you're iterating over `j` from `0` to `y`.
- `for k in range(z+1)` means for each pair of `i` and `j`, you're iterating over `k` from `0` to `z`.
- `if i + j + k != n` is the **condition**. You’re only including the values `[i, j, k]` where the sum `i + j + k` is **not equal** to `n`.

<br>

### Equivalent Without List Comprehension

Without list comprehension, you would achieve this using **nested loops**, like this:

```python
result = []
for i in range(x+1):
    for j in range(y+1):
        for k in range(z+1):
            if i + j + k != n:
                result.append([i, j, k])
```

<br>

### How is the list comprehension different?

- **Without list comprehension**, you use multiple lines of code and multiple explicit loops (`for` loops) to iterate and then append results to the list.
- **With list comprehension**, you achieve the same result in **one single line**. Python handles the loop and the condition internally, making the code cleaner and more concise.

<br>

## Run

```bash
$ python3 list_comprehension.py
1
1
1
3
[[0, 0, 0], [0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0]]
```

<br>

## Unit Tests

```bash
$ python3 -m unittest test_list_comprehension.py
.....
-----------------------------------------------------------------------

Ran 5 tests in 0.000s

OK
```
