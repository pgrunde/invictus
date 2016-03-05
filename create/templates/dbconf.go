package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const dbconfTemplateText = `development:
    driver: postgres
    open: host=localhost port=5432 dbname={{.DbName}}_dev user={{.PostgresUser}} password={{.PostgresPassword}} sslmode=disable

testing:
    driver: postgres
    open: host=localhost port=5432 dbname={{.DbName}}_test user={{.PostgresUser}} password={{.PostgresPassword}} sslmode=disable
`

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
	file, err := os.Create(fmt.Sprintf("%s/%s", fullpath, projectName) + "/db/dbconf.example.yml")
	if err != nil {
		log.Fatal("cannot create a dbconf.example.yml file")
	}
	defer file.Close()
	dbconfTemplate.Execute(file, attr)

	if dbname == "" {
		attr.DbName = projectName
	}
	if pgusername != "" {
		attr.PostgresUser = pgusername
	}
	if pgpassword != "" {
		attr.PostgresPassword = pgpassword
	}
	file, err = os.Create(fmt.Sprintf("%s/%s", fullpath, projectName) + "/db/dbconf.yml")
	if err != nil {
		log.Fatal("cannot create a dbconf.yml file")
	}
	defer file.Close()
	dbconfTemplate.Execute(file, attr)

}
