# Function to generate all substrings using list comprehension
def generate_substrings_with_comprehension(test_str):
    return [test_str[i: j] for i in range(len(test_str)) for j in range(i + 1, len(test_str) + 1)]


# Function to generate all substrings using nested loops (non-comprehension)
def generate_substrings_with_loops(test_str):
    substrings = []
    # Outer loop for starting index of substring
    for i in range(len(test_str)):
        # Inner loop for ending index of substring
        for j in range(i + 1, len(test_str) + 1):
            # Slice the string from i to j
            substrings.append(test_str[i:j])
    return substrings


# Main function to demonstrate both methods
def main():
    # initializing string 
    test_str = "abcd"

    # printing original string 
    print("The original string is: " + str(test_str))

    # Generate substrings using list comprehension + slicing
    substrings_comprehension = generate_substrings_with_comprehension(test_str)
    print("Substrings using list comprehension:")
    print(substrings_comprehension)

    # Generate substrings using nested loops
    substrings_loops = generate_substrings_with_loops(test_str)
    print("Substrings using nested loops:")
    print(substrings_loops)


# Entry point
if __name__ == "__main__":
    main()
