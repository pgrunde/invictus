package newtemps

import (
	"fmt"
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
	path := fmt.Sprintf("%s/%s/server/api/request.go", fullpath, projectName)
	writeFile(requestTemplate, path, attr)
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
