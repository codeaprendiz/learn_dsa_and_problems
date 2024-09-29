# Python code to sort the tuples using second element 
# of sublist Inplace way to sort using sort()
def Sort(sub_li):
 
    # The key is set to sort using the second element of each sublist
    # 'lambda x: x[1]' means for each sublist 'x', we're using the second element (index 1) to sort 
    sub_li.sort(key = lambda x: x[1])
    return sub_li
 
def students_with_second_lowest_grade(student_list):
    Sort(student_list)
    l=len(student_list)
    lowest=student_list[0][1]
    ans=[]
    i=0
    while i<l and student_list[i][1]==lowest:
            i+=1
    
    if i==l:
        return
        
    second_lowest=student_list[i][1]
    while i<l and student_list[i][1]==second_lowest:
        ans.append(student_list[i][0])
        i+=1
        
    ans.sort()
    for name in ans:
        print(name)
        
if __name__ == '__main__':
    student_list = []
    for _ in range(int(input())):
        name = input()
        score = float(input())
        student_list.append([name, score])
        
    students_with_second_lowest_grade(student_list)
