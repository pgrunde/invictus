package create

import (
	"fmt"
	"log"
	"os"
)

func createDbFolder(s, path string) {
	createFolder(s, path, "db")
}

func createMigrationsFolder(s, path string) {
	createFolder(s, path, "db/migrations")
}

func createServerFolder(s, path string) {
	createFolder(s, path, "server")
}

func createApiFolder(s, path string) {
	createFolder(s, path, "server/api")
}

func createParamsFolder(s, path string) {
	createFolder(s, path, "server/params")
}

func createFolder(s, path, folder string) {
	err := os.Mkdir(fmt.Sprintf("%s/%s/"+folder, path, s), 0744)
	if err != nil {
		log.Fatalf("cannot create a /%s folder at path %s%s: %s", folder, path, s, err)
	}
}
