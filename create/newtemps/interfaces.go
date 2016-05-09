package newtemps

import (
	"fmt"
	"text/template"
)

var interfaceTemplate = template.Must(
	template.New("interface").Parse(interfaceTemplateText),
)

func CreateInterface(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/api/interface.go", fullpath, projectName)
	writeFile(interfaceTemplate, path, nil)
}

const interfaceTemplateText = `package api

import (
	"net/http"
	"net/url"
)

// Rest is the common interface for REST-ful resources
type Rest interface {
	Name() string
	List(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Options() Option
	SetAPI(*API)
	API() *API
}

type APIer interface {
	Prefix() string
	SetPrefix(string) APIer
	Link(string, ...url.Values) Link
	RootURL() string
	Add(Rest, ...string) error
	Handle(http.ResponseWriter, *Request) error
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Handle interface {
	Rest
}
`
