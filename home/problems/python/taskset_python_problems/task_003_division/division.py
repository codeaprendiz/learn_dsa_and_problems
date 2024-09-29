def division(a,b):
    return int(a/b), float(a/b)

if __name__ == '__main__':
    a = int(input())
    b = int(input())
    c,d=division(a,b)
    print(c)
    print(d)