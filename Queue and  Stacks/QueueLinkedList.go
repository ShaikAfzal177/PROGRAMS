package main

import "fmt"

type Node struct {
	val  int
	next *Node
}
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (L *LinkedList) isempty() bool {
	return L.size == 0
}
func (L *LinkedList) display() {
	p := L.head
	for p != nil {
		fmt.Printf("%d <-", p.val)
		p = p.next
	}
	fmt.Println("")
}
func (L *LinkedList) enque(e int) {
	newest := &Node{e, nil}
	if L.isempty() {
		L.head = newest
	} else {
		L.tail.next = newest

	}
	L.tail = newest
	L.size++

}
func (L *LinkedList) deque() int {
	if L.isempty() {
		fmt.Println("Queue is empty")
		return 0
	}
	e := L.head.val
	L.head = L.head.next
	L.size--
	if L.isempty() {
		L.tail = nil
	}
	return e

}
func (L *LinkedList) top() int {
	if L.isempty() {
		fmt.Println("Queue is empty")
		return 0
	}
	e := L.tail.val
	return e
}

func main() {
	n := LinkedList{}
	fmt.Println(n)
	fmt.Println(n.isempty())
	n.enque(2)
	n.enque(3)
	n.enque(4)
	n.enque(5)
	n.enque(6)
	n.display()
	fmt.Println(n.deque())
	n.display()
	n.top()
}
