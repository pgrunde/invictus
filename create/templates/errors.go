package templates

import (
	"fmt"
	"text/template"
)

var errorsTemplate = template.Must(
	template.New("errors").Parse(errorsTemplateText),
)

func CreateErrors(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/api/errors.go", fullpath, projectName)
	writeFile(errorsTemplate, path, nil)
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
