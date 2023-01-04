package main

import "fmt"

func main() {

	for i := 2; i < 100; i++ {

		count := 0

		for j := 2; j < i/2; j++ {

			if i%j == 0 {

				count += 1

				break

			}

		}

		if count == 0 {

			fmt.Printf("%d ", i)

		}

	}

}
