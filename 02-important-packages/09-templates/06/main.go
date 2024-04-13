package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	})
	t = template.Must(t.ParseFiles(templates...))

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
