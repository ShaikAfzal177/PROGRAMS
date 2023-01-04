def ncd(n):
    n+=1 
    while True:
        z=str(n)
        x=0 
        for i in range(len(z)-1):
            if int(z[i])==int(z[i+1]):
                n+=1 
                x=1
            
        if x==0:
            return n
print(ncd(1123))