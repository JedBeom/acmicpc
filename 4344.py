# 4344 평균은 넘겠지

n = int(input())

for _ in range(n):
    i = list(map(int,input().split()))
    avg = sum(i[1:])/i[0]
    cnt = 0
    for k in i[1:]:
        if k > avg:
            cnt += 1
    print(format(cnt/len(i[1:])*100, ".3f"), "%", sep="")
