package main

import "fmt"

func selectionSort(items []int) {
	n := len(items)
	for i := 0; i < n-1; i++ {
		position := i
		for j := i + 1; j < n; j++ {
			if items[j] < items[position] {
				position = j
			}
		}
		temp := items[i]
		items[i] = items[position]
		items[position] = temp

	}

}

func main() {
	items := []int{1, 2, 6, 8, 5, 4, 3, 7}
	fmt.Println(items)
	selectionSort(items)
	fmt.Println(items)
}
