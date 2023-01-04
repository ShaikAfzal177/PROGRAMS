n=int(input("enter the number:"))
for i in range(n):
    space=" "*(n-(i+1))
    star="*"*(i+1)
    print(space+star)
for i in range(n):
    s=" "*(n-1) 
    st="*"*(n-i)
    print(s+st)
