def find_runner_up_score(n, arr):
    # remove duplicates
    unique_scores = set(arr)
    # remove max
    unique_scores.remove(max(unique_scores))
    return max(unique_scores)

if __name__ == '__main__':
    n = int(input())
    arr = map(int, input().split())
    ans = find_runner_up_score(n,arr)
    print(ans)