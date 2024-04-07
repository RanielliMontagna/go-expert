package main

func sum(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	number1 := 10
	number2 := 20
	sum(&number1, &number2)

	println(number1)
	println(number2)
}
