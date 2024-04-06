package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(5, 54, 24, 9, 10))
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
