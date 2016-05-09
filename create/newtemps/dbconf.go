package newtemps

import (
	"fmt"
	"text/template"
)

var dbconfTemplate = template.Must(
	template.New("dbconf").Parse(dbconfTemplateText),
)

func CreateDbConf(projectName, fullpath, dbname, pgusername, pgpassword string) {
	attr := struct {
		ProjectName      string
		DbName           string
		PostgresUser     string
		PostgresPassword string
	}{
		ProjectName:      projectName,
		DbName:           dbname,
		PostgresUser:     "postgres",
		PostgresPassword: "",
	}
	path := fmt.Sprintf("%s/%s/db/dbconf.example.yml", fullpath, projectName)
	writeFile(dbconfTemplate, path, attr)

	if dbname == "" {
		attr.DbName = projectName
	}
	if pgusername != "" {
		attr.PostgresUser = pgusername
	}
	if pgpassword != "" {
		attr.PostgresPassword = pgpassword
	}
	path = fmt.Sprintf("%s/%s/db/dbconf.yml", fullpath, projectName)
	writeFile(dbconfTemplate, path, attr)
}

const dbconfTemplateText = `development:
    driver: postgres
    open: host=localhost port=5432 dbname={{.DbName}}_dev user={{.PostgresUser}} password={{.PostgresPassword}} sslmode=disable

testing:
    driver: postgres
    open: host=localhost port=5432 dbname={{.DbName}}_test user={{.PostgresUser}} password={{.PostgresPassword}} sslmode=disable
`
