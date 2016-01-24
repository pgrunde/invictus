package templates

import (
	"log"
	"os"
	"text/template"
)

const mainTemplateText = `package main

import (
	"log"
	"os"
	"path/filepath"

	"{{.ProjectInGopath}}/{{.ProjectName}}/server"

	"github.com/aodin/sol"
	_ "github.com/aodin/sol/postgres"
	"github.com/aodin/volta/config"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "img"
	app.Usage = "runs the image manipulation server"
	app.Version = server.VERSION
	app.Action = startServer
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

func startServer(c *cli.Context) {
	logF := c.String("log")
	file := c.String("config")
	// Set the log output - if no path given, use stdout
	if logF != "" {
		if err := os.MkdirAll(filepath.Dir(logF), 0776); err != nil {
			log.Panic(err)
		}
		l, err := os.OpenFile(logF, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer l.Close()
		log.SetOutput(l)
	}
	conf, conn := Connect(file)
	defer conn.Close()
	log.Panic(server.New(conf, conn).ListenAndServe())
}

func Connect(file string) (config.Config, sol.Conn) {
	// Parse the given configuration file
	conf, err := config.ParseFile(file)
	if err != nil {
		log.Panicf("{{.ProjectName}}: could not parse configuration: %s", err)
	}

	// Connect to the database
	conn, err := sol.Open(conf.Database.Driver, conf.Database.Credentials())
	if err != nil {
		log.Panicf("{{.ProjectName}}: could not connect to the database: %s", err)
	}
	return conf, conn
}
`

var mainTemplate = template.Must(
	template.New("main").Parse(mainTemplateText),
)

func CreateMain(projectName, currentGoPath string) {
	attr := struct {
		ProjectName string
		ProjectInGopath string
	}{
		ProjectName: projectName,
		ProjectInGopath: currentGoPath,
	}
	file, err := os.Create(projectName + "/main.go")
	if err != nil {
		log.Fatal("cannot create a main.go file")
	}
	defer file.Close()
	mainTemplate.Execute(file, attr)
}
