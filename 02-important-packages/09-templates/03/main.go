package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Python", 50},
		{"Java", 40},
		{"JavaScript", 100},
	})

	if err != nil {
		panic(err)
	}
}
