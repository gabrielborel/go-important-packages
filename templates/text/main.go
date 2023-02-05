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
	tmp := template.New("CourseTemplate")
	tmp, err := tmp.Parse("Course: {{.Name}} - Total Hours: {{.TotalHours}}.")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
