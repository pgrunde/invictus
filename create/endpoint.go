package create

import (
	"os"
	"log"
	"fmt"

	"github.com/pgrunde/invictus/create/endtemps"
)

func NewEndpoint(projectName, name, dirPath string) error {
	fullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gopath := buildGoPath(projectName, fullPath, os.Getenv("GOPATH"))
	fmt.Println("gopath", gopath)
	fmt.Println("fullPath", fullPath)
	endtemps.CreateEndpoint(projectName,
													fullPath,
													gopath,
													fmt.Sprintf("%s/%s.go", dirPath, name),
													dirPath,
												)
	return nil
}
