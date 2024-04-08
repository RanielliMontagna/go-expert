package main

import "fmt"

func main() {
	var name interface{} = "Ranni Montagna"
	println(name.(string))

	res, ok := name.(int)

	fmt.Printf("res: %v, ok: %v\n", res, ok)
}
