package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Courses{
			{"Go", 40},
			{"Java", 60},
			{"Node.js", 70},
			{"Typescript", 90},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
