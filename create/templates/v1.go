package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var v1template = template.Must(
	template.New("v1").Parse(v1templateText),
)

func CreateV1(s, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s", fullpath, s) + "/v1/v1.go")
	if err != nil {
		log.Fatal("cannot create a /v1/v1.go file")
	}
	defer file.Close()
	v1template.Execute(file, nil)
}

const v1templateText = `package v1

var Prefix = "/api/v1/"
`
