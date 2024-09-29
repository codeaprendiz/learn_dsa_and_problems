
def list_comprehension(x,y,z,n):
    # General Syntax of a List Comprehension:
    # [expression for item in iterable if condition]
    result = [[i,j,k] for i in range(x+1) for j in range(y+1) for k in range(z+1) if i+j+k != n ]
    return result

if __name__ == '__main__':
    x = int(input())
    y = int(input())
    z = int(input())
    n = int(input())
    result=list_comprehension(x,y,z,n)
    print(result)
