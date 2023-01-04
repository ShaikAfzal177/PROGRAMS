def bubblesort(a):
    n=len(a) 
    for j in range(n-1,0,-1):
        for i in range(j):
            if a[i]>a[i+1]:
                temp=a[i] 
                a[i]=a[i+1]
                a[i+1]=temp

A = [3, 5, 8, 9, 6, 2]
print('Original Array:',A)
bubblesort(A)
print('Sorted Array:',A)
