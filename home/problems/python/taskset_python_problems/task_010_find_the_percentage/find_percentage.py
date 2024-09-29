
def avg(marks):
    return f"{sum(marks)/len(marks):.2f}"


if __name__ == '__main__':
    n = int(input())
    student_marks = {}
    for _ in range(n):
        name, *line = input().split()
        scores = list(map(float, line))
        student_marks[name] = scores
    query_name = input()
    
    ans=avg(student_marks[query_name])
    print(ans)   