package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var optionMethodTemplate = template.Must(
	template.New("optionMethod").Parse(optionMethodTemplateText),
)

func CreateOptionMethod(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/optionmethod.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/optionmethod.go file: %s", err)
	}
	defer file.Close()
	optionMethodTemplate.Execute(file, nil)
}

const optionMethodTemplateText = `package api

import (
	"log"
)

type OptionMethod struct {
	Example
	Supported   bool                   `+"`json:\"supported\"`"+`
	Description string                 `+"`json:\"description,omitempty\"`"+`
	Parameters  map[string]OptionParam `+"`json:\"parameters,omitempty\"`"+`
}

func (o *OptionMethod) AddParam(name, dataType, description string) {
	if _, exists := o.Parameters[name]; exists {
		log.Panic("Param with name %s already exists")
	}
	var example string
	if dataType == "bool" {
		example = "true"
	} else if dataType == "int" {
		example = "42"
	} else {
		example = "argument"
	}
	o.Parameters[name] = OptionParam{
		Type:        dataType,
		Description: description,
		ExampleURL:  o.ExampleURL + Link("?"+name+"="+example),
	}
}

type Example struct {
	ExampleURL      Link        `+"`json:\"example_url,omitempty\"`"+`
	ExampleBody     interface{} `+"`json:\"example_body,omitempty\"`"+`
	ExampleResponse interface{} `+"`json:\"example_response,omitempty\"`"+`
}

func (e *Example) SetBody(body interface{}) {
	e.ExampleBody = body
	return
}

func (e *Example) SetResponse(resp interface{}) {
	e.ExampleResponse = resp
	return
}

type OptionParam struct {
	Type        string `+"`json:\"type\"`"+` // indicates boolean, integer, string, etc
	Description string `+"`json:\"description\"`"+`
	ExampleURL  Link   `+"`json:\"example_url,omitempty\"`"+`
}
`
