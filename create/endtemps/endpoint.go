package endtemps

import (
	"fmt"
	"text/template"
)

var endpointTemplate = template.Must(
	template.New("endpoint").Parse(endpointTemplateText),
)

// CreateEndpoint builds a generic endpoint. It assumes that dirPath
// ends with the name of of file to be written.
func CreateEndpoint(fullpath, gopath, dirPath, packageName string) {
	attr := struct {
		PackageName string
		ProjectInGopath string
	}{
		PackageName: packageName,
		ProjectInGopath: gopath,
	}
	path := fmt.Sprintf("%s/%s", fullpath, dirPath)
	writeFile(endpointTemplate, path, attr)
	//endpointTemplate.Execute(os.Stdout, attr)
}

var endpointTemplateText = `package {{.PackageName}}

import (
	"net/http"

	"github.com/aodin/sol"

	"{{.ProjectInGopath}}/server/api"
)

type ExampleAPI struct {
	*api.Endpoint
	conn sol.Conn
}

func (c *ExampleAPI) List(w http.ResponseWriter, r *http.Request) {
	example := struct{
		ID int64
		Name string
	}{
		ID: 1,
		Name: "Invictus Example",
	}
	api.Respond(w, example)
	return
}


func (c *ExampleAPI) Options() api.Option {
	o := api.NewOption(c.API().Link(c.Name()), "example endpoint")
	return o
}

func Example(conn sol.Conn) *ExampleAPI {
	return &ExampleAPI{
		Endpoint: api.NewEndpoint("example"),
		conn: conn,
	}
}
`
