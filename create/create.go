package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pgrunde/invictus/create/templates"
)

var badWindowsNames = []string{"com1", "com2", "com3", "com4", "com5", "com6", "com7", "com8", "com9", "lpt1", "lpt2", "lpt3", "lpt4", "lpt5", "lpt6", "lpt7", "lpt8", "lpt9", "con", "nul", "prn"}

var badCharacters = []rune{'^', '/', '?', '<', '>', '\\', ':', '*', '|', '"', '.'}

func NewProject(s, dbname string) (err error) {
	s = strings.ToLower(s)
	if err = hasIllegalFilename(s); err != nil {
		return fmt.Errorf("Given project name %s has an illegal filename: %s", s, err)
	}
	err = os.Mkdir("."+string(filepath.Separator)+s, 0744)
	if os.IsExist(err) {
		return fmt.Errorf("A directory of that name already exists")
	}
	GenerateNew(s, dbname)
	return nil
}

// GenerateNew takes in a project name and creates a folder
// with enclosing files
func GenerateNew(projectName, dbname string) {
	currentFullPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentGoPath := buildGoPath(projectName, currentFullPath, os.Getenv("GOPATH"))

	templates.CreateMain(projectName, currentGoPath)
	templates.CreateSettings(projectName, dbname)

	createServerFolder(projectName, currentFullPath)
	templates.CreateServer(projectName, currentFullPath, currentGoPath)

	createApiFolder(projectName, currentFullPath)
	createParamsFolder(projectName, currentFullPath)
}

// bulidGoPath assumes that imports follow GOPATH + "/src"
func buildGoPath(s, w, g string) string {
	return w[len(g) + 5:] + "/" + s
}

func hasIllegalFilename(s string) error {
	if s == "" {
		return fmt.Errorf("No project name given. Please specify a project name")
	}
	if len(s) > 255 {
		return fmt.Errorf("Project name cannot exceed 255 characters")
	}
	for _, bwn := range badWindowsNames {
		if s == bwn {
			return fmt.Errorf("Given name contains a reserved Windows operating system file name")
		}
	}
	for _, r := range badCharacters {
		if strings.ContainsRune(s, r) {
			return fmt.Errorf("Found illegal character %s in filename", string(r))
		}
	}
	return nil
}
