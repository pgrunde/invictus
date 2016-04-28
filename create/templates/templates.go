package templates

import (
	"log"
	"os"
	"text/template"
)

func writeFile(temp *template.Template, path string, attr interface{}) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Cannot create a %s file: %s", path, err)
	}
	defer file.Close()
	temp.Execute(file, attr)
}
