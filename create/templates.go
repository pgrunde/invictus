package create

import (
	"text/template"
)

const mainTemplateText = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("{{.Arg}}")
}
`

var mainTemplate = template.Must(
	template.New("main").Parse(mainTemplateText),
)
