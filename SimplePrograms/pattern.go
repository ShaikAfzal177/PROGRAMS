package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Println("enter the number: ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		space := strings.Repeat(" ", (n - (i + 1)))
		star := strings.Repeat("*", (i + 1))
		fmt.Println(space + star)
	}
	for i := 0; i < n; i++ {
		space := strings.Repeat(" ", (n - 1))
		star := strings.Repeat("*", (n - i))
		fmt.Println(space + star)
	}
}
