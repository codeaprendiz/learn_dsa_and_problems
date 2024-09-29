#!/bin/python3

import os

#
# Complete the 'sherlockAndAnagrams' function below.
#
# The function is expected to return an INTEGER.
# The function accepts STRING s as parameter.
#

from collections import defaultdict  # Importing defaultdict

def sherlockAndAnagrams(s):
    # Dictionary to store the frequency of sorted substrings
    substring_dict = defaultdict(int)

    # Generate all possible substrings
    for i in range(len(s)):
        for j in range(i + 1, len(s) + 1):
            # Sort the substring and use it as the key
            substring = ''.join(sorted(s[i:j]))
            substring_dict[substring] += 1

    # Calculate the number of anagrammatic pairs
    anagram_pairs = 0
    for count in substring_dict.values():
        # If count > 1, we can have anagram pairs
        anagram_pairs += (count * (count - 1)) // 2 # // (floor division) ensures the result is always an integer by rounding down if necessary. 
    
    return anagram_pairs

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    for q_itr in range(q):
        s = input()

        result = sherlockAndAnagrams(s)

        fptr.write(str(result) + '\n')

    fptr.close()
