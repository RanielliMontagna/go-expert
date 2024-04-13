package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Hello, World!"))
	// size, err := f.WriteString("Hello, World!")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created with size: %d bytes\n", size)

	f.Close()

	// Read
	file, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File content: %s\n", string(file))

	arquive2, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquive2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")

	if err != nil {
		panic(err)
	}
}
