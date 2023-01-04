package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}
func (n *Node) Print() {
	if n.Left != nil {
		n.Left.Print()
	}
	fmt.Printf("%d ", n.Value)
	if n.Right != nil {
		n.Right.Print()
	}
}

func main() {
	tree := &Node{Value: 20}
	tree.Insert(15)
	tree.Insert(8)
	tree.Insert(17)
	tree.Insert(25)
	tree.Insert(24)
	tree.Print()
}
