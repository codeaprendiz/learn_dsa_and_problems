def printsquares(n):
    ans=[]
    i=0
    while i<n:
        ans.append(i*i)
        print(i*i)
        i+=1
    return ans

if __name__ == '__main__':
    n = int(input())
    printsquares(n)
        