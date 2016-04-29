package templates

import (
	"fmt"
	"text/template"
)

var paramsTemplate = template.Must(
	template.New("params").Parse(paramsTemplateText),
)

func CreateParams(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/params/params.go", fullpath, projectName)
	writeFile(paramsTemplate, path, nil)
}

const paramsTemplateText = `package params

type Param struct {
	Key, Value string
}

type Params []Param
`
