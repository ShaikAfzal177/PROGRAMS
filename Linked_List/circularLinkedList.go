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
	if L.isEmpty() {
		return
	}
	i := 0

	for i < L.size {
		fmt.Printf("%d ->", p.val)
		p = p.next
		i += 1
	}
	fmt.Println("")
}
func (L *LinkedList) addLast(val int) {
	newest := &Node{val, nil}
	if L.isEmpty() {

		L.head = newest
		L.head.next = newest

	} else {
		newest.next = L.tail.next
		L.tail.next = newest
	}
	L.tail = newest
	L.size++
}
func (L *LinkedList) addFirst(val int) {
	newest := &Node{val, nil}
	if L.isEmpty() {
		L.head = newest
		L.head.next = newest
		L.tail = newest
	} else {
		newest.next = L.head
		L.tail.next = newest
		L.head = newest
	}
	L.size++

}
func (L *LinkedList) addany(e, pos int) {
	newest := &Node{e, nil}
	if L.isEmpty() {
		L.head = newest
		L.tail = newest
	}
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
func (L *LinkedList) search(key int) bool {
	p := L.head
	i := 0
	for i < L.size {
		if p.val == key {
			return true

		}
		p = p.next
		i++

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
	p.next = L.head
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
	L.tail.next = L.head
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
	fmt.Println("gocircularLinkedList")
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
