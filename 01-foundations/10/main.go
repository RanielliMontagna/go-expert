package main

import "fmt"

func main() {

	total := func() int {
		return sum(5, 54, 24, 9, 10) * 2
	}()

	fmt.Println(total)

}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
