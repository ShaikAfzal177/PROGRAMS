package main

import "fmt"

func bubbleSort(items []int) {
	n := len(items)
	for j := n - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if items[i] > items[i+1] {
				temp := items[i]
				items[i] = items[i+1]
				items[i+1] = temp
			}
		}
	}

}

func main() {
	items := []int{1, 2, 6, 8, 5, 4, 3, 7}
	fmt.Println(items)
	bubbleSort(items)
	fmt.Println(items)

}
