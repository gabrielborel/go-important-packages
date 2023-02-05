package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name       string
	TotalHours int
}

func main() {
	course := Course{Name: "Golang", TotalHours: 40}
	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Total Hours: {{.TotalHours}}."))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
