package create

import (
	"log"
	"os"
	"text/template"
)

const mainTemplateText = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("{{.ProjectName}}")
}
`

var mainTemplate = template.Must(
	template.New("main").Parse(mainTemplateText),
)

func createMain(s string) {
	attr := struct {
		ProjectName string
	}{
		ProjectName: s,
	}
	file, err := os.Create(s + "/main.go")
	if err != nil {
		log.Fatal("cannot create a main.go file")
	}
	defer file.Close()
	mainTemplate.Execute(file, attr)
}
