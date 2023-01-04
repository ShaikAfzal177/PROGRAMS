def mergesort(a,left, right):
    if left<right:
        mid=(left+right)//2 
        mergesort(a,left,mid) 
        mergesort(a,mid+1,right)
        merge(a,left,mid,right) 
def merge(a,left,mid,right):
    i=left 
    j=mid+1 
    k=left 
    B=[0]*(right+1)
    while i<=mid and j<=right :
        if a[i]<a[j]:
            B[k]=a[i]
            i+=1 
        else:
            B[k]=a[j] 
            j+=1
        k=k+1 
    while i<=mid:
        B[k]=a[i]
        i+=1 
        k+=1 
    while j<=right:
        B[k]=a[j] 
        j+=1 
        k+=1 
    for x in range(left,right+1):
        a[x]=B[x]



A = [3, 5, 8, 9, 6, 2]
print('Original Array:',A)
mergesort(A,0,len(A)-1)
print('Sorted Array:',A)
