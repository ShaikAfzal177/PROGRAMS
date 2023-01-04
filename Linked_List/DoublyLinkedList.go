package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
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
	newest := &Node{val, nil, nil}
	if L.isEmpty() {

		L.head = newest

	} else {

		L.tail.next = newest
		newest.prev = L.tail

	}
	L.tail = newest
	L.size++
}
func (L *LinkedList) addFirst(val int) {
	newest := &Node{val, nil, nil}
	if L.isEmpty() {
		L.head = newest

		L.tail = newest
	} else {
		newest.next = L.head
		L.head.prev = newest
		L.head = newest
	}
	L.size++

}
func (L *LinkedList) addany(e, pos int) {
	newest := &Node{e, nil, nil}
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
	p.next.prev = newest
	newest.prev = p
	p.next = newest
	L.size++

}
func (L *LinkedList) search(key int) bool {
	p := L.head
	i := 0
	for i < L.size {
		if L.head.val == key {
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

	e := L.tail.val
	L.tail = L.tail.prev
	L.tail.next = nil
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
	if L.head == nil {
		L.size--
		L.tail = nil
		return e
	}
	L.head.prev = nil
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
	p.next.prev = p
	L.size--

}

func main() {
	fmt.Println("goDoublyLinkedList")
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
