package main

import "fmt"

func partition(a []int, low, high int) int {
	pivot := a[low]
	i := low + 1
	j := high
	for {
		for i <= j && a[i] <= pivot {
			i += 1
		}

		for i <= j && a[j] > pivot {
			j -= 1
		}
		if i <= j {
			a[i], a[j] = a[j], a[i]
		} else {
			break
		}
	}
	a[low], a[j] = a[j], a[low]
	return j

}
func quicksort(a []int, low, high int) {
	if low < high {
		pi := partition(a, low, high)
		quicksort(a, low, pi-1)
		quicksort(a, pi+1, high)

	}
}

func main() {
	items := []int{1, 2, 6, 8, 5, 4, 3, 7}
	fmt.Println(items)
	quicksort(items, 0, len(items)-1)
	fmt.Println(items)

}
