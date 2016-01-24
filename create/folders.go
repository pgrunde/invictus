package create

import (
	"fmt"
	"os"
	"log"
)

func createServerFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server", path, s), 0744)
	if err != nil {
		log.Fatal("cannot create a /server folder at path "+path+s)
	}
}

func createApiFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server/api", path, s), 0744)
	if err != nil {
		log.Fatal("cannot create a /server/api folder at path "+path+s)
	}
}

func createParamsFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server/params", path, s), 0744)
	if err != nil {
		log.Fatal("cannot create a /server/params folder at path "+path+s)
	}
}

