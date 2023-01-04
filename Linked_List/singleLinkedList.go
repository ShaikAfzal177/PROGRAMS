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

func (L *LinkedList) isEmpty() bool {
	return L.size == 0
}
func (L *LinkedList) display() {
	p := L.head
	for p != nil {
		fmt.Printf("%d ->", p.val)
		p = p.next
	}
	fmt.Println("")
}
func (L *LinkedList) addLast(val int) {
	newest := &Node{val, nil}
	if L.isEmpty() {
		L.head = newest
	} else {
		L.tail.next = newest
	}
	L.tail = newest
	L.size++
}
func (L *LinkedList) addFirst(val int) {
	newest := &Node{val, nil}
	if L.isEmpty() {
		L.head = newest
		L.tail = newest
	} else {
		newest.next = L.head
		L.head = newest
	}
	L.size++

}
func (L *LinkedList) addany(e, pos int) {
	newest := &Node{e, nil}
	p := L.head
	i := 1
	for i < pos-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next = newest
	L.size++

}
func (L LinkedList) search(key int) bool {
	for L.head != nil {
		if L.head.val == key {
			return true
		}
		L.head = L.head.next

	}
	return false
}
func (L *LinkedList) removeLast() int {
	if L.isEmpty() {
		fmt.Println("list is empty")
		return -1
	}
	i := 1
	p := L.head
	for i < L.size-1 {
		p = p.next
		i++
	}
	e := L.tail.val
	L.tail = p
	p.next = nil
	L.size--

	if L.size == 0 {
		L.tail = nil
	}
	return e
}
func (L *LinkedList) removeFirst() int {
	if L.isEmpty() {
		fmt.Println("list is empty")
		return -1
	}
	e := L.head.val
	L.head = L.head.next
	L.size--
	return e
}
func (L *LinkedList) removeany(pos int) {
	p := L.head
	i := 1
	for i < pos-1 {
		p = p.next
		i++
	}
	p.next = p.next.next
	L.size--

}

func main() {
	fmt.Println("goSingleLinkedList")
	n := LinkedList{}
	fmt.Println(n)
	fmt.Println(n.size)
	n.addLast(34)

	fmt.Println(n.size)

	n.display()
	n.addLast(20)

	fmt.Println(n.size)

	n.display()
	n.addLast(50)

	fmt.Println(n.size)
	n.display()
	n.addFirst(5)
	n.addFirst(1)
	fmt.Println(n.size)
	n.display()
	fmt.Println(n.search(-20))
	fmt.Println(n.removeLast())
	fmt.Println(n.size)
	n.display()
	fmt.Println(n.removeFirst())
	fmt.Println(n.size)
	n.display()
	n.addany(25, 2)
	fmt.Println(n.size)
	n.display()
	n.removeany(1)
	fmt.Println(n.size)
	n.display()
	fmt.Println(n.removeFirst())
	fmt.Println(n.removeFirst())
	fmt.Println(n.size)
	n.display()
	n.removeFirst()
	fmt.Println(n.size)
	n.display()

}
