#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'isBalanced' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def isBalanced(s):
    # Stack to keep track of opening brackets
    stack = []
    # Dictionary to map closing brackets to their corresponding opening brackets
    matching_brackets = {')': '(', '}': '{', ']': '['}
    
    for char in s:
        if char in '({[':  # If it's an opening bracket, push to stack
            stack.append(char)
        elif char in ')}]':  # If it's a closing bracket
            # If stack is empty or top of stack doesn't match, it's unbalanced
            if not stack or stack.pop() != matching_brackets[char]:
                return "NO"
    
    # If the stack is empty, the brackets are balanced, ternary expression (or conditional expression
    return "YES" if not stack else "NO"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input().strip())

    for t_itr in range(t):
        s = input()

        result = isBalanced(s)

        fptr.write(result + '\n')

    fptr.close()
