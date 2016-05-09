package newtemps

import (
	"fmt"
	"text/template"
)

var initTemplate = template.Must(
	template.New("init").Parse(initTemplateText),
)

func CreateInit(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/db/init.go", fullpath, projectName)
	writeFile(initTemplate, path, nil)
}

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
