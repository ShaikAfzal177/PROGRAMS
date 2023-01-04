def insertionsort(a):
    n=len(a)
    for i in range(1,n):
        cvalue=a[i]
        position=i 
        while position>0 and a[position-1]>cvalue:
            a[position]=a[position-1]
            position=position-1 
        a[position]=cvalue
    


A = [3,5,8,9,6,2]
print('Original Array:',A)
insertionsort(A)
print('Sorted Array:',A)
