package templates

import (
	"fmt"
	"text/template"
)

var mainTemplate = template.Must(
	template.New("main").Parse(mainTemplateText),
)

func CreateMain(projectName, currentGoPath string) {
	attr := struct {
		ProjectName     string
		ProjectInGopath string
	}{
		ProjectName:     projectName,
		ProjectInGopath: currentGoPath,
	}
	path := fmt.Sprintf("%s/main.go", projectName)
	writeFile(mainTemplate, path, attr)
}

const mainTemplateText = `package main

import (
	"os"

	"{{.ProjectInGopath}}/server"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "img"
	app.Usage = "runs the image manipulation server"
	app.Version = server.VERSION
	app.Action = server.Start
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log, l",
			Value: "",
			Usage: "Sets the log output file path",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "./settings.json",
			Usage: "Sets the configuration file",
		},
	}
	app.Run(os.Args)
}
`
