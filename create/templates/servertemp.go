package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

const serverTemplateText = `package server

import (
	"log"
	"net/http"

	"{{.ProjectInGopath}}/server/api"
	"{{.ProjectInGopath}}/server/api/v1"

	"github.com/aodin/sol"
	"github.com/aodin/volta/config"
)

const VERSION = "0.0.1"

type Server struct {
	Config config.Config
	Conn   sol.Conn
}

// ListenAndServe starts the server and serves forever
func (server *Server) ListenAndServe() error {
	log.Printf("server: serving on address %s\n", server.Config.Address())
	return http.ListenAndServe(server.Config.Address(), nil)
}

// New creates a new Server
func New(conf config.Config, conn sol.Conn) *Server {
	API := api.New(conf, conn).SetPrefix(v1.Prefix)

	API.Add(v1.Imgs(conn), "id")

	http.Handle("/", http.RedirectHandler(v1.Prefix, 302))
	http.Handle(v1.Prefix, API)
	http.HandleFunc("/favicon.ico", favicon)

	return &Server{Config: conf, Conn: conn}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	return
}
`

var serverTemplate = template.Must(
	template.New("server").Parse(serverTemplateText),
)

func CreateServerFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server", path, s), 0744)
	if err != nil {
		log.Fatal("cannot create a /server folder at path "+path+s)
	}
}

func CreateServer(s, fullpath, gopath string) {
	attr := struct {
		ProjectInGopath string
	}{
		ProjectInGopath: gopath,
	}
	file, err := os.Create(fmt.Sprintf("%s/%s", fullpath, s) + "/server/server.go")
	if err != nil {
		log.Fatal("cannot create a /server/server.go file")
	}
	defer file.Close()
	serverTemplate.Execute(file, attr)
}
