arr = [28,31,31,29,30]

arr.sort()

repeat = 0

miss = 0


for i in range(len(arr)):

    if i == len(arr)-1:

        break

    if arr[i] == arr[i+1]:

        repeat = arr[i]

    if arr[i+1] - arr[i]!= 1:

        miss = arr[i]+1

print(repeat, miss)