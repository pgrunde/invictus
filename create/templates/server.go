package templates

import (
	"fmt"
	"text/template"
)

var serverTemplate = template.Must(
	template.New("server").Parse(serverTemplateText),
)

func CreateServer(projectName, fullpath, gopath string) {
	attr := struct {
		ProjectInGopath string
	}{
		ProjectInGopath: gopath,
	}
	path := fmt.Sprintf("%s/%s/server/server.go", fullpath, projectName)
	writeFile(serverTemplate, path, attr)
}

const serverTemplateText = `package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"{{.ProjectInGopath}}/v1"
	"{{.ProjectInGopath}}/server/api"

	"github.com/aodin/sol"
	_ "github.com/aodin/sol/postgres"
	"github.com/aodin/volta/config"
	"github.com/codegangsta/cli"
)

const VERSION = "0.0.1"

type Server struct {
	Config config.Config
	Conn   sol.Conn
}

// New creates a new Server
func New(conf config.Config, conn sol.Conn) *Server {
	API := api.New(conf, conn).SetPrefix(v1.Prefix)

	// Controller Resources
	// Example of how to add a controller resource from package v1
	// API.Add(v1.Example(conn), "id")
	http.Handle(v1.Prefix, API)

	// Redirects root requests to api V1
	http.Handle("/", http.RedirectHandler(v1.Prefix, 302))

	http.HandleFunc("/favicon.ico", favicon)
	return &Server{Config: conf, Conn: conn}
}

// ListenAndServe starts the server and serves forever
func (server *Server) ListenAndServe() error {
	log.Printf("server: serving on address %s\n", server.Config.Address())
	return http.ListenAndServe(server.Config.Address(), nil)
}

func Start(c *cli.Context) {
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
	log.Panic(New(conf, conn).ListenAndServe())
}

func Connect(file string) (config.Config, sol.Conn) {
	// Parse the given configuration file
	conf, err := config.ParseFile(file)
	if err != nil {
		log.Panicf("img: could not parse configuration: %s", err)
	}

	// Connect to the database
	conn, err := sol.Open(conf.Database.Driver, conf.Database.Credentials())
	if err != nil {
		log.Panicf("img: could not connect to the database: %s", err)
	}
	return conf, conn
}

func favicon(w http.ResponseWriter, r *http.Request) {
	return
}
`
