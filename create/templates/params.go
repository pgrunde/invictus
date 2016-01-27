package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const paramsTemplateText = `package params

type Param struct {
	Key, Value string
}

type Params []Param
`

var paramsTemplate = template.Must(
	template.New("params").Parse(paramsTemplateText),
)

func CreateParams(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/params/params.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/params/params.go file: %s", err)
	}
	defer file.Close()
	paramsTemplate.Execute(file, nil)
}
