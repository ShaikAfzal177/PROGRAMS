package main

import "fmt"

func nonrepeating(s string) {
	var arr []int
	for _, i := range s {
		count := 0
		for _, j := range s {
			if i == j {
				count++
			}
		}
		arr = append(arr, count)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			fmt.Println(string(s[i]))
			break
		}
	}

}
func main() {
	var s string
	fmt.Println("enter a word: ")
	fmt.Scanf("%s", &s)
	nonrepeating(s)

}
