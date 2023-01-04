def selectionsort(a):
    n=len(a)
    for i in range(n-1):
        position=i 
        for j in range(i+1,n):
            if a[j]<a[position]: 
                position=j 
        temp=a[i] 
        a[i]=a[position]
        a[position]=temp


A = [3, 5, 8, 9, 6, 2]
print('Original Array:',A)
selectionsort(A)
print('Sorted Array:',A)
