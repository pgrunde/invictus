package templates

import (
	"fmt"
	"text/template"
)

var responseTemplate = template.Must(
	template.New("response").Parse(responseTemplateText),
)

func CreateResponse(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/api/response.go", fullpath, projectName)
	writeFile(responseTemplate, path, nil)
}

const responseTemplateText = `package api

import (
	"net/http"
)

type Response struct {
	Error Error       ` + "`json:\"error,omitempty\"`" + `
	Data  interface{} ` + "`json:\"data\"`" + `
}

func (r Response) IsEmpty() bool {
	if r.Data == nil {
		return true
	}
	results, ok := r.Data.([]interface{})
	if !ok {
		return false
	}
	if len(results) == 0 {
		return true
	}
	return false
}

func Respond(w http.ResponseWriter, results interface{}) {
	var resp Response
	resp.Data = results
	Write(w, resp)
	return
}

func Empty() Response {
	return Response{
		Data: make([]interface{}, 0),
	}
}

func Unsupported(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Error: MetaError(404, "The method %s is disabled for %s", r.Method, r.URL.Path),
		Data:  make([]interface{}, 0),
	}
	Write(w, response)
	return
}
`
