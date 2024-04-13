package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go", 40}
	tmp := template.New("TemplateCourse")
	tmp, _ = tmp.Parse("Course: {{.Name}}\nWorkload: {{.Workload}} hours")

	err := tmp.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
