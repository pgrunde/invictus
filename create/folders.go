package create

import (
	"fmt"
	"log"
	"os"
)

func createDbFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/db", path, s), 0744)
	if err != nil {
		log.Fatalf("cannot create a /db folder at path %s%s: %s", path, s, err)
	}
}

func createMigrationsFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/db/migrations", path, s), 0744)
	if err != nil {
		log.Fatalf("cannot create a /db/migrations folder at path %s%s: %s", path, s, err)
	}
}

func createServerFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server", path, s), 0744)
	if err != nil {
		log.Fatalf("cannot create a /server folder at path %s%s: %s"+path+s, err)
	}
}

func createApiFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server/api", path, s), 0744)
	if err != nil {
		log.Fatalf("cannot create a /server/api folder at path %s%s: %s", path, s, err)
	}
}

func createParamsFolder(s, path string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/server/params", path, s), 0744)
	if err != nil {
		log.Fatalf("cannot create a /server/params folder at path %s%s: %s", path, s, err)
	}
}
