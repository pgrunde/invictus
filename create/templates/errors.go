package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var errorsTemplate = template.Must(
	template.New("errors").Parse(errorsTemplateText),
)

func CreateErrors(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/errors.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/errors.go file: %s", err)
	}
	defer file.Close()
	errorsTemplate.Execute(file, nil)
}

const errorsTemplateText = `package api

import "fmt"

type Error struct {
	Code   int              ` + "`json:\"-\"`" + `
	Meta   []string         ` + "`json:\"meta\"`" + `
	Fields map[string]error ` + "`json:\"fields\"`" + `
}

func MetaError(code int, msg string, args ...interface{}) *Error {
	return &Error{
		Code:   code,
		Meta:   []string{fmt.Sprintf(msg, args...)},
		Fields: make(map[string]error),
	}
}

func FieldErrors(code int, errorSet map[string]error) *Error {
	return &Error{
		Code:   code,
		Meta:   []string{},
		Fields: errorSet,
	}
}

func (e *Error) Exists() bool {
	return len(e.Fields) != 0 || len(e.Meta) != 0
}
`
