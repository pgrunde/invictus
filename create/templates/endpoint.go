package templates

import (
	"fmt"
	"text/template"
)

var endpointTemplate = template.Must(
	template.New("endpoint").Parse(endpointTemplateText),
)

func CreateEndpoint(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/api/endpoint.go", fullpath, projectName)
	writeFile(endpointTemplate, path, nil)
}

var endpointTemplateText = `package api

import (
	"net/http"
)

type Endpoint struct {
	name string
	api  *API
}

func (c Endpoint) API() *API {
	return c.api
}

func (c Endpoint) Name() string {
	return c.name
}

func (c *Endpoint) SetAPI(api *API) {
	c.api = api
}

func (c *Endpoint) List(w http.ResponseWriter, r *http.Request) {
	Unsupported(w, r)
}

func (c *Endpoint) Post(w http.ResponseWriter, r *http.Request) {
	Unsupported(w, r)
}

func (c *Endpoint) Get(w http.ResponseWriter, r *http.Request) {
	Unsupported(w, r)
}

func (c *Endpoint) Patch(w http.ResponseWriter, r *http.Request) {
	Unsupported(w, r)
}

func (c *Endpoint) Delete(w http.ResponseWriter, r *http.Request) {
	Unsupported(w, r)
}

func (c *Endpoint) Options() Option {
	return NewOption(c.api.Link(c.name), "")
}

func NewEndpoint(name string) *Endpoint {
	return &Endpoint{name: name}
}
`
