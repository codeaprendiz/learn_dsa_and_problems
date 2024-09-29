#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'repeatedString' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts following parameters:
#  1. STRING s
#  2. LONG_INTEGER n
#

def repeatedString(s, n):
    # Write your code here
    ans=0
    i=1
    count_a=s.count("a")

    number_of_whole_substrings=int(n/len(s))

    ans=count_a*number_of_whole_substrings
    remainder=n%len(s)

    i=0
    while i<remainder:
        if s[i] == 'a':
            ans+=1
        i+=1

    return ans
            

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    n = int(input().strip())

    result = repeatedString(s, n)

    fptr.write(str(result) + '\n')

    fptr.close()
