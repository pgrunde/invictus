package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var responseTemplate = template.Must(
	template.New("response").Parse(responseTemplateText),
)

func CreateResponse(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/response.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/response.go file: %s", err)
	}
	defer file.Close()
	responseTemplate.Execute(file, nil)
}

const responseTemplateText = `package api

import (
	"fmt"
	"net/http"
)

type Response struct {
	Error   *Error      ` + "`json:\"error,omitempty\"`" + `
	Results interface{} ` + "`json:\"results\"`" + `
}

func (r Response) IsEmpty() bool {
	if r.Results == nil {
		return true
	}
	results, ok := r.Results.([]interface{})
	if !ok {
		return false
	}
	if len(results) == 0 {
		return true
	}
	return false
}

func Respond(results interface{}) (r Response) {
	r.Results = results
	return
}

func Empty() Response {
	return Response{
		Results: make([]interface{}, 0),
	}
}

func Unsupported(r *Request) (Response, *Error) {
	msg := fmt.Sprintf(
		"The method %s is disabled for %s",
		r.Request.Method,
		r.Request.URL.Path,
	)
	return Empty(), MetaError(http.StatusMethodNotAllowed, msg)
}
`
