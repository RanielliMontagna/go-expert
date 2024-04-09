package main

func main() {
	a := 1
	b := 2
	c := 3

	if a < b && c > a {
		println("Both conditions are true")
	}

	if a < b {
		println("a is less than b")
	} else {
		println("a is not less than b")
	}

	switch a {
	case 1:
		println("a is 1")
	case 2:
		println("a is 2")
	default:
		println("a is not 1 or 2")
	}

}
