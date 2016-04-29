package templates

import (
	"fmt"
	"text/template"
)

var v1template = template.Must(
	template.New("v1").Parse(v1templateText),
)

func CreateV1(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/v1/v1.go", fullpath, projectName)
	writeFile(v1template, path, nil)
}

const v1templateText = `package v1

var Prefix = "/api/v1/"
`
