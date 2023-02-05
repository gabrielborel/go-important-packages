package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 60},
		{"Node.js", 70},
		{"Typescript", 90},
	})
	if err != nil {
		panic(err)
	}
}
