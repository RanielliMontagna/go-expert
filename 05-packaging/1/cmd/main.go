package main

import (
	"fmt"
	"goexpert/packaging/1/math"
)

func main() {
	fmt.Println("Hello, World!")

	// Create a new Math struct and call the Add method
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
