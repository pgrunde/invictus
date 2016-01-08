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
func GenerateNew(s, dbname string) {
	templates.CreateMain(s)
	templates.CreateSettings(s, dbname)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	templates.CreateServerFolder(s, currentPath)
	templates.CreateServer(s, currentPath)
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
