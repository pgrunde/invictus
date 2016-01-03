package create

import (
	"log"
	"os"
	"text/template"
)

const exampleSettingsTemplateText = `{
    "domain": "localhost",
    "proxy_domain": "",
    "port": 3001,
    "proxy_port": 3000,
    "smtp": {
        "port": 587,
        "user": "",
        "password": "",
        "host": "smtp.sendgrid.net",
        "from": "test_from@example.com",
        "alias": "{{.ProjectName}}"
    },
    "database": {
        "driver": "postgres",
        "host": "localhost",
        "port": 5432,
        "name": "{{.DbName}}_dev",
        "user": "postgres",
        "password": "",
        "ssl_mode": "disable"
    },
    "cookie": {
        "age": 7776000000000000,
        "name": "{{.ProjectName}}id",
        "path": "/"
    },
    "metadata": {
    }
}
`

var exampleSettingsTemplate = template.Must(
	template.New("examplesettings").Parse(exampleSettingsTemplateText),
)

func createExampleSettings(s, dbname string) {
	attr := struct {
		ProjectName string
		DbName      string
	}{
		ProjectName: s,
		DbName:      dbname,
	}
	file, err := os.Create(s + "/settings.example.json")
	if err != nil {
		log.Fatal("cannot create a settings.example.json file")
	}
	defer file.Close()
	exampleSettingsTemplate.Execute(file, attr)
}
