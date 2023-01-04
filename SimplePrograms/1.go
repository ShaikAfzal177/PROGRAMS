package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(2 + 5)
	a := 5
	p := &a
	fmt.Println(*p)
}
