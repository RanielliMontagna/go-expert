package main

func main() {
	// Memory -> Address -> Value
	a := 10
	var pointer *int = &a

	*pointer = 20
	println(a)

	b := &a
	*b = 30

	println(a)
}
