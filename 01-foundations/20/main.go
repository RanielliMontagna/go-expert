package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"one", "two", "three"}
	for i, number := range numbers {
		println(i, number)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Infinite loop")
	}
}
