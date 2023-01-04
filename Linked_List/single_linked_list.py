class  _Node :
    __slots__ = "_element","_next"
    def __init__(self,element,next):
        self._element=element 
        self._next=next 
class Linked_List:
    def __init__(self):
        self.head=None 
        self.tail=None 
        self.size=0
    def __len__(self):
        return self.size 
    def isempty(self):
        return self.size==0 
    def addlast(self, e):
        newest=_Node(e,None)
        if self.isempty():
            self.head=newest
        else:
            self.tail._next=newest
        
        self.tail=newest 
        self.size+=1
    def addfirst(self, e):
        newest=_Node(e,None) 
        if self.isempty():
            self.head=newest 
            self.tail=newest 
        else:
            newest._next=self.head 
            self.head=newest 
        self.size+=1 
    def addany(self,e,position):
        newest=_Node(e,None)
        p=self.head 
        i=1 
        while i<position-1:
            p=p._next 
            i+=1 

        newest._next=p._next 
        p._next=newest 
        self.size+=1 
    def removefist(self):
        if self.isempty():
            print("List is empty")
            return 
        e=self.head._element 
        self.head=self.head._next 
        self.size-=1 
        if self.isempty():
            self.tail=None 
        return e 
    def removelast(self):
        if self.isempty():
            print("List is empty")
            return 
        p=self.head 
        i=1 
        while i<len(self)-1:
            p=p._next 
            i+=1 
        self.tail=p 
        e=p._next._element 
        self.tail._next=None 
        self.size-=1 
        return e

    def removeany(self,position):
        if self.isempty():
            print("List is empty")
            return 
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
        while p:
            print(p._element, end="--->")
            p=p._next 
        print()
    def search(self,key):
        p=self.head 
        index=0 
        while p:
            if p._element==key:
                return index 
            p=p.next 
            index+=1 
        return -1 

L = Linked_List()
L.addfirst(4)
L.addlast(7)
L.addlast(4)
L.addlast(12)
L.addlast(8)
L.addlast(3)
L.display()
print('Size:',len(L))
ele = L.removeany(3)
L.display()
print('Size:',len(L))
print('Removed Element:',ele)

    







