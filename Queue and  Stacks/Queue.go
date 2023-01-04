package main

import "fmt"

type Queue []int

func (q *Queue) isempty() bool {
	return len(*q) == 0
}
func (q *Queue) enque(e int) {
	*q = append(*q, e)
}
func (q *Queue) deque() (int, bool) {
	if q.isempty() {
		return 0, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}
func (q *Queue) top() (int, bool) {
	if q.isempty() {
		return 0, false
	} else {
		d := *q
		fmt.Println(d[len(d)-1])
		return 0, true
	}
}

func main() {
	var n Queue
	fmt.Println(n)
	fmt.Println(n.isempty())
	n.enque(2)
	n.enque(3)
	n.enque(4)
	n.enque(5)
	n.enque(6)
	fmt.Println(n)
	n.deque()
	fmt.Println(n)
	n.top()
}
