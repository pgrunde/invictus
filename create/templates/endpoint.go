package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var endpointTemplate = template.Must(
	template.New("endpoint").Parse(endpointTemplateText),
)

func CreateEndpoint(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/endpoint.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/endpoint.go file: %s", err)
	}
	defer file.Close()
	endpointTemplate.Execute(file, nil)
}

var endpointTemplateText = `package api

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

func (c *Endpoint) List(r *Request) (Response, *Error) {
	return Unsupported(r)
}

func (c *Endpoint) Post(r *Request) (Response, *Error) {
	return Unsupported(r)
}

func (c *Endpoint) Get(r *Request) (Response, *Error) {
	return Unsupported(r)
}

func (c *Endpoint) Patch(r *Request) (Response, *Error) {
	return Unsupported(r)
}

func (c *Endpoint) Delete(r *Request) (Response, *Error) {
	return Unsupported(r)
}

func (c *Endpoint) Options() Option {
	return NewOption(c.api.Link(c.name), "")
}

func NewEndpoint(name string) *Endpoint {
	return &Endpoint{name: name}
}
`
