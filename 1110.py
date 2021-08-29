o = int(input())
i = 0

n = o
while True:
    i+=1
    k = n//10 + n%10
    print(f"k: {k}")
    nn = n[len(n)-1] + k[len(k)-1]
    print(f"nn: {nn}")
    
    if o == nn:
        print(i)
        break
    n = nn
    print(f"new n: {n}")


