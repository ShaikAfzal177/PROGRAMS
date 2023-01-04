package main

import "fmt"

func insertionSort(items []int) {
	n := len(items)
	for i := 1; i < n; i++ {

		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1

		}

	}
}

func main() {
	items := []int{1, 2, 6, 8, 5, 4, 3, 7}
	fmt.Println(items)
	insertionSort(items)
	fmt.Println(items)

}
