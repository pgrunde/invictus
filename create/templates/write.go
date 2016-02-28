package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var writeTemplate = template.Must(
	template.New("write").Parse(writeTemplateText),
)

func CreateWrite(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/write.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/errors.go file: %s", err)
	}
	defer file.Close()
	writeTemplate.Execute(file, nil)
}

const writeTemplateText = `package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	if response.Error != nil {
		w.WriteHeader(response.Error.Code)
	} else if response.IsEmpty() {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Panicf("api: could not JSON encode response: %s", err)
	}
	w.Write(b)
}
`
