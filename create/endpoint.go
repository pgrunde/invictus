package create

import (
	"os"
	"log"
	"fmt"

	"github.com/pgrunde/invictus/create/endtemps"
)

func NewEndpoint(projectName, name, dirPath string) error {
	currentFullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gopath := buildGoPath(projectName, currentFullPath, os.Getenv("GOPATH"))
	endtemps.CreateEndpoint(projectName,
													currentFullPath,
													gopath,
													fmt.Sprintf("%s/%s.go", dirPath, name),
													name,
												)
	return nil
}
