
def checksum(a,b):
    return a+b,a-b,a*b

if __name__ == '__main__':
    a = int(input())
    b = int(input())
    c,d,e=checksum(a,b)
    print(c)
    print(d)
    print(e)
