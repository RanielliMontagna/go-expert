package main

import (
	"18/math"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := math.Sum(10, 20)

	fmt.Printf("Sum: %d\n", s)

	fmt.Println(uuid.New())
}
