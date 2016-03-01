package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var requestTemplate = template.Must(
	template.New("request").Parse(requestTemplateText),
)

func CreateRequest(projectName, fullpath, gopath string) {
	attr := struct {
		ProjectInGopath string
	}{
		ProjectInGopath: gopath,
	}
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/request.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/request.go file: %s", err)
	}
	defer file.Close()
	requestTemplate.Execute(file, attr)
}

const requestTemplateText = `package api

import (
	"net/http"

	"{{.ProjectInGopath}}/server/params"
)

type Request struct {
	*http.Request
	Params params.Params
}

func NewRequest(r *http.Request, ps ...params.Param) *Request {
	return &Request{Request: r, Params: ps}
}
`
