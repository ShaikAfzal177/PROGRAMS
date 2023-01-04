class _Node:
    __slots__ ="_element","_next"
    def __init__(self,element, next):
        self._element=element 
        self._next=next 
class DEQueLinked: 
    def __init__(self):
        self._head =None 
        self._tail=None 
        self._size=0 
    def __len__(self):
        return self._size 
    def isempty(self):
        return self._size==0 
    def addlast(self,e):
        newest=_Node(e,None) 
        if self.isempty():
            self._head=newest 
        else:
            self._tail._next=newest 
        self._tail=newest 
        self._size+=1 
    def addfirst(self,e):
        newest=_Node(e,None) 
        if self.isempty():
            self._head=newest 
            self._tail=newest 
        else:
            newest._next=self._head 
            self._head =newest 
        self._size +=1 
    def removefirst(self):
        if self.isempty():
            print("Deque is empty") 
            return 
        e=self._head._element 
        self._head =self._head._next 
        self._size -=1 
        if self.isempty():
            self._tail=None 
        return e 
    def removelast(self):
        if self.isempty():
            print("Deque is empty") 
            return 
        p=self._head 
        i=1 
        while i< len(self)-1:
            p=p._next 
            i+=1 
        e=p._next._element 
        self._tail=p 
        self._tail._next=None 
        self._size-=1 
        return e 
    def first(self):
        if self.isempty():
            print("DEque is empty") 
            return 
        return self._head._element 
    def last(self):
        if self.isempty():
            print("Deque is empty") 
            return 
        return self._tail._element 
    def display(self):
        p=self._head 
        while p:
            print(p._element, end ="-->")
            p=p._next 
        print()



D = DEQueLinked()
D.addfirst(5)
D.addfirst(3)
D.addlast(7)
D.addlast(12)
D.display()
print('Length:',len(D))
print(D.removefirst())
print(D.removelast())
D.display()
print(D.first())
print(D.last())
