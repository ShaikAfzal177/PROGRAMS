package main

import "fmt"

func BinarySearch(items []int, value int) {
	min := 0
	max := len(items)
	mid := 0
	for min < max {
		mid = (min + max) / 2
		if value == items[mid] {
			break
		} else if value > items[mid] {
			min = mid + 1

		} else {
			max = mid - 1
		}
	}
	if min >= len(items) {
		fmt.Println("item not found")
	} else {
		fmt.Println("item found at =", mid+1)
	}

}
func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	BinarySearch(items, 5)
}
