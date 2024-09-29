# Sherlock and Anagrams

- [Sherlock and Anagrams](#sherlock-and-anagrams)
  - [Theory](#theory)
  - [Unit Tests](#unit-tests)

[hackerrank.com » Sherlock and Anagrams](https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem)

<br>

## Theory

**Understanding Anagram Pairs**

- If a particular sorted substring appears `count` times, how many distinct pairs of anagrams can we form?

- To calculate this, we use the combinatorics formula for choosing 2 items from `count`:

> [!TIP] Number of pairs = C(count, 2) = (count * (count - 1)) / 2

<br>

## Unit Tests

```bash
$ python3 -m unittest test_sherlock_and_anagrams.py 
.......
----------------------------------------------------------------------
Ran 7 tests in 0.000s

OK
```
