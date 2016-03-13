package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var interfaceTemplate = template.Must(
	template.New("interface").Parse(interfaceTemplateText),
)

func CreateInterface(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/interface.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/interface.go file: %s", err)
	}
	defer file.Close()
	interfaceTemplate.Execute(file, nil)
}

const interfaceTemplateText = `package api

import (
	"net/http"
	"net/url"
)

type Rest interface {
	Name() string
	List(*Request) (Response, *Error)
	Post(*Request) (Response, *Error)
	Get(*Request) (Response, *Error)
	Patch(*Request) (Response, *Error)
	Delete(*Request) (Response, *Error)
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
