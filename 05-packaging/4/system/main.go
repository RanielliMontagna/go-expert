package main

import (
	"fmt"
	"goexpert/packaging/3/math"

	"github.com/google/uuid"
)

func main() {
	m := math.Math{}
	fmt.Println(m.Add())

	print(uuid.New().String())
}
