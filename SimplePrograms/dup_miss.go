package main

import (
	"fmt"

	"sort"
)

func main() {

	arr := []int{28, 31, 31, 29, 30}

	sort.Ints(arr)

	var repeat, miss int

	for i := 0; i < len(arr); i++ {

		if i == len(arr)-1 {

			break

		}

		if arr[i] == arr[i+1] {

			repeat = arr[i]

		}

		if arr[i+1]-arr[i] != 1 {

			miss = arr[i] + 1

		}

	}

	fmt.Println(repeat, miss)

}
