package newtemps

import (
	"fmt"
	"text/template"
)

var settingsTemplate = template.Must(
	template.New("settings").Parse(settingsTemplateText),
)

func CreateSettings(projectName, dbname string) {
	if dbname == "" {
		dbname = projectName
	}
	attr := struct {
		ProjectName string
		DbName      string
	}{
		ProjectName: projectName,
		DbName:      dbname,
	}
	path := fmt.Sprintf("%s/settings.example.json", projectName)
	writeFile(settingsTemplate, path, attr)

	path = fmt.Sprintf("%s/settings.json", projectName)
	writeFile(settingsTemplate, path, attr)
}

const settingsTemplateText = `{
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
