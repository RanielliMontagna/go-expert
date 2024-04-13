package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := t.Execute(w, Courses{
			{"Go", 40},
			{"Python", 50},
			{"Java", 40},
			{"JavaScript", 100},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
