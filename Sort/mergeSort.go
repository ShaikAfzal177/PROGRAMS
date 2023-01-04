package main

import "fmt"

func mergesort(a []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergesort(a, left, mid)
		mergesort(a, mid+1, right)
		merge(a, left, mid, right)
	}
}
func merge(a []int, left, mid, right int) {
	i := left
	j := mid + 1
	k := left
	B := make([]int, right+1)
	for i <= mid && j <= right {

		if a[i] < a[j] {
			B[k] = a[i]
			i += 1
		} else {
			B[k] = a[j]
			j += 1
		}
		k = k + 1
	}
	for i <= mid {
		B[k] = a[i]
		i += 1
		k += 1
	}
	for j <= right {
		B[k] = a[j]
		j += 1
		k += 1
	}
	for x := left; x <= right; x++ {

		a[x] = B[x]
	}
}

func main() {
	items := []int{1, 2, 6, 8, 5, 4, 3, 7}
	fmt.Println(items)
	mergesort(items, 0, len(items)-1)
	fmt.Println(items)

}
