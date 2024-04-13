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

	t := template.Must(template.New("TemplateCourse").Parse("Course: {{.Name}}\nWorkload: {{.Workload}} hours"))

	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
