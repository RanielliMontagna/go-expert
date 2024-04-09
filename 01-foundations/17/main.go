package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func sum[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{
		"Ranni":  1000,
		"Priya":  2000,
		"Shivam": 3000,
	}

	m2 := map[string]float64{
		"Ranni":  1000.0,
		"Priya":  2000.0,
		"Shivam": 3000.0,
	}

	m3 := map[string]MyNumber{
		"Ranni":  1000,
		"Priya":  2000,
		"Shivam": 3000,
	}

	fmt.Printf("Sum of Integers: %d\n", sum(m))
	fmt.Printf("Sum of Floats: %.2f\n", sum(m2))
	fmt.Printf("Sum of MyNumber: %d\n", sum(m3))

	fmt.Printf("Comparing Integers: %t\n", compare(10, 10.0))
}
