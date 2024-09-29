def printfunc(n):
    ans=[]
    i=1
    while i<=n:
        print(i,end="")
        ans.append(i)
        i+=1    
    return ans

if __name__ == '__main__':
    n = int(input())
    printfunc(n)