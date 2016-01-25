package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const initTemplateText = `package db

import (
	"github.com/aodin/sol"
	"github.com/aodin/sol/types"
)

var (
	// Models
	// Example *sol.TableElem
)

func init() {
	// Example = sol.Table("images",
	// 	sol.Column("id", types.Integer().NotNull()),
	// 	sol.Column("name", types.Varchar().Limit(32).NotNull()),
	// 	sol.PrimaryKey("id"),
	// )
}
`

var initTemplate = template.Must(
	template.New("init").Parse(initTemplateText),
)

func CreateInit(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s", fullpath, projectName) + "/db/init.go")
	if err != nil {
		log.Fatalf("cannot create a init.go file: %s", err)
	}
	defer file.Close()
	initTemplate.Execute(file, nil)
}
