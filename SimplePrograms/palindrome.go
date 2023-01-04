package main

import "fmt"

func palindrome(m, n int) []int {
	result := []int{}
	for i := m; i <= n; i++ {
		temp := i
		rev := 0
		for a := i; a > 0; a = a / 10 {
			rev = rev*10 + a%10

		}
		if rev == temp {
			result = append(result, temp)

		}

	}
	return result

}

func main() {
	var m, n int
	fmt.Println("enter the first number and second number:")
	fmt.Scanf("%d %d", &m, &n)
	fmt.Println(palindrome(m, n))

}
