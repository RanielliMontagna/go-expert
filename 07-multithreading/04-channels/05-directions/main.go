package main

import "fmt"

func receice(name string, ch chan<- string) {
	ch <- name
}

func read(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string, 1)

	go receice("Hello", ch)
	read(ch)
}
