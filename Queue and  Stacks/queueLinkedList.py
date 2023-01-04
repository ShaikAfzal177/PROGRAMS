class _Node:
    __slots__ = '_element','_next' 
    def __init__(self,element,next):
        self._element=element 
        self._next =next 
class QueuesLinked():
    def __init__(self):
        self._head=None 
        self._tail=None
        self._size=0 
    def __len__(self):
        return self._size 
    def isempty(self):
        return self._size==0 
    def enqueue(self,e):
        newest=_Node(e,None)
        if self.isempty():
            self._head=newest
        else:
            self._tail._next=newest 
        self._tail=newest 
        self._size+=1
    def dequeue(self):
        if self.isempty():
            print("Queue is empty") 
            return 
        e=self._head._element 
        self._head=self._head._next 
        self._size -=1 
        if self.isempty():
            self._tail=None 
        return e 
    def first(self):
        if self.isempty():
            print("Queue is empty") 
            return 
        return self._head._element 
    def display(self):
        p=self._head 
        while p:
            print(p._element ,end="<--")
            p=p._next 
        print()




Q = QueuesLinked()
Q.enqueue(5)
Q.enqueue(3)
print('Queue:')
Q.display()
print('Queue Length:', len(Q))
ele = Q.dequeue()
print('Queue:')
Q.display()
print('Queue Length:', len(Q))
print('Removed Element:',ele)
print('Is Queue Empty:',Q.isempty())
Q.enqueue(7)
print('Queue:')
Q.display()
print('Queue Length:', len(Q))
Q.enqueue(12)
print('Queue:')
Q.display()
print('Queue Length:', len(Q))
ele = Q.dequeue()
print('Queue:')
Q.display()
print('Queue Length:', len(Q))
print('Removed Element:',ele)
print('First Element:',Q.first())
