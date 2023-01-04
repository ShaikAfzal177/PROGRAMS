package main

import "fmt"

func linearSearch(items []int, num int) bool {
	for _, item := range items {
		if item == num {
			return true
		}
	}
	return false
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if linearSearch(items, 5) {
		fmt.Println("Item found")
	} else {
		fmt.Println("Item not found")
	}
}
