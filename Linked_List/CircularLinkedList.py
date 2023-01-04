class _Node:

    __slots__ ="_element", "_next"

    def __init__(self, element, next):
        self._element=element 
        self._next=next 
class  CircularLinkedList:
    def __init__(self):
        self.head=None 
        self.tail=None 
        self.size=0 
    def __len__(self):
        return self.size 
    def isempty(self):
        return self.size==0 
    def addlast(self,e):
        newest=_Node(e, None)
        if self.isempty():
            newest._next=newest
            self.head=newest 

        else:
            newest._next =self.tail._next
            self.tail._next=newest

        self.tail=newest 
        self.size+=1 
    def addfirst(self,e):
        newest=_Node(e,None) 
        if self.isempty():
            newest._next=newest 
            self.head=newest 
            self.tail=newest 
        else:
            newest._next=self.head 
            self.tail._next=newest 
            self.head=newest 
            self.size+=1 
    def addany(self, e,position):
        newest=_Node(e,None) 
        p=self.head 
        i=1 
        while i<position-1:
            p=p._next 
            i+=1 
        newest._next=p._next 
        p._next=newest 
        self.size+=1
    def removefirst(self):
        if self.isempty():
            print("List is empty")
            return 
        e=self.head._element
        self.tail._next=self.head._next  
        self.head=self.head._next 
        self.size-=1 
        if self.isempty():
            self.head=None 
            self.tail=None 
        return e 
    def removelast(self):
        if self.isempty():
            print("list is empty")
            return 
        p=self.head 
        i=1 
        while i<len(self)-1:
            p=p._next 
            i+=1 
        self.tail=p 
        e=p._next._element 
        self.tail._next=self.head 
        self.size-=1 
        return e 
    def removeany(self, position):
        p=self.head 
        i=1 
        while i<position-1:
            p=p._next 
            i+=1 
        e=p._next._element 
        p._next=p._next._next 
        self.size-=1 
        return e 
    def display(self):
        p=self.head 
        i=0 
        while i<len(self):
            print(p._element, end ="-->")
            i+=1 
            p=p._next
        print()
    def search(self,key):
        p=self.head 
        index=0 
        while index<len(self):
            if p._element==key:
                return index 
            p=p._next 
            index+=1
        return -1





    
C = CircularLinkedList()
C.addlast(7)
C.addlast(4)
C.addlast(12)
C.addlast(8)
C.addfirst(4)
C.addany(1,3)
C.addlast(3)
C.display()
print('Size:',len(C))
ele = C.removeany(3)
C.display()
print('Size:',len(C))
print('Removed Element:',ele)

